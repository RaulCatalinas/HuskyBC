use crate::{constants::COMMANDS, types::CommandResult};

pub fn handle_help(_: &[String]) -> CommandResult {
    println!("Usage: huskybc [options]\n");
    println!("Command line for easy Husky configuration\n");

    for cmd in COMMANDS {
        println!(
            "  {}, {}: {}",
            cmd.full_command, cmd.command_alias, cmd.description
        );
    }

    Ok(())
}
