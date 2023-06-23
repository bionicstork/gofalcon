// Code generated by go-swagger; DO NOT EDIT.

package device_control_policies

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

// SetDeviceControlPoliciesPrecedenceReader is a Reader for the SetDeviceControlPoliciesPrecedence structure.
type SetDeviceControlPoliciesPrecedenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetDeviceControlPoliciesPrecedenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetDeviceControlPoliciesPrecedenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetDeviceControlPoliciesPrecedenceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSetDeviceControlPoliciesPrecedenceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewSetDeviceControlPoliciesPrecedenceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSetDeviceControlPoliciesPrecedenceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /policy/entities/device-control-precedence/v1] setDeviceControlPoliciesPrecedence", response, response.Code())
	}
}

// NewSetDeviceControlPoliciesPrecedenceOK creates a SetDeviceControlPoliciesPrecedenceOK with default headers values
func NewSetDeviceControlPoliciesPrecedenceOK() *SetDeviceControlPoliciesPrecedenceOK {
	return &SetDeviceControlPoliciesPrecedenceOK{}
}

/*
SetDeviceControlPoliciesPrecedenceOK describes a response with status code 200, with default header values.

OK
*/
type SetDeviceControlPoliciesPrecedenceOK struct {

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

// IsSuccess returns true when this set device control policies precedence o k response has a 2xx status code
func (o *SetDeviceControlPoliciesPrecedenceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set device control policies precedence o k response has a 3xx status code
func (o *SetDeviceControlPoliciesPrecedenceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set device control policies precedence o k response has a 4xx status code
func (o *SetDeviceControlPoliciesPrecedenceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this set device control policies precedence o k response has a 5xx status code
func (o *SetDeviceControlPoliciesPrecedenceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this set device control policies precedence o k response a status code equal to that given
func (o *SetDeviceControlPoliciesPrecedenceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the set device control policies precedence o k response
func (o *SetDeviceControlPoliciesPrecedenceOK) Code() int {
	return 200
}

func (o *SetDeviceControlPoliciesPrecedenceOK) Error() string {
	return fmt.Sprintf("[POST /policy/entities/device-control-precedence/v1][%d] setDeviceControlPoliciesPrecedenceOK  %+v", 200, o.Payload)
}

func (o *SetDeviceControlPoliciesPrecedenceOK) String() string {
	return fmt.Sprintf("[POST /policy/entities/device-control-precedence/v1][%d] setDeviceControlPoliciesPrecedenceOK  %+v", 200, o.Payload)
}

func (o *SetDeviceControlPoliciesPrecedenceOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *SetDeviceControlPoliciesPrecedenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSetDeviceControlPoliciesPrecedenceBadRequest creates a SetDeviceControlPoliciesPrecedenceBadRequest with default headers values
func NewSetDeviceControlPoliciesPrecedenceBadRequest() *SetDeviceControlPoliciesPrecedenceBadRequest {
	return &SetDeviceControlPoliciesPrecedenceBadRequest{}
}

/*
SetDeviceControlPoliciesPrecedenceBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SetDeviceControlPoliciesPrecedenceBadRequest struct {

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

// IsSuccess returns true when this set device control policies precedence bad request response has a 2xx status code
func (o *SetDeviceControlPoliciesPrecedenceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set device control policies precedence bad request response has a 3xx status code
func (o *SetDeviceControlPoliciesPrecedenceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set device control policies precedence bad request response has a 4xx status code
func (o *SetDeviceControlPoliciesPrecedenceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this set device control policies precedence bad request response has a 5xx status code
func (o *SetDeviceControlPoliciesPrecedenceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this set device control policies precedence bad request response a status code equal to that given
func (o *SetDeviceControlPoliciesPrecedenceBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the set device control policies precedence bad request response
func (o *SetDeviceControlPoliciesPrecedenceBadRequest) Code() int {
	return 400
}

func (o *SetDeviceControlPoliciesPrecedenceBadRequest) Error() string {
	return fmt.Sprintf("[POST /policy/entities/device-control-precedence/v1][%d] setDeviceControlPoliciesPrecedenceBadRequest  %+v", 400, o.Payload)
}

func (o *SetDeviceControlPoliciesPrecedenceBadRequest) String() string {
	return fmt.Sprintf("[POST /policy/entities/device-control-precedence/v1][%d] setDeviceControlPoliciesPrecedenceBadRequest  %+v", 400, o.Payload)
}

func (o *SetDeviceControlPoliciesPrecedenceBadRequest) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *SetDeviceControlPoliciesPrecedenceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSetDeviceControlPoliciesPrecedenceForbidden creates a SetDeviceControlPoliciesPrecedenceForbidden with default headers values
func NewSetDeviceControlPoliciesPrecedenceForbidden() *SetDeviceControlPoliciesPrecedenceForbidden {
	return &SetDeviceControlPoliciesPrecedenceForbidden{}
}

/*
SetDeviceControlPoliciesPrecedenceForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SetDeviceControlPoliciesPrecedenceForbidden struct {

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

// IsSuccess returns true when this set device control policies precedence forbidden response has a 2xx status code
func (o *SetDeviceControlPoliciesPrecedenceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set device control policies precedence forbidden response has a 3xx status code
func (o *SetDeviceControlPoliciesPrecedenceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set device control policies precedence forbidden response has a 4xx status code
func (o *SetDeviceControlPoliciesPrecedenceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this set device control policies precedence forbidden response has a 5xx status code
func (o *SetDeviceControlPoliciesPrecedenceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this set device control policies precedence forbidden response a status code equal to that given
func (o *SetDeviceControlPoliciesPrecedenceForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the set device control policies precedence forbidden response
func (o *SetDeviceControlPoliciesPrecedenceForbidden) Code() int {
	return 403
}

func (o *SetDeviceControlPoliciesPrecedenceForbidden) Error() string {
	return fmt.Sprintf("[POST /policy/entities/device-control-precedence/v1][%d] setDeviceControlPoliciesPrecedenceForbidden  %+v", 403, o.Payload)
}

func (o *SetDeviceControlPoliciesPrecedenceForbidden) String() string {
	return fmt.Sprintf("[POST /policy/entities/device-control-precedence/v1][%d] setDeviceControlPoliciesPrecedenceForbidden  %+v", 403, o.Payload)
}

func (o *SetDeviceControlPoliciesPrecedenceForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *SetDeviceControlPoliciesPrecedenceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSetDeviceControlPoliciesPrecedenceTooManyRequests creates a SetDeviceControlPoliciesPrecedenceTooManyRequests with default headers values
func NewSetDeviceControlPoliciesPrecedenceTooManyRequests() *SetDeviceControlPoliciesPrecedenceTooManyRequests {
	return &SetDeviceControlPoliciesPrecedenceTooManyRequests{}
}

/*
SetDeviceControlPoliciesPrecedenceTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type SetDeviceControlPoliciesPrecedenceTooManyRequests struct {

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

// IsSuccess returns true when this set device control policies precedence too many requests response has a 2xx status code
func (o *SetDeviceControlPoliciesPrecedenceTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set device control policies precedence too many requests response has a 3xx status code
func (o *SetDeviceControlPoliciesPrecedenceTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set device control policies precedence too many requests response has a 4xx status code
func (o *SetDeviceControlPoliciesPrecedenceTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this set device control policies precedence too many requests response has a 5xx status code
func (o *SetDeviceControlPoliciesPrecedenceTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this set device control policies precedence too many requests response a status code equal to that given
func (o *SetDeviceControlPoliciesPrecedenceTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the set device control policies precedence too many requests response
func (o *SetDeviceControlPoliciesPrecedenceTooManyRequests) Code() int {
	return 429
}

func (o *SetDeviceControlPoliciesPrecedenceTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /policy/entities/device-control-precedence/v1][%d] setDeviceControlPoliciesPrecedenceTooManyRequests  %+v", 429, o.Payload)
}

func (o *SetDeviceControlPoliciesPrecedenceTooManyRequests) String() string {
	return fmt.Sprintf("[POST /policy/entities/device-control-precedence/v1][%d] setDeviceControlPoliciesPrecedenceTooManyRequests  %+v", 429, o.Payload)
}

func (o *SetDeviceControlPoliciesPrecedenceTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *SetDeviceControlPoliciesPrecedenceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSetDeviceControlPoliciesPrecedenceInternalServerError creates a SetDeviceControlPoliciesPrecedenceInternalServerError with default headers values
func NewSetDeviceControlPoliciesPrecedenceInternalServerError() *SetDeviceControlPoliciesPrecedenceInternalServerError {
	return &SetDeviceControlPoliciesPrecedenceInternalServerError{}
}

/*
SetDeviceControlPoliciesPrecedenceInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type SetDeviceControlPoliciesPrecedenceInternalServerError struct {

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

// IsSuccess returns true when this set device control policies precedence internal server error response has a 2xx status code
func (o *SetDeviceControlPoliciesPrecedenceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set device control policies precedence internal server error response has a 3xx status code
func (o *SetDeviceControlPoliciesPrecedenceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set device control policies precedence internal server error response has a 4xx status code
func (o *SetDeviceControlPoliciesPrecedenceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this set device control policies precedence internal server error response has a 5xx status code
func (o *SetDeviceControlPoliciesPrecedenceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this set device control policies precedence internal server error response a status code equal to that given
func (o *SetDeviceControlPoliciesPrecedenceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the set device control policies precedence internal server error response
func (o *SetDeviceControlPoliciesPrecedenceInternalServerError) Code() int {
	return 500
}

func (o *SetDeviceControlPoliciesPrecedenceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /policy/entities/device-control-precedence/v1][%d] setDeviceControlPoliciesPrecedenceInternalServerError  %+v", 500, o.Payload)
}

func (o *SetDeviceControlPoliciesPrecedenceInternalServerError) String() string {
	return fmt.Sprintf("[POST /policy/entities/device-control-precedence/v1][%d] setDeviceControlPoliciesPrecedenceInternalServerError  %+v", 500, o.Payload)
}

func (o *SetDeviceControlPoliciesPrecedenceInternalServerError) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *SetDeviceControlPoliciesPrecedenceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
