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

// S3BucketPolicyCondition Information about policy conditions based on various condition operators and condition keys.
//
// swagger:model s3_bucket_policy_condition
type S3BucketPolicyCondition struct {

	// Condition operator that is applied to the specified condition key.
	// Example: ip_address
	// Enum: [ip_address not_ip_address string_equals string_not_equals string_equals_ignore_case string_not_equals_ignore_case string_like string_not_like numeric_equals numeric_not_equals numeric_greater_than numeric_greater_than_equals numeric_less_than numeric_less_than_equals]
	Operator *string `json:"operator,omitempty"`

	// An array of delimiters that are compared with the delimiter value specified at the time of execution of an S3-based command, using the condition operator specified.
	//
	// Example: ["/"]
	S3BucketPolicyConditionInlineDelimiters []*string `json:"delimiters,omitempty"`

	// An array of maximum keys that are allowed or denied to be retrieved using an S3 list operation, based on the condition operator specified.
	//
	// Example: [1000]
	S3BucketPolicyConditionInlineMaxKeys []*int64 `json:"max_keys,omitempty"`

	// An array of prefixes that are compared with the input prefix value specified at the time of execution of an S3-based command, using the condition operator specified.
	//
	// Example: ["pref"]
	S3BucketPolicyConditionInlinePrefixes []*string `json:"prefixes,omitempty"`

	// An array of IP address ranges that are compared with the IP address of a source command at the time of execution of an S3-based command, using the condition operator specified.
	//
	// Example: ["1.1.1.1","1.2.2.0/24"]
	S3BucketPolicyConditionInlineSourceIps []*string `json:"source_ips,omitempty"`

	// An array of usernames that a current user in the context is evaluated against using the condition operators.
	//
	// Example: ["user1"]
	S3BucketPolicyConditionInlineUsernames []*string `json:"usernames,omitempty"`
}

