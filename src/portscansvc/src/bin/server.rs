use portscansvc::{
    portscan::PortScanComponent,
    init_grpc_server,
    telemetry::init_telemetry,
    config::Config,
};

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

        // instantiate portscancomponent with default values and start grpc server
        let portscan = PortScanComponent::default();
        init_grpc_server(portscan, addr).await?;

        Ok(())
}