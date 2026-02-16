use std::{process::exit, thread::spawn};

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

    let thread_1 = spawn(move || lint_staged::config(ctx));
    let thread_2 = spawn(move || commitlint::config(ctx));

    if let Err(e) = thread_1.join() {
        eprintln!("✗ Full configuration thread 1 panicked: {:?}", e);
        exit(1);
    }

    if let Err(e) = thread_2.join() {
        eprintln!("✗ Full configuration thread 2 panicked: {:?}", e);
        exit(1);
    }
}
