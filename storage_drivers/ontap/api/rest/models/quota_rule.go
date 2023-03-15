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

// QuotaRule quota rule
// Example: {"files":{"hard_limit":"100","soft_limit":"80"},"qtree":{"id":"1","name":"qt1"},"space":{"hard_limit":"1222800","soft_limit":"51200"},"svm":{"name":"svm1"},"type":"user","user_mapping":"on","users":[{"name":"fred"}],"uuid":"264a9e0b-2e03-11e9-a610-005056a7b72d","volume":{"name":"fv"}}
//
// swagger:model quota_rule
type QuotaRule struct {

	// links
	Links *QuotaRuleInlineLinks `json:"_links,omitempty"`

	// files
	Files *QuotaRuleInlineFiles `json:"files,omitempty"`

	// group
	Group *QuotaRuleInlineGroup `json:"group,omitempty"`

	// qtree
	Qtree *QuotaRuleInlineQtree `json:"qtree,omitempty"`

	// This parameter specifies the target user to which the user quota policy rule applies. This parameter takes single or multiple user names or identifiers. This parameter is valid only for the POST operation of a user quota policy rule. If this parameter is used as an input to create a group or a tree quota policy rule, the POST operation will fail with an appropriate error. For POST, this input parameter takes either a user name or a user identifier, not both. For default quota rules, the user name must be chosen and specified as "". For explicit user quota rules, this parameter can indicate either a user name or user identifier. The user name can be a UNIX user name or a Windows user name. If a name contains a space, enclose the entire value in quotes. A UNIX user name cannot include a backslash (\) or an @ sign; user names with these characters are treated as Windows names. The user identifer can be a UNIX user identifier or a Windows security identifier. For multi-user quota, this parameter can contain multiple user targets separated by a comma.
	QuotaRuleInlineUsers []*QuotaRuleInlineUsersInlineArrayItem `json:"users,omitempty"`

	// space
	Space *QuotaRuleInlineSpace `json:"space,omitempty"`

	// svm
	Svm *QuotaRuleInlineSvm `json:"svm,omitempty"`

	// This parameter specifies the quota policy rule type. This is required in POST only and can take either one of the \"user\", \"group\" or \"tree\" values.
	// Enum: [tree user group]
	Type *string `json:"type,omitempty"`

	// This parameter enables user mapping for user quota policy rules. This is valid in POST or PATCH for user quota policy rules only.
	UserMapping *bool `json:"user_mapping,omitempty"`

	// Unique identifier for the quota policy rule. This field is generated when the quota policy rule is created.
	// Example: 5f1d13a7-f401-11e8-ac1a-005056a7c3b9
	// Read Only: true
	UUID *string `json:"uuid,omitempty"`

	// volume
	Volume *QuotaRuleInlineVolume `json:"volume,omitempty"`
}

// Validate validates this quota rule
func (m *QuotaRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQtree(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuotaRuleInlineUsers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuotaRule) validateLinks(formats strfmt.Registry) error {
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

func (m *QuotaRule) validateFiles(formats strfmt.Registry) error {
	if swag.IsZero(m.Files) { // not required
		return nil
	}

	if m.Files != nil {
		if err := m.Files.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("files")
			}
			return err
		}
	}

	return nil
}

func (m *QuotaRule) validateGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if m.Group != nil {
		if err := m.Group.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

func (m *QuotaRule) validateQtree(formats strfmt.Registry) error {
	if swag.IsZero(m.Qtree) { // not required
		return nil
	}

	if m.Qtree != nil {
		if err := m.Qtree.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qtree")
			}
			return err
		}
	}

	return nil
}

