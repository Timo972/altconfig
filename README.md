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
voice: {
    port: 7789
    ip: 127.0.0.1
    external: true
}
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
	// possible converts: string, int, bool, map[string]interface{}, []interface{}
	port, ok := val.(int)
	if !ok {
		println("could not convert port to int")
	} else {
		println(port)
	}
	voice, err := cfg.Get("voice")
	if err != nil {
		panic(err)
	}
	voicePort := voice.(map[string]interface{})["port"]
	println(voicePort.(int))

	cfg.Set("test", "Hello World")

	// saves changes to the file you opened
	err = cfg.Save(false, false)
	if err != nil {
		panic(err)
	}
	
	// serializes into string
	content := cfg.Serialize(false, false)
	println(content)
}
```