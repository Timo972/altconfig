# cfg-reader -- alt-config bindings for go
Go package for reading alt-config format config files.
Used by alt:V Multiplayer.  
Maybe you want to use this package along the [alt:V Go module](https://github.com/shockdev04/altv-go-module).

## Usage
```js
// server.cfg
name: Test
port: 7788
announce: true
token: 'my token with special chars'
modules: [
    go-module
]
resources: [
    test
]
```
```go
// main.go
package main

import (
	"github.com/Timo972/cfg-reader"
)

func main() {
	cfg := cfg_reader.NewConfig("server.cfg")
	if cfg == nil {
		println("file not found")
		return
	}
	val, err := cfg.Get("port")
	if err != nil {
		panic(err)
	}
	// possible converts: string, int, bool, cfg_reader.Dict = map[string]interface{}, cfg_reader.List = []interface{}
	port, ok := val.(int)
	if ! ok {
		println("could not convert port to int")
	} else {
		println(port)
	}
}
```