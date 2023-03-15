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

// NewMetroclusterDrGroupCollectionGetParams creates a new MetroclusterDrGroupCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMetroclusterDrGroupCollectionGetParams() *MetroclusterDrGroupCollectionGetParams {
	return &MetroclusterDrGroupCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMetroclusterDrGroupCollectionGetParamsWithTimeout creates a new MetroclusterDrGroupCollectionGetParams object
// with the ability to set a timeout on a request.
func NewMetroclusterDrGroupCollectionGetParamsWithTimeout(timeout time.Duration) *MetroclusterDrGroupCollectionGetParams {
	return &MetroclusterDrGroupCollectionGetParams{
		timeout: timeout,
	}
}

// NewMetroclusterDrGroupCollectionGetParamsWithContext creates a new MetroclusterDrGroupCollectionGetParams object
// with the ability to set a context for a request.
func NewMetroclusterDrGroupCollectionGetParamsWithContext(ctx context.Context) *MetroclusterDrGroupCollectionGetParams {
	return &MetroclusterDrGroupCollectionGetParams{
		Context: ctx,
	}
}

// NewMetroclusterDrGroupCollectionGetParamsWithHTTPClient creates a new MetroclusterDrGroupCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewMetroclusterDrGroupCollectionGetParamsWithHTTPClient(client *http.Client) *MetroclusterDrGroupCollectionGetParams {
	return &MetroclusterDrGroupCollectionGetParams{
		HTTPClient: client,
	}
}

/*
MetroclusterDrGroupCollectionGetParams contains all the parameters to send to the API endpoint

	for the metrocluster dr group collection get operation.

	Typically these are written to a http.Request.
*/
type MetroclusterDrGroupCollectionGetParams struct {

	/* DrPairsNodeName.

	   Filter by dr_pairs.node.name
	*/
	DrPairsNodeName *string

	/* DrPairsNodeUUID.

	   Filter by dr_pairs.node.uuid
	*/
	DrPairsNodeUUID *string

	/* DrPairsPartnerName.

	   Filter by dr_pairs.partner.name
	*/
	DrPairsPartnerName *string

	/* DrPairsPartnerUUID.

	   Filter by dr_pairs.partner.uuid
	*/
	DrPairsPartnerUUID *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* ID.

	   Filter by id
	*/
	ID *int64

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* PartnerClusterName.

	   Filter by partner_cluster.name
	*/
	PartnerClusterName *string

	/* PartnerClusterUUID.

	   Filter by partner_cluster.uuid
	*/
	PartnerClusterUUID *string

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the metrocluster dr group collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MetroclusterDrGroupCollectionGetParams) WithDefaults() *MetroclusterDrGroupCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the metrocluster dr group collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MetroclusterDrGroupCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := MetroclusterDrGroupCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) WithTimeout(timeout time.Duration) *MetroclusterDrGroupCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) WithContext(ctx context.Context) *MetroclusterDrGroupCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) WithHTTPClient(client *http.Client) *MetroclusterDrGroupCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDrPairsNodeName adds the drPairsNodeName to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) WithDrPairsNodeName(drPairsNodeName *string) *MetroclusterDrGroupCollectionGetParams {
	o.SetDrPairsNodeName(drPairsNodeName)
	return o
}

// SetDrPairsNodeName adds the drPairsNodeName to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) SetDrPairsNodeName(drPairsNodeName *string) {
	o.DrPairsNodeName = drPairsNodeName
}

// WithDrPairsNodeUUID adds the drPairsNodeUUID to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) WithDrPairsNodeUUID(drPairsNodeUUID *string) *MetroclusterDrGroupCollectionGetParams {
	o.SetDrPairsNodeUUID(drPairsNodeUUID)
	return o
}

// SetDrPairsNodeUUID adds the drPairsNodeUuid to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) SetDrPairsNodeUUID(drPairsNodeUUID *string) {
	o.DrPairsNodeUUID = drPairsNodeUUID
}

// WithDrPairsPartnerName adds the drPairsPartnerName to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) WithDrPairsPartnerName(drPairsPartnerName *string) *MetroclusterDrGroupCollectionGetParams {
	o.SetDrPairsPartnerName(drPairsPartnerName)
	return o
}

// SetDrPairsPartnerName adds the drPairsPartnerName to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) SetDrPairsPartnerName(drPairsPartnerName *string) {
	o.DrPairsPartnerName = drPairsPartnerName
}

// WithDrPairsPartnerUUID adds the drPairsPartnerUUID to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) WithDrPairsPartnerUUID(drPairsPartnerUUID *string) *MetroclusterDrGroupCollectionGetParams {
	o.SetDrPairsPartnerUUID(drPairsPartnerUUID)
	return o
}

// SetDrPairsPartnerUUID adds the drPairsPartnerUuid to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) SetDrPairsPartnerUUID(drPairsPartnerUUID *string) {
	o.DrPairsPartnerUUID = drPairsPartnerUUID
}

