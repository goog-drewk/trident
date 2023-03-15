// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// NewPortCollectionGetParams creates a new PortCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPortCollectionGetParams() *PortCollectionGetParams {
	return &PortCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPortCollectionGetParamsWithTimeout creates a new PortCollectionGetParams object
// with the ability to set a timeout on a request.
func NewPortCollectionGetParamsWithTimeout(timeout time.Duration) *PortCollectionGetParams {
	return &PortCollectionGetParams{
		timeout: timeout,
	}
}

// NewPortCollectionGetParamsWithContext creates a new PortCollectionGetParams object
// with the ability to set a context for a request.
func NewPortCollectionGetParamsWithContext(ctx context.Context) *PortCollectionGetParams {
	return &PortCollectionGetParams{
		Context: ctx,
	}
}

// NewPortCollectionGetParamsWithHTTPClient creates a new PortCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPortCollectionGetParamsWithHTTPClient(client *http.Client) *PortCollectionGetParams {
	return &PortCollectionGetParams{
		HTTPClient: client,
	}
}

/*
PortCollectionGetParams contains all the parameters to send to the API endpoint

	for the port collection get operation.

	Typically these are written to a http.Request.
*/
type PortCollectionGetParams struct {

	/* BoardName.

	   Filter by board_name
	*/
	BoardName *string

	/* CableIdentifier.

	   Filter by cable.identifier
	*/
	CableIdentifier *string

	/* CableLength.

	   Filter by cable.length
	*/
	CableLength *string

	/* CablePartNumber.

	   Filter by cable.part_number
	*/
	CablePartNumber *string

	/* CableSerialNumber.

	   Filter by cable.serial_number
	*/
	CableSerialNumber *string

	/* Description.

	   Filter by description
	*/
	Description *string

	/* Enabled.

	   Filter by enabled
	*/
	Enabled *bool

	/* ErrorCorrectiveAction.

	   Filter by error.corrective_action
	*/
	ErrorCorrectiveAction *string

	/* ErrorMessage.

	   Filter by error.message
	*/
	ErrorMessage *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* FirmwareVersion.

	   Filter by firmware_version
	*/
	FirmwareVersion *string

	/* InUse.

	   Filter by in_use
	*/
	InUse *bool

	/* MacAddress.

	   Filter by mac_address
	*/
	MacAddress *string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* Mode.

	   Filter by mode
	*/
	Mode *string

	/* Name.

	   Filter by name
	*/
	Name *string

	/* NodeName.

	   Filter by node.name
	*/
	NodeName *string

	/* NodeUUID.

	   Filter by node.uuid
	*/
	NodeUUID *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* PartNumber.

	   Filter by part_number
	*/
	PartNumber *string

	/* Redundant.

	   Filter by redundant
	*/
	Redundant *bool

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

	/* SerialNumber.

	   Filter by serial_number
	*/
	SerialNumber *string

	/* Speed.

	   Filter by speed
	*/
	Speed *float64

	/* State.

	   Filter by state
	*/
	State *string

	/* Type.

	   Filter by type
	*/
	Type *string

	/* Wwn.

	   Filter by wwn
	*/
	Wwn *string

	/* Wwpn.

	   Filter by wwpn
	*/
	Wwpn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the port collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PortCollectionGetParams) WithDefaults() *PortCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the port collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PortCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := PortCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the port collection get params
