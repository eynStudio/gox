package strx

import (
	"bytes"
	"unicode"
)

func UnderScoreCase(n string) string {
	var buf bytes.Buffer
	for i, it := range n {
		if i == 0 {
			buf.WriteRune(unicode.ToLower(it))
		} else if unicode.IsUpper(it) {
			buf.WriteRune('_')
			buf.WriteRune(unicode.ToLower(it))
		} else {
			buf.WriteRune(it)
		}
	}
	return buf.String()
}

func LowerCamel(s string) string { return toCamel(unicode.ToLower, s) }
func UpperCamel(s string) string { return toCamel(unicode.ToUpper, s) }

func toCamel(first func(r rune) rune, s string) string {
	var buf bytes.Buffer
	lastIs_ := false
	for i, it := range s {
		switch {
		case i == 0:
			buf.WriteRune(first(it))
		case it == '_':
			lastIs_ = true
		case lastIs_:
			buf.WriteRune(unicode.ToUpper(it))
			lastIs_ = false
		default:
			buf.WriteRune(it)
		}
	}
	return buf.String()
}

func Wrap(str, wrap string) string { return wrap + str + wrap }
