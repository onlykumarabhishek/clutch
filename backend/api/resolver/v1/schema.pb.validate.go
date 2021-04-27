// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: resolver/v1/schema.proto

package resolverv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on StringField with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *StringField) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Placeholder

	// no validation rules for DefaultValue

	return nil
}

// StringFieldValidationError is the validation error returned by
// StringField.Validate if the designated constraints aren't met.
type StringFieldValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StringFieldValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StringFieldValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StringFieldValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StringFieldValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StringFieldValidationError) ErrorName() string { return "StringFieldValidationError" }

// Error satisfies the builtin error interface
func (e StringFieldValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStringField.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StringFieldValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StringFieldValidationError{}

// Validate checks the field values on Option with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Option) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for DisplayName

	switch m.Value.(type) {

	case *Option_StringValue:
		// no validation rules for StringValue

	}

	return nil
}

// OptionValidationError is the validation error returned by Option.Validate if
// the designated constraints aren't met.
type OptionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OptionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OptionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OptionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OptionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OptionValidationError) ErrorName() string { return "OptionValidationError" }

// Error satisfies the builtin error interface
func (e OptionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOption.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OptionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OptionValidationError{}

// Validate checks the field values on OptionField with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *OptionField) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for IncludeAllOption

	for idx, item := range m.GetOptions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OptionFieldValidationError{
					field:  fmt.Sprintf("Options[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// OptionFieldValidationError is the validation error returned by
// OptionField.Validate if the designated constraints aren't met.
type OptionFieldValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OptionFieldValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OptionFieldValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OptionFieldValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OptionFieldValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OptionFieldValidationError) ErrorName() string { return "OptionFieldValidationError" }

// Error satisfies the builtin error interface
func (e OptionFieldValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOptionField.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OptionFieldValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OptionFieldValidationError{}

// Validate checks the field values on Field with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Field) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FieldValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// FieldValidationError is the validation error returned by Field.Validate if
// the designated constraints aren't met.
type FieldValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FieldValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FieldValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FieldValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FieldValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FieldValidationError) ErrorName() string { return "FieldValidationError" }

// Error satisfies the builtin error interface
func (e FieldValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sField.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FieldValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FieldValidationError{}

// Validate checks the field values on FieldMetadata with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *FieldMetadata) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for DisplayName

	// no validation rules for Required

	switch m.Type.(type) {

	case *FieldMetadata_StringField:

		if v, ok := interface{}(m.GetStringField()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FieldMetadataValidationError{
					field:  "StringField",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *FieldMetadata_OptionField:

		if v, ok := interface{}(m.GetOptionField()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FieldMetadataValidationError{
					field:  "OptionField",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// FieldMetadataValidationError is the validation error returned by
// FieldMetadata.Validate if the designated constraints aren't met.
type FieldMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FieldMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FieldMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FieldMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FieldMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FieldMetadataValidationError) ErrorName() string { return "FieldMetadataValidationError" }

// Error satisfies the builtin error interface
func (e FieldMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFieldMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FieldMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FieldMetadataValidationError{}

// Validate checks the field values on SearchMetadata with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SearchMetadata) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Enabled

	// no validation rules for AutocompleteEnabled

	return nil
}

// SearchMetadataValidationError is the validation error returned by
// SearchMetadata.Validate if the designated constraints aren't met.
type SearchMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchMetadataValidationError) ErrorName() string { return "SearchMetadataValidationError" }

// Error satisfies the builtin error interface
func (e SearchMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchMetadataValidationError{}

// Validate checks the field values on SchemaMetadata with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SchemaMetadata) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for DisplayName

	// no validation rules for Searchable

	if v, ok := interface{}(m.GetSearch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SchemaMetadataValidationError{
				field:  "Search",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SchemaMetadataValidationError is the validation error returned by
// SchemaMetadata.Validate if the designated constraints aren't met.
type SchemaMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SchemaMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SchemaMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SchemaMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SchemaMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SchemaMetadataValidationError) ErrorName() string { return "SchemaMetadataValidationError" }

// Error satisfies the builtin error interface
func (e SchemaMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSchemaMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SchemaMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SchemaMetadataValidationError{}

// Validate checks the field values on Schema with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Schema) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for TypeUrl

	if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SchemaValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetFields() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SchemaValidationError{
					field:  fmt.Sprintf("Fields[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SchemaValidationError{
				field:  "Error",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SchemaValidationError is the validation error returned by Schema.Validate if
// the designated constraints aren't met.
type SchemaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SchemaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SchemaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SchemaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SchemaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SchemaValidationError) ErrorName() string { return "SchemaValidationError" }

// Error satisfies the builtin error interface
func (e SchemaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSchema.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SchemaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SchemaValidationError{}
