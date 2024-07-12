use tracing_log::log;
use tokio::sync::mpsc;
use futures::StreamExt;
use std::{fs::File, io::{BufRead, BufReader}};
use tonic::{Request, Response, Status};
use tokio_stream::wrappers::ReceiverStream;
use crate::{errors::SourceFileError, subdomain::api::brute::v1::{brute_service_server::BruteService, BruteForceRequest, BruteForceResponse}};

#[derive(Debug, Default, Clone)] 
pub struct BruteForceComponent {}

#[derive(Debug)]
struct SourceFile {
    name: String
}

impl SourceFile {
    fn new(name: &str) -> SourceFile {
        SourceFile { name: name.to_string() }
    }

    fn reader(&self) -> Result<BufReader<File>, SourceFileError>{
        let f = File::open(&self.name)?;
        let reader  = BufReader::new(f);
        Ok(reader)
    }
}

#[tonic::async_trait]
impl BruteService for BruteForceComponent {
    #[tracing::instrument(skip_all)]
    async fn get_subdomains_by_brute_force(
        &self, 
        request: Request<BruteForceRequest>,
    ) -> Result<Response<BruteForceResponse>, Status>{
        log::info!("Performing subdomain enumeration by using brute force from words.txt");
        
        let subdomains: Vec<String>;
        let _source = SourceFile::new("./words.txt");
        println!("{:?}", _source.reader());
        match _source.reader() {
            Ok(reader) => {

                let buffer: usize = 1000000;
                let (input_tx, input_rx ) = mpsc::channel::<String>(buffer);
                let (output_tx, output_rx ) = mpsc::channel::<String>(buffer);

                
            tokio::spawn(async move {
                for line in reader.lines() {
                    if let Ok(line) = line{
                        log::info!("{:?} sending", &line);
                        let _ = input_tx.send(line).await;
                    }
                }
                drop(input_tx);
            });
    
        
            let target = request.get_ref().target.as_str();
            if target.is_empty() {
                return Err(Status::invalid_argument("invalid argument provided"));
            };
            let input_rx_stream = ReceiverStream::new(input_rx);
            input_rx_stream
                .for_each_concurrent(buffer, |prefix| {
                    let output_tx = output_tx.clone();
                    async move {
                        let subdomain = format!("{}.{}", prefix, &target);
                        log::info!("bruteforce subdomain: {}", &subdomain);
                        let _ = output_tx.send(subdomain).await;
                    }
                })
                .await;
                
                drop(output_tx);

                let output_rx_stream = ReceiverStream::new(output_rx);
                subdomains = output_rx_stream.collect::<Vec<String>>().await;

                log::info!("bruteforce scan done!");
            }

            Err(e) =>  {
                log::error!("error trying to get reader -> {}", e);
                return Err(Status::internal("error trying to get reader"))
            }
        }  

        Ok(Response::new(BruteForceResponse{subdomains: subdomains}))
    }
}

