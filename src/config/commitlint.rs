use crate::{
    types::{CliContext, PackageManager},
    utils::fs::write_file,
};
use std::{process::exit, thread::spawn};

pub fn config(ctx: CliContext) {
    let commitlint_command = match ctx.package_manager {
        PackageManager::Npm => "npx --no -- commitlint --edit $1",
        PackageManager::Pnpm => "pnpm dlx commitlint --edit $1",
        PackageManager::Bun => "bunx commitlint --edit $1",
        PackageManager::Yarn => "yarn commitlint --edit $1",
    };

    let thread_1 = spawn(move || {
        write_file(
            ".",
            "commitlint.config.mjs",
            "export default { extends: ['@commitlint/config-conventional'] };",
        );
    });

    let thread_2 = spawn(move || {
        write_file(".husky", "commit-msg", commitlint_command);
    });

    if let Err(e) = thread_1.join() {
        eprintln!("✗ Thread 1 panicked: {:?}", e);
        exit(1);
    }

    if let Err(e) = thread_2.join() {
        eprintln!("✗ Thread 2 panicked: {:?}", e);
        exit(1);
    }
}
