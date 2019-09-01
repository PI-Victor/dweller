#[allow(dead_code)]
pub enum ProviderType {
    Libvirt,
    DigitalOcean,
    AWS,
    GCP,
    Azure
}

pub trait Provider{}
