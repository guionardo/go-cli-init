package pkg

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
var camelCaseRegex = regexp.MustCompile(`(?m)([A-Z]{1}[a-z0-9]{0,60})`)


func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())

	// Create a slice of runes
	charset := make([]rune, length)

	for i := range charset {
		charset[i] = rune(rand.Intn(26) + 97)
	}

	// Display the character
	return string(charset)
}

func join(words []string) (result string) {
	w2 := make([]string, 0, len(words))
	for _, word := range words {
		word = strings.ToLower(strings.Trim(word, " \t\r\n"))
		if len(word) == 0 {
			continue
		}
		w2 = append(w2, word)
	}
	return strings.Join(w2, "-")
}

func CreateSlug(text string) (result string) {
	text = nonAlphanumericRegex.ReplaceAllString(text, "")

	// tries to split by spaces
	words := strings.Split(text, " ")
	if len(words) > 1 {
		return join(words)
	}

	// tries to split by camel case
	words = camelCaseRegex.FindAllString(text, -1)
	if len(words) > 1 {
		return join(words)
	}

	return text
}
