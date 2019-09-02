#[derive(Serialize, Deserialize, Debug)]
enum Kind {
    Domain,
    Pool,
    Network,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct Network<K> {
    pub kind: K,
    pub name: String 
}

#[derive(Serialize, Deserialize, Debug)]
pub struct StoragePool<K> {
    pub kind: K,
    pub name: String,
    pub allocation: u16,
    pub capacity: u16,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct Domain<K> {
    pub kind: K,
    pub domain_type: String,
    pub name: String,
    pub vcpu: u8,
    pub memory: u8,
    pub disks: StorageDisks,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct StorageDisks {}

#[derive(Serialize, Deserialize, Debug)]
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
