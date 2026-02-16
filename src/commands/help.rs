use colored::Colorize;

use crate::constants::COMMANDS;

pub fn execute() {
    println!("\n{}", "HuskyBC".bold());
    println!("{}\n", "Command line for easy Husky configuration".dimmed());

    println!("{}", "Usage:".bold());
    println!("  huskybc [command]\n");

    println!("{}", "Commands:".bold());

    let max_command_width = COMMANDS
        .iter()
        .map(|cmd| cmd.full_command.len())
        .max()
        .unwrap_or(0);

    let max_alias_width = COMMANDS
        .iter()
        .map(|cmd| cmd.command_alias.len())
        .max()
        .unwrap_or(0);

    for cmd in COMMANDS {
        println!(
            "  {:command_width$}  {:alias_width$}  {}",
            cmd.full_command,
            cmd.command_alias.dimmed(),
            cmd.description.dimmed(),
            command_width = max_command_width,
            alias_width = max_alias_width
        );
    }
}
