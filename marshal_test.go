package altconfig

import (
	"github.com/timo972/altconfig/schemes"
	"testing"
)

func TestMarshal(t *testing.T) {
	_, err := Marshal(schemes.ServerConfig{
		Name:        "alt:V Server",
		Slots:       4096,
		Description: "Default alt:V Server",
		Modules:     []string{"js-module", "csharp-module"},
		Resources:   []string{"resource1", "resource2"},
		Announce:    true,
		Voice: schemes.VoiceConfig{
			Bitrate:            64000,
			ExternalSecret:     "secret123",
			ExternalHost:       "xx.xx.xx.xx",
			ExternalPort:       7798,
			ExternalPublicHost: "xx.xx.xx.xx",
			ExternalPublicPort: 7799,
		},
	})

	if err != nil {
		t.Error(err)
	}
}
