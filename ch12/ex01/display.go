package display

import (
	"fmt"
	"reflect"
	"strconv"
)

type target struct {
	name string
	x    interface{}
}

type queue struct {
	count   int
	targets []target
}

func (q *queue) push(x interface{}) string {
	q.count++
	name := fmt.Sprintf("$%d", q.count)
	q.targets = append(q.targets, target{name, x})
	return name
}

func (q *queue) isEmpty() bool {
	return len(q.targets) == 0
}

func (q *queue) pop() target {
	ret := q.targets[0]
	q.targets = q.targets[1:]
	return ret
}

func Display(name string, x interface{}) {
	q := &queue{0, []target{{name, x}}}

	for !q.isEmpty() {
		t := q.pop()
		fmt.Printf("Display %s (%T):\n", t.name, t.x)
		display(t.name, reflect.ValueOf(t.x), q)
		if !q.isEmpty() {
			fmt.Println()
		}
	}
}

func formatAtom(v reflect.Value, q *queue) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		if v.Bool() {
			return "true"
		}
		return "false"
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr,
		reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	case reflect.Array, reflect.Struct:
		return q.push(v.Interface())
	default: // reflect.Interface
		return v.Type().String() + " value"
	}
}

func display(path string, v reflect.Value, q *queue) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i), q)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i), q)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path,
				formatAtom(key, q)), v.MapIndex(key), q)
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem(), q)
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem(), q)
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v, q))
	}
}
