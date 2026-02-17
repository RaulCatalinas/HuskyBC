use crate::{
    config::{commitlint, husky},
    types::CliContext,
    utils::{
        npm::install_dependencies,
        terminal::{start_spinner, stop_spinner},
    },
};

pub fn execute(ctx: CliContext) {
    let mut spinner = start_spinner("Installing Husky + Commitlint...");

    install_dependencies(
        ctx.package_manager,
        &[
            "husky",
            "@commitlint/cli",
            "@commitlint/config-conventional",
        ],
    );

    stop_spinner(&spinner, "Dependencies installed");

    spinner = start_spinner("Setting up Git hooks and commit linting...");

    husky::config(ctx);
    commitlint::config(ctx);

    stop_spinner(&spinner, "Husky + Commitlint ready");
}
