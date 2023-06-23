// Code generated by go-swagger; DO NOT EDIT.

package real_time_response_admin

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

// RTRExecuteAdminCommandReader is a Reader for the RTRExecuteAdminCommand structure.
type RTRExecuteAdminCommandReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RTRExecuteAdminCommandReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRTRExecuteAdminCommandCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRTRExecuteAdminCommandBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRTRExecuteAdminCommandForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRTRExecuteAdminCommandTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /real-time-response/entities/admin-command/v1] RTR-ExecuteAdminCommand", response, response.Code())
	}
}

// NewRTRExecuteAdminCommandCreated creates a RTRExecuteAdminCommandCreated with default headers values
func NewRTRExecuteAdminCommandCreated() *RTRExecuteAdminCommandCreated {
	return &RTRExecuteAdminCommandCreated{}
}

/*
RTRExecuteAdminCommandCreated describes a response with status code 201, with default header values.

Created
*/
type RTRExecuteAdminCommandCreated struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainCommandExecuteResponseWrapper
}

// IsSuccess returns true when this r t r execute admin command created response has a 2xx status code
func (o *RTRExecuteAdminCommandCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this r t r execute admin command created response has a 3xx status code
func (o *RTRExecuteAdminCommandCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r execute admin command created response has a 4xx status code
func (o *RTRExecuteAdminCommandCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this r t r execute admin command created response has a 5xx status code
func (o *RTRExecuteAdminCommandCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r execute admin command created response a status code equal to that given
func (o *RTRExecuteAdminCommandCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the r t r execute admin command created response
func (o *RTRExecuteAdminCommandCreated) Code() int {
	return 201
}

func (o *RTRExecuteAdminCommandCreated) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/admin-command/v1][%d] rTRExecuteAdminCommandCreated  %+v", 201, o.Payload)
}

func (o *RTRExecuteAdminCommandCreated) String() string {
	return fmt.Sprintf("[POST /real-time-response/entities/admin-command/v1][%d] rTRExecuteAdminCommandCreated  %+v", 201, o.Payload)
}

func (o *RTRExecuteAdminCommandCreated) GetPayload() *models.DomainCommandExecuteResponseWrapper {
	return o.Payload
}

func (o *RTRExecuteAdminCommandCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainCommandExecuteResponseWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRExecuteAdminCommandBadRequest creates a RTRExecuteAdminCommandBadRequest with default headers values
func NewRTRExecuteAdminCommandBadRequest() *RTRExecuteAdminCommandBadRequest {
	return &RTRExecuteAdminCommandBadRequest{}
}

/*
RTRExecuteAdminCommandBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RTRExecuteAdminCommandBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

// IsSuccess returns true when this r t r execute admin command bad request response has a 2xx status code
func (o *RTRExecuteAdminCommandBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r execute admin command bad request response has a 3xx status code
func (o *RTRExecuteAdminCommandBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r execute admin command bad request response has a 4xx status code
func (o *RTRExecuteAdminCommandBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r execute admin command bad request response has a 5xx status code
func (o *RTRExecuteAdminCommandBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r execute admin command bad request response a status code equal to that given
func (o *RTRExecuteAdminCommandBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the r t r execute admin command bad request response
func (o *RTRExecuteAdminCommandBadRequest) Code() int {
	return 400
}

func (o *RTRExecuteAdminCommandBadRequest) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/admin-command/v1][%d] rTRExecuteAdminCommandBadRequest  %+v", 400, o.Payload)
}

func (o *RTRExecuteAdminCommandBadRequest) String() string {
	return fmt.Sprintf("[POST /real-time-response/entities/admin-command/v1][%d] rTRExecuteAdminCommandBadRequest  %+v", 400, o.Payload)
}

func (o *RTRExecuteAdminCommandBadRequest) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRExecuteAdminCommandBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRExecuteAdminCommandForbidden creates a RTRExecuteAdminCommandForbidden with default headers values
func NewRTRExecuteAdminCommandForbidden() *RTRExecuteAdminCommandForbidden {
	return &RTRExecuteAdminCommandForbidden{}
}

/*
RTRExecuteAdminCommandForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RTRExecuteAdminCommandForbidden struct {

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

// IsSuccess returns true when this r t r execute admin command forbidden response has a 2xx status code
func (o *RTRExecuteAdminCommandForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r execute admin command forbidden response has a 3xx status code
func (o *RTRExecuteAdminCommandForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r execute admin command forbidden response has a 4xx status code
func (o *RTRExecuteAdminCommandForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r execute admin command forbidden response has a 5xx status code
func (o *RTRExecuteAdminCommandForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r execute admin command forbidden response a status code equal to that given
func (o *RTRExecuteAdminCommandForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the r t r execute admin command forbidden response
func (o *RTRExecuteAdminCommandForbidden) Code() int {
	return 403
}

func (o *RTRExecuteAdminCommandForbidden) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/admin-command/v1][%d] rTRExecuteAdminCommandForbidden  %+v", 403, o.Payload)
}

func (o *RTRExecuteAdminCommandForbidden) String() string {
	return fmt.Sprintf("[POST /real-time-response/entities/admin-command/v1][%d] rTRExecuteAdminCommandForbidden  %+v", 403, o.Payload)
}

func (o *RTRExecuteAdminCommandForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRExecuteAdminCommandForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRExecuteAdminCommandTooManyRequests creates a RTRExecuteAdminCommandTooManyRequests with default headers values
func NewRTRExecuteAdminCommandTooManyRequests() *RTRExecuteAdminCommandTooManyRequests {
	return &RTRExecuteAdminCommandTooManyRequests{}
}

/*
RTRExecuteAdminCommandTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RTRExecuteAdminCommandTooManyRequests struct {

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

// IsSuccess returns true when this r t r execute admin command too many requests response has a 2xx status code
func (o *RTRExecuteAdminCommandTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r execute admin command too many requests response has a 3xx status code
func (o *RTRExecuteAdminCommandTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r execute admin command too many requests response has a 4xx status code
func (o *RTRExecuteAdminCommandTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r execute admin command too many requests response has a 5xx status code
func (o *RTRExecuteAdminCommandTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r execute admin command too many requests response a status code equal to that given
func (o *RTRExecuteAdminCommandTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the r t r execute admin command too many requests response
func (o *RTRExecuteAdminCommandTooManyRequests) Code() int {
	return 429
}

func (o *RTRExecuteAdminCommandTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /real-time-response/entities/admin-command/v1][%d] rTRExecuteAdminCommandTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRExecuteAdminCommandTooManyRequests) String() string {
	return fmt.Sprintf("[POST /real-time-response/entities/admin-command/v1][%d] rTRExecuteAdminCommandTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRExecuteAdminCommandTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRExecuteAdminCommandTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
