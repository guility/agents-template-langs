fn main() -> lungo_build::Result<()> {
    lungo_build::configure()
        .deny_axioms(true)
        .deny_unsafe(true)
        .require_claims("Shapes.area")
        .compile_lean("lean")
}
