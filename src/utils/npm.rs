use crate::cli::prompts::select_package_manger;

pub fn get_install_command() -> (String, Vec<&'static str>) {
    let package_manger = select_package_manger();

    let mut install_command_args: Vec<&'static str> = vec![];

    if package_manger == "npm" {
        install_command_args.push("install");
    } else {
        install_command_args.push("add");
    }

    install_command_args.push("-D");

    return (package_manger, install_command_args);
}
