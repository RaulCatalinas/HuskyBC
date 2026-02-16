use crate::config::{commitlint, husky, lint_staged};

pub fn execute() {
    husky::config();
    lint_staged::config();
    commitlint::config();
}
