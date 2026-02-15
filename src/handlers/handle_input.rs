use crate::{constants::COMMANDS, handlers};

pub fn handle_input(args: Vec<String>) {
    if args.len() <= 1 {
        handlers::handle_help();
        return;
    }

    let cmd = COMMANDS
        .iter()
        .find(|cmd| cmd.full_command == args[1] || cmd.command_alias == args[1]);

    match cmd {
        Some(cmd) => {
            (cmd.handler)();
        }
        None => {
            eprintln!("Error: Unknown command '{}'", args[1]);
            handlers::handle_help();
        }
    }
}
