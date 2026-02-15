use crate::{handlers::core::common::get_install_command, utils::execute_command};

pub fn handle_husky_config() {
    let (package_manager, mut args) = get_install_command();

    args.push("husky");

    execute_command(&package_manager, &args);

    if package_manager == "npm" {
        execute_command("npx", &["husky", "init"]);
        return;
    }

    if package_manager == "pnpm" {
        execute_command("npx", &["exec", "husky", "init"]);
        return;
    }

    if package_manager == "bun" {
        execute_command("bunx", &["husky", "init"]);
        return;
    }

    todo!("Add logic to configure husky with yarn")
}
