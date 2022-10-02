package forms

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

// Holds the errors, embedding url.Values for convenience
type Form struct {
	url.Values
	Errors errors
}

// Initialize a new Form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Mark the given fields as required.
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank.")
		}
	}
}

// Set a max length for a given field.
func (f *Form) MaxLength(field string, n int) {
	value := f.Get(field)

	if value == "" {
		return
	}

	if utf8.RuneCountInString(value) > n {
		f.Errors.Add(field, fmt.Sprintf("This field is too long (maximum %d characters).", n))
	}
}

// Restrict a field to a given set of options.
func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)

	if value == "" {
		return
	}

	for _, opt := range opts {
		if value == opt {
			return
		}
	}

	f.Errors.Add(field, `This field is invalid`)
}

// Restrict a field to a minimum length.
func (f *Form) MinLength(field string, n int) {
	value := f.Get(field)

	if value == "" {
		return
	}

	if utf8.RuneCountInString(value) < n {
		f.Errors.Add(field, fmt.Sprintf("This field is too short (minimum %d characters).", n))
	}
}

// Restrict a field to a RegEx patttern.
func (f *Form) MatchesPattern(field string, pattern *regexp.Regexp) {
	value := f.Get(field)

	if value == "" {
		return
	}

	if !pattern.MatchString(value) {
		f.Errors.Add(field, "This field is invalid")
	}
}

// Check if the form has no errors.
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