func (m *QuotaRule) validateQuotaRuleInlineUsers(formats strfmt.Registry) error {
	if swag.IsZero(m.QuotaRuleInlineUsers) { // not required
		return nil
	}

	for i := 0; i < len(m.QuotaRuleInlineUsers); i++ {
		if swag.IsZero(m.QuotaRuleInlineUsers[i]) { // not required
			continue
		}

		if m.QuotaRuleInlineUsers[i] != nil {
			if err := m.QuotaRuleInlineUsers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QuotaRule) validateSpace(formats strfmt.Registry) error {
	if swag.IsZero(m.Space) { // not required
		return nil
	}

	if m.Space != nil {
		if err := m.Space.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("space")
			}
			return err
		}
	}

	return nil
}

func (m *QuotaRule) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

var quotaRuleTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tree","user","group"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		quotaRuleTypeTypePropEnum = append(quotaRuleTypeTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// quota_rule
	// QuotaRule
	// type
	// Type
	// tree
	// END DEBUGGING
	// QuotaRuleTypeTree captures enum value "tree"
	QuotaRuleTypeTree string = "tree"

	// BEGIN DEBUGGING
	// quota_rule
	// QuotaRule
	// type
	// Type
	// user
	// END DEBUGGING
	// QuotaRuleTypeUser captures enum value "user"
	QuotaRuleTypeUser string = "user"

	// BEGIN DEBUGGING
	// quota_rule
	// QuotaRule
	// type
	// Type
	// group
	// END DEBUGGING
	// QuotaRuleTypeGroup captures enum value "group"
	QuotaRuleTypeGroup string = "group"
)

// prop value enum
func (m *QuotaRule) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, quotaRuleTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *QuotaRule) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *QuotaRule) validateVolume(formats strfmt.Registry) error {
	if swag.IsZero(m.Volume) { // not required
		return nil
	}

	if m.Volume != nil {
		if err := m.Volume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this quota rule based on the context it is used
func (m *QuotaRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQtree(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQuotaRuleInlineUsers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuotaRule) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *QuotaRule) contextValidateFiles(ctx context.Context, formats strfmt.Registry) error {

	if m.Files != nil {
		if err := m.Files.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("files")
			}
			return err
		}
	}

	return nil
}

func (m *QuotaRule) contextValidateGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.Group != nil {
		if err := m.Group.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

func (m *QuotaRule) contextValidateQtree(ctx context.Context, formats strfmt.Registry) error {

	if m.Qtree != nil {
		if err := m.Qtree.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qtree")
			}
			return err
		}
	}

	return nil
}

func (m *QuotaRule) contextValidateQuotaRuleInlineUsers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.QuotaRuleInlineUsers); i++ {

		if m.QuotaRuleInlineUsers[i] != nil {
			if err := m.QuotaRuleInlineUsers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QuotaRule) contextValidateSpace(ctx context.Context, formats strfmt.Registry) error {

	if m.Space != nil {
		if err := m.Space.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("space")
			}
			return err
		}
	}

	return nil
}

