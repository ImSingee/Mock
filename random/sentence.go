package random

import _ "embed"

import "strings"

//go:embed sentences.txt
var sentences string

var Sentences []string

func init() {
	Sentences = strings.Split(sentences, "\n")
}

func Sentence() string {
	return Sentences[r.Intn(len(Sentences))]
}
