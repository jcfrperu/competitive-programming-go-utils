package cp

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// thanks to and based on https://github.com/agrison/go-commons-lang

// AppendIfMissing appends a suffix to a string if missing.
func AppendIfMissing(str string, suffix string, suffixes ...string) string {
	if str == "" || EndsWith(str, suffix) {
		return str
	}
	for i := range suffixes {
		if EndsWith(str, suffixes[i]) {
			return str
		}
	}
	return str + suffix
}

// AppendIfMissingIgnoreCase appends a suffix to a string if missing (ignoring case).
func AppendIfMissingIgnoreCase(str string, suffix string, suffixes ...string) string {
	if str == "" || EndsWithIgnoreCase(str, suffix) {
		return str
	}
	for i := range suffixes {
		if EndsWithIgnoreCase(str, suffixes[i]) {
			return str
		}
	}
	return str + suffix
}

// Chomp removes one newline from end of a string if it's there, otherwise leave it alone.
// A newline is "\n", "\r", or "\r\n".
func Chomp(str string) string {
	if str == "" {
		return str
	}
	if len(str) == 1 && (str[0:1] == "\n" || str[0:1] == "\r") {
		return ""
	}
	if EndsWith(str, "\r\n") {
		return str[:len(str)-2]
	}
	if EndsWith(str, "\n") || EndsWith(str, "\r") {
		return str[:len(str)-1]
	}
	return str
}

// Contains checks if string contains a search string.
func Contains(str string, search string) bool {
	return strings.Contains(str, search)
}

// ContainsAny checks if the string contains any of the string in the given array.
func ContainsAny(str string, search ...string) bool {
	for _, s := range search {
		if Contains(str, s) {
			return true
		}
	}
	return false
}

// ContainsAnyCharacter checks if the string contains any of the character in the given string.
func ContainsAnyCharacter(str string, search string) bool {
	for _, c := range search {
		if Contains(str, (string)(c)) {
			return true
		}
	}
	return false
}

// ContainsIgnoreCase checks if the string contains the searched string ignoring case.
func ContainsIgnoreCase(str string, search string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(search))
}

// ContainsNone checks if the string contains no occurrence of searched string.
func ContainsNone(str string, search ...string) bool {
	return !ContainsAny(str, search...)
}

// ContainsNoneCharacter checks if the string contains no occurrence of searched string.
func ContainsNoneCharacter(str string, search string) bool {
	return !ContainsAnyCharacter(str, search)
}

func stringInSlice(a string, list []string) bool {
	for i := range list {
		if list[i] == a {
			return true
		}
	}
	return false
}

// ContainsOnly checks if a string contains only some strings.
func ContainsOnly(str string, search ...string) bool {
	for _, c := range str {
		if !stringInSlice((string)(c), search) {
			return false
		}
	}
	return true
}

// IsAllLowerCase checks if the string contains only lowercase characters.
func IsAllLowerCase(str string) bool {
	if str == "" {
		return false
	}
	for _, c := range str {
		if !unicode.IsLower(c) {
			return false
		}
	}
	return true
}

// IsAllUpperCase checks if the string contains only uppercase characters.
func IsAllUpperCase(str string) bool {
	if str == "" {
		return false
	}
	for _, c := range str {
		if !unicode.IsUpper(c) {
			return false
		}
	}
	return true
}

