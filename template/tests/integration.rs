use lean_rust_component::shapes::{Rect, area};
use lungo::Nat;

#[test]
fn test_area_large_natural_numbers() {
    let rectangle = Rect {
        width: Nat::from(u64::MAX),
        height: Nat::from(2u64),
    };
    assert_eq!(area(rectangle), Nat::from(u64::MAX) * Nat::from(2u64));
}
