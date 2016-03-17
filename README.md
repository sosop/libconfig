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


