package internal

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

func GetTagName(t reflect.StructField) string {
	name := t.Tag.Get("alt")
	if name == "" {
		name = t.Name
	}

	return name
}

// Emitter struct
type Emitter struct {
	Buffer *bytes.Buffer
}

// NewEmitter Create new emitter instance
func NewEmitter() *Emitter {
	return &Emitter{Buffer: bytes.NewBuffer(make([]byte, 0))}
}

func (e *Emitter) writeValue(scalar string, isLast bool) {
	specials := strings.Count(scalar, " ")

	if specials > 0 {
		e.Buffer.WriteString("'" + scalar + "'")
	} else {
		e.Buffer.WriteString(scalar)
	}

	if !isLast {
		e.Buffer.WriteString(",\n")
	} else {
		e.Buffer.WriteRune('\n')
	}
}

// Emit Serialize node
func (e *Emitter) Emit(v interface{}, indent uint, isLast bool) error {
	_indent := strings.Repeat(" ", int(indent*2))

	rt := reflect.TypeOf(v)
	rv := reflect.ValueOf(v)
	kind := rt.Kind()
	if kind == reflect.Pointer {
		rt = rt.Elem()
		rv = rv.Elem()
		kind = rt.Kind()
	}

	if kind == reflect.String {
		str := Escape(rv.String())
		e.writeValue(str, isLast)
	} else if kind == reflect.Bool {
		str := fmt.Sprintf("%v", rv.Bool())
		e.writeValue(str, isLast)
	} else if kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 || kind == reflect.Int32 || kind == reflect.Int64 {
		str := fmt.Sprintf("%v", rv.Int())
		e.writeValue(str, isLast)
	} else if kind == reflect.Uint || kind == reflect.Uint8 || kind == reflect.Uint16 || kind == reflect.Uint32 || kind == reflect.Uint64 {
		str := fmt.Sprintf("%v", rv.Uint())
		e.writeValue(str, isLast)
	} else if kind == reflect.Float32 || kind == reflect.Float64 {
		str := fmt.Sprintf("%v", rv.Float())
		e.writeValue(str, isLast)
	} else if kind == reflect.Slice || kind == reflect.Array {
		e.Buffer.WriteString("[\n")

		l := rv.Len()
		for i := 0; i < l; i++ {
			e.Buffer.WriteString(_indent)
			err := e.Emit(rv.Index(i).Interface(), indent+1, i >= l-1)
			if err != nil {
				return err
			}
		}

		e.Buffer.WriteString(strings.Repeat(" ", int((indent-1)*2)))
		e.Buffer.WriteString("]")

		if !isLast {
			e.Buffer.WriteString(",\n")
		} else {
			e.Buffer.WriteRune('\n')
		}
	} else if kind == reflect.Struct {
		if indent > 0 {
			e.Buffer.WriteString("{\n")
		}

		fc := rv.NumField()
		for i := 0; i < fc; i++ {
			field := rv.Field(i)
			ft := rt.Field(i)
			key := GetTagName(ft)

			if !ft.IsExported() || field.IsZero() {
				continue
			}

			e.Buffer.WriteString(_indent + key + ": ")
			err := e.Emit(field.Interface(), indent+1, i >= fc-1)
			if err != nil {
				return err
			}
		}

		if indent > 0 {
			e.Buffer.WriteString(strings.Repeat(" ", int((indent-1)*2)))
			e.Buffer.WriteRune('}')
		}

		if !isLast {
			e.Buffer.WriteString(",\n")
		}
	}

	return nil
}
