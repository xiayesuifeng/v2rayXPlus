package core

import (
	"gitlab.com/xiayesuifeng/v2rayxplus/conf"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func CopyFile(src, target string) error {
	srcConfig, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcConfig.Close()

	targetConfig, err := os.OpenFile(target, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer targetConfig.Close()

	_, err = io.Copy(targetConfig, srcConfig)
	if err != nil {
		return err
	}

	return nil
}

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

func StartSerive(service string) bool {
	return exec.Command("systemctl", "start", service).Run() == nil
}

func RestartService(service string) bool {
	return exec.Command("systemctl", "restart", service).Run() == nil
}

func StopService(service string) bool {
	return exec.Command("systemctl", "stop", service).Run() == nil
}

func StatusService(service string) (exited, enable bool) {
	bytes, err := exec.Command("systemctl", "status", service).CombinedOutput()
	if err != nil {
		return exited, enable
	}

	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Loaded:") {
			tmp := strings.Split(line, ";")
			if len(tmp) > 1 {
				if strings.Contains(tmp[1], "enable") {
					enable = true
				}
			}
		} else if strings.HasPrefix(line, "Active:") {
			exited = strings.Contains(line, "exited")
		}
	}

	return exited, enable
}
