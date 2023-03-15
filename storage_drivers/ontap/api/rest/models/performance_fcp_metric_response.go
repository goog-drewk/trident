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

// PerformanceFcpMetricResponse performance fcp metric response
//
// swagger:model performance_fcp_metric_response
type PerformanceFcpMetricResponse struct {

	// links
	Links *PerformanceFcpMetricResponseInlineLinks `json:"_links,omitempty"`

	// Number of records
	// Example: 1
	NumRecords *int64 `json:"num_records,omitempty"`

	// performance fcp metric response inline records
	PerformanceFcpMetricResponseInlineRecords []*PerformanceFcpMetricResponseInlineRecordsInlineArrayItem `json:"records,omitempty"`
}

// Validate validates this performance fcp metric response
func (m *PerformanceFcpMetricResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerformanceFcpMetricResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceFcpMetricResponse) validateLinks(formats strfmt.Registry) error {
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

func (m *PerformanceFcpMetricResponse) validatePerformanceFcpMetricResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.PerformanceFcpMetricResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(m.PerformanceFcpMetricResponseInlineRecords); i++ {
		if swag.IsZero(m.PerformanceFcpMetricResponseInlineRecords[i]) { // not required
			continue
		}

		if m.PerformanceFcpMetricResponseInlineRecords[i] != nil {
			if err := m.PerformanceFcpMetricResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this performance fcp metric response based on the context it is used
func (m *PerformanceFcpMetricResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePerformanceFcpMetricResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceFcpMetricResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceFcpMetricResponse) contextValidatePerformanceFcpMetricResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PerformanceFcpMetricResponseInlineRecords); i++ {

		if m.PerformanceFcpMetricResponseInlineRecords[i] != nil {
			if err := m.PerformanceFcpMetricResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (m *PerformanceFcpMetricResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceFcpMetricResponse) UnmarshalBinary(b []byte) error {
	var res PerformanceFcpMetricResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceFcpMetricResponseInlineLinks performance fcp metric response inline links
//
// swagger:model performance_fcp_metric_response_inline__links
type PerformanceFcpMetricResponseInlineLinks struct {

	// next
	Next *Href `json:"next,omitempty"`

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this performance fcp metric response inline links
func (m *PerformanceFcpMetricResponseInlineLinks) Validate(formats strfmt.Registry) error {
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

func (m *PerformanceFcpMetricResponseInlineLinks) validateNext(formats strfmt.Registry) error {
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

func (m *PerformanceFcpMetricResponseInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this performance fcp metric response inline links based on the context it is used
func (m *PerformanceFcpMetricResponseInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *PerformanceFcpMetricResponseInlineLinks) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceFcpMetricResponseInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PerformanceFcpMetricResponseInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceFcpMetricResponseInlineLinks) UnmarshalBinary(b []byte) error {
	var res PerformanceFcpMetricResponseInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceFcpMetricResponseInlineRecordsInlineArrayItem Performance numbers, such as IOPS latency and throughput, for SVM protocols.
//
// swagger:model performance_fcp_metric_response_inline_records_inline_array_item
type PerformanceFcpMetricResponseInlineRecordsInlineArrayItem struct {

	// links
	Links *PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineLinks `json:"_links,omitempty"`

	// The duration over which this sample is calculated. The time durations are represented in the ISO-8601 standard format. Samples can be calculated over the following durations:
	//
	// Example: PT15S
	// Read Only: true
	// Enum: [PT15S PT4M PT30M PT2H P1D PT5M]
	Duration *string `json:"duration,omitempty"`

	// iops
	Iops *PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineIops `json:"iops,omitempty"`

	// latency
	Latency *PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineLatency `json:"latency,omitempty"`

	// Any errors associated with the sample. For example, if the aggregation of data over multiple nodes fails then any of the partial errors might be returned, "ok" on success, or "error" on any internal uncategorized failure. Whenever a sample collection is missed but done at a later time, it is back filled to the previous 15 second timestamp and tagged with "backfilled_data". "Inconsistent_ delta_time" is encountered when the time between two collections is not the same for all nodes. Therefore, the aggregated value might be over or under inflated. "Negative_delta" is returned when an expected monotonically increasing value has decreased in value. "Inconsistent_old_data" is returned when one or more nodes do not have the latest data.
	// Example: ok
	// Read Only: true
	// Enum: [ok error partial_no_data partial_no_response partial_other_error negative_delta not_found backfilled_data inconsistent_delta_time inconsistent_old_data partial_no_uuid]
	Status *string `json:"status,omitempty"`

	// svm
	Svm *PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineSvm `json:"svm,omitempty"`

	// throughput
	Throughput *PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineThroughput `json:"throughput,omitempty"`

	// The timestamp of the performance data.
	// Example: 2017-01-25T11:20:13Z
	// Read Only: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this performance fcp metric response inline records inline array item
func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItem) Validate(formats strfmt.Registry) error {
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

	if err := m.validateSvm(formats); err != nil {
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

func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItem) validateLinks(formats strfmt.Registry) error {
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

var performanceFcpMetricResponseInlineRecordsInlineArrayItemTypeDurationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PT15S","PT4M","PT30M","PT2H","P1D","PT5M"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceFcpMetricResponseInlineRecordsInlineArrayItemTypeDurationPropEnum = append(performanceFcpMetricResponseInlineRecordsInlineArrayItemTypeDurationPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// performance_fcp_metric_response_inline_records_inline_array_item
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT15S
	// END DEBUGGING
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItemDurationPT15S captures enum value "PT15S"
	PerformanceFcpMetricResponseInlineRecordsInlineArrayItemDurationPT15S string = "PT15S"

	// BEGIN DEBUGGING
	// performance_fcp_metric_response_inline_records_inline_array_item
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT4M
	// END DEBUGGING
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItemDurationPT4M captures enum value "PT4M"
	PerformanceFcpMetricResponseInlineRecordsInlineArrayItemDurationPT4M string = "PT4M"

	// BEGIN DEBUGGING
	// performance_fcp_metric_response_inline_records_inline_array_item
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT30M
	// END DEBUGGING
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItemDurationPT30M captures enum value "PT30M"
	PerformanceFcpMetricResponseInlineRecordsInlineArrayItemDurationPT30M string = "PT30M"

	// BEGIN DEBUGGING
	// performance_fcp_metric_response_inline_records_inline_array_item
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT2H
	// END DEBUGGING
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItemDurationPT2H captures enum value "PT2H"
	PerformanceFcpMetricResponseInlineRecordsInlineArrayItemDurationPT2H string = "PT2H"

	// BEGIN DEBUGGING
	// performance_fcp_metric_response_inline_records_inline_array_item
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// P1D
	// END DEBUGGING
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItemDurationP1D captures enum value "P1D"
	PerformanceFcpMetricResponseInlineRecordsInlineArrayItemDurationP1D string = "P1D"

	// BEGIN DEBUGGING
	// performance_fcp_metric_response_inline_records_inline_array_item
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItem
	// duration
	// Duration
	// PT5M
	// END DEBUGGING
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItemDurationPT5M captures enum value "PT5M"
	PerformanceFcpMetricResponseInlineRecordsInlineArrayItemDurationPT5M string = "PT5M"
)

// prop value enum
func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItem) validateDurationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceFcpMetricResponseInlineRecordsInlineArrayItemTypeDurationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItem) validateDuration(formats strfmt.Registry) error {
	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	// value enum
	if err := m.validateDurationEnum("duration", "body", *m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItem) validateIops(formats strfmt.Registry) error {
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

func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItem) validateLatency(formats strfmt.Registry) error {
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

var performanceFcpMetricResponseInlineRecordsInlineArrayItemTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ok","error","partial_no_data","partial_no_response","partial_other_error","negative_delta","not_found","backfilled_data","inconsistent_delta_time","inconsistent_old_data","partial_no_uuid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceFcpMetricResponseInlineRecordsInlineArrayItemTypeStatusPropEnum = append(performanceFcpMetricResponseInlineRecordsInlineArrayItemTypeStatusPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// performance_fcp_metric_response_inline_records_inline_array_item
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// ok
	// END DEBUGGING
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItemStatusOk captures enum value "ok"
	PerformanceFcpMetricResponseInlineRecordsInlineArrayItemStatusOk string = "ok"

	// BEGIN DEBUGGING
	// performance_fcp_metric_response_inline_records_inline_array_item
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// error
	// END DEBUGGING
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItemStatusError captures enum value "error"
	PerformanceFcpMetricResponseInlineRecordsInlineArrayItemStatusError string = "error"

	// BEGIN DEBUGGING
	// performance_fcp_metric_response_inline_records_inline_array_item
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// partial_no_data
	// END DEBUGGING
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItemStatusPartialNoData captures enum value "partial_no_data"
	PerformanceFcpMetricResponseInlineRecordsInlineArrayItemStatusPartialNoData string = "partial_no_data"

	// BEGIN DEBUGGING
	// performance_fcp_metric_response_inline_records_inline_array_item
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// partial_no_response
	// END DEBUGGING
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItemStatusPartialNoResponse captures enum value "partial_no_response"
	PerformanceFcpMetricResponseInlineRecordsInlineArrayItemStatusPartialNoResponse string = "partial_no_response"

	// BEGIN DEBUGGING
	// performance_fcp_metric_response_inline_records_inline_array_item
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// partial_other_error
	// END DEBUGGING
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItemStatusPartialOtherError captures enum value "partial_other_error"
	PerformanceFcpMetricResponseInlineRecordsInlineArrayItemStatusPartialOtherError string = "partial_other_error"

	// BEGIN DEBUGGING
	// performance_fcp_metric_response_inline_records_inline_array_item
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// negative_delta
	// END DEBUGGING
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItemStatusNegativeDelta captures enum value "negative_delta"
	PerformanceFcpMetricResponseInlineRecordsInlineArrayItemStatusNegativeDelta string = "negative_delta"

	// BEGIN DEBUGGING
	// performance_fcp_metric_response_inline_records_inline_array_item
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// not_found
	// END DEBUGGING
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItemStatusNotFound captures enum value "not_found"
	PerformanceFcpMetricResponseInlineRecordsInlineArrayItemStatusNotFound string = "not_found"

	// BEGIN DEBUGGING
	// performance_fcp_metric_response_inline_records_inline_array_item
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// backfilled_data
	// END DEBUGGING
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItemStatusBackfilledData captures enum value "backfilled_data"
	PerformanceFcpMetricResponseInlineRecordsInlineArrayItemStatusBackfilledData string = "backfilled_data"

	// BEGIN DEBUGGING
	// performance_fcp_metric_response_inline_records_inline_array_item
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// inconsistent_delta_time
	// END DEBUGGING
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItemStatusInconsistentDeltaTime captures enum value "inconsistent_delta_time"
	PerformanceFcpMetricResponseInlineRecordsInlineArrayItemStatusInconsistentDeltaTime string = "inconsistent_delta_time"

	// BEGIN DEBUGGING
	// performance_fcp_metric_response_inline_records_inline_array_item
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// inconsistent_old_data
	// END DEBUGGING
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItemStatusInconsistentOldData captures enum value "inconsistent_old_data"
	PerformanceFcpMetricResponseInlineRecordsInlineArrayItemStatusInconsistentOldData string = "inconsistent_old_data"

	// BEGIN DEBUGGING
	// performance_fcp_metric_response_inline_records_inline_array_item
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItem
	// status
	// Status
	// partial_no_uuid
	// END DEBUGGING
	// PerformanceFcpMetricResponseInlineRecordsInlineArrayItemStatusPartialNoUUID captures enum value "partial_no_uuid"
	PerformanceFcpMetricResponseInlineRecordsInlineArrayItemStatusPartialNoUUID string = "partial_no_uuid"
)

// prop value enum
func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItem) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceFcpMetricResponseInlineRecordsInlineArrayItemTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItem) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItem) validateSvm(formats strfmt.Registry) error {
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

func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItem) validateThroughput(formats strfmt.Registry) error {
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

func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItem) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this performance fcp metric response inline records inline array item based on the context it is used
func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateThroughput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItem) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItem) contextValidateDuration(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItem) contextValidateIops(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItem) contextValidateLatency(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItem) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItem) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItem) contextValidateThroughput(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItem) contextValidateTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItem) UnmarshalBinary(b []byte) error {
	var res PerformanceFcpMetricResponseInlineRecordsInlineArrayItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineIops The rate of I/O operations observed at the storage object.
//
// swagger:model performance_fcp_metric_response_inline_records_inline_array_item_inline_iops
type PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineIops struct {

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

// Validate validates this performance fcp metric response inline records inline array item inline iops
func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineIops) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance fcp metric response inline records inline array item inline iops based on the context it is used
func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineIops) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineIops) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineIops) UnmarshalBinary(b []byte) error {
	var res PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineIops
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineLatency The round trip latency in microseconds observed at the storage object.
//
// swagger:model performance_fcp_metric_response_inline_records_inline_array_item_inline_latency
type PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineLatency struct {

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

// Validate validates this performance fcp metric response inline records inline array item inline latency
func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineLatency) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance fcp metric response inline records inline array item inline latency based on the context it is used
func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineLatency) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineLatency) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineLatency) UnmarshalBinary(b []byte) error {
	var res PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineLatency
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineLinks performance fcp metric response inline records inline array item inline links
//
// swagger:model performance_fcp_metric_response_inline_records_inline_array_item_inline__links
type PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this performance fcp metric response inline records inline array item inline links
func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this performance fcp metric response inline records inline array item inline links based on the context it is used
func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineLinks) UnmarshalBinary(b []byte) error {
	var res PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineSvm performance fcp metric response inline records inline array item inline svm
//
// swagger:model performance_fcp_metric_response_inline_records_inline_array_item_inline_svm
type PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineSvm struct {

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this performance fcp metric response inline records inline array item inline svm
func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineSvm) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this performance fcp metric response inline records inline array item inline svm based on context it is used
func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineSvm) UnmarshalBinary(b []byte) error {
	var res PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineThroughput The rate of throughput bytes per second observed at the storage object.
//
// swagger:model performance_fcp_metric_response_inline_records_inline_array_item_inline_throughput
type PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineThroughput struct {

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

// Validate validates this performance fcp metric response inline records inline array item inline throughput
func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineThroughput) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this performance fcp metric response inline records inline array item inline throughput based on the context it is used
func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineThroughput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineThroughput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineThroughput) UnmarshalBinary(b []byte) error {
	var res PerformanceFcpMetricResponseInlineRecordsInlineArrayItemInlineThroughput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
