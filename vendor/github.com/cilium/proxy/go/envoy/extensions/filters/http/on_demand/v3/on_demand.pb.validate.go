// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/on_demand/v3/on_demand.proto

package on_demandv3

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

// Validate checks the field values on OnDemandCds with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OnDemandCds) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OnDemandCds with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OnDemandCdsMultiError, or
// nil if none found.
func (m *OnDemandCds) ValidateAll() error {
	return m.validate(true)
}

func (m *OnDemandCds) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetSource() == nil {
		err := OnDemandCdsValidationError{
			field:  "Source",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetSource()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OnDemandCdsValidationError{
					field:  "Source",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OnDemandCdsValidationError{
					field:  "Source",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OnDemandCdsValidationError{
				field:  "Source",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ResourcesLocator

	if all {
		switch v := interface{}(m.GetTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OnDemandCdsValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OnDemandCdsValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OnDemandCdsValidationError{
				field:  "Timeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return OnDemandCdsMultiError(errors)
	}

	return nil
}

// OnDemandCdsMultiError is an error wrapping multiple validation errors
// returned by OnDemandCds.ValidateAll() if the designated constraints aren't met.
type OnDemandCdsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OnDemandCdsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OnDemandCdsMultiError) AllErrors() []error { return m }

// OnDemandCdsValidationError is the validation error returned by
// OnDemandCds.Validate if the designated constraints aren't met.
type OnDemandCdsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OnDemandCdsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OnDemandCdsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OnDemandCdsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OnDemandCdsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OnDemandCdsValidationError) ErrorName() string { return "OnDemandCdsValidationError" }

// Error satisfies the builtin error interface
func (e OnDemandCdsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOnDemandCds.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OnDemandCdsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OnDemandCdsValidationError{}

// Validate checks the field values on OnDemand with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OnDemand) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OnDemand with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OnDemandMultiError, or nil
// if none found.
func (m *OnDemand) ValidateAll() error {
	return m.validate(true)
}

func (m *OnDemand) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetOdcds()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OnDemandValidationError{
					field:  "Odcds",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OnDemandValidationError{
					field:  "Odcds",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOdcds()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OnDemandValidationError{
				field:  "Odcds",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return OnDemandMultiError(errors)
	}

	return nil
}

// OnDemandMultiError is an error wrapping multiple validation errors returned
// by OnDemand.ValidateAll() if the designated constraints aren't met.
type OnDemandMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OnDemandMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OnDemandMultiError) AllErrors() []error { return m }

// OnDemandValidationError is the validation error returned by
// OnDemand.Validate if the designated constraints aren't met.
type OnDemandValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OnDemandValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OnDemandValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OnDemandValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OnDemandValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OnDemandValidationError) ErrorName() string { return "OnDemandValidationError" }

// Error satisfies the builtin error interface
func (e OnDemandValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOnDemand.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OnDemandValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OnDemandValidationError{}

// Validate checks the field values on PerRouteConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PerRouteConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PerRouteConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PerRouteConfigMultiError,
// or nil if none found.
func (m *PerRouteConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *PerRouteConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetOdcds()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PerRouteConfigValidationError{
					field:  "Odcds",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PerRouteConfigValidationError{
					field:  "Odcds",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOdcds()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PerRouteConfigValidationError{
				field:  "Odcds",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PerRouteConfigMultiError(errors)
	}

	return nil
}

// PerRouteConfigMultiError is an error wrapping multiple validation errors
// returned by PerRouteConfig.ValidateAll() if the designated constraints
// aren't met.
type PerRouteConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PerRouteConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PerRouteConfigMultiError) AllErrors() []error { return m }

// PerRouteConfigValidationError is the validation error returned by
// PerRouteConfig.Validate if the designated constraints aren't met.
type PerRouteConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PerRouteConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PerRouteConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PerRouteConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PerRouteConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PerRouteConfigValidationError) ErrorName() string { return "PerRouteConfigValidationError" }

// Error satisfies the builtin error interface
func (e PerRouteConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPerRouteConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PerRouteConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PerRouteConfigValidationError{}