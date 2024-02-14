package kalam

import (
	"fmt"
	"testing"

	"github.com/amrojjeh/kalam/assert"
)

func TestPunctuationRegex(t *testing.T) {
	reg := PunctuationRegex()
	assert.Equal(t, reg.MatchString(string(ArabicComma)), true)
	assert.Equal(t, reg.MatchString(string("!")), false)
}

func TestRemoveExtraWhitespace(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected string
	}{
		{
			name:     "Normal",
			content:  "This is a test",
			expected: "This is a test",
		},
		{
			name:     "Whitespace",
			content:  " \t  ",
			expected: "",
		},
		{
			name:     "Mixed",
			content:  "  wow   how is  this? ",
			expected: "wow how is this?",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, RemoveExtraWhitespace(test.content), test.expected)
		})
	}
}

func TestIsContentClean(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected bool
	}{
		{
			name:     "Normal",
			content:  "هذا بيت",
			expected: true,
		},
		{
			name:     "English",
			content:  "Test",
			expected: false,
		},
		{
			name:     "Mixed",
			content:  "هذا test",
			expected: false,
		},
		{
			name:     "Empty",
			content:  "",
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, IsContentClean(test.content), test.expected)
		})
	}
}

func TestBuckWalter(t *testing.T) {
	assert.Equal(t, FromBuckwalter("h*A byt."), fmt.Sprintf("%c%c%c %c%c%c.",
		Heh, Thal, Alef, Beh, Yeh, Teh))
}
