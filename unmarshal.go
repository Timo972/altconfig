package altconfig

import "github.com/timo972/altconfig/internal"

func Unmarshal(data []byte, v interface{}) error {
	p := internal.NewParser(data)

	err := p.Tokenize()
	if err != nil {
		return err
	}

	return p.Parse(v)
}
