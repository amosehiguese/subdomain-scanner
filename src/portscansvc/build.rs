use std::{env, path::PathBuf};

fn main(){
    let out_dir = PathBuf::from(env::var("OUT_DIR").unwrap());
    tonic_build::configure()
        .file_descriptor_set_path(out_dir.join("portscan_server.bin"))
        .compile(&["proto/api/v1/portscan.proto"], &["proto"])
        .unwrap();
}
