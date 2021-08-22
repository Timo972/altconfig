package cfg_reader

import (
	"errors"
	"strings"
)

type Emitter struct {
	Stream string
}

// Create new emitter instance
func NewEmitter() *Emitter {
	return &Emitter{Stream: ""}
}

// Serialize node
func (e *Emitter) Emit(node *Node, indent uint, isLast bool, useCommas bool, useApostrophes bool) error {
	_indent := strings.Repeat(" ", int(indent * 2))

	if node.IsScalar() {
		str, ok := node.ToString()
		if !ok {
			return errors.New("failed to emit")
		}
		specialChars := strings.Count(str, " ")
		if useApostrophes || specialChars > 0 {
			e.Stream += "'" + Escape(str) + "'\n"
		} else {
			e.Stream += Escape(str) + "\n"
		}
	} else if node.IsList() {
		e.Stream += "[\n"

		list, ok := node.ToList()
		if !ok {
			return errors.New("failed to emit list")
		}
		end := len(list) - 1
		for i, node := range list {
			e.Stream += _indent
			e.Emit(node, indent + 1, i == end, useCommas, useApostrophes)
		}
		e.Stream += strings.Repeat(" ", int((indent - 1) * 2))
		if isLast || !useCommas {
			e.Stream += "]\n"
		} else {
			e.Stream += "],\n"
		}
	} else if node.IsDict() {
		if indent > 0 {
			e.Stream += "{\n"
		}

		dict, ok := node.ToDict()
		if !ok {
			return errors.New("failed to emit dict")
		}
		i := 0
		endIdx := len(dict) - 1
		for key, node := range dict {
			if node == nil || node.IsNone() {
				continue
			}
			e.Stream += _indent + key + ":"
			e.Emit(node, indent + 1, i == endIdx, useCommas, useApostrophes)
			i++
		}

		if indent > 0 {
			e.Stream += strings.Repeat(" ", int((indent - 1) * 2))
			if isLast || !useCommas {
				e.Stream += "}\n"
			} else {
				e.Stream += "},\n"
			}
		}
	}
	return nil
}

// Get the serialized string
func (e *Emitter) String() string {
	return e.Stream
}