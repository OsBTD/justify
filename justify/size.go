package module

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetSize() (int, int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	size := strings.Split(strings.TrimSpace(string(output)), " ")
	TerminalWidth, err := strconv.Atoi(size[1])
	if err != nil {
		fmt.Println("Invalid terminal width")
	}
	fmt.Println(TerminalWidth)
	content, err := os.ReadFile("standard.txt")
	if err != nil {
		log.Fatal("Error : couldn't read file")
	}
	Ascii := strings.Split(string(content), "\n")
	var Char int = 32
	input := os.Args[1]
	Replace := map[rune][]string{}

	for i := 0; i < len(Ascii); i += 9 {
		if i+9 <= len(Ascii)-1 {
			Replace[rune(Char)] = Ascii[i+1 : i+9]
		}
		if Char <= 126 {
			Char++
		}

	}
	inputsplit := strings.Split(input, "\\n")
	var length int = 0
	var countspace int
	for _, line := range inputsplit {
		if line == "" {
			fmt.Println()
			continue
		}
		for i := 0; i < 1; i++ {
			for j := 0; j < len(line); j++ {
				inputrune := rune(line[j])
				length += len(Replace[inputrune][i])
				if inputrune == 32 {
					countspace++
				}

			}
			length = length - countspace*6
			fmt.Println()

		}
	}
	return length, countspace
}
