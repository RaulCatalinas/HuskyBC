use std::process::{Command, Stdio, exit};

pub fn execute_command(program: &str, args: &[&str]) {
    let output = Command::new(program)
        .args(args)
        .stdout(Stdio::piped())
        .stderr(Stdio::piped())
        .output();

    match output {
        Ok(output) => {
            if !output.status.success() {
                eprintln!(
                    "✗ Command '{}' failed\n{}",
                    program,
                    String::from_utf8_lossy(&output.stderr)
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
