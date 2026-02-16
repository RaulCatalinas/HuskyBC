use crate::{
    types::{CliContext, PackageManager},
    utils::fs::{write_file, write_json_file},
};

pub fn config(ctx: CliContext) {
    let lint_stage_content = match ctx.package_manager {
        PackageManager::Npm => "npx lint-staged",
        PackageManager::Pnpm => "pnpm dlx lint-staged",
        PackageManager::Bun => "bunx lint-staged",
        PackageManager::Yarn => "yarn dlx lint-staged",
    };

    write_json_file(
        ".",
        ".lintstagedrc",
        r#"
            {
                "lint-staged": {
                    "*": [
                        "prettier --write --ignore-unknown"
                    ]
                }
            }
        "#,
    );

    write_file(".husky", "pre-commit", lint_stage_content);
}
