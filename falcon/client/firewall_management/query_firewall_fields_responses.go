// Code generated by go-swagger; DO NOT EDIT.

package firewall_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// QueryFirewallFieldsReader is a Reader for the QueryFirewallFields structure.
type QueryFirewallFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryFirewallFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryFirewallFieldsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewQueryFirewallFieldsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryFirewallFieldsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /fwmgr/queries/firewall-fields/v1] query-firewall-fields", response, response.Code())
	}
}

// NewQueryFirewallFieldsOK creates a QueryFirewallFieldsOK with default headers values
func NewQueryFirewallFieldsOK() *QueryFirewallFieldsOK {
	return &QueryFirewallFieldsOK{}
}

/*
QueryFirewallFieldsOK describes a response with status code 200, with default header values.

OK
*/
type QueryFirewallFieldsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.FwmgrMsaspecQueryResponse
}

// IsSuccess returns true when this query firewall fields o k response has a 2xx status code
func (o *QueryFirewallFieldsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query firewall fields o k response has a 3xx status code
func (o *QueryFirewallFieldsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query firewall fields o k response has a 4xx status code
func (o *QueryFirewallFieldsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query firewall fields o k response has a 5xx status code
func (o *QueryFirewallFieldsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query firewall fields o k response a status code equal to that given
func (o *QueryFirewallFieldsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query firewall fields o k response
func (o *QueryFirewallFieldsOK) Code() int {
	return 200
}

func (o *QueryFirewallFieldsOK) Error() string {
	return fmt.Sprintf("[GET /fwmgr/queries/firewall-fields/v1][%d] queryFirewallFieldsOK  %+v", 200, o.Payload)
}

func (o *QueryFirewallFieldsOK) String() string {
	return fmt.Sprintf("[GET /fwmgr/queries/firewall-fields/v1][%d] queryFirewallFieldsOK  %+v", 200, o.Payload)
}

func (o *QueryFirewallFieldsOK) GetPayload() *models.FwmgrMsaspecQueryResponse {
	return o.Payload
}

func (o *QueryFirewallFieldsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.FwmgrMsaspecQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryFirewallFieldsForbidden creates a QueryFirewallFieldsForbidden with default headers values
func NewQueryFirewallFieldsForbidden() *QueryFirewallFieldsForbidden {
	return &QueryFirewallFieldsForbidden{}
}

/*
QueryFirewallFieldsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryFirewallFieldsForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this query firewall fields forbidden response has a 2xx status code
func (o *QueryFirewallFieldsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query firewall fields forbidden response has a 3xx status code
func (o *QueryFirewallFieldsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query firewall fields forbidden response has a 4xx status code
func (o *QueryFirewallFieldsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query firewall fields forbidden response has a 5xx status code
func (o *QueryFirewallFieldsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query firewall fields forbidden response a status code equal to that given
func (o *QueryFirewallFieldsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query firewall fields forbidden response
func (o *QueryFirewallFieldsForbidden) Code() int {
	return 403
}

func (o *QueryFirewallFieldsForbidden) Error() string {
	return fmt.Sprintf("[GET /fwmgr/queries/firewall-fields/v1][%d] queryFirewallFieldsForbidden  %+v", 403, o.Payload)
}

func (o *QueryFirewallFieldsForbidden) String() string {
	return fmt.Sprintf("[GET /fwmgr/queries/firewall-fields/v1][%d] queryFirewallFieldsForbidden  %+v", 403, o.Payload)
}

func (o *QueryFirewallFieldsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryFirewallFieldsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryFirewallFieldsTooManyRequests creates a QueryFirewallFieldsTooManyRequests with default headers values
func NewQueryFirewallFieldsTooManyRequests() *QueryFirewallFieldsTooManyRequests {
	return &QueryFirewallFieldsTooManyRequests{}
}

/*
QueryFirewallFieldsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryFirewallFieldsTooManyRequests struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this query firewall fields too many requests response has a 2xx status code
func (o *QueryFirewallFieldsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query firewall fields too many requests response has a 3xx status code
func (o *QueryFirewallFieldsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query firewall fields too many requests response has a 4xx status code
func (o *QueryFirewallFieldsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query firewall fields too many requests response has a 5xx status code
func (o *QueryFirewallFieldsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query firewall fields too many requests response a status code equal to that given
func (o *QueryFirewallFieldsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query firewall fields too many requests response
func (o *QueryFirewallFieldsTooManyRequests) Code() int {
	return 429
}

func (o *QueryFirewallFieldsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /fwmgr/queries/firewall-fields/v1][%d] queryFirewallFieldsTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryFirewallFieldsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /fwmgr/queries/firewall-fields/v1][%d] queryFirewallFieldsTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryFirewallFieldsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryFirewallFieldsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
