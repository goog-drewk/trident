// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EmsActionParameter ems action parameter
//
// swagger:model ems_action_parameter
type EmsActionParameter struct {

	// description
	Description *EmsActionParameterInlineDescription `json:"description,omitempty"`

	// Specifies the possible values of the parameter.
	// Example: ["value-1","value-2"]
	// Read Only: true
	EmsActionParameterInlineEnum []*string `json:"enum,omitempty"`

	// Specifies whether the "maximum" value is excluded in the parameter value range.
	// Read Only: true
	ExclusiveMaximum *bool `json:"exclusiveMaximum,omitempty"`

	// Specifies whether the "minimum" value is excluded in the parameter value range.
	// Read Only: true
	ExclusiveMinimum *bool `json:"exclusiveMinimum,omitempty"`

	// An optional modifier that serves as a hint at the content and format of the parameter.
	// Example: date-time
	// Read Only: true
	Format *string `json:"format,omitempty"`

	// help
	Help *EmsActionParameterInlineHelp `json:"help,omitempty"`

	// If the type of the parameter is an array, this specifies the type of items in the form of a JSON object, {"type":"type-value"}, where the type-value is one of the values for the type property.
	// Example: {"type":"string"}
	// Read Only: true
	Items *string `json:"items,omitempty"`

	// Specifies the maximum length of an array type parameter.
	// Read Only: true
	MaxItems *int64 `json:"maxItems,omitempty"`

	// Specifies the maximum length of a string type parameter.
	// Read Only: true
	MaxLength *int64 `json:"maxLength,omitempty"`

	// Specifies the maximum value of the parameter.
	// Read Only: true
	Maximum *int64 `json:"maximum,omitempty"`

	// Specifies the minimum length of an array type parameter.
	// Read Only: true
	MinItems *int64 `json:"minItems,omitempty"`

	// Specifies the minimum length of a string type parameter.
	// Read Only: true
	MinLength *int64 `json:"minLength,omitempty"`

	// Specifies the minimum value of the parameter.
	// Read Only: true
	Minimum *int64 `json:"minimum,omitempty"`

	// Parameter name.
	// Example: schedule-at
	// Read Only: true
	Name *string `json:"name,omitempty"`

	// Specifies where the parameter is placed when invoking the action.
	// Example: body
	// Read Only: true
	// Enum: [body query]
	ParamIn *string `json:"param_in,omitempty"`

	// title
	Title *EmsActionParameterInlineTitle `json:"title,omitempty"`

	// Parameter type.
	// Example: string
	// Read Only: true
	// Enum: [string number integer boolean array]
	Type *string `json:"type,omitempty"`

	// validation error message
	ValidationErrorMessage *EmsActionParameterInlineValidationErrorMessage `json:"validation_error_message,omitempty"`
}

// Validate validates this ems action parameter
func (m *EmsActionParameter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHelp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParamIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidationErrorMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsActionParameter) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if m.Description != nil {
		if err := m.Description.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("description")
			}
			return err
		}
	}

	return nil
}

func (m *EmsActionParameter) validateHelp(formats strfmt.Registry) error {
	if swag.IsZero(m.Help) { // not required
		return nil
	}

	if m.Help != nil {
		if err := m.Help.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("help")
			}
			return err
		}
	}

	return nil
}

var emsActionParameterTypeParamInPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["body","query"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		emsActionParameterTypeParamInPropEnum = append(emsActionParameterTypeParamInPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// ems_action_parameter
	// EmsActionParameter
	// param_in
	// ParamIn
	// body
	// END DEBUGGING
	// EmsActionParameterParamInBody captures enum value "body"
	EmsActionParameterParamInBody string = "body"

	// BEGIN DEBUGGING
	// ems_action_parameter
	// EmsActionParameter
	// param_in
	// ParamIn
	// query
	// END DEBUGGING
	// EmsActionParameterParamInQuery captures enum value "query"
	EmsActionParameterParamInQuery string = "query"
)

// prop value enum
func (m *EmsActionParameter) validateParamInEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, emsActionParameterTypeParamInPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EmsActionParameter) validateParamIn(formats strfmt.Registry) error {
	if swag.IsZero(m.ParamIn) { // not required
		return nil
	}

	// value enum
	if err := m.validateParamInEnum("param_in", "body", *m.ParamIn); err != nil {
		return err
	}

	return nil
}

