use std::process::exit;

use crate::{
    cli::{
        prompts::{select_package_manger, will_be_published_on_npm},
        wizard::select_config_option,
    },
    constants::WIZARD_OPTION_FUNCTIONS,
    types::{CliContext, PackageManager},
    utils::fs::is_nodejs_project,
};

pub fn execute() {
    if !is_nodejs_project() {
        eprintln!("❌ Error: package.json not found");
        eprintln!("Make sure you're in a Node.js project");

        exit(1);
    }

    let package_manager = select_package_manger();

    let will_publish = if package_manager == PackageManager::Yarn {
        will_be_published_on_npm()
    } else {
        false
    };

    let ctx = CliContext {
        package_manager,
        will_publish,
    };

    let choice = select_config_option();

    let func = WIZARD_OPTION_FUNCTIONS.get(&choice).unwrap();

    func(ctx);
}
