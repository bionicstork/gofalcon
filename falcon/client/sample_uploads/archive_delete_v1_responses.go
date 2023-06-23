// Code generated by go-swagger; DO NOT EDIT.

package sample_uploads

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

// ArchiveDeleteV1Reader is a Reader for the ArchiveDeleteV1 structure.
type ArchiveDeleteV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ArchiveDeleteV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewArchiveDeleteV1Accepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewArchiveDeleteV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewArchiveDeleteV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewArchiveDeleteV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewArchiveDeleteV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewArchiveDeleteV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /archives/entities/archives/v1] ArchiveDeleteV1", response, response.Code())
	}
}

// NewArchiveDeleteV1Accepted creates a ArchiveDeleteV1Accepted with default headers values
func NewArchiveDeleteV1Accepted() *ArchiveDeleteV1Accepted {
	return &ArchiveDeleteV1Accepted{}
}

/*
ArchiveDeleteV1Accepted describes a response with status code 202, with default header values.

OK
*/
type ArchiveDeleteV1Accepted struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64
}

// IsSuccess returns true when this archive delete v1 accepted response has a 2xx status code
func (o *ArchiveDeleteV1Accepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this archive delete v1 accepted response has a 3xx status code
func (o *ArchiveDeleteV1Accepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive delete v1 accepted response has a 4xx status code
func (o *ArchiveDeleteV1Accepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this archive delete v1 accepted response has a 5xx status code
func (o *ArchiveDeleteV1Accepted) IsServerError() bool {
	return false
}

// IsCode returns true when this archive delete v1 accepted response a status code equal to that given
func (o *ArchiveDeleteV1Accepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the archive delete v1 accepted response
func (o *ArchiveDeleteV1Accepted) Code() int {
	return 202
}

func (o *ArchiveDeleteV1Accepted) Error() string {
	return fmt.Sprintf("[DELETE /archives/entities/archives/v1][%d] archiveDeleteV1Accepted ", 202)
}

func (o *ArchiveDeleteV1Accepted) String() string {
	return fmt.Sprintf("[DELETE /archives/entities/archives/v1][%d] archiveDeleteV1Accepted ", 202)
}

func (o *ArchiveDeleteV1Accepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	return nil
}

// NewArchiveDeleteV1BadRequest creates a ArchiveDeleteV1BadRequest with default headers values
func NewArchiveDeleteV1BadRequest() *ArchiveDeleteV1BadRequest {
	return &ArchiveDeleteV1BadRequest{}
}

/*
ArchiveDeleteV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ArchiveDeleteV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64
}

// IsSuccess returns true when this archive delete v1 bad request response has a 2xx status code
func (o *ArchiveDeleteV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this archive delete v1 bad request response has a 3xx status code
func (o *ArchiveDeleteV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive delete v1 bad request response has a 4xx status code
func (o *ArchiveDeleteV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this archive delete v1 bad request response has a 5xx status code
func (o *ArchiveDeleteV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this archive delete v1 bad request response a status code equal to that given
func (o *ArchiveDeleteV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the archive delete v1 bad request response
func (o *ArchiveDeleteV1BadRequest) Code() int {
	return 400
}

func (o *ArchiveDeleteV1BadRequest) Error() string {
	return fmt.Sprintf("[DELETE /archives/entities/archives/v1][%d] archiveDeleteV1BadRequest ", 400)
}

func (o *ArchiveDeleteV1BadRequest) String() string {
	return fmt.Sprintf("[DELETE /archives/entities/archives/v1][%d] archiveDeleteV1BadRequest ", 400)
}

func (o *ArchiveDeleteV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	return nil
}

// NewArchiveDeleteV1Forbidden creates a ArchiveDeleteV1Forbidden with default headers values
func NewArchiveDeleteV1Forbidden() *ArchiveDeleteV1Forbidden {
	return &ArchiveDeleteV1Forbidden{}
}

/*
ArchiveDeleteV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ArchiveDeleteV1Forbidden struct {

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

// IsSuccess returns true when this archive delete v1 forbidden response has a 2xx status code
func (o *ArchiveDeleteV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this archive delete v1 forbidden response has a 3xx status code
func (o *ArchiveDeleteV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive delete v1 forbidden response has a 4xx status code
func (o *ArchiveDeleteV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this archive delete v1 forbidden response has a 5xx status code
func (o *ArchiveDeleteV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this archive delete v1 forbidden response a status code equal to that given
func (o *ArchiveDeleteV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the archive delete v1 forbidden response
func (o *ArchiveDeleteV1Forbidden) Code() int {
	return 403
}

func (o *ArchiveDeleteV1Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /archives/entities/archives/v1][%d] archiveDeleteV1Forbidden  %+v", 403, o.Payload)
}

func (o *ArchiveDeleteV1Forbidden) String() string {
	return fmt.Sprintf("[DELETE /archives/entities/archives/v1][%d] archiveDeleteV1Forbidden  %+v", 403, o.Payload)
}

func (o *ArchiveDeleteV1Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ArchiveDeleteV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewArchiveDeleteV1NotFound creates a ArchiveDeleteV1NotFound with default headers values
func NewArchiveDeleteV1NotFound() *ArchiveDeleteV1NotFound {
	return &ArchiveDeleteV1NotFound{}
}

/*
ArchiveDeleteV1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type ArchiveDeleteV1NotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64
}

// IsSuccess returns true when this archive delete v1 not found response has a 2xx status code
func (o *ArchiveDeleteV1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this archive delete v1 not found response has a 3xx status code
func (o *ArchiveDeleteV1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive delete v1 not found response has a 4xx status code
func (o *ArchiveDeleteV1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this archive delete v1 not found response has a 5xx status code
func (o *ArchiveDeleteV1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this archive delete v1 not found response a status code equal to that given
func (o *ArchiveDeleteV1NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the archive delete v1 not found response
func (o *ArchiveDeleteV1NotFound) Code() int {
	return 404
}

func (o *ArchiveDeleteV1NotFound) Error() string {
	return fmt.Sprintf("[DELETE /archives/entities/archives/v1][%d] archiveDeleteV1NotFound ", 404)
}

func (o *ArchiveDeleteV1NotFound) String() string {
	return fmt.Sprintf("[DELETE /archives/entities/archives/v1][%d] archiveDeleteV1NotFound ", 404)
}

func (o *ArchiveDeleteV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	return nil
}

// NewArchiveDeleteV1TooManyRequests creates a ArchiveDeleteV1TooManyRequests with default headers values
func NewArchiveDeleteV1TooManyRequests() *ArchiveDeleteV1TooManyRequests {
	return &ArchiveDeleteV1TooManyRequests{}
}

/*
ArchiveDeleteV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type ArchiveDeleteV1TooManyRequests struct {

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

// IsSuccess returns true when this archive delete v1 too many requests response has a 2xx status code
func (o *ArchiveDeleteV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this archive delete v1 too many requests response has a 3xx status code
func (o *ArchiveDeleteV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive delete v1 too many requests response has a 4xx status code
func (o *ArchiveDeleteV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this archive delete v1 too many requests response has a 5xx status code
func (o *ArchiveDeleteV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this archive delete v1 too many requests response a status code equal to that given
func (o *ArchiveDeleteV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the archive delete v1 too many requests response
func (o *ArchiveDeleteV1TooManyRequests) Code() int {
	return 429
}

func (o *ArchiveDeleteV1TooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /archives/entities/archives/v1][%d] archiveDeleteV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *ArchiveDeleteV1TooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /archives/entities/archives/v1][%d] archiveDeleteV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *ArchiveDeleteV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *ArchiveDeleteV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewArchiveDeleteV1InternalServerError creates a ArchiveDeleteV1InternalServerError with default headers values
func NewArchiveDeleteV1InternalServerError() *ArchiveDeleteV1InternalServerError {
	return &ArchiveDeleteV1InternalServerError{}
}

/*
ArchiveDeleteV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ArchiveDeleteV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64
}

// IsSuccess returns true when this archive delete v1 internal server error response has a 2xx status code
func (o *ArchiveDeleteV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this archive delete v1 internal server error response has a 3xx status code
func (o *ArchiveDeleteV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this archive delete v1 internal server error response has a 4xx status code
func (o *ArchiveDeleteV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this archive delete v1 internal server error response has a 5xx status code
func (o *ArchiveDeleteV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this archive delete v1 internal server error response a status code equal to that given
func (o *ArchiveDeleteV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the archive delete v1 internal server error response
func (o *ArchiveDeleteV1InternalServerError) Code() int {
	return 500
}

func (o *ArchiveDeleteV1InternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /archives/entities/archives/v1][%d] archiveDeleteV1InternalServerError ", 500)
}

func (o *ArchiveDeleteV1InternalServerError) String() string {
	return fmt.Sprintf("[DELETE /archives/entities/archives/v1][%d] archiveDeleteV1InternalServerError ", 500)
}

func (o *ArchiveDeleteV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	return nil
}
