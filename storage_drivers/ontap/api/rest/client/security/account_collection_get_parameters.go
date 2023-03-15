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
)

// NewAccountCollectionGetParams creates a new AccountCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccountCollectionGetParams() *AccountCollectionGetParams {
	return &AccountCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccountCollectionGetParamsWithTimeout creates a new AccountCollectionGetParams object
// with the ability to set a timeout on a request.
func NewAccountCollectionGetParamsWithTimeout(timeout time.Duration) *AccountCollectionGetParams {
	return &AccountCollectionGetParams{
		timeout: timeout,
	}
}

// NewAccountCollectionGetParamsWithContext creates a new AccountCollectionGetParams object
// with the ability to set a context for a request.
func NewAccountCollectionGetParamsWithContext(ctx context.Context) *AccountCollectionGetParams {
	return &AccountCollectionGetParams{
		Context: ctx,
	}
}

// NewAccountCollectionGetParamsWithHTTPClient creates a new AccountCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccountCollectionGetParamsWithHTTPClient(client *http.Client) *AccountCollectionGetParams {
	return &AccountCollectionGetParams{
		HTTPClient: client,
	}
}

/*
AccountCollectionGetParams contains all the parameters to send to the API endpoint

	for the account collection get operation.

	Typically these are written to a http.Request.
*/
type AccountCollectionGetParams struct {

	/* ApplicationsApplication.

	   Filter by applications.application
	*/
	ApplicationsApplication *string

	/* ApplicationsAuthenticationMethods.

	   Filter by applications.authentication_methods
	*/
	ApplicationsAuthenticationMethods *string

	/* ApplicationsSecondAuthenticationMethod.

	   Filter by applications.second_authentication_method
	*/
	ApplicationsSecondAuthenticationMethod *string

	/* Comment.

	   Filter by comment
	*/
	Comment *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* LdapFastbind.

	   Filter by ldap_fastbind
	*/
	LdapFastbind *bool

	/* Locked.

	   Filter by locked
	*/
	Locked *bool

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* Name.

	   Filter by name
	*/
	Name *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* OwnerName.

	   Filter by owner.name
	*/
	OwnerName *string

	/* OwnerUUID.

	   Filter by owner.uuid
	*/
	OwnerUUID *string

	/* PasswordHashAlgorithm.

	   Filter by password_hash_algorithm
	*/
	PasswordHashAlgorithm *string

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

	/* RoleName.

	   Filter by role.name
	*/
	RoleName *string

	/* Scope.

	   Filter by scope
	*/
	Scope *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the account collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountCollectionGetParams) WithDefaults() *AccountCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the account collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := AccountCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the account collection get params
