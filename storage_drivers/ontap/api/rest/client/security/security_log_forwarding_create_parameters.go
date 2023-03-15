// Code generated by go-swagger; DO NOT EDIT.

package security

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

// NewSecurityLogForwardingCreateParams creates a new SecurityLogForwardingCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityLogForwardingCreateParams() *SecurityLogForwardingCreateParams {
	return &SecurityLogForwardingCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityLogForwardingCreateParamsWithTimeout creates a new SecurityLogForwardingCreateParams object
// with the ability to set a timeout on a request.
func NewSecurityLogForwardingCreateParamsWithTimeout(timeout time.Duration) *SecurityLogForwardingCreateParams {
	return &SecurityLogForwardingCreateParams{
		timeout: timeout,
	}
}

// NewSecurityLogForwardingCreateParamsWithContext creates a new SecurityLogForwardingCreateParams object
// with the ability to set a context for a request.
func NewSecurityLogForwardingCreateParamsWithContext(ctx context.Context) *SecurityLogForwardingCreateParams {
	return &SecurityLogForwardingCreateParams{
		Context: ctx,
	}
}

// NewSecurityLogForwardingCreateParamsWithHTTPClient creates a new SecurityLogForwardingCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityLogForwardingCreateParamsWithHTTPClient(client *http.Client) *SecurityLogForwardingCreateParams {
	return &SecurityLogForwardingCreateParams{
		HTTPClient: client,
	}
}

/*
SecurityLogForwardingCreateParams contains all the parameters to send to the API endpoint

	for the security log forwarding create operation.

	Typically these are written to a http.Request.
*/
type SecurityLogForwardingCreateParams struct {

	/* Force.

	   Skip the Connectivity Test
	*/
	Force *bool

	/* Info.

	   Remote syslog/splunk server information
	*/
	Info *models.SecurityAuditLogForward

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security log forwarding create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityLogForwardingCreateParams) WithDefaults() *SecurityLogForwardingCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security log forwarding create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityLogForwardingCreateParams) SetDefaults() {
	var (
		forceDefault = bool(false)

		returnRecordsDefault = bool(false)

		returnTimeoutDefault = int64(0)
	)

	val := SecurityLogForwardingCreateParams{
		Force:         &forceDefault,
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the security log forwarding create params
func (o *SecurityLogForwardingCreateParams) WithTimeout(timeout time.Duration) *SecurityLogForwardingCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security log forwarding create params
func (o *SecurityLogForwardingCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security log forwarding create params
func (o *SecurityLogForwardingCreateParams) WithContext(ctx context.Context) *SecurityLogForwardingCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security log forwarding create params
func (o *SecurityLogForwardingCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security log forwarding create params
func (o *SecurityLogForwardingCreateParams) WithHTTPClient(client *http.Client) *SecurityLogForwardingCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security log forwarding create params
func (o *SecurityLogForwardingCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the security log forwarding create params
func (o *SecurityLogForwardingCreateParams) WithForce(force *bool) *SecurityLogForwardingCreateParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the security log forwarding create params
func (o *SecurityLogForwardingCreateParams) SetForce(force *bool) {
	o.Force = force
}

// WithInfo adds the info to the security log forwarding create params
func (o *SecurityLogForwardingCreateParams) WithInfo(info *models.SecurityAuditLogForward) *SecurityLogForwardingCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the security log forwarding create params
func (o *SecurityLogForwardingCreateParams) SetInfo(info *models.SecurityAuditLogForward) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the security log forwarding create params
func (o *SecurityLogForwardingCreateParams) WithReturnRecords(returnRecords *bool) *SecurityLogForwardingCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the security log forwarding create params
func (o *SecurityLogForwardingCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the security log forwarding create params
func (o *SecurityLogForwardingCreateParams) WithReturnTimeout(returnTimeout *int64) *SecurityLogForwardingCreateParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the security log forwarding create params
func (o *SecurityLogForwardingCreateParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityLogForwardingCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Force != nil {

		// query param force
		var qrForce bool

		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {

			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}
	}
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
