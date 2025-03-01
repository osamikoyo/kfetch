package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getCPUInfoLinux() (string, error) {
	data, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		return "", fmt.Errorf("ошибка при чтении /proc/cpuinfo: %v", err)
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "model name") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1]), nil
			}
		}
	}

	return "", fmt.Errorf("информация о процессоре не найдена")
}
