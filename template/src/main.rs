use lean_rust_component::shapes::{Rect, area};
use lungo::Nat;

fn main() {
    let rectangle = Rect {
        width: Nat::from(3u64),
        height: Nat::from(4u64),
    };
    println!("{}", area(rectangle));
}
