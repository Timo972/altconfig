package cfg_reader

import "errors"

type Config struct {
	Node Dict
}

func NewConfig(file string) *Config {
	node, err := ReadFile(file)
	if err != nil {
		panic(err)
	}
	config := &Config{}
	dict, ok := node.ToDict()
	if !ok {
		println("could not convert to dict")
		return nil
	}
	config.Node = dict
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

func (c Config) Get(key string) (interface{}, error) {
	node := c.Node[key]

	if node == nil {
		return nil, errors.New("key not found")
	}

	value := Convert(node)

	return value, nil
}

func (c Config) Set(key string, value interface{}) {
	c.Node[key] = NewNode(value)
}
