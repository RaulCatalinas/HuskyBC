use crate::constants::GITHUB_REPO_URL;

pub fn handle_collaborate() {
    println!("Opening HuskyBC repository...");

    match open::that(GITHUB_REPO_URL) {
        Ok(_) => {
            println!("✓ Repository opened in your default browser");
            println!("{}", GITHUB_REPO_URL);
        }
        Err(e) => {
            eprintln!("✗ Failed to open browser: {}", e);
            println!("Please visit manually: {}", GITHUB_REPO_URL);
        }
    }
}
