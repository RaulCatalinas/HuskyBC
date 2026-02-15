use phf::{Map, phf_map};

use crate::handlers::core::{
    handle_full_configuration, handle_husky_commitlint, handle_husky_config,
    handle_husky_lint_staged,
};

pub const WIZARD_OPTION_FUNCTIONS: Map<&'static str, fn()> = phf_map! {
    "Only Husky" => handle_husky_config,
    "Husky + Commitlint" => handle_husky_commitlint,
    "Husky + Lint-staged" => handle_husky_lint_staged,
    "Husky + Commitlint + Lint-staged (Full)" => handle_full_configuration,
};
