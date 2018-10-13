package core

import (
	"errors"
	"os/exec"
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
