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

// UploadSampleV3Reader is a Reader for the UploadSampleV3 structure.
type UploadSampleV3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadSampleV3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUploadSampleV3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUploadSampleV3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUploadSampleV3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUploadSampleV3TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUploadSampleV3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /samples/entities/samples/v3] UploadSampleV3", response, response.Code())
	}
}

// NewUploadSampleV3OK creates a UploadSampleV3OK with default headers values
func NewUploadSampleV3OK() *UploadSampleV3OK {
	return &UploadSampleV3OK{}
}

/*
UploadSampleV3OK describes a response with status code 200, with default header values.

OK
*/
type UploadSampleV3OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ClientSampleMetadataResponseV2
}

// IsSuccess returns true when this upload sample v3 o k response has a 2xx status code
func (o *UploadSampleV3OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upload sample v3 o k response has a 3xx status code
func (o *UploadSampleV3OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload sample v3 o k response has a 4xx status code
func (o *UploadSampleV3OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload sample v3 o k response has a 5xx status code
func (o *UploadSampleV3OK) IsServerError() bool {
	return false
}

// IsCode returns true when this upload sample v3 o k response a status code equal to that given
func (o *UploadSampleV3OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the upload sample v3 o k response
func (o *UploadSampleV3OK) Code() int {
	return 200
}

func (o *UploadSampleV3OK) Error() string {
	return fmt.Sprintf("[POST /samples/entities/samples/v3][%d] uploadSampleV3OK  %+v", 200, o.Payload)
}

func (o *UploadSampleV3OK) String() string {
	return fmt.Sprintf("[POST /samples/entities/samples/v3][%d] uploadSampleV3OK  %+v", 200, o.Payload)
}

func (o *UploadSampleV3OK) GetPayload() *models.ClientSampleMetadataResponseV2 {
	return o.Payload
}

func (o *UploadSampleV3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ClientSampleMetadataResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadSampleV3BadRequest creates a UploadSampleV3BadRequest with default headers values
func NewUploadSampleV3BadRequest() *UploadSampleV3BadRequest {
	return &UploadSampleV3BadRequest{}
}

/*
UploadSampleV3BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UploadSampleV3BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ClientSampleMetadataResponseV2
}

// IsSuccess returns true when this upload sample v3 bad request response has a 2xx status code
func (o *UploadSampleV3BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload sample v3 bad request response has a 3xx status code
func (o *UploadSampleV3BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload sample v3 bad request response has a 4xx status code
func (o *UploadSampleV3BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload sample v3 bad request response has a 5xx status code
func (o *UploadSampleV3BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this upload sample v3 bad request response a status code equal to that given
func (o *UploadSampleV3BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the upload sample v3 bad request response
func (o *UploadSampleV3BadRequest) Code() int {
	return 400
}

func (o *UploadSampleV3BadRequest) Error() string {
	return fmt.Sprintf("[POST /samples/entities/samples/v3][%d] uploadSampleV3BadRequest  %+v", 400, o.Payload)
}

func (o *UploadSampleV3BadRequest) String() string {
	return fmt.Sprintf("[POST /samples/entities/samples/v3][%d] uploadSampleV3BadRequest  %+v", 400, o.Payload)
}

func (o *UploadSampleV3BadRequest) GetPayload() *models.ClientSampleMetadataResponseV2 {
	return o.Payload
}

func (o *UploadSampleV3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ClientSampleMetadataResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadSampleV3Forbidden creates a UploadSampleV3Forbidden with default headers values
func NewUploadSampleV3Forbidden() *UploadSampleV3Forbidden {
	return &UploadSampleV3Forbidden{}
}

/*
UploadSampleV3Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UploadSampleV3Forbidden struct {

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

// IsSuccess returns true when this upload sample v3 forbidden response has a 2xx status code
func (o *UploadSampleV3Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload sample v3 forbidden response has a 3xx status code
func (o *UploadSampleV3Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload sample v3 forbidden response has a 4xx status code
func (o *UploadSampleV3Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload sample v3 forbidden response has a 5xx status code
func (o *UploadSampleV3Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this upload sample v3 forbidden response a status code equal to that given
func (o *UploadSampleV3Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the upload sample v3 forbidden response
func (o *UploadSampleV3Forbidden) Code() int {
	return 403
}

func (o *UploadSampleV3Forbidden) Error() string {
	return fmt.Sprintf("[POST /samples/entities/samples/v3][%d] uploadSampleV3Forbidden  %+v", 403, o.Payload)
}

func (o *UploadSampleV3Forbidden) String() string {
	return fmt.Sprintf("[POST /samples/entities/samples/v3][%d] uploadSampleV3Forbidden  %+v", 403, o.Payload)
}

func (o *UploadSampleV3Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UploadSampleV3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUploadSampleV3TooManyRequests creates a UploadSampleV3TooManyRequests with default headers values
func NewUploadSampleV3TooManyRequests() *UploadSampleV3TooManyRequests {
	return &UploadSampleV3TooManyRequests{}
}

/*
UploadSampleV3TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UploadSampleV3TooManyRequests struct {

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

// IsSuccess returns true when this upload sample v3 too many requests response has a 2xx status code
func (o *UploadSampleV3TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload sample v3 too many requests response has a 3xx status code
func (o *UploadSampleV3TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload sample v3 too many requests response has a 4xx status code
func (o *UploadSampleV3TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload sample v3 too many requests response has a 5xx status code
func (o *UploadSampleV3TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this upload sample v3 too many requests response a status code equal to that given
func (o *UploadSampleV3TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the upload sample v3 too many requests response
func (o *UploadSampleV3TooManyRequests) Code() int {
	return 429
}

func (o *UploadSampleV3TooManyRequests) Error() string {
	return fmt.Sprintf("[POST /samples/entities/samples/v3][%d] uploadSampleV3TooManyRequests  %+v", 429, o.Payload)
}

func (o *UploadSampleV3TooManyRequests) String() string {
	return fmt.Sprintf("[POST /samples/entities/samples/v3][%d] uploadSampleV3TooManyRequests  %+v", 429, o.Payload)
}

func (o *UploadSampleV3TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UploadSampleV3TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUploadSampleV3InternalServerError creates a UploadSampleV3InternalServerError with default headers values
func NewUploadSampleV3InternalServerError() *UploadSampleV3InternalServerError {
	return &UploadSampleV3InternalServerError{}
}

/*
UploadSampleV3InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UploadSampleV3InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ClientSampleMetadataResponseV2
}

// IsSuccess returns true when this upload sample v3 internal server error response has a 2xx status code
func (o *UploadSampleV3InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload sample v3 internal server error response has a 3xx status code
func (o *UploadSampleV3InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload sample v3 internal server error response has a 4xx status code
func (o *UploadSampleV3InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload sample v3 internal server error response has a 5xx status code
func (o *UploadSampleV3InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this upload sample v3 internal server error response a status code equal to that given
func (o *UploadSampleV3InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the upload sample v3 internal server error response
func (o *UploadSampleV3InternalServerError) Code() int {
	return 500
}

func (o *UploadSampleV3InternalServerError) Error() string {
	return fmt.Sprintf("[POST /samples/entities/samples/v3][%d] uploadSampleV3InternalServerError  %+v", 500, o.Payload)
}

func (o *UploadSampleV3InternalServerError) String() string {
	return fmt.Sprintf("[POST /samples/entities/samples/v3][%d] uploadSampleV3InternalServerError  %+v", 500, o.Payload)
}

func (o *UploadSampleV3InternalServerError) GetPayload() *models.ClientSampleMetadataResponseV2 {
	return o.Payload
}

func (o *UploadSampleV3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ClientSampleMetadataResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
