package validator

import (
	"regexp"
	"slices"
	"strings"
	"unicode/utf8"
)

// Regular expression to check if an email address is valid.
var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Return true if value contains at least n characters
func MinChars(values string, n int) bool {
	return utf8.RuneCountInString(values) >= n
}

func MatchesPattern(rx *regexp.Regexp, value string) bool {
	return rx.MatchString(value)
}

// Define a new Validator type which contains a map of validation errors.
type Validator struct {
	FieldErrors map[string]string
}

// Returns true if there are no errors, otherwise returns false.
func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0
}

// Checks if the form field is empty. If it is empty, add the message to the
// map of errors.
func (v *Validator) AddFieldError(key, message string) {
	// initialze the map if it does not exist
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}

	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}

// Only add err msg if validation check is not 'ok'
func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldError(key, message)
	}
}

// Return true is a value is not an empty string
func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

// Return true if a value contains no more than n characters
func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

// Return true if a value is in a list of specific permitted values
func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	return slices.Contains(permittedValues, value)
}
