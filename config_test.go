package altconfig

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
)

const (
	configSampleName    = "alt-config.cfg"
	configSampleContent = `
name: Test,
port: 7788,
announce: true,
modules: [
	go
],
resources: [],
voice: {
	external: true,
	ip: 192.168.178.1,
	port: 7789,
	blacklist: [
		someone
	]
}
`
)

func writeConfigSample() {
	err := ioutil.WriteFile(configSampleName, []byte(configSampleContent), 0644)
	if err != nil {
		fmt.Errorf("could not create sample config file")
	}
}

func deleteConfigSample() {
	os.Remove(configSampleName)
}

func TestMain(m *testing.M) {
	writeConfigSample()
	code := m.Run()
	deleteConfigSample()
	os.Exit(code)
}

func TestNewConfig(t *testing.T) {
	cfg, err := New(configSampleName)
	if err != nil {
		t.Errorf(err.Error())
	}

	if cfg.Name != configSampleName {
		t.Errorf("got file name %v, want %v", cfg.Name, configSampleName)
	}
}

func TestConfig_Get(t *testing.T) {
	cfg, err := New(configSampleName)
	if err != nil {
		t.Errorf(err.Error())
	}

	val, err := cfg.Get("name")
	if err != nil {
		t.Errorf(err.Error())
	}

	str := val.(string)
	if str != "Test" {
		t.Errorf("got value %v, want Test", str)
	}
}

func TestConfig_GetBool(t *testing.T) {
	cfg, err := New(configSampleName)
	if err != nil {
		t.Errorf(err.Error())
	}

	val, err := cfg.GetBool("announce")
	if err != nil {
		t.Errorf(err.Error())
	}

	if val != true {
		t.Errorf("got value %v, want true", val)
	}
}

func TestConfig_GetInt(t *testing.T) {
	cfg, err := New(configSampleName)
	if err != nil {
		t.Errorf(err.Error())
	}

	val, err := cfg.GetInt("port")
	if err != nil {
		t.Errorf(err.Error())
	}

	if val != 7788 {
		t.Errorf("got value %v, want 7788", val)
	}
}

func TestConfig_GetString(t *testing.T) {
	cfg, err := New(configSampleName)
	if err != nil {
		t.Errorf(err.Error())
	}

	val, err := cfg.GetString("name")
	if err != nil {
		t.Errorf(val)
	}

	if val != "Test" {
		t.Errorf("got value %v, want Test", val)
	}
}

func TestConfig_GetList(t *testing.T) {
	cfg, err := New(configSampleName)
	if err != nil {
		t.Errorf(err.Error())
	}

	val, err := cfg.GetList("modules")
	if err != nil {
		t.Errorf(err.Error())
	}

	moduleOne, ok := val[0].(string)
	if !ok {
		t.Errorf("could not convert member of list")
	}

	if moduleOne != "go" {
		t.Errorf("got value %v, want go", moduleOne)
	}
}

func TestConfig_GetDict(t *testing.T) {
	cfg, err := New(configSampleName)
	if err != nil {
		t.Errorf(err.Error())
	}

	voice, err := cfg.GetDict("voice")
	if err != nil {
		t.Errorf(err.Error())
	}

	external, ok := voice["external"].(bool)
	if !ok {
		t.Errorf("could not convert member of dict")
	}

	if external != true {
		t.Errorf("got value %v, want true", external)
	}
}

func TestConfig_Set(t *testing.T) {
	cfg, err := New(configSampleName)
	if err != nil {
		t.Errorf(err.Error())
	}

	scenarios := []struct {
		key   string
		value interface{}
	}{
		{"newBool", true},
		{"newInt", 7},
		{"newString", "Hello World"},
		{"newDict", make(map[string]string)},
		{"newList", make([]string, 0)},
	}

	for _, pair := range scenarios {
		err = cfg.Set(pair.key, pair.value)
		if err != nil {
			t.Error(err.Error())
		}

		value, err := cfg.Get(pair.key)
		if err != nil {
			t.Errorf(err.Error())
		}

		switch value.(type) {
		case string:
			if v := value.(string); v != pair.value.(string) {
				t.Errorf("got value %v, want %v", v, pair.value.(string))
			}
		case int:
			if v := value.(int); v != pair.value.(int) {
				t.Errorf("got value %v, want %v", v, pair.value.(int))
			}
		case bool:
			if v := value.(bool); v != pair.value.(bool) {
				t.Errorf("got value %v, want %v", v, pair.value.(bool))
			}
		case map[string]interface{}:
		case []interface{}:
		case nil:
			t.Errorf("value not set for key %v (received nil value)", pair.key)
		default:
			t.Errorf("unkown type %v at key %v", reflect.TypeOf(value), pair.key)
		}
	}
}

func TestConfig_Serialize(t *testing.T) {
	cfg, err := New(configSampleName)
	if err != nil {
		t.Errorf(err.Error())
	}

	err = cfg.Set("serializeTest", true)
	if err != nil {
		t.Error(err.Error())
	}

	content, err := cfg.Serialize(false, false)
	if err != nil {
		t.Error(err.Error())
	}

	if !strings.Contains(content, "serializeTest: true") {
		t.Errorf("could not find newly setted value in serialization")
	}
}

func TestConfig_Save(t *testing.T) {
	cfg, err := New(configSampleName)
	if err != nil {
		t.Error(err.Error())
	}

	err = cfg.Set("saveTest", true)
	if err != nil {
		t.Error(err.Error())
	}

	err = cfg.Save(false, false)
	if err != nil {
		t.Error(err.Error())
	}

	bytes, err := ioutil.ReadFile(configSampleName)
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(string(bytes), "saveTest: true") {
		t.Errorf("could not find newly setted value in saved file")
	}
}