func (o *AccountCollectionGetParams) WithTimeout(timeout time.Duration) *AccountCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account collection get params
func (o *AccountCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account collection get params
func (o *AccountCollectionGetParams) WithContext(ctx context.Context) *AccountCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account collection get params
func (o *AccountCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account collection get params
func (o *AccountCollectionGetParams) WithHTTPClient(client *http.Client) *AccountCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account collection get params
func (o *AccountCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationsApplication adds the applicationsApplication to the account collection get params
func (o *AccountCollectionGetParams) WithApplicationsApplication(applicationsApplication *string) *AccountCollectionGetParams {
	o.SetApplicationsApplication(applicationsApplication)
	return o
}

// SetApplicationsApplication adds the applicationsApplication to the account collection get params
func (o *AccountCollectionGetParams) SetApplicationsApplication(applicationsApplication *string) {
	o.ApplicationsApplication = applicationsApplication
}

// WithApplicationsAuthenticationMethods adds the applicationsAuthenticationMethods to the account collection get params
func (o *AccountCollectionGetParams) WithApplicationsAuthenticationMethods(applicationsAuthenticationMethods *string) *AccountCollectionGetParams {
	o.SetApplicationsAuthenticationMethods(applicationsAuthenticationMethods)
	return o
}

// SetApplicationsAuthenticationMethods adds the applicationsAuthenticationMethods to the account collection get params
func (o *AccountCollectionGetParams) SetApplicationsAuthenticationMethods(applicationsAuthenticationMethods *string) {
	o.ApplicationsAuthenticationMethods = applicationsAuthenticationMethods
}

// WithApplicationsSecondAuthenticationMethod adds the applicationsSecondAuthenticationMethod to the account collection get params
func (o *AccountCollectionGetParams) WithApplicationsSecondAuthenticationMethod(applicationsSecondAuthenticationMethod *string) *AccountCollectionGetParams {
	o.SetApplicationsSecondAuthenticationMethod(applicationsSecondAuthenticationMethod)
	return o
}

// SetApplicationsSecondAuthenticationMethod adds the applicationsSecondAuthenticationMethod to the account collection get params
func (o *AccountCollectionGetParams) SetApplicationsSecondAuthenticationMethod(applicationsSecondAuthenticationMethod *string) {
	o.ApplicationsSecondAuthenticationMethod = applicationsSecondAuthenticationMethod
}

// WithComment adds the comment to the account collection get params
func (o *AccountCollectionGetParams) WithComment(comment *string) *AccountCollectionGetParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the account collection get params
func (o *AccountCollectionGetParams) SetComment(comment *string) {
	o.Comment = comment
}

// WithFields adds the fields to the account collection get params
func (o *AccountCollectionGetParams) WithFields(fields []string) *AccountCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the account collection get params
func (o *AccountCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithLdapFastbind adds the ldapFastbind to the account collection get params
func (o *AccountCollectionGetParams) WithLdapFastbind(ldapFastbind *bool) *AccountCollectionGetParams {
	o.SetLdapFastbind(ldapFastbind)
	return o
}

// SetLdapFastbind adds the ldapFastbind to the account collection get params
func (o *AccountCollectionGetParams) SetLdapFastbind(ldapFastbind *bool) {
	o.LdapFastbind = ldapFastbind
}

// WithLocked adds the locked to the account collection get params
func (o *AccountCollectionGetParams) WithLocked(locked *bool) *AccountCollectionGetParams {
	o.SetLocked(locked)
	return o
}

// SetLocked adds the locked to the account collection get params
func (o *AccountCollectionGetParams) SetLocked(locked *bool) {
	o.Locked = locked
}

// WithMaxRecords adds the maxRecords to the account collection get params
func (o *AccountCollectionGetParams) WithMaxRecords(maxRecords *int64) *AccountCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the account collection get params
func (o *AccountCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithName adds the name to the account collection get params
func (o *AccountCollectionGetParams) WithName(name *string) *AccountCollectionGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the account collection get params
func (o *AccountCollectionGetParams) SetName(name *string) {
	o.Name = name
}

// WithOrderBy adds the orderBy to the account collection get params
func (o *AccountCollectionGetParams) WithOrderBy(orderBy []string) *AccountCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the account collection get params
func (o *AccountCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithOwnerName adds the ownerName to the account collection get params
func (o *AccountCollectionGetParams) WithOwnerName(ownerName *string) *AccountCollectionGetParams {
	o.SetOwnerName(ownerName)
	return o
}

// SetOwnerName adds the ownerName to the account collection get params
func (o *AccountCollectionGetParams) SetOwnerName(ownerName *string) {
	o.OwnerName = ownerName
}

// WithOwnerUUID adds the ownerUUID to the account collection get params
func (o *AccountCollectionGetParams) WithOwnerUUID(ownerUUID *string) *AccountCollectionGetParams {
	o.SetOwnerUUID(ownerUUID)
	return o
}

// SetOwnerUUID adds the ownerUuid to the account collection get params
func (o *AccountCollectionGetParams) SetOwnerUUID(ownerUUID *string) {
	o.OwnerUUID = ownerUUID
}

// WithPasswordHashAlgorithm adds the passwordHashAlgorithm to the account collection get params
func (o *AccountCollectionGetParams) WithPasswordHashAlgorithm(passwordHashAlgorithm *string) *AccountCollectionGetParams {
	o.SetPasswordHashAlgorithm(passwordHashAlgorithm)
	return o
}

// SetPasswordHashAlgorithm adds the passwordHashAlgorithm to the account collection get params
func (o *AccountCollectionGetParams) SetPasswordHashAlgorithm(passwordHashAlgorithm *string) {
	o.PasswordHashAlgorithm = passwordHashAlgorithm
}

// WithReturnRecords adds the returnRecords to the account collection get params
func (o *AccountCollectionGetParams) WithReturnRecords(returnRecords *bool) *AccountCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the account collection get params
func (o *AccountCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the account collection get params
func (o *AccountCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *AccountCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the account collection get params
func (o *AccountCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithRoleName adds the roleName to the account collection get params
func (o *AccountCollectionGetParams) WithRoleName(roleName *string) *AccountCollectionGetParams {
	o.SetRoleName(roleName)
	return o
}

// SetRoleName adds the roleName to the account collection get params
func (o *AccountCollectionGetParams) SetRoleName(roleName *string) {
	o.RoleName = roleName
}

// WithScope adds the scope to the account collection get params
func (o *AccountCollectionGetParams) WithScope(scope *string) *AccountCollectionGetParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the account collection get params
func (o *AccountCollectionGetParams) SetScope(scope *string) {
	o.Scope = scope
}

// WriteToRequest writes these params to a swagger request
func (o *AccountCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ApplicationsApplication != nil {

		// query param applications.application
		var qrApplicationsApplication string

		if o.ApplicationsApplication != nil {
			qrApplicationsApplication = *o.ApplicationsApplication
		}
		qApplicationsApplication := qrApplicationsApplication
		if qApplicationsApplication != "" {

			if err := r.SetQueryParam("applications.application", qApplicationsApplication); err != nil {
				return err
			}
		}
	}

	if o.ApplicationsAuthenticationMethods != nil {

		// query param applications.authentication_methods
		var qrApplicationsAuthenticationMethods string

		if o.ApplicationsAuthenticationMethods != nil {
			qrApplicationsAuthenticationMethods = *o.ApplicationsAuthenticationMethods
		}
		qApplicationsAuthenticationMethods := qrApplicationsAuthenticationMethods
		if qApplicationsAuthenticationMethods != "" {

			if err := r.SetQueryParam("applications.authentication_methods", qApplicationsAuthenticationMethods); err != nil {
				return err
			}
		}
	}

	if o.ApplicationsSecondAuthenticationMethod != nil {

		// query param applications.second_authentication_method
		var qrApplicationsSecondAuthenticationMethod string

		if o.ApplicationsSecondAuthenticationMethod != nil {
			qrApplicationsSecondAuthenticationMethod = *o.ApplicationsSecondAuthenticationMethod
		}
		qApplicationsSecondAuthenticationMethod := qrApplicationsSecondAuthenticationMethod
		if qApplicationsSecondAuthenticationMethod != "" {

			if err := r.SetQueryParam("applications.second_authentication_method", qApplicationsSecondAuthenticationMethod); err != nil {
				return err
			}
		}
	}

	if o.Comment != nil {

		// query param comment
		var qrComment string

		if o.Comment != nil {
			qrComment = *o.Comment
		}
		qComment := qrComment
		if qComment != "" {

			if err := r.SetQueryParam("comment", qComment); err != nil {
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

	if o.LdapFastbind != nil {

		// query param ldap_fastbind
		var qrLdapFastbind bool

		if o.LdapFastbind != nil {
			qrLdapFastbind = *o.LdapFastbind
		}
		qLdapFastbind := swag.FormatBool(qrLdapFastbind)
		if qLdapFastbind != "" {

			if err := r.SetQueryParam("ldap_fastbind", qLdapFastbind); err != nil {
				return err
			}
		}
	}

	if o.Locked != nil {

		// query param locked
		var qrLocked bool

		if o.Locked != nil {
			qrLocked = *o.Locked
		}
		qLocked := swag.FormatBool(qrLocked)
		if qLocked != "" {

			if err := r.SetQueryParam("locked", qLocked); err != nil {
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

	if o.OrderBy != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.OwnerName != nil {

		// query param owner.name
		var qrOwnerName string

		if o.OwnerName != nil {
			qrOwnerName = *o.OwnerName
		}
		qOwnerName := qrOwnerName
		if qOwnerName != "" {

			if err := r.SetQueryParam("owner.name", qOwnerName); err != nil {
				return err
			}
		}
	}

	if o.OwnerUUID != nil {

		// query param owner.uuid
		var qrOwnerUUID string

		if o.OwnerUUID != nil {
			qrOwnerUUID = *o.OwnerUUID
		}
		qOwnerUUID := qrOwnerUUID
		if qOwnerUUID != "" {

			if err := r.SetQueryParam("owner.uuid", qOwnerUUID); err != nil {
				return err
			}
		}
	}

	if o.PasswordHashAlgorithm != nil {

		// query param password_hash_algorithm
		var qrPasswordHashAlgorithm string

		if o.PasswordHashAlgorithm != nil {
			qrPasswordHashAlgorithm = *o.PasswordHashAlgorithm
		}
		qPasswordHashAlgorithm := qrPasswordHashAlgorithm
		if qPasswordHashAlgorithm != "" {

			if err := r.SetQueryParam("password_hash_algorithm", qPasswordHashAlgorithm); err != nil {
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

	if o.RoleName != nil {

		// query param role.name
		var qrRoleName string

		if o.RoleName != nil {
			qrRoleName = *o.RoleName
		}
		qRoleName := qrRoleName
		if qRoleName != "" {

			if err := r.SetQueryParam("role.name", qRoleName); err != nil {
				return err
			}
		}
	}

	if o.Scope != nil {

		// query param scope
		var qrScope string

		if o.Scope != nil {
			qrScope = *o.Scope
		}
		qScope := qrScope
		if qScope != "" {

			if err := r.SetQueryParam("scope", qScope); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamAccountCollectionGet binds the parameter fields
func (o *AccountCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamAccountCollectionGet binds the parameter order_by
func (o *AccountCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
