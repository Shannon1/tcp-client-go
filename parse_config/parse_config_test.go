package parse_config

import (
	"testing"
)


func TestParse_config(t *testing.T) {
	filePath := "../config/config.yaml"
	config, err := Parse_config(filePath)
	if err != nil {
		t.Error("Parse_Config error.")
		t.Error(err.Error())
		return
	}

	t.Log(config)


}
