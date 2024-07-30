use std::sync;
use std::net::SocketAddr;
use envconfig::Envconfig;

static CONFIG: sync::OnceLock<Config> = sync::OnceLock::new();

#[derive(Envconfig, Debug, Clone)]
pub struct TelemetryConfig {
    #[envconfig(from = "TRACING_ENABLED", default = "0")]
    pub otel_tracing_enabled: String,

    #[envconfig(from = "OTEL_COLLECTOR_ADDR", default = "http://jaeger-otel.jaeger.svc.cluster.local:14278/api/traces")]

    pub otel_endpoint: String,
}


#[derive(Envconfig, Debug, Clone)]
pub struct Config{
    #[envconfig(nested=true)]
    pub telemetry: TelemetryConfig,

    #[envconfig(
        from = "LISTEN_ADDR",
        default = "0.0.0.0:50052"
    )]
    pub listen_addr: SocketAddr,
}

impl Config {
    pub fn get() -> &'static Self {
        CONFIG.get_or_init(|| Config::init_from_env().unwrap())
    }

    pub fn set(config: Config) -> &'static Self {
        match CONFIG.get() {
            None => {
                CONFIG.set(config).expect("Failed to set config value");
                Config::get()
            }
            Some(v) => {
                panic!("Config value is already set {:?}", v)
            }
        }
    }
}
