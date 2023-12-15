package kalam

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

const (
	Shadda = string(rune(0x0651))

	Sukoon   = string(rune(0x0652))
	Damma    = string(rune(0x064F))
	Fatha    = string(rune(0x064E))
	Kasra    = string(rune(0x0650))
	Dammatan = string(rune(0x064C))
	Fathatan = string(rune(0x064B))
	Kasratan = string(rune(0x064D))

	Placeholder = string(rune(0x25CC))

	SuperscriptAlef = string(rune(0x670))
)

var GrammaticalTags = []string{
	"اسم مرفوع",
	"اسم منصوب",
	"اسم مجرور",

	" فعل مرفوع",
	"فعل منصوب",
	"فعل مجزوم",

	"مبني",
	"توابع",

	"مضارع مرفوع",
	"مضارع منصوب بحرف النصب",
	"مضارع مجزوم بحرف الجزم",
	"مضارع مجزوم بأداة الشرط الجازم",

	"مبتدأ",
	"خبر ",
	"فاعل نائب",
	"فاعل",
	"اسم كان وأخواتها",
	"خبر إن وأخواتها",

	"مفعول به",
	"مفعول به ثان",
	"مفعول به ثالث",
	"مفعول فيه",
	"مفعول مطلق",
	"مفعول لأجله ",
	"مفعول معه ",
	"حال ",
	"تمييز ",
	"مستثنى ",
	"حصر ",
	"منادى",
	" اسم وخبر ظن وأخواتها ",
	"إسم وخبر حرف نفي",
	"اسم إن وأخواتها",
	" خبر كان واخواتها",

	"إسم مجرور",
	"مضاف إليه",

	"نعت",
	"إسم معطوف",
	"توكيد",
	"بدل",
}

// This is preferred over unicode.IsSpace since we have our own whitespace rules
func IsWhitespace(letter rune) bool {
	return letter == ' '
}

var Punctuation = regexp.MustCompile("[\\.:«»،\"—]")

func IsPunctuation(letter rune) bool {
	return Punctuation.MatchString(string(letter))
}

// IsArabicLetter checks if a letter is part of the classical Arabic script.
// It returns false for tashkeel
func IsArabicLetter(letter rune) bool {
	if letter >= 0x0621 && letter <= 0x063A {
		return true
	}
	if letter >= 0x0641 && letter <= 0x064A {
		return true
	}
	return false
}

// IsVowel checks if the character is a fatha, kasra, damma, or sukoon, with
// their tanween variations. It returns false for shadda and long vowels like
// the alef.
func IsVowel(letter rune) bool {
	sl := string(letter)
	return sl == Sukoon || sl == Damma || sl == Fatha || sl == Kasra ||
		sl == Dammatan || sl == Fathatan || sl == Kasratan
}

// IsShadda checks if the character is a shadda.
func IsShadda(letter rune) bool {
	return string(letter) == Shadda
}

// TODO(Amr Ojjeh): Rewrite clean content into two functions:
// one that cleans content and another that checks if the content is clean.
// It currently does both

// CleanContent cleans a string, ensuring that there are no double spaces
// and that all characters are either Arabic letter, an accepted whitespace,
// or an accepted punctuation. All tashkeel would be deleted.
func CleanContent(content string) (string, error) {
	for _, c := range content {
		if !(IsArabicLetter(c) || IsWhitespace(c) || IsPunctuation(c)) {
			return "", errors.New(fmt.Sprintf("kalam: %v is an invalid letter", c))
		}
	}

	// Remove double spaces
	r, _ := regexp.Compile(" +")
	content = r.ReplaceAllString(content, " ")

	// Trim sentence
	content = strings.TrimFunc(content, unicode.IsSpace)
	return content, nil
}
