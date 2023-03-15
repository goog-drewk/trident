// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SoftwareStatusDetailsReference software status details reference
//
// swagger:model software_status_details_reference
type SoftwareStatusDetailsReference struct {

	// action
	Action *SoftwareStatusDetailsReferenceInlineAction `json:"action,omitempty"`

	// End time for each status phase.
	// Example: 2019-02-02T19:00:00Z
	// Read Only: true
	// Format: date-time
	EndTime *strfmt.DateTime `json:"end_time,omitempty"`

	// issue
	Issue *SoftwareStatusDetailsReferenceInlineIssue `json:"issue,omitempty"`

	// Name of the phase to be retrieved for status details.
	// Example: initialize
	// Read Only: true
	Name *string `json:"name,omitempty"`

	// node
	Node *SoftwareStatusDetailsReferenceInlineNode `json:"node,omitempty"`

	// Start time for each status phase.
	// Example: 2019-02-02T19:00:00Z
	// Read Only: true
	// Format: date-time
	StartTime *strfmt.DateTime `json:"start_time,omitempty"`

	// Status of the phase
	// Example: failed
	// Read Only: true
	// Enum: [in_progress waiting paused_by_user paused_on_error completed canceled failed pause_pending cancel_pending]
	State *string `json:"state,omitempty"`
}

// Validate validates this software status details reference
func (m *SoftwareStatusDetailsReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoftwareStatusDetailsReference) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if m.Action != nil {
		if err := m.Action.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action")
			}
			return err
		}
	}

	return nil
}

func (m *SoftwareStatusDetailsReference) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("end_time", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareStatusDetailsReference) validateIssue(formats strfmt.Registry) error {
	if swag.IsZero(m.Issue) { // not required
		return nil
	}

	if m.Issue != nil {
		if err := m.Issue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("issue")
			}
			return err
		}
	}

	return nil
}

func (m *SoftwareStatusDetailsReference) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

func (m *SoftwareStatusDetailsReference) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("start_time", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var softwareStatusDetailsReferenceTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["in_progress","waiting","paused_by_user","paused_on_error","completed","canceled","failed","pause_pending","cancel_pending"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		softwareStatusDetailsReferenceTypeStatePropEnum = append(softwareStatusDetailsReferenceTypeStatePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// software_status_details_reference
	// SoftwareStatusDetailsReference
	// state
	// State
	// in_progress
	// END DEBUGGING
	// SoftwareStatusDetailsReferenceStateInProgress captures enum value "in_progress"
	SoftwareStatusDetailsReferenceStateInProgress string = "in_progress"

	// BEGIN DEBUGGING
	// software_status_details_reference
	// SoftwareStatusDetailsReference
	// state
	// State
	// waiting
	// END DEBUGGING
	// SoftwareStatusDetailsReferenceStateWaiting captures enum value "waiting"
	SoftwareStatusDetailsReferenceStateWaiting string = "waiting"

	// BEGIN DEBUGGING
	// software_status_details_reference
	// SoftwareStatusDetailsReference
	// state
	// State
	// paused_by_user
	// END DEBUGGING
	// SoftwareStatusDetailsReferenceStatePausedByUser captures enum value "paused_by_user"
	SoftwareStatusDetailsReferenceStatePausedByUser string = "paused_by_user"

	// BEGIN DEBUGGING
	// software_status_details_reference
	// SoftwareStatusDetailsReference
	// state
	// State
	// paused_on_error
	// END DEBUGGING
	// SoftwareStatusDetailsReferenceStatePausedOnError captures enum value "paused_on_error"
	SoftwareStatusDetailsReferenceStatePausedOnError string = "paused_on_error"

	// BEGIN DEBUGGING
	// software_status_details_reference
	// SoftwareStatusDetailsReference
	// state
	// State
	// completed
	// END DEBUGGING
	// SoftwareStatusDetailsReferenceStateCompleted captures enum value "completed"
	SoftwareStatusDetailsReferenceStateCompleted string = "completed"

	// BEGIN DEBUGGING
	// software_status_details_reference
	// SoftwareStatusDetailsReference
	// state
	// State
	// canceled
	// END DEBUGGING
	// SoftwareStatusDetailsReferenceStateCanceled captures enum value "canceled"
	SoftwareStatusDetailsReferenceStateCanceled string = "canceled"

	// BEGIN DEBUGGING
	// software_status_details_reference
	// SoftwareStatusDetailsReference
	// state
	// State
	// failed
	// END DEBUGGING
	// SoftwareStatusDetailsReferenceStateFailed captures enum value "failed"
	SoftwareStatusDetailsReferenceStateFailed string = "failed"

	// BEGIN DEBUGGING
	// software_status_details_reference
	// SoftwareStatusDetailsReference
	// state
	// State
	// pause_pending
	// END DEBUGGING
	// SoftwareStatusDetailsReferenceStatePausePending captures enum value "pause_pending"
	SoftwareStatusDetailsReferenceStatePausePending string = "pause_pending"

	// BEGIN DEBUGGING
	// software_status_details_reference
	// SoftwareStatusDetailsReference
	// state
	// State
	// cancel_pending
	// END DEBUGGING
	// SoftwareStatusDetailsReferenceStateCancelPending captures enum value "cancel_pending"
	SoftwareStatusDetailsReferenceStateCancelPending string = "cancel_pending"
)

// prop value enum
func (m *SoftwareStatusDetailsReference) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, softwareStatusDetailsReferenceTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SoftwareStatusDetailsReference) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this software status details reference based on the context it is used
func (m *SoftwareStatusDetailsReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEndTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIssue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStartTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoftwareStatusDetailsReference) contextValidateAction(ctx context.Context, formats strfmt.Registry) error {

	if m.Action != nil {
		if err := m.Action.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("action")
			}
			return err
		}
	}

	return nil
}

