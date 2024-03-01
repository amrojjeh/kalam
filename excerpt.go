package kalam

import (
	"fmt"
	"strings"
)

type LetterPack struct {
	Letter          rune
	Vowel           rune
	Shadda          bool
	SuperscriptAlef bool
}

// UnpointedString returns the word without any vowels. The shadda is shown
// if showShadda is true.
func Unpointed(pointedWord string, showShadda bool) string {
	res := ""
	for _, l := range pointedWord {
		c := string(l)
		if !IsShortVowel(l) {
			if (showShadda && l == Shadda) || l != Shadda {
				res += c
			}
		}
	}
	return res
}

// Base returns a new word which does not have the last letter of w
func Base(word string) string {
	res := ""
	letters := LetterPacks(word)
	for _, l := range letters[0 : len(letters)-1] {
		res += l.String()
	}
	return res
}

// Termination returns the last letter of w
func Termination(word string) LetterPack {
	letters := LetterPacks(word)
	return letters[len(letters)-1]
}

// IsValid checks if every Arabic letter in pointedWord has a vowel, and that each letter
// only has one optional vowel and one optional shadda
// IsValid makes a call to IsContentClean
func IsValid(pointedWord string) bool {
	l := false
	v := false
	s := false
	superscript := false
	for _, c := range pointedWord {
		if c == Shadda {
			if l == false || s == true {
				return false
			}
			s = true
		} else if c == SuperscriptAlef {
			if l == false || superscript {
				return false
			}
			superscript = true
		} else if vowels[c] {
			if l == false || v == true {
				return false
			}
			v = true
		} else {
			l = true
			v = false
			s = false
		}
	}
	return IsContentClean(pointedWord)
}

// LetterPacks breaks down each letter from pointedWord into a LetterPack struct
// LetterPacks assumes pointedWord is valid
func LetterPacks(pointedWord string) []LetterPack {
	letters := []LetterPack{}
	letter := LetterPack{
		Vowel: 0,
	}
	for _, l := range pointedWord {
		if l == Shadda {
			letter.Shadda = true
		} else if l == SuperscriptAlef {
			letter.SuperscriptAlef = true
		} else if vowels[l] {
			letter.Vowel = l
		} else {
			if letter.Letter != 0 {
				letters = append(letters, letter)
				letter = LetterPack{
					Vowel: Sukoon,
				}
			}
			letter.Letter = l
		}
	}
	letters = append(letters, letter)
	return letters
}

func LetterPacksToString(ls []LetterPack) string {
	var b strings.Builder
	for _, l := range ls {
		b.WriteString(l.String())
	}
	return b.String()
}

func (l LetterPack) String() string {
	shadda := ""
	superscript := ""
	vowel := ""
	if l.Vowel != 0 {
		vowel = string(l.Vowel)
	}
	if l.Shadda {
		shadda = string(Shadda)
	}
	if l.SuperscriptAlef {
		superscript = string(SuperscriptAlef)
	}
	return fmt.Sprintf("%c%s%s%s", l.Letter, vowel, shadda, superscript)
}

func (l LetterPack) Unpointed(showShadda bool) string {
	if !l.Shadda || !showShadda {
		return string(l.Letter)
	}

	return string(l.Letter) + string(Shadda)
}

func LetterPackFromString(str string) LetterPack {
	letter := LetterPack{}
	for _, l := range str {
		switch l {
		case Shadda:
			letter.Shadda = true
		case Sukoon, Damma, Fatha, Kasra, Dammatan, Fathatan, Kasratan:
			letter.Vowel = l
		default:
			letter.Letter = l
		}
	}
	return letter
}

func (l LetterPack) EqualTo(o LetterPack) bool {
	return l.Shadda == o.Shadda && l.Letter == o.Letter && l.Vowel == o.Vowel
}
