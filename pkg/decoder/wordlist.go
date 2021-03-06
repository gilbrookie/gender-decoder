package decoder

// SOURCE: https://github.com/lovedaybrooke/gender-decoder/blob/master/app/wordlists.py

// These words are written as the stem to make it easier to match all variants.
// In other words, the suffix is intentionally left out.

var feminineCodedWords = []string{
	"agree",
	"affectionate",
	"child",
	"cheer",
	"collab",
	"commit",
	"communal",
	"compassion",
	"connect",
	"considerate",
	"cooperat",
	"co-operat",
	"depend",
	"emotiona",
	"empath",
	"feel",
	"flatterable",
	"gentle",
	"honest",
	"interpersonal",
	"interdependen",
	"interpersona",
	"inter-personal",
	"inter-dependen",
	"inter-persona",
	"kind",
	"kinship",
	"loyal",
	"modesty",
	"nag",
	"nurtur",
	"pleasant",
	"polite",
	"quiet",
	"respon",
	"sensitiv",
	"submissive",
	"support",
	"sympath",
	"tender",
	"together",
	"trust",
	"understand",
	"warm",
	"whin",
	"enthusias",
	"inclusive",
	"yield",
	"share",
	"sharin",
}

var masculineCodedWords = []string{
	"active",
	"adventurous",
	"aggress",
	"ambitio",
	"analy",
	"assert",
	"athlet",
	"autonom",
	"battle",
	"boast",
	"challeng",
	"champion",
	"compet",
	"confident",
	"courag",
	"decid",
	"decision",
	"decisive",
	"defend",
	"determin",
	"domina",
	"dominant",
	"driven",
	"fearless",
	"fight",
	"force",
	"greedy",
	"head-strong",
	"headstrong",
	"hierarch",
	"hostil",
	"impulsive",
	"independen",
	"individual",
	"intellect",
	"lead",
	"logic",
	"objective",
	"opinion",
	"outspoken",
	"persist",
	"principle",
	"reckless",
	"self-confiden",
	"self-relian",
	"self-sufficien",
	"selfconfiden",
	"selfrelian",
	"selfsufficien",
	"stubborn",
	"superior",
	"unreasonab",
}

var hyphenatedCodedWords = []string{
	"co-operat",
	"inter-personal",
	"inter-dependen",
	"inter-persona",
	"self-confiden",
	"self-relian",
	"self-sufficien",
}

type Wordlist struct {
	feminineCoded, masculineCoded, hyphenatedCoded []string
}

func NewWordlist() *Wordlist {
	return &Wordlist{
		feminineCoded:   feminineCodedWords,
		masculineCoded:  masculineCodedWords,
		hyphenatedCoded: hyphenatedCodedWords,
	}
}
