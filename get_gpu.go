package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func extractGPUModel(input string) (string, error) {
	re := regexp.MustCompile(`\[(.*?)\]`)
	matches := re.FindStringSubmatch(input)
	if len(matches) < 2 {
		return "", fmt.Errorf("не удалось извлечь модель GPU")
	}
	return matches[1], nil
}


func getGPUInfoLinux() (string, error) {
	cmd := exec.Command("lspci")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении lspci: %v", err)
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "VGA compatible controller") || strings.Contains(line, "3D controller") {
			return extractGPUModel(line)
		}
	}

	return "", fmt.Errorf("GPU не найдено")
}