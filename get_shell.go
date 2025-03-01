package main

import (
	"fmt"
	"os"
	"path/filepath"
)
func getShellFromEnv() (string, error) {
	// Получаем значение переменной окружения SHELL
	shellPath := os.Getenv("SHELL")
	if shellPath == "" {
		return "", fmt.Errorf("переменная SHELL не установлена")
	}

	// Извлекаем имя оболочки из пути
	shell := filepath.Base(shellPath)
	return shell, nil
}
