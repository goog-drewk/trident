// Code generated by go-swagger; DO NOT EDIT.

package networking

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

// NewNetworkIPRouteGetParams creates a new NetworkIPRouteGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNetworkIPRouteGetParams() *NetworkIPRouteGetParams {
	return &NetworkIPRouteGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkIPRouteGetParamsWithTimeout creates a new NetworkIPRouteGetParams object
// with the ability to set a timeout on a request.
func NewNetworkIPRouteGetParamsWithTimeout(timeout time.Duration) *NetworkIPRouteGetParams {
	return &NetworkIPRouteGetParams{
		timeout: timeout,
	}
}

// NewNetworkIPRouteGetParamsWithContext creates a new NetworkIPRouteGetParams object
// with the ability to set a context for a request.
func NewNetworkIPRouteGetParamsWithContext(ctx context.Context) *NetworkIPRouteGetParams {
	return &NetworkIPRouteGetParams{
		Context: ctx,
	}
}

// NewNetworkIPRouteGetParamsWithHTTPClient creates a new NetworkIPRouteGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNetworkIPRouteGetParamsWithHTTPClient(client *http.Client) *NetworkIPRouteGetParams {
	return &NetworkIPRouteGetParams{
		HTTPClient: client,
	}
}

/*
NetworkIPRouteGetParams contains all the parameters to send to the API endpoint

	for the network ip route get operation.

	Typically these are written to a http.Request.
*/
type NetworkIPRouteGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* UUID.

	   Route UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the network ip route get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkIPRouteGetParams) WithDefaults() *NetworkIPRouteGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the network ip route get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkIPRouteGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the network ip route get params
func (o *NetworkIPRouteGetParams) WithTimeout(timeout time.Duration) *NetworkIPRouteGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network ip route get params
func (o *NetworkIPRouteGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network ip route get params
func (o *NetworkIPRouteGetParams) WithContext(ctx context.Context) *NetworkIPRouteGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network ip route get params
func (o *NetworkIPRouteGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the network ip route get params
func (o *NetworkIPRouteGetParams) WithHTTPClient(client *http.Client) *NetworkIPRouteGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the network ip route get params
func (o *NetworkIPRouteGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the network ip route get params
func (o *NetworkIPRouteGetParams) WithFields(fields []string) *NetworkIPRouteGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the network ip route get params
func (o *NetworkIPRouteGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithUUID adds the uuid to the network ip route get params
func (o *NetworkIPRouteGetParams) WithUUID(uuid string) *NetworkIPRouteGetParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the network ip route get params
func (o *NetworkIPRouteGetParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkIPRouteGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamNetworkIPRouteGet binds the parameter fields
func (o *NetworkIPRouteGetParams) bindParamFields(formats strfmt.Registry) []string {
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
