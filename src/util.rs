use super::providers::virt::Libvirt;

use config::{Config, ConfigError, Environment, File};

#[derive(Deserialize, Debug)]
pub struct Configuration {
    libvirt: Option<Libvirt>
}

impl Default for Configuration {
    fn default() -> Self {
        Self{
            libvirt: Some(Libvirt::default())
        }
    }
}

impl Configuration {
    pub fn new(path: &str) -> Result<Self, ConfigError> {
        let mut c = Config::new();
        c.merge(File::with_name(path))?;
        c.merge(Environment::with_prefix("MOAI_CONFIG"))?;
        c.try_into()
    }
}
