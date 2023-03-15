// Code generated by go-swagger; DO NOT EDIT.

package object_store

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

// NewS3PolicyGetParams creates a new S3PolicyGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewS3PolicyGetParams() *S3PolicyGetParams {
	return &S3PolicyGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewS3PolicyGetParamsWithTimeout creates a new S3PolicyGetParams object
// with the ability to set a timeout on a request.
func NewS3PolicyGetParamsWithTimeout(timeout time.Duration) *S3PolicyGetParams {
	return &S3PolicyGetParams{
		timeout: timeout,
	}
}

// NewS3PolicyGetParamsWithContext creates a new S3PolicyGetParams object
// with the ability to set a context for a request.
func NewS3PolicyGetParamsWithContext(ctx context.Context) *S3PolicyGetParams {
	return &S3PolicyGetParams{
		Context: ctx,
	}
}

// NewS3PolicyGetParamsWithHTTPClient creates a new S3PolicyGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewS3PolicyGetParamsWithHTTPClient(client *http.Client) *S3PolicyGetParams {
	return &S3PolicyGetParams{
		HTTPClient: client,
	}
}

/*
S3PolicyGetParams contains all the parameters to send to the API endpoint

	for the s3 policy get operation.

	Typically these are written to a http.Request.
*/
type S3PolicyGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Name.

	   Policy name
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

// WithDefaults hydrates default values in the s3 policy get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3PolicyGetParams) WithDefaults() *S3PolicyGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s3 policy get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3PolicyGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the s3 policy get params
func (o *S3PolicyGetParams) WithTimeout(timeout time.Duration) *S3PolicyGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s3 policy get params
func (o *S3PolicyGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s3 policy get params
func (o *S3PolicyGetParams) WithContext(ctx context.Context) *S3PolicyGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s3 policy get params
func (o *S3PolicyGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s3 policy get params
func (o *S3PolicyGetParams) WithHTTPClient(client *http.Client) *S3PolicyGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s3 policy get params
func (o *S3PolicyGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the s3 policy get params
func (o *S3PolicyGetParams) WithFields(fields []string) *S3PolicyGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the s3 policy get params
func (o *S3PolicyGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithName adds the name to the s3 policy get params
func (o *S3PolicyGetParams) WithName(name string) *S3PolicyGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the s3 policy get params
func (o *S3PolicyGetParams) SetName(name string) {
	o.Name = name
}

// WithSvmUUID adds the svmUUID to the s3 policy get params
func (o *S3PolicyGetParams) WithSvmUUID(svmUUID string) *S3PolicyGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the s3 policy get params
func (o *S3PolicyGetParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *S3PolicyGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamS3PolicyGet binds the parameter fields
func (o *S3PolicyGetParams) bindParamFields(formats strfmt.Registry) []string {
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
