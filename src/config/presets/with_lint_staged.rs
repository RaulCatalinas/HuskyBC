use crate::{
    config::{husky, lint_staged},
    types::CliContext,
    utils::npm::install_dependencies,
};

pub fn execute(ctx: CliContext) {
    install_dependencies(ctx.package_manager, &["husky", "lint-staged"]);

    husky::config(ctx);
    lint_staged::config(ctx);
}
