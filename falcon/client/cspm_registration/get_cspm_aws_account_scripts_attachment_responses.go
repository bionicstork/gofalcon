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

// GetCSPMAWSAccountScriptsAttachmentReader is a Reader for the GetCSPMAWSAccountScriptsAttachment structure.
type GetCSPMAWSAccountScriptsAttachmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCSPMAWSAccountScriptsAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCSPMAWSAccountScriptsAttachmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCSPMAWSAccountScriptsAttachmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCSPMAWSAccountScriptsAttachmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCSPMAWSAccountScriptsAttachmentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCSPMAWSAccountScriptsAttachmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cloud-connect-cspm-aws/entities/user-scripts-download/v1] GetCSPMAwsAccountScriptsAttachment", response, response.Code())
	}
}

// NewGetCSPMAWSAccountScriptsAttachmentOK creates a GetCSPMAWSAccountScriptsAttachmentOK with default headers values
func NewGetCSPMAWSAccountScriptsAttachmentOK() *GetCSPMAWSAccountScriptsAttachmentOK {
	return &GetCSPMAWSAccountScriptsAttachmentOK{}
}

/*
GetCSPMAWSAccountScriptsAttachmentOK describes a response with status code 200, with default header values.

OK
*/
type GetCSPMAWSAccountScriptsAttachmentOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAWSProvisionGetAccountScriptResponseV2
}

// IsSuccess returns true when this get Cspm Aws account scripts attachment o k response has a 2xx status code
func (o *GetCSPMAWSAccountScriptsAttachmentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Cspm Aws account scripts attachment o k response has a 3xx status code
func (o *GetCSPMAWSAccountScriptsAttachmentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Cspm Aws account scripts attachment o k response has a 4xx status code
func (o *GetCSPMAWSAccountScriptsAttachmentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Cspm Aws account scripts attachment o k response has a 5xx status code
func (o *GetCSPMAWSAccountScriptsAttachmentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Cspm Aws account scripts attachment o k response a status code equal to that given
func (o *GetCSPMAWSAccountScriptsAttachmentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Cspm Aws account scripts attachment o k response
func (o *GetCSPMAWSAccountScriptsAttachmentOK) Code() int {
	return 200
}

func (o *GetCSPMAWSAccountScriptsAttachmentOK) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-aws/entities/user-scripts-download/v1][%d] getCspmAwsAccountScriptsAttachmentOK  %+v", 200, o.Payload)
}

func (o *GetCSPMAWSAccountScriptsAttachmentOK) String() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-aws/entities/user-scripts-download/v1][%d] getCspmAwsAccountScriptsAttachmentOK  %+v", 200, o.Payload)
}

func (o *GetCSPMAWSAccountScriptsAttachmentOK) GetPayload() *models.RegistrationAWSProvisionGetAccountScriptResponseV2 {
	return o.Payload
}

func (o *GetCSPMAWSAccountScriptsAttachmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAWSProvisionGetAccountScriptResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMAWSAccountScriptsAttachmentBadRequest creates a GetCSPMAWSAccountScriptsAttachmentBadRequest with default headers values
func NewGetCSPMAWSAccountScriptsAttachmentBadRequest() *GetCSPMAWSAccountScriptsAttachmentBadRequest {
	return &GetCSPMAWSAccountScriptsAttachmentBadRequest{}
}

/*
GetCSPMAWSAccountScriptsAttachmentBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetCSPMAWSAccountScriptsAttachmentBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAWSProvisionGetAccountScriptResponseV2
}

// IsSuccess returns true when this get Cspm Aws account scripts attachment bad request response has a 2xx status code
func (o *GetCSPMAWSAccountScriptsAttachmentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Cspm Aws account scripts attachment bad request response has a 3xx status code
func (o *GetCSPMAWSAccountScriptsAttachmentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Cspm Aws account scripts attachment bad request response has a 4xx status code
func (o *GetCSPMAWSAccountScriptsAttachmentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Cspm Aws account scripts attachment bad request response has a 5xx status code
func (o *GetCSPMAWSAccountScriptsAttachmentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get Cspm Aws account scripts attachment bad request response a status code equal to that given
func (o *GetCSPMAWSAccountScriptsAttachmentBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get Cspm Aws account scripts attachment bad request response
func (o *GetCSPMAWSAccountScriptsAttachmentBadRequest) Code() int {
	return 400
}

func (o *GetCSPMAWSAccountScriptsAttachmentBadRequest) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-aws/entities/user-scripts-download/v1][%d] getCspmAwsAccountScriptsAttachmentBadRequest  %+v", 400, o.Payload)
}

func (o *GetCSPMAWSAccountScriptsAttachmentBadRequest) String() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-aws/entities/user-scripts-download/v1][%d] getCspmAwsAccountScriptsAttachmentBadRequest  %+v", 400, o.Payload)
}

func (o *GetCSPMAWSAccountScriptsAttachmentBadRequest) GetPayload() *models.RegistrationAWSProvisionGetAccountScriptResponseV2 {
	return o.Payload
}

func (o *GetCSPMAWSAccountScriptsAttachmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAWSProvisionGetAccountScriptResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCSPMAWSAccountScriptsAttachmentForbidden creates a GetCSPMAWSAccountScriptsAttachmentForbidden with default headers values
func NewGetCSPMAWSAccountScriptsAttachmentForbidden() *GetCSPMAWSAccountScriptsAttachmentForbidden {
	return &GetCSPMAWSAccountScriptsAttachmentForbidden{}
}

/*
GetCSPMAWSAccountScriptsAttachmentForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCSPMAWSAccountScriptsAttachmentForbidden struct {

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

// IsSuccess returns true when this get Cspm Aws account scripts attachment forbidden response has a 2xx status code
func (o *GetCSPMAWSAccountScriptsAttachmentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Cspm Aws account scripts attachment forbidden response has a 3xx status code
func (o *GetCSPMAWSAccountScriptsAttachmentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Cspm Aws account scripts attachment forbidden response has a 4xx status code
func (o *GetCSPMAWSAccountScriptsAttachmentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Cspm Aws account scripts attachment forbidden response has a 5xx status code
func (o *GetCSPMAWSAccountScriptsAttachmentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get Cspm Aws account scripts attachment forbidden response a status code equal to that given
func (o *GetCSPMAWSAccountScriptsAttachmentForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get Cspm Aws account scripts attachment forbidden response
func (o *GetCSPMAWSAccountScriptsAttachmentForbidden) Code() int {
	return 403
}

func (o *GetCSPMAWSAccountScriptsAttachmentForbidden) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-aws/entities/user-scripts-download/v1][%d] getCspmAwsAccountScriptsAttachmentForbidden  %+v", 403, o.Payload)
}

func (o *GetCSPMAWSAccountScriptsAttachmentForbidden) String() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-aws/entities/user-scripts-download/v1][%d] getCspmAwsAccountScriptsAttachmentForbidden  %+v", 403, o.Payload)
}

func (o *GetCSPMAWSAccountScriptsAttachmentForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCSPMAWSAccountScriptsAttachmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCSPMAWSAccountScriptsAttachmentTooManyRequests creates a GetCSPMAWSAccountScriptsAttachmentTooManyRequests with default headers values
func NewGetCSPMAWSAccountScriptsAttachmentTooManyRequests() *GetCSPMAWSAccountScriptsAttachmentTooManyRequests {
	return &GetCSPMAWSAccountScriptsAttachmentTooManyRequests{}
}

/*
GetCSPMAWSAccountScriptsAttachmentTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetCSPMAWSAccountScriptsAttachmentTooManyRequests struct {

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

// IsSuccess returns true when this get Cspm Aws account scripts attachment too many requests response has a 2xx status code
func (o *GetCSPMAWSAccountScriptsAttachmentTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Cspm Aws account scripts attachment too many requests response has a 3xx status code
func (o *GetCSPMAWSAccountScriptsAttachmentTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Cspm Aws account scripts attachment too many requests response has a 4xx status code
func (o *GetCSPMAWSAccountScriptsAttachmentTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Cspm Aws account scripts attachment too many requests response has a 5xx status code
func (o *GetCSPMAWSAccountScriptsAttachmentTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get Cspm Aws account scripts attachment too many requests response a status code equal to that given
func (o *GetCSPMAWSAccountScriptsAttachmentTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get Cspm Aws account scripts attachment too many requests response
func (o *GetCSPMAWSAccountScriptsAttachmentTooManyRequests) Code() int {
	return 429
}

func (o *GetCSPMAWSAccountScriptsAttachmentTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-aws/entities/user-scripts-download/v1][%d] getCspmAwsAccountScriptsAttachmentTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCSPMAWSAccountScriptsAttachmentTooManyRequests) String() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-aws/entities/user-scripts-download/v1][%d] getCspmAwsAccountScriptsAttachmentTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetCSPMAWSAccountScriptsAttachmentTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCSPMAWSAccountScriptsAttachmentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCSPMAWSAccountScriptsAttachmentInternalServerError creates a GetCSPMAWSAccountScriptsAttachmentInternalServerError with default headers values
func NewGetCSPMAWSAccountScriptsAttachmentInternalServerError() *GetCSPMAWSAccountScriptsAttachmentInternalServerError {
	return &GetCSPMAWSAccountScriptsAttachmentInternalServerError{}
}

/*
GetCSPMAWSAccountScriptsAttachmentInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetCSPMAWSAccountScriptsAttachmentInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAWSProvisionGetAccountScriptResponseV2
}

// IsSuccess returns true when this get Cspm Aws account scripts attachment internal server error response has a 2xx status code
func (o *GetCSPMAWSAccountScriptsAttachmentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Cspm Aws account scripts attachment internal server error response has a 3xx status code
func (o *GetCSPMAWSAccountScriptsAttachmentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Cspm Aws account scripts attachment internal server error response has a 4xx status code
func (o *GetCSPMAWSAccountScriptsAttachmentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Cspm Aws account scripts attachment internal server error response has a 5xx status code
func (o *GetCSPMAWSAccountScriptsAttachmentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get Cspm Aws account scripts attachment internal server error response a status code equal to that given
func (o *GetCSPMAWSAccountScriptsAttachmentInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get Cspm Aws account scripts attachment internal server error response
func (o *GetCSPMAWSAccountScriptsAttachmentInternalServerError) Code() int {
	return 500
}

func (o *GetCSPMAWSAccountScriptsAttachmentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-aws/entities/user-scripts-download/v1][%d] getCspmAwsAccountScriptsAttachmentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCSPMAWSAccountScriptsAttachmentInternalServerError) String() string {
	return fmt.Sprintf("[GET /cloud-connect-cspm-aws/entities/user-scripts-download/v1][%d] getCspmAwsAccountScriptsAttachmentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCSPMAWSAccountScriptsAttachmentInternalServerError) GetPayload() *models.RegistrationAWSProvisionGetAccountScriptResponseV2 {
	return o.Payload
}

func (o *GetCSPMAWSAccountScriptsAttachmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAWSProvisionGetAccountScriptResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
