import Lungo

namespace Shapes

structure Rect where
  width : Nat
  height : Nat

def area (rectangle : Rect) : Nat :=
  rectangle.width * rectangle.height

@[lungo_claim "lungo.law" subject area]
theorem area_matches_product (width height : Nat) :
    area ⟨width, height⟩ = width * height := rfl

end Shapes
