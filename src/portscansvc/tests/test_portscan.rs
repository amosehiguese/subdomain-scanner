use portscansvc::subdomain::api::portscan::v1::port_scan_service_client::PortScanServiceClient;
use portscansvc::subdomain::api::portscan::v1::PortScanRequest;
use tokio::sync::mpsc;
use std::net::SocketAddr;
use tonic::transport::Server;
use portscansvc::portscan::PortScanComponent;
use portscansvc::subdomain::api::portscan::v1::port_scan_service_server::PortScanServiceServer;

#[tokio::test]
async fn test_port_scan() {
    let addr: SocketAddr = "[::1]:3668".parse().unwrap();
    let test_portscan_svc_component = PortScanComponent::default();

    // Start grpc server in a background task
    let (tx, mut rx) = mpsc::channel::<()>(1);
    tokio::spawn(async move {
        Server::builder()
            .add_service(PortScanServiceServer::new(test_portscan_svc_component))
            .serve_with_shutdown(addr, async {
                rx.recv().await;
            }).await.unwrap();
    });

    // wait for server to start
    tokio::time::sleep(std::time::Duration::from_millis(100)).await;

    // Create a grpc client
    let mut client = PortScanServiceClient::connect(format!("http://{}", addr)).await.expect("failed to connect to server");

    let outbound = async_stream::stream! {

        for _i in 1..3 {
            let req = PortScanRequest{
                host: "vmrw.com".to_string(),
            };

            yield req;
        }
    };

    let response = client.scan_for_open_ports(outbound).await.unwrap();
    let mut inbound = response.into_inner();

    let mut got  = vec![];
    while let Some(value) = inbound.message().await.unwrap() {
        got.push(value);
    }

    assert!(got.len() > 1);
    // shut down server
    tx.send(()).await.unwrap();
}