func (o *PortCollectionGetParams) WithTimeout(timeout time.Duration) *PortCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the port collection get params
func (o *PortCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the port collection get params
func (o *PortCollectionGetParams) WithContext(ctx context.Context) *PortCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the port collection get params
func (o *PortCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the port collection get params
func (o *PortCollectionGetParams) WithHTTPClient(client *http.Client) *PortCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the port collection get params
func (o *PortCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBoardName adds the boardName to the port collection get params
func (o *PortCollectionGetParams) WithBoardName(boardName *string) *PortCollectionGetParams {
	o.SetBoardName(boardName)
	return o
}

// SetBoardName adds the boardName to the port collection get params
func (o *PortCollectionGetParams) SetBoardName(boardName *string) {
	o.BoardName = boardName
}

// WithCableIdentifier adds the cableIdentifier to the port collection get params
func (o *PortCollectionGetParams) WithCableIdentifier(cableIdentifier *string) *PortCollectionGetParams {
	o.SetCableIdentifier(cableIdentifier)
	return o
}

// SetCableIdentifier adds the cableIdentifier to the port collection get params
func (o *PortCollectionGetParams) SetCableIdentifier(cableIdentifier *string) {
	o.CableIdentifier = cableIdentifier
}

// WithCableLength adds the cableLength to the port collection get params
func (o *PortCollectionGetParams) WithCableLength(cableLength *string) *PortCollectionGetParams {
	o.SetCableLength(cableLength)
	return o
}

// SetCableLength adds the cableLength to the port collection get params
func (o *PortCollectionGetParams) SetCableLength(cableLength *string) {
	o.CableLength = cableLength
}

// WithCablePartNumber adds the cablePartNumber to the port collection get params
func (o *PortCollectionGetParams) WithCablePartNumber(cablePartNumber *string) *PortCollectionGetParams {
	o.SetCablePartNumber(cablePartNumber)
	return o
}

// SetCablePartNumber adds the cablePartNumber to the port collection get params
func (o *PortCollectionGetParams) SetCablePartNumber(cablePartNumber *string) {
	o.CablePartNumber = cablePartNumber
}

// WithCableSerialNumber adds the cableSerialNumber to the port collection get params
func (o *PortCollectionGetParams) WithCableSerialNumber(cableSerialNumber *string) *PortCollectionGetParams {
	o.SetCableSerialNumber(cableSerialNumber)
	return o
}

// SetCableSerialNumber adds the cableSerialNumber to the port collection get params
func (o *PortCollectionGetParams) SetCableSerialNumber(cableSerialNumber *string) {
	o.CableSerialNumber = cableSerialNumber
}

// WithDescription adds the description to the port collection get params
func (o *PortCollectionGetParams) WithDescription(description *string) *PortCollectionGetParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the port collection get params
func (o *PortCollectionGetParams) SetDescription(description *string) {
	o.Description = description
}

// WithEnabled adds the enabled to the port collection get params
func (o *PortCollectionGetParams) WithEnabled(enabled *bool) *PortCollectionGetParams {
	o.SetEnabled(enabled)
	return o
}

// SetEnabled adds the enabled to the port collection get params
func (o *PortCollectionGetParams) SetEnabled(enabled *bool) {
	o.Enabled = enabled
}

// WithErrorCorrectiveAction adds the errorCorrectiveAction to the port collection get params
func (o *PortCollectionGetParams) WithErrorCorrectiveAction(errorCorrectiveAction *string) *PortCollectionGetParams {
	o.SetErrorCorrectiveAction(errorCorrectiveAction)
	return o
}

// SetErrorCorrectiveAction adds the errorCorrectiveAction to the port collection get params
func (o *PortCollectionGetParams) SetErrorCorrectiveAction(errorCorrectiveAction *string) {
	o.ErrorCorrectiveAction = errorCorrectiveAction
}

// WithErrorMessage adds the errorMessage to the port collection get params
func (o *PortCollectionGetParams) WithErrorMessage(errorMessage *string) *PortCollectionGetParams {
	o.SetErrorMessage(errorMessage)
	return o
}

// SetErrorMessage adds the errorMessage to the port collection get params
func (o *PortCollectionGetParams) SetErrorMessage(errorMessage *string) {
	o.ErrorMessage = errorMessage
}

// WithFields adds the fields to the port collection get params
func (o *PortCollectionGetParams) WithFields(fields []string) *PortCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the port collection get params
func (o *PortCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithFirmwareVersion adds the firmwareVersion to the port collection get params
func (o *PortCollectionGetParams) WithFirmwareVersion(firmwareVersion *string) *PortCollectionGetParams {
	o.SetFirmwareVersion(firmwareVersion)
	return o
}

// SetFirmwareVersion adds the firmwareVersion to the port collection get params
func (o *PortCollectionGetParams) SetFirmwareVersion(firmwareVersion *string) {
	o.FirmwareVersion = firmwareVersion
}

// WithInUse adds the inUse to the port collection get params
func (o *PortCollectionGetParams) WithInUse(inUse *bool) *PortCollectionGetParams {
	o.SetInUse(inUse)
	return o
}

// SetInUse adds the inUse to the port collection get params
func (o *PortCollectionGetParams) SetInUse(inUse *bool) {
	o.InUse = inUse
}

// WithMacAddress adds the macAddress to the port collection get params
func (o *PortCollectionGetParams) WithMacAddress(macAddress *string) *PortCollectionGetParams {
	o.SetMacAddress(macAddress)
	return o
}

// SetMacAddress adds the macAddress to the port collection get params
func (o *PortCollectionGetParams) SetMacAddress(macAddress *string) {
	o.MacAddress = macAddress
}

// WithMaxRecords adds the maxRecords to the port collection get params
func (o *PortCollectionGetParams) WithMaxRecords(maxRecords *int64) *PortCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the port collection get params
func (o *PortCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithMode adds the mode to the port collection get params
func (o *PortCollectionGetParams) WithMode(mode *string) *PortCollectionGetParams {
	o.SetMode(mode)
	return o
}

// SetMode adds the mode to the port collection get params
func (o *PortCollectionGetParams) SetMode(mode *string) {
	o.Mode = mode
}

// WithName adds the name to the port collection get params
func (o *PortCollectionGetParams) WithName(name *string) *PortCollectionGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the port collection get params
func (o *PortCollectionGetParams) SetName(name *string) {
	o.Name = name
}

// WithNodeName adds the nodeName to the port collection get params
func (o *PortCollectionGetParams) WithNodeName(nodeName *string) *PortCollectionGetParams {
	o.SetNodeName(nodeName)
	return o
}

// SetNodeName adds the nodeName to the port collection get params
func (o *PortCollectionGetParams) SetNodeName(nodeName *string) {
	o.NodeName = nodeName
}

// WithNodeUUID adds the nodeUUID to the port collection get params
func (o *PortCollectionGetParams) WithNodeUUID(nodeUUID *string) *PortCollectionGetParams {
	o.SetNodeUUID(nodeUUID)
	return o
}

// SetNodeUUID adds the nodeUuid to the port collection get params
func (o *PortCollectionGetParams) SetNodeUUID(nodeUUID *string) {
	o.NodeUUID = nodeUUID
}

// WithOrderBy adds the orderBy to the port collection get params
func (o *PortCollectionGetParams) WithOrderBy(orderBy []string) *PortCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the port collection get params
func (o *PortCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithPartNumber adds the partNumber to the port collection get params
func (o *PortCollectionGetParams) WithPartNumber(partNumber *string) *PortCollectionGetParams {
	o.SetPartNumber(partNumber)
	return o
}

// SetPartNumber adds the partNumber to the port collection get params
func (o *PortCollectionGetParams) SetPartNumber(partNumber *string) {
	o.PartNumber = partNumber
}

// WithRedundant adds the redundant to the port collection get params
func (o *PortCollectionGetParams) WithRedundant(redundant *bool) *PortCollectionGetParams {
	o.SetRedundant(redundant)
	return o
}

// SetRedundant adds the redundant to the port collection get params
func (o *PortCollectionGetParams) SetRedundant(redundant *bool) {
	o.Redundant = redundant
}

// WithReturnRecords adds the returnRecords to the port collection get params
func (o *PortCollectionGetParams) WithReturnRecords(returnRecords *bool) *PortCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the port collection get params
func (o *PortCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the port collection get params
func (o *PortCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *PortCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the port collection get params
func (o *PortCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialNumber adds the serialNumber to the port collection get params
func (o *PortCollectionGetParams) WithSerialNumber(serialNumber *string) *PortCollectionGetParams {
	o.SetSerialNumber(serialNumber)
	return o
}

// SetSerialNumber adds the serialNumber to the port collection get params
func (o *PortCollectionGetParams) SetSerialNumber(serialNumber *string) {
	o.SerialNumber = serialNumber
}

// WithSpeed adds the speed to the port collection get params
func (o *PortCollectionGetParams) WithSpeed(speed *float64) *PortCollectionGetParams {
	o.SetSpeed(speed)
	return o
}

// SetSpeed adds the speed to the port collection get params
func (o *PortCollectionGetParams) SetSpeed(speed *float64) {
	o.Speed = speed
}

// WithState adds the state to the port collection get params
func (o *PortCollectionGetParams) WithState(state *string) *PortCollectionGetParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the port collection get params
func (o *PortCollectionGetParams) SetState(state *string) {
	o.State = state
}

// WithType adds the typeVar to the port collection get params
func (o *PortCollectionGetParams) WithType(typeVar *string) *PortCollectionGetParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the port collection get params
func (o *PortCollectionGetParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithWwn adds the wwn to the port collection get params
func (o *PortCollectionGetParams) WithWwn(wwn *string) *PortCollectionGetParams {
	o.SetWwn(wwn)
	return o
}

// SetWwn adds the wwn to the port collection get params
func (o *PortCollectionGetParams) SetWwn(wwn *string) {
	o.Wwn = wwn
}

// WithWwpn adds the wwpn to the port collection get params
func (o *PortCollectionGetParams) WithWwpn(wwpn *string) *PortCollectionGetParams {
	o.SetWwpn(wwpn)
	return o
}

// SetWwpn adds the wwpn to the port collection get params
func (o *PortCollectionGetParams) SetWwpn(wwpn *string) {
	o.Wwpn = wwpn
}

// WriteToRequest writes these params to a swagger request
func (o *PortCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BoardName != nil {

		// query param board_name
		var qrBoardName string

		if o.BoardName != nil {
			qrBoardName = *o.BoardName
		}
		qBoardName := qrBoardName
		if qBoardName != "" {

			if err := r.SetQueryParam("board_name", qBoardName); err != nil {
				return err
			}
		}
	}

	if o.CableIdentifier != nil {

		// query param cable.identifier
		var qrCableIdentifier string

		if o.CableIdentifier != nil {
			qrCableIdentifier = *o.CableIdentifier
		}
		qCableIdentifier := qrCableIdentifier
		if qCableIdentifier != "" {

			if err := r.SetQueryParam("cable.identifier", qCableIdentifier); err != nil {
				return err
			}
		}
	}

	if o.CableLength != nil {

		// query param cable.length
		var qrCableLength string

		if o.CableLength != nil {
			qrCableLength = *o.CableLength
		}
		qCableLength := qrCableLength
		if qCableLength != "" {

			if err := r.SetQueryParam("cable.length", qCableLength); err != nil {
				return err
			}
		}
	}

	if o.CablePartNumber != nil {

		// query param cable.part_number
		var qrCablePartNumber string

		if o.CablePartNumber != nil {
			qrCablePartNumber = *o.CablePartNumber
		}
		qCablePartNumber := qrCablePartNumber
		if qCablePartNumber != "" {

			if err := r.SetQueryParam("cable.part_number", qCablePartNumber); err != nil {
				return err
			}
		}
	}

	if o.CableSerialNumber != nil {

		// query param cable.serial_number
		var qrCableSerialNumber string

		if o.CableSerialNumber != nil {
			qrCableSerialNumber = *o.CableSerialNumber
		}
		qCableSerialNumber := qrCableSerialNumber
		if qCableSerialNumber != "" {

			if err := r.SetQueryParam("cable.serial_number", qCableSerialNumber); err != nil {
				return err
			}
		}
	}

	if o.Description != nil {

		// query param description
		var qrDescription string

		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}
	}

	if o.Enabled != nil {

		// query param enabled
		var qrEnabled bool

		if o.Enabled != nil {
			qrEnabled = *o.Enabled
		}
		qEnabled := swag.FormatBool(qrEnabled)
		if qEnabled != "" {

			if err := r.SetQueryParam("enabled", qEnabled); err != nil {
				return err
			}
		}
	}

	if o.ErrorCorrectiveAction != nil {

		// query param error.corrective_action
		var qrErrorCorrectiveAction string

		if o.ErrorCorrectiveAction != nil {
			qrErrorCorrectiveAction = *o.ErrorCorrectiveAction
		}
		qErrorCorrectiveAction := qrErrorCorrectiveAction
		if qErrorCorrectiveAction != "" {

			if err := r.SetQueryParam("error.corrective_action", qErrorCorrectiveAction); err != nil {
				return err
			}
		}
	}

	if o.ErrorMessage != nil {

		// query param error.message
		var qrErrorMessage string

		if o.ErrorMessage != nil {
			qrErrorMessage = *o.ErrorMessage
		}
		qErrorMessage := qrErrorMessage
		if qErrorMessage != "" {

			if err := r.SetQueryParam("error.message", qErrorMessage); err != nil {
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

	if o.FirmwareVersion != nil {

		// query param firmware_version
		var qrFirmwareVersion string

		if o.FirmwareVersion != nil {
			qrFirmwareVersion = *o.FirmwareVersion
		}
		qFirmwareVersion := qrFirmwareVersion
		if qFirmwareVersion != "" {

			if err := r.SetQueryParam("firmware_version", qFirmwareVersion); err != nil {
				return err
			}
		}
	}

	if o.InUse != nil {

		// query param in_use
		var qrInUse bool

		if o.InUse != nil {
			qrInUse = *o.InUse
		}
		qInUse := swag.FormatBool(qrInUse)
		if qInUse != "" {

			if err := r.SetQueryParam("in_use", qInUse); err != nil {
				return err
			}
		}
	}

	if o.MacAddress != nil {

		// query param mac_address
		var qrMacAddress string

		if o.MacAddress != nil {
			qrMacAddress = *o.MacAddress
		}
		qMacAddress := qrMacAddress
		if qMacAddress != "" {

			if err := r.SetQueryParam("mac_address", qMacAddress); err != nil {
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

	if o.Mode != nil {

		// query param mode
		var qrMode string

		if o.Mode != nil {
			qrMode = *o.Mode
		}
		qMode := qrMode
		if qMode != "" {

			if err := r.SetQueryParam("mode", qMode); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.NodeName != nil {

		// query param node.name
		var qrNodeName string

		if o.NodeName != nil {
			qrNodeName = *o.NodeName
		}
		qNodeName := qrNodeName
		if qNodeName != "" {

			if err := r.SetQueryParam("node.name", qNodeName); err != nil {
				return err
			}
		}
	}

	if o.NodeUUID != nil {

		// query param node.uuid
		var qrNodeUUID string

		if o.NodeUUID != nil {
			qrNodeUUID = *o.NodeUUID
		}
		qNodeUUID := qrNodeUUID
		if qNodeUUID != "" {

			if err := r.SetQueryParam("node.uuid", qNodeUUID); err != nil {
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

	if o.PartNumber != nil {

		// query param part_number
		var qrPartNumber string

		if o.PartNumber != nil {
			qrPartNumber = *o.PartNumber
		}
		qPartNumber := qrPartNumber
		if qPartNumber != "" {

			if err := r.SetQueryParam("part_number", qPartNumber); err != nil {
				return err
			}
		}
	}

	if o.Redundant != nil {

		// query param redundant
		var qrRedundant bool

		if o.Redundant != nil {
			qrRedundant = *o.Redundant
		}
		qRedundant := swag.FormatBool(qrRedundant)
		if qRedundant != "" {

			if err := r.SetQueryParam("redundant", qRedundant); err != nil {
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

	if o.SerialNumber != nil {

		// query param serial_number
		var qrSerialNumber string

		if o.SerialNumber != nil {
			qrSerialNumber = *o.SerialNumber
		}
		qSerialNumber := qrSerialNumber
		if qSerialNumber != "" {

			if err := r.SetQueryParam("serial_number", qSerialNumber); err != nil {
				return err
			}
		}
	}

	if o.Speed != nil {

		// query param speed
		var qrSpeed float64

		if o.Speed != nil {
			qrSpeed = *o.Speed
		}
		qSpeed := swag.FormatFloat64(qrSpeed)
		if qSpeed != "" {

			if err := r.SetQueryParam("speed", qSpeed); err != nil {
				return err
			}
		}
	}

	if o.State != nil {

		// query param state
		var qrState string

		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {

			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}
	}

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if o.Wwn != nil {

		// query param wwn
		var qrWwn string

		if o.Wwn != nil {
			qrWwn = *o.Wwn
		}
		qWwn := qrWwn
		if qWwn != "" {

			if err := r.SetQueryParam("wwn", qWwn); err != nil {
				return err
			}
		}
	}

	if o.Wwpn != nil {

		// query param wwpn
		var qrWwpn string

		if o.Wwpn != nil {
			qrWwpn = *o.Wwpn
		}
		qWwpn := qrWwpn
		if qWwpn != "" {

			if err := r.SetQueryParam("wwpn", qWwpn); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamPortCollectionGet binds the parameter fields
func (o *PortCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamPortCollectionGet binds the parameter order_by
func (o *PortCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