// IsAlpha checks if the string contains only Unicode letters.
func IsAlpha(str string) bool {
	if str == "" {
		return false
	}
	for _, c := range str {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}

// IsAlphanumeric checks if the string contains only Unicode letters and digits.
func IsAlphanumeric(str string) bool {
	if str == "" {
		return false
	}
	for _, c := range str {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

// IsAlphaSpace checks if the string contains only Unicode letters and spaces.
func IsAlphaSpace(str string) bool {
	if str == "" {
		return false
	}
	for _, c := range str {
		if !unicode.IsLetter(c) && !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

// IsAlphanumericSpace checks if the string contains only Unicode letters, digits and spaces.
func IsAlphanumericSpace(str string) bool {
	if str == "" {
		return false
	}
	for _, c := range str {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) && !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

// IsEmpty checks if a string is empty.
func IsEmpty(s string) bool {
	if s == "" {
		return true
	}
	return false
}

// IsNotEmpty checks if a string is not empty.
func IsNotEmpty(s string) bool {
	return s != ""
}

// IsAnyEmpty checks if any one of the given strings are empty.
func IsAnyEmpty(strings ...string) bool {
	for i := range strings {
		if strings[i] == "" {
			return true
		}
	}
	return false
}

// IsNoneEmpty checks if none of the strings are empty.
func IsNoneEmpty(strings ...string) bool {
	for i := range strings {
		if strings[i] == "" {
			return false
		}
	}
	return true
}

// IsBlank checks if a string is whitespace or empty
func IsBlank(s string) bool {
	if s == "" {
		return true
	}
	if regexp.MustCompile(`^\s+$`).MatchString(s) {
		return true
	}
	return false
}

// IsNotBlank checks if a string is not empty or containing only whitespaces.
func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

// IsAnyBlank checks if any one of the strings are empty or containing only whitespaces.
func IsAnyBlank(strings ...string) bool {
	for i := range strings {
		if IsBlank(strings[i]) {
			return true
		}
	}
	return false
}

// IsNoneBlank checks if none of the strings are empty or containing only whitespaces.
func IsNoneBlank(strings ...string) bool {
	for i := range strings {
		if IsBlank(strings[i]) {
			return false
		}
	}
	return true
}

// IsNumeric checks if the string contains only digits.
func IsNumeric(str string) bool {
	for _, c := range str {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

// IsNumericSpace checks if the string contains only digits and whitespace.
func IsNumericSpace(str string) bool {
	for _, c := range str {
		if !unicode.IsDigit(c) && !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

// IsWhitespace checks if the string contains only whitespace.
func IsWhitespace(str string) bool {
	for _, c := range str {
		if !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

// IndexAt return first match of substr starting from fromIndex
func IndexAt(s, substr string, fromIndex int) int {
	if fromIndex >= 0 && fromIndex < len(s) {
		index := strings.Index(s[fromIndex:], substr)
		if index >= 0 {
			return index + fromIndex
		}
	}
	return -1
}

// Left gets the leftmost len characters of a string.
func Left(str string, size int) string {
	if str == "" || size < 0 {
		return ""
	}
	if len(str) <= size {
		return str
	}
	return str[0:size]
}

// LowerCase converts a string to lower case.
func LowerCase(str string) string {
	return strings.ToLower(str)
}

// UpperCase converts a string to upper case.
func UpperCase(str string) string {
	return strings.ToUpper(str)
}

// Mid gets size characters from the middle of a string.
func Mid(str string, pos int, size int) string {
	if str == "" || size < 0 || pos > len(str) {
		return ""
	}
	if pos < 0 {
		pos = 0
	}
	if len(str) <= pos+size {
		return str[pos:]
	}
	return str[pos : pos+size]
}

// Overlay overlays part of a string with another string.
func Overlay(str string, overlay string, start int, end int) string {
	strLen := len(str)
	// guards
	if start < 0 {
		start = 0
	}
	if start > strLen {
		start = strLen
	}
	if end < 0 {
		end = 0
	}
	if end > strLen {
		end = strLen
	}
	if start > end {
		start, end = end, start
	}
	return str[:start] + overlay + str[end:]
}

// Remove removes all occurrences of a substring from within the source string.
func Remove(str string, remove string) string {
	if str == "" {
		return str
	}
	index := strings.Index(str, remove)
	for index > -1 {
		str = str[:index] + str[index+len(remove):]
		index = strings.Index(str, remove)
	}
	return str
}

// RemoveEnd removes a substring only if it is at the end of a source string,
// otherwise returns the source string.
func RemoveEnd(str string, remove string) string {
	if str == "" || remove == "" {
		return str
	}
	if EndsWith(str, remove) {
		return str[:len(str)-len(remove)]
	}
	return str
}

// RemoveEndIgnoreCase is the case in sensitive removal of a substring if it is at
// the end of a source string, otherwise returns the source string.
func RemoveEndIgnoreCase(str string, remove string) string {
	if str == "" || remove == "" {
		return str
	}
	if EndsWithIgnoreCase(str, remove) {
		return str[:len(str)-len(remove)]
	}
	return str
}

// RemovePattern removes each substring of the source string that matches
// the given regular expression
func RemovePattern(str string, pattern string) string {
	return regexp.MustCompile(pattern).ReplaceAllString(str, "")
}

// RemoveStart removes a substring only if it is at the beginning of a source string,
// otherwise returns the source string
func RemoveStart(str string, remove string) string {
	if str == "" || remove == "" {
		return str
	}
	if StartsWith(str, remove) {
		return str[len(remove)+1:]
	}
	return str
}

// RemoveStartIgnoreCase is the case in sensitive removal of a substring if it is at
// the beginning of a source string, otherwise returns the source string.
func RemoveStartIgnoreCase(str string, remove string) string {
	if str == "" || remove == "" {
		return str
	}
	if StartsWithIgnoreCase(str, remove) {
		return str[len(remove)+1:]
	}
	return str
}

// Repeat repeats a string `repeat` times to form a new string.
func Repeat(str string, repeat int) string {
	buff := ""
	for repeat > 0 {
		repeat = repeat - 1
		buff += str
	}
	return buff
}

// RepeatWithSeparator repeats a string `repeat` times to form a new String,
// with a string separator injected each time.
func RepeatWithSeparator(str string, sep string, repeat int) string {
	buff := ""
	for repeat > 0 {
		repeat = repeat - 1
		buff += str
		if repeat > 0 {
			buff += sep
		}
	}
	return buff
}

// Reverse reverses a string.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// ReverseDelimited reverses a string separated by a delimiter.
func ReverseDelimited(str string, del string) string {
	s := strings.Split(str, del)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return strings.Join(s, del)
}

// Right gets the rightmost len characters of a string.
func Right(str string, size int) string {
	if str == "" || size < 0 {
		return ""
	}
	if len(str) <= size {
		return str
	}
	return str[len(str)-size:]
}

// Strip strips whitespace from the start and end of a String.
func Strip(str string) string {
	return regexp.MustCompile(`^\s+|\s+$`).ReplaceAllString(str, "")
}

// StripEnd strips whitespace from the end of a String.
func StripEnd(str string) string {
	return regexp.MustCompile(`\s+$`).ReplaceAllString(str, "")
}

// StripStart strips whitespace from the start of a String.
func StripStart(str string) string {
	return regexp.MustCompile(`^\s+`).ReplaceAllString(str, "")
}

// SubstringAfter gets the substring after the first occurrence of a separator.
func SubstringAfter(str string, sep string) string {
	idx := strings.Index(str, sep)
	if idx == -1 {
		return str
	}
	return str[idx+len(sep):]
}

// SubstringAfterLast gets the substring after the last occurrence of a separator.
func SubstringAfterLast(str string, sep string) string {
	idx := strings.LastIndex(str, sep)
	if idx == -1 {
		return str
	}
	return str[idx+len(sep):]
}

// SubstringBefore gets the substring before the first occurrence of a separator.
func SubstringBefore(str string, sep string) string {
	idx := strings.Index(str, sep)
	if idx == -1 {
		return str
	}
	return str[:idx]
}

// SubstringBeforeLast gets the substring before the last occurrence of a separator.
func SubstringBeforeLast(str string, sep string) string {
	idx := strings.LastIndex(str, sep)
	if idx == -1 {
		return str
	}
	return str[:idx]
}

// Trim removes control characters from both ends of this string.
func Trim(str string) string {
	return strings.TrimSpace(str)
}

// startsWith internal method to check if a string starts with a specified prefix ignoring case or not.
func startsWith(str string, prefix string, ignoreCase bool) bool {
	if str == "" || prefix == "" {
		return str == "" && prefix == ""
	}
	if utf8.RuneCountInString(prefix) > utf8.RuneCountInString(str) {
		return false
	}
	if ignoreCase {
		return strings.HasPrefix(strings.ToLower(str), strings.ToLower(prefix))
	}
	return strings.HasPrefix(str, prefix)
}

// StartsWith check if a string starts with a specified prefix.
func StartsWith(str string, prefix string) bool {
	return startsWith(str, prefix, false)
}

// StartsWithIgnoreCase case in sensitive check if a string starts with a specified prefix.
func StartsWithIgnoreCase(str string, prefix string) bool {
	return startsWith(str, prefix, true)
}

// StartsWithAny check if a string starts with any of an array of specified strings.
func StartsWithAny(str string, prefixes ...string) bool {
	for i := range prefixes {
		if startsWith(str, prefixes[i], false) {
			return true
		}
	}
	return false
}

// StartsWithAnyIgnoreCase check if a string starts with any of an array of specified strings (ignoring case).
func StartsWithAnyIgnoreCase(str string, prefixes ...string) bool {
	for i := range prefixes {
		if startsWith(str, prefixes[i], true) {
			return true
		}
	}
	return false
}

// Internal method to check if a string ends with a specified suffix ignoring case or not.
func endsWith(str string, suffix string, ignoreCase bool) bool {
	if str == "" || suffix == "" {
		return str == "" && suffix == ""
	}
	if utf8.RuneCountInString(suffix) > utf8.RuneCountInString(str) {
		return false
	}
	if ignoreCase {
		return strings.HasSuffix(strings.ToLower(str), strings.ToLower(suffix))
	}
	return strings.HasSuffix(str, suffix)
}

// EndsWith check if a string ends with a specified suffix.
func EndsWith(str string, suffix string) bool {
	return endsWith(str, suffix, false)
}

// EndsWithIgnoreCase case in sensitive check if a string ends with a specified suffix.
func EndsWithIgnoreCase(str string, suffix string) bool {
	return endsWith(str, suffix, true)
}

// EndsWithAny check if a string ends with any of an array of specified strings.
func EndsWithAny(str string, suffixes ...string) bool {
	for i := range suffixes {
		if endsWith(str, suffixes[i], false) {
			return true
		}
	}
	return false
}

// EndsWithAnyIgnoreCase check if a string ends with any of an array of specified strings (ignoring case).
func EndsWithAnyIgnoreCase(str string, suffixes ...string) bool {
	for i := range suffixes {
		if endsWith(str, suffixes[i], true) {
			return true
		}
	}
	return false
}

// DefaultIfEmpty returns either the passed in String, or if the String is Empty, the value of defaultStr.
func DefaultIfEmpty(str string, defaultStr string) string {
	if str == "" {
		return defaultStr
	}
	return str
}

// DefaultIfBlank returns either the passed in String, or if the String is Blank, the value of defaultStr.
func DefaultIfBlank(str string, defaultStr string) string {
	if IsBlank(str) {
		return defaultStr
	}
	return str
}

// prependIfMissing prepends the prefix to the start of the string if the string does not already start with any of the prefixes.
func prependIfMissing(str string, prefix string, ignoreCase bool, prefixes ...string) string {
	if prefix == "" || startsWith(str, prefix, ignoreCase) {
		return str
	}
	for i := range prefixes {
		if prefixes[i] == "" || startsWith(str, prefixes[i], ignoreCase) {
			return str
		}
	}
	return prefix + str
}

// PrependIfMissing prepends the prefix to the start of the string if the string does not already start with any of the prefixes.
func PrependIfMissing(str string, prefix string, prefixes ...string) string {
	return prependIfMissing(str, prefix, false, prefixes...)
}

// PrependIfMissingIgnoreCase prepends the prefix to the start of the string if the string does not already start, case-in sensitive, with any of the prefixes.
func PrependIfMissingIgnoreCase(str string, prefix string, prefixes ...string) string {
	return prependIfMissing(str, prefix, true, prefixes...)
}
