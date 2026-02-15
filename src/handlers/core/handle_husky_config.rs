use crate::{handlers::core::common::get_install_command, utils::execute_command};

pub fn handle_husky_config() {
    let (install_command, mut args) = get_install_command();

    args.push("husky");

    execute_command(&install_command, &args, None);

    // TODO: Add husky's configuration logic
}
