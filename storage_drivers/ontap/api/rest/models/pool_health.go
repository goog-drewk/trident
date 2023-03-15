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

// PoolHealth pool health
//
// swagger:model pool_health
type PoolHealth struct {

	// Indicates whether the storage pool is able to participate in provisioning operations.
	// Read Only: true
	IsHealthy *bool `json:"is_healthy,omitempty"`

	// The state of the shared storage pool.
	// Read Only: true
	// Enum: [normal degraded creating deleting reassigning growing]
	State *string `json:"state,omitempty"`

	// Indicates why the storage pool is unhealthy. This property is not returned for healthy storage pools.
	// Read Only: true
	UnhealthyReason *Error `json:"unhealthy_reason,omitempty"`
}

// Validate validates this pool health
func (m *PoolHealth) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnhealthyReason(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var poolHealthTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["normal","degraded","creating","deleting","reassigning","growing"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		poolHealthTypeStatePropEnum = append(poolHealthTypeStatePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// pool_health
	// PoolHealth
	// state
	// State
	// normal
	// END DEBUGGING
	// PoolHealthStateNormal captures enum value "normal"
	PoolHealthStateNormal string = "normal"

	// BEGIN DEBUGGING
	// pool_health
	// PoolHealth
	// state
	// State
	// degraded
	// END DEBUGGING
	// PoolHealthStateDegraded captures enum value "degraded"
	PoolHealthStateDegraded string = "degraded"

	// BEGIN DEBUGGING
	// pool_health
	// PoolHealth
	// state
	// State
	// creating
	// END DEBUGGING
	// PoolHealthStateCreating captures enum value "creating"
	PoolHealthStateCreating string = "creating"

	// BEGIN DEBUGGING
	// pool_health
	// PoolHealth
	// state
	// State
	// deleting
	// END DEBUGGING
	// PoolHealthStateDeleting captures enum value "deleting"
	PoolHealthStateDeleting string = "deleting"

	// BEGIN DEBUGGING
	// pool_health
	// PoolHealth
	// state
	// State
	// reassigning
	// END DEBUGGING
	// PoolHealthStateReassigning captures enum value "reassigning"
	PoolHealthStateReassigning string = "reassigning"

	// BEGIN DEBUGGING
	// pool_health
	// PoolHealth
	// state
	// State
	// growing
	// END DEBUGGING
	// PoolHealthStateGrowing captures enum value "growing"
	PoolHealthStateGrowing string = "growing"
)

// prop value enum
func (m *PoolHealth) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, poolHealthTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PoolHealth) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

func (m *PoolHealth) validateUnhealthyReason(formats strfmt.Registry) error {
	if swag.IsZero(m.UnhealthyReason) { // not required
		return nil
	}

	if m.UnhealthyReason != nil {
		if err := m.UnhealthyReason.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unhealthy_reason")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this pool health based on the context it is used
func (m *PoolHealth) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIsHealthy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUnhealthyReason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PoolHealth) contextValidateIsHealthy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "is_healthy", "body", m.IsHealthy); err != nil {
		return err
	}

	return nil
}

func (m *PoolHealth) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *PoolHealth) contextValidateUnhealthyReason(ctx context.Context, formats strfmt.Registry) error {

	if m.UnhealthyReason != nil {
		if err := m.UnhealthyReason.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unhealthy_reason")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PoolHealth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PoolHealth) UnmarshalBinary(b []byte) error {
	var res PoolHealth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
