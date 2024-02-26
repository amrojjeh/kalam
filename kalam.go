package kalam

import (
	"regexp"
	"strings"
)

type declention struct {
	Declention string
	Functions  []string
}

var Declention = []declention{
	{
		Declention: FromBuckwalter("Asm mrfwE"),
		Functions: []string{
			FromBuckwalter("mbtd>"),
			FromBuckwalter("xbr"),
			FromBuckwalter("fAEl nA}b"),
			FromBuckwalter("fAEl"),
			FromBuckwalter("Asm kAn w>xwAthA"),
			FromBuckwalter("Asm <n w>xwAthA"),
		},
	},
	{
		Declention: FromBuckwalter("Asm mnSwb"),
		Functions: []string{
			FromBuckwalter("mfEwl bh"),
			FromBuckwalter("mfEwl bh vAn"),
			FromBuckwalter("mfEwl bh vAlv"),
			FromBuckwalter("mfEwl fyh"),
			FromBuckwalter("mfEwl mTlq"),
			FromBuckwalter("mfEwl l>jlh"),
			FromBuckwalter("mfEwl mEh"),
			FromBuckwalter("HAl"),
			FromBuckwalter("tmyyz"),
			FromBuckwalter("mstvnY"),
			FromBuckwalter("HSr"),
			FromBuckwalter("mnAdY"),
			FromBuckwalter("Asm wxbr Zn w>xwAthA"),
			FromBuckwalter("Asm wxbr Hrf nfy"),
			FromBuckwalter("Asm <n w>xwAthA"),
			FromBuckwalter("xbr kAn w>xwAthA"),
		},
	},
	{
		Declention: FromBuckwalter("Asm mjrwr"),
		Functions: []string{
			FromBuckwalter("Asm bEd Hrf jr"),
			FromBuckwalter("mDAf <lyh"),
		},
	},

	{
		Declention: FromBuckwalter("fEl mrfwE"),
		Functions: []string{
			FromBuckwalter("mDArE mrfwE"),
		},
	},
	{
		Declention: FromBuckwalter("fEl mnSwb"),
		Functions: []string{
			FromBuckwalter("mDArE mnSwb bHrf AlnSb"),
		},
	},
	{
		Declention: FromBuckwalter("fEl mjzwm"),
		Functions: []string{
			FromBuckwalter("mDArE mjzwm bHrf Aljzm"),
			FromBuckwalter("mDArE mjzwm b>dAp Al$rT AljAzm"),
		},
	},

	{
		// TODO(Amr Ojjeh): Add functions
		Declention: FromBuckwalter("mbny"),
		Functions:  []string{},
	},
	{
		Declention: FromBuckwalter("tAbE"),
		Functions: []string{
			FromBuckwalter("nEt"),
			FromBuckwalter("Asm mETwf"),
			FromBuckwalter("twkyd"),
			FromBuckwalter("bdl"),
		},
	},
}

// IsWhitespace is preferred over unicode.IsSpace since we have our own whitespace rules
func IsWhitespace(letter rune) bool {
	return letter == ' '
}

func PunctuationRegex() *regexp.Regexp {
	str := ""
	for key := range punctuation {
		str += string(key)
	}
	e := regexp.MustCompile("[" + str + "]")
	return e
}

// IsPunctuation cheks if a latter is part of the accepted punctuation
func IsPunctuation(letter rune) bool {
	return punctuation[letter]
}

// IsArabicLetter checks if a letter is part of the classical Arabic script.
// It returns false for tashkeel
func IsArabicLetter(letter rune) bool {
	return letters[letter]
}

// IsVowel checks if the character is a fatha, kasra, damma, or sukoon, with
// their tanween variations. It returns false for shadda and long vowels like
// the alef.
func IsShortVowel(letter rune) bool {
	return vowels[letter]
}

func IsTanween(letter rune) bool {
	return letter == Fathatan ||
		letter == Dammatan ||
		letter == Kasratan
}

// IsShadda checks if the character is a shadda.
func IsShadda(letter rune) bool {
	return letter == Shadda
}

// RemoveExtraWhitespace removes unnecessary whitespace, ensuring that there
// are no double spaces and no beginning/ending whitespace.
func RemoveExtraWhitespace(content string) string {
	// Remove double spaces
	r, _ := regexp.Compile(" +")
	content = r.ReplaceAllString(content, " ")

	content = strings.TrimSpace(content)
	return content
}

// IsContentClean ensures that all characters conform to Kalam's character
// set
func IsContentClean(content string) bool {
	for _, c := range content {
		if !(IsArabicLetter(c) || IsWhitespace(c) || IsPunctuation(c)) {
			return false
		}
	}
	return true
}

func FromBuckwalter(sen string) string {
	res := ""
	for _, l := range sen {
		b, ok := buckwalter[l]
		if ok {
			res += string(b)
		} else {
			res += string(l)
		}
	}
	return res
}
