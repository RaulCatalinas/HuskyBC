use crate::{
    handlers::{handle_collaborate, handle_help, handle_init, handle_version},
    types::CliCommand,
};

pub const COMMANDS: &[CliCommand] = &[
    CliCommand {
        full_command: "--help",
        command_alias: "-h",
        description: "Display this help message",
        handler: handle_help,
    },
    CliCommand {
        full_command: "--version",
        command_alias: "-v",
        description: "Show the current version",
        handler: handle_version,
    },
    CliCommand {
        full_command: "--init",
        command_alias: "-i",
        description: "Initialize Husky hooks in your Git repository",
        handler: handle_init,
    },
    CliCommand {
        full_command: "--collaborate",
        command_alias: "-c",
        description: "Open the GitHub repository in your browser",
        handler: handle_collaborate,
    },
];
