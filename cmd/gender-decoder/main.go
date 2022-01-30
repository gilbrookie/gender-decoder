package main

import (
	"flag"
	decoder2 "github.com/gilbrookie/gender-decoder/pkg/decoder"
	"log"
)

func main() {
	var input = flag.String("i", "", "Input file/folder location")
	flag.Parse()

	log.Println("Got Input", *input)

	cfg := decoder2.NewDecoderConfig(input)
	decoder, err := decoder2.NewDecoder(cfg)
	if err != nil {
		log.Fatalln("Failed to start decoder, err:", err)
	}

	// Run the assessment
	decoder.Assess()

	// print the results
	decoder.ShowResults()
}
