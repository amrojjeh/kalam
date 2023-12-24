package kalam

import "fmt"

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

func (w Word) String() string {
	return w.PointedWord
}

// UnpointedString returns the word without any vowels. The shadda is shown
// if showShadda is true.
func (w Word) UnpointedString(showShadda bool) string {
	res := ""
	for _, l := range w.PointedWord {
		c := string(l)
		if c != Sukoon && c != Damma && c != Fatha && c != Kasra &&
			c != Dammatan && c != Fathatan && c != Kasratan {
			res += c
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
		switch string(c) {
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
		switch string(l) {
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

// TODO(Amr Ojjeh): Write Sentence.String and Excerpt.String

func (l LetterPack) String() string {
	if l.Shadda {
		return fmt.Sprintf("%c%c%s", l.Letter, l.Vowel, Shadda)
	}
	return fmt.Sprintf("%c%c", l.Letter, l.Vowel)
}
