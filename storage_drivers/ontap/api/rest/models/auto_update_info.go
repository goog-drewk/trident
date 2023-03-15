// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AutoUpdateInfo auto update info
//
// swagger:model auto_update_info
type AutoUpdateInfo struct {

	// links
	Links *AutoUpdateInfoInlineLinks `json:"_links,omitempty"`

	// Flag indicating feature state.
	// Example: true
	Enabled *bool `json:"enabled,omitempty"`

	// eula
	Eula *AutoUpdateInfoInlineEula `json:"eula,omitempty"`
}

// Validate validates this auto update info
func (m *AutoUpdateInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEula(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoUpdateInfo) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *AutoUpdateInfo) validateEula(formats strfmt.Registry) error {
	if swag.IsZero(m.Eula) { // not required
		return nil
	}

	if m.Eula != nil {
		if err := m.Eula.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eula")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this auto update info based on the context it is used
func (m *AutoUpdateInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEula(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoUpdateInfo) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *AutoUpdateInfo) contextValidateEula(ctx context.Context, formats strfmt.Registry) error {

	if m.Eula != nil {
		if err := m.Eula.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eula")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AutoUpdateInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutoUpdateInfo) UnmarshalBinary(b []byte) error {
	var res AutoUpdateInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AutoUpdateInfoInlineEula auto update info inline eula
//
// swagger:model auto_update_info_inline_eula
type AutoUpdateInfoInlineEula struct {

	// Flag indicating the End User License Agreement (EULA) acceptance. When the feature is enabled, it is assumed that the EULA is accepted.
	// Example: true
	// Read Only: true
	Accepted *bool `json:"accepted,omitempty"`

	// IP Address from where the EULA was accepted.
	// Example: 192.168.1.125
	// Read Only: true
	AcceptedIPAddress *string `json:"accepted_ip_address,omitempty"`

	// Date and time when the EULA was accepted.
	// Example: 2020-12-01T09:12:23-04:00
	// Read Only: true
	// Format: date-time
	AcceptedTimestamp *strfmt.DateTime `json:"accepted_timestamp,omitempty"`

	// User ID that provided the EULA acceptance.
	// Example: admin
	// Read Only: true
	UserIDAccepted *string `json:"user_id_accepted,omitempty"`
}

// Validate validates this auto update info inline eula
func (m *AutoUpdateInfoInlineEula) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptedTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoUpdateInfoInlineEula) validateAcceptedTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.AcceptedTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("eula"+"."+"accepted_timestamp", "body", "date-time", m.AcceptedTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this auto update info inline eula based on the context it is used
func (m *AutoUpdateInfoInlineEula) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccepted(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAcceptedIPAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAcceptedTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserIDAccepted(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoUpdateInfoInlineEula) contextValidateAccepted(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "eula"+"."+"accepted", "body", m.Accepted); err != nil {
		return err
	}

	return nil
}

func (m *AutoUpdateInfoInlineEula) contextValidateAcceptedIPAddress(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "eula"+"."+"accepted_ip_address", "body", m.AcceptedIPAddress); err != nil {
		return err
	}

	return nil
}

func (m *AutoUpdateInfoInlineEula) contextValidateAcceptedTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "eula"+"."+"accepted_timestamp", "body", m.AcceptedTimestamp); err != nil {
		return err
	}

	return nil
}

func (m *AutoUpdateInfoInlineEula) contextValidateUserIDAccepted(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "eula"+"."+"user_id_accepted", "body", m.UserIDAccepted); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AutoUpdateInfoInlineEula) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutoUpdateInfoInlineEula) UnmarshalBinary(b []byte) error {
	var res AutoUpdateInfoInlineEula
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AutoUpdateInfoInlineLinks auto update info inline links
//
// swagger:model auto_update_info_inline__links
type AutoUpdateInfoInlineLinks struct {

	// self
	Self *AutoUpdateInfoInlineLinksInlineSelf `json:"self,omitempty"`
}

// Validate validates this auto update info inline links
func (m *AutoUpdateInfoInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoUpdateInfoInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this auto update info inline links based on the context it is used
func (m *AutoUpdateInfoInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoUpdateInfoInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AutoUpdateInfoInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutoUpdateInfoInlineLinks) UnmarshalBinary(b []byte) error {
	var res AutoUpdateInfoInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AutoUpdateInfoInlineLinksInlineSelf auto update info inline links inline self
//
// swagger:model auto_update_info_inline__links_inline_self
type AutoUpdateInfoInlineLinksInlineSelf struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this auto update info inline links inline self
func (m *AutoUpdateInfoInlineLinksInlineSelf) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoUpdateInfoInlineLinksInlineSelf) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this auto update info inline links inline self based on the context it is used
func (m *AutoUpdateInfoInlineLinksInlineSelf) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoUpdateInfoInlineLinksInlineSelf) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AutoUpdateInfoInlineLinksInlineSelf) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutoUpdateInfoInlineLinksInlineSelf) UnmarshalBinary(b []byte) error {
	var res AutoUpdateInfoInlineLinksInlineSelf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
