// Code generated by go-swagger; DO NOT EDIT.

package name_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewNetgroupsSettingsCollectionGetParams creates a new NetgroupsSettingsCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNetgroupsSettingsCollectionGetParams() *NetgroupsSettingsCollectionGetParams {
	return &NetgroupsSettingsCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNetgroupsSettingsCollectionGetParamsWithTimeout creates a new NetgroupsSettingsCollectionGetParams object
// with the ability to set a timeout on a request.
func NewNetgroupsSettingsCollectionGetParamsWithTimeout(timeout time.Duration) *NetgroupsSettingsCollectionGetParams {
	return &NetgroupsSettingsCollectionGetParams{
		timeout: timeout,
	}
}

// NewNetgroupsSettingsCollectionGetParamsWithContext creates a new NetgroupsSettingsCollectionGetParams object
// with the ability to set a context for a request.
func NewNetgroupsSettingsCollectionGetParamsWithContext(ctx context.Context) *NetgroupsSettingsCollectionGetParams {
	return &NetgroupsSettingsCollectionGetParams{
		Context: ctx,
	}
}

// NewNetgroupsSettingsCollectionGetParamsWithHTTPClient creates a new NetgroupsSettingsCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNetgroupsSettingsCollectionGetParamsWithHTTPClient(client *http.Client) *NetgroupsSettingsCollectionGetParams {
	return &NetgroupsSettingsCollectionGetParams{
		HTTPClient: client,
	}
}

/*
NetgroupsSettingsCollectionGetParams contains all the parameters to send to the API endpoint

	for the netgroups settings collection get operation.

	Typically these are written to a http.Request.
*/
type NetgroupsSettingsCollectionGetParams struct {

	/* Enabled.

	   Filter by enabled
	*/
	Enabled *bool

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* NegativeCacheEnabledByhost.

	   Filter by negative_cache_enabled_byhost
	*/
	NegativeCacheEnabledByhost *bool

	/* NegativeTTLByhost.

	   Filter by negative_ttl_byhost
	*/
	NegativeTTLByhost *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeout *int64

	/* SvmName.

	   Filter by svm.name
	*/
	SvmName *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SvmUUID *string

	/* TTLByhost.

	   Filter by ttl_byhost
	*/
	TTLByhost *string

	/* TTLForMembers.

	   Filter by ttl_for_members
	*/
	TTLForMembers *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the netgroups settings collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetgroupsSettingsCollectionGetParams) WithDefaults() *NetgroupsSettingsCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the netgroups settings collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetgroupsSettingsCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := NetgroupsSettingsCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) WithTimeout(timeout time.Duration) *NetgroupsSettingsCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) WithContext(ctx context.Context) *NetgroupsSettingsCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) WithHTTPClient(client *http.Client) *NetgroupsSettingsCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnabled adds the enabled to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) WithEnabled(enabled *bool) *NetgroupsSettingsCollectionGetParams {
	o.SetEnabled(enabled)
	return o
}

// SetEnabled adds the enabled to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) SetEnabled(enabled *bool) {
	o.Enabled = enabled
}

// WithFields adds the fields to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) WithFields(fields []string) *NetgroupsSettingsCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) WithMaxRecords(maxRecords *int64) *NetgroupsSettingsCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithNegativeCacheEnabledByhost adds the negativeCacheEnabledByhost to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) WithNegativeCacheEnabledByhost(negativeCacheEnabledByhost *bool) *NetgroupsSettingsCollectionGetParams {
	o.SetNegativeCacheEnabledByhost(negativeCacheEnabledByhost)
	return o
}

// SetNegativeCacheEnabledByhost adds the negativeCacheEnabledByhost to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) SetNegativeCacheEnabledByhost(negativeCacheEnabledByhost *bool) {
	o.NegativeCacheEnabledByhost = negativeCacheEnabledByhost
}

// WithNegativeTTLByhost adds the negativeTTLByhost to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) WithNegativeTTLByhost(negativeTTLByhost *string) *NetgroupsSettingsCollectionGetParams {
	o.SetNegativeTTLByhost(negativeTTLByhost)
	return o
}

// SetNegativeTTLByhost adds the negativeTtlByhost to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) SetNegativeTTLByhost(negativeTTLByhost *string) {
	o.NegativeTTLByhost = negativeTTLByhost
}

