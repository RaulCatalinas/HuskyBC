pub mod constants;
pub mod handlers;
pub mod types;

fn main() {
    let args: Vec<String> = std::env::args().collect();

    handlers::handle_input(args);
}
