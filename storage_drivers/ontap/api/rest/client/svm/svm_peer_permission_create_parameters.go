// Code generated by go-swagger; DO NOT EDIT.

package svm

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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewSvmPeerPermissionCreateParams creates a new SvmPeerPermissionCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSvmPeerPermissionCreateParams() *SvmPeerPermissionCreateParams {
	return &SvmPeerPermissionCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSvmPeerPermissionCreateParamsWithTimeout creates a new SvmPeerPermissionCreateParams object
// with the ability to set a timeout on a request.
func NewSvmPeerPermissionCreateParamsWithTimeout(timeout time.Duration) *SvmPeerPermissionCreateParams {
	return &SvmPeerPermissionCreateParams{
		timeout: timeout,
	}
}

// NewSvmPeerPermissionCreateParamsWithContext creates a new SvmPeerPermissionCreateParams object
// with the ability to set a context for a request.
func NewSvmPeerPermissionCreateParamsWithContext(ctx context.Context) *SvmPeerPermissionCreateParams {
	return &SvmPeerPermissionCreateParams{
		Context: ctx,
	}
}

// NewSvmPeerPermissionCreateParamsWithHTTPClient creates a new SvmPeerPermissionCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSvmPeerPermissionCreateParamsWithHTTPClient(client *http.Client) *SvmPeerPermissionCreateParams {
	return &SvmPeerPermissionCreateParams{
		HTTPClient: client,
	}
}

/*
SvmPeerPermissionCreateParams contains all the parameters to send to the API endpoint

	for the svm peer permission create operation.

	Typically these are written to a http.Request.
*/
type SvmPeerPermissionCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.SvmPeerPermission

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the svm peer permission create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmPeerPermissionCreateParams) WithDefaults() *SvmPeerPermissionCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the svm peer permission create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmPeerPermissionCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := SvmPeerPermissionCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the svm peer permission create params
func (o *SvmPeerPermissionCreateParams) WithTimeout(timeout time.Duration) *SvmPeerPermissionCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the svm peer permission create params
func (o *SvmPeerPermissionCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the svm peer permission create params
func (o *SvmPeerPermissionCreateParams) WithContext(ctx context.Context) *SvmPeerPermissionCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the svm peer permission create params
func (o *SvmPeerPermissionCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the svm peer permission create params
func (o *SvmPeerPermissionCreateParams) WithHTTPClient(client *http.Client) *SvmPeerPermissionCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the svm peer permission create params
func (o *SvmPeerPermissionCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the svm peer permission create params
func (o *SvmPeerPermissionCreateParams) WithInfo(info *models.SvmPeerPermission) *SvmPeerPermissionCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the svm peer permission create params
func (o *SvmPeerPermissionCreateParams) SetInfo(info *models.SvmPeerPermission) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the svm peer permission create params
func (o *SvmPeerPermissionCreateParams) WithReturnRecords(returnRecords *bool) *SvmPeerPermissionCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the svm peer permission create params
func (o *SvmPeerPermissionCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *SvmPeerPermissionCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
