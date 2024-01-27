package ascii

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	asciiStart      = ' '
	asciiEnd        = '~'
	bannerHeight    = 9 
	invalidPosition = -1
)

var asciiPositions = generateAsciiPositions()

func generateAsciiPositions() map[rune]int {
	positions := make(map[rune]int)
	asciiRange := asciiEnd - asciiStart + 1
	for i := 0; i < int(asciiRange); i++ {
		positions[rune(asciiStart)+rune(i)] = 1 + (bannerHeight * i)
	}
	return positions
}

func checkAsciiPos(symb rune) int {
	if pos, exists := asciiPositions[symb]; exists {
		return pos
	}
	return invalidPosition
}

func PrintAscii(banner, inputTxt string) (string, error) {
	splitBanner := strings.Split(banner, "\n")
	lines := strings.Split(inputTxt, "\\n")

	var resultBuilder strings.Builder

	for _, line := range lines {
		for row := 0; row < bannerHeight; row++ {
			for _, char := range line {
				charPosition := checkAsciiPos(char)
				if charPosition == invalidPosition {
					return "", fmt.Errorf("invalid character: %c", char)
				}
				resultBuilder.WriteString(splitBanner[charPosition+row])
			}
			if row != bannerHeight-1 {
				resultBuilder.WriteString("\n")
			}
		}
	}

	return resultBuilder.String(), nil
}

func LoadBanner(banner string) (string, error) {
	file, err := os.Open(banner)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var builder strings.Builder
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
		builder.WriteString("\n")
	}
	return builder.String(), nil
}
