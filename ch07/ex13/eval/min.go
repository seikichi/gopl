package eval

import (
	"fmt"
	"math"
)

type min struct {
	x, y Expr
}

func (e min) String() string {
	return fmt.Sprintf("(min %c %s)", e.x, e.y)
}

func (b min) Eval(env Env) float64 {
	return math.Min(b.x.Eval(env), b.y.Eval(env))
}

func (b min) Check(vars map[Var]bool) error {
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)
}
