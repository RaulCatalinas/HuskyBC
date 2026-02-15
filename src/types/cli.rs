pub struct CliCommand {
    pub full_command: &'static str,
    pub command_alias: &'static str,
    pub description: &'static str,
    pub handler: fn(&[String]) -> (),
}
