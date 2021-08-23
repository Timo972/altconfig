package cfgreader

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
func WriteFile(fileName string, content string) error {
	err := ioutil.WriteFile(fileName, []byte(content), 0644)
	return err
}