package eval

import (
	"math"
	"testing"
)

func TestMinEval(t *testing.T) {
	tests := []struct {
		expr Expr
		env  Env
		want float64
	}{
		{&min{literal(1.0), literal(2.0)}, Env{}, 1.0},
		{&min{Var("x"), literal(2.0)}, Env{"x": 2.1}, 2.0},
	}

	for _, tt := range tests {
		if got := tt.expr.Eval(tt.env); math.Abs(tt.want-got) > 1e-7 {
			t.Errorf("%#v.Eval(%#v) = %v; want %v", tt.expr, tt.env, got, tt.want)
		}
	}
}
