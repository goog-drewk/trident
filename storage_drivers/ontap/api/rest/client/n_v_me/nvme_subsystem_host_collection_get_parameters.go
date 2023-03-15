// Code generated by go-swagger; DO NOT EDIT.

package n_v_me

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

// NewNvmeSubsystemHostCollectionGetParams creates a new NvmeSubsystemHostCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNvmeSubsystemHostCollectionGetParams() *NvmeSubsystemHostCollectionGetParams {
	return &NvmeSubsystemHostCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNvmeSubsystemHostCollectionGetParamsWithTimeout creates a new NvmeSubsystemHostCollectionGetParams object
// with the ability to set a timeout on a request.
func NewNvmeSubsystemHostCollectionGetParamsWithTimeout(timeout time.Duration) *NvmeSubsystemHostCollectionGetParams {
	return &NvmeSubsystemHostCollectionGetParams{
		timeout: timeout,
	}
}

// NewNvmeSubsystemHostCollectionGetParamsWithContext creates a new NvmeSubsystemHostCollectionGetParams object
// with the ability to set a context for a request.
func NewNvmeSubsystemHostCollectionGetParamsWithContext(ctx context.Context) *NvmeSubsystemHostCollectionGetParams {
	return &NvmeSubsystemHostCollectionGetParams{
		Context: ctx,
	}
}

// NewNvmeSubsystemHostCollectionGetParamsWithHTTPClient creates a new NvmeSubsystemHostCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNvmeSubsystemHostCollectionGetParamsWithHTTPClient(client *http.Client) *NvmeSubsystemHostCollectionGetParams {
	return &NvmeSubsystemHostCollectionGetParams{
		HTTPClient: client,
	}
}

/*
NvmeSubsystemHostCollectionGetParams contains all the parameters to send to the API endpoint

	for the nvme subsystem host collection get operation.

	Typically these are written to a http.Request.
*/
type NvmeSubsystemHostCollectionGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

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

	/* SubsystemUUID.

	   The unique identifier of the NVMe subsystem.

	*/
	SubsystemUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nvme subsystem host collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NvmeSubsystemHostCollectionGetParams) WithDefaults() *NvmeSubsystemHostCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nvme subsystem host collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NvmeSubsystemHostCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := NvmeSubsystemHostCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) WithTimeout(timeout time.Duration) *NvmeSubsystemHostCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) WithContext(ctx context.Context) *NvmeSubsystemHostCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) WithHTTPClient(client *http.Client) *NvmeSubsystemHostCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) WithFields(fields []string) *NvmeSubsystemHostCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) WithMaxRecords(maxRecords *int64) *NvmeSubsystemHostCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) WithOrderBy(orderBy []string) *NvmeSubsystemHostCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) WithReturnRecords(returnRecords *bool) *NvmeSubsystemHostCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *NvmeSubsystemHostCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSubsystemUUID adds the subsystemUUID to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) WithSubsystemUUID(subsystemUUID string) *NvmeSubsystemHostCollectionGetParams {
	o.SetSubsystemUUID(subsystemUUID)
	return o
}

// SetSubsystemUUID adds the subsystemUuid to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) SetSubsystemUUID(subsystemUUID string) {
	o.SubsystemUUID = subsystemUUID
}

// WriteToRequest writes these params to a swagger request
func (o *NvmeSubsystemHostCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param subsystem.uuid
	if err := r.SetPathParam("subsystem.uuid", o.SubsystemUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamNvmeSubsystemHostCollectionGet binds the parameter fields
func (o *NvmeSubsystemHostCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamNvmeSubsystemHostCollectionGet binds the parameter order_by
func (o *NvmeSubsystemHostCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
