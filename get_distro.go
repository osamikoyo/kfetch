package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getDistroInfo() (string, error) {
	content, err := ioutil.ReadFile("/etc/os-release")
	if err == nil {
		lines := strings.Split(string(content), "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "PRETTY_NAME=") {
				return strings.Trim(strings.TrimPrefix(line, "PRETTY_NAME="), "\""), nil
			}
		}
	}

	content, err = ioutil.ReadFile("/etc/lsb-release")
	if err == nil {
		lines := strings.Split(string(content), "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "DISTRIB_DESCRIPTION=") {
				return strings.Trim(strings.TrimPrefix(line, "DISTRIB_DESCRIPTION="), "\""), nil
			}
		}
	}

	content, err = ioutil.ReadFile("/etc/issue")
	if err == nil {
		return strings.TrimSpace(string(content)), nil
	}

	return "", fmt.Errorf("не удалось определить дистрибутив")
}
