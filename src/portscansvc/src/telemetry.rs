use crate::config::TelemetryConfig;
use tracing::Subscriber;
use tracing_log::log;
use tracing_subscriber::{layer::SubscriberExt, registry::LookupSpan, util::SubscriberInitExt, EnvFilter, filter::LevelFilter};

pub fn init_telemetry(config: &TelemetryConfig) {
    if config.otel_tracing_enabled {
        init_telemetry_tracing(config);
        log::info!("Tracing is enabled");
    } else {
        init_regular_logging();
        log::info!("Tracing is disable");
    }
}

pub fn init_regular_logging() {
    let filter = EnvFilter::builder()
        .with_default_directive(LevelFilter::INFO.into())
        .from_env_lossy();

    tracing_subscriber::registry()
        .with(formatting_layer())
        .with(filter)
        .init();
}


fn init_telemetry_tracing(config: &TelemetryConfig){ 
    std::env::set_var("OTEL_EXPORTER_OTLP_ENDPOINT", config.otel_endpoint.clone());
    std::env::set_var("OTEL_EXPORTER_OTLP_TRACES_PROTOCOL", "grpc");

    std::env::set_var("OTEL_SERVICE_NAME", env!("CARGO_PKG_NAME"));

    tracing_subscriber::registry()
        .with(init_tracing_opentelemetry::tracing_subscriber_ext::build_otel_layer().unwrap())
        .with(init_tracing_opentelemetry::tracing_subscriber_ext::build_loglevel_filter_layer())
        .with(formatting_layer())
        .try_init()
        .unwrap()

}

fn formatting_layer<S>() -> Box<dyn tracing_subscriber::layer::Layer<S> + Send + Sync + 'static>
where
    S: Subscriber + for<'a> LookupSpan<'a>,
{
    Box::new(
        tracing_subscriber::fmt::Layer::default()
            .with_ansi(true)
            .with_target(true)
            .with_line_number(true)
            .with_thread_ids(true)
            .with_thread_names(true)
            .compact(),
    )
}
