use crate::{
    config::{husky, lint_staged},
    types::CliContext,
    utils::{
        npm::install_dependencies,
        terminal::{start_spinner, stop_spinner},
    },
};

pub fn execute(ctx: CliContext) {
    let mut spinner = start_spinner("Installing Husky + Lint-staged...");

    install_dependencies(ctx.package_manager, &["husky", "lint-staged"]);

    stop_spinner(&spinner, "Dependencies installed");

    spinner = start_spinner("Setting up Git hooks and staged linting...");

    husky::config(ctx);
    lint_staged::config(ctx);

    stop_spinner(&spinner, "Husky + Lint-staged ready");
}
