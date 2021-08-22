package cfg_reader

import "errors"

type Config struct {
	Node *Node
	Name string
}

func NewConfig(file string) *Config {
	node, err := ReadFile(file)
	if err != nil {
		panic(err)
	}
	config := &Config{}
	config.Node = node
	config.Name = file
	return config
}

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

func (c *Config) Get(key string) (interface{}, error) {

	dict, ok := c.Node.ToDict()
	if !ok {
		return nil, errors.New("could not convert to dict")
	}

	node := dict[key]

	if node == nil {
		return nil, errors.New("key not found")
	}

	value := Convert(node)

	return value, nil
}

func (c *Config) Set(key string, value interface{}) {
	node := NewNode(value)
	dict := c.Node.Value.(Dict)
	dict[key] = node
	println(c.Node.Value.(Dict))
}

func (c *Config) Serialize(useCommas bool, useApostrophes bool) string {
	emitter := NewEmitter()
	emitter.Emit(c.Node, 0, true, useCommas, useApostrophes)
	return emitter.String()
}

func (c *Config) Save(useCommas bool, useApostrophes bool) error {
	emitter := NewEmitter()
	emitter.Emit(c.Node, 0, true, useCommas, useApostrophes)
	err := WriteFile(c.Name, emitter.String())
	return err
}