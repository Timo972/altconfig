package internal

import "io/ioutil"

// ReadFile into node
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

// WriteFile write file with content
func WriteFile(fileName string, content []byte) error {
	err := ioutil.WriteFile(fileName, content, 0644)
	return err
}
