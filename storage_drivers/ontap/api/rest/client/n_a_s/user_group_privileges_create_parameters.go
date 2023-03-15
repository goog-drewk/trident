// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// NewUserGroupPrivilegesCreateParams creates a new UserGroupPrivilegesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserGroupPrivilegesCreateParams() *UserGroupPrivilegesCreateParams {
	return &UserGroupPrivilegesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserGroupPrivilegesCreateParamsWithTimeout creates a new UserGroupPrivilegesCreateParams object
// with the ability to set a timeout on a request.
func NewUserGroupPrivilegesCreateParamsWithTimeout(timeout time.Duration) *UserGroupPrivilegesCreateParams {
	return &UserGroupPrivilegesCreateParams{
		timeout: timeout,
	}
}

// NewUserGroupPrivilegesCreateParamsWithContext creates a new UserGroupPrivilegesCreateParams object
// with the ability to set a context for a request.
func NewUserGroupPrivilegesCreateParamsWithContext(ctx context.Context) *UserGroupPrivilegesCreateParams {
	return &UserGroupPrivilegesCreateParams{
		Context: ctx,
	}
}

// NewUserGroupPrivilegesCreateParamsWithHTTPClient creates a new UserGroupPrivilegesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserGroupPrivilegesCreateParamsWithHTTPClient(client *http.Client) *UserGroupPrivilegesCreateParams {
	return &UserGroupPrivilegesCreateParams{
		HTTPClient: client,
	}
}

/*
UserGroupPrivilegesCreateParams contains all the parameters to send to the API endpoint

	for the user group privileges create operation.

	Typically these are written to a http.Request.
*/
type UserGroupPrivilegesCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.UserGroupPrivileges

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user group privileges create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserGroupPrivilegesCreateParams) WithDefaults() *UserGroupPrivilegesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user group privileges create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserGroupPrivilegesCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := UserGroupPrivilegesCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the user group privileges create params
func (o *UserGroupPrivilegesCreateParams) WithTimeout(timeout time.Duration) *UserGroupPrivilegesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user group privileges create params
func (o *UserGroupPrivilegesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user group privileges create params
func (o *UserGroupPrivilegesCreateParams) WithContext(ctx context.Context) *UserGroupPrivilegesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user group privileges create params
func (o *UserGroupPrivilegesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user group privileges create params
func (o *UserGroupPrivilegesCreateParams) WithHTTPClient(client *http.Client) *UserGroupPrivilegesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user group privileges create params
func (o *UserGroupPrivilegesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the user group privileges create params
func (o *UserGroupPrivilegesCreateParams) WithInfo(info *models.UserGroupPrivileges) *UserGroupPrivilegesCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the user group privileges create params
func (o *UserGroupPrivilegesCreateParams) SetInfo(info *models.UserGroupPrivileges) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the user group privileges create params
func (o *UserGroupPrivilegesCreateParams) WithReturnRecords(returnRecords *bool) *UserGroupPrivilegesCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the user group privileges create params
func (o *UserGroupPrivilegesCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *UserGroupPrivilegesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