// WithFields adds the fields to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) WithFields(fields []string) *MetroclusterDrGroupCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithID adds the id to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) WithID(id *int64) *MetroclusterDrGroupCollectionGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) SetID(id *int64) {
	o.ID = id
}

// WithMaxRecords adds the maxRecords to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) WithMaxRecords(maxRecords *int64) *MetroclusterDrGroupCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) WithOrderBy(orderBy []string) *MetroclusterDrGroupCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithPartnerClusterName adds the partnerClusterName to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) WithPartnerClusterName(partnerClusterName *string) *MetroclusterDrGroupCollectionGetParams {
	o.SetPartnerClusterName(partnerClusterName)
	return o
}

// SetPartnerClusterName adds the partnerClusterName to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) SetPartnerClusterName(partnerClusterName *string) {
	o.PartnerClusterName = partnerClusterName
}

// WithPartnerClusterUUID adds the partnerClusterUUID to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) WithPartnerClusterUUID(partnerClusterUUID *string) *MetroclusterDrGroupCollectionGetParams {
	o.SetPartnerClusterUUID(partnerClusterUUID)
	return o
}

// SetPartnerClusterUUID adds the partnerClusterUuid to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) SetPartnerClusterUUID(partnerClusterUUID *string) {
	o.PartnerClusterUUID = partnerClusterUUID
}

// WithReturnRecords adds the returnRecords to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) WithReturnRecords(returnRecords *bool) *MetroclusterDrGroupCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *MetroclusterDrGroupCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the metrocluster dr group collection get params
func (o *MetroclusterDrGroupCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *MetroclusterDrGroupCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DrPairsNodeName != nil {

		// query param dr_pairs.node.name
		var qrDrPairsNodeName string

		if o.DrPairsNodeName != nil {
			qrDrPairsNodeName = *o.DrPairsNodeName
		}
		qDrPairsNodeName := qrDrPairsNodeName
		if qDrPairsNodeName != "" {

			if err := r.SetQueryParam("dr_pairs.node.name", qDrPairsNodeName); err != nil {
				return err
			}
		}
	}

	if o.DrPairsNodeUUID != nil {

		// query param dr_pairs.node.uuid
		var qrDrPairsNodeUUID string

		if o.DrPairsNodeUUID != nil {
			qrDrPairsNodeUUID = *o.DrPairsNodeUUID
		}
		qDrPairsNodeUUID := qrDrPairsNodeUUID
		if qDrPairsNodeUUID != "" {

			if err := r.SetQueryParam("dr_pairs.node.uuid", qDrPairsNodeUUID); err != nil {
				return err
			}
		}
	}

	if o.DrPairsPartnerName != nil {

		// query param dr_pairs.partner.name
		var qrDrPairsPartnerName string

		if o.DrPairsPartnerName != nil {
			qrDrPairsPartnerName = *o.DrPairsPartnerName
		}
		qDrPairsPartnerName := qrDrPairsPartnerName
		if qDrPairsPartnerName != "" {

			if err := r.SetQueryParam("dr_pairs.partner.name", qDrPairsPartnerName); err != nil {
				return err
			}
		}
	}

	if o.DrPairsPartnerUUID != nil {

		// query param dr_pairs.partner.uuid
		var qrDrPairsPartnerUUID string

		if o.DrPairsPartnerUUID != nil {
			qrDrPairsPartnerUUID = *o.DrPairsPartnerUUID
		}
		qDrPairsPartnerUUID := qrDrPairsPartnerUUID
		if qDrPairsPartnerUUID != "" {

			if err := r.SetQueryParam("dr_pairs.partner.uuid", qDrPairsPartnerUUID); err != nil {
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

	if o.ID != nil {

		// query param id
		var qrID int64

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := swag.FormatInt64(qrID)
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
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

	if o.PartnerClusterName != nil {

		// query param partner_cluster.name
		var qrPartnerClusterName string

		if o.PartnerClusterName != nil {
			qrPartnerClusterName = *o.PartnerClusterName
		}
		qPartnerClusterName := qrPartnerClusterName
		if qPartnerClusterName != "" {

			if err := r.SetQueryParam("partner_cluster.name", qPartnerClusterName); err != nil {
				return err
			}
		}
	}

	if o.PartnerClusterUUID != nil {

		// query param partner_cluster.uuid
		var qrPartnerClusterUUID string

		if o.PartnerClusterUUID != nil {
			qrPartnerClusterUUID = *o.PartnerClusterUUID
		}
		qPartnerClusterUUID := qrPartnerClusterUUID
		if qPartnerClusterUUID != "" {

			if err := r.SetQueryParam("partner_cluster.uuid", qPartnerClusterUUID); err != nil {
				return err
			}
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

// bindParamMetroclusterDrGroupCollectionGet binds the parameter fields
func (o *MetroclusterDrGroupCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamMetroclusterDrGroupCollectionGet binds the parameter order_by
func (o *MetroclusterDrGroupCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
