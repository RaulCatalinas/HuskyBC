use crate::handlers::core::handle_husky_config;

pub fn handle_husky_lint_staged() {
    handle_husky_config();

    println!("Configuring lint-staged...");
}
