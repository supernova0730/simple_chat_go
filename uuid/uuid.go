package uuid

import (
	"math/rand"
	"strings"
)

func getRandomCharacter(a, b rune) rune {
	return rune(rand.Intn(int(b-a+1)) + int(a))
}

func getRandomUpperCaseLetter() rune {
	return getRandomCharacter('A', 'Z')
}

func getRandomLowerCaseLetter() rune {
	return getRandomCharacter('a', 'z')
}

func getRandomDigit() rune {
	return getRandomCharacter('0', '9')
}

func GenerateUUID(n int) string {
	var result strings.Builder

	for i := 0; i < n; i++ {
		r := rand.Intn(3)
		switch r {
		case 0:
			result.WriteRune(getRandomUpperCaseLetter())
		case 1:
			result.WriteRune(getRandomLowerCaseLetter())
		case 2:
			result.WriteRune(getRandomDigit())
		}
	}

	return result.String()
}
