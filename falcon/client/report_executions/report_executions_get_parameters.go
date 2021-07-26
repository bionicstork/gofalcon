// Code generated by go-swagger; DO NOT EDIT.

package report_executions

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

// NewReportExecutionsGetParams creates a new ReportExecutionsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReportExecutionsGetParams() *ReportExecutionsGetParams {
	return &ReportExecutionsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReportExecutionsGetParamsWithTimeout creates a new ReportExecutionsGetParams object
// with the ability to set a timeout on a request.
func NewReportExecutionsGetParamsWithTimeout(timeout time.Duration) *ReportExecutionsGetParams {
	return &ReportExecutionsGetParams{
		timeout: timeout,
	}
}

// NewReportExecutionsGetParamsWithContext creates a new ReportExecutionsGetParams object
// with the ability to set a context for a request.
func NewReportExecutionsGetParamsWithContext(ctx context.Context) *ReportExecutionsGetParams {
	return &ReportExecutionsGetParams{
		Context: ctx,
	}
}

// NewReportExecutionsGetParamsWithHTTPClient creates a new ReportExecutionsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewReportExecutionsGetParamsWithHTTPClient(client *http.Client) *ReportExecutionsGetParams {
	return &ReportExecutionsGetParams{
		HTTPClient: client,
	}
}

/* ReportExecutionsGetParams contains all the parameters to send to the API endpoint
   for the report executions get operation.

   Typically these are written to a http.Request.
*/
type ReportExecutionsGetParams struct {

	/* XCSUSERID.

	   The user id (not used with API client)
	*/
	XCSUSERID *string

	/* XCSUSERUUID.

	   The user uuid (not used with API client)
	*/
	XCSUSERUUID *string

	/* Ids.

	   The report_execution id to get details about.
	*/
	Ids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the report executions get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReportExecutionsGetParams) WithDefaults() *ReportExecutionsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the report executions get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReportExecutionsGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the report executions get params
func (o *ReportExecutionsGetParams) WithTimeout(timeout time.Duration) *ReportExecutionsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the report executions get params
func (o *ReportExecutionsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the report executions get params
func (o *ReportExecutionsGetParams) WithContext(ctx context.Context) *ReportExecutionsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the report executions get params
func (o *ReportExecutionsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the report executions get params
func (o *ReportExecutionsGetParams) WithHTTPClient(client *http.Client) *ReportExecutionsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the report executions get params
func (o *ReportExecutionsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXCSUSERID adds the xCSUSERID to the report executions get params
func (o *ReportExecutionsGetParams) WithXCSUSERID(xCSUSERID *string) *ReportExecutionsGetParams {
	o.SetXCSUSERID(xCSUSERID)
	return o
}

// SetXCSUSERID adds the xCSUSERId to the report executions get params
func (o *ReportExecutionsGetParams) SetXCSUSERID(xCSUSERID *string) {
	o.XCSUSERID = xCSUSERID
}

// WithXCSUSERUUID adds the xCSUSERUUID to the report executions get params
func (o *ReportExecutionsGetParams) WithXCSUSERUUID(xCSUSERUUID *string) *ReportExecutionsGetParams {
	o.SetXCSUSERUUID(xCSUSERUUID)
	return o
}

// SetXCSUSERUUID adds the xCSUSERUuid to the report executions get params
func (o *ReportExecutionsGetParams) SetXCSUSERUUID(xCSUSERUUID *string) {
	o.XCSUSERUUID = xCSUSERUUID
}

// WithIds adds the ids to the report executions get params
func (o *ReportExecutionsGetParams) WithIds(ids []string) *ReportExecutionsGetParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the report executions get params
func (o *ReportExecutionsGetParams) SetIds(ids []string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *ReportExecutionsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XCSUSERID != nil {

		// header param X-CS-USERID
		if err := r.SetHeaderParam("X-CS-USERID", *o.XCSUSERID); err != nil {
			return err
		}
	}

	if o.XCSUSERUUID != nil {

		// header param X-CS-USERUUID
		if err := r.SetHeaderParam("X-CS-USERUUID", *o.XCSUSERUUID); err != nil {
			return err
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

// bindParamReportExecutionsGet binds the parameter ids
func (o *ReportExecutionsGetParams) bindParamIds(formats strfmt.Registry) []string {
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