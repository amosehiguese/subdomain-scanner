use brutesvc::config::Config;
use brutesvc::init_grpc_server;
use brutesvc::brute::BruteForceComponent;
use brutesvc::telemetry::init_telemetry;


#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // load env variables from .env file
    dotenvy::dotenv().expect("got an error trying to read .env");

    // get configuration and set up local config variales
    let config = Config::get();
    let addr = config.listen_addr;   

    println!("Starting...");

    // initialize telemetry
    init_telemetry(&config.telemetry);

    // instantiate bruteforce component with default values and start grpc server
    let brutesvc = BruteForceComponent::default();
    init_grpc_server(brutesvc, addr).await?;

    Ok(())
}