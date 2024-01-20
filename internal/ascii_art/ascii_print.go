package utils

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
func checkAsciiPos(symb rune) int {

	if pos, exists := asciiPositions[symb]; exists {
		return pos
	}
	return -1
}

// преобразуем входной текст, используя заданный баннер
func printAscii(banner, inputTxt string) (string, error) {
	// разделяем баннер на строки
	splitBanner := strings.Split(banner, "\n")
	// разделяем входной текст на строки
	lines := strings.Split(inputTxt, "\\n")

	var resultBuilder strings.Builder

	for _, line := range lines { // итерация по строкам входного текста
		for row := 0; row < 9; row++ { // проходим по каждой их 9 строк баннера
			for _, char := range line { // проходим по символам в текущей строке входного текста
				charPosition := checkAsciiPos(char)
				if charPosition == -1 {
					return "", fmt.Errorf("invalid character : %c", char)
				}
				resultBuilder.WriteString(splitBanner[charPosition+row])
			}
			if row != 8 {
				resultBuilder.WriteString("\n")
			}
		}
	}
	return resultBuilder.String(), nil
}
