package kalam

import (
	"fmt"
	"strings"
)

type Excerpt struct {
	Name      string
	Sentences []Sentence
}

type Sentence struct {
	Words []Word
}

type Word struct {
	PointedWord string
	Tags        []string
	Punctuation bool

	Ignore bool

	// Preceding is true if it preceeds another word or punctuation without
	// any space.
	Preceding bool
}

type LetterPack struct {
	Letter rune
	Vowel  rune
	Shadda bool
}

func (e Excerpt) String() string {
	res := ""
	for _, s := range e.Sentences {
		res += s.String()
	}
	return res
}

func (e Excerpt) Unpointed(showShadda bool) string {
	res := ""
	for _, s := range e.Sentences {
		res += s.Unpointed(showShadda)
	}
	return res
}

func (s Sentence) String() string {
	res := ""
	for _, w := range s.Words {
		res += w.String()
		if !w.Preceding {
			res += " "
		}
	}
	return strings.TrimSpace(res)
}

func (s Sentence) Unpointed(showShadda bool) string {
	res := ""
	for _, w := range s.Words {
		res += w.Unpointed(showShadda)
		if !w.Preceding {
			res += " "
		}
	}
	return res
}

func (w Word) String() string {
	return w.PointedWord
}

// UnpointedString returns the word without any vowels. The shadda is shown
// if showShadda is true.
func (w Word) Unpointed(showShadda bool) string {
	res := ""
	for _, l := range w.PointedWord {
		c := string(l)
		if !IsShortVowel(l) {
			if (showShadda && l == Shadda) || l != Shadda {
				res += c
			}
		}
	}
	return res
}

// IsValid checks if every Arabic letter in w has a vowel, and that each letter
// only has one vowel and only one optional shadda
// IsValid makes a call to IsContentClean
func (w Word) IsValid() bool {
	l := false
	v := false
	s := false
	for _, c := range w.PointedWord {
		switch c {
		case Shadda:
			if l == false || s == true {
				return false
			}
			s = true
		case Sukoon, Damma, Fatha, Kasra, Dammatan, Fathatan, Kasratan:
			if l == false || v == true {
				return false
			}
			v = true
		default:
			if l == true && v == false {
				return false
			}
			l = true
			v = false
			s = false
		}
	}
	return IsContentClean(w.PointedWord)
}

// LetterPacks breaks down each letter from w into a LetterPack struct
// LetterPacks assumes w is valid
func (w Word) LetterPacks() []LetterPack {
	letters := []LetterPack{}
	letter := LetterPack{}
	for _, l := range w.PointedWord {
		switch l {
		case Shadda:
			letter.Shadda = true
		case Sukoon, Damma, Fatha, Kasra, Dammatan, Fathatan, Kasratan:
			letter.Vowel = l
		default:
			if letter.Letter != 0 {
				letters = append(letters, letter)
				letter = LetterPack{}
			}
			letter.Letter = l
		}
	}
	return letters
}

func (l LetterPack) String() string {
	if l.Shadda {
		return fmt.Sprintf("%c%c%c", l.Letter, l.Vowel, Shadda)
	}
	return fmt.Sprintf("%c%c", l.Letter, l.Vowel)
}
