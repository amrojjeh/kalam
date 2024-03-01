package kalam

import (
	"fmt"
	"testing"

	"github.com/amrojjeh/kalam/assert"
)

func TestBase(t *testing.T) {
	w := FromBuckwalter("bayotN")
	assert.Equal(t, Base(w), FromBuckwalter("bayo"))
}

func TestTermination(t *testing.T) {
	w := FromBuckwalter("bayotN")
	wt := Termination(w)

	assert.Equal(t, wt.Letter, Teh)
	assert.Equal(t, wt.Shadda, false)
	assert.Equal(t, wt.Vowel, Dammatan)
}

func TestLetterPackString(t *testing.T) {
	l := LetterPack{
		Letter: Yeh,
		Vowel:  Fatha,
		Shadda: true,
	}

	assert.Equal(t, l.String(), fmt.Sprintf("%c%c%c", Yeh, Fatha, Shadda))
}
