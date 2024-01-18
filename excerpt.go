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

type ExcerptIterator struct {
	Excerpt   Excerpt
	SentenceI int
	WordI     int
	Index     int
}

func (e Excerpt) String() string {
	res := ""
	for _, s := range e.Sentences {
		res += s.String() + " "
	}
	return RemoveExtraWhitespace(res)
}

func (e Excerpt) Unpointed(showShadda bool) string {
	res := ""
	for _, s := range e.Sentences {
		res += s.Unpointed(showShadda) + " "
	}
	return RemoveExtraWhitespace(res)
}

// Iterator returns an ExcerptIterator which points to the first quizzable word
func (e Excerpt) Iterator() (ExcerptIterator, bool) {
	i := ExcerptIterator{Excerpt: e}
	if i.Word().Quizzable() {
		return i, true
	} else {
		i, f := i.Next()
		i.Index = 0
		return i, f
	}
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

// Base returns a new word which does not have the last letter of w
func (w Word) Base() Word {
	res := ""
	letters := w.LetterPacks()
	for _, l := range letters[0 : len(letters)-1] {
		res += l.String()
	}
	w.PointedWord = res
	return w
}

// Termination returns the last letter of w
func (w Word) Termination() LetterPack {
	letters := w.LetterPacks()
	return letters[len(letters)-1]
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
	letters = append(letters, letter)
	return letters
}

func (w Word) Quizzable() bool {
	return !w.Punctuation && !w.Ignore
}

func (l LetterPack) String() string {
	if l.Shadda {
		return fmt.Sprintf("%c%c%c", l.Letter, l.Vowel, Shadda)
	}
	return fmt.Sprintf("%c%c", l.Letter, l.Vowel)
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

// Next returns the next quizzable word. If there are no more words, it returns true
func (i ExcerptIterator) Next() (ExcerptIterator, bool) {
	i, found := i.nextWord()
	if found {
		return i, true
	}

	for i.SentenceI < len(i.Excerpt.Sentences)-1 {
		i.SentenceI += 1
		i.WordI = 0
		if i.Word().Quizzable() {
			i.Index += 1
			return i, true
		}
		i, found = i.nextWord()
		if found {
			return i, true
		}
	}
	return i, false
}

func (i ExcerptIterator) nextWord() (ExcerptIterator, bool) {
	for wi, w := range i.Sentence().Words[i.WordI:] {
		if wi == 0 {
			continue
		}
		if w.Quizzable() {
			return ExcerptIterator{
				Excerpt:   i.Excerpt,
				SentenceI: i.SentenceI,
				WordI:     i.WordI + wi,
				Index:     i.Index + 1,
			}, true
		}
	}
	return i, false
}

func (i ExcerptIterator) Sentence() Sentence {
	return i.Excerpt.Sentences[i.SentenceI]
}

func (i ExcerptIterator) Word() Word {
	return i.Sentence().Words[i.WordI]
}
