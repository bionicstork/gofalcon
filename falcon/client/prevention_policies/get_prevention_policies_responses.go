// Code generated by go-swagger; DO NOT EDIT.

package prevention_policies

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

// GetPreventionPoliciesReader is a Reader for the GetPreventionPolicies structure.
type GetPreventionPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPreventionPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPreventionPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetPreventionPoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPreventionPoliciesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPreventionPoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPreventionPoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /policy/entities/prevention/v1] getPreventionPolicies", response, response.Code())
	}
}

// NewGetPreventionPoliciesOK creates a GetPreventionPoliciesOK with default headers values
func NewGetPreventionPoliciesOK() *GetPreventionPoliciesOK {
	return &GetPreventionPoliciesOK{}
}

/*
GetPreventionPoliciesOK describes a response with status code 200, with default header values.

OK
*/
type GetPreventionPoliciesOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.PreventionRespV1
}

// IsSuccess returns true when this get prevention policies o k response has a 2xx status code
func (o *GetPreventionPoliciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get prevention policies o k response has a 3xx status code
func (o *GetPreventionPoliciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get prevention policies o k response has a 4xx status code
func (o *GetPreventionPoliciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get prevention policies o k response has a 5xx status code
func (o *GetPreventionPoliciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get prevention policies o k response a status code equal to that given
func (o *GetPreventionPoliciesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get prevention policies o k response
func (o *GetPreventionPoliciesOK) Code() int {
	return 200
}

func (o *GetPreventionPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /policy/entities/prevention/v1][%d] getPreventionPoliciesOK  %+v", 200, o.Payload)
}

func (o *GetPreventionPoliciesOK) String() string {
	return fmt.Sprintf("[GET /policy/entities/prevention/v1][%d] getPreventionPoliciesOK  %+v", 200, o.Payload)
}

func (o *GetPreventionPoliciesOK) GetPayload() *models.PreventionRespV1 {
	return o.Payload
}

func (o *GetPreventionPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.PreventionRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPreventionPoliciesForbidden creates a GetPreventionPoliciesForbidden with default headers values
func NewGetPreventionPoliciesForbidden() *GetPreventionPoliciesForbidden {
	return &GetPreventionPoliciesForbidden{}
}

/*
GetPreventionPoliciesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetPreventionPoliciesForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

// IsSuccess returns true when this get prevention policies forbidden response has a 2xx status code
func (o *GetPreventionPoliciesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get prevention policies forbidden response has a 3xx status code
func (o *GetPreventionPoliciesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get prevention policies forbidden response has a 4xx status code
func (o *GetPreventionPoliciesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get prevention policies forbidden response has a 5xx status code
func (o *GetPreventionPoliciesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get prevention policies forbidden response a status code equal to that given
func (o *GetPreventionPoliciesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get prevention policies forbidden response
func (o *GetPreventionPoliciesForbidden) Code() int {
	return 403
}

func (o *GetPreventionPoliciesForbidden) Error() string {
	return fmt.Sprintf("[GET /policy/entities/prevention/v1][%d] getPreventionPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *GetPreventionPoliciesForbidden) String() string {
	return fmt.Sprintf("[GET /policy/entities/prevention/v1][%d] getPreventionPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *GetPreventionPoliciesForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetPreventionPoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPreventionPoliciesNotFound creates a GetPreventionPoliciesNotFound with default headers values
func NewGetPreventionPoliciesNotFound() *GetPreventionPoliciesNotFound {
	return &GetPreventionPoliciesNotFound{}
}

/*
GetPreventionPoliciesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetPreventionPoliciesNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.PreventionRespV1
}

// IsSuccess returns true when this get prevention policies not found response has a 2xx status code
func (o *GetPreventionPoliciesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get prevention policies not found response has a 3xx status code
func (o *GetPreventionPoliciesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get prevention policies not found response has a 4xx status code
func (o *GetPreventionPoliciesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get prevention policies not found response has a 5xx status code
func (o *GetPreventionPoliciesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get prevention policies not found response a status code equal to that given
func (o *GetPreventionPoliciesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get prevention policies not found response
func (o *GetPreventionPoliciesNotFound) Code() int {
	return 404
}

func (o *GetPreventionPoliciesNotFound) Error() string {
	return fmt.Sprintf("[GET /policy/entities/prevention/v1][%d] getPreventionPoliciesNotFound  %+v", 404, o.Payload)
}

func (o *GetPreventionPoliciesNotFound) String() string {
	return fmt.Sprintf("[GET /policy/entities/prevention/v1][%d] getPreventionPoliciesNotFound  %+v", 404, o.Payload)
}

func (o *GetPreventionPoliciesNotFound) GetPayload() *models.PreventionRespV1 {
	return o.Payload
}

func (o *GetPreventionPoliciesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.PreventionRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPreventionPoliciesTooManyRequests creates a GetPreventionPoliciesTooManyRequests with default headers values
func NewGetPreventionPoliciesTooManyRequests() *GetPreventionPoliciesTooManyRequests {
	return &GetPreventionPoliciesTooManyRequests{}
}

/*
GetPreventionPoliciesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetPreventionPoliciesTooManyRequests struct {

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

// IsSuccess returns true when this get prevention policies too many requests response has a 2xx status code
func (o *GetPreventionPoliciesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get prevention policies too many requests response has a 3xx status code
func (o *GetPreventionPoliciesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get prevention policies too many requests response has a 4xx status code
func (o *GetPreventionPoliciesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get prevention policies too many requests response has a 5xx status code
func (o *GetPreventionPoliciesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get prevention policies too many requests response a status code equal to that given
func (o *GetPreventionPoliciesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get prevention policies too many requests response
func (o *GetPreventionPoliciesTooManyRequests) Code() int {
	return 429
}

func (o *GetPreventionPoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/entities/prevention/v1][%d] getPreventionPoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPreventionPoliciesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /policy/entities/prevention/v1][%d] getPreventionPoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetPreventionPoliciesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetPreventionPoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPreventionPoliciesInternalServerError creates a GetPreventionPoliciesInternalServerError with default headers values
func NewGetPreventionPoliciesInternalServerError() *GetPreventionPoliciesInternalServerError {
	return &GetPreventionPoliciesInternalServerError{}
}

/*
GetPreventionPoliciesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetPreventionPoliciesInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.PreventionRespV1
}

// IsSuccess returns true when this get prevention policies internal server error response has a 2xx status code
func (o *GetPreventionPoliciesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get prevention policies internal server error response has a 3xx status code
func (o *GetPreventionPoliciesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get prevention policies internal server error response has a 4xx status code
func (o *GetPreventionPoliciesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get prevention policies internal server error response has a 5xx status code
func (o *GetPreventionPoliciesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get prevention policies internal server error response a status code equal to that given
func (o *GetPreventionPoliciesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get prevention policies internal server error response
func (o *GetPreventionPoliciesInternalServerError) Code() int {
	return 500
}

func (o *GetPreventionPoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/entities/prevention/v1][%d] getPreventionPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPreventionPoliciesInternalServerError) String() string {
	return fmt.Sprintf("[GET /policy/entities/prevention/v1][%d] getPreventionPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPreventionPoliciesInternalServerError) GetPayload() *models.PreventionRespV1 {
	return o.Payload
}

func (o *GetPreventionPoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.PreventionRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
