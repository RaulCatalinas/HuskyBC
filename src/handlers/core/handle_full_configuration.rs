use crate::handlers::core::handle_husky_config;

pub fn handle_full_configuration() {
    handle_husky_config();

    println!("Configuring commitlint + lint-staged...");
}
