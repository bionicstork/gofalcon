// Code generated by go-swagger; DO NOT EDIT.

package user_management

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

// QueryUserV1Reader is a Reader for the QueryUserV1 structure.
type QueryUserV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryUserV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryUserV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryUserV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryUserV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryUserV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryUserV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /user-management/queries/users/v1] queryUserV1", response, response.Code())
	}
}

// NewQueryUserV1OK creates a QueryUserV1OK with default headers values
func NewQueryUserV1OK() *QueryUserV1OK {
	return &QueryUserV1OK{}
}

/*
QueryUserV1OK describes a response with status code 200, with default header values.

OK
*/
type QueryUserV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecQueryResponse
}

// IsSuccess returns true when this query user v1 o k response has a 2xx status code
func (o *QueryUserV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query user v1 o k response has a 3xx status code
func (o *QueryUserV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query user v1 o k response has a 4xx status code
func (o *QueryUserV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query user v1 o k response has a 5xx status code
func (o *QueryUserV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this query user v1 o k response a status code equal to that given
func (o *QueryUserV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query user v1 o k response
func (o *QueryUserV1OK) Code() int {
	return 200
}

func (o *QueryUserV1OK) Error() string {
	return fmt.Sprintf("[GET /user-management/queries/users/v1][%d] queryUserV1OK  %+v", 200, o.Payload)
}

func (o *QueryUserV1OK) String() string {
	return fmt.Sprintf("[GET /user-management/queries/users/v1][%d] queryUserV1OK  %+v", 200, o.Payload)
}

func (o *QueryUserV1OK) GetPayload() *models.MsaspecQueryResponse {
	return o.Payload
}

func (o *QueryUserV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryUserV1BadRequest creates a QueryUserV1BadRequest with default headers values
func NewQueryUserV1BadRequest() *QueryUserV1BadRequest {
	return &QueryUserV1BadRequest{}
}

/*
QueryUserV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryUserV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this query user v1 bad request response has a 2xx status code
func (o *QueryUserV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query user v1 bad request response has a 3xx status code
func (o *QueryUserV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query user v1 bad request response has a 4xx status code
func (o *QueryUserV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query user v1 bad request response has a 5xx status code
func (o *QueryUserV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query user v1 bad request response a status code equal to that given
func (o *QueryUserV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the query user v1 bad request response
func (o *QueryUserV1BadRequest) Code() int {
	return 400
}

func (o *QueryUserV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /user-management/queries/users/v1][%d] queryUserV1BadRequest  %+v", 400, o.Payload)
}

func (o *QueryUserV1BadRequest) String() string {
	return fmt.Sprintf("[GET /user-management/queries/users/v1][%d] queryUserV1BadRequest  %+v", 400, o.Payload)
}

func (o *QueryUserV1BadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *QueryUserV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryUserV1Forbidden creates a QueryUserV1Forbidden with default headers values
func NewQueryUserV1Forbidden() *QueryUserV1Forbidden {
	return &QueryUserV1Forbidden{}
}

/*
QueryUserV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryUserV1Forbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this query user v1 forbidden response has a 2xx status code
func (o *QueryUserV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query user v1 forbidden response has a 3xx status code
func (o *QueryUserV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query user v1 forbidden response has a 4xx status code
func (o *QueryUserV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query user v1 forbidden response has a 5xx status code
func (o *QueryUserV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query user v1 forbidden response a status code equal to that given
func (o *QueryUserV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query user v1 forbidden response
func (o *QueryUserV1Forbidden) Code() int {
	return 403
}

func (o *QueryUserV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /user-management/queries/users/v1][%d] queryUserV1Forbidden  %+v", 403, o.Payload)
}

func (o *QueryUserV1Forbidden) String() string {
	return fmt.Sprintf("[GET /user-management/queries/users/v1][%d] queryUserV1Forbidden  %+v", 403, o.Payload)
}

func (o *QueryUserV1Forbidden) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *QueryUserV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryUserV1TooManyRequests creates a QueryUserV1TooManyRequests with default headers values
func NewQueryUserV1TooManyRequests() *QueryUserV1TooManyRequests {
	return &QueryUserV1TooManyRequests{}
}

/*
QueryUserV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryUserV1TooManyRequests struct {

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

// IsSuccess returns true when this query user v1 too many requests response has a 2xx status code
func (o *QueryUserV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query user v1 too many requests response has a 3xx status code
func (o *QueryUserV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query user v1 too many requests response has a 4xx status code
func (o *QueryUserV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query user v1 too many requests response has a 5xx status code
func (o *QueryUserV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query user v1 too many requests response a status code equal to that given
func (o *QueryUserV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query user v1 too many requests response
func (o *QueryUserV1TooManyRequests) Code() int {
	return 429
}

func (o *QueryUserV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /user-management/queries/users/v1][%d] queryUserV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryUserV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /user-management/queries/users/v1][%d] queryUserV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryUserV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryUserV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryUserV1InternalServerError creates a QueryUserV1InternalServerError with default headers values
func NewQueryUserV1InternalServerError() *QueryUserV1InternalServerError {
	return &QueryUserV1InternalServerError{}
}

/*
QueryUserV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryUserV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this query user v1 internal server error response has a 2xx status code
func (o *QueryUserV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query user v1 internal server error response has a 3xx status code
func (o *QueryUserV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query user v1 internal server error response has a 4xx status code
func (o *QueryUserV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query user v1 internal server error response has a 5xx status code
func (o *QueryUserV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query user v1 internal server error response a status code equal to that given
func (o *QueryUserV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the query user v1 internal server error response
func (o *QueryUserV1InternalServerError) Code() int {
	return 500
}

func (o *QueryUserV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /user-management/queries/users/v1][%d] queryUserV1InternalServerError  %+v", 500, o.Payload)
}

func (o *QueryUserV1InternalServerError) String() string {
	return fmt.Sprintf("[GET /user-management/queries/users/v1][%d] queryUserV1InternalServerError  %+v", 500, o.Payload)
}

func (o *QueryUserV1InternalServerError) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *QueryUserV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
