package cfgreader

import "strconv"

type Type = uint8
type Scalar = string
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

type Node struct {
	Type Type
	Value interface{}
}

// NewNode create new node
func NewNode(val interface{}) *Node {
	switch val.(type) {
	case int:
	case uint:
	case float32:
	case float64:
	case bool:
		return &Node{Type: SCALAR, Value: val}
	case []int:
	case []uint:
	case []float32:
	case []float64:
	case []bool:
		values := val.([]interface{})
		nodes := make([]*Node, len(values))
		for i, value := range values {
			nodes[i] = NewNode(value)
		}
		return &Node{Type: LIST, Value: nodes}
	case map[string]int:
	case map[string]uint:
	case map[string]float32:
	case map[string]float64:
	case map[string]bool:
		values := val.(map[string]interface{})
		node := &Node{Type: DICT, Value: make(Dict)}
		for key, val := range values {
			node.Value.(map[string]interface{})[key] = NewNode(val)
		}
		return &Node{Type: DICT, Value: node}
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

func (n Node) ToBool() (bool, bool) {
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
}

// ToNumber convert node value to number
func (n Node) ToNumber() (int, bool) {
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