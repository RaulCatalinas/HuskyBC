use crate::{types::PackageManager, utils::execute::execute_command};

pub fn install_dependencies(package_manager: PackageManager, deps: &[&str]) {
    let mut args: Vec<&str> = vec![];

    if package_manager == PackageManager::Npm {
        args.push("install");
    } else {
        args.push("add");
    }

    args.push("-D");
    args.extend(deps);

    execute_command(package_manager.as_str(), &args);
}
