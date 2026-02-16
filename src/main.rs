pub mod cli;
pub mod commands;
pub mod config;
pub mod constants;
pub mod types;
pub mod utils;

fn main() {
    cli::input::parse_and_dispatch();
}
