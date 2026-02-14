use crate::handlers::{handle_collaborate, handle_help, handle_init, handle_version};
use crate::types::CliCommand;

pub const COMMANDS: &[CliCommand] = &[
    CliCommand {
        full_command: "--help",
        command_alias: "-h",
        description: "Display help information",
        handler: handle_help,
    },
    CliCommand {
        full_command: "--collaborate",
        command_alias: "-c",
        description: "Collaborate with team",
        handler: handle_collaborate,
    },
    CliCommand {
        full_command: "--version",
        command_alias: "-v",
        description: "Show version",
        handler: handle_version,
    },
    CliCommand {
        full_command: "--init",
        command_alias: "-i",
        description: "Initialize Husky configuration in your project",
        handler: handle_init,
    },
];
