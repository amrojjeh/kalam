package kalam

import (
	"regexp"
)

var Cases = []string{
	FromBuckwalter("Asm mrfwE"),
	FromBuckwalter("Asm mnSwb"),
	FromBuckwalter("Asm mjrwr"),
	FromBuckwalter("fEl mrfwE"),
	FromBuckwalter("fEl mnSwb"),
	FromBuckwalter("fEl mjzwm"),
	FromBuckwalter("mbny"),
	FromBuckwalter("tAbE"),
}

var States = map[string][]string{
	Cases[0]: {
		FromBuckwalter("mbtd>"),
		FromBuckwalter("xbr"),
		FromBuckwalter("fAEl nA}b"),
		FromBuckwalter("fAEl"),
		FromBuckwalter("Asm kAn w>xwAthA"),
		FromBuckwalter("Asm <n w>xwAthA"),
	},
	Cases[1]: {
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
	Cases[2]: {
		FromBuckwalter("Asm bEd Hrf jr"),
		FromBuckwalter("mDAf <lyh"),
	},

	Cases[3]: {
		FromBuckwalter("mDArE mrfwE"),
	},
	Cases[4]: {
		FromBuckwalter("mDArE mnSwb bHrf AlnSb"),
	},

	Cases[5]: {
		FromBuckwalter("mDArE mjzwm bHrf Aljzm"),
		FromBuckwalter("mDArE mjzwm b>dAp Al$rT AljAzm"),
	},

	// TODO(Amr Ojjeh): Add functions
	Cases[6]: {},

	Cases[7]: {
		FromBuckwalter("nEt"),
		FromBuckwalter("Asm mETwf"),
		FromBuckwalter("twkyd"),
		FromBuckwalter("bdl"),
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
		b, ok := toBuckwalter[l]
		if ok {
			res += string(b)
		} else {
			res += string(l)
		}
	}
	return res
}

func ToBuckwalter(sen string) string {
	res := ""
	for _, l := range sen {
		b, ok := fromBuckwalter[l]
		if ok {
			res += string(b)
		} else {
			res += string(l)
		}
	}
	return res
}
