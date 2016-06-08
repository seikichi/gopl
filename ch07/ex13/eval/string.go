package eval

import (
	"bytes"
	"fmt"
)

func (e Var) String() string {
	return string(e)
}

func (e literal) String() string {
	return fmt.Sprintf("%g", e)
}

func (e unary) String() string {
	return fmt.Sprintf("(%c%s)", e.op, e.x)
}

func (e binary) String() string {
	return fmt.Sprintf("(%s %c %s)", e.x, e.op, e.y)
}

func (e call) String() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%s(", e.fn)
	for i, arg := range e.args {
		if i > 0 {
			buf.WriteString(", ")
		}
		write(buf, arg)
	}
	buf.WriteByte(')')
	return buf.String()
}
