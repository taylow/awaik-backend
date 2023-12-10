// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/monitor/v1/query_service.proto

package monitorv1

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

// Validate checks the field values on ReadRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ReadRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReadRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ReadRequestMultiError, or
// nil if none found.
func (m *ReadRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ReadRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return ReadRequestMultiError(errors)
	}

	return nil
}

// ReadRequestMultiError is an error wrapping multiple validation errors
// returned by ReadRequest.ValidateAll() if the designated constraints aren't met.
type ReadRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReadRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReadRequestMultiError) AllErrors() []error { return m }

// ReadRequestValidationError is the validation error returned by
// ReadRequest.Validate if the designated constraints aren't met.
type ReadRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadRequestValidationError) ErrorName() string { return "ReadRequestValidationError" }

// Error satisfies the builtin error interface
func (e ReadRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadRequestValidationError{}

// Validate checks the field values on ReadResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ReadResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReadResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ReadResponseMultiError, or
// nil if none found.
func (m *ReadResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ReadResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if all {
		switch v := interface{}(m.GetMonitor()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ReadResponseValidationError{
					field:  "Monitor",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ReadResponseValidationError{
					field:  "Monitor",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMonitor()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ReadResponseValidationError{
				field:  "Monitor",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ReadResponseMultiError(errors)
	}

	return nil
}

// ReadResponseMultiError is an error wrapping multiple validation errors
// returned by ReadResponse.ValidateAll() if the designated constraints aren't met.
type ReadResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReadResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReadResponseMultiError) AllErrors() []error { return m }

// ReadResponseValidationError is the validation error returned by
// ReadResponse.Validate if the designated constraints aren't met.
type ReadResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadResponseValidationError) ErrorName() string { return "ReadResponseValidationError" }

// Error satisfies the builtin error interface
func (e ReadResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadResponseValidationError{}

// Validate checks the field values on ListByProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListByProjectRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListByProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListByProjectRequestMultiError, or nil if none found.
func (m *ListByProjectRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListByProjectRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ProjectId

	if len(errors) > 0 {
		return ListByProjectRequestMultiError(errors)
	}

	return nil
}

// ListByProjectRequestMultiError is an error wrapping multiple validation
// errors returned by ListByProjectRequest.ValidateAll() if the designated
// constraints aren't met.
type ListByProjectRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListByProjectRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListByProjectRequestMultiError) AllErrors() []error { return m }

// ListByProjectRequestValidationError is the validation error returned by
// ListByProjectRequest.Validate if the designated constraints aren't met.
type ListByProjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListByProjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListByProjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListByProjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListByProjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListByProjectRequestValidationError) ErrorName() string {
	return "ListByProjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListByProjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListByProjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListByProjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListByProjectRequestValidationError{}

// Validate checks the field values on ListByProjectResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListByProjectResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListByProjectResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListByProjectResponseMultiError, or nil if none found.
func (m *ListByProjectResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListByProjectResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	for idx, item := range m.GetMonitors() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListByProjectResponseValidationError{
						field:  fmt.Sprintf("Monitors[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListByProjectResponseValidationError{
						field:  fmt.Sprintf("Monitors[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListByProjectResponseValidationError{
					field:  fmt.Sprintf("Monitors[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListByProjectResponseMultiError(errors)
	}

	return nil
}

// ListByProjectResponseMultiError is an error wrapping multiple validation
// errors returned by ListByProjectResponse.ValidateAll() if the designated
// constraints aren't met.
type ListByProjectResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListByProjectResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListByProjectResponseMultiError) AllErrors() []error { return m }

// ListByProjectResponseValidationError is the validation error returned by
// ListByProjectResponse.Validate if the designated constraints aren't met.
type ListByProjectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListByProjectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListByProjectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListByProjectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListByProjectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListByProjectResponseValidationError) ErrorName() string {
	return "ListByProjectResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListByProjectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListByProjectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListByProjectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListByProjectResponseValidationError{}
