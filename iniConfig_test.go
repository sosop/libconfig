package libconfig

import "testing"

var iniConf *IniConfig

func TestParseIni(t *testing.T) {
	iniConf = NewIniConfig("iniConfig.conf")
}

func TestSet(t *testing.T) {
	iniConf.Set("testStr", "string")
	iniConf.Set("testInt", "1")
	iniConf.Set("testBool", "true")
}

func TestGet(t *testing.T) {
	if iniConf.GetString("testStr") != "string" {
		t.Fatal("error get string")
	}
	if iniConf.GetInt("testInt") != 1 {
		t.Fatal("error get int")
	}
	if iniConf.GetBool("testBool") != true {
		t.Fatal("error get bool")
	}
	if iniConf.GetString("appname") != "test" {
		t.Fatal("error get string")
	}
}
