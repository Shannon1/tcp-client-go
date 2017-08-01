package parse_config

import (
	"testing"
)


func TestParseConfig(t *testing.T) {
	filePath := "../config/config.yaml"
	config, err := ParseConfig(filePath)
	if err != nil {
		t.Error("Parse_Config error.")
		t.Error(err.Error())
		return
	}

	t.Log(config)


}
