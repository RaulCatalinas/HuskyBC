use std::time::Duration;

use indicatif::{ProgressBar, ProgressStyle};

pub fn start_spinner(msg: &str) -> ProgressBar {
    let spinner = ProgressBar::new_spinner();

    spinner.set_style(
        ProgressStyle::default_spinner()
            .template("{spinner:.green} {msg}")
            .unwrap(),
    );
    spinner.set_message(msg.to_string());
    spinner.enable_steady_tick(Duration::from_millis(100));

    spinner
}

pub fn stop_spinner(spinner: &ProgressBar, msg: &str) {
    spinner.finish_with_message(msg.to_string());
}
