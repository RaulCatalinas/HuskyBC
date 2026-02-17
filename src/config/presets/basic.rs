use crate::{
    config::husky,
    types::CliContext,
    utils::{
        npm::install_dependencies,
        terminal::{start_spinner, stop_spinner},
    },
};

pub fn execute(ctx: CliContext) {
    let mut spinner = start_spinner("Installing Husky...");

    install_dependencies(ctx.package_manager, &["husky"]);

    stop_spinner(&spinner, "Husky installed");

    spinner = start_spinner("Setting up Git hooks...");

    husky::config(ctx);

    stop_spinner(&spinner, "Git hooks ready");
}
