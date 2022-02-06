package decoder

import (
	"fmt"
)

var feminineLeaning = `
	This job description uses more words that are stereotypically feminine
	than words that are stereotypically masculine. Research	suggests
	this will have only a slight effect on how appealing the job is
	to men, and will encourage women applicants. `

var masculineLeaning = `
	This job description uses more words that are stereotypically masculine
	than words that are stereotypically feminine. It risks putting women off
	applying, but will probably encourage men to apply. `

var neutralMessage = `
	This job description uses an equal number of words that are
	stereotypically masculine and stereotypically feminine. It probably won't
	be off-putting to men or women applicants. `

var cleanMessage = `
	This job description doesn't use any words that are stereotypically
	masculine and stereotypically feminine. It probably won't be off-putting
	to men or women applicants. `

type Results struct {
	filepath             string
	masculineCodedWords  map[string][]string
	feminineCodedWords   map[string][]string
	hyphenatedCodedWords map[string][]string
}

type ResultResponse struct {
	File           string              `json:"file"`
	Result         string              `json:"result"`
	Description    string              `json:"description"`
	MasculineWords map[string][]string `json:"masculine_words"`
	FeminineWords  map[string][]string `json:"feminine_words"`
}

func NewResults(filepath string) *Results {
	return &Results{
		filepath: filepath,
		// initialize the maps
		// The intent is to use the map keys like a Set to avoid duplicate entries
		masculineCodedWords:  make(map[string][]string),
		feminineCodedWords:   make(map[string][]string),
		hyphenatedCodedWords: make(map[string][]string),
	}
}

func (r *Results) foundMasculineWord(w, listedWord string) {
	r.masculineCodedWords[listedWord] = append(r.masculineCodedWords[listedWord], w)
}
func (r *Results) foundFeminineWord(w, listedWord string) {
	r.feminineCodedWords[listedWord] = append(r.feminineCodedWords[listedWord], w)
}
func (r *Results) foundHyphenatedWord(w, listedWord string) {
	r.hyphenatedCodedWords[listedWord] = append(r.hyphenatedCodedWords[listedWord], w)
}

func (r *Results) getCount(m map[string][]string) int {
	var sum int
	for _, words := range m {
		sum = sum + len(words)
	}
	return sum
}

func PrintMap(m map[string][]string) {
	for k, words := range m {
		fmt.Printf("%s => %s\n", k, words)
	}
}

func (r *Results) Explain() *ResultResponse {
	var result, explanation string

	masculineWords := r.getCount(r.masculineCodedWords)
	feminineWords := r.getCount(r.feminineCodedWords)

	if masculineWords > feminineWords {
		// mostly masculine
		result = "masculine"
		explanation = masculineLeaning
	} else if masculineWords < feminineWords {
		// mostly feminine
		result = "feminine"
		explanation = feminineLeaning
	} else {
		// neutral
		result = "neutral"
		explanation = neutralMessage
	}

	if masculineWords == 0 && feminineWords == 0 {
		// clean !!!
		result = "clean"
		explanation = cleanMessage
	}

	fmt.Println("")
	fmt.Println("File:", r.filepath)
	fmt.Println("Result:", result)
	fmt.Println("Explanation:", explanation)
	fmt.Println("Masculine words:")
	PrintMap(r.masculineCodedWords)
	fmt.Println("Feminine words:")
	PrintMap(r.feminineCodedWords)
	// fmt.Println("masculine words: ", getKeys(r.masculineCodedWords))
	// fmt.Println("feminine words", getKeys(r.feminineCodedWords))

	res := ResultResponse{
		File:           r.filepath,
		Result:         result,
		Description:    explanation,
		MasculineWords: r.masculineCodedWords,
		FeminineWords:  r.feminineCodedWords,
	}
	return &res
}
