// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// NewClusterCollectionPerformanceMetricsGetParams creates a new ClusterCollectionPerformanceMetricsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClusterCollectionPerformanceMetricsGetParams() *ClusterCollectionPerformanceMetricsGetParams {
	return &ClusterCollectionPerformanceMetricsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClusterCollectionPerformanceMetricsGetParamsWithTimeout creates a new ClusterCollectionPerformanceMetricsGetParams object
// with the ability to set a timeout on a request.
func NewClusterCollectionPerformanceMetricsGetParamsWithTimeout(timeout time.Duration) *ClusterCollectionPerformanceMetricsGetParams {
	return &ClusterCollectionPerformanceMetricsGetParams{
		timeout: timeout,
	}
}

// NewClusterCollectionPerformanceMetricsGetParamsWithContext creates a new ClusterCollectionPerformanceMetricsGetParams object
// with the ability to set a context for a request.
func NewClusterCollectionPerformanceMetricsGetParamsWithContext(ctx context.Context) *ClusterCollectionPerformanceMetricsGetParams {
	return &ClusterCollectionPerformanceMetricsGetParams{
		Context: ctx,
	}
}

// NewClusterCollectionPerformanceMetricsGetParamsWithHTTPClient creates a new ClusterCollectionPerformanceMetricsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewClusterCollectionPerformanceMetricsGetParamsWithHTTPClient(client *http.Client) *ClusterCollectionPerformanceMetricsGetParams {
	return &ClusterCollectionPerformanceMetricsGetParams{
		HTTPClient: client,
	}
}

/*
ClusterCollectionPerformanceMetricsGetParams contains all the parameters to send to the API endpoint

	for the cluster collection performance metrics get operation.

	Typically these are written to a http.Request.
*/
type ClusterCollectionPerformanceMetricsGetParams struct {

	/* Duration.

	   Filter by duration
	*/
	Duration *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* Interval.

	     The time range for the data. Examples can be 1h, 1d, 1m, 1w, or 1y.
	The period for each time range is specified as follows:
	* 1h: Metrics over the most recent hour sampled over 15 seconds.
	* 1d: Metrics over the most recent day sampled over 5 minutes.
	* 1w: Metrics over the most recent week sampled over 30 minutes.
	* 1m: Metrics over the most recent month sampled over 2 hours.
	* 1y: Metrics over the most recent year sampled over a day.


	     Default: "1h"
	*/
	Interval *string

	/* IopsOther.

	   Filter by iops.other
	*/
	IopsOther *int64

	/* IopsRead.

	   Filter by iops.read
	*/
	IopsRead *int64

	/* IopsTotal.

	   Filter by iops.total
	*/
	IopsTotal *int64

	/* IopsWrite.

	   Filter by iops.write
	*/
	IopsWrite *int64

	/* LatencyOther.

	   Filter by latency.other
	*/
	LatencyOther *int64

	/* LatencyRead.

	   Filter by latency.read
	*/
	LatencyRead *int64

	/* LatencyTotal.

	   Filter by latency.total
	*/
	LatencyTotal *int64

	/* LatencyWrite.

	   Filter by latency.write
	*/
	LatencyWrite *int64

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

	/* Status.

	   Filter by status
	*/
	Status *string

	/* ThroughputOther.

	   Filter by throughput.other
	*/
	ThroughputOther *int64

	/* ThroughputRead.

	   Filter by throughput.read
	*/
	ThroughputRead *int64

	/* ThroughputTotal.

	   Filter by throughput.total
	*/
	ThroughputTotal *int64

	/* ThroughputWrite.

	   Filter by throughput.write
	*/
	ThroughputWrite *int64

	/* Timestamp.

	   Filter by timestamp
	*/
	Timestamp *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cluster collection performance metrics get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterCollectionPerformanceMetricsGetParams) WithDefaults() *ClusterCollectionPerformanceMetricsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cluster collection performance metrics get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterCollectionPerformanceMetricsGetParams) SetDefaults() {
	var (
		intervalDefault = string("1h")

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := ClusterCollectionPerformanceMetricsGetParams{
		Interval:      &intervalDefault,
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) WithTimeout(timeout time.Duration) *ClusterCollectionPerformanceMetricsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) WithContext(ctx context.Context) *ClusterCollectionPerformanceMetricsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) WithHTTPClient(client *http.Client) *ClusterCollectionPerformanceMetricsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDuration adds the duration to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) WithDuration(duration *string) *ClusterCollectionPerformanceMetricsGetParams {
	o.SetDuration(duration)
	return o
}