func (m *SoftwareStatusDetailsReference) contextValidateEndTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "end_time", "body", m.EndTime); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareStatusDetailsReference) contextValidateIssue(ctx context.Context, formats strfmt.Registry) error {

	if m.Issue != nil {
		if err := m.Issue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("issue")
			}
			return err
		}
	}

	return nil
}

func (m *SoftwareStatusDetailsReference) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareStatusDetailsReference) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if m.Node != nil {
		if err := m.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

func (m *SoftwareStatusDetailsReference) contextValidateStartTime(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "start_time", "body", m.StartTime); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareStatusDetailsReference) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SoftwareStatusDetailsReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareStatusDetailsReference) UnmarshalBinary(b []byte) error {
	var res SoftwareStatusDetailsReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SoftwareStatusDetailsReferenceInlineAction software status details reference inline action
//
// swagger:model software_status_details_reference_inline_action
type SoftwareStatusDetailsReferenceInlineAction struct {

	// Error code corresponding the status error
	Code *int64 `json:"code,omitempty"`

	// Corrective action to be taken to resolve the status error.
	Message *string `json:"message,omitempty"`
}

// Validate validates this software status details reference inline action
func (m *SoftwareStatusDetailsReferenceInlineAction) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this software status details reference inline action based on the context it is used
func (m *SoftwareStatusDetailsReferenceInlineAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SoftwareStatusDetailsReferenceInlineAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareStatusDetailsReferenceInlineAction) UnmarshalBinary(b []byte) error {
	var res SoftwareStatusDetailsReferenceInlineAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SoftwareStatusDetailsReferenceInlineIssue software status details reference inline issue
//
// swagger:model software_status_details_reference_inline_issue
type SoftwareStatusDetailsReferenceInlineIssue struct {

	// Error code corresponding to update status
	// Example: 10551399
	Code *int64 `json:"code,omitempty"`

	// Update status details
	// Example: Image update complete
	Message *string `json:"message,omitempty"`
}

// Validate validates this software status details reference inline issue
func (m *SoftwareStatusDetailsReferenceInlineIssue) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this software status details reference inline issue based on the context it is used
func (m *SoftwareStatusDetailsReferenceInlineIssue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SoftwareStatusDetailsReferenceInlineIssue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareStatusDetailsReferenceInlineIssue) UnmarshalBinary(b []byte) error {
	var res SoftwareStatusDetailsReferenceInlineIssue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SoftwareStatusDetailsReferenceInlineNode software status details reference inline node
//
// swagger:model software_status_details_reference_inline_node
type SoftwareStatusDetailsReferenceInlineNode struct {

	// Name of the node to be retrieved for status details.
	// Example: node1
	// Read Only: true
	Name *string `json:"name,omitempty"`
}

// Validate validates this software status details reference inline node
func (m *SoftwareStatusDetailsReferenceInlineNode) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this software status details reference inline node based on the context it is used
func (m *SoftwareStatusDetailsReferenceInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoftwareStatusDetailsReferenceInlineNode) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "node"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SoftwareStatusDetailsReferenceInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareStatusDetailsReferenceInlineNode) UnmarshalBinary(b []byte) error {
	var res SoftwareStatusDetailsReferenceInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
