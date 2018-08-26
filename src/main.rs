extern crate clap;
extern crate config;
extern crate serde;

#[macro_use]
extern crate serde_derive;

mod configuration;
mod controller;
mod core;
mod providers;

use clap::{App, AppSettings, Arg, SubCommand};

use configuration::Configuration;
use core::bootstrap;

// TODO: this should probably be moved somewhere else.
const VERSION: &str = "v0.1.0-alpha";

fn main() {
  let new_cli = App::new("dweller")
    .author("Cloudflavor Org")
    .version(VERSION)
    .about("Manage opinionated infrastructure for private and public cloud")
    .setting(AppSettings::SubcommandRequiredElseHelp)
    .subcommand(
      SubCommand::with_name("new")
        .display_order(1)
        .about("provision a new cluster"),
    )
    .subcommand(
      SubCommand::with_name("mod")
        .display_order(2)
        .about("modify an existing cluster"),
    )
    .subcommand(
      SubCommand::with_name("delete")
        .display_order(3)
        .about("provision a new cluster"),
    )
    .arg(
      Arg::with_name("provider")
        .help("specify which provider to use. See the docs for supported providers")
        .required(true),
    )
    .arg(
      Arg::with_name("config")
        .short("c")
        .long("config")
        .value_name("JSON, TOML, YAML, HJSON, INI - configuration")
        .help("Path to config file")
        .takes_value(true),
    )
    .get_matches();

  if let Some(matches) = new_cli.subcommand_matches("new") {
    let config = Configuration::new(matches.value_of("config").unwrap());
    bootstrap(config.unwrap());
  }
}
