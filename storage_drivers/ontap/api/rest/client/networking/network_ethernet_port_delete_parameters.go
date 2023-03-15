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
)

// NewNetworkEthernetPortDeleteParams creates a new NetworkEthernetPortDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNetworkEthernetPortDeleteParams() *NetworkEthernetPortDeleteParams {
	return &NetworkEthernetPortDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkEthernetPortDeleteParamsWithTimeout creates a new NetworkEthernetPortDeleteParams object
// with the ability to set a timeout on a request.
func NewNetworkEthernetPortDeleteParamsWithTimeout(timeout time.Duration) *NetworkEthernetPortDeleteParams {
	return &NetworkEthernetPortDeleteParams{
		timeout: timeout,
	}
}

// NewNetworkEthernetPortDeleteParamsWithContext creates a new NetworkEthernetPortDeleteParams object
// with the ability to set a context for a request.
func NewNetworkEthernetPortDeleteParamsWithContext(ctx context.Context) *NetworkEthernetPortDeleteParams {
	return &NetworkEthernetPortDeleteParams{
		Context: ctx,
	}
}

// NewNetworkEthernetPortDeleteParamsWithHTTPClient creates a new NetworkEthernetPortDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewNetworkEthernetPortDeleteParamsWithHTTPClient(client *http.Client) *NetworkEthernetPortDeleteParams {
	return &NetworkEthernetPortDeleteParams{
		HTTPClient: client,
	}
}

/*
NetworkEthernetPortDeleteParams contains all the parameters to send to the API endpoint

	for the network ethernet port delete operation.

	Typically these are written to a http.Request.
*/
type NetworkEthernetPortDeleteParams struct {

	/* UUID.

	   Port UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the network ethernet port delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkEthernetPortDeleteParams) WithDefaults() *NetworkEthernetPortDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the network ethernet port delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkEthernetPortDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the network ethernet port delete params
func (o *NetworkEthernetPortDeleteParams) WithTimeout(timeout time.Duration) *NetworkEthernetPortDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network ethernet port delete params
func (o *NetworkEthernetPortDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network ethernet port delete params
func (o *NetworkEthernetPortDeleteParams) WithContext(ctx context.Context) *NetworkEthernetPortDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network ethernet port delete params
func (o *NetworkEthernetPortDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the network ethernet port delete params
func (o *NetworkEthernetPortDeleteParams) WithHTTPClient(client *http.Client) *NetworkEthernetPortDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the network ethernet port delete params
func (o *NetworkEthernetPortDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUID adds the uuid to the network ethernet port delete params
func (o *NetworkEthernetPortDeleteParams) WithUUID(uuid string) *NetworkEthernetPortDeleteParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the network ethernet port delete params
func (o *NetworkEthernetPortDeleteParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkEthernetPortDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
