use std::process::{Command, Stdio, exit};

pub fn execute_command(program: &str, args: &[&str]) {
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

                exit(1);
            }
        }
        Err(e) => {
            eprintln!("✗ Failed to execute '{}': {}", program, e);

            exit(1);
        }
    }
}
