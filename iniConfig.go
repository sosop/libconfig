package libconfig

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// IniConfig kv的配置文件
type IniConfig struct {
	filename string
	entry    map[string]interface{}
}

// NewIniConfig 构造器，读取文件
func NewIniConfig(filename string) *IniConfig {
	iniConfig := &IniConfig{filename, make(map[string]interface{}, 32)}
	iniConfig.parse()
	return iniConfig
}

// NewIniConfigAsReader 构造器，输入流
func NewIniConfigAsReader(reader io.Reader) *IniConfig {
	iniConfig := &IniConfig{entry: make(map[string]interface{}, 32)}
	iniConfig.parseReader(reader)
	return iniConfig
}

func (c *IniConfig) parse() {
	file, err := os.Open(c.filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	c.parseReader(file)
}

func (c *IniConfig) parseReader(reader io.Reader) {
	scan := bufio.NewScanner(reader)
	tagName := ""
	for scan.Scan() {
		line := scan.Text()
		n := len(line)
		if n > 0 && (line[0] == '[' && line[n-1] == ']') {
			tagName = line[1 : n-1]
			continue
		}
		if line = strings.TrimSpace(line); strings.Contains(line, "=") {
			kv := strings.SplitN(line, "=", 2)
			key := strings.TrimSpace(kv[0])
			if tagName != "" {
				key = tagName + "::" + key
			}
			c.entry[key] = strings.TrimSpace(kv[1])
		}
	}
	if err := scan.Err(); err != nil {
		panic(err)
	}
}

// GetString 获取字符串
func (c *IniConfig) GetString(key string, defaultValue ...string) string {
	if val, ok := c.entry[key]; ok {
		return val.(string)
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return ""
}

// GetBool 获取bool值
func (c *IniConfig) GetBool(key string, defaultValue ...bool) bool {
	if val, ok := c.entry[key]; ok {
		ret, err := strconv.ParseBool(val.(string))
		if err != nil {
			panic(err)
		}
		return ret
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return false
}

// GetInt 获取整型
func (c *IniConfig) GetInt(key string, defaultValue ...int) int {
	if val, ok := c.entry[key]; ok {
		ret, err := strconv.Atoi(val.(string))
		if err != nil {
			panic(err)
		}
		return ret
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return 0
}

// Set 设置
func (c *IniConfig) Set(key string, value interface{}) {
	c.entry[strings.TrimSpace(key)] = value
}
