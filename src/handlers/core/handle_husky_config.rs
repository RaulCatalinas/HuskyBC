use crate::{
    cli::will_be_published_on_npm,
    handlers::core::common::get_install_command,
    utils::{execute_command, write_scripts_in_package_json},
};

pub fn handle_husky_config() {
    let (package_manager, mut args) = get_install_command();

    args.push("husky");

    execute_command(&package_manager, &args);

    if package_manager == "npm" {
        execute_command("npx", &["husky", "init"]);
        return;
    }

    if package_manager == "pnpm" {
        execute_command("pnpm", &["exec", "husky", "init"]);
        return;
    }

    if package_manager == "bun" {
        execute_command("bunx", &["husky", "init"]);
        return;
    }

    let mut scripts = vec![("postinstall", "husky")];

    if will_be_published_on_npm() {
        scripts.push(("prepack", "pinst --disable"));
        scripts.push(("postpack", "pinst --enable"));
    }

    write_scripts_in_package_json(scripts);
}
