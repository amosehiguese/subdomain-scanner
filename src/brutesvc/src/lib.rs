pub mod brute;
pub mod config;
pub mod errors;
pub mod telemetry;



use anyhow::Ok;
use tonic::transport::Server;
use std::net::SocketAddr;
use subdomain::brute_service_server::BruteServiceServer;

pub mod subdomain {
    tonic::include_proto!("subdomain");
}

pub async fn init_grpc_server(brutesvc: brute::BruteForceComponent, addr: SocketAddr) -> Result<(), anyhow::Error> {
    let (_, health_service) = tonic_health::server::health_reporter();
    tracing_log::log::info!("Starting grpc server at {:?}...", addr);
    Server::builder()
        .add_service(health_service)
        .add_service(BruteServiceServer::new(brutesvc))
        .serve(addr)
        .await?;
    Ok(())
}
