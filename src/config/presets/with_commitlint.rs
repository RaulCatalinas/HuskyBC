use crate::{
    config::{commitlint, husky},
    types::CliContext,
    utils::npm::install_dependencies,
};
use std::{process::exit, thread::spawn};

pub fn execute(ctx: CliContext) {
    install_dependencies(
        ctx.package_manager,
        &[
            "husky",
            "@commitlint/cli",
            "@commitlint/config-conventional",
        ],
    );

    let thread_1 = spawn(move || husky::config(ctx));
    let thread_2 = spawn(move || commitlint::config(ctx));

    if let Err(e) = thread_1.join() {
        eprintln!("✗ Commitlint thread 1 panicked: {:?}", e);
        exit(1);
    }

    if let Err(e) = thread_2.join() {
        eprintln!("✗ Commitlint thread 2 panicked: {:?}", e);
        exit(1);
    }
}
