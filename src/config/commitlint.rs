use crate::{
    types::{CliContext, PackageManager},
    utils::{
        fs::write_file,
        terminal::{start_spinner, stop_spinner},
    },
};

pub fn config(ctx: CliContext) {
    let spinner = start_spinner("Configuring Commitlint");

    let commitlint_command = match ctx.package_manager {
        PackageManager::Npm => "npx --no -- commitlint --edit $1",
        PackageManager::Pnpm => "pnpm dlx commitlint --edit $1",
        PackageManager::Bun => "bunx commitlint --edit $1",
        PackageManager::Yarn => "yarn dlx commitlint --edit $1",
    };

    write_file(
        ".",
        "commitlint.config.mjs",
        "export default { extends: ['@commitlint/config-conventional'] };",
    );
    write_file(".husky", "commit-msg", commitlint_command);

    stop_spinner(&spinner, "Commitlint successfully configured");
}
