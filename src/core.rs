use configuration::{Configuration};


pub trait Provider {
  fn bootstrap_cluster(&self);
  fn modify_cluster(&self);
  fn get_cluster_info(&self);
}

pub trait Cluster {
  fn get_nodes(&self);
  fn check_nodes(&self);
  fn modify_nodes(&self);
}

pub trait Node {
  fn start(&self);
  fn stop(&self);
  fn suspend(&self);
  fn status(&self);
}

pub fn bootstrap(config: Configuration) {
  
}