// SetDuration adds the duration to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) SetDuration(duration *string) {
	o.Duration = duration
}

// WithFields adds the fields to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) WithFields(fields []string) *ClusterCollectionPerformanceMetricsGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithInterval adds the interval to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) WithInterval(interval *string) *ClusterCollectionPerformanceMetricsGetParams {
	o.SetInterval(interval)
	return o
}

// SetInterval adds the interval to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) SetInterval(interval *string) {
	o.Interval = interval
}

// WithIopsOther adds the iopsOther to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) WithIopsOther(iopsOther *int64) *ClusterCollectionPerformanceMetricsGetParams {
	o.SetIopsOther(iopsOther)
	return o
}

// SetIopsOther adds the iopsOther to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) SetIopsOther(iopsOther *int64) {
	o.IopsOther = iopsOther
}

// WithIopsRead adds the iopsRead to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) WithIopsRead(iopsRead *int64) *ClusterCollectionPerformanceMetricsGetParams {
	o.SetIopsRead(iopsRead)
	return o
}

// SetIopsRead adds the iopsRead to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) SetIopsRead(iopsRead *int64) {
	o.IopsRead = iopsRead
}

// WithIopsTotal adds the iopsTotal to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) WithIopsTotal(iopsTotal *int64) *ClusterCollectionPerformanceMetricsGetParams {
	o.SetIopsTotal(iopsTotal)
	return o
}

// SetIopsTotal adds the iopsTotal to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) SetIopsTotal(iopsTotal *int64) {
	o.IopsTotal = iopsTotal
}

// WithIopsWrite adds the iopsWrite to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) WithIopsWrite(iopsWrite *int64) *ClusterCollectionPerformanceMetricsGetParams {
	o.SetIopsWrite(iopsWrite)
	return o
}

// SetIopsWrite adds the iopsWrite to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) SetIopsWrite(iopsWrite *int64) {
	o.IopsWrite = iopsWrite
}

// WithLatencyOther adds the latencyOther to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) WithLatencyOther(latencyOther *int64) *ClusterCollectionPerformanceMetricsGetParams {
	o.SetLatencyOther(latencyOther)
	return o
}

// SetLatencyOther adds the latencyOther to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) SetLatencyOther(latencyOther *int64) {
	o.LatencyOther = latencyOther
}

// WithLatencyRead adds the latencyRead to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) WithLatencyRead(latencyRead *int64) *ClusterCollectionPerformanceMetricsGetParams {
	o.SetLatencyRead(latencyRead)
	return o
}

// SetLatencyRead adds the latencyRead to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) SetLatencyRead(latencyRead *int64) {
	o.LatencyRead = latencyRead
}

// WithLatencyTotal adds the latencyTotal to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) WithLatencyTotal(latencyTotal *int64) *ClusterCollectionPerformanceMetricsGetParams {
	o.SetLatencyTotal(latencyTotal)
	return o
}

// SetLatencyTotal adds the latencyTotal to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) SetLatencyTotal(latencyTotal *int64) {
	o.LatencyTotal = latencyTotal
}

// WithLatencyWrite adds the latencyWrite to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) WithLatencyWrite(latencyWrite *int64) *ClusterCollectionPerformanceMetricsGetParams {
	o.SetLatencyWrite(latencyWrite)
	return o
}

// SetLatencyWrite adds the latencyWrite to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) SetLatencyWrite(latencyWrite *int64) {
	o.LatencyWrite = latencyWrite
}

// WithMaxRecords adds the maxRecords to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) WithMaxRecords(maxRecords *int64) *ClusterCollectionPerformanceMetricsGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) WithOrderBy(orderBy []string) *ClusterCollectionPerformanceMetricsGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) WithReturnRecords(returnRecords *bool) *ClusterCollectionPerformanceMetricsGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) WithReturnTimeout(returnTimeout *int64) *ClusterCollectionPerformanceMetricsGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithStatus adds the status to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) WithStatus(status *string) *ClusterCollectionPerformanceMetricsGetParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) SetStatus(status *string) {
	o.Status = status
}