// Validate validates this s3 bucket policy condition
func (m *S3BucketPolicyCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var s3BucketPolicyConditionTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ip_address","not_ip_address","string_equals","string_not_equals","string_equals_ignore_case","string_not_equals_ignore_case","string_like","string_not_like","numeric_equals","numeric_not_equals","numeric_greater_than","numeric_greater_than_equals","numeric_less_than","numeric_less_than_equals"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		s3BucketPolicyConditionTypeOperatorPropEnum = append(s3BucketPolicyConditionTypeOperatorPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// s3_bucket_policy_condition
	// S3BucketPolicyCondition
	// operator
	// Operator
	// ip_address
	// END DEBUGGING
	// S3BucketPolicyConditionOperatorIPAddress captures enum value "ip_address"
	S3BucketPolicyConditionOperatorIPAddress string = "ip_address"

	// BEGIN DEBUGGING
	// s3_bucket_policy_condition
	// S3BucketPolicyCondition
	// operator
	// Operator
	// not_ip_address
	// END DEBUGGING
	// S3BucketPolicyConditionOperatorNotIPAddress captures enum value "not_ip_address"
	S3BucketPolicyConditionOperatorNotIPAddress string = "not_ip_address"

	// BEGIN DEBUGGING
	// s3_bucket_policy_condition
	// S3BucketPolicyCondition
	// operator
	// Operator
	// string_equals
	// END DEBUGGING
	// S3BucketPolicyConditionOperatorStringEquals captures enum value "string_equals"
	S3BucketPolicyConditionOperatorStringEquals string = "string_equals"

	// BEGIN DEBUGGING
	// s3_bucket_policy_condition
	// S3BucketPolicyCondition
	// operator
	// Operator
	// string_not_equals
	// END DEBUGGING
	// S3BucketPolicyConditionOperatorStringNotEquals captures enum value "string_not_equals"
	S3BucketPolicyConditionOperatorStringNotEquals string = "string_not_equals"

	// BEGIN DEBUGGING
	// s3_bucket_policy_condition
	// S3BucketPolicyCondition
	// operator
	// Operator
	// string_equals_ignore_case
	// END DEBUGGING
	// S3BucketPolicyConditionOperatorStringEqualsIgnoreCase captures enum value "string_equals_ignore_case"
	S3BucketPolicyConditionOperatorStringEqualsIgnoreCase string = "string_equals_ignore_case"

	// BEGIN DEBUGGING
	// s3_bucket_policy_condition
	// S3BucketPolicyCondition
	// operator
	// Operator
	// string_not_equals_ignore_case
	// END DEBUGGING
	// S3BucketPolicyConditionOperatorStringNotEqualsIgnoreCase captures enum value "string_not_equals_ignore_case"
	S3BucketPolicyConditionOperatorStringNotEqualsIgnoreCase string = "string_not_equals_ignore_case"

	// BEGIN DEBUGGING
	// s3_bucket_policy_condition
	// S3BucketPolicyCondition
	// operator
	// Operator
	// string_like
	// END DEBUGGING
	// S3BucketPolicyConditionOperatorStringLike captures enum value "string_like"
	S3BucketPolicyConditionOperatorStringLike string = "string_like"

	// BEGIN DEBUGGING
	// s3_bucket_policy_condition
	// S3BucketPolicyCondition
	// operator
	// Operator
	// string_not_like
	// END DEBUGGING
	// S3BucketPolicyConditionOperatorStringNotLike captures enum value "string_not_like"
	S3BucketPolicyConditionOperatorStringNotLike string = "string_not_like"

	// BEGIN DEBUGGING
	// s3_bucket_policy_condition
	// S3BucketPolicyCondition
	// operator
	// Operator
	// numeric_equals
	// END DEBUGGING
	// S3BucketPolicyConditionOperatorNumericEquals captures enum value "numeric_equals"
	S3BucketPolicyConditionOperatorNumericEquals string = "numeric_equals"

	// BEGIN DEBUGGING
	// s3_bucket_policy_condition
	// S3BucketPolicyCondition
	// operator
	// Operator
	// numeric_not_equals
	// END DEBUGGING
	// S3BucketPolicyConditionOperatorNumericNotEquals captures enum value "numeric_not_equals"
	S3BucketPolicyConditionOperatorNumericNotEquals string = "numeric_not_equals"

	// BEGIN DEBUGGING
	// s3_bucket_policy_condition
	// S3BucketPolicyCondition
	// operator
	// Operator
	// numeric_greater_than
	// END DEBUGGING
	// S3BucketPolicyConditionOperatorNumericGreaterThan captures enum value "numeric_greater_than"
	S3BucketPolicyConditionOperatorNumericGreaterThan string = "numeric_greater_than"

	// BEGIN DEBUGGING
	// s3_bucket_policy_condition
	// S3BucketPolicyCondition
	// operator
	// Operator
	// numeric_greater_than_equals
	// END DEBUGGING
	// S3BucketPolicyConditionOperatorNumericGreaterThanEquals captures enum value "numeric_greater_than_equals"
	S3BucketPolicyConditionOperatorNumericGreaterThanEquals string = "numeric_greater_than_equals"

	// BEGIN DEBUGGING
	// s3_bucket_policy_condition
	// S3BucketPolicyCondition
	// operator
	// Operator
	// numeric_less_than
	// END DEBUGGING
	// S3BucketPolicyConditionOperatorNumericLessThan captures enum value "numeric_less_than"
	S3BucketPolicyConditionOperatorNumericLessThan string = "numeric_less_than"

	// BEGIN DEBUGGING
	// s3_bucket_policy_condition
	// S3BucketPolicyCondition
	// operator
	// Operator
	// numeric_less_than_equals
	// END DEBUGGING
	// S3BucketPolicyConditionOperatorNumericLessThanEquals captures enum value "numeric_less_than_equals"
	S3BucketPolicyConditionOperatorNumericLessThanEquals string = "numeric_less_than_equals"
)

// prop value enum
func (m *S3BucketPolicyCondition) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, s3BucketPolicyConditionTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *S3BucketPolicyCondition) validateOperator(formats strfmt.Registry) error {
	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", *m.Operator); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this s3 bucket policy condition based on context it is used
func (m *S3BucketPolicyCondition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *S3BucketPolicyCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3BucketPolicyCondition) UnmarshalBinary(b []byte) error {
	var res S3BucketPolicyCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
