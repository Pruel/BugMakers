package ascii

import (
	"fmt"
	"strings"
)

// будем использовать мапу, для хранения позиций (это ускорит поиск позиции)
var asciiPositions map[rune]int

// особенная функция в GO, в нашем случае она позволит заполнить asciiPositions данными еще до начала выполнения основной команды
// (не имеет явного вызова)
func init() {

	asciiPositions = make(map[rune]int)

	for i, char := range asciiStorage() {
		asciiPositions[char] = 1 + (9 * i)
	}
}

func asciiStorage() (storage string) {

	for i := ' '; i <= '~'; i++ {
		storage += string(i)
	}
	return storage
}

// проверяем, существует ли позиция
func checkPosition(char rune, Ascii string) int {
	for i, character := range Ascii {
		if character == char {
			return 1 + (9 * i)
		}
	}
	return -1
}

// преобразуем входной текст, используя заданный баннер
func PrintAscii(banner, inputText string) (string, error) {
	ourFileTransform := strings.Split(banner, "\n")
	asciiStorage :=  asciiStorage()

	lines := strings.Split(inputText, "\\n")
	var resultBuilder strings.Builder

	for _, line := range lines {
		for row := 0; row < 9; row++ {
			for _, char := range line {
				charPosition := checkPosition(char, asciiStorage)
				if charPosition == -1 {
					return "", fmt.Errorf("Invalid character : %c", char)
				}
				resultBuilder.WriteString(ourFileTransform[charPosition+row])
			}
			if row != 8 {
				resultBuilder.WriteString("\n")
			}
		}
	}
	return resultBuilder.String(), nil
}

