package decoder

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), err
}

func isDocx(pathToFile string) bool {
	return filepath.Ext(pathToFile) == ".docx" || filepath.Ext(pathToFile) == ".DOCX"
}

func isTxt(pathToFile string) bool {
	return filepath.Ext(pathToFile) == ".txt"
}

// normalize fixes quotation marks in document
func normalizeQuotes(in rune) rune {
	switch in {
	case '“', '”':
		return '"'
	case '‘', '’':
		return '\''
	}
	return in
}

// cleans raw xml data from tags, brackets, punctuation, excess whitespace
func normalizeAll(text string) string {
	// Remove all the xml tags
	brackets := regexp.MustCompile("<.*?>")
	text = brackets.ReplaceAllString(text, "")
	// replace quotes with and escaped quote
	quotes := regexp.MustCompile("&quot;")
	text = quotes.ReplaceAllString(text, "\"")
	// remove punctuation, and replace with an empty space
	punctuation := regexp.MustCompile(`[?.,;:)(-]`)
	text = punctuation.ReplaceAllString(text, " ")
	// Remove any extra whitespace
	whitespace := regexp.MustCompile(`\s+`)
	text = whitespace.ReplaceAllString(text, " ")

	return strings.Map(normalizeQuotes, text)
}

func getKeys(m map[string][]string) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