// WithThroughputOther adds the throughputOther to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) WithThroughputOther(throughputOther *int64) *ClusterCollectionPerformanceMetricsGetParams {
	o.SetThroughputOther(throughputOther)
	return o
}

// SetThroughputOther adds the throughputOther to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) SetThroughputOther(throughputOther *int64) {
	o.ThroughputOther = throughputOther
}

// WithThroughputRead adds the throughputRead to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) WithThroughputRead(throughputRead *int64) *ClusterCollectionPerformanceMetricsGetParams {
	o.SetThroughputRead(throughputRead)
	return o
}

// SetThroughputRead adds the throughputRead to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) SetThroughputRead(throughputRead *int64) {
	o.ThroughputRead = throughputRead
}

// WithThroughputTotal adds the throughputTotal to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) WithThroughputTotal(throughputTotal *int64) *ClusterCollectionPerformanceMetricsGetParams {
	o.SetThroughputTotal(throughputTotal)
	return o
}

// SetThroughputTotal adds the throughputTotal to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) SetThroughputTotal(throughputTotal *int64) {
	o.ThroughputTotal = throughputTotal
}

// WithThroughputWrite adds the throughputWrite to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) WithThroughputWrite(throughputWrite *int64) *ClusterCollectionPerformanceMetricsGetParams {
	o.SetThroughputWrite(throughputWrite)
	return o
}

// SetThroughputWrite adds the throughputWrite to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) SetThroughputWrite(throughputWrite *int64) {
	o.ThroughputWrite = throughputWrite
}

// WithTimestamp adds the timestamp to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) WithTimestamp(timestamp *string) *ClusterCollectionPerformanceMetricsGetParams {
	o.SetTimestamp(timestamp)
	return o
}

