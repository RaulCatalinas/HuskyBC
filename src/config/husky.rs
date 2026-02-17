use crate::{
    cli::prompts::will_be_published_on_npm,
    types::{CliContext, PackageManager},
    utils::{execute::execute_command, fs::write_scripts_in_package_json},
};

pub fn config(ctx: CliContext) {
    if ctx.package_manager == PackageManager::Npm {
        execute_command("npx", &["husky", "init"]);
        return;
    }

    if ctx.package_manager == PackageManager::Pnpm {
        execute_command("pnpm", &["exec", "husky", "init"]);
        return;
    }

    if ctx.package_manager == PackageManager::Bun {
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
