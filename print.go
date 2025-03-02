package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

const art = `
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⣀⣀⠀⠀⠀⢀⣠⣤⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⡀⠀⠿⠿⣿⣿⣷⣤⠈⠻⡿⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣷⡀⠀⠀⣎⣹⣿⣿⣷⣄⡀⠰⠂⣤⣤⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⣶⣤⡀⠀⠀⠹⣿⣿⣤⣾⣿⣿⣿⣿⣿⣿⣿⣄⠀⠙⠋⠀⠀⡀⠀⠀⠀
⠀⠀⠀⠀⡜⢙⠆⠀⠙⣿⣿⣆⠀⠀⠙⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡆⠀⣀⠀⠈⠁⠀⠀⠀
⠀⠀⠀⢠⡇⢸⠀⠀⠀⢸⣿⣿⡄⠀⠀⠈⣿⣿⡟⢿⣿⣿⣿⣿⣿⣿⣇⡄⠀⣀⠀⠀⠀⠀⠀
⠀⠀⠀⢸⡆⠘⡤⠀⠁⠉⠉⠉⠓⢄⠀⠀⠙⣿⣷⣀⣿⣿⣿⣿⣿⣋⣿⣿⢃⠻⠃⠀⠀⠀⠀
⠀⠀⠀⠈⢧⠀⠈⠀⠀⠀⠀⢰⣶⣌⢣⠀⠀⠈⠻⠿⠿⠿⣿⣿⣿⣿⣿⣿⣶⣿⣷⣄⠀⠀⠀
⠀⠀⠀⠀⢸⠃⠀⠀⠀⠀⠀⠈⣿⣿⠄⡇⠰⢆⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⡁⠀⠀⠀
⠀⠀⠀⠀⠈⡄⠀⠀⠀⠀⠀⠀⠈⠉⡰⠃⠀⠀⠀⠸⠟⢐⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⠀⠀⠀
⠀⠀⠀⠀⠀⠈⠢⣤⣀⡀⣠⣤⣶⣿⣄⡀⠀⠀⠀⣤⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⢸⣷⣿⣿⣿⣿⣿⣷⠀⠀⠀⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠟⠉⠀⠀⠀
⠀⠀⠀⠐⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠉⠙⠿⠿⡿⠟⠛⠛⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⢠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⠟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠙⠋⠿⣿⣿⣿⣿⣿⣿⣿⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠟⠋⠹⠟⠛⠻⡿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`

func PrintInfo(info string) {

	artLines := strings.Split(strings.TrimSpace(art), "\n")
	infoLines := strings.Split(strings.TrimSpace(info), "\n")

	maxArtWidth := 0
	for _, line := range artLines {
		if len(line) > maxArtWidth {
			maxArtWidth = len(line)
		}
	}
	for i := range artLines {
		artLines[i] = artLines[i] + strings.Repeat(" ", maxArtWidth-len(artLines[i]))
	}

	maxLines := len(artLines)
	if len(infoLines) > maxLines {
		maxLines = len(infoLines)
	}

	// Выравниваем количество строк
	for len(artLines) < maxLines {
		artLines = append(artLines, strings.Repeat(" ", maxArtWidth)) 
	}
	for len(infoLines) < maxLines {
		infoLines = append(infoLines, "") 
	}

	// Выводим ASCII-арт и текст построчно
	artColor := color.New(color.FgHiYellow).SprintFunc()
	textColor := color.New(color.FgHiYellow).SprintFunc()

// Выводим с цветами
for i := 0; i < maxLines; i++ {
    fmt.Printf("%s  %s\n", artColor(artLines[i]), textColor(infoLines[i]))
}
}