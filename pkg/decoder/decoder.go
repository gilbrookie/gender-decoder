package decoder

import (
	"fmt"
	"github.com/nguyenthenguyen/docx"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Decoder interface {
	Assess(file string, list *Wordlist) (*Results, error)
}

type Config struct {
	input    *string
	wordlist *Wordlist
}

func NewDecoderConfig(input *string) *Config {
	return &Config{input, nil}
}

type DecodeWrapper struct {
	cfg           *Config
	filesToAssess map[string]Decoder
	results       []*Results
}

func (dw *DecodeWrapper) Assess() {

	list := NewWordlist()
	for file, decoder := range dw.filesToAssess {
		log.Println("Assessing file", file)
		res, err := decoder.Assess(file, list)
		if err != nil {
			log.Printf("Caught error assessing file %s, %v", file, err)
			continue
		}
		dw.results = append(dw.results, res)
	}
}

func (dw *DecodeWrapper) ShowResults() {
	for _, r := range dw.results {
		r.Explain()
	}
}

func NewDecoder(cfg *Config) (*DecodeWrapper, error) {
	var decodeWrapper DecodeWrapper
	decodeWrapper.filesToAssess = make(map[string]Decoder)

	log.Println("searching for input files...")
	// If the input is a folder, iterate through the files in the directory and assess each one
	isDir, err := isDirectory(*cfg.input)
	if err != nil {
		return nil, fmt.Errorf("failed input file stat %v", err)
	}
	if isDir {
		// walk the directory assessing the files contained
		err := filepath.Walk(*cfg.input, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Fatalf(err.Error())
			}
			if !info.IsDir() {
				log.Printf("Found: file name: %s\n", info.Name())
				decodeWrapper.addFileToAssessment(path)
			}
			return nil
		})

		if err != nil {
			return nil, fmt.Errorf("failed reading through directory %v", err)
		}
	} else {
		// If the input is a file, open the file and process it
		decodeWrapper.addFileToAssessment(*cfg.input)
	}

	return &decodeWrapper, nil
}

func (dw *DecodeWrapper) addFileToAssessment(file string) {
	var d Decoder
	if isDocx(file) {
		d = &docxDecoder{}
	} else if isTxt(file) {
		d = &textDecoder{}
	} else {
		log.Println("ERROR Usupported file (not .txt/.docx), skipping")
		return
	}
	dw.filesToAssess[file] = d
}

type textDecoder struct{}

func (td *textDecoder) Assess(file string, wordlist *Wordlist) (*Results, error) {
	log.Println(file)
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	cleanedContent := strings.Split(normalizeAll(string(content)), " ")
	res := compareToWordList(file, cleanedContent, wordlist)

	return res, nil
}

type docxDecoder struct{}

func (dd *docxDecoder) Assess(file string, wordlist *Wordlist) (*Results, error) {
	// Read from docx file
	r, err := docx.ReadDocxFile(file)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	content := r.Editable().GetContent()

	// clean up content so we have just the text to process, then split the words into a list
	cleanedContent := strings.Split(normalizeAll(content), " ")

	res := compareToWordList(file, cleanedContent, wordlist)

	return res, nil
}

func compareToWordList(file string, content []string, wordlist *Wordlist) *Results {
	res := NewResults(file)
	for _, word := range content {
		for _, listedWord := range wordlist.masculineCoded {
			if strings.HasPrefix(word, listedWord) {
				res.foundMasculineWord(word)
			}
		}
		for _, listedWord := range wordlist.feminineCoded {
			if strings.HasPrefix(word, listedWord) {
				res.foundFeminineWord(word)
			}
		}
		for _, listedWord := range wordlist.hyphenatedCoded {
			if strings.HasPrefix(word, listedWord) {
				res.foundHyphenatedWord(word)
			}
		}
	}
	return res
}
