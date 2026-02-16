use crate::types::PackageManager;

pub struct CliCommand {
    pub full_command: &'static str,
    pub command_alias: &'static str,
    pub description: &'static str,
    pub handler: fn() -> (),
}

#[derive(Debug, Clone, Copy)]
pub struct CliContext {
    pub package_manager: PackageManager,
    pub will_publish: bool,
}
