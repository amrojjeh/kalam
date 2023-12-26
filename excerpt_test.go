package kalam

import (
	"fmt"
	"testing"

	"github.com/amrojjeh/kalam/assert"
)

func TestExcerptString(t *testing.T) {
	e := Excerpt{
		Name: "Test",
		Sentences: []Sentence{
			{
				Words: []Word{
					{
						PointedWord: FromBuckwalter("ha*aAo"),
						Tags:        []string{},
						Punctuation: false,
						Ignore:      true,
						Preceding:   false,
					},
					{
						PointedWord: FromBuckwalter("bayotN"),
						Tags:        []string{},
						Punctuation: false,
						Ignore:      false,
						Preceding:   true,
					},
					{
						PointedWord: ".",
						Tags:        []string{},
						Punctuation: true,
						Ignore:      true,
						Preceding:   false,
					},
				},
			},
		},
	}

	assert.Equal(t, e.String(), FromBuckwalter("ha*aAo bayotN."))
}

func TestWordBase(t *testing.T) {
	w := Word{
		PointedWord: FromBuckwalter("bayotN"),
	}

	assert.Equal(t, w.Base().PointedWord, FromBuckwalter("bayo"))
}

func TestWordTermination(t *testing.T) {
	w := Word{
		PointedWord: FromBuckwalter("bayotN"),
	}

	assert.Equal(t, w.Termination().Letter, Teh)
	assert.Equal(t, w.Termination().Shadda, false)
	assert.Equal(t, w.Termination().Vowel, Dammatan)
}

func TestLetterPackString(t *testing.T) {
	l := LetterPack{
		Letter: Yeh,
		Vowel:  Fatha,
		Shadda: true,
	}

	assert.Equal(t, l.String(), fmt.Sprintf("%c%c%c", Yeh, Fatha, Shadda))
}
