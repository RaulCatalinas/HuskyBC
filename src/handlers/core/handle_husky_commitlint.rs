use crate::handlers::core::handle_husky_config;

pub fn handle_husky_commitlint() {
    handle_husky_config();

    println!("Configuring commitlint...");
}
