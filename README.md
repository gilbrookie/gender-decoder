# Gender-Decoder

A command line tool that inspects job description for gender-coded language.

Inspired by a research paper written by Danielle Gaucher, Justin Friesen, and Aaron C. Kay from 2011, called [Evidence That Gendered Wording in Job Advertisements Exists and Sustains Gender Inequality (Journal of Personality and Social Psychology](http://gender-decoder.katmatfield.com/static/documents/Gaucher-Friesen-Kay-JPSP-Gendered-Wording-in-Job-ads.pdf), July 2011, Vol 101(1), p109-28).

Written in Go.  Currently, supports .docx and .txt files.  Can read single file or a directory of files

Shout out to http://gender-decoder.katmatfield.com/ 

## Usage

```bash
go install github.com/gilbrookie/gender-decoder
````
or 
```bash
git clone https://github.com/gilbrookie/gender-decoder
```

Run the tool
```
# Build
go build -o gender-decoder cmd/gender-decoder/main.go
# execute
./gender-decoder -i <file or folder to scan>
```