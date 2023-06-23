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

// QueryPreventionPoliciesReader is a Reader for the QueryPreventionPolicies structure.
type QueryPreventionPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryPreventionPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryPreventionPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryPreventionPoliciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryPreventionPoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryPreventionPoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryPreventionPoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /policy/queries/prevention/v1] queryPreventionPolicies", response, response.Code())
	}
}

// NewQueryPreventionPoliciesOK creates a QueryPreventionPoliciesOK with default headers values
func NewQueryPreventionPoliciesOK() *QueryPreventionPoliciesOK {
	return &QueryPreventionPoliciesOK{}
}

/*
QueryPreventionPoliciesOK describes a response with status code 200, with default header values.

OK
*/
type QueryPreventionPoliciesOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query prevention policies o k response has a 2xx status code
func (o *QueryPreventionPoliciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query prevention policies o k response has a 3xx status code
func (o *QueryPreventionPoliciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query prevention policies o k response has a 4xx status code
func (o *QueryPreventionPoliciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query prevention policies o k response has a 5xx status code
func (o *QueryPreventionPoliciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query prevention policies o k response a status code equal to that given
func (o *QueryPreventionPoliciesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query prevention policies o k response
func (o *QueryPreventionPoliciesOK) Code() int {
	return 200
}

func (o *QueryPreventionPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /policy/queries/prevention/v1][%d] queryPreventionPoliciesOK  %+v", 200, o.Payload)
}

func (o *QueryPreventionPoliciesOK) String() string {
	return fmt.Sprintf("[GET /policy/queries/prevention/v1][%d] queryPreventionPoliciesOK  %+v", 200, o.Payload)
}

func (o *QueryPreventionPoliciesOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryPreventionPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryPreventionPoliciesBadRequest creates a QueryPreventionPoliciesBadRequest with default headers values
func NewQueryPreventionPoliciesBadRequest() *QueryPreventionPoliciesBadRequest {
	return &QueryPreventionPoliciesBadRequest{}
}

/*
QueryPreventionPoliciesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryPreventionPoliciesBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query prevention policies bad request response has a 2xx status code
func (o *QueryPreventionPoliciesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query prevention policies bad request response has a 3xx status code
func (o *QueryPreventionPoliciesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query prevention policies bad request response has a 4xx status code
func (o *QueryPreventionPoliciesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query prevention policies bad request response has a 5xx status code
func (o *QueryPreventionPoliciesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query prevention policies bad request response a status code equal to that given
func (o *QueryPreventionPoliciesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the query prevention policies bad request response
func (o *QueryPreventionPoliciesBadRequest) Code() int {
	return 400
}

func (o *QueryPreventionPoliciesBadRequest) Error() string {
	return fmt.Sprintf("[GET /policy/queries/prevention/v1][%d] queryPreventionPoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *QueryPreventionPoliciesBadRequest) String() string {
	return fmt.Sprintf("[GET /policy/queries/prevention/v1][%d] queryPreventionPoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *QueryPreventionPoliciesBadRequest) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryPreventionPoliciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryPreventionPoliciesForbidden creates a QueryPreventionPoliciesForbidden with default headers values
func NewQueryPreventionPoliciesForbidden() *QueryPreventionPoliciesForbidden {
	return &QueryPreventionPoliciesForbidden{}
}

/*
QueryPreventionPoliciesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryPreventionPoliciesForbidden struct {

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

// IsSuccess returns true when this query prevention policies forbidden response has a 2xx status code
func (o *QueryPreventionPoliciesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query prevention policies forbidden response has a 3xx status code
func (o *QueryPreventionPoliciesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query prevention policies forbidden response has a 4xx status code
func (o *QueryPreventionPoliciesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query prevention policies forbidden response has a 5xx status code
func (o *QueryPreventionPoliciesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query prevention policies forbidden response a status code equal to that given
func (o *QueryPreventionPoliciesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query prevention policies forbidden response
func (o *QueryPreventionPoliciesForbidden) Code() int {
	return 403
}

func (o *QueryPreventionPoliciesForbidden) Error() string {
	return fmt.Sprintf("[GET /policy/queries/prevention/v1][%d] queryPreventionPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *QueryPreventionPoliciesForbidden) String() string {
	return fmt.Sprintf("[GET /policy/queries/prevention/v1][%d] queryPreventionPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *QueryPreventionPoliciesForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryPreventionPoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryPreventionPoliciesTooManyRequests creates a QueryPreventionPoliciesTooManyRequests with default headers values
func NewQueryPreventionPoliciesTooManyRequests() *QueryPreventionPoliciesTooManyRequests {
	return &QueryPreventionPoliciesTooManyRequests{}
}

/*
QueryPreventionPoliciesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryPreventionPoliciesTooManyRequests struct {

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

// IsSuccess returns true when this query prevention policies too many requests response has a 2xx status code
func (o *QueryPreventionPoliciesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query prevention policies too many requests response has a 3xx status code
func (o *QueryPreventionPoliciesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query prevention policies too many requests response has a 4xx status code
func (o *QueryPreventionPoliciesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query prevention policies too many requests response has a 5xx status code
func (o *QueryPreventionPoliciesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query prevention policies too many requests response a status code equal to that given
func (o *QueryPreventionPoliciesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query prevention policies too many requests response
func (o *QueryPreventionPoliciesTooManyRequests) Code() int {
	return 429
}

func (o *QueryPreventionPoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/queries/prevention/v1][%d] queryPreventionPoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryPreventionPoliciesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /policy/queries/prevention/v1][%d] queryPreventionPoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryPreventionPoliciesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryPreventionPoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryPreventionPoliciesInternalServerError creates a QueryPreventionPoliciesInternalServerError with default headers values
func NewQueryPreventionPoliciesInternalServerError() *QueryPreventionPoliciesInternalServerError {
	return &QueryPreventionPoliciesInternalServerError{}
}

/*
QueryPreventionPoliciesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryPreventionPoliciesInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query prevention policies internal server error response has a 2xx status code
func (o *QueryPreventionPoliciesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query prevention policies internal server error response has a 3xx status code
func (o *QueryPreventionPoliciesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query prevention policies internal server error response has a 4xx status code
func (o *QueryPreventionPoliciesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query prevention policies internal server error response has a 5xx status code
func (o *QueryPreventionPoliciesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query prevention policies internal server error response a status code equal to that given
func (o *QueryPreventionPoliciesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the query prevention policies internal server error response
func (o *QueryPreventionPoliciesInternalServerError) Code() int {
	return 500
}

func (o *QueryPreventionPoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/queries/prevention/v1][%d] queryPreventionPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryPreventionPoliciesInternalServerError) String() string {
	return fmt.Sprintf("[GET /policy/queries/prevention/v1][%d] queryPreventionPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryPreventionPoliciesInternalServerError) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryPreventionPoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
