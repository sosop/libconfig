package libconfig

import (
	"encoding/xml"
	"testing"
)

type XConfig struct {
	XMLName xml.Name `xml:"config"`
	AppName string   `xml:"appname"`
	Host    string   `xml:"host"`
	Port    int      `xml:"port"`
}

func TestParseXml(t *testing.T) {
	xConfig := &XConfig{}
	NewXmlConfig("xmlConfig.xml", xConfig)
	if xConfig.AppName != "testXml" || xConfig.Host != "0.0.0.0" || xConfig.Port != 8888 {
		t.Fatal("parse xml error")
	}
}
