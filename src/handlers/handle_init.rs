use crate::{cli::select_config_option, constants::WIZARD_OPTION_FUNCTIONS};

pub fn handle_init() {
    let choice = select_config_option();

    let func = WIZARD_OPTION_FUNCTIONS.get(&choice).unwrap();

    func();
}
