use crate::consts;
use tracing_log::log;
use tokio::{net::TcpStream, sync::mpsc};
use futures::stream::StreamExt;
use tokio_stream::Stream;
use tonic::{Request, Response, Status, Streaming};
use std::{net::{ToSocketAddrs, SocketAddr}, pin::Pin};
use crate::subdomain::api::portscan::v1::{port_scan_service_server::PortScanService, Subdomain, Port, PortScanRequest};


#[derive(Debug, Default)]
pub struct PortScanComponent {}

#[tonic::async_trait]
impl PortScanService for PortScanComponent {
    type ScanForOpenPortsStream = Pin<Box<dyn Stream<Item = Result<Subdomain, Status>> + Send + 'static>>;

    async fn scan_for_open_ports(
        &self,
        request: Request<Streaming<PortScanRequest>>,
    )-> Result<Response<Self::ScanForOpenPortsStream>, Status>{
        log::info!("Scanning for open ports...");

        let buffer = 100000;
        let mut  stream = request.into_inner();
        let output = async_stream::try_stream!{
            while let Some(req) = stream.next().await {
                let req = req.unwrap();
                let address = format!("{}:1024", req.host)
                    .to_socket_addrs()
                    .unwrap()
                    .next();


                if let None = address {
                    yield Subdomain { domain: req.host.clone(), ports: vec![] };
                    continue
                }

                let (input_tx, input_rx) = mpsc::channel::<u32>(buffer);
                let (output_tx, output_rx) = mpsc::channel::<Port>(buffer);
                
                tokio::spawn(async move {
                    for port in consts::PORTS_LIST{
                        let _ = input_tx.send(*port).await;

                    }
                    drop(input_tx);
                });

                let input_receiver_stream = tokio_stream::wrappers::ReceiverStream::new(input_rx);
                input_receiver_stream
                    .for_each_concurrent(buffer, |port| {
                        let output_tx = output_tx.clone();

                        async move {
                            let port: Port = scan_port(address.unwrap(), port).await;
                            if port.conn_open{
                                let _ = output_tx.send(port).await;
                            }
                        }
                    })
                    .await;
                drop(output_tx);

                let output_receiver_stream = tokio_stream::wrappers::ReceiverStream::new(output_rx);
                

                yield Subdomain { domain: req.host.clone(), ports: output_receiver_stream.collect::<Vec<Port>>().await };
            }
        };

        Ok(Response::new(Box::pin(output) as Self::ScanForOpenPortsStream))
    }
}

async fn scan_port(mut addr: SocketAddr, port: u32) -> Port {
    let timeout = tokio::time::Duration::from_secs(3);
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



