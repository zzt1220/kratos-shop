// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/shop/v1/shop.proto

package v1

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

// Validate checks the field values on CreateAddressReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateAddressReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateAddressReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateAddressReqMultiError, or nil if none found.
func (m *CreateAddressReq) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateAddressReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Uid

	// no validation rules for Name

	// no validation rules for Mobile

	// no validation rules for Province

	// no validation rules for City

	// no validation rules for Districts

	// no validation rules for Address

	// no validation rules for PostCode

	// no validation rules for IsDefault

	if len(errors) > 0 {
		return CreateAddressReqMultiError(errors)
	}
	return nil
}

// CreateAddressReqMultiError is an error wrapping multiple validation errors
// returned by CreateAddressReq.ValidateAll() if the designated constraints
// aren't met.
type CreateAddressReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateAddressReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateAddressReqMultiError) AllErrors() []error { return m }

// CreateAddressReqValidationError is the validation error returned by
// CreateAddressReq.Validate if the designated constraints aren't met.
type CreateAddressReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateAddressReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateAddressReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateAddressReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateAddressReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateAddressReqValidationError) ErrorName() string { return "CreateAddressReqValidationError" }

// Error satisfies the builtin error interface
func (e CreateAddressReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateAddressReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateAddressReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateAddressReqValidationError{}

// Validate checks the field values on UpdateAddressReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateAddressReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateAddressReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateAddressReqMultiError, or nil if none found.
func (m *UpdateAddressReq) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateAddressReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Uid

	// no validation rules for Name

	// no validation rules for Mobile

	// no validation rules for Province

	// no validation rules for City

	// no validation rules for Districts

	// no validation rules for Address

	// no validation rules for PostCode

	// no validation rules for IsDefault

	// no validation rules for Id

	if len(errors) > 0 {
		return UpdateAddressReqMultiError(errors)
	}
	return nil
}

// UpdateAddressReqMultiError is an error wrapping multiple validation errors
// returned by UpdateAddressReq.ValidateAll() if the designated constraints
// aren't met.
type UpdateAddressReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateAddressReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateAddressReqMultiError) AllErrors() []error { return m }

// UpdateAddressReqValidationError is the validation error returned by
// UpdateAddressReq.Validate if the designated constraints aren't met.
type UpdateAddressReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateAddressReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateAddressReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateAddressReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateAddressReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateAddressReqValidationError) ErrorName() string { return "UpdateAddressReqValidationError" }

// Error satisfies the builtin error interface
func (e UpdateAddressReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateAddressReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateAddressReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateAddressReqValidationError{}

// Validate checks the field values on AddressInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddressInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddressInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AddressInfoMultiError, or
// nil if none found.
func (m *AddressInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *AddressInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Mobile

	// no validation rules for Province

	// no validation rules for City

	// no validation rules for Districts

	// no validation rules for Address

	// no validation rules for PostCode

	// no validation rules for IsDefault

	if len(errors) > 0 {
		return AddressInfoMultiError(errors)
	}
	return nil
}

// AddressInfoMultiError is an error wrapping multiple validation errors
// returned by AddressInfo.ValidateAll() if the designated constraints aren't met.
type AddressInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddressInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddressInfoMultiError) AllErrors() []error { return m }

// AddressInfoValidationError is the validation error returned by
// AddressInfo.Validate if the designated constraints aren't met.
type AddressInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddressInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddressInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddressInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddressInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddressInfoValidationError) ErrorName() string { return "AddressInfoValidationError" }

// Error satisfies the builtin error interface
func (e AddressInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddressInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddressInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddressInfoValidationError{}

// Validate checks the field values on ListAddressReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListAddressReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListAddressReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListAddressReqMultiError,
// or nil if none found.
func (m *ListAddressReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ListAddressReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Uid

	if len(errors) > 0 {
		return ListAddressReqMultiError(errors)
	}
	return nil
}

// ListAddressReqMultiError is an error wrapping multiple validation errors
// returned by ListAddressReq.ValidateAll() if the designated constraints
// aren't met.
type ListAddressReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListAddressReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListAddressReqMultiError) AllErrors() []error { return m }

// ListAddressReqValidationError is the validation error returned by
// ListAddressReq.Validate if the designated constraints aren't met.
type ListAddressReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAddressReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAddressReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAddressReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAddressReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAddressReqValidationError) ErrorName() string { return "ListAddressReqValidationError" }

// Error satisfies the builtin error interface
func (e ListAddressReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAddressReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAddressReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAddressReqValidationError{}

// Validate checks the field values on ListAddressReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListAddressReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListAddressReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListAddressReplyMultiError, or nil if none found.
func (m *ListAddressReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListAddressReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetResults() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListAddressReplyValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListAddressReplyValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListAddressReplyValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListAddressReplyMultiError(errors)
	}
	return nil
}

// ListAddressReplyMultiError is an error wrapping multiple validation errors
// returned by ListAddressReply.ValidateAll() if the designated constraints
// aren't met.
type ListAddressReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListAddressReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListAddressReplyMultiError) AllErrors() []error { return m }

// ListAddressReplyValidationError is the validation error returned by
// ListAddressReply.Validate if the designated constraints aren't met.
type ListAddressReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListAddressReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListAddressReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListAddressReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListAddressReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListAddressReplyValidationError) ErrorName() string { return "ListAddressReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListAddressReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListAddressReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListAddressReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListAddressReplyValidationError{}

// Validate checks the field values on AddressReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddressReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddressReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AddressReqMultiError, or
// nil if none found.
func (m *AddressReq) ValidateAll() error {
	return m.validate(true)
}

func (m *AddressReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Uid

	if len(errors) > 0 {
		return AddressReqMultiError(errors)
	}
	return nil
}

// AddressReqMultiError is an error wrapping multiple validation errors
// returned by AddressReq.ValidateAll() if the designated constraints aren't met.
type AddressReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddressReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddressReqMultiError) AllErrors() []error { return m }

// AddressReqValidationError is the validation error returned by
// AddressReq.Validate if the designated constraints aren't met.
type AddressReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddressReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddressReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddressReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddressReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddressReqValidationError) ErrorName() string { return "AddressReqValidationError" }

// Error satisfies the builtin error interface
func (e AddressReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddressReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddressReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddressReqValidationError{}

// Validate checks the field values on CheckResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CheckResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CheckResponseMultiError, or
// nil if none found.
func (m *CheckResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if len(errors) > 0 {
		return CheckResponseMultiError(errors)
	}
	return nil
}

// CheckResponseMultiError is an error wrapping multiple validation errors
// returned by CheckResponse.ValidateAll() if the designated constraints
// aren't met.
type CheckResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckResponseMultiError) AllErrors() []error { return m }

// CheckResponseValidationError is the validation error returned by
// CheckResponse.Validate if the designated constraints aren't met.
type CheckResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckResponseValidationError) ErrorName() string { return "CheckResponseValidationError" }

// Error satisfies the builtin error interface
func (e CheckResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckResponseValidationError{}

// Validate checks the field values on RegisterReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RegisterReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RegisterReplyMultiError, or
// nil if none found.
func (m *RegisterReply) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Mobile

	// no validation rules for Username

	// no validation rules for Token

	// no validation rules for ExpiredAt

	if len(errors) > 0 {
		return RegisterReplyMultiError(errors)
	}
	return nil
}

// RegisterReplyMultiError is an error wrapping multiple validation errors
// returned by RegisterReply.ValidateAll() if the designated constraints
// aren't met.
type RegisterReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterReplyMultiError) AllErrors() []error { return m }

// RegisterReplyValidationError is the validation error returned by
// RegisterReply.Validate if the designated constraints aren't met.
type RegisterReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterReplyValidationError) ErrorName() string { return "RegisterReplyValidationError" }

