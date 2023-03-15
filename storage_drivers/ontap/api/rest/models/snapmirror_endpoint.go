// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SnapmirrorEndpoint Endpoint of a SnapMirror relationship. For a GET request, the property "cluster" is populated when the endpoint is on a remote cluster. A POST request to create the destination SVM endpoint or to establish an SVM DR relationship must have the property "cluster" populated with the remote cluster details. A POST request to create the destination FlexVol volume, FlexGroup volume, Consistency Group, ONTAP S3 bucket and NON-ONTAP object-store endpoints can optionally specify the "cluster" property when the source SVM and the destination SVM are peered. A POST request to establish a SnapMirror relationship between the source endpoint and destination endpoint and when the source SVM and the destination SVM are not peered, must specify the "cluster" property for the remote endpoint.
//
// swagger:model snapmirror_endpoint
type SnapmirrorEndpoint struct {

	// cluster
	Cluster *SnapmirrorEndpointInlineCluster `json:"cluster,omitempty"`

	// Optional property to specify the IPSpace of the SVM.
	// Example: Default
	Ipspace *string `json:"ipspace,omitempty"`

	// ONTAP FlexVol/FlexGroup - svm1:volume1
	// ONTAP SVM               - svm1:
	// ONTAP Consistency Group - svm1:/cg/cg_name
	// ONTAP S3                - svm1:/bucket/bucket1
	// NON-ONTAP               - objstore1:/objstore
	//
	// Example: svm1:volume1
	Path *string `json:"path,omitempty"`

	// Mandatory property for a Consistency Group endpoint. Specifies the list of FlexVol volumes for a Consistency Group.
	SnapmirrorEndpointInlineConsistencyGroupVolumes []*SnapmirrorEndpointInlineConsistencyGroupVolumesInlineArrayItem `json:"consistency_group_volumes,omitempty"`

	// svm
	Svm *SnapmirrorEndpointInlineSvm `json:"svm,omitempty"`
}

