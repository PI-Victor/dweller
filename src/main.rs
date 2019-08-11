extern crate clap;
extern crate serde;
#[macro_use]
extern crate serde_derive;
#[macro_use]
extern crate log;
extern crate env_logger;

mod util;
mod providers;

use clap::{App, AppSettings, Arg, SubCommand};
use util::Configuration;

const VERSION: &str = "v0.1.0-alpha";

const ASCIIART: &str = r#"
                       _ 
                      (_) 
  _ __ ___   ___   __ _ _ 
 | '_ ` _ \ / _ \ / _` | |
 | | | | | | (_) | (_| | |
 |_| |_| |_|\___/ \__,_|_|
                          
Manage opinionated infrastructure for private cloud"                          
"#;

fn main() {
    let provs: [&str; 1] = ["libvirt"];
    let matches = App::new("moai")
        .author("Cloudflavor Org")
        .version(VERSION)
        .about(ASCIIART)
        .setting(AppSettings::SubcommandRequiredElseHelp)
        .arg(Arg::with_name("v")
                .help("verbosity level 0-4")
                .short("v")
                .multiple(true))
        .subcommand(SubCommand::with_name("create")
                .display_order(1)
                .about("provision a new cluster"))
        .subcommand(SubCommand::with_name("modify")
                .display_order(2)
                .about("modify an existing cluster"))
        .subcommand(SubCommand::with_name("delete")
                .display_order(3)
                .about("delete a cluster"))
        .arg(Arg::with_name("provider")
                .help("specify which provider to use. See the docs for supported providers")
                .default_value("libvirt")
                .possible_values(&provs))
        .arg(Arg::with_name("config")
                .short("c")
                .long("config")
                .value_name("JSON, TOML, YAML, HJSON, INI - configuration")
                .help("Path to config file")
                .takes_value(true)
                .required(true))
        .get_matches();

    let log_level = match matches.occurrences_of("v") {
        0 => log::LevelFilter::Error,
        1 => log::LevelFilter::Warn,
        2 => log::LevelFilter::Info,
        3 => log::LevelFilter::Debug,
        _ => log::LevelFilter::Trace,
    };

    env_logger::Builder::from_default_env()
        .filter(Some(module_path!()), log_level)
        .init();
    let mut config = Configuration::default();

    if let Some(matches) = matches.subcommand_matches("new") {
        config = Configuration::new(matches.value_of("config").unwrap())
            .unwrap();
    }
    debug!("Loaded configuration: {:?}", config);
}
