package kalam

import (
	"fmt"
	"testing"

	"github.com/amrojjeh/kalam/assert"
)

func FactoryExcerpt() Excerpt {
	return Excerpt{
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
			{
				Words: []Word{
					{
						PointedWord: FromBuckwalter("ha*ihi"),
						Tags:        []string{},
						Punctuation: false,
						Ignore:      true,
						Preceding:   false,
					},
					{
						PointedWord: FromBuckwalter("musolimapN"),
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
}

func TestExcerptString(t *testing.T) {
	e := FactoryExcerpt()
	assert.Equal(t, e.String(), FromBuckwalter("ha*aAo bayotN. ha*ihi musolimapN."))
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

func TestExcerptIterator(t *testing.T) {
	e := FactoryExcerpt()
	i, found := e.Iterator()
	assert.Equal(t, found, true)
	assert.Equal(t, i.Word().PointedWord, FromBuckwalter("bayotN"))
	assert.Equal(t, i.Index, 0)
	assert.Equal(t, i.WordI, 1)
	assert.Equal(t, i.SentenceI, 0)
	i, found = i.Next()
	assert.Equal(t, found, true)
	assert.Equal(t, i.Word().PointedWord, FromBuckwalter("musolimapN"))
	assert.Equal(t, i.Index, 1)
	assert.Equal(t, i.WordI, 1)
	assert.Equal(t, i.SentenceI, 1)
	i, found = i.Next()
	assert.Equal(t, found, false)
	assert.Equal(t, i.Word().PointedWord, FromBuckwalter("musolimapN"))
	assert.Equal(t, i.Index, 1)
	assert.Equal(t, i.WordI, 1)
	assert.Equal(t, i.SentenceI, 1)

	e = Excerpt{
		Name: "Test",
		Sentences: []Sentence{
			{
				Words: []Word{
					{
						PointedWord: "يَنْكِحُ",
						Tags:        []string{},
						Punctuation: false,
						Ignore:      false,
						Preceding:   true,
					},
					{
						PointedWord: "هَاْ",
						Tags:        []string{},
						Punctuation: false,
						Ignore:      true,
						Preceding:   true,
					},
					{
						PointedWord: "،",
						Tags:        []string{},
						Punctuation: true,
						Ignore:      false, // should this be true?
						Preceding:   false,
					},
				},
			},
			{
				Words: []Word{
					{
						PointedWord: "فَهِجْرَتُ",
						Tags:        []string{},
						Punctuation: false,
						Ignore:      false,
						Preceding:   true,
					},
					{
						PointedWord: "هُ",
						Tags:        []string{},
						Punctuation: false,
						Ignore:      false,
						Preceding:   false,
					},
				},
			},
		},
	}

	i, found = e.Iterator()
	assert.Equal(t, found, true)
	assert.Equal(t, i.Word().String(), "يَنْكِحُ")
	assert.Equal(t, i.Index, 0)

	i, found = i.Next()
	assert.Equal(t, i.Word().String(), "فَهِجْرَتُ")
	assert.Equal(t, found, true)
	assert.Equal(t, i.Index, 1)
}
