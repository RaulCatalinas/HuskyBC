use std::process::exit;

use crate::{
    cli::select_config_option, constants::WIZARD_OPTION_FUNCTIONS, utils::is_nodejs_project,
};

pub fn handle_init() {
    if !is_nodejs_project() {
        eprintln!("❌ Error: package.json not found");
        eprintln!("Make sure you're in a Node.js project");

        exit(1);
    }

    let choice = select_config_option();

    let func = WIZARD_OPTION_FUNCTIONS.get(&choice).unwrap();

    func();
}
