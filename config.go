package cfgreader

import (
	"errors"
	"fmt"
	"reflect"
)

// Config config struct
type Config struct {
	Node *Node
	Name string
}

// NewConfig create config
func NewConfig(file string) (*Config, error) {
	node, err := ReadFile(file)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	config.Node = node
	config.Name = file

	return config, nil
}

// Convert node to go value
func Convert(node *Node) interface{} {
	for i := 0; i < 5; i++ {
		switch i {
		case 0:
			bool, ok := node.ToBool()
			if ok {
				return bool
			}
		case 1:
			int, ok := node.ToNumber()
			if ok {
				return int
			}
		case 2:
			string, ok := node.ToString()
			if ok {
				return string
			}
		case 3:
			list, ok := node.ToList()
			if ok {
				values := make([]interface{}, 0)
				for _, node := range list {
					val := Convert(node)
					values = append(values, val)
				}
				return values
			}
		case 4:
			dict, ok := node.ToDict()
			if ok {
				values := make(map[string]interface{})
				for key, node := range dict {
					val := Convert(node)
					values[key] = val
				}
				return values
			}
		}
	}
	return nil
}

// Get value at key
func (c *Config) Get(key string) (interface{}, error) {

	dict, ok := c.Node.ToDict()
	if !ok {
		return nil, errors.New("could not convert to dict")
	}

	node := dict[key]

	if node == nil {
		return nil, fmt.Errorf("key %v not found", key)
	}

	value := Convert(node)

	return value, nil
}

func (c *Config) GetString(key string) (string, error) {
	val, err := c.Get(key)
	if err != nil {
		return "", err
	}

	str, ok := val.(string)
	if !ok {
		return "", errors.New("could not convert to string")
	}

	return str, nil
}

func (c *Config) GetInt(key string) (int, error) {
	val, err := c.Get(key)
	if err != nil {
		return 0, err
	}

	i, ok := val.(int)
	if !ok {
		return 0, errors.New("could not convert to int")
	}

	return i, nil
}

func (c *Config) GetBool(key string) (bool, error) {
	val, err := c.Get(key)
	if err != nil {
		return false, err
	}

	b, ok := val.(bool)
	if !ok {
		return false, errors.New("could not convert to bool")
	}

	return b, nil
}

func (c *Config) GetDict(key string) (map[string]interface{}, error) {
	val, err := c.Get(key)
	if err != nil {
		return nil, err
	}

	dict, ok := val.(map[string]interface{})
	if !ok {
		return nil, errors.New("could not convert to dict")
	}

	return dict, nil
}

func (c *Config) GetList(key string) ([]interface{}, error) {
	val, err := c.Get(key)
	if err != nil {
		return nil, err
	}

	list, ok := val.([]interface{})
	if !ok {
		return nil, errors.New("could not convert to list")
	}

	return list, nil
}

// Set value at key
func (c *Config) Set(key string, value interface{}) error {
	node := NewNode(value)
	if node == nil {
		return fmt.Errorf("unsupported value type %v", reflect.TypeOf(value))
	}
	dict := c.Node.Value.(Dict)
	dict[key] = node
	//c.Node.Value = dict

	return nil
}

// Serialize config to string
func (c *Config) Serialize(useCommas bool, useApostrophes bool) string {
	emitter := NewEmitter()
	emitter.Emit(c.Node, 0, true, useCommas, useApostrophes)
	return emitter.String()
}

// Save config to file
func (c *Config) Save(useCommas bool, useApostrophes bool) error {
	emitter := NewEmitter()
	emitter.Emit(c.Node, 0, true, useCommas, useApostrophes)
	err := WriteFile(c.Name, emitter.String())
	return err
}