package weightconv

import "fmt"

type Kilogram float64
type Pound float64

func (g Kilogram) String() string { return fmt.Sprintf("%.2fkg", g) }
func (p Pound) String() string    { return fmt.Sprintf("%.2flb", p) }

func KToP(g Kilogram) Pound { return Pound(g * 2.2046) }
func PToK(p Pound) Kilogram { return Kilogram(p / 2.2046) }