// Validate validates this snapmirror endpoint
func (m *SnapmirrorEndpoint) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapmirrorEndpointInlineConsistencyGroupVolumes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorEndpoint) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *SnapmirrorEndpoint) validateSnapmirrorEndpointInlineConsistencyGroupVolumes(formats strfmt.Registry) error {
	if swag.IsZero(m.SnapmirrorEndpointInlineConsistencyGroupVolumes) { // not required
		return nil
	}

	for i := 0; i < len(m.SnapmirrorEndpointInlineConsistencyGroupVolumes); i++ {
		if swag.IsZero(m.SnapmirrorEndpointInlineConsistencyGroupVolumes[i]) { // not required
			continue
		}

		if m.SnapmirrorEndpointInlineConsistencyGroupVolumes[i] != nil {
			if err := m.SnapmirrorEndpointInlineConsistencyGroupVolumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("consistency_group_volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnapmirrorEndpoint) validateSvm(formats strfmt.Registry) error {
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

// ContextValidate validate this snapmirror endpoint based on the context it is used
func (m *SnapmirrorEndpoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapmirrorEndpointInlineConsistencyGroupVolumes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorEndpoint) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *SnapmirrorEndpoint) contextValidateSnapmirrorEndpointInlineConsistencyGroupVolumes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SnapmirrorEndpointInlineConsistencyGroupVolumes); i++ {

		if m.SnapmirrorEndpointInlineConsistencyGroupVolumes[i] != nil {
			if err := m.SnapmirrorEndpointInlineConsistencyGroupVolumes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("consistency_group_volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnapmirrorEndpoint) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *SnapmirrorEndpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapmirrorEndpoint) UnmarshalBinary(b []byte) error {
	var res SnapmirrorEndpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapmirrorEndpointInlineCluster snapmirror endpoint inline cluster
//
// swagger:model snapmirror_endpoint_inline_cluster
type SnapmirrorEndpointInlineCluster struct {

	// links
	Links *SnapmirrorEndpointInlineClusterInlineLinks `json:"_links,omitempty"`

	// name
	// Example: cluster1
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Format: uuid
	UUID *strfmt.UUID `json:"uuid,omitempty"`
}

// Validate validates this snapmirror endpoint inline cluster
func (m *SnapmirrorEndpointInlineCluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorEndpointInlineCluster) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *SnapmirrorEndpointInlineCluster) validateUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.UUID) { // not required
		return nil
	}

	if err := validate.FormatOf("cluster"+"."+"uuid", "body", "uuid", m.UUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this snapmirror endpoint inline cluster based on the context it is used
func (m *SnapmirrorEndpointInlineCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorEndpointInlineCluster) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapmirrorEndpointInlineCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapmirrorEndpointInlineCluster) UnmarshalBinary(b []byte) error {
	var res SnapmirrorEndpointInlineCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapmirrorEndpointInlineClusterInlineLinks snapmirror endpoint inline cluster inline links
//
// swagger:model snapmirror_endpoint_inline_cluster_inline__links
type SnapmirrorEndpointInlineClusterInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this snapmirror endpoint inline cluster inline links
func (m *SnapmirrorEndpointInlineClusterInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorEndpointInlineClusterInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this snapmirror endpoint inline cluster inline links based on the context it is used
func (m *SnapmirrorEndpointInlineClusterInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorEndpointInlineClusterInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapmirrorEndpointInlineClusterInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapmirrorEndpointInlineClusterInlineLinks) UnmarshalBinary(b []byte) error {
	var res SnapmirrorEndpointInlineClusterInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapmirrorEndpointInlineConsistencyGroupVolumesInlineArrayItem snapmirror endpoint inline consistency group volumes inline array item
//
// swagger:model snapmirror_endpoint_inline_consistency_group_volumes_inline_array_item
type SnapmirrorEndpointInlineConsistencyGroupVolumesInlineArrayItem struct {

	// The name of the volume.
	// Example: volume1
	Name *string `json:"name,omitempty"`

	// Unique identifier for the volume. This corresponds to the instance-uuid that is exposed in the CLI and ONTAPI. It does not change due to a volume move.
	// Example: 028baa66-41bd-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this snapmirror endpoint inline consistency group volumes inline array item
func (m *SnapmirrorEndpointInlineConsistencyGroupVolumesInlineArrayItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this snapmirror endpoint inline consistency group volumes inline array item based on context it is used
func (m *SnapmirrorEndpointInlineConsistencyGroupVolumesInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SnapmirrorEndpointInlineConsistencyGroupVolumesInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapmirrorEndpointInlineConsistencyGroupVolumesInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res SnapmirrorEndpointInlineConsistencyGroupVolumesInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapmirrorEndpointInlineSvm snapmirror endpoint inline svm
//
// swagger:model snapmirror_endpoint_inline_svm
type SnapmirrorEndpointInlineSvm struct {

	// links
	Links *SnapmirrorEndpointInlineSvmInlineLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name *string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this snapmirror endpoint inline svm
func (m *SnapmirrorEndpointInlineSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorEndpointInlineSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this snapmirror endpoint inline svm based on the context it is used
func (m *SnapmirrorEndpointInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorEndpointInlineSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SnapmirrorEndpointInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapmirrorEndpointInlineSvm) UnmarshalBinary(b []byte) error {
	var res SnapmirrorEndpointInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SnapmirrorEndpointInlineSvmInlineLinks snapmirror endpoint inline svm inline links
//
// swagger:model snapmirror_endpoint_inline_svm_inline__links
type SnapmirrorEndpointInlineSvmInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this snapmirror endpoint inline svm inline links
func (m *SnapmirrorEndpointInlineSvmInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorEndpointInlineSvmInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this snapmirror endpoint inline svm inline links based on the context it is used
func (m *SnapmirrorEndpointInlineSvmInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapmirrorEndpointInlineSvmInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SnapmirrorEndpointInlineSvmInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapmirrorEndpointInlineSvmInlineLinks) UnmarshalBinary(b []byte) error {
	var res SnapmirrorEndpointInlineSvmInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
