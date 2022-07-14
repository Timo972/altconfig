package altconfig

import "github.com/timo972/altconfig/internal"

func Marshal(v interface{}) ([]byte, error) {
	e := internal.NewEmitter()
	err := e.Emit(v, 0, true)
	if err != nil {
		return nil, err
	}

	return e.Buffer.Bytes(), nil
}
