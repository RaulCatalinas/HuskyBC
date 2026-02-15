pub mod core;
mod handle_collaborate;
mod handle_help;
mod handle_init;
mod handle_input;
mod handle_version;

pub use handle_collaborate::handle_collaborate;
pub use handle_help::handle_help;
pub use handle_init::handle_init;
pub use handle_input::handle_input;
pub use handle_version::handle_version;
