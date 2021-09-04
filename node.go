package cfgreader

import (
	"strconv"
)

// Type enum
type Type = uint8
// Scalar default type
type Scalar = string
// List default type
type List = []*Node
// Dict default type
type Dict = map[string]*Node

const (
	// NONE node type
	NONE Type = iota
	// SCALAR node type
	SCALAR
	// LIST node type
	LIST
	// DICT node type
	DICT
)

// Node struct
type Node struct {
	Type Type
	Value interface{}
}

// NewNode create new node
func NewNode(val interface{}) *Node {
	switch val.(type) {
	case string:
		return &Node{Type: SCALAR, Value: val.(string)}
	case bool:
		return &Node{Type: SCALAR, Value: strconv.FormatBool(val.(bool))}
	case int, int8, int16, int32, int64:
		var value string
		if v, ok := val.(int); ok {
			value = strconv.FormatInt(int64(v), 10)
		} else if v, ok := val.(int8); ok {
			value = strconv.FormatInt(int64(v), 10)
		} else if v, ok := val.(int16); ok {
			value = strconv.FormatInt(int64(v), 10)
		} else if v, ok := val.(int32); ok {
			value = strconv.FormatInt(int64(v), 10)
		} else if v, ok := val.(int64); ok {
			value = strconv.FormatInt(v, 10)
		}
		return &Node{Type: SCALAR, Value: value}
	case uint, uint8, uint16, uint32, uint64:
		var value string
		if v, ok := val.(uint); ok {
			value = strconv.FormatUint(uint64(v), 10)
		} else if v, ok := val.(uint8); ok {
			value = strconv.FormatUint(uint64(v), 10)
		} else if v, ok := val.(uint16); ok {
			value = strconv.FormatUint(uint64(v), 10)
		} else if v, ok := val.(uint32); ok {
			value = strconv.FormatUint(uint64(v), 10)
		} else if v, ok := val.(uint64); ok {
			value = strconv.FormatUint(v, 10)
		}
		return &Node{Type: SCALAR, Value: value}
	case float32, float64:
		var float float64
		var bitSize int
		if v, ok := val.(float32); ok {
			float = float64(v)
			bitSize = 32
		} else if v, ok := val.(float64); ok {
			float = v
			bitSize = 64
		}
		return &Node{Type: SCALAR, Value: strconv.FormatFloat(float, 'E', -1, bitSize)}
	case []string:
		values := val.([]string)
		nodes := make([]*Node, len(values))
		for i, value := range values {
			nodes[i] = NewNode(value)
		}
		return &Node{Type: LIST, Value: nodes}
	case map[string]string:
		values := val.(map[string]string)
		node := &Node{Type: DICT, Value: make(Dict)}
		for key, val := range values {
			node.Value.(Dict)[key] = NewNode(val)
		}
		return node
	}
	return nil
}

// IsNone check if node value is nil
func (n Node) IsNone() bool {
	return n.Type == NONE
}

// IsScalar check if node value is string
func (n Node) IsScalar() bool {
	return n.Type == SCALAR
}

// IsList check if node value is list
func (n Node) IsList() bool {
	return n.Type == LIST
}

// IsDict check if node value is dict
func (n Node) IsDict() bool {
	return n.Type == DICT
}

//TODO revert node converts when there is an actual fix for NewNode

// ToBool convert node value to bool
func (n Node) ToBool() (bool, bool) {
	//bool, ok := n.Value.(bool)
	//if !ok {
		val, ok := n.Value.(string)
		if !ok {
			return false, ok
		}
		bool := false
		if val == "true" || val == "yes" {
			bool = true
		} else if val == "false" || val == "no" {
			bool = false
		} else {
			ok = false
		}
		return bool, ok
	//}
	//return bool, ok
}

// ToNumber convert node value to number
func (n Node) ToNumber() (int, bool) {
	val, ok := n.Value.(int)
	if !ok {
		str, ok := n.Value.(string)
		if !ok {
			return 0, ok
		}
		val, err := strconv.Atoi(str)
		if err != nil {
			return 0, false
		}
		return val, ok
	}
	return val, ok
}

// ToString convert node value to string
func (n Node) ToString() (string, bool) {
	val, ok := n.Value.(string)
	return val, ok
}

// ToDict convert node value to dict
func (n Node) ToDict() (Dict, bool) {
	val, ok := n.Value.(Dict)
	return val, ok
}

// ToList convert node value to list
func (n Node) ToList() (List, bool) {
	val, ok := n.Value.(List)
	return val, ok
}

// At value at idx
func (n Node) At(idx uint) *Node{
	val, ok := n.Value.(List)
	if !ok {
		return nil
	}
	return val[idx]
}

// Get value from key
func (n Node) Get(key string) *Node {
	val, ok := n.Value.(Dict)
	if !ok {
		return nil
	}
	return val[key]
}