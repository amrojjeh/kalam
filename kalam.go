package kalam

import (
	"regexp"
	"strings"
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
	"اسم وخبر حرف نفي",
	"اسم إن وأخواتها",
	" خبر كان واخواتها",

	"مضاف إليه",
	"اسم بعد حرف جر",

	"نعت",
	"اسم معطوف",
	"توكيد",
	"بدل",
}

// IsWhitespace is preferred over unicode.IsSpace since we have our own whitespace rules
func IsWhitespace(letter rune) bool {
	return letter == ' '
}

func PunctuationRegex() (*regexp.Regexp, error) {
	str := ""
	for key := range punctuation {
		str += string(key)
	}
	e, err := regexp.Compile("[" + str + "]")
	if err != nil {
		return nil, err
	}
	return e, nil
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
