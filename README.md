# alt-config port in go
Go package for reading alt-config format config files.
Used by alt:V Multiplayer.  
Maybe you want to use this package along the [alt:V Go module](https://github.com/shockdev04/altv-go-module).

[altMP alt-config (C++)](https://github.com/altmp/alt-config)

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
	"github.com/Timo972/altconfig"
)

func main() {
	cfg, err := altconfig.NewConfig("server.cfg")
	if err != nil {
		// Error if file not found or not permitted to read
		panic(err)
	}
	
	// Get a value
	val, err := cfg.Get("port")
	if err != nil {
		panic(err)
	}
	// Possible converts: string, int, bool, map[string]interface{}, []interface{}
	port, ok := val.(int)
	if !ok {
		println("could not convert port to int")
	} else {
		println(port)
	}
	
	// Get a value of type (Config_GetInt, Config_GetBool, Config_GetString, Config_GetDict, Config_GetList)
	portValue, err := cfg.GetInt("port")
	if err != nil {
		// Error if key is not found or value is not of the type you want to get
		panic(err)
    }
	
	// Use portValue here
	println(portValue)
	
	voice, err := cfg.Get("voice")
	if err != nil {
		panic(err)
	}
	voicePort := voice.(map[string]interface{})["port"]
	println(voicePort.(int))

	cfg.Set("test", "Hello World")

	// Save changes to the file you opened
	err = cfg.Save(false, false)
	if err != nil {
		panic(err)
	}
	
	// Serializes into string
	content := cfg.Serialize(false, false)
	println(content)
}
```