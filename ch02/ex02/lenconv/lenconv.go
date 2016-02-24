package lenconv

import "fmt"

type Meter float64
type Feet float64

func (m Meter) String() string { return fmt.Sprintf("%.2fm", m) }
func (f Feet) String() string  { return fmt.Sprintf("%.2fft", f) }

func MToF(m Meter) Feet { return Feet(m * 3.2808) }
func FToM(f Feet) Meter { return Meter(f / 3.2808) }