func (m *EmsActionParameter) validateTitle(formats strfmt.Registry) error {
	if swag.IsZero(m.Title) { // not required
		return nil
	}

	if m.Title != nil {
		if err := m.Title.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("title")
			}
			return err
		}
	}

	return nil
}

var emsActionParameterTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["string","number","integer","boolean","array"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		emsActionParameterTypeTypePropEnum = append(emsActionParameterTypeTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// ems_action_parameter
	// EmsActionParameter
	// type
	// Type
	// string
	// END DEBUGGING
	// EmsActionParameterTypeString captures enum value "string"
	EmsActionParameterTypeString string = "string"

	// BEGIN DEBUGGING
	// ems_action_parameter
	// EmsActionParameter
	// type
	// Type
	// number
	// END DEBUGGING
	// EmsActionParameterTypeNumber captures enum value "number"
	EmsActionParameterTypeNumber string = "number"

	// BEGIN DEBUGGING
	// ems_action_parameter
	// EmsActionParameter
	// type
	// Type
	// integer
	// END DEBUGGING
	// EmsActionParameterTypeInteger captures enum value "integer"
	EmsActionParameterTypeInteger string = "integer"

	// BEGIN DEBUGGING
	// ems_action_parameter
	// EmsActionParameter
	// type
	// Type
	// boolean
	// END DEBUGGING
	// EmsActionParameterTypeBoolean captures enum value "boolean"
	EmsActionParameterTypeBoolean string = "boolean"

	// BEGIN DEBUGGING
	// ems_action_parameter
	// EmsActionParameter
	// type
	// Type
	// array
	// END DEBUGGING
	// EmsActionParameterTypeArray captures enum value "array"
	EmsActionParameterTypeArray string = "array"
)

// prop value enum
func (m *EmsActionParameter) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, emsActionParameterTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EmsActionParameter) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *EmsActionParameter) validateValidationErrorMessage(formats strfmt.Registry) error {
	if swag.IsZero(m.ValidationErrorMessage) { // not required
		return nil
	}

	if m.ValidationErrorMessage != nil {
		if err := m.ValidationErrorMessage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("validation_error_message")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ems action parameter based on the context it is used
func (m *EmsActionParameter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmsActionParameterInlineEnum(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExclusiveMaximum(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExclusiveMinimum(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFormat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHelp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxLength(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaximum(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMinItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMinLength(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMinimum(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParamIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTitle(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValidationErrorMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsActionParameter) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if m.Description != nil {
		if err := m.Description.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("description")
			}
			return err
		}
	}

	return nil
}

func (m *EmsActionParameter) contextValidateEmsActionParameterInlineEnum(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "enum", "body", []*string(m.EmsActionParameterInlineEnum)); err != nil {
		return err
	}

	return nil
}

func (m *EmsActionParameter) contextValidateExclusiveMaximum(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "exclusiveMaximum", "body", m.ExclusiveMaximum); err != nil {
		return err
	}

	return nil
}

func (m *EmsActionParameter) contextValidateExclusiveMinimum(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "exclusiveMinimum", "body", m.ExclusiveMinimum); err != nil {
		return err
	}

	return nil
}

func (m *EmsActionParameter) contextValidateFormat(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "format", "body", m.Format); err != nil {
		return err
	}

	return nil
}

func (m *EmsActionParameter) contextValidateHelp(ctx context.Context, formats strfmt.Registry) error {

	if m.Help != nil {
		if err := m.Help.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("help")
			}
			return err
		}
	}

	return nil
}

func (m *EmsActionParameter) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "items", "body", m.Items); err != nil {
		return err
	}

	return nil
}

func (m *EmsActionParameter) contextValidateMaxItems(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "maxItems", "body", m.MaxItems); err != nil {
		return err
	}

	return nil
}

func (m *EmsActionParameter) contextValidateMaxLength(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "maxLength", "body", m.MaxLength); err != nil {
		return err
	}

	return nil
}

func (m *EmsActionParameter) contextValidateMaximum(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "maximum", "body", m.Maximum); err != nil {
		return err
	}

	return nil
}

