use phf::{Map, phf_map};

use crate::config::presets::{basic, full, with_commitlint, with_lint_staged};

pub const WIZARD_OPTION_FUNCTIONS: Map<&'static str, fn()> = phf_map! {
    "Only Husky" => basic::execute,
    "Husky + Commitlint" => with_commitlint::execute,
    "Husky + Lint-staged" => with_lint_staged::execute,
    "Husky + Commitlint + Lint-staged (Full)" => full::execute,
};
