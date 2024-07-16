pub mod telemetry;
pub mod portscan;
pub mod config;
pub mod consts;

use anyhow::Ok;
use tonic::transport::Server;
use std::net::SocketAddr;
use subdomain::port_scan_service_server::PortScanServiceServer;


pub mod subdomain {
    tonic::include_proto!("subdomain");
}

pub async fn init_grpc_server(portscan: portscan::PortScanComponent, addr: SocketAddr) -> Result<(), anyhow::Error> {
    let (_, health_service) = tonic_health::server::health_reporter();
    tracing_log::log::info!("Starting grpc server at {:?}...", addr);
    Server::builder()
        .add_service(health_service)
        .add_service(PortScanServiceServer::new(portscan))
        .serve(addr)
        .await?;
    Ok(())
}
