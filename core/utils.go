package core

import (
	"gitlab.com/xiayesuifeng/v2rayxplus/conf"
	"os"
	"strconv"
)

func GetConfigName() (name, path string) {
	for i := 0; ; i++ {
		tmp := ""
		if i > 0 {
			tmp = strconv.FormatInt(int64(i), 10)
		}
		path := conf.ConfigPath + "/v2ray/config" + tmp + ".json"
		_, err := os.Stat(path)
		if err != nil && os.IsNotExist(err) {
			return "config" + tmp, path
		}
	}
}
