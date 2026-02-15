use std::path::Path;

pub fn is_nodejs_project() -> bool {
    Path::new("package.json").exists()
}
