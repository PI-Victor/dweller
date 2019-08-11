#[derive(Deserialize, Debug)]
pub struct Libvirt {
    pub uri: String
}

impl Default for Libvirt {
    fn default() -> Self {
        Self{
            uri: "qemu://system".to_string()
        }
    }
}