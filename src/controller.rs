#[derive(Debug, Deserialize)]
pub struct Provider {
  pub connection_uri: String,
}

#[derive(Debug, Deserialize)]
pub struct Client {
  pub providers: Vec<Provider>,
}