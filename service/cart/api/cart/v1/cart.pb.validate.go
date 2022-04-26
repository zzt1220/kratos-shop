// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: cart/v1/cart.proto

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

// Validate checks the field values on CartInfoReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CartInfoReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CartInfoReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CartInfoReplyMultiError, or
// nil if none found.
func (m *CartInfoReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CartInfoReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for UserId

	// no validation rules for GoodsId

	// no validation rules for GoodsSn

	// no validation rules for GoodsName

	// no validation rules for SkuId

	// no validation rules for GoodsPrice

	// no validation rules for GoodsNum

	// no validation rules for IsSelect

	if len(errors) > 0 {
		return CartInfoReplyMultiError(errors)
	}

	return nil
}

// CartInfoReplyMultiError is an error wrapping multiple validation errors
// returned by CartInfoReply.ValidateAll() if the designated constraints
// aren't met.
type CartInfoReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CartInfoReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CartInfoReplyMultiError) AllErrors() []error { return m }

// CartInfoReplyValidationError is the validation error returned by
// CartInfoReply.Validate if the designated constraints aren't met.
type CartInfoReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CartInfoReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CartInfoReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CartInfoReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CartInfoReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CartInfoReplyValidationError) ErrorName() string { return "CartInfoReplyValidationError" }

