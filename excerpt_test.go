package kalam

import (
	"testing"

	"github.com/amrojjeh/kalam/assert"
)

func TestLetterPackString(t *testing.T) {
	l := LetterPack{
		Letter: 'ي',
		Vowel:  'َ',
		Shadda: true,
	}

	assert.Equal(t, l.String(), "يَّ")
}
