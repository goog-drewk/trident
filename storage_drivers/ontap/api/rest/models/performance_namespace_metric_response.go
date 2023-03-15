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

// PerformanceNamespaceMetricResponse performance namespace metric response
//
// swagger:model performance_namespace_metric_response
type PerformanceNamespaceMetricResponse struct {

	// links
	Links *PerformanceNamespaceMetricResponseInlineLinks `json:"_links,omitempty"`

	// Number of records
	// Example: 1
	NumRecords *int64 `json:"num_records,omitempty"`

	// performance namespace metric response inline records
	PerformanceNamespaceMetricResponseInlineRecords []*PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem `json:"records,omitempty"`
}

// Validate validates this performance namespace metric response
func (m *PerformanceNamespaceMetricResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerformanceNamespaceMetricResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceNamespaceMetricResponse) validateLinks(formats strfmt.Registry) error {
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

func (m *PerformanceNamespaceMetricResponse) validatePerformanceNamespaceMetricResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.PerformanceNamespaceMetricResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(m.PerformanceNamespaceMetricResponseInlineRecords); i++ {
		if swag.IsZero(m.PerformanceNamespaceMetricResponseInlineRecords[i]) { // not required
			continue
		}

		if m.PerformanceNamespaceMetricResponseInlineRecords[i] != nil {
			if err := m.PerformanceNamespaceMetricResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this performance namespace metric response based on the context it is used
func (m *PerformanceNamespaceMetricResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePerformanceNamespaceMetricResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceNamespaceMetricResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceNamespaceMetricResponse) contextValidatePerformanceNamespaceMetricResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PerformanceNamespaceMetricResponseInlineRecords); i++ {

		if m.PerformanceNamespaceMetricResponseInlineRecords[i] != nil {
			if err := m.PerformanceNamespaceMetricResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceNamespaceMetricResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceNamespaceMetricResponse) UnmarshalBinary(b []byte) error {
	var res PerformanceNamespaceMetricResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceNamespaceMetricResponseInlineLinks performance namespace metric response inline links
//
// swagger:model performance_namespace_metric_response_inline__links
type PerformanceNamespaceMetricResponseInlineLinks struct {

	// next
	Next *Href `json:"next,omitempty"`

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this performance namespace metric response inline links
func (m *PerformanceNamespaceMetricResponseInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceNamespaceMetricResponseInlineLinks) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceNamespaceMetricResponseInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this performance namespace metric response inline links based on the context it is used
func (m *PerformanceNamespaceMetricResponseInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceNamespaceMetricResponseInlineLinks) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

	if m.Next != nil {
		if err := m.Next.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceNamespaceMetricResponseInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PerformanceNamespaceMetricResponseInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceNamespaceMetricResponseInlineLinks) UnmarshalBinary(b []byte) error {
	var res PerformanceNamespaceMetricResponseInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem Performance numbers, such as IOPS latency and throughput, for SVM protocols.
//
// swagger:model performance_namespace_metric_response_inline_records_inline_array_item
type PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem struct {

	// links
	Links *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineLinks `json:"_links,omitempty"`

	// The duration over which this sample is calculated. The time durations are represented in the ISO-8601 standard format. Samples can be calculated over the following durations:
	//
	// Example: PT15S
	// Read Only: true
	// Enum: [PT15S PT4M PT30M PT2H P1D PT5M]
	Duration *string `json:"duration,omitempty"`

	// iops
	Iops *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineIops `json:"iops,omitempty"`

	// latency
	Latency *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineLatency `json:"latency,omitempty"`

	// Any errors associated with the sample. For example, if the aggregation of data over multiple nodes fails then any of the partial errors might be returned, "ok" on success, or "error" on any internal uncategorized failure. Whenever a sample collection is missed but done at a later time, it is back filled to the previous 15 second timestamp and tagged with "backfilled_data". "Inconsistent_ delta_time" is encountered when the time between two collections is not the same for all nodes. Therefore, the aggregated value might be over or under inflated. "Negative_delta" is returned when an expected monotonically increasing value has decreased in value. "Inconsistent_old_data" is returned when one or more nodes do not have the latest data.
	// Example: ok
	// Read Only: true
	// Enum: [ok error partial_no_data partial_no_response partial_other_error negative_delta not_found backfilled_data inconsistent_delta_time inconsistent_old_data partial_no_uuid]
	Status *string `json:"status,omitempty"`

	// throughput
	Throughput *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineThroughput `json:"throughput,omitempty"`

	// The timestamp of the performance data.
	// Example: 2017-01-25T11:20:13Z
	// Read Only: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`

	// The unique identifier of the NVMe namespace.
	//
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this performance namespace metric response inline records inline array item
func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIops(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateThroughput(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem) validateLinks(formats strfmt.Registry) error {
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

var performanceNamespaceMetricResponseInlineRecordsInlineArrayItemTypeDurationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PT15S","PT4M","PT30M","PT2H","P1D","PT5M"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceNamespaceMetricResponseInlineRecordsInlineArrayItemTypeDurationPropEnum = append(performanceNamespaceMetricResponseInlineRecordsInlineArrayItemTypeDurationPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// performance_namespace_metric_response_inline_records_inline_array_item
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT15S
	// END DEBUGGING
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemDurationPT15S captures enum value "PT15S"
	PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemDurationPT15S string = "PT15S"

	// BEGIN DEBUGGING
	// performance_namespace_metric_response_inline_records_inline_array_item
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT4M
	// END DEBUGGING
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemDurationPT4M captures enum value "PT4M"
	PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemDurationPT4M string = "PT4M"

	// BEGIN DEBUGGING
	// performance_namespace_metric_response_inline_records_inline_array_item
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT30M
	// END DEBUGGING
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemDurationPT30M captures enum value "PT30M"
	PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemDurationPT30M string = "PT30M"

	// BEGIN DEBUGGING
	// performance_namespace_metric_response_inline_records_inline_array_item
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT2H
	// END DEBUGGING
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemDurationPT2H captures enum value "PT2H"
	PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemDurationPT2H string = "PT2H"

	// BEGIN DEBUGGING
	// performance_namespace_metric_response_inline_records_inline_array_item
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// P1D
	// END DEBUGGING
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemDurationP1D captures enum value "P1D"
	PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemDurationP1D string = "P1D"

	// BEGIN DEBUGGING
	// performance_namespace_metric_response_inline_records_inline_array_item
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT5M
	// END DEBUGGING
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemDurationPT5M captures enum value "PT5M"
	PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemDurationPT5M string = "PT5M"
)

// prop value enum
func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem) validateDurationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceNamespaceMetricResponseInlineRecordsInlineArrayItemTypeDurationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem) validateDuration(formats strfmt.Registry) error {
	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	// value enum
	if err := m.validateDurationEnum("duration", "body", *m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem) validateIops(formats strfmt.Registry) error {
	if swag.IsZero(m.Iops) { // not required
		return nil
	}

	if m.Iops != nil {
		if err := m.Iops.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem) validateLatency(formats strfmt.Registry) error {
	if swag.IsZero(m.Latency) { // not required
		return nil
	}

	if m.Latency != nil {
		if err := m.Latency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latency")
			}
			return err
		}
	}

	return nil
}

var performanceNamespaceMetricResponseInlineRecordsInlineArrayItemTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","error","partial_no_data","partial_no_response","partial_other_error","negative_delta","not_found","backfilled_data","inconsistent_delta_time","inconsistent_old_data","partial_no_uuid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceNamespaceMetricResponseInlineRecordsInlineArrayItemTypeStatusPropEnum = append(performanceNamespaceMetricResponseInlineRecordsInlineArrayItemTypeStatusPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// performance_namespace_metric_response_inline_records_inline_array_item
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// ok
	// END DEBUGGING
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemStatusOk captures enum value "ok"
	PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemStatusOk string = "ok"

	// BEGIN DEBUGGING
	// performance_namespace_metric_response_inline_records_inline_array_item
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// error
	// END DEBUGGING
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemStatusError captures enum value "error"
	PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemStatusError string = "error"

	// BEGIN DEBUGGING
	// performance_namespace_metric_response_inline_records_inline_array_item
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// partial_no_data
	// END DEBUGGING
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemStatusPartialNoData captures enum value "partial_no_data"
	PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemStatusPartialNoData string = "partial_no_data"

	// BEGIN DEBUGGING
	// performance_namespace_metric_response_inline_records_inline_array_item
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// partial_no_response
	// END DEBUGGING
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemStatusPartialNoResponse captures enum value "partial_no_response"
	PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemStatusPartialNoResponse string = "partial_no_response"

	// BEGIN DEBUGGING
	// performance_namespace_metric_response_inline_records_inline_array_item
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// partial_other_error
	// END DEBUGGING
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemStatusPartialOtherError captures enum value "partial_other_error"
	PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemStatusPartialOtherError string = "partial_other_error"

	// BEGIN DEBUGGING
	// performance_namespace_metric_response_inline_records_inline_array_item
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// negative_delta
	// END DEBUGGING
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemStatusNegativeDelta captures enum value "negative_delta"
	PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemStatusNegativeDelta string = "negative_delta"

	// BEGIN DEBUGGING
	// performance_namespace_metric_response_inline_records_inline_array_item
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// not_found
	// END DEBUGGING
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemStatusNotFound captures enum value "not_found"
	PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemStatusNotFound string = "not_found"

	// BEGIN DEBUGGING
	// performance_namespace_metric_response_inline_records_inline_array_item
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// backfilled_data
	// END DEBUGGING
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemStatusBackfilledData captures enum value "backfilled_data"
	PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemStatusBackfilledData string = "backfilled_data"

	// BEGIN DEBUGGING
	// performance_namespace_metric_response_inline_records_inline_array_item
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// inconsistent_delta_time
	// END DEBUGGING
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemStatusInconsistentDeltaTime captures enum value "inconsistent_delta_time"
	PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemStatusInconsistentDeltaTime string = "inconsistent_delta_time"

	// BEGIN DEBUGGING
	// performance_namespace_metric_response_inline_records_inline_array_item
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// inconsistent_old_data
	// END DEBUGGING
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemStatusInconsistentOldData captures enum value "inconsistent_old_data"
	PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemStatusInconsistentOldData string = "inconsistent_old_data"

	// BEGIN DEBUGGING
	// performance_namespace_metric_response_inline_records_inline_array_item
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// partial_no_uuid
	// END DEBUGGING
	// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemStatusPartialNoUUID captures enum value "partial_no_uuid"
	PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemStatusPartialNoUUID string = "partial_no_uuid"
)

// prop value enum
func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceNamespaceMetricResponseInlineRecordsInlineArrayItemTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem) validateThroughput(formats strfmt.Registry) error {
	if swag.IsZero(m.Throughput) { // not required
		return nil
	}

	if m.Throughput != nil {
		if err := m.Throughput.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this performance namespace metric response inline records inline array item based on the context it is used
func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDuration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIops(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLatency(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThroughput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem) contextValidateDuration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem) contextValidateIops(ctx context.Context, formats strfmt.Registry) error {

	if m.Iops != nil {
		if err := m.Iops.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("iops")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem) contextValidateLatency(ctx context.Context, formats strfmt.Registry) error {

	if m.Latency != nil {
		if err := m.Latency.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latency")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem) contextValidateThroughput(ctx context.Context, formats strfmt.Registry) error {

	if m.Throughput != nil {
		if err := m.Throughput.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("throughput")
			}
			return err
		}
	}

	return nil
}

func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem) contextValidateTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineIops The rate of I/O operations observed at the storage object.
//
// swagger:model performance_namespace_metric_response_inline_records_inline_array_item_inline_iops
type PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineIops struct {

	// Performance metric for other I/O operations. Other I/O operations can be metadata operations, such as directory lookups and so on.
	Other *int64 `json:"other,omitempty"`

	// Performance metric for read I/O operations.
	// Example: 200
	Read *int64 `json:"read,omitempty"`

	// Performance metric aggregated over all types of I/O operations.
	// Example: 1000
	Total *int64 `json:"total,omitempty"`

	// Peformance metric for write I/O operations.
	// Example: 100
	Write *int64 `json:"write,omitempty"`
}

// Validate validates this performance namespace metric response inline records inline array item inline iops
func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineIops) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance namespace metric response inline records inline array item inline iops based on the context it is used
func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineIops) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineIops) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineIops) UnmarshalBinary(b []byte) error {
	var res PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineIops
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineLatency The round trip latency in microseconds observed at the storage object.
//
// swagger:model performance_namespace_metric_response_inline_records_inline_array_item_inline_latency
type PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineLatency struct {

	// Performance metric for other I/O operations. Other I/O operations can be metadata operations, such as directory lookups and so on.
	Other *int64 `json:"other,omitempty"`

	// Performance metric for read I/O operations.
	// Example: 200
	Read *int64 `json:"read,omitempty"`

	// Performance metric aggregated over all types of I/O operations.
	// Example: 1000
	Total *int64 `json:"total,omitempty"`

	// Peformance metric for write I/O operations.
	// Example: 100
	Write *int64 `json:"write,omitempty"`
}

// Validate validates this performance namespace metric response inline records inline array item inline latency
func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineLatency) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance namespace metric response inline records inline array item inline latency based on the context it is used
func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineLatency) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineLatency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineLatency) UnmarshalBinary(b []byte) error {
	var res PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineLatency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineLinks performance namespace metric response inline records inline array item inline links
//
// swagger:model performance_namespace_metric_response_inline_records_inline_array_item_inline__links
type PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this performance namespace metric response inline records inline array item inline links
func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this performance namespace metric response inline records inline array item inline links based on the context it is used
func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineLinks) UnmarshalBinary(b []byte) error {
	var res PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineThroughput The rate of throughput bytes per second observed at the storage object.
//
// swagger:model performance_namespace_metric_response_inline_records_inline_array_item_inline_throughput
type PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineThroughput struct {

	// Performance metric for read I/O operations.
	// Example: 200
	Read *int64 `json:"read,omitempty"`

	// Performance metric aggregated over all types of I/O operations.
	// Example: 1000
	Total *int64 `json:"total,omitempty"`

	// Peformance metric for write I/O operations.
	// Example: 100
	Write *int64 `json:"write,omitempty"`
}

// Validate validates this performance namespace metric response inline records inline array item inline throughput
func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineThroughput) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance namespace metric response inline records inline array item inline throughput based on the context it is used
func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineThroughput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineThroughput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineThroughput) UnmarshalBinary(b []byte) error {
	var res PerformanceNamespaceMetricResponseInlineRecordsInlineArrayItemInlineThroughput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