func (m *EmsActionParameter) contextValidateMinItems(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "minItems", "body", m.MinItems); err != nil {
		return err
	}

	return nil
}

func (m *EmsActionParameter) contextValidateMinLength(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "minLength", "body", m.MinLength); err != nil {
		return err
	}

	return nil
}

func (m *EmsActionParameter) contextValidateMinimum(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "minimum", "body", m.Minimum); err != nil {
		return err
	}

	return nil
}

func (m *EmsActionParameter) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *EmsActionParameter) contextValidateParamIn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "param_in", "body", m.ParamIn); err != nil {
		return err
	}

	return nil
}

func (m *EmsActionParameter) contextValidateTitle(ctx context.Context, formats strfmt.Registry) error {

	if m.Title != nil {
		if err := m.Title.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("title")
			}
			return err
		}
	}

	return nil
}

func (m *EmsActionParameter) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *EmsActionParameter) contextValidateValidationErrorMessage(ctx context.Context, formats strfmt.Registry) error {

	if m.ValidationErrorMessage != nil {
		if err := m.ValidationErrorMessage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("validation_error_message")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsActionParameter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsActionParameter) UnmarshalBinary(b []byte) error {
	var res EmsActionParameter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsActionParameterInlineDescription Description of the parameter that is presented in user facing applications.
//
// swagger:model ems_action_parameter_inline_description
type EmsActionParameterInlineDescription struct {

	// Message arguments
	// Read Only: true
	Arguments []*EmsActionParameterDescriptionArgumentsItems0 `json:"arguments,omitempty"`

	// Unique message code.
	// Example: 4
	// Read Only: true
	Code *string `json:"code,omitempty"`

	// User message.
	// Example: entry doesn't exist
	// Read Only: true
	Message *string `json:"message,omitempty"`
}

// Validate validates this ems action parameter inline description
func (m *EmsActionParameterInlineDescription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArguments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsActionParameterInlineDescription) validateArguments(formats strfmt.Registry) error {
	if swag.IsZero(m.Arguments) { // not required
		return nil
	}

	for i := 0; i < len(m.Arguments); i++ {
		if swag.IsZero(m.Arguments[i]) { // not required
			continue
		}

		if m.Arguments[i] != nil {
			if err := m.Arguments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("description" + "." + "arguments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ems action parameter inline description based on the context it is used
func (m *EmsActionParameterInlineDescription) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArguments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsActionParameterInlineDescription) contextValidateArguments(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "description"+"."+"arguments", "body", []*EmsActionParameterDescriptionArgumentsItems0(m.Arguments)); err != nil {
		return err
	}

	for i := 0; i < len(m.Arguments); i++ {

		if m.Arguments[i] != nil {
			if err := m.Arguments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("description" + "." + "arguments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EmsActionParameterInlineDescription) contextValidateCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "description"+"."+"code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *EmsActionParameterInlineDescription) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "description"+"."+"message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsActionParameterInlineDescription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsActionParameterInlineDescription) UnmarshalBinary(b []byte) error {
	var res EmsActionParameterInlineDescription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsActionParameterDescriptionArgumentsItems0 ems action parameter description arguments items0
//
// swagger:model EmsActionParameterDescriptionArgumentsItems0
type EmsActionParameterDescriptionArgumentsItems0 struct {

	// Argument code
	// Read Only: true
	Code *string `json:"code,omitempty"`

	// Message argument
	// Read Only: true
	Message *string `json:"message,omitempty"`
}

// Validate validates this ems action parameter description arguments items0
func (m *EmsActionParameterDescriptionArgumentsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this ems action parameter description arguments items0 based on the context it is used
func (m *EmsActionParameterDescriptionArgumentsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsActionParameterDescriptionArgumentsItems0) contextValidateCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *EmsActionParameterDescriptionArgumentsItems0) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsActionParameterDescriptionArgumentsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsActionParameterDescriptionArgumentsItems0) UnmarshalBinary(b []byte) error {
	var res EmsActionParameterDescriptionArgumentsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsActionParameterInlineHelp Help message of the parameter that is presented in user facing applications.
//
// swagger:model ems_action_parameter_inline_help
type EmsActionParameterInlineHelp struct {

	// Message arguments
	// Read Only: true
	Arguments []*EmsActionParameterHelpArgumentsItems0 `json:"arguments,omitempty"`

	// Unique message code.
	// Example: 4
	// Read Only: true
	Code *string `json:"code,omitempty"`

	// User message.
	// Example: entry doesn't exist
	// Read Only: true
	Message *string `json:"message,omitempty"`
}

// Validate validates this ems action parameter inline help
func (m *EmsActionParameterInlineHelp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArguments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsActionParameterInlineHelp) validateArguments(formats strfmt.Registry) error {
	if swag.IsZero(m.Arguments) { // not required
		return nil
	}

	for i := 0; i < len(m.Arguments); i++ {
		if swag.IsZero(m.Arguments[i]) { // not required
			continue
		}

		if m.Arguments[i] != nil {
			if err := m.Arguments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("help" + "." + "arguments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ems action parameter inline help based on the context it is used
func (m *EmsActionParameterInlineHelp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArguments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsActionParameterInlineHelp) contextValidateArguments(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "help"+"."+"arguments", "body", []*EmsActionParameterHelpArgumentsItems0(m.Arguments)); err != nil {
		return err
	}

	for i := 0; i < len(m.Arguments); i++ {

		if m.Arguments[i] != nil {
			if err := m.Arguments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("help" + "." + "arguments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EmsActionParameterInlineHelp) contextValidateCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "help"+"."+"code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *EmsActionParameterInlineHelp) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "help"+"."+"message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsActionParameterInlineHelp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsActionParameterInlineHelp) UnmarshalBinary(b []byte) error {
	var res EmsActionParameterInlineHelp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsActionParameterHelpArgumentsItems0 ems action parameter help arguments items0
//
// swagger:model EmsActionParameterHelpArgumentsItems0
type EmsActionParameterHelpArgumentsItems0 struct {

	// Argument code
	// Read Only: true
	Code *string `json:"code,omitempty"`

	// Message argument
	// Read Only: true
	Message *string `json:"message,omitempty"`
}

// Validate validates this ems action parameter help arguments items0
func (m *EmsActionParameterHelpArgumentsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this ems action parameter help arguments items0 based on the context it is used
func (m *EmsActionParameterHelpArgumentsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsActionParameterHelpArgumentsItems0) contextValidateCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *EmsActionParameterHelpArgumentsItems0) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsActionParameterHelpArgumentsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsActionParameterHelpArgumentsItems0) UnmarshalBinary(b []byte) error {
	var res EmsActionParameterHelpArgumentsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsActionParameterInlineTitle Title of the parameter that is presented in user facing applications.
//
// swagger:model ems_action_parameter_inline_title
type EmsActionParameterInlineTitle struct {

	// Message arguments
	// Read Only: true
	Arguments []*EmsActionParameterTitleArgumentsItems0 `json:"arguments,omitempty"`

	// Unique message code.
	// Example: 4
	// Read Only: true
	Code *string `json:"code,omitempty"`

	// User message.
	// Example: entry doesn't exist
	// Read Only: true
	Message *string `json:"message,omitempty"`
}

// Validate validates this ems action parameter inline title
func (m *EmsActionParameterInlineTitle) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArguments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsActionParameterInlineTitle) validateArguments(formats strfmt.Registry) error {
	if swag.IsZero(m.Arguments) { // not required
		return nil
	}

	for i := 0; i < len(m.Arguments); i++ {
		if swag.IsZero(m.Arguments[i]) { // not required
			continue
		}

		if m.Arguments[i] != nil {
			if err := m.Arguments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("title" + "." + "arguments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ems action parameter inline title based on the context it is used
func (m *EmsActionParameterInlineTitle) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArguments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsActionParameterInlineTitle) contextValidateArguments(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "title"+"."+"arguments", "body", []*EmsActionParameterTitleArgumentsItems0(m.Arguments)); err != nil {
		return err
	}

	for i := 0; i < len(m.Arguments); i++ {

		if m.Arguments[i] != nil {
			if err := m.Arguments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("title" + "." + "arguments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EmsActionParameterInlineTitle) contextValidateCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "title"+"."+"code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *EmsActionParameterInlineTitle) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "title"+"."+"message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsActionParameterInlineTitle) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsActionParameterInlineTitle) UnmarshalBinary(b []byte) error {
	var res EmsActionParameterInlineTitle
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsActionParameterTitleArgumentsItems0 ems action parameter title arguments items0
//
// swagger:model EmsActionParameterTitleArgumentsItems0
type EmsActionParameterTitleArgumentsItems0 struct {

	// Argument code
	// Read Only: true
	Code *string `json:"code,omitempty"`

	// Message argument
	// Read Only: true
	Message *string `json:"message,omitempty"`
}

// Validate validates this ems action parameter title arguments items0
func (m *EmsActionParameterTitleArgumentsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this ems action parameter title arguments items0 based on the context it is used
func (m *EmsActionParameterTitleArgumentsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsActionParameterTitleArgumentsItems0) contextValidateCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *EmsActionParameterTitleArgumentsItems0) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsActionParameterTitleArgumentsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsActionParameterTitleArgumentsItems0) UnmarshalBinary(b []byte) error {
	var res EmsActionParameterTitleArgumentsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsActionParameterInlineValidationErrorMessage Error message that is presented in user facing applications, in cases where parameter validation fails.
//
// swagger:model ems_action_parameter_inline_validation_error_message
type EmsActionParameterInlineValidationErrorMessage struct {

	// Message arguments
	// Read Only: true
	Arguments []*EmsActionParameterValidationErrorMessageArgumentsItems0 `json:"arguments,omitempty"`

	// Unique message code.
	// Example: 4
	// Read Only: true
	Code *string `json:"code,omitempty"`

	// User message.
	// Example: entry doesn't exist
	// Read Only: true
	Message *string `json:"message,omitempty"`
}

// Validate validates this ems action parameter inline validation error message
func (m *EmsActionParameterInlineValidationErrorMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArguments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsActionParameterInlineValidationErrorMessage) validateArguments(formats strfmt.Registry) error {
	if swag.IsZero(m.Arguments) { // not required
		return nil
	}

	for i := 0; i < len(m.Arguments); i++ {
		if swag.IsZero(m.Arguments[i]) { // not required
			continue
		}

		if m.Arguments[i] != nil {
			if err := m.Arguments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validation_error_message" + "." + "arguments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ems action parameter inline validation error message based on the context it is used
func (m *EmsActionParameterInlineValidationErrorMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArguments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsActionParameterInlineValidationErrorMessage) contextValidateArguments(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "validation_error_message"+"."+"arguments", "body", []*EmsActionParameterValidationErrorMessageArgumentsItems0(m.Arguments)); err != nil {
		return err
	}

	for i := 0; i < len(m.Arguments); i++ {

		if m.Arguments[i] != nil {
			if err := m.Arguments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validation_error_message" + "." + "arguments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EmsActionParameterInlineValidationErrorMessage) contextValidateCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "validation_error_message"+"."+"code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *EmsActionParameterInlineValidationErrorMessage) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "validation_error_message"+"."+"message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsActionParameterInlineValidationErrorMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsActionParameterInlineValidationErrorMessage) UnmarshalBinary(b []byte) error {
	var res EmsActionParameterInlineValidationErrorMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// EmsActionParameterValidationErrorMessageArgumentsItems0 ems action parameter validation error message arguments items0
//
// swagger:model EmsActionParameterValidationErrorMessageArgumentsItems0
type EmsActionParameterValidationErrorMessageArgumentsItems0 struct {

	// Argument code
	// Read Only: true
	Code *string `json:"code,omitempty"`

	// Message argument
	// Read Only: true
	Message *string `json:"message,omitempty"`
}

// Validate validates this ems action parameter validation error message arguments items0
func (m *EmsActionParameterValidationErrorMessageArgumentsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this ems action parameter validation error message arguments items0 based on the context it is used
func (m *EmsActionParameterValidationErrorMessageArgumentsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmsActionParameterValidationErrorMessageArgumentsItems0) contextValidateCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *EmsActionParameterValidationErrorMessageArgumentsItems0) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmsActionParameterValidationErrorMessageArgumentsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmsActionParameterValidationErrorMessageArgumentsItems0) UnmarshalBinary(b []byte) error {
	var res EmsActionParameterValidationErrorMessageArgumentsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
