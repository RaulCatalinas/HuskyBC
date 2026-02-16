use crate::{
    types::PackageManager,
    utils::{
        execute::execute_command,
        terminal::{start_spinner, stop_spinner},
    },
};

pub fn install_dependencies(package_manager: PackageManager, deps: &[&str]) {
    let spinner = start_spinner("Installing dependencies");

    let mut args: Vec<&str> = vec![];

    if package_manager == PackageManager::Npm {
        args.push("install");
    } else {
        args.push("add");
    }

    args.push("-D");
    args.extend(deps);

    execute_command(package_manager.as_str(), &args);

    stop_spinner(&spinner, "Dependencies successfully installed");
}