func (m *QuotaRule) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {
		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

func (m *QuotaRule) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

func (m *QuotaRule) contextValidateVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.Volume != nil {
		if err := m.Volume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QuotaRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuotaRule) UnmarshalBinary(b []byte) error {
	var res QuotaRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QuotaRuleInlineFiles quota rule inline files
//
// swagger:model quota_rule_inline_files
type QuotaRuleInlineFiles struct {

	// This parameter specifies the hard limit for files. This is valid in POST or PATCH.
	HardLimit *int64 `json:"hard_limit,omitempty"`

	// This parameter specifies the soft limit for files. This is valid in POST or PATCH.
	SoftLimit *int64 `json:"soft_limit,omitempty"`
}

// Validate validates this quota rule inline files
func (m *QuotaRuleInlineFiles) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this quota rule inline files based on context it is used
func (m *QuotaRuleInlineFiles) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QuotaRuleInlineFiles) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuotaRuleInlineFiles) UnmarshalBinary(b []byte) error {
	var res QuotaRuleInlineFiles
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QuotaRuleInlineGroup This parameter specifies the target group to which the group quota policy rule applies. This parameter takes a group name or identifier. This parameter is only valid for the POST operation of a group quota policy rule. The POST operation will fail with an appropriate error if this parameter is used as an input to create a user or a tree quota policy rule. This input parameter for POST takes either a group name or a group identifier, but not both. For default quota rules, the group name must be chosen and should be specified as "". For explicit group quota rules, this parameter can contain a UNIX group name or a UNIX group identifier.
//
// swagger:model quota_rule_inline_group
type QuotaRuleInlineGroup struct {

	// Quota target group ID
	ID *string `json:"id,omitempty"`

	// Quota target group name
	Name *string `json:"name,omitempty"`
}

// Validate validates this quota rule inline group
func (m *QuotaRuleInlineGroup) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this quota rule inline group based on context it is used
func (m *QuotaRuleInlineGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QuotaRuleInlineGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuotaRuleInlineGroup) UnmarshalBinary(b []byte) error {
	var res QuotaRuleInlineGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QuotaRuleInlineLinks quota rule inline links
//
// swagger:model quota_rule_inline__links
type QuotaRuleInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this quota rule inline links
func (m *QuotaRuleInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuotaRuleInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this quota rule inline links based on the context it is used
func (m *QuotaRuleInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuotaRuleInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *QuotaRuleInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuotaRuleInlineLinks) UnmarshalBinary(b []byte) error {
	var res QuotaRuleInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QuotaRuleInlineQtree This parameter specifies the target qtree to which the user/group/tree quota policy rule applies. For a user/group quota policy rule at qtree level, this parameter takes a qtree name and is valid in GET or POST. For a user/group quota policy rule at volume level, this parameter is not valid in GET or POST. For a tree quota policy rule, this parameter is mandatory and is valid in both POST and GET. For a default tree quota policy rule, this parameter needs to be specified as "". For a tree quota policy rule at qtree level, this parameter takes a qtree name and is valid in GET or POST.
//
// swagger:model quota_rule_inline_qtree
type QuotaRuleInlineQtree struct {

	// links
	Links *QuotaRuleInlineQtreeInlineLinks `json:"_links,omitempty"`

	// The unique identifier for a qtree.
	// Example: 1
	// Read Only: true
	ID *int64 `json:"id,omitempty"`

	// The name of the qtree.
	// Example: qt1
	Name *string `json:"name,omitempty"`
}

// Validate validates this quota rule inline qtree
func (m *QuotaRuleInlineQtree) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuotaRuleInlineQtree) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qtree" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this quota rule inline qtree based on the context it is used
func (m *QuotaRuleInlineQtree) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuotaRuleInlineQtree) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qtree" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *QuotaRuleInlineQtree) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "qtree"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QuotaRuleInlineQtree) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuotaRuleInlineQtree) UnmarshalBinary(b []byte) error {
	var res QuotaRuleInlineQtree
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QuotaRuleInlineQtreeInlineLinks quota rule inline qtree inline links
//
// swagger:model quota_rule_inline_qtree_inline__links
type QuotaRuleInlineQtreeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this quota rule inline qtree inline links
func (m *QuotaRuleInlineQtreeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuotaRuleInlineQtreeInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qtree" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this quota rule inline qtree inline links based on the context it is used
func (m *QuotaRuleInlineQtreeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuotaRuleInlineQtreeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qtree" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QuotaRuleInlineQtreeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuotaRuleInlineQtreeInlineLinks) UnmarshalBinary(b []byte) error {
	var res QuotaRuleInlineQtreeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QuotaRuleInlineSpace quota rule inline space
//
// swagger:model quota_rule_inline_space
type QuotaRuleInlineSpace struct {

	// This parameter specifies the space hard limit, in bytes. If less than 1024 bytes, the value is rounded up to 1024 bytes. Valid in POST or PATCH. For a POST operation where the parameter is either empty or set to -1, no limit is applied. For a PATCH operation where a limit is configured, use a value of -1 to clear the limit.
	HardLimit *int64 `json:"hard_limit,omitempty"`

	// This parameter specifies the space soft limit, in bytes. If less than 1024 bytes, the value is rounded up to 1024 bytes. Valid in POST or PATCH. For a POST operation where the parameter is either empty or set to -1, no limit is applied. For a PATCH operation where a limit is configured, use a value of -1 to clear the limit.
	SoftLimit *int64 `json:"soft_limit,omitempty"`
}

// Validate validates this quota rule inline space
func (m *QuotaRuleInlineSpace) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this quota rule inline space based on context it is used
func (m *QuotaRuleInlineSpace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QuotaRuleInlineSpace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuotaRuleInlineSpace) UnmarshalBinary(b []byte) error {
	var res QuotaRuleInlineSpace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QuotaRuleInlineSvm quota rule inline svm
//
// swagger:model quota_rule_inline_svm
type QuotaRuleInlineSvm struct {

	// links
	Links *QuotaRuleInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this quota rule inline svm
func (m *QuotaRuleInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuotaRuleInlineSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this quota rule inline svm based on the context it is used
func (m *QuotaRuleInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuotaRuleInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QuotaRuleInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuotaRuleInlineSvm) UnmarshalBinary(b []byte) error {
	var res QuotaRuleInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QuotaRuleInlineSvmInlineLinks quota rule inline svm inline links
//
// swagger:model quota_rule_inline_svm_inline__links
type QuotaRuleInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this quota rule inline svm inline links
func (m *QuotaRuleInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuotaRuleInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this quota rule inline svm inline links based on the context it is used
func (m *QuotaRuleInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuotaRuleInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QuotaRuleInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuotaRuleInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res QuotaRuleInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QuotaRuleInlineUsersInlineArrayItem quota rule inline users inline array item
//
// swagger:model quota_rule_inline_users_inline_array_item
type QuotaRuleInlineUsersInlineArrayItem struct {

	// Quota target user ID
	ID *string `json:"id,omitempty"`

	// Quota target user name
	Name *string `json:"name,omitempty"`
}

// Validate validates this quota rule inline users inline array item
func (m *QuotaRuleInlineUsersInlineArrayItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this quota rule inline users inline array item based on context it is used
func (m *QuotaRuleInlineUsersInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QuotaRuleInlineUsersInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuotaRuleInlineUsersInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res QuotaRuleInlineUsersInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QuotaRuleInlineVolume quota rule inline volume
//
// swagger:model quota_rule_inline_volume
type QuotaRuleInlineVolume struct {

	// links
	Links *QuotaRuleInlineVolumeInlineLinks `json:"_links,omitempty"`

	// The name of the volume.
	// Example: volume1
	Name *string `json:"name,omitempty"`

	// Unique identifier for the volume. This corresponds to the instance-uuid that is exposed in the CLI and ONTAPI. It does not change due to a volume move.
	// Example: 028baa66-41bd-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this quota rule inline volume
func (m *QuotaRuleInlineVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuotaRuleInlineVolume) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this quota rule inline volume based on the context it is used
func (m *QuotaRuleInlineVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuotaRuleInlineVolume) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QuotaRuleInlineVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuotaRuleInlineVolume) UnmarshalBinary(b []byte) error {
	var res QuotaRuleInlineVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// QuotaRuleInlineVolumeInlineLinks quota rule inline volume inline links
//
// swagger:model quota_rule_inline_volume_inline__links
type QuotaRuleInlineVolumeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this quota rule inline volume inline links
func (m *QuotaRuleInlineVolumeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuotaRuleInlineVolumeInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this quota rule inline volume inline links based on the context it is used
func (m *QuotaRuleInlineVolumeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuotaRuleInlineVolumeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("volume" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QuotaRuleInlineVolumeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuotaRuleInlineVolumeInlineLinks) UnmarshalBinary(b []byte) error {
	var res QuotaRuleInlineVolumeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
