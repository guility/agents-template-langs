use lean_rust_component::shapes::{Rect, area};
use lungo::Nat;

#[test]
fn test_area_three_by_four() {
    let rectangle = Rect {
        width: Nat::from(3u64),
        height: Nat::from(4u64),
    };
    assert_eq!(area(rectangle), Nat::from(12u64));
}
