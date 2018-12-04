package core

import (
	"gitlab.com/xiayesuifeng/v2rayxplus/conf"
	"io"
	"log"
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

func GetSu() string {
	su, err := exec.LookPath("kdesu")
	if err != nil {
		su, err = exec.LookPath("pkexec")
		if err != nil {
			return ""
		}
	}

	return su
}

func CopyConfigToEtc(config string) bool {
	su := GetSu()
	if su == "" {
		return false
	}

	cmd := exec.Command(su, "cp", conf.V2rayConfigPath+"/"+config+".json", "/etc/v2ray/config.json")
	_, err := cmd.CombinedOutput()
	return err == nil
}

func StartV2rayXPlusSerive(config string) bool {
	CopyConfigToEtc(config)
	return StartService("v2rayxplus")
}

func RestartV2rayXPlusSerive(config string) bool {
	CopyConfigToEtc(config)

	if !StopService("v2rayxplus") {
		return false
	}

	return StartService("v2rayxplus")
}

func StatusV2rayXPlusSerive() (exited, enable bool) {
	return StatusService("v2rayxplus")
}

func StopV2rayXPlusSerive() bool {
	return StopService("v2rayxplus")
}

func StartService(service string) bool {
	if os.Getuid() != 0 {
		su := GetSu()
		if su == "" {
			return false
		}

		return exec.Command(su, "systemctl", "start", service).Run() == nil
	} else {
		return exec.Command("systemctl", "start", service).Run() == nil
	}
}

func RestartService(service string) bool {
	if os.Getuid() != 0 {
		su := GetSu()
		if su == "" {
			return false
		}

		return exec.Command(su, "systemctl", "restart", service).Run() == nil
	} else {
		return exec.Command("systemctl", "restart", service).Run() == nil
	}
}

func StopService(service string) bool {
	if os.Getuid() != 0 {
		su := GetSu()
		if su == "" {
			return false
		}

		return exec.Command(su, "systemctl", "stop", service).Run() == nil
	} else {
		return exec.Command("systemctl", "stop", service).Run() == nil
	}
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
			exited = strings.Contains(line, "exited") || strings.Contains(line, "running")
		}
	}

	return exited, enable
}

func InitIpTables() {
	if exited, _ := StatusService("iptables"); !exited {
		if _, err := os.Stat("/etc/iptable/iptables.rules"); os.IsNotExist(err) {
			if _, err := os.Stat("/etc/iptable/empty.rules"); err == nil {
				CopyFile("/etc/iptable/empty.rules", "/etc/iptable/iptables.rules")
				StartService("iptables")
			}
		}
	}

	bytes, err := exec.Command("sysctl", "net.ipv4.ip_forward").CombinedOutput()
	result := strings.TrimSpace(string(bytes))
	if err == nil {
		if strings.HasSuffix(result, "0") {
			log.Println(string(bytes))
			exec.Command("sysctl", "net.ipv4.ip_forward=1").Run()
		}
	}
}
