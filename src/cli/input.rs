use std::env::args;

use crate::{commands::help, constants::COMMANDS};

pub fn parse_and_dispatch() {
    let args: Vec<String> = args().collect();

    if args.len() <= 1 {
        help::execute();
        return;
    }

    dispatch_command(args[1].trim());
}

fn dispatch_command(command: &str) {
    let cmd = COMMANDS
        .iter()
        .find(|cmd| cmd.full_command == command || cmd.command_alias == command);

    match cmd {
        Some(cmd) => (cmd.handler)(),
        None => {
            eprintln!("Error: Unknown command '{}'", command);
            help::execute();
        }
    }
}
