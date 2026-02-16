use std::process::exit;

use inquire::Select;

use crate::constants::WIZARD_OPTION_FUNCTIONS;

pub fn select_config_option() -> String {
    let options = WIZARD_OPTION_FUNCTIONS.keys().copied().collect();

    let ans = Select::new("What do you want to configure?", options).prompt();

    match ans {
        Ok(choice) => choice.to_string(),
        Err(e) => {
            eprintln!("Selection cancelled or failed: {}", e);

            exit(1);
        }
    }
}
