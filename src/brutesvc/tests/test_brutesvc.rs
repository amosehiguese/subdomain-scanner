use tonic::Request;
use brutesvc::brute;
use tokio::sync::mpsc;
use tonic::transport::Server;
use brutesvc::subdomain::{brute_service_client::BruteServiceClient, brute_service_server::BruteServiceServer};
use brutesvc::subdomain::BruteForceRequest;

#[tokio::test]
async fn test_brute_valid_target(){

    let addr = "0.0.0.0:3661".parse().unwrap();


    let test_brute_svc_component = brute::BruteForceComponent::default();

    // Start grpc server in a background task
    let (tx, mut rx) = mpsc::channel::<()>(1);
    tokio::spawn(async move {
        Server::builder()
            .add_service(BruteServiceServer::new(test_brute_svc_component))
            .serve_with_shutdown(addr, async {
                rx.recv().await;
            }).await.unwrap();
    });

    // wait for server to start
    tokio::time::sleep(std::time::Duration::from_millis(100)).await;

    // Create a grpc client
    let mut client = BruteServiceClient::connect(format!("http://{}", addr)).await.expect("failed to connect to server");

    struct Data {
        target: String,
        success: bool,
        expected_code: tonic::Code
    }

    let test_cases  = vec![
        Data{
            target: "vmrw.com".to_string(),
            success: true,
            expected_code: tonic::Code::Ok,
        },
        Data{
            target: "".to_string(),
            success: false,
            expected_code: tonic::Code::InvalidArgument,
        }

    ];

    // Run test requests
    for Data { target, success, expected_code } in test_cases {

        let request = Request::new(BruteForceRequest{
            target: target.clone(),
        });

        let response = client.get_subdomains_by_brute_force(request).await;
        if success {
            assert!(response.is_ok(), "expected a successful response for target: {}", target);
            if let Ok(e) = response {
                assert!(e.into_inner().subdomains.len() > 1)
            }

        } else {
            assert!(response.is_err(), "expected an error for target: {}", target);
            if let Err(e) = response {
                assert_eq!(e.code(), expected_code);
            }
        }
    }

    // shut down server
    tx.send(()).await.unwrap();
}

