use std::process::{Command, Stdio, exit};

pub fn execute_command(program: &str, args: &[&str], exit_on_failure: Option<bool>) {
    match Command::new(program)
        .args(args)
        .stdout(Stdio::piped())
        .stderr(Stdio::piped())
        .output()
    {
        Ok(output) => {
            if output.status.success() {
                println!(
                    "✓ Command '{} {}' executed successfully",
                    program,
                    args.join(" ")
                );
            } else {
                eprintln!(
                    "✗ Command '{}' failed with status: {}",
                    program, output.status
                );

                if exit_on_failure.unwrap_or(false) {
                    exit(1);
                }
            }
        }
        Err(e) => {
            eprintln!("✗ Failed to execute '{}': {}", program, e);
            if exit_on_failure.unwrap_or(false) {
                exit(1);
            }
        }
    }
}
