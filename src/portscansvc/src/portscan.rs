use crate::consts;
use tracing_log::log;
use tokio::{net::TcpStream, sync::mpsc};
use futures::stream::StreamExt;
use tonic::{Request, Response, Status};
use std::{net::{ToSocketAddrs, SocketAddr}};
use crate::subdomain::{port_scan_service_server::PortScanService, Subdomain, Port, PortScanRequest, PortScanResponse};


#[derive(Debug, Default)]
pub struct PortScanComponent {}

#[tonic::async_trait]
impl PortScanService for PortScanComponent {

    async fn scan_for_open_ports(
        &self,
        request: Request<PortScanRequest>,
    )-> Result<Response<PortScanResponse>, Status> {
        log::info!("Scanning for open ports...");

        let buffer = 100000;
        let req = request.into_inner();
        let mut output: Vec<Subdomain> = Vec::new();

        for subdomain in req.hosts.iter() {
            let address = format!("{}:1024", subdomain)
                .to_socket_addrs()
                .ok()
                .and_then(|mut addrs| addrs.next());

            if address.is_none() {
                let subd = Subdomain { domain: subdomain.to_string(), ports: vec![] };
                output.push(subd);
                continue
            }

            let address = address.unwrap();

            let (input_tx, input_rx) = mpsc::channel::<u32>(buffer);
            let (output_tx, output_rx) = mpsc::channel::<Port>(buffer);

            tokio::spawn(async move{
                for port in consts::PORTS_LIST{
                    let _ = input_tx.send(*port).await;
                }
                drop(input_tx);
            });

            let input_receiver_stream = tokio_stream::wrappers::ReceiverStream::new(input_rx);

            input_receiver_stream
                .for_each_concurrent(buffer, |port| {
                    let output_tx = output_tx.clone();
                    let address = address.clone();

                    async move {
                        let port: Port = scan_port(address, port).await;
                        if port.conn_open {
                            let _ = output_tx.send(port).await;
                        }
                    }
                })
                .await;
            drop(output_tx);

            let output_receiver_stream = tokio_stream::wrappers::ReceiverStream::new(output_rx);
            let ports: Vec<Port> = output_receiver_stream.collect().await;

            log::info!("scan completed for {:?}", subdomain);
            let subd = Subdomain { domain: subdomain.to_string(), ports: ports};

            output.push(subd);
        }

        Ok(Response::new(PortScanResponse{subdomains: output}))
    }
}

async fn scan_port(mut addr: SocketAddr, port: u32) -> Port {
    let timeout = tokio::time::Duration::from_millis(500);
    let mut is_open = false;

    addr.set_port(port as u16);

    if tokio::time::timeout(timeout, TcpStream::connect(&addr)).await.is_ok(){
        is_open = true;
    }

    Port{
        port,
        conn_open: is_open
    }
}



