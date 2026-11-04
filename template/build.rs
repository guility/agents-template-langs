fn main() -> lungo_build::Result<()> {
    lungo_build::configure()
        .deny_axioms(true)
        .deny_unsafe(true)
        .require_claims(".")
        .compile_lean("lean")
}