// SetTimestamp adds the timestamp to the cluster collection performance metrics get params
func (o *ClusterCollectionPerformanceMetricsGetParams) SetTimestamp(timestamp *string) {
	o.Timestamp = timestamp
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterCollectionPerformanceMetricsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Duration != nil {

		// query param duration
		var qrDuration string

		if o.Duration != nil {
			qrDuration = *o.Duration
		}
		qDuration := qrDuration
		if qDuration != "" {

			if err := r.SetQueryParam("duration", qDuration); err != nil {
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

	if o.Interval != nil {

		// query param interval
		var qrInterval string

		if o.Interval != nil {
			qrInterval = *o.Interval
		}
		qInterval := qrInterval
		if qInterval != "" {

			if err := r.SetQueryParam("interval", qInterval); err != nil {
				return err
			}
		}
	}

	if o.IopsOther != nil {

		// query param iops.other
		var qrIopsOther int64

		if o.IopsOther != nil {
			qrIopsOther = *o.IopsOther
		}
		qIopsOther := swag.FormatInt64(qrIopsOther)
		if qIopsOther != "" {

			if err := r.SetQueryParam("iops.other", qIopsOther); err != nil {
				return err
			}
		}
	}

	if o.IopsRead != nil {

		// query param iops.read
		var qrIopsRead int64

		if o.IopsRead != nil {
			qrIopsRead = *o.IopsRead
		}
		qIopsRead := swag.FormatInt64(qrIopsRead)
		if qIopsRead != "" {

			if err := r.SetQueryParam("iops.read", qIopsRead); err != nil {
				return err
			}
		}
	}

	if o.IopsTotal != nil {

		// query param iops.total
		var qrIopsTotal int64

		if o.IopsTotal != nil {
			qrIopsTotal = *o.IopsTotal
		}
		qIopsTotal := swag.FormatInt64(qrIopsTotal)
		if qIopsTotal != "" {

			if err := r.SetQueryParam("iops.total", qIopsTotal); err != nil {
				return err
			}
		}
	}

	if o.IopsWrite != nil {

		// query param iops.write
		var qrIopsWrite int64

		if o.IopsWrite != nil {
			qrIopsWrite = *o.IopsWrite
		}
		qIopsWrite := swag.FormatInt64(qrIopsWrite)
		if qIopsWrite != "" {

			if err := r.SetQueryParam("iops.write", qIopsWrite); err != nil {
				return err
			}
		}
	}

	if o.LatencyOther != nil {

		// query param latency.other
		var qrLatencyOther int64

		if o.LatencyOther != nil {
			qrLatencyOther = *o.LatencyOther
		}
		qLatencyOther := swag.FormatInt64(qrLatencyOther)
		if qLatencyOther != "" {

			if err := r.SetQueryParam("latency.other", qLatencyOther); err != nil {
				return err
			}
		}
	}

	if o.LatencyRead != nil {

		// query param latency.read
		var qrLatencyRead int64

		if o.LatencyRead != nil {
			qrLatencyRead = *o.LatencyRead
		}
		qLatencyRead := swag.FormatInt64(qrLatencyRead)
		if qLatencyRead != "" {

			if err := r.SetQueryParam("latency.read", qLatencyRead); err != nil {
				return err
			}
		}
	}

	if o.LatencyTotal != nil {

		// query param latency.total
		var qrLatencyTotal int64

		if o.LatencyTotal != nil {
			qrLatencyTotal = *o.LatencyTotal
		}
		qLatencyTotal := swag.FormatInt64(qrLatencyTotal)
		if qLatencyTotal != "" {

			if err := r.SetQueryParam("latency.total", qLatencyTotal); err != nil {
				return err
			}
		}
	}

	if o.LatencyWrite != nil {

		// query param latency.write
		var qrLatencyWrite int64

		if o.LatencyWrite != nil {
			qrLatencyWrite = *o.LatencyWrite
		}
		qLatencyWrite := swag.FormatInt64(qrLatencyWrite)
		if qLatencyWrite != "" {

			if err := r.SetQueryParam("latency.write", qLatencyWrite); err != nil {
				return err
			}
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

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if o.ThroughputOther != nil {

		// query param throughput.other
		var qrThroughputOther int64

		if o.ThroughputOther != nil {
			qrThroughputOther = *o.ThroughputOther
		}
		qThroughputOther := swag.FormatInt64(qrThroughputOther)
		if qThroughputOther != "" {

			if err := r.SetQueryParam("throughput.other", qThroughputOther); err != nil {
				return err
			}
		}
	}

	if o.ThroughputRead != nil {

		// query param throughput.read
		var qrThroughputRead int64

		if o.ThroughputRead != nil {
			qrThroughputRead = *o.ThroughputRead
		}
		qThroughputRead := swag.FormatInt64(qrThroughputRead)
		if qThroughputRead != "" {

			if err := r.SetQueryParam("throughput.read", qThroughputRead); err != nil {
				return err
			}
		}
	}

	if o.ThroughputTotal != nil {

		// query param throughput.total
		var qrThroughputTotal int64

		if o.ThroughputTotal != nil {
			qrThroughputTotal = *o.ThroughputTotal
		}
		qThroughputTotal := swag.FormatInt64(qrThroughputTotal)
		if qThroughputTotal != "" {

			if err := r.SetQueryParam("throughput.total", qThroughputTotal); err != nil {
				return err
			}
		}
	}

	if o.ThroughputWrite != nil {

		// query param throughput.write
		var qrThroughputWrite int64

		if o.ThroughputWrite != nil {
			qrThroughputWrite = *o.ThroughputWrite
		}
		qThroughputWrite := swag.FormatInt64(qrThroughputWrite)
		if qThroughputWrite != "" {

			if err := r.SetQueryParam("throughput.write", qThroughputWrite); err != nil {
				return err
			}
		}
	}

	if o.Timestamp != nil {

		// query param timestamp
		var qrTimestamp string

		if o.Timestamp != nil {
			qrTimestamp = *o.Timestamp
		}
		qTimestamp := qrTimestamp
		if qTimestamp != "" {

			if err := r.SetQueryParam("timestamp", qTimestamp); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamClusterCollectionPerformanceMetricsGet binds the parameter fields
func (o *ClusterCollectionPerformanceMetricsGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamClusterCollectionPerformanceMetricsGet binds the parameter order_by
func (o *ClusterCollectionPerformanceMetricsGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
