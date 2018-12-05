package config

import (
	"testing"
)

func Test_Config(t *testing.T) {
	c, err := LoadConfigWithDefaults("../display.yml")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", c)
}