// Error satisfies the builtin error interface
func (e RegisterReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterReplyValidationError{}

// Validate checks the field values on RegisterReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RegisterReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RegisterReqMultiError, or
// nil if none found.
func (m *RegisterReq) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetMobile()) != 11 {
		err := RegisterReqValidationError{
			field:  "Mobile",
			reason: "value length must be 11 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if l := utf8.RuneCountInString(m.GetUsername()); l < 3 || l > 15 {
		err := RegisterReqValidationError{
			field:  "Username",
			reason: "value length must be between 3 and 15 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetPassword()) < 8 {
		err := RegisterReqValidationError{
			field:  "Password",
			reason: "value length must be at least 8 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RegisterReqMultiError(errors)
	}
	return nil
}

// RegisterReqMultiError is an error wrapping multiple validation errors
// returned by RegisterReq.ValidateAll() if the designated constraints aren't met.
type RegisterReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterReqMultiError) AllErrors() []error { return m }

// RegisterReqValidationError is the validation error returned by
// RegisterReq.Validate if the designated constraints aren't met.
type RegisterReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterReqValidationError) ErrorName() string { return "RegisterReqValidationError" }

// Error satisfies the builtin error interface
func (e RegisterReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterReqValidationError{}

// Validate checks the field values on LoginReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginReqMultiError, or nil
// if none found.
func (m *LoginReq) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetMobile()) != 11 {
		err := LoginReqValidationError{
			field:  "Mobile",
			reason: "value length must be 11 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if utf8.RuneCountInString(m.GetPassword()) < 8 {
		err := LoginReqValidationError{
			field:  "Password",
			reason: "value length must be at least 8 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetCaptcha()) != 5 {
		err := LoginReqValidationError{
			field:  "Captcha",
			reason: "value length must be 5 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if utf8.RuneCountInString(m.GetCaptchaId()) < 1 {
		err := LoginReqValidationError{
			field:  "CaptchaId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return LoginReqMultiError(errors)
	}
	return nil
}

// LoginReqMultiError is an error wrapping multiple validation errors returned
// by LoginReq.ValidateAll() if the designated constraints aren't met.
type LoginReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginReqMultiError) AllErrors() []error { return m }

// LoginReqValidationError is the validation error returned by
// LoginReq.Validate if the designated constraints aren't met.
type LoginReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginReqValidationError) ErrorName() string { return "LoginReqValidationError" }

// Error satisfies the builtin error interface
func (e LoginReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginReqValidationError{}

// Validate checks the field values on UserDetailResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UserDetailResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserDetailResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserDetailResponseMultiError, or nil if none found.
func (m *UserDetailResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UserDetailResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Mobile

	// no validation rules for NickName

	// no validation rules for Birthday

	// no validation rules for Gender

	// no validation rules for Role

	if len(errors) > 0 {
		return UserDetailResponseMultiError(errors)
	}
	return nil
}

// UserDetailResponseMultiError is an error wrapping multiple validation errors
// returned by UserDetailResponse.ValidateAll() if the designated constraints
// aren't met.
type UserDetailResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserDetailResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserDetailResponseMultiError) AllErrors() []error { return m }

// UserDetailResponseValidationError is the validation error returned by
// UserDetailResponse.Validate if the designated constraints aren't met.
type UserDetailResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserDetailResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserDetailResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserDetailResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserDetailResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserDetailResponseValidationError) ErrorName() string {
	return "UserDetailResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UserDetailResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserDetailResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserDetailResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserDetailResponseValidationError{}

// Validate checks the field values on CaptchaReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CaptchaReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CaptchaReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CaptchaReplyMultiError, or
// nil if none found.
func (m *CaptchaReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CaptchaReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CaptchaId

	// no validation rules for PicPath

	if len(errors) > 0 {
		return CaptchaReplyMultiError(errors)
	}
	return nil
}

// CaptchaReplyMultiError is an error wrapping multiple validation errors
// returned by CaptchaReply.ValidateAll() if the designated constraints aren't met.
type CaptchaReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CaptchaReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CaptchaReplyMultiError) AllErrors() []error { return m }

// CaptchaReplyValidationError is the validation error returned by
// CaptchaReply.Validate if the designated constraints aren't met.
type CaptchaReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CaptchaReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CaptchaReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CaptchaReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CaptchaReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CaptchaReplyValidationError) ErrorName() string { return "CaptchaReplyValidationError" }

// Error satisfies the builtin error interface
func (e CaptchaReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCaptchaReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CaptchaReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CaptchaReplyValidationError{}
