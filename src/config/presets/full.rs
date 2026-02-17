use crate::{
    config::{commitlint, husky, lint_staged},
    types::CliContext,
    utils::{
        npm::install_dependencies,
        terminal::{start_spinner, stop_spinner},
    },
};

pub fn execute(ctx: CliContext) {
    let mut spinner = start_spinner("Installing Husky + Commitlint + Lint-staged...");

    install_dependencies(
        ctx.package_manager,
        &[
            "husky",
            "@commitlint/cli",
            "@commitlint/config-conventional",
            "lint-staged",
        ],
    );

    stop_spinner(&spinner, "Dependencies installed");

    spinner = start_spinner("Setting up full Git hooks workflow...");

    husky::config(ctx);
    lint_staged::config(ctx);
    commitlint::config(ctx);

    stop_spinner(&spinner, "Full configuration ready");
}
