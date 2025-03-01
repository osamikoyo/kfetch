package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func getDesktopEnvironment() (string,error) {
		// Запускаем команду ps для получения списка процессов
		cmd := exec.Command("ps", "-e")
		output, err := cmd.Output()
		if err != nil {
			return "", fmt.Errorf("ошибка при выполнении ps: %v", err)
		}
	
		// Анализируем вывод
		lines := strings.Split(string(output), "\n")
		for _, line := range lines {
			if strings.Contains(line, "gnome-shell") {
				return "GNOME", nil
			} else if strings.Contains(line, "plasmashell") {
				return "KDE Plasma", nil
			} else if strings.Contains(line, "xfce4-session") {
				return "XFCE", nil
			} else if strings.Contains(line, "cinnamon-session") {
				return "Cinnamon", nil
			} else if strings.Contains(line, "mate-session") {
				return "MATE", nil
			} else if strings.Contains(line, "lxqt-session") {
				return "LXQt", nil
			} else if strings.Contains(line, "budgie-panel") {
				return "Budgie", nil
			}
		}
	
		return "", fmt.Errorf("не удалось определить окружение рабочего стола")
}