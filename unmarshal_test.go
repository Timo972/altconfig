package altconfig

import (
	"github.com/timo972/altconfig/schemes"
	"testing"
)

var (
	TestRaw = []byte(`
name: 'alt:V Server',
players: 4096,
description: 'Default alt:V Server',
announce: true,
modules: [
  js-module,
  csharp-module
],
resources: [
  resource1,
  resource2
],
voice: {
  bitrate: 64000,
  externalSecret: secret123,
  externalHost: xx.xx.xx.xx,
  externalPort: 7798,
  externalPublicHost: xx.xx.xx.xx,
  externalPublicPort: 7799
},
Test: true

`)
)

func TestUnmarshal(t *testing.T) {
	var conf schemes.ServerConfig

	err := Unmarshal(TestRaw, &conf)

	if err != nil {
		t.Error(err)
	}
}