// Error satisfies the builtin error interface
func (e CartInfoReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCartInfoReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CartInfoReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CartInfoReplyValidationError{}

// Validate checks the field values on CreateCartRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateCartRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateCartRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateCartRequestMultiError, or nil if none found.
func (m *CreateCartRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateCartRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if m.GetUserId() <= 0 {
		err := CreateCartRequestValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetGoodsId() <= 0 {
		err := CreateCartRequestValidationError{
			field:  "GoodsId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetGoodsSn()) < 1 {
		err := CreateCartRequestValidationError{
			field:  "GoodsSn",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetGoodsName()) < 1 {
		err := CreateCartRequestValidationError{
			field:  "GoodsName",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetSkuId() <= 0 {
		err := CreateCartRequestValidationError{
			field:  "SkuId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetGoodsPrice() <= 0 {
		err := CreateCartRequestValidationError{
			field:  "GoodsPrice",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetGoodsNum() <= 0 {
		err := CreateCartRequestValidationError{
			field:  "GoodsNum",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetIsSelect() != true {
		err := CreateCartRequestValidationError{
			field:  "IsSelect",
			reason: "value must equal true",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateCartRequestMultiError(errors)
	}

	return nil
}

// CreateCartRequestMultiError is an error wrapping multiple validation errors
// returned by CreateCartRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateCartRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateCartRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateCartRequestMultiError) AllErrors() []error { return m }

// CreateCartRequestValidationError is the validation error returned by
// CreateCartRequest.Validate if the designated constraints aren't met.
type CreateCartRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCartRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCartRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCartRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCartRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCartRequestValidationError) ErrorName() string {
	return "CreateCartRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCartRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCartRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCartRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCartRequestValidationError{}

// Validate checks the field values on UpdateCartRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateCartRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateCartRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateCartRequestMultiError, or nil if none found.
func (m *UpdateCartRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateCartRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for GoodsNum

	if len(errors) > 0 {
		return UpdateCartRequestMultiError(errors)
	}

	return nil
}

// UpdateCartRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateCartRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateCartRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateCartRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateCartRequestMultiError) AllErrors() []error { return m }

// UpdateCartRequestValidationError is the validation error returned by
// UpdateCartRequest.Validate if the designated constraints aren't met.
type UpdateCartRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateCartRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateCartRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateCartRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateCartRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateCartRequestValidationError) ErrorName() string {
	return "UpdateCartRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateCartRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateCartRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateCartRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateCartRequestValidationError{}

// Validate checks the field values on UpdateCartReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateCartReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateCartReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateCartReplyMultiError, or nil if none found.
func (m *UpdateCartReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateCartReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateCartReplyMultiError(errors)
	}

	return nil
}

// UpdateCartReplyMultiError is an error wrapping multiple validation errors
// returned by UpdateCartReply.ValidateAll() if the designated constraints
// aren't met.
type UpdateCartReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateCartReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateCartReplyMultiError) AllErrors() []error { return m }

// UpdateCartReplyValidationError is the validation error returned by
// UpdateCartReply.Validate if the designated constraints aren't met.
type UpdateCartReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateCartReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateCartReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateCartReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateCartReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateCartReplyValidationError) ErrorName() string { return "UpdateCartReplyValidationError" }

// Error satisfies the builtin error interface
func (e UpdateCartReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateCartReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateCartReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateCartReplyValidationError{}

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

// Validate checks the field values on DeleteCartRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteCartRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteCartRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteCartRequestMultiError, or nil if none found.
func (m *DeleteCartRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteCartRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteCartRequestMultiError(errors)
	}

	return nil
}

// DeleteCartRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteCartRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteCartRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteCartRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteCartRequestMultiError) AllErrors() []error { return m }

// DeleteCartRequestValidationError is the validation error returned by
// DeleteCartRequest.Validate if the designated constraints aren't met.
type DeleteCartRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteCartRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteCartRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteCartRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteCartRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteCartRequestValidationError) ErrorName() string {
	return "DeleteCartRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteCartRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteCartRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteCartRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteCartRequestValidationError{}

// Validate checks the field values on DeleteCartReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteCartReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteCartReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteCartReplyMultiError, or nil if none found.
func (m *DeleteCartReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteCartReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteCartReplyMultiError(errors)
	}

	return nil
}

// DeleteCartReplyMultiError is an error wrapping multiple validation errors
// returned by DeleteCartReply.ValidateAll() if the designated constraints
// aren't met.
type DeleteCartReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteCartReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteCartReplyMultiError) AllErrors() []error { return m }

// DeleteCartReplyValidationError is the validation error returned by
// DeleteCartReply.Validate if the designated constraints aren't met.
type DeleteCartReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteCartReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteCartReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteCartReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteCartReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteCartReplyValidationError) ErrorName() string { return "DeleteCartReplyValidationError" }

// Error satisfies the builtin error interface
func (e DeleteCartReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteCartReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteCartReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteCartReplyValidationError{}

// Validate checks the field values on GetCartRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetCartRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCartRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetCartRequestMultiError,
// or nil if none found.
func (m *GetCartRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCartRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetCartRequestMultiError(errors)
	}

	return nil
}

// GetCartRequestMultiError is an error wrapping multiple validation errors
// returned by GetCartRequest.ValidateAll() if the designated constraints
// aren't met.
type GetCartRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCartRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCartRequestMultiError) AllErrors() []error { return m }

// GetCartRequestValidationError is the validation error returned by
// GetCartRequest.Validate if the designated constraints aren't met.
type GetCartRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCartRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCartRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCartRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCartRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCartRequestValidationError) ErrorName() string { return "GetCartRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetCartRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCartRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCartRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCartRequestValidationError{}

// Validate checks the field values on GetCartReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetCartReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCartReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetCartReplyMultiError, or
// nil if none found.
func (m *GetCartReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCartReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetCartReplyMultiError(errors)
	}

	return nil
}

// GetCartReplyMultiError is an error wrapping multiple validation errors
// returned by GetCartReply.ValidateAll() if the designated constraints aren't met.
type GetCartReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCartReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCartReplyMultiError) AllErrors() []error { return m }

// GetCartReplyValidationError is the validation error returned by
// GetCartReply.Validate if the designated constraints aren't met.
type GetCartReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCartReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCartReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCartReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCartReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCartReplyValidationError) ErrorName() string { return "GetCartReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetCartReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCartReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCartReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCartReplyValidationError{}

// Validate checks the field values on ListCartRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListCartRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListCartRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListCartRequestMultiError, or nil if none found.
func (m *ListCartRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListCartRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListCartRequestMultiError(errors)
	}

	return nil
}

// ListCartRequestMultiError is an error wrapping multiple validation errors
// returned by ListCartRequest.ValidateAll() if the designated constraints
// aren't met.
type ListCartRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListCartRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListCartRequestMultiError) AllErrors() []error { return m }

// ListCartRequestValidationError is the validation error returned by
// ListCartRequest.Validate if the designated constraints aren't met.
type ListCartRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCartRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCartRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCartRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCartRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCartRequestValidationError) ErrorName() string { return "ListCartRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListCartRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCartRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCartRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCartRequestValidationError{}

// Validate checks the field values on CartListReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CartListReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CartListReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CartListReplyMultiError, or
// nil if none found.
func (m *CartListReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CartListReply) validate(all bool) error {
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
					errors = append(errors, CartListReplyValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CartListReplyValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CartListReplyValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CartListReplyMultiError(errors)
	}

	return nil
}

// CartListReplyMultiError is an error wrapping multiple validation errors
// returned by CartListReply.ValidateAll() if the designated constraints
// aren't met.
type CartListReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CartListReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CartListReplyMultiError) AllErrors() []error { return m }

// CartListReplyValidationError is the validation error returned by
// CartListReply.Validate if the designated constraints aren't met.
type CartListReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CartListReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CartListReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CartListReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CartListReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CartListReplyValidationError) ErrorName() string { return "CartListReplyValidationError" }

// Error satisfies the builtin error interface
func (e CartListReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCartListReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CartListReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CartListReplyValidationError{}
