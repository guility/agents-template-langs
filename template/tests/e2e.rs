use std::process::Command;

#[test]
fn test_cli_prints_compiled_lean_area() {
    let output = Command::new(env!("CARGO_BIN_EXE_lean-rust-component"))
        .output()
        .expect("component binary should start");
    assert!(output.status.success());
    assert_eq!(String::from_utf8_lossy(&output.stdout).trim(), "12");
}
