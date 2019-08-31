extern crate clap;
extern crate serde;
#[macro_use]
extern crate serde_derive;
#[macro_use]
extern crate log;
extern crate env_logger;
extern crate failure;

mod util;
mod providers;

use clap::{App, AppSettings, Arg, SubCommand};

use util::open_file;

const VERSION: &str = "v0.1.0-alpha";

const ASCIIART: &str = r#"
                       _ 
                      (_) 
  _ __ ___   ___   __ _ _ 
 | '_ ` _ \ / _ \ / _` | |
 | | | | | | (_) | (_| | |
 |_| |_| |_|\___/ \__,_|_|

Manage public/private cloud as declarative resources"
"#;

fn main() {
    let filename = Arg::with_name("filename")
        .help("resource file to apply")
        .long("filename")
        .short("f")
        .required(true)
        .takes_value(true);

    let matches = App::new("moai")
        .author("Cloudflavor Org")
        .version(VERSION)
        .about(ASCIIART)
        .setting(AppSettings::SubcommandRequiredElseHelp)
        .arg(Arg::with_name("verbosity")
                .help("verbosity level 0-4")
                .long("verbosity")
                .short("v")
                .multiple(true))
        .arg(Arg::with_name("workspace")
                .help("choose a workspace")
                .long("workspace")
                .short("w")
                .global(true))
        .subcommand(SubCommand::with_name("create")
                .display_order(1)
                .about("create new resource(s)")
                .arg(&filename))
        .subcommand(SubCommand::with_name("apply")
                .display_order(2)
                .about("modify (patch) existing resources")
                .arg(&filename))
        .subcommand(SubCommand::with_name("delete")
                .display_order(3)
                .about("delete existing resources")
                .arg(&filename))
        .get_matches();

    let log_level = match matches.occurrences_of("verbosity") {
        0 => log::LevelFilter::Error,
        1 => log::LevelFilter::Warn,
        2 => log::LevelFilter::Info,
        3 => log::LevelFilter::Debug,
        _ => log::LevelFilter::Trace,
    };

    env_logger::Builder::from_default_env()
        .filter(Some(module_path!()), log_level)
        .init();

    let mut file_path = matches.value_of("filename");

    match matches.subcommand_name().unwrap() {
        "create" => create_resources(file_path.unwrap()),
        "apply" => debug!("apply"),
        "delete" => debug!("delete"),
        _ => ()
    }
}

fn create_resources(file_path: &str) {
    let deserialize = open_file(file_path).unwrap();
}