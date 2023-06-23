// Code generated by go-swagger; DO NOT EDIT.

package cspm_registration

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

// GetConfigurationDetectionIDsV2Reader is a Reader for the GetConfigurationDetectionIDsV2 structure.
type GetConfigurationDetectionIDsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConfigurationDetectionIDsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConfigurationDetectionIDsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConfigurationDetectionIDsV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConfigurationDetectionIDsV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConfigurationDetectionIDsV2TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConfigurationDetectionIDsV2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /detects/queries/iom/v2] GetConfigurationDetectionIDsV2", response, response.Code())
	}
}

// NewGetConfigurationDetectionIDsV2OK creates a GetConfigurationDetectionIDsV2OK with default headers values
func NewGetConfigurationDetectionIDsV2OK() *GetConfigurationDetectionIDsV2OK {
	return &GetConfigurationDetectionIDsV2OK{}
}

/*
GetConfigurationDetectionIDsV2OK describes a response with status code 200, with default header values.

OK
*/
type GetConfigurationDetectionIDsV2OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationIOMEventIDsResponseV2
}

// IsSuccess returns true when this get configuration detection i ds v2 o k response has a 2xx status code
func (o *GetConfigurationDetectionIDsV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get configuration detection i ds v2 o k response has a 3xx status code
func (o *GetConfigurationDetectionIDsV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configuration detection i ds v2 o k response has a 4xx status code
func (o *GetConfigurationDetectionIDsV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get configuration detection i ds v2 o k response has a 5xx status code
func (o *GetConfigurationDetectionIDsV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get configuration detection i ds v2 o k response a status code equal to that given
func (o *GetConfigurationDetectionIDsV2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get configuration detection i ds v2 o k response
func (o *GetConfigurationDetectionIDsV2OK) Code() int {
	return 200
}

func (o *GetConfigurationDetectionIDsV2OK) Error() string {
	return fmt.Sprintf("[GET /detects/queries/iom/v2][%d] getConfigurationDetectionIDsV2OK  %+v", 200, o.Payload)
}

func (o *GetConfigurationDetectionIDsV2OK) String() string {
	return fmt.Sprintf("[GET /detects/queries/iom/v2][%d] getConfigurationDetectionIDsV2OK  %+v", 200, o.Payload)
}

func (o *GetConfigurationDetectionIDsV2OK) GetPayload() *models.RegistrationIOMEventIDsResponseV2 {
	return o.Payload
}

func (o *GetConfigurationDetectionIDsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationIOMEventIDsResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationDetectionIDsV2BadRequest creates a GetConfigurationDetectionIDsV2BadRequest with default headers values
func NewGetConfigurationDetectionIDsV2BadRequest() *GetConfigurationDetectionIDsV2BadRequest {
	return &GetConfigurationDetectionIDsV2BadRequest{}
}

/*
GetConfigurationDetectionIDsV2BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetConfigurationDetectionIDsV2BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecMetaInfo
}

// IsSuccess returns true when this get configuration detection i ds v2 bad request response has a 2xx status code
func (o *GetConfigurationDetectionIDsV2BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get configuration detection i ds v2 bad request response has a 3xx status code
func (o *GetConfigurationDetectionIDsV2BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configuration detection i ds v2 bad request response has a 4xx status code
func (o *GetConfigurationDetectionIDsV2BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get configuration detection i ds v2 bad request response has a 5xx status code
func (o *GetConfigurationDetectionIDsV2BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get configuration detection i ds v2 bad request response a status code equal to that given
func (o *GetConfigurationDetectionIDsV2BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get configuration detection i ds v2 bad request response
func (o *GetConfigurationDetectionIDsV2BadRequest) Code() int {
	return 400
}

func (o *GetConfigurationDetectionIDsV2BadRequest) Error() string {
	return fmt.Sprintf("[GET /detects/queries/iom/v2][%d] getConfigurationDetectionIDsV2BadRequest  %+v", 400, o.Payload)
}

func (o *GetConfigurationDetectionIDsV2BadRequest) String() string {
	return fmt.Sprintf("[GET /detects/queries/iom/v2][%d] getConfigurationDetectionIDsV2BadRequest  %+v", 400, o.Payload)
}

func (o *GetConfigurationDetectionIDsV2BadRequest) GetPayload() *models.MsaspecMetaInfo {
	return o.Payload
}

func (o *GetConfigurationDetectionIDsV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecMetaInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationDetectionIDsV2Forbidden creates a GetConfigurationDetectionIDsV2Forbidden with default headers values
func NewGetConfigurationDetectionIDsV2Forbidden() *GetConfigurationDetectionIDsV2Forbidden {
	return &GetConfigurationDetectionIDsV2Forbidden{}
}

/*
GetConfigurationDetectionIDsV2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetConfigurationDetectionIDsV2Forbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecMetaInfo
}

// IsSuccess returns true when this get configuration detection i ds v2 forbidden response has a 2xx status code
func (o *GetConfigurationDetectionIDsV2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get configuration detection i ds v2 forbidden response has a 3xx status code
func (o *GetConfigurationDetectionIDsV2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configuration detection i ds v2 forbidden response has a 4xx status code
func (o *GetConfigurationDetectionIDsV2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get configuration detection i ds v2 forbidden response has a 5xx status code
func (o *GetConfigurationDetectionIDsV2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get configuration detection i ds v2 forbidden response a status code equal to that given
func (o *GetConfigurationDetectionIDsV2Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get configuration detection i ds v2 forbidden response
func (o *GetConfigurationDetectionIDsV2Forbidden) Code() int {
	return 403
}

func (o *GetConfigurationDetectionIDsV2Forbidden) Error() string {
	return fmt.Sprintf("[GET /detects/queries/iom/v2][%d] getConfigurationDetectionIDsV2Forbidden  %+v", 403, o.Payload)
}

func (o *GetConfigurationDetectionIDsV2Forbidden) String() string {
	return fmt.Sprintf("[GET /detects/queries/iom/v2][%d] getConfigurationDetectionIDsV2Forbidden  %+v", 403, o.Payload)
}

func (o *GetConfigurationDetectionIDsV2Forbidden) GetPayload() *models.MsaspecMetaInfo {
	return o.Payload
}

func (o *GetConfigurationDetectionIDsV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecMetaInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationDetectionIDsV2TooManyRequests creates a GetConfigurationDetectionIDsV2TooManyRequests with default headers values
func NewGetConfigurationDetectionIDsV2TooManyRequests() *GetConfigurationDetectionIDsV2TooManyRequests {
	return &GetConfigurationDetectionIDsV2TooManyRequests{}
}

/*
GetConfigurationDetectionIDsV2TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetConfigurationDetectionIDsV2TooManyRequests struct {

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

// IsSuccess returns true when this get configuration detection i ds v2 too many requests response has a 2xx status code
func (o *GetConfigurationDetectionIDsV2TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get configuration detection i ds v2 too many requests response has a 3xx status code
func (o *GetConfigurationDetectionIDsV2TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configuration detection i ds v2 too many requests response has a 4xx status code
func (o *GetConfigurationDetectionIDsV2TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get configuration detection i ds v2 too many requests response has a 5xx status code
func (o *GetConfigurationDetectionIDsV2TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get configuration detection i ds v2 too many requests response a status code equal to that given
func (o *GetConfigurationDetectionIDsV2TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get configuration detection i ds v2 too many requests response
func (o *GetConfigurationDetectionIDsV2TooManyRequests) Code() int {
	return 429
}

func (o *GetConfigurationDetectionIDsV2TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /detects/queries/iom/v2][%d] getConfigurationDetectionIDsV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConfigurationDetectionIDsV2TooManyRequests) String() string {
	return fmt.Sprintf("[GET /detects/queries/iom/v2][%d] getConfigurationDetectionIDsV2TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConfigurationDetectionIDsV2TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetConfigurationDetectionIDsV2TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetConfigurationDetectionIDsV2InternalServerError creates a GetConfigurationDetectionIDsV2InternalServerError with default headers values
func NewGetConfigurationDetectionIDsV2InternalServerError() *GetConfigurationDetectionIDsV2InternalServerError {
	return &GetConfigurationDetectionIDsV2InternalServerError{}
}

/*
GetConfigurationDetectionIDsV2InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetConfigurationDetectionIDsV2InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecMetaInfo
}

// IsSuccess returns true when this get configuration detection i ds v2 internal server error response has a 2xx status code
func (o *GetConfigurationDetectionIDsV2InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get configuration detection i ds v2 internal server error response has a 3xx status code
func (o *GetConfigurationDetectionIDsV2InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configuration detection i ds v2 internal server error response has a 4xx status code
func (o *GetConfigurationDetectionIDsV2InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get configuration detection i ds v2 internal server error response has a 5xx status code
func (o *GetConfigurationDetectionIDsV2InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get configuration detection i ds v2 internal server error response a status code equal to that given
func (o *GetConfigurationDetectionIDsV2InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get configuration detection i ds v2 internal server error response
func (o *GetConfigurationDetectionIDsV2InternalServerError) Code() int {
	return 500
}

func (o *GetConfigurationDetectionIDsV2InternalServerError) Error() string {
	return fmt.Sprintf("[GET /detects/queries/iom/v2][%d] getConfigurationDetectionIDsV2InternalServerError  %+v", 500, o.Payload)
}

func (o *GetConfigurationDetectionIDsV2InternalServerError) String() string {
	return fmt.Sprintf("[GET /detects/queries/iom/v2][%d] getConfigurationDetectionIDsV2InternalServerError  %+v", 500, o.Payload)
}

func (o *GetConfigurationDetectionIDsV2InternalServerError) GetPayload() *models.MsaspecMetaInfo {
	return o.Payload
}

func (o *GetConfigurationDetectionIDsV2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecMetaInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
