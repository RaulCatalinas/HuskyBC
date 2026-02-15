use crate::cli::select_package_manger;

pub fn get_install_command() -> (String, Vec<&'static str>) {
    let packege_manger = select_package_manger();

    let mut insrall_command_args: Vec<&'static str> = vec![];

    if packege_manger == "npm" {
        insrall_command_args.push("install");
    } else {
        insrall_command_args.push("add");
    }

    insrall_command_args.push("-D");

    return (packege_manger, insrall_command_args);
}
