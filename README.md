### 主要针对配置文件解析
#### 一、解析ini文件

app.ini

```
appname = sosop_app
mode = dev



[prod]
db = mysql@172.78.66.88:3306/dbname
port = 80

[dev]
db = mysql@localhost:3306/dbname
port = 8080

[test]
db = mysql@172.78.66.28:3306/dbname
port = 8088
```

```
package main

import (
	"fmt"
	"github.com/sosop/libconfig"
)

func main() {
	iniConfig := libconfig.NewIniConfig("app.ini")
	appname := iniConfig.GetString("appname")
	mode := iniConfig.GetString("mode")
	devDB := iniConfig.GetString("dev::db")
	testPort := iniConfig.GetInt("test::port")
	fmt.Println(appname, mode, devDB, testPort)
}
```

输出：sosop_app dev mysql@localhost:3306/dbname 8088


#### 二、解析json

config.json

```
{
	"redisCluster": [
		{"host": "192.168.1.100", "port": 6379}, 
		{"host": "192.168.1.101", "port": 6380},
		{"host": "192.168.1.102", "port": 6381}],
	"dbCluster": [
		{"host": "172.20.10.8", "port": 3306}, 
		{"host": "172.20.10.9", "port": 3308},
		{"host": "172.20.10.10", "port": 3310}
	]
}
```


```
type StoreConfig struct {
	RedisCluster []struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"redisCluster`
	DbCluster []struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"dbCluster"`
}

func main() {
	storeConf := &StoreConfig{}
	libconfig.NewJsonConfig("config.json", storeConf)
	fmt.Println(*storeConf)
}
```
输出：{[{192.168.1.100 6379} {192.168.1.101 6380} {192.168.1.102 6381}] [{172.20.10.8 3306} {172.20.10.9 3308} {172.20.10.10 3310}]}


