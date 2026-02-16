use crate::{
    config::{commitlint, husky, lint_staged},
    types::CliContext,
    utils::npm::install_dependencies,
};

pub fn execute(ctx: CliContext) {
    install_dependencies(
        ctx.package_manager,
        &[
            "husky",
            "@commitlint/cli",
            "@commitlint/config-conventional",
            "lint-staged",
        ],
    );

    husky::config(ctx);
    lint_staged::config(ctx);
    commitlint::config(ctx);
}
