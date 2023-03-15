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
)

// NewUserGroupPrivilegesGetParams creates a new UserGroupPrivilegesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserGroupPrivilegesGetParams() *UserGroupPrivilegesGetParams {
	return &UserGroupPrivilegesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserGroupPrivilegesGetParamsWithTimeout creates a new UserGroupPrivilegesGetParams object
// with the ability to set a timeout on a request.
func NewUserGroupPrivilegesGetParamsWithTimeout(timeout time.Duration) *UserGroupPrivilegesGetParams {
	return &UserGroupPrivilegesGetParams{
		timeout: timeout,
	}
}

// NewUserGroupPrivilegesGetParamsWithContext creates a new UserGroupPrivilegesGetParams object
// with the ability to set a context for a request.
func NewUserGroupPrivilegesGetParamsWithContext(ctx context.Context) *UserGroupPrivilegesGetParams {
	return &UserGroupPrivilegesGetParams{
		Context: ctx,
	}
}

// NewUserGroupPrivilegesGetParamsWithHTTPClient creates a new UserGroupPrivilegesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserGroupPrivilegesGetParamsWithHTTPClient(client *http.Client) *UserGroupPrivilegesGetParams {
	return &UserGroupPrivilegesGetParams{
		HTTPClient: client,
	}
}

/*
UserGroupPrivilegesGetParams contains all the parameters to send to the API endpoint

	for the user group privileges get operation.

	Typically these are written to a http.Request.
*/
type UserGroupPrivilegesGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Name.

	   Local or Active Directory user or group name.
	*/
	Name string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the user group privileges get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserGroupPrivilegesGetParams) WithDefaults() *UserGroupPrivilegesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user group privileges get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserGroupPrivilegesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user group privileges get params
func (o *UserGroupPrivilegesGetParams) WithTimeout(timeout time.Duration) *UserGroupPrivilegesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user group privileges get params
func (o *UserGroupPrivilegesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user group privileges get params
func (o *UserGroupPrivilegesGetParams) WithContext(ctx context.Context) *UserGroupPrivilegesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user group privileges get params
func (o *UserGroupPrivilegesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user group privileges get params
func (o *UserGroupPrivilegesGetParams) WithHTTPClient(client *http.Client) *UserGroupPrivilegesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user group privileges get params
func (o *UserGroupPrivilegesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the user group privileges get params
func (o *UserGroupPrivilegesGetParams) WithFields(fields []string) *UserGroupPrivilegesGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the user group privileges get params
func (o *UserGroupPrivilegesGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithName adds the name to the user group privileges get params
func (o *UserGroupPrivilegesGetParams) WithName(name string) *UserGroupPrivilegesGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the user group privileges get params
func (o *UserGroupPrivilegesGetParams) SetName(name string) {
	o.Name = name
}

// WithSvmUUID adds the svmUUID to the user group privileges get params
func (o *UserGroupPrivilegesGetParams) WithSvmUUID(svmUUID string) *UserGroupPrivilegesGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the user group privileges get params
func (o *UserGroupPrivilegesGetParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *UserGroupPrivilegesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamUserGroupPrivilegesGet binds the parameter fields
func (o *UserGroupPrivilegesGetParams) bindParamFields(formats strfmt.Registry) []string {
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
