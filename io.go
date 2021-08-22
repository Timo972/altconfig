package cfg_reader

import "io/ioutil"

func ReadFile(fileName string) (*Node, error) {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	content := string(data)

	parser := NewParser(content)
	node, err := parser.Parse()

	if err != nil {
		return nil, err
	}

	return node, err
}

func WriteFile(fileName string, content string) error {
	err := ioutil.WriteFile(fileName, []byte(content), 0644)
	return err
}