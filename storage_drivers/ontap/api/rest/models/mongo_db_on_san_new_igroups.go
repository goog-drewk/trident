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

// MongoDbOnSanNewIgroups The list of initiator groups to create.
//
// swagger:model mongo_db_on_san_new_igroups
type MongoDbOnSanNewIgroups struct {

	// A comment available for use by the administrator.
	Comment *string `json:"comment,omitempty"`

	// mongo db on san new igroups inline igroups
	MongoDbOnSanNewIgroupsInlineIgroups []*MongoDbOnSanNewIgroupsInlineIgroupsInlineArrayItem `json:"igroups,omitempty"`

	// mongo db on san new igroups inline initiator objects
	MongoDbOnSanNewIgroupsInlineInitiatorObjects []*MongoDbOnSanNewIgroupsInlineInitiatorObjectsInlineArrayItem `json:"initiator_objects,omitempty"`

	// mongo db on san new igroups inline initiators
	MongoDbOnSanNewIgroupsInlineInitiators []*string `json:"initiators,omitempty"`

	// The name of the new initiator group.
	// Required: true
	// Max Length: 96
	// Min Length: 1
	Name *string `json:"name"`

	// The name of the host OS accessing the application. The default value is the host OS that is running the application.
	// Enum: [hyper_v linux solaris vmware windows xen]
	OsType *string `json:"os_type,omitempty"`

	// The protocol of the new initiator group.
	// Enum: [fcp iscsi mixed]
	Protocol *string `json:"protocol,omitempty"`
}

