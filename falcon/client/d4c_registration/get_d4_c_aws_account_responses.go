// Code generated by go-swagger; DO NOT EDIT.

package d4c_registration

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

// GetD4CAWSAccountReader is a Reader for the GetD4CAWSAccount structure.
type GetD4CAWSAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetD4CAWSAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetD4CAWSAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewGetD4CAWSAccountMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetD4CAWSAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetD4CAWSAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetD4CAWSAccountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetD4CAWSAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /cloud-connect-aws/entities/account/v2] GetD4CAwsAccount", response, response.Code())
	}
}

// NewGetD4CAWSAccountOK creates a GetD4CAWSAccountOK with default headers values
func NewGetD4CAWSAccountOK() *GetD4CAWSAccountOK {
	return &GetD4CAWSAccountOK{}
}

/*
GetD4CAWSAccountOK describes a response with status code 200, with default header values.

OK
*/
type GetD4CAWSAccountOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAWSAccountResponseV2
}

// IsSuccess returns true when this get d4 c Aws account o k response has a 2xx status code
func (o *GetD4CAWSAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get d4 c Aws account o k response has a 3xx status code
func (o *GetD4CAWSAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d4 c Aws account o k response has a 4xx status code
func (o *GetD4CAWSAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get d4 c Aws account o k response has a 5xx status code
func (o *GetD4CAWSAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get d4 c Aws account o k response a status code equal to that given
func (o *GetD4CAWSAccountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get d4 c Aws account o k response
func (o *GetD4CAWSAccountOK) Code() int {
	return 200
}

func (o *GetD4CAWSAccountOK) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/entities/account/v2][%d] getD4CAwsAccountOK  %+v", 200, o.Payload)
}

func (o *GetD4CAWSAccountOK) String() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/entities/account/v2][%d] getD4CAwsAccountOK  %+v", 200, o.Payload)
}

func (o *GetD4CAWSAccountOK) GetPayload() *models.RegistrationAWSAccountResponseV2 {
	return o.Payload
}

func (o *GetD4CAWSAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAWSAccountResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetD4CAWSAccountMultiStatus creates a GetD4CAWSAccountMultiStatus with default headers values
func NewGetD4CAWSAccountMultiStatus() *GetD4CAWSAccountMultiStatus {
	return &GetD4CAWSAccountMultiStatus{}
}

/*
GetD4CAWSAccountMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type GetD4CAWSAccountMultiStatus struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAWSAccountResponseV2
}

// IsSuccess returns true when this get d4 c Aws account multi status response has a 2xx status code
func (o *GetD4CAWSAccountMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get d4 c Aws account multi status response has a 3xx status code
func (o *GetD4CAWSAccountMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d4 c Aws account multi status response has a 4xx status code
func (o *GetD4CAWSAccountMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this get d4 c Aws account multi status response has a 5xx status code
func (o *GetD4CAWSAccountMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this get d4 c Aws account multi status response a status code equal to that given
func (o *GetD4CAWSAccountMultiStatus) IsCode(code int) bool {
	return code == 207
}

// Code gets the status code for the get d4 c Aws account multi status response
func (o *GetD4CAWSAccountMultiStatus) Code() int {
	return 207
}

func (o *GetD4CAWSAccountMultiStatus) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/entities/account/v2][%d] getD4CAwsAccountMultiStatus  %+v", 207, o.Payload)
}

func (o *GetD4CAWSAccountMultiStatus) String() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/entities/account/v2][%d] getD4CAwsAccountMultiStatus  %+v", 207, o.Payload)
}

func (o *GetD4CAWSAccountMultiStatus) GetPayload() *models.RegistrationAWSAccountResponseV2 {
	return o.Payload
}

func (o *GetD4CAWSAccountMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAWSAccountResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetD4CAWSAccountBadRequest creates a GetD4CAWSAccountBadRequest with default headers values
func NewGetD4CAWSAccountBadRequest() *GetD4CAWSAccountBadRequest {
	return &GetD4CAWSAccountBadRequest{}
}

/*
GetD4CAWSAccountBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetD4CAWSAccountBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAWSAccountResponseV2
}

// IsSuccess returns true when this get d4 c Aws account bad request response has a 2xx status code
func (o *GetD4CAWSAccountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get d4 c Aws account bad request response has a 3xx status code
func (o *GetD4CAWSAccountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d4 c Aws account bad request response has a 4xx status code
func (o *GetD4CAWSAccountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get d4 c Aws account bad request response has a 5xx status code
func (o *GetD4CAWSAccountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get d4 c Aws account bad request response a status code equal to that given
func (o *GetD4CAWSAccountBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get d4 c Aws account bad request response
func (o *GetD4CAWSAccountBadRequest) Code() int {
	return 400
}

func (o *GetD4CAWSAccountBadRequest) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/entities/account/v2][%d] getD4CAwsAccountBadRequest  %+v", 400, o.Payload)
}

func (o *GetD4CAWSAccountBadRequest) String() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/entities/account/v2][%d] getD4CAwsAccountBadRequest  %+v", 400, o.Payload)
}

func (o *GetD4CAWSAccountBadRequest) GetPayload() *models.RegistrationAWSAccountResponseV2 {
	return o.Payload
}

func (o *GetD4CAWSAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAWSAccountResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetD4CAWSAccountForbidden creates a GetD4CAWSAccountForbidden with default headers values
func NewGetD4CAWSAccountForbidden() *GetD4CAWSAccountForbidden {
	return &GetD4CAWSAccountForbidden{}
}

/*
GetD4CAWSAccountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetD4CAWSAccountForbidden struct {

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

// IsSuccess returns true when this get d4 c Aws account forbidden response has a 2xx status code
func (o *GetD4CAWSAccountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get d4 c Aws account forbidden response has a 3xx status code
func (o *GetD4CAWSAccountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d4 c Aws account forbidden response has a 4xx status code
func (o *GetD4CAWSAccountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get d4 c Aws account forbidden response has a 5xx status code
func (o *GetD4CAWSAccountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get d4 c Aws account forbidden response a status code equal to that given
func (o *GetD4CAWSAccountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get d4 c Aws account forbidden response
func (o *GetD4CAWSAccountForbidden) Code() int {
	return 403
}

func (o *GetD4CAWSAccountForbidden) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/entities/account/v2][%d] getD4CAwsAccountForbidden  %+v", 403, o.Payload)
}

func (o *GetD4CAWSAccountForbidden) String() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/entities/account/v2][%d] getD4CAwsAccountForbidden  %+v", 403, o.Payload)
}

func (o *GetD4CAWSAccountForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetD4CAWSAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetD4CAWSAccountTooManyRequests creates a GetD4CAWSAccountTooManyRequests with default headers values
func NewGetD4CAWSAccountTooManyRequests() *GetD4CAWSAccountTooManyRequests {
	return &GetD4CAWSAccountTooManyRequests{}
}

/*
GetD4CAWSAccountTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetD4CAWSAccountTooManyRequests struct {

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

// IsSuccess returns true when this get d4 c Aws account too many requests response has a 2xx status code
func (o *GetD4CAWSAccountTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get d4 c Aws account too many requests response has a 3xx status code
func (o *GetD4CAWSAccountTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d4 c Aws account too many requests response has a 4xx status code
func (o *GetD4CAWSAccountTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get d4 c Aws account too many requests response has a 5xx status code
func (o *GetD4CAWSAccountTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get d4 c Aws account too many requests response a status code equal to that given
func (o *GetD4CAWSAccountTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get d4 c Aws account too many requests response
func (o *GetD4CAWSAccountTooManyRequests) Code() int {
	return 429
}

func (o *GetD4CAWSAccountTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/entities/account/v2][%d] getD4CAwsAccountTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetD4CAWSAccountTooManyRequests) String() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/entities/account/v2][%d] getD4CAwsAccountTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetD4CAWSAccountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetD4CAWSAccountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetD4CAWSAccountInternalServerError creates a GetD4CAWSAccountInternalServerError with default headers values
func NewGetD4CAWSAccountInternalServerError() *GetD4CAWSAccountInternalServerError {
	return &GetD4CAWSAccountInternalServerError{}
}

/*
GetD4CAWSAccountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetD4CAWSAccountInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAWSAccountResponseV2
}

// IsSuccess returns true when this get d4 c Aws account internal server error response has a 2xx status code
func (o *GetD4CAWSAccountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get d4 c Aws account internal server error response has a 3xx status code
func (o *GetD4CAWSAccountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get d4 c Aws account internal server error response has a 4xx status code
func (o *GetD4CAWSAccountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get d4 c Aws account internal server error response has a 5xx status code
func (o *GetD4CAWSAccountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get d4 c Aws account internal server error response a status code equal to that given
func (o *GetD4CAWSAccountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get d4 c Aws account internal server error response
func (o *GetD4CAWSAccountInternalServerError) Code() int {
	return 500
}

func (o *GetD4CAWSAccountInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/entities/account/v2][%d] getD4CAwsAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *GetD4CAWSAccountInternalServerError) String() string {
	return fmt.Sprintf("[GET /cloud-connect-aws/entities/account/v2][%d] getD4CAwsAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *GetD4CAWSAccountInternalServerError) GetPayload() *models.RegistrationAWSAccountResponseV2 {
	return o.Payload
}

func (o *GetD4CAWSAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAWSAccountResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
