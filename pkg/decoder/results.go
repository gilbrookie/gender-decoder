package decoder

import (
	"log"
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
	masculineCodedWords  map[string]bool
	feminineCodedWords   map[string]bool
	hyphenatedCodedWords map[string]bool
}

func NewResults(filepath string) *Results {
	return &Results{
		filepath: filepath,
		// initialize the maps
		// The intent is to use the map keys like a Set to avoid duplicate entries
		masculineCodedWords:  make(map[string]bool),
		feminineCodedWords:   make(map[string]bool),
		hyphenatedCodedWords: make(map[string]bool),
	}
}
func (r *Results) foundMasculineWord(w string) {
	r.masculineCodedWords[w] = true
}
func (r *Results) foundFeminineWord(w string) {
	r.feminineCodedWords[w] = true
}
func (r *Results) foundHyphenatedWord(w string) {
	r.hyphenatedCodedWords[w] = true
}

func (r *Results) Explain() {
	var result, explanation string

	if len(r.masculineCodedWords) > len(r.feminineCodedWords) {
		// mostly masculine
		result = "masculine"
		explanation = masculineLeaning
	} else if len(r.feminineCodedWords) > len(masculineCodedWords) {
		// mostly feminine
		result = "feminine"
		explanation = feminineLeaning
	} else if len(r.masculineCodedWords) == 0 && len(r.feminineCodedWords) == 0 {
		// clean !!!
		result = "clean"
		explanation = cleanMessage
	} else {
		// neutral
		result = "neutral"
		explanation = neutralMessage
	}

	log.Println("------------------------------------")
	log.Println("--> File:", r.filepath)
	log.Println("--> Result:", result)
	log.Println("--> Explanation:", explanation)
	log.Println("--> masculine words: ", getKeys(r.masculineCodedWords))
	log.Println("--> feminine words", getKeys(r.feminineCodedWords))
	log.Println("--> hyphenated words", getKeys(r.hyphenatedCodedWords))

}
