use crate::{
    commands::{collaborate, help, init, version},
    types::CliCommand,
};

pub const COMMANDS: &[CliCommand] = &[
    CliCommand {
        full_command: "--help",
        command_alias: "-h",
        description: "Display this help message",
        handler: help::execute,
    },
    CliCommand {
        full_command: "--version",
        command_alias: "-v",
        description: "Show the current version",
        handler: version::execute,
    },
    CliCommand {
        full_command: "--init",
        command_alias: "-i",
        description: "Initialize Husky hooks in your Git repository",
        handler: init::execute,
    },
    CliCommand {
        full_command: "--collaborate",
        command_alias: "-c",
        description: "Open the GitHub repository in your browser",
        handler: collaborate::execute,
    },
];
