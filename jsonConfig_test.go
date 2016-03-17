package libconfig

import "testing"

type JConfig struct {
	AppName string `json:"appname"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

func TestParseJson(t *testing.T) {
	jConfig := &JConfig{}
	NewJsonConfig("jsonConfig.json", jConfig)
	if jConfig.AppName != "testJson" || jConfig.Host != "0.0.0.0" || jConfig.Port != 8888 {
		t.Fatal("parse json error")
	}
}