// Validate validates this mongo db on san new igroups
func (m *MongoDbOnSanNewIgroups) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMongoDbOnSanNewIgroupsInlineIgroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMongoDbOnSanNewIgroupsInlineInitiatorObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MongoDbOnSanNewIgroups) validateMongoDbOnSanNewIgroupsInlineIgroups(formats strfmt.Registry) error {
	if swag.IsZero(m.MongoDbOnSanNewIgroupsInlineIgroups) { // not required
		return nil
	}

	for i := 0; i < len(m.MongoDbOnSanNewIgroupsInlineIgroups); i++ {
		if swag.IsZero(m.MongoDbOnSanNewIgroupsInlineIgroups[i]) { // not required
			continue
		}

		if m.MongoDbOnSanNewIgroupsInlineIgroups[i] != nil {
			if err := m.MongoDbOnSanNewIgroupsInlineIgroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("igroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MongoDbOnSanNewIgroups) validateMongoDbOnSanNewIgroupsInlineInitiatorObjects(formats strfmt.Registry) error {
	if swag.IsZero(m.MongoDbOnSanNewIgroupsInlineInitiatorObjects) { // not required
		return nil
	}

	for i := 0; i < len(m.MongoDbOnSanNewIgroupsInlineInitiatorObjects); i++ {
		if swag.IsZero(m.MongoDbOnSanNewIgroupsInlineInitiatorObjects[i]) { // not required
			continue
		}

		if m.MongoDbOnSanNewIgroupsInlineInitiatorObjects[i] != nil {
			if err := m.MongoDbOnSanNewIgroupsInlineInitiatorObjects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("initiator_objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MongoDbOnSanNewIgroups) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 96); err != nil {
		return err
	}

	return nil
}

var mongoDbOnSanNewIgroupsTypeOsTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["hyper_v","linux","solaris","vmware","windows","xen"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mongoDbOnSanNewIgroupsTypeOsTypePropEnum = append(mongoDbOnSanNewIgroupsTypeOsTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// mongo_db_on_san_new_igroups
	// MongoDbOnSanNewIgroups
	// os_type
	// OsType
	// hyper_v
	// END DEBUGGING
	// MongoDbOnSanNewIgroupsOsTypeHyperv captures enum value "hyper_v"
	MongoDbOnSanNewIgroupsOsTypeHyperv string = "hyper_v"

	// BEGIN DEBUGGING
	// mongo_db_on_san_new_igroups
	// MongoDbOnSanNewIgroups
	// os_type
	// OsType
	// linux
	// END DEBUGGING
	// MongoDbOnSanNewIgroupsOsTypeLinux captures enum value "linux"
	MongoDbOnSanNewIgroupsOsTypeLinux string = "linux"

	// BEGIN DEBUGGING
	// mongo_db_on_san_new_igroups
	// MongoDbOnSanNewIgroups
	// os_type
	// OsType
	// solaris
	// END DEBUGGING
	// MongoDbOnSanNewIgroupsOsTypeSolaris captures enum value "solaris"
	MongoDbOnSanNewIgroupsOsTypeSolaris string = "solaris"

	// BEGIN DEBUGGING
	// mongo_db_on_san_new_igroups
	// MongoDbOnSanNewIgroups
	// os_type
	// OsType
	// vmware
	// END DEBUGGING
	// MongoDbOnSanNewIgroupsOsTypeVmware captures enum value "vmware"
	MongoDbOnSanNewIgroupsOsTypeVmware string = "vmware"

	// BEGIN DEBUGGING
	// mongo_db_on_san_new_igroups
	// MongoDbOnSanNewIgroups
	// os_type
	// OsType
	// windows
	// END DEBUGGING
	// MongoDbOnSanNewIgroupsOsTypeWindows captures enum value "windows"
	MongoDbOnSanNewIgroupsOsTypeWindows string = "windows"

	// BEGIN DEBUGGING
	// mongo_db_on_san_new_igroups
	// MongoDbOnSanNewIgroups
	// os_type
	// OsType
	// xen
	// END DEBUGGING
	// MongoDbOnSanNewIgroupsOsTypeXen captures enum value "xen"
	MongoDbOnSanNewIgroupsOsTypeXen string = "xen"
)

// prop value enum
func (m *MongoDbOnSanNewIgroups) validateOsTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, mongoDbOnSanNewIgroupsTypeOsTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MongoDbOnSanNewIgroups) validateOsType(formats strfmt.Registry) error {
	if swag.IsZero(m.OsType) { // not required
		return nil
	}

	// value enum
	if err := m.validateOsTypeEnum("os_type", "body", *m.OsType); err != nil {
		return err
	}

	return nil
}

var mongoDbOnSanNewIgroupsTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["fcp","iscsi","mixed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mongoDbOnSanNewIgroupsTypeProtocolPropEnum = append(mongoDbOnSanNewIgroupsTypeProtocolPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// mongo_db_on_san_new_igroups
	// MongoDbOnSanNewIgroups
	// protocol
	// Protocol
	// fcp
	// END DEBUGGING
	// MongoDbOnSanNewIgroupsProtocolFcp captures enum value "fcp"
	MongoDbOnSanNewIgroupsProtocolFcp string = "fcp"

	// BEGIN DEBUGGING
	// mongo_db_on_san_new_igroups
	// MongoDbOnSanNewIgroups
	// protocol
	// Protocol
	// iscsi
	// END DEBUGGING
	// MongoDbOnSanNewIgroupsProtocolIscsi captures enum value "iscsi"
	MongoDbOnSanNewIgroupsProtocolIscsi string = "iscsi"

	// BEGIN DEBUGGING
	// mongo_db_on_san_new_igroups
	// MongoDbOnSanNewIgroups
	// protocol
	// Protocol
	// mixed
	// END DEBUGGING
	// MongoDbOnSanNewIgroupsProtocolMixed captures enum value "mixed"
	MongoDbOnSanNewIgroupsProtocolMixed string = "mixed"
)

// prop value enum
func (m *MongoDbOnSanNewIgroups) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, mongoDbOnSanNewIgroupsTypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MongoDbOnSanNewIgroups) validateProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", *m.Protocol); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this mongo db on san new igroups based on the context it is used
func (m *MongoDbOnSanNewIgroups) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMongoDbOnSanNewIgroupsInlineIgroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMongoDbOnSanNewIgroupsInlineInitiatorObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MongoDbOnSanNewIgroups) contextValidateMongoDbOnSanNewIgroupsInlineIgroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MongoDbOnSanNewIgroupsInlineIgroups); i++ {

		if m.MongoDbOnSanNewIgroupsInlineIgroups[i] != nil {
			if err := m.MongoDbOnSanNewIgroupsInlineIgroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("igroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MongoDbOnSanNewIgroups) contextValidateMongoDbOnSanNewIgroupsInlineInitiatorObjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MongoDbOnSanNewIgroupsInlineInitiatorObjects); i++ {

		if m.MongoDbOnSanNewIgroupsInlineInitiatorObjects[i] != nil {
			if err := m.MongoDbOnSanNewIgroupsInlineInitiatorObjects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("initiator_objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MongoDbOnSanNewIgroups) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MongoDbOnSanNewIgroups) UnmarshalBinary(b []byte) error {
	var res MongoDbOnSanNewIgroups
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MongoDbOnSanNewIgroupsInlineIgroupsInlineArrayItem mongo db on san new igroups inline igroups inline array item
//
// swagger:model mongo_db_on_san_new_igroups_inline_igroups_inline_array_item
type MongoDbOnSanNewIgroupsInlineIgroupsInlineArrayItem struct {

	// The name of an igroup to nest within a parent igroup. Mutually exclusive with initiators and initiator_objects.
	Name *string `json:"name,omitempty"`

	// The UUID of an igroup to nest within a parent igroup Usage: &lt;UUID&gt;
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this mongo db on san new igroups inline igroups inline array item
func (m *MongoDbOnSanNewIgroupsInlineIgroupsInlineArrayItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this mongo db on san new igroups inline igroups inline array item based on context it is used
func (m *MongoDbOnSanNewIgroupsInlineIgroupsInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MongoDbOnSanNewIgroupsInlineIgroupsInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MongoDbOnSanNewIgroupsInlineIgroupsInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res MongoDbOnSanNewIgroupsInlineIgroupsInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MongoDbOnSanNewIgroupsInlineInitiatorObjectsInlineArrayItem mongo db on san new igroups inline initiator objects inline array item
//
// swagger:model mongo_db_on_san_new_igroups_inline_initiator_objects_inline_array_item
type MongoDbOnSanNewIgroupsInlineInitiatorObjectsInlineArrayItem struct {

	// A comment available for use by the administrator.
	Comment *string `json:"comment,omitempty"`

	// The WWPN, IQN, or Alias of the initiator. Mutually exclusive with nested igroups and the initiators array.
	Name *string `json:"name,omitempty"`
}

// Validate validates this mongo db on san new igroups inline initiator objects inline array item
func (m *MongoDbOnSanNewIgroupsInlineInitiatorObjectsInlineArrayItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this mongo db on san new igroups inline initiator objects inline array item based on context it is used
func (m *MongoDbOnSanNewIgroupsInlineInitiatorObjectsInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MongoDbOnSanNewIgroupsInlineInitiatorObjectsInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MongoDbOnSanNewIgroupsInlineInitiatorObjectsInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res MongoDbOnSanNewIgroupsInlineInitiatorObjectsInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
