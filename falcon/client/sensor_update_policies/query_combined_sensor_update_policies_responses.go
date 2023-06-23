// Code generated by go-swagger; DO NOT EDIT.

package sensor_update_policies

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

// QueryCombinedSensorUpdatePoliciesReader is a Reader for the QueryCombinedSensorUpdatePolicies structure.
type QueryCombinedSensorUpdatePoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryCombinedSensorUpdatePoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryCombinedSensorUpdatePoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryCombinedSensorUpdatePoliciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryCombinedSensorUpdatePoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryCombinedSensorUpdatePoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryCombinedSensorUpdatePoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /policy/combined/sensor-update/v1] queryCombinedSensorUpdatePolicies", response, response.Code())
	}
}

// NewQueryCombinedSensorUpdatePoliciesOK creates a QueryCombinedSensorUpdatePoliciesOK with default headers values
func NewQueryCombinedSensorUpdatePoliciesOK() *QueryCombinedSensorUpdatePoliciesOK {
	return &QueryCombinedSensorUpdatePoliciesOK{}
}

/*
QueryCombinedSensorUpdatePoliciesOK describes a response with status code 200, with default header values.

OK
*/
type QueryCombinedSensorUpdatePoliciesOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.SensorUpdateRespV1
}

// IsSuccess returns true when this query combined sensor update policies o k response has a 2xx status code
func (o *QueryCombinedSensorUpdatePoliciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query combined sensor update policies o k response has a 3xx status code
func (o *QueryCombinedSensorUpdatePoliciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined sensor update policies o k response has a 4xx status code
func (o *QueryCombinedSensorUpdatePoliciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query combined sensor update policies o k response has a 5xx status code
func (o *QueryCombinedSensorUpdatePoliciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined sensor update policies o k response a status code equal to that given
func (o *QueryCombinedSensorUpdatePoliciesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query combined sensor update policies o k response
func (o *QueryCombinedSensorUpdatePoliciesOK) Code() int {
	return 200
}

func (o *QueryCombinedSensorUpdatePoliciesOK) Error() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update/v1][%d] queryCombinedSensorUpdatePoliciesOK  %+v", 200, o.Payload)
}

func (o *QueryCombinedSensorUpdatePoliciesOK) String() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update/v1][%d] queryCombinedSensorUpdatePoliciesOK  %+v", 200, o.Payload)
}

func (o *QueryCombinedSensorUpdatePoliciesOK) GetPayload() *models.SensorUpdateRespV1 {
	return o.Payload
}

func (o *QueryCombinedSensorUpdatePoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.SensorUpdateRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedSensorUpdatePoliciesBadRequest creates a QueryCombinedSensorUpdatePoliciesBadRequest with default headers values
func NewQueryCombinedSensorUpdatePoliciesBadRequest() *QueryCombinedSensorUpdatePoliciesBadRequest {
	return &QueryCombinedSensorUpdatePoliciesBadRequest{}
}

/*
QueryCombinedSensorUpdatePoliciesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryCombinedSensorUpdatePoliciesBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.SensorUpdateRespV1
}

// IsSuccess returns true when this query combined sensor update policies bad request response has a 2xx status code
func (o *QueryCombinedSensorUpdatePoliciesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined sensor update policies bad request response has a 3xx status code
func (o *QueryCombinedSensorUpdatePoliciesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined sensor update policies bad request response has a 4xx status code
func (o *QueryCombinedSensorUpdatePoliciesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this query combined sensor update policies bad request response has a 5xx status code
func (o *QueryCombinedSensorUpdatePoliciesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined sensor update policies bad request response a status code equal to that given
func (o *QueryCombinedSensorUpdatePoliciesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the query combined sensor update policies bad request response
func (o *QueryCombinedSensorUpdatePoliciesBadRequest) Code() int {
	return 400
}

func (o *QueryCombinedSensorUpdatePoliciesBadRequest) Error() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update/v1][%d] queryCombinedSensorUpdatePoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *QueryCombinedSensorUpdatePoliciesBadRequest) String() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update/v1][%d] queryCombinedSensorUpdatePoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *QueryCombinedSensorUpdatePoliciesBadRequest) GetPayload() *models.SensorUpdateRespV1 {
	return o.Payload
}

func (o *QueryCombinedSensorUpdatePoliciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.SensorUpdateRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCombinedSensorUpdatePoliciesForbidden creates a QueryCombinedSensorUpdatePoliciesForbidden with default headers values
func NewQueryCombinedSensorUpdatePoliciesForbidden() *QueryCombinedSensorUpdatePoliciesForbidden {
	return &QueryCombinedSensorUpdatePoliciesForbidden{}
}

/*
QueryCombinedSensorUpdatePoliciesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryCombinedSensorUpdatePoliciesForbidden struct {

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

// IsSuccess returns true when this query combined sensor update policies forbidden response has a 2xx status code
func (o *QueryCombinedSensorUpdatePoliciesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined sensor update policies forbidden response has a 3xx status code
func (o *QueryCombinedSensorUpdatePoliciesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined sensor update policies forbidden response has a 4xx status code
func (o *QueryCombinedSensorUpdatePoliciesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query combined sensor update policies forbidden response has a 5xx status code
func (o *QueryCombinedSensorUpdatePoliciesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined sensor update policies forbidden response a status code equal to that given
func (o *QueryCombinedSensorUpdatePoliciesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query combined sensor update policies forbidden response
func (o *QueryCombinedSensorUpdatePoliciesForbidden) Code() int {
	return 403
}

func (o *QueryCombinedSensorUpdatePoliciesForbidden) Error() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update/v1][%d] queryCombinedSensorUpdatePoliciesForbidden  %+v", 403, o.Payload)
}

func (o *QueryCombinedSensorUpdatePoliciesForbidden) String() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update/v1][%d] queryCombinedSensorUpdatePoliciesForbidden  %+v", 403, o.Payload)
}

func (o *QueryCombinedSensorUpdatePoliciesForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryCombinedSensorUpdatePoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryCombinedSensorUpdatePoliciesTooManyRequests creates a QueryCombinedSensorUpdatePoliciesTooManyRequests with default headers values
func NewQueryCombinedSensorUpdatePoliciesTooManyRequests() *QueryCombinedSensorUpdatePoliciesTooManyRequests {
	return &QueryCombinedSensorUpdatePoliciesTooManyRequests{}
}

/*
QueryCombinedSensorUpdatePoliciesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryCombinedSensorUpdatePoliciesTooManyRequests struct {

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

// IsSuccess returns true when this query combined sensor update policies too many requests response has a 2xx status code
func (o *QueryCombinedSensorUpdatePoliciesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined sensor update policies too many requests response has a 3xx status code
func (o *QueryCombinedSensorUpdatePoliciesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined sensor update policies too many requests response has a 4xx status code
func (o *QueryCombinedSensorUpdatePoliciesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query combined sensor update policies too many requests response has a 5xx status code
func (o *QueryCombinedSensorUpdatePoliciesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query combined sensor update policies too many requests response a status code equal to that given
func (o *QueryCombinedSensorUpdatePoliciesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query combined sensor update policies too many requests response
func (o *QueryCombinedSensorUpdatePoliciesTooManyRequests) Code() int {
	return 429
}

func (o *QueryCombinedSensorUpdatePoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update/v1][%d] queryCombinedSensorUpdatePoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryCombinedSensorUpdatePoliciesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update/v1][%d] queryCombinedSensorUpdatePoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryCombinedSensorUpdatePoliciesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryCombinedSensorUpdatePoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryCombinedSensorUpdatePoliciesInternalServerError creates a QueryCombinedSensorUpdatePoliciesInternalServerError with default headers values
func NewQueryCombinedSensorUpdatePoliciesInternalServerError() *QueryCombinedSensorUpdatePoliciesInternalServerError {
	return &QueryCombinedSensorUpdatePoliciesInternalServerError{}
}

/*
QueryCombinedSensorUpdatePoliciesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryCombinedSensorUpdatePoliciesInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.SensorUpdateRespV1
}

// IsSuccess returns true when this query combined sensor update policies internal server error response has a 2xx status code
func (o *QueryCombinedSensorUpdatePoliciesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query combined sensor update policies internal server error response has a 3xx status code
func (o *QueryCombinedSensorUpdatePoliciesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query combined sensor update policies internal server error response has a 4xx status code
func (o *QueryCombinedSensorUpdatePoliciesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query combined sensor update policies internal server error response has a 5xx status code
func (o *QueryCombinedSensorUpdatePoliciesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query combined sensor update policies internal server error response a status code equal to that given
func (o *QueryCombinedSensorUpdatePoliciesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the query combined sensor update policies internal server error response
func (o *QueryCombinedSensorUpdatePoliciesInternalServerError) Code() int {
	return 500
}

func (o *QueryCombinedSensorUpdatePoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update/v1][%d] queryCombinedSensorUpdatePoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryCombinedSensorUpdatePoliciesInternalServerError) String() string {
	return fmt.Sprintf("[GET /policy/combined/sensor-update/v1][%d] queryCombinedSensorUpdatePoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryCombinedSensorUpdatePoliciesInternalServerError) GetPayload() *models.SensorUpdateRespV1 {
	return o.Payload
}

func (o *QueryCombinedSensorUpdatePoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.SensorUpdateRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
