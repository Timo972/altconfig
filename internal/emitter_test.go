package internal

import (
	"testing"
)

func TestNewEmitter(t *testing.T) {
	e := NewEmitter()

	if e == nil {
		t.Error("NewEmitter returned nil")
	}
}

func TestEmitter_writeValue(t *testing.T) {
	scenarios := []struct {
		value  string
		isLast bool
		buffer string
	}{
		{"scalarvalue", true, "scalarvalue\n"},
		{"sv", false, "sv,\n"},
	}

	for _, v := range scenarios {
		e := NewEmitter()

		e.writeValue(v.value, v.isLast)

		buf := e.Buffer.String()
		if buf != v.buffer {
			t.Errorf("Emitter_writeValue was incorrect, got: '%s', want: '%s'", buf, v.buffer)
		}
	}
}

// TODO: cover structs & slices
func TestEmitter_Emit(t *testing.T) {
	scenarios := []struct {
		input  interface{}
		indent uint
		isLast bool
		buffer string
	}{
		// indent does not matter for scalars
		{"string", 0, false, "string,\n"},
		{"string", 0, true, "string\n"},
		{"str ing", 0, false, "'str ing',\n"},
		{"str ing", 0, true, "'str ing'\n"},
		{true, 0, false, "true,\n"},
		{true, 0, true, "true\n"},
		{false, 0, false, "false,\n"},
		{false, 0, true, "false\n"},
		{2.3, 0, false, "2.3,\n"},
		{2.3, 0, true, "2.3\n"},
		{42, 0, false, "42,\n"},
		{42, 0, true, "42\n"},
	}

	for _, v := range scenarios {
		e := NewEmitter()

		err := e.Emit(v.input, v.indent, v.isLast)
		if err != nil {
			t.Error(err)
		}

		buf := e.Buffer.String()
		if buf != v.buffer {
			t.Errorf("Emitter_Emit was incorrect, got: '%s', want: '%s'", buf, v.buffer)
		}
	}
}
