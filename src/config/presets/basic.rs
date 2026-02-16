use crate::{config::husky, types::CliContext, utils::npm::install_dependencies};

pub fn execute(ctx: CliContext) {
    install_dependencies(ctx.package_manager, &["husky"]);

    husky::config(ctx);
}
