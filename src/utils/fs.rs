use serde_json::{Value, json};
use std::{fs, io::Result, path::Path, process::exit};

pub fn is_nodejs_project() -> bool {
    Path::new("package.json").exists()
}

fn handle_io_result<T>(result: Result<T>, success_msg: &str, error_context: &str) {
    match result {
        Ok(_) => println!("✓ {}", success_msg),
        Err(e) => {
            eprintln!("✗ {}: {}", error_context, e);
            exit(1);
        }
    }
}

fn parse_json_file(file_name: &str) -> Value {
    let content = match fs::read_to_string(file_name) {
        Ok(text) => text,
        Err(e) => {
            eprintln!("✗ Error reading package.json: {}", e);
            exit(1);
        }
    };

    match serde_json::from_str::<Value>(&content) {
        Ok(json) => json,
        Err(e) => {
            eprintln!("✗ Error parsing package.json: {}", e);
            exit(1);
        }
    }
}

pub fn write_scripts_in_package_json(scripts: Vec<(&str, &str)>) {
    let mut package_json = parse_json_file("package.json");

    if package_json.get("scripts").is_none() {
        package_json["scripts"] = json!({});
    }

    let scripts_section = package_json["scripts"]
        .as_object_mut()
        .expect("scripts should be an object");

    for (key, value) in scripts {
        scripts_section.insert(key.to_string(), json!(value));
    }

    let updated_json = match serde_json::to_string_pretty(&package_json) {
        Ok(json) => json,
        Err(e) => {
            eprintln!("✗ Error serializing JSON: {}", e);
            exit(1);
        }
    };

    handle_io_result(
        fs::write("package.json", updated_json),
        "✓ Scripts added to package.json",
        "✗ Error writing package.json",
    );
}

pub fn write_file(path: &str, file_name: &str, content: &str, success_msg: &str) {
    let dir_path = Path::new(path);
    let full_path = dir_path.join(file_name);

    if !dir_path.exists() {
        handle_io_result(
            fs::create_dir_all(dir_path),
            &format!("Created directory {}", dir_path.display()),
            &format!("Failed to create directory {}", dir_path.display()),
        );
    }

    handle_io_result(
        fs::write(&full_path, content),
        success_msg,
        &format!("Failed to write {}", full_path.display()),
    );
}

pub fn write_json_file(path: &str, file_name: &str, content: &str) {
    let json_value: Value = match serde_json::from_str(content) {
        Ok(json) => json,
        Err(e) => {
            eprintln!("✗ Invalid JSON content: {}", e);
            exit(1);
        }
    };

    let formatted = match serde_json::to_string_pretty(&json_value) {
        Ok(json) => json,
        Err(e) => {
            eprintln!("✗ Error serializing JSON: {}", e);
            exit(1);
        }
    };

    write_file(
        path,
        file_name,
        &formatted,
        &format!("✓ JSON written to {}", file_name),
    );
}
