package core

import (
	"errors"
	"gitlab.com/xiayesuifeng/v2rayxplus/conf"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetVension() (string, error) {
	cmd := exec.Command("v2ray", "-version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.New("v2ray not exist")
	}

	str := string(out)
	if !strings.Contains(str, "V2Ray") {
		return "", errors.New("v2ray not exist")
	}

	strs := strings.Split(str, " ")

	if len(strs) < 2 {
		return "", errors.New("v2ray not exist")
	}

	return strs[1], nil
}

func StartV2ray(config string) error {
	if _, err := os.Stat(config); err != nil {
		return errors.New("config file not found")
	}

	if err := CopyFile(config, "/etc/v2ray/config.json"); err != nil {
		return err
	}

	if exited, _ := StatusService("v2ray"); exited {
		RestartService("v2ray")
	} else {
		StartSerive("v2ray")
	}

	InitIpTables()

	if err := AddIpTablesRules(); err != nil {
		return err
	}

	return nil
}

func AddIpTablesRules() error {
	sh := `iptables -t nat -N V2RAY
iptables -t nat -A V2RAY -d 127.0.0.0/8 -j RETURN 
iptables -t nat -A V2RAY -d 192.168.0.0/16 -j RETURN 
iptables -t nat -A V2RAY -p tcp -j RETURN -m mark --mark 0xff
iptables -t nat -A V2RAY -p tcp -j REDIRECT --to-ports {port}
iptables -t nat -A PREROUTING -p tcp -j V2RAY
iptables -t nat -A OUTPUT -p tcp -j V2RAY
ip rule add fwmark 1 table 100
ip route add local 0.0.0.0/0 dev lo table 100
iptables -t mangle -N V2RAY_MASK
iptables -t mangle -A V2RAY_MASK -d 127.0.0.0/8 -j RETURN
iptables -t mangle -A V2RAY_MASK -d 192.168.0.0/16 -j RETURN
iptables -t mangle -A V2RAY_MASK -p udp -j TPROXY --on-port {port} --tproxy-mark 1
iptables -t mangle -A PREROUTING -p udp -j V2RAY_MASK`
	sh = strings.Replace(sh, "{port}", strconv.FormatInt(int64(conf.Conf.ListerPort), 10), -1)
	cmd := exec.Command("bash", "-c", sh)
	return cmd.Run()
}

func RemoveIpTablesRules() error {
	sh := `iptables -t nat -F V2RAY
iptables -t nat -D PREROUTING -p tcp -j V2RAY
iptables -t nat -D OUTPUT -p tcp -j V2RAY
iptables -t nat -X V2RAY
iptables -t mangle -F V2RAY_MASK
iptables -t mangle -D PREROUTING -p udp -j V2RAY_MASK
iptables -t mangle -X V2RAY_MASK
ip rule del fwmark 1 table 100
ip route del local 0.0.0.0/0 dev lo table 100`
	cmd := exec.Command("bash", "-c", sh)
	return cmd.Run()
}
