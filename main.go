package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/fatih/color"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	distro, err := getDistroInfo()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	gpu, err := getGPUInfoLinux()
	if err != nil{
		fmt.Println("Ошибка:",err)
	}

	cpu, err := getCPUInfoLinux()
	if err != nil{
		fmt.Println("Ошибка:", err)
	}


	shell, err := getShellFromEnv()
	if err != nil{
		fmt.Println("Ошибка:", err)
	}

	total, used, err := getMemoryInfoLinux()
	if err != nil{
		fmt.Println("Ошибка:", err)
	}

	de, err := getDesktopEnvironment()
	if err != nil{
		fmt.Println("Ошибка:", err)
	}

	artColor := color.New(color.FgHiMagenta).SprintFunc()
	bold := color.New(color.Bold).SprintlnFunc()

	version := runtime.Version()

	text := fmt.Sprintf(`
%s
%s

%s: %s
%s: %s
%s: %s
%s: %s
%s: %s
%s: %s
%s: %d/%d
%s: %s
`, artColor(bold(hostname)), strings.Repeat("-", 30), artColor("Distro"), distro, artColor("Kernel"),GetKernel(), 
artColor("Gpu"),gpu, artColor("Cpu"), cpu, artColor("De"),de, artColor("Shell") ,shell, artColor("Memory") ,used/1024/1024, total/1024/1024,
artColor("Go version"), version)

	PrintInfo(artColor(text))
}