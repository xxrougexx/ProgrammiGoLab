package amazon

import "fmt"
import "math"

func (x Box) Volume() float64 {
  return float64(x.L) * float64(x.H) * float64(x.D)
}

func (x Sphere) Volume() float64 {
  return float64(4 / 3) * float64(x.R) * float64(math.Pi)
}

func (x Box) Length() int {
  return x.L
}

func (x Sphere) Length() int {
  return 2 * x.R
}

func (x Box) String() string {
  return fmt.Sprintf("Lunghezza: %d, Altezza: %d, Profondità: %d", x.L, x.H, x.D)
}
