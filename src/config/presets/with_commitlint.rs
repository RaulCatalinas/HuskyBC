use crate::{
    config::{commitlint, husky},
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
        ],
    );

    husky::config(ctx);
    commitlint::config(ctx);
}
