package stringx

import (
	"errors"
	"regexp"
	"strings"
	"unicode"

	"manlu.org/tao/core/lang"
)

var (
	// ErrInvalidStartPosition is an error that indicates the start position is invalid.
	ErrInvalidStartPosition = errors.New("start position is invalid")
	// ErrInvalidStopPosition is an error that indicates the stop position is invalid.
	ErrInvalidStopPosition = errors.New("stop position is invalid")

	// regex for convert string.
	toSnakeReg  = regexp.MustCompile("[A-Z][a-z]")
	toCamelRegs = map[string]*regexp.Regexp{
		" ": regexp.MustCompile(" +[a-zA-Z]"),
		"-": regexp.MustCompile("-+[a-zA-Z]"),
		"_": regexp.MustCompile("_+[a-zA-Z]"),
	}
)

// Contains checks if str is in list.
func Contains(list []string, str string) bool {
	for _, each := range list {
		if each == str {
			return true
		}
	}

	return false
}

// Filter filters chars from s with given filter function.
func Filter(s string, filter func(r rune) bool) string {
	var n int
	chars := []rune(s)
	for i, x := range chars {
		if n < i {
			chars[n] = x
		}
		if !filter(x) {
			n++
		}
	}

	return string(chars[:n])
}

// FirstN returns first n runes from s.
func FirstN(s string, n int, ellipsis ...string) string {
	var i int

	for j := range s {
		if i == n {
			ret := s[:j]
			for _, each := range ellipsis {
				ret += each
			}
			return ret
		}
		i++
	}

	return s
}

// HasEmpty checks if there are empty strings in args.
func HasEmpty(args ...string) bool {
	for _, arg := range args {
		if len(arg) == 0 {
			return true
		}
	}

	return false
}

// NotEmpty checks if all strings are not empty in args.
func NotEmpty(args ...string) bool {
	return !HasEmpty(args...)
}

// Remove removes given strs from strings.
func Remove(strings []string, strs ...string) []string {
	out := append([]string(nil), strings...)

	for _, str := range strs {
		var n int
		for _, v := range out {
			if v != str {
				out[n] = v
				n++
			}
		}
		out = out[:n]
	}

	return out
}

// Reverse reverses s.
func Reverse(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}

// Substr returns runes between start and stop [start, stop) regardless of the chars are ascii or utf8.
func Substr(str string, start, stop int) (string, error) {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return "", ErrInvalidStartPosition
	}

	if stop < 0 || stop > length {
		return "", ErrInvalidStopPosition
	}

	return string(rs[start:stop]), nil
}

// TakeOne returns valid string if not empty or later one.
func TakeOne(valid, or string) string {
	if len(valid) > 0 {
		return valid
	}

	return or
}

// TakeWithPriority returns the first not empty result from fns.
func TakeWithPriority(fns ...func() string) string {
	for _, fn := range fns {
		val := fn()
		if len(val) > 0 {
			return val
		}
	}

	return ""
}

// Union merges the strings in first and second.
func Union(first, second []string) []string {
	set := make(map[string]lang.PlaceholderType)

	for _, each := range first {
		set[each] = lang.Placeholder
	}
	for _, each := range second {
		set[each] = lang.Placeholder
	}

	merged := make([]string, 0, len(set))
	for k := range set {
		merged = append(merged, k)
	}

	return merged
}

// IsNumeric checks whether the given string s is numeric.
func IsNumeric(s string) bool {
	length := len(s)
	if length == 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '-' && i == 0 {
			continue
		}
		if s[i] == '.' {
			if i > 0 && i < len(s)-1 {
				continue
			} else {
				return false
			}
		}
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

// LowerFirst lower first char
func LowerFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	rs := []rune(s)
	f := rs[0]
	if 'A' <= f && f <= 'Z' {
		return string(unicode.ToLower(f)) + string(rs[1:])
	}

	return s
}

// UpperFirst upper first char
func UpperFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	rs := []rune(s)
	f := rs[0]
	if 'a' <= f && f <= 'z' {
		return string(unicode.ToUpper(f)) + string(rs[1:])
	}

	return s
}

// SnakeCase convert. eg "RangePrice" -> "range_price"
func SnakeCase(s string, sep ...string) string {
	sepChar := "_"
	if len(sep) > 0 {
		sepChar = sep[0]
	}

	newStr := toSnakeReg.ReplaceAllStringFunc(s, func(s string) string {
		return sepChar + LowerFirst(s)
	})

	return strings.TrimLeft(newStr, sepChar)
}

// CamelCase convert string to camel case.
// Support:
// 	"range_price" -> "rangePrice"
// 	"range price" -> "rangePrice"
// 	"range-price" -> "rangePrice"
func CamelCase(s string, sep ...string) string {
	sepChar := "_"
	if len(sep) > 0 {
		sepChar = sep[0]
	}

	// Not contains sep char
	if !strings.Contains(s, sepChar) {
		return s
	}

	// Get regexp instance
	rgx, ok := toCamelRegs[sepChar]
	if !ok {
		rgx = regexp.MustCompile(regexp.QuoteMeta(sepChar) + "+[a-zA-Z]")
	}

	return rgx.ReplaceAllStringFunc(s, func(s string) string {
		s = strings.TrimLeft(s, sepChar)
		return UpperFirst(s)
	})
}
