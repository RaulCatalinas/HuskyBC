use std::process::exit;

use inquire::{Confirm, Select};

use crate::types::PackageManager;

pub fn select_package_manger() -> PackageManager {
    let options = PackageManager::ALL.to_vec();

    let ans = Select::new("Which package manager do you want to use?", options).prompt();

    match ans {
        Ok(choice) => choice,
        Err(e) => {
            eprintln!("Selection cancelled or failed: {}", e);
            exit(1);
        }
    }
}

pub fn will_be_published_on_npm() -> bool {
    let ans = Confirm::new("Publish to npm?")
        .with_default(false)
        .with_help_message("Adds prepack/postpack scripts for Yarn")
        .prompt();

    match ans {
        Ok(published) => published,
        Err(e) => {
            eprintln!("Selection cancelled: {}", e);
            exit(1);
        }
    }
}
