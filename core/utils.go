package core

import (
	"gitlab.com/xiayesuifeng/v2rayxplus/conf"
	"os"
	"strconv"
)

func GetConfigName() string {
	for i := 0; ; i++ {
		tmp := ""
		if i > 0 {
			tmp = strconv.FormatInt(int64(i), 10)
		}
		name := conf.ConfigPath + "/config" + tmp + ".json"
		_, err := os.Stat(name)
		if err != nil && os.IsNotExist(err) {
			return name
		}
	}
}
