use crate::config::{husky, lint_staged};

pub fn execute() {
    husky::config();
    lint_staged::config();
}
