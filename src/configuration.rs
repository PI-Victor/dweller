extern crate config;


use config::{ConfigError, Config, File, Environment};


#[derive(Debug, Deserialize)]
pub struct Configuration {
  #[serde(default)]
  pub ssl_certificates: String,
  pub libvirt: Libvirt,
}

#[derive(Debug, Deserialize)]
pub struct ConnectionUri {}

#[derive(Debug, Deserialize)]
pub struct Libvirt {
  pub uri: Vec<ConnectionUri>,
  pub nodes: Vec<Node>,
}

#[derive(Debug, Deserialize)]
pub struct Node {}

impl Configuration {
    pub fn new(path: &str) -> Result<Self, ConfigError> {
        let mut c = Config::new();
        c.merge(File::with_name(path))?;
        c.merge(Environment::with_prefix("DWELLER"))?;
        c.try_into()
    }
}