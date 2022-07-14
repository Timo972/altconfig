# alt-config port in go
Go package for reading alt-config format config files which are used by alt:V Multiplayer.  
Maybe you want to use this package along the [alt:V Go module](https://github.com/timo972/altv-go).

[altMP alt-config (C++)](https://github.com/altmp/alt-config)

## Usage
```bash
# server.cfg
name: Test
port: 7788
announce: true
token: 'my token with special chars'
modules: [
    'go-module'
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
	"fmt"
	"github.com/timo972/altconfig/schemes"
	"io/ioutil"
	"github.com/timo972/altconfig"
)

const (
	ServerConfigFile = "server.cfg"
)

func main() {
	data, err := ioutil.ReadFile(ServerConfigFile)
	if err != nil {
		// Error if file not found or not permitted to read
		panic(err)
	}

	// schemes package contains predefined alt:V Config files (server.cfg, resource.cfg, stream.cfg voice.cfg, altv.cfg)
	var config schemes.ServerConfig
	err = altconfig.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	// Access server.cfg name field
	fmt.Printf("Server Name: %s\n", config.Name)

	// Set server port
	config.Port = 7787
	
	data, err = altconfig.Marshal(config)
	if err != nil {
		panic(err)
    }
	
	// Save server.cfg with port changed to 7787
	err = ioutil.WriteFile(ServerConfigFile, data, 0644)
	if err != nil {
		panic(err)
    }
}
```