// WithOrderBy adds the orderBy to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) WithOrderBy(orderBy []string) *NetgroupsSettingsCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) WithReturnRecords(returnRecords *bool) *NetgroupsSettingsCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *NetgroupsSettingsCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSvmName adds the svmName to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) WithSvmName(svmName *string) *NetgroupsSettingsCollectionGetParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) WithSvmUUID(svmUUID *string) *NetgroupsSettingsCollectionGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WithTTLByhost adds the tTLByhost to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) WithTTLByhost(tTLByhost *string) *NetgroupsSettingsCollectionGetParams {
	o.SetTTLByhost(tTLByhost)
	return o
}

// SetTTLByhost adds the ttlByhost to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) SetTTLByhost(tTLByhost *string) {
	o.TTLByhost = tTLByhost
}

// WithTTLForMembers adds the tTLForMembers to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) WithTTLForMembers(tTLForMembers *string) *NetgroupsSettingsCollectionGetParams {
	o.SetTTLForMembers(tTLForMembers)
	return o
}

// SetTTLForMembers adds the ttlForMembers to the netgroups settings collection get params
func (o *NetgroupsSettingsCollectionGetParams) SetTTLForMembers(tTLForMembers *string) {
	o.TTLForMembers = tTLForMembers
}

// WriteToRequest writes these params to a swagger request
func (o *NetgroupsSettingsCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Enabled != nil {

		// query param enabled
		var qrEnabled bool

		if o.Enabled != nil {
			qrEnabled = *o.Enabled
		}
		qEnabled := swag.FormatBool(qrEnabled)
		if qEnabled != "" {

			if err := r.SetQueryParam("enabled", qEnabled); err != nil {
				return err
			}
		}
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.MaxRecords != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecords != nil {
			qrMaxRecords = *o.MaxRecords
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.NegativeCacheEnabledByhost != nil {

		// query param negative_cache_enabled_byhost
		var qrNegativeCacheEnabledByhost bool

		if o.NegativeCacheEnabledByhost != nil {
			qrNegativeCacheEnabledByhost = *o.NegativeCacheEnabledByhost
		}
		qNegativeCacheEnabledByhost := swag.FormatBool(qrNegativeCacheEnabledByhost)
		if qNegativeCacheEnabledByhost != "" {

			if err := r.SetQueryParam("negative_cache_enabled_byhost", qNegativeCacheEnabledByhost); err != nil {
				return err
			}
		}
	}

	if o.NegativeTTLByhost != nil {

		// query param negative_ttl_byhost
		var qrNegativeTTLByhost string

		if o.NegativeTTLByhost != nil {
			qrNegativeTTLByhost = *o.NegativeTTLByhost
		}
		qNegativeTTLByhost := qrNegativeTTLByhost
		if qNegativeTTLByhost != "" {

			if err := r.SetQueryParam("negative_ttl_byhost", qNegativeTTLByhost); err != nil {
				return err
			}
		}
	}

	if o.OrderBy != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.ReturnRecords != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecords != nil {
			qrReturnRecords = *o.ReturnRecords
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeout != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeout != nil {
			qrReturnTimeout = *o.ReturnTimeout
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if o.SvmName != nil {

		// query param svm.name
		var qrSvmName string

		if o.SvmName != nil {
			qrSvmName = *o.SvmName
		}
		qSvmName := qrSvmName
		if qSvmName != "" {

			if err := r.SetQueryParam("svm.name", qSvmName); err != nil {
				return err
			}
		}
	}

	if o.SvmUUID != nil {

		// query param svm.uuid
		var qrSvmUUID string

		if o.SvmUUID != nil {
			qrSvmUUID = *o.SvmUUID
		}
		qSvmUUID := qrSvmUUID
		if qSvmUUID != "" {

			if err := r.SetQueryParam("svm.uuid", qSvmUUID); err != nil {
				return err
			}
		}
	}

	if o.TTLByhost != nil {

		// query param ttl_byhost
		var qrTTLByhost string

		if o.TTLByhost != nil {
			qrTTLByhost = *o.TTLByhost
		}
		qTTLByhost := qrTTLByhost
		if qTTLByhost != "" {

			if err := r.SetQueryParam("ttl_byhost", qTTLByhost); err != nil {
				return err
			}
		}
	}

	if o.TTLForMembers != nil {

		// query param ttl_for_members
		var qrTTLForMembers string

		if o.TTLForMembers != nil {
			qrTTLForMembers = *o.TTLForMembers
		}
		qTTLForMembers := qrTTLForMembers
		if qTTLForMembers != "" {

			if err := r.SetQueryParam("ttl_for_members", qTTLForMembers); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamNetgroupsSettingsCollectionGet binds the parameter fields
func (o *NetgroupsSettingsCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}

// bindParamNetgroupsSettingsCollectionGet binds the parameter order_by
func (o *NetgroupsSettingsCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderBy

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}
