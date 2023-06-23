// Code generated by go-swagger; DO NOT EDIT.

package user_management

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

// NewEntitiesRolesV1Params creates a new EntitiesRolesV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEntitiesRolesV1Params() *EntitiesRolesV1Params {
	return &EntitiesRolesV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewEntitiesRolesV1ParamsWithTimeout creates a new EntitiesRolesV1Params object
// with the ability to set a timeout on a request.
func NewEntitiesRolesV1ParamsWithTimeout(timeout time.Duration) *EntitiesRolesV1Params {
	return &EntitiesRolesV1Params{
		timeout: timeout,
	}
}

// NewEntitiesRolesV1ParamsWithContext creates a new EntitiesRolesV1Params object
// with the ability to set a context for a request.
func NewEntitiesRolesV1ParamsWithContext(ctx context.Context) *EntitiesRolesV1Params {
	return &EntitiesRolesV1Params{
		Context: ctx,
	}
}

// NewEntitiesRolesV1ParamsWithHTTPClient creates a new EntitiesRolesV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewEntitiesRolesV1ParamsWithHTTPClient(client *http.Client) *EntitiesRolesV1Params {
	return &EntitiesRolesV1Params{
		HTTPClient: client,
	}
}

/*
EntitiesRolesV1Params contains all the parameters to send to the API endpoint

	for the entities roles v1 operation.

	Typically these are written to a http.Request.
*/
type EntitiesRolesV1Params struct {

	/* CID.

	   Customer ID to get available roles for. Empty CID would result in Role IDs for current CID in view.
	*/
	CID *string

	/* Ids.

	   ID of a role. Find a role ID from `/user-management/queries/roles/v1`.
	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the entities roles v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EntitiesRolesV1Params) WithDefaults() *EntitiesRolesV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the entities roles v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EntitiesRolesV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the entities roles v1 params
func (o *EntitiesRolesV1Params) WithTimeout(timeout time.Duration) *EntitiesRolesV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the entities roles v1 params
func (o *EntitiesRolesV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the entities roles v1 params
func (o *EntitiesRolesV1Params) WithContext(ctx context.Context) *EntitiesRolesV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the entities roles v1 params
func (o *EntitiesRolesV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the entities roles v1 params
func (o *EntitiesRolesV1Params) WithHTTPClient(client *http.Client) *EntitiesRolesV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the entities roles v1 params
func (o *EntitiesRolesV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCID adds the cid to the entities roles v1 params
func (o *EntitiesRolesV1Params) WithCID(cid *string) *EntitiesRolesV1Params {
	o.SetCID(cid)
	return o
}

// SetCID adds the cid to the entities roles v1 params
func (o *EntitiesRolesV1Params) SetCID(cid *string) {
	o.CID = cid
}

// WithIds adds the ids to the entities roles v1 params
func (o *EntitiesRolesV1Params) WithIds(ids []string) *EntitiesRolesV1Params {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the entities roles v1 params
func (o *EntitiesRolesV1Params) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *EntitiesRolesV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CID != nil {

		// query param cid
		var qrCID string

		if o.CID != nil {
			qrCID = *o.CID
		}
		qCID := qrCID
		if qCID != "" {

			if err := r.SetQueryParam("cid", qCID); err != nil {
				return err
			}
		}
	}

	if o.Ids != nil {

		// binding items for ids
		joinedIds := o.bindParamIds(reg)

		// query array param ids
		if err := r.SetQueryParam("ids", joinedIds...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamEntitiesRolesV1 binds the parameter ids
func (o *EntitiesRolesV1Params) bindParamIds(formats strfmt.Registry) []string {
	idsIR := o.Ids

	var idsIC []string
	for _, idsIIR := range idsIR { // explode []string

		idsIIV := idsIIR // string as string
		idsIC = append(idsIC, idsIIV)
	}

	// items.CollectionFormat: "multi"
	idsIS := swag.JoinByFormat(idsIC, "multi")

	return idsIS
}
