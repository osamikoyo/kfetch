package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getMemoryInfoLinux() (total, used uint64, err error) {
	data, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка при чтении /proc/meminfo: %v", err)
	}

	lines := strings.Split(string(data), "\n")
	var memTotal, memAvailable uint64
	for _, line := range lines {
		if strings.HasPrefix(line, "MemTotal:") {
			memTotal = parseMemInfoLine(line)
		} else if strings.HasPrefix(line, "MemAvailable:") {
			memAvailable = parseMemInfoLine(line)
		}
	}

	if memTotal == 0 || memAvailable == 0 {
		return 0, 0, fmt.Errorf("не удалось получить информацию о памяти")
	}

	used = memTotal - memAvailable
	return memTotal, used, nil
}

func parseMemInfoLine(line string) uint64 {
	parts := strings.Fields(line)
	if len(parts) < 2 {
		return 0
	}
	value, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return 0
	}
	return value * 1024
}
