use crate::config::{commitlint, husky};

pub fn execute() {
    husky::config();
    commitlint::config();
}
