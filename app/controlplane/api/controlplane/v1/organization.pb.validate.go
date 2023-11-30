// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: controlplane/v1/organization.proto

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

// define the regex for a UUID once up-front
var _organization_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on
// OrganizationServiceListMembershipsRequest with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OrganizationServiceListMembershipsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// OrganizationServiceListMembershipsRequest with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// OrganizationServiceListMembershipsRequestMultiError, or nil if none found.
func (m *OrganizationServiceListMembershipsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *OrganizationServiceListMembershipsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return OrganizationServiceListMembershipsRequestMultiError(errors)
	}

	return nil
}

// OrganizationServiceListMembershipsRequestMultiError is an error wrapping
// multiple validation errors returned by
// OrganizationServiceListMembershipsRequest.ValidateAll() if the designated
// constraints aren't met.
type OrganizationServiceListMembershipsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrganizationServiceListMembershipsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrganizationServiceListMembershipsRequestMultiError) AllErrors() []error { return m }

// OrganizationServiceListMembershipsRequestValidationError is the validation
// error returned by OrganizationServiceListMembershipsRequest.Validate if the
// designated constraints aren't met.
type OrganizationServiceListMembershipsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrganizationServiceListMembershipsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrganizationServiceListMembershipsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrganizationServiceListMembershipsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrganizationServiceListMembershipsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrganizationServiceListMembershipsRequestValidationError) ErrorName() string {
	return "OrganizationServiceListMembershipsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e OrganizationServiceListMembershipsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrganizationServiceListMembershipsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrganizationServiceListMembershipsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrganizationServiceListMembershipsRequestValidationError{}

// Validate checks the field values on
// OrganizationServiceListMembershipsResponse with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OrganizationServiceListMembershipsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// OrganizationServiceListMembershipsResponse with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// OrganizationServiceListMembershipsResponseMultiError, or nil if none found.
func (m *OrganizationServiceListMembershipsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *OrganizationServiceListMembershipsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetResult() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, OrganizationServiceListMembershipsResponseValidationError{
						field:  fmt.Sprintf("Result[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, OrganizationServiceListMembershipsResponseValidationError{
						field:  fmt.Sprintf("Result[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OrganizationServiceListMembershipsResponseValidationError{
					field:  fmt.Sprintf("Result[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return OrganizationServiceListMembershipsResponseMultiError(errors)
	}

	return nil
}

// OrganizationServiceListMembershipsResponseMultiError is an error wrapping
// multiple validation errors returned by
// OrganizationServiceListMembershipsResponse.ValidateAll() if the designated
// constraints aren't met.
type OrganizationServiceListMembershipsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrganizationServiceListMembershipsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrganizationServiceListMembershipsResponseMultiError) AllErrors() []error { return m }

// OrganizationServiceListMembershipsResponseValidationError is the validation
// error returned by OrganizationServiceListMembershipsResponse.Validate if
// the designated constraints aren't met.
type OrganizationServiceListMembershipsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrganizationServiceListMembershipsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrganizationServiceListMembershipsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrganizationServiceListMembershipsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrganizationServiceListMembershipsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrganizationServiceListMembershipsResponseValidationError) ErrorName() string {
	return "OrganizationServiceListMembershipsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e OrganizationServiceListMembershipsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrganizationServiceListMembershipsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrganizationServiceListMembershipsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrganizationServiceListMembershipsResponseValidationError{}

// Validate checks the field values on OrganizationServiceUpdateRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *OrganizationServiceUpdateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrganizationServiceUpdateRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// OrganizationServiceUpdateRequestMultiError, or nil if none found.
func (m *OrganizationServiceUpdateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *OrganizationServiceUpdateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetId()); err != nil {
		err = OrganizationServiceUpdateRequestValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.Name != nil {
		// no validation rules for Name
	}

	if len(errors) > 0 {
		return OrganizationServiceUpdateRequestMultiError(errors)
	}

	return nil
}

func (m *OrganizationServiceUpdateRequest) _validateUuid(uuid string) error {
	if matched := _organization_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// OrganizationServiceUpdateRequestMultiError is an error wrapping multiple
// validation errors returned by
// OrganizationServiceUpdateRequest.ValidateAll() if the designated
// constraints aren't met.
type OrganizationServiceUpdateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrganizationServiceUpdateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrganizationServiceUpdateRequestMultiError) AllErrors() []error { return m }

// OrganizationServiceUpdateRequestValidationError is the validation error
// returned by OrganizationServiceUpdateRequest.Validate if the designated
// constraints aren't met.
type OrganizationServiceUpdateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrganizationServiceUpdateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrganizationServiceUpdateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrganizationServiceUpdateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrganizationServiceUpdateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrganizationServiceUpdateRequestValidationError) ErrorName() string {
	return "OrganizationServiceUpdateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e OrganizationServiceUpdateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrganizationServiceUpdateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrganizationServiceUpdateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrganizationServiceUpdateRequestValidationError{}

// Validate checks the field values on OrganizationServiceUpdateResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *OrganizationServiceUpdateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrganizationServiceUpdateResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// OrganizationServiceUpdateResponseMultiError, or nil if none found.
func (m *OrganizationServiceUpdateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *OrganizationServiceUpdateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResult()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OrganizationServiceUpdateResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OrganizationServiceUpdateResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrganizationServiceUpdateResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return OrganizationServiceUpdateResponseMultiError(errors)
	}

	return nil
}

// OrganizationServiceUpdateResponseMultiError is an error wrapping multiple
// validation errors returned by
// OrganizationServiceUpdateResponse.ValidateAll() if the designated
// constraints aren't met.
type OrganizationServiceUpdateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrganizationServiceUpdateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrganizationServiceUpdateResponseMultiError) AllErrors() []error { return m }

// OrganizationServiceUpdateResponseValidationError is the validation error
// returned by OrganizationServiceUpdateResponse.Validate if the designated
// constraints aren't met.
type OrganizationServiceUpdateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrganizationServiceUpdateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrganizationServiceUpdateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrganizationServiceUpdateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrganizationServiceUpdateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrganizationServiceUpdateResponseValidationError) ErrorName() string {
	return "OrganizationServiceUpdateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e OrganizationServiceUpdateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrganizationServiceUpdateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrganizationServiceUpdateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrganizationServiceUpdateResponseValidationError{}

// Validate checks the field values on SetCurrentMembershipRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SetCurrentMembershipRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SetCurrentMembershipRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SetCurrentMembershipRequestMultiError, or nil if none found.
func (m *SetCurrentMembershipRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SetCurrentMembershipRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetMembershipId()); err != nil {
		err = SetCurrentMembershipRequestValidationError{
			field:  "MembershipId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SetCurrentMembershipRequestMultiError(errors)
	}

	return nil
}

func (m *SetCurrentMembershipRequest) _validateUuid(uuid string) error {
	if matched := _organization_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// SetCurrentMembershipRequestMultiError is an error wrapping multiple
// validation errors returned by SetCurrentMembershipRequest.ValidateAll() if
// the designated constraints aren't met.
type SetCurrentMembershipRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SetCurrentMembershipRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SetCurrentMembershipRequestMultiError) AllErrors() []error { return m }

// SetCurrentMembershipRequestValidationError is the validation error returned
// by SetCurrentMembershipRequest.Validate if the designated constraints
// aren't met.
type SetCurrentMembershipRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetCurrentMembershipRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetCurrentMembershipRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetCurrentMembershipRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetCurrentMembershipRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetCurrentMembershipRequestValidationError) ErrorName() string {
	return "SetCurrentMembershipRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SetCurrentMembershipRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetCurrentMembershipRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetCurrentMembershipRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetCurrentMembershipRequestValidationError{}

// Validate checks the field values on SetCurrentMembershipResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SetCurrentMembershipResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SetCurrentMembershipResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SetCurrentMembershipResponseMultiError, or nil if none found.
func (m *SetCurrentMembershipResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SetCurrentMembershipResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResult()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SetCurrentMembershipResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SetCurrentMembershipResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SetCurrentMembershipResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SetCurrentMembershipResponseMultiError(errors)
	}

	return nil
}

// SetCurrentMembershipResponseMultiError is an error wrapping multiple
// validation errors returned by SetCurrentMembershipResponse.ValidateAll() if
// the designated constraints aren't met.
type SetCurrentMembershipResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SetCurrentMembershipResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SetCurrentMembershipResponseMultiError) AllErrors() []error { return m }

// SetCurrentMembershipResponseValidationError is the validation error returned
// by SetCurrentMembershipResponse.Validate if the designated constraints
// aren't met.
type SetCurrentMembershipResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetCurrentMembershipResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetCurrentMembershipResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetCurrentMembershipResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetCurrentMembershipResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetCurrentMembershipResponseValidationError) ErrorName() string {
	return "SetCurrentMembershipResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SetCurrentMembershipResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetCurrentMembershipResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetCurrentMembershipResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetCurrentMembershipResponseValidationError{}
