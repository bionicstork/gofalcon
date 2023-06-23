// Code generated by go-swagger; DO NOT EDIT.

package mssp

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

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// NewCreateUserGroupsParams creates a new CreateUserGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateUserGroupsParams() *CreateUserGroupsParams {
	return &CreateUserGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUserGroupsParamsWithTimeout creates a new CreateUserGroupsParams object
// with the ability to set a timeout on a request.
func NewCreateUserGroupsParamsWithTimeout(timeout time.Duration) *CreateUserGroupsParams {
	return &CreateUserGroupsParams{
		timeout: timeout,
	}
}

// NewCreateUserGroupsParamsWithContext creates a new CreateUserGroupsParams object
// with the ability to set a context for a request.
func NewCreateUserGroupsParamsWithContext(ctx context.Context) *CreateUserGroupsParams {
	return &CreateUserGroupsParams{
		Context: ctx,
	}
}

// NewCreateUserGroupsParamsWithHTTPClient creates a new CreateUserGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateUserGroupsParamsWithHTTPClient(client *http.Client) *CreateUserGroupsParams {
	return &CreateUserGroupsParams{
		HTTPClient: client,
	}
}

/*
CreateUserGroupsParams contains all the parameters to send to the API endpoint

	for the create user groups operation.

	Typically these are written to a http.Request.
*/
type CreateUserGroupsParams struct {

	/* Body.

	   Only 'name' and/or 'description' fields are required. Remaining are assigned by the system.
	*/
	Body *models.DomainUserGroupsRequestV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create user groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUserGroupsParams) WithDefaults() *CreateUserGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create user groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUserGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create user groups params
func (o *CreateUserGroupsParams) WithTimeout(timeout time.Duration) *CreateUserGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create user groups params
func (o *CreateUserGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create user groups params
func (o *CreateUserGroupsParams) WithContext(ctx context.Context) *CreateUserGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create user groups params
func (o *CreateUserGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create user groups params
func (o *CreateUserGroupsParams) WithHTTPClient(client *http.Client) *CreateUserGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create user groups params
func (o *CreateUserGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create user groups params
func (o *CreateUserGroupsParams) WithBody(body *models.DomainUserGroupsRequestV1) *CreateUserGroupsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create user groups params
func (o *CreateUserGroupsParams) SetBody(body *models.DomainUserGroupsRequestV1) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUserGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
