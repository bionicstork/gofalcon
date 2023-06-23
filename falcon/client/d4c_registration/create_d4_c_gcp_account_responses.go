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

// CreateD4CGCPAccountReader is a Reader for the CreateD4CGCPAccount structure.
type CreateD4CGCPAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateD4CGCPAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateD4CGCPAccountCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewCreateD4CGCPAccountMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateD4CGCPAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateD4CGCPAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateD4CGCPAccountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateD4CGCPAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /cloud-connect-gcp/entities/account/v1] CreateD4CGCPAccount", response, response.Code())
	}
}

// NewCreateD4CGCPAccountCreated creates a CreateD4CGCPAccountCreated with default headers values
func NewCreateD4CGCPAccountCreated() *CreateD4CGCPAccountCreated {
	return &CreateD4CGCPAccountCreated{}
}

/*
CreateD4CGCPAccountCreated describes a response with status code 201, with default header values.

Created
*/
type CreateD4CGCPAccountCreated struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPAccountResponseV1
}

// IsSuccess returns true when this create d4 c Gcp account created response has a 2xx status code
func (o *CreateD4CGCPAccountCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create d4 c Gcp account created response has a 3xx status code
func (o *CreateD4CGCPAccountCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create d4 c Gcp account created response has a 4xx status code
func (o *CreateD4CGCPAccountCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create d4 c Gcp account created response has a 5xx status code
func (o *CreateD4CGCPAccountCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create d4 c Gcp account created response a status code equal to that given
func (o *CreateD4CGCPAccountCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create d4 c Gcp account created response
func (o *CreateD4CGCPAccountCreated) Code() int {
	return 201
}

func (o *CreateD4CGCPAccountCreated) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-gcp/entities/account/v1][%d] createD4CGcpAccountCreated  %+v", 201, o.Payload)
}

func (o *CreateD4CGCPAccountCreated) String() string {
	return fmt.Sprintf("[POST /cloud-connect-gcp/entities/account/v1][%d] createD4CGcpAccountCreated  %+v", 201, o.Payload)
}

func (o *CreateD4CGCPAccountCreated) GetPayload() *models.RegistrationGCPAccountResponseV1 {
	return o.Payload
}

func (o *CreateD4CGCPAccountCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPAccountResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateD4CGCPAccountMultiStatus creates a CreateD4CGCPAccountMultiStatus with default headers values
func NewCreateD4CGCPAccountMultiStatus() *CreateD4CGCPAccountMultiStatus {
	return &CreateD4CGCPAccountMultiStatus{}
}

/*
CreateD4CGCPAccountMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type CreateD4CGCPAccountMultiStatus struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPAccountResponseV1
}

// IsSuccess returns true when this create d4 c Gcp account multi status response has a 2xx status code
func (o *CreateD4CGCPAccountMultiStatus) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create d4 c Gcp account multi status response has a 3xx status code
func (o *CreateD4CGCPAccountMultiStatus) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create d4 c Gcp account multi status response has a 4xx status code
func (o *CreateD4CGCPAccountMultiStatus) IsClientError() bool {
	return false
}

// IsServerError returns true when this create d4 c Gcp account multi status response has a 5xx status code
func (o *CreateD4CGCPAccountMultiStatus) IsServerError() bool {
	return false
}

// IsCode returns true when this create d4 c Gcp account multi status response a status code equal to that given
func (o *CreateD4CGCPAccountMultiStatus) IsCode(code int) bool {
	return code == 207
}

// Code gets the status code for the create d4 c Gcp account multi status response
func (o *CreateD4CGCPAccountMultiStatus) Code() int {
	return 207
}

func (o *CreateD4CGCPAccountMultiStatus) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-gcp/entities/account/v1][%d] createD4CGcpAccountMultiStatus  %+v", 207, o.Payload)
}

func (o *CreateD4CGCPAccountMultiStatus) String() string {
	return fmt.Sprintf("[POST /cloud-connect-gcp/entities/account/v1][%d] createD4CGcpAccountMultiStatus  %+v", 207, o.Payload)
}

func (o *CreateD4CGCPAccountMultiStatus) GetPayload() *models.RegistrationGCPAccountResponseV1 {
	return o.Payload
}

func (o *CreateD4CGCPAccountMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPAccountResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateD4CGCPAccountBadRequest creates a CreateD4CGCPAccountBadRequest with default headers values
func NewCreateD4CGCPAccountBadRequest() *CreateD4CGCPAccountBadRequest {
	return &CreateD4CGCPAccountBadRequest{}
}

/*
CreateD4CGCPAccountBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateD4CGCPAccountBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPAccountResponseV1
}

// IsSuccess returns true when this create d4 c Gcp account bad request response has a 2xx status code
func (o *CreateD4CGCPAccountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create d4 c Gcp account bad request response has a 3xx status code
func (o *CreateD4CGCPAccountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create d4 c Gcp account bad request response has a 4xx status code
func (o *CreateD4CGCPAccountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create d4 c Gcp account bad request response has a 5xx status code
func (o *CreateD4CGCPAccountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create d4 c Gcp account bad request response a status code equal to that given
func (o *CreateD4CGCPAccountBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create d4 c Gcp account bad request response
func (o *CreateD4CGCPAccountBadRequest) Code() int {
	return 400
}

func (o *CreateD4CGCPAccountBadRequest) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-gcp/entities/account/v1][%d] createD4CGcpAccountBadRequest  %+v", 400, o.Payload)
}

func (o *CreateD4CGCPAccountBadRequest) String() string {
	return fmt.Sprintf("[POST /cloud-connect-gcp/entities/account/v1][%d] createD4CGcpAccountBadRequest  %+v", 400, o.Payload)
}

func (o *CreateD4CGCPAccountBadRequest) GetPayload() *models.RegistrationGCPAccountResponseV1 {
	return o.Payload
}

func (o *CreateD4CGCPAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPAccountResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateD4CGCPAccountForbidden creates a CreateD4CGCPAccountForbidden with default headers values
func NewCreateD4CGCPAccountForbidden() *CreateD4CGCPAccountForbidden {
	return &CreateD4CGCPAccountForbidden{}
}

/*
CreateD4CGCPAccountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateD4CGCPAccountForbidden struct {

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

// IsSuccess returns true when this create d4 c Gcp account forbidden response has a 2xx status code
func (o *CreateD4CGCPAccountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create d4 c Gcp account forbidden response has a 3xx status code
func (o *CreateD4CGCPAccountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create d4 c Gcp account forbidden response has a 4xx status code
func (o *CreateD4CGCPAccountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create d4 c Gcp account forbidden response has a 5xx status code
func (o *CreateD4CGCPAccountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create d4 c Gcp account forbidden response a status code equal to that given
func (o *CreateD4CGCPAccountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create d4 c Gcp account forbidden response
func (o *CreateD4CGCPAccountForbidden) Code() int {
	return 403
}

func (o *CreateD4CGCPAccountForbidden) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-gcp/entities/account/v1][%d] createD4CGcpAccountForbidden  %+v", 403, o.Payload)
}

func (o *CreateD4CGCPAccountForbidden) String() string {
	return fmt.Sprintf("[POST /cloud-connect-gcp/entities/account/v1][%d] createD4CGcpAccountForbidden  %+v", 403, o.Payload)
}

func (o *CreateD4CGCPAccountForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CreateD4CGCPAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateD4CGCPAccountTooManyRequests creates a CreateD4CGCPAccountTooManyRequests with default headers values
func NewCreateD4CGCPAccountTooManyRequests() *CreateD4CGCPAccountTooManyRequests {
	return &CreateD4CGCPAccountTooManyRequests{}
}

/*
CreateD4CGCPAccountTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateD4CGCPAccountTooManyRequests struct {

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

// IsSuccess returns true when this create d4 c Gcp account too many requests response has a 2xx status code
func (o *CreateD4CGCPAccountTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create d4 c Gcp account too many requests response has a 3xx status code
func (o *CreateD4CGCPAccountTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create d4 c Gcp account too many requests response has a 4xx status code
func (o *CreateD4CGCPAccountTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create d4 c Gcp account too many requests response has a 5xx status code
func (o *CreateD4CGCPAccountTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create d4 c Gcp account too many requests response a status code equal to that given
func (o *CreateD4CGCPAccountTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the create d4 c Gcp account too many requests response
func (o *CreateD4CGCPAccountTooManyRequests) Code() int {
	return 429
}

func (o *CreateD4CGCPAccountTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-gcp/entities/account/v1][%d] createD4CGcpAccountTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateD4CGCPAccountTooManyRequests) String() string {
	return fmt.Sprintf("[POST /cloud-connect-gcp/entities/account/v1][%d] createD4CGcpAccountTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateD4CGCPAccountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CreateD4CGCPAccountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateD4CGCPAccountInternalServerError creates a CreateD4CGCPAccountInternalServerError with default headers values
func NewCreateD4CGCPAccountInternalServerError() *CreateD4CGCPAccountInternalServerError {
	return &CreateD4CGCPAccountInternalServerError{}
}

/*
CreateD4CGCPAccountInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateD4CGCPAccountInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationGCPAccountResponseV1
}

// IsSuccess returns true when this create d4 c Gcp account internal server error response has a 2xx status code
func (o *CreateD4CGCPAccountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create d4 c Gcp account internal server error response has a 3xx status code
func (o *CreateD4CGCPAccountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create d4 c Gcp account internal server error response has a 4xx status code
func (o *CreateD4CGCPAccountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create d4 c Gcp account internal server error response has a 5xx status code
func (o *CreateD4CGCPAccountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create d4 c Gcp account internal server error response a status code equal to that given
func (o *CreateD4CGCPAccountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create d4 c Gcp account internal server error response
func (o *CreateD4CGCPAccountInternalServerError) Code() int {
	return 500
}

func (o *CreateD4CGCPAccountInternalServerError) Error() string {
	return fmt.Sprintf("[POST /cloud-connect-gcp/entities/account/v1][%d] createD4CGcpAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateD4CGCPAccountInternalServerError) String() string {
	return fmt.Sprintf("[POST /cloud-connect-gcp/entities/account/v1][%d] createD4CGcpAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateD4CGCPAccountInternalServerError) GetPayload() *models.RegistrationGCPAccountResponseV1 {
	return o.Payload
}

func (o *CreateD4CGCPAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationGCPAccountResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
