// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: bookstore.proto

package gen

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Shelf with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Shelf) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Shelf with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ShelfMultiError, or nil if none found.
func (m *Shelf) ValidateAll() error {
	return m.validate(true)
}

func (m *Shelf) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Theme

	if len(errors) > 0 {
		return ShelfMultiError(errors)
	}

	return nil
}

// ShelfMultiError is an error wrapping multiple validation errors returned by
// Shelf.ValidateAll() if the designated constraints aren't met.
type ShelfMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ShelfMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ShelfMultiError) AllErrors() []error { return m }

// ShelfValidationError is the validation error returned by Shelf.Validate if
// the designated constraints aren't met.
type ShelfValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ShelfValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ShelfValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ShelfValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ShelfValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ShelfValidationError) ErrorName() string { return "ShelfValidationError" }

// Error satisfies the builtin error interface
func (e ShelfValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sShelf.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ShelfValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ShelfValidationError{}

// Validate checks the field values on Book with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Book) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Book with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in BookMultiError, or nil if none found.
func (m *Book) ValidateAll() error {
	return m.validate(true)
}

func (m *Book) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Author

	// no validation rules for Title

	if len(errors) > 0 {
		return BookMultiError(errors)
	}

	return nil
}

// BookMultiError is an error wrapping multiple validation errors returned by
// Book.ValidateAll() if the designated constraints aren't met.
type BookMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BookMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BookMultiError) AllErrors() []error { return m }

// BookValidationError is the validation error returned by Book.Validate if the
// designated constraints aren't met.
type BookValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BookValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BookValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BookValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BookValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BookValidationError) ErrorName() string { return "BookValidationError" }

// Error satisfies the builtin error interface
func (e BookValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBook.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BookValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BookValidationError{}

// Validate checks the field values on ListShelvesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListShelvesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListShelvesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListShelvesResponseMultiError, or nil if none found.
func (m *ListShelvesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListShelvesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetShelves() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListShelvesResponseValidationError{
						field:  fmt.Sprintf("Shelves[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListShelvesResponseValidationError{
						field:  fmt.Sprintf("Shelves[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListShelvesResponseValidationError{
					field:  fmt.Sprintf("Shelves[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListShelvesResponseMultiError(errors)
	}

	return nil
}

// ListShelvesResponseMultiError is an error wrapping multiple validation
// errors returned by ListShelvesResponse.ValidateAll() if the designated
// constraints aren't met.
type ListShelvesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListShelvesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListShelvesResponseMultiError) AllErrors() []error { return m }

// ListShelvesResponseValidationError is the validation error returned by
// ListShelvesResponse.Validate if the designated constraints aren't met.
type ListShelvesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListShelvesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListShelvesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListShelvesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListShelvesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListShelvesResponseValidationError) ErrorName() string {
	return "ListShelvesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListShelvesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListShelvesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListShelvesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListShelvesResponseValidationError{}

// Validate checks the field values on CreateShelfRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateShelfRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateShelfRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateShelfRequestMultiError, or nil if none found.
func (m *CreateShelfRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateShelfRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetShelf()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateShelfRequestValidationError{
					field:  "Shelf",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateShelfRequestValidationError{
					field:  "Shelf",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetShelf()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateShelfRequestValidationError{
				field:  "Shelf",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateShelfRequestMultiError(errors)
	}

	return nil
}

// CreateShelfRequestMultiError is an error wrapping multiple validation errors
// returned by CreateShelfRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateShelfRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateShelfRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateShelfRequestMultiError) AllErrors() []error { return m }

// CreateShelfRequestValidationError is the validation error returned by
// CreateShelfRequest.Validate if the designated constraints aren't met.
type CreateShelfRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateShelfRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateShelfRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateShelfRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateShelfRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateShelfRequestValidationError) ErrorName() string {
	return "CreateShelfRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateShelfRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateShelfRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateShelfRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateShelfRequestValidationError{}

// Validate checks the field values on GetShelfRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetShelfRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetShelfRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetShelfRequestMultiError, or nil if none found.
func (m *GetShelfRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetShelfRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Shelf

	if len(errors) > 0 {
		return GetShelfRequestMultiError(errors)
	}

	return nil
}

// GetShelfRequestMultiError is an error wrapping multiple validation errors
// returned by GetShelfRequest.ValidateAll() if the designated constraints
// aren't met.
type GetShelfRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetShelfRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetShelfRequestMultiError) AllErrors() []error { return m }

// GetShelfRequestValidationError is the validation error returned by
// GetShelfRequest.Validate if the designated constraints aren't met.
type GetShelfRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetShelfRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetShelfRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetShelfRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetShelfRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetShelfRequestValidationError) ErrorName() string { return "GetShelfRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetShelfRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetShelfRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetShelfRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetShelfRequestValidationError{}

// Validate checks the field values on DeleteShelfRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteShelfRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteShelfRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteShelfRequestMultiError, or nil if none found.
func (m *DeleteShelfRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteShelfRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Shelf

	if len(errors) > 0 {
		return DeleteShelfRequestMultiError(errors)
	}

	return nil
}

// DeleteShelfRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteShelfRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteShelfRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteShelfRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteShelfRequestMultiError) AllErrors() []error { return m }

// DeleteShelfRequestValidationError is the validation error returned by
// DeleteShelfRequest.Validate if the designated constraints aren't met.
type DeleteShelfRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteShelfRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteShelfRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteShelfRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteShelfRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteShelfRequestValidationError) ErrorName() string {
	return "DeleteShelfRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteShelfRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteShelfRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteShelfRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteShelfRequestValidationError{}

// Validate checks the field values on CreateBookRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateBookRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateBookRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateBookRequestMultiError, or nil if none found.
func (m *CreateBookRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateBookRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Shelf

	if all {
		switch v := interface{}(m.GetBook()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateBookRequestValidationError{
					field:  "Book",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateBookRequestValidationError{
					field:  "Book",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBook()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateBookRequestValidationError{
				field:  "Book",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateBookRequestMultiError(errors)
	}

	return nil
}

// CreateBookRequestMultiError is an error wrapping multiple validation errors
// returned by CreateBookRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateBookRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateBookRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateBookRequestMultiError) AllErrors() []error { return m }

// CreateBookRequestValidationError is the validation error returned by
// CreateBookRequest.Validate if the designated constraints aren't met.
type CreateBookRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateBookRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateBookRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateBookRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateBookRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateBookRequestValidationError) ErrorName() string {
	return "CreateBookRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateBookRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateBookRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateBookRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateBookRequestValidationError{}
