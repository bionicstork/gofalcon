// Code generated by go-swagger; DO NOT EDIT.

package falcon_container_image

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

// DeleteRegistryEntitiesReader is a Reader for the DeleteRegistryEntities structure.
type DeleteRegistryEntitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRegistryEntitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRegistryEntitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteRegistryEntitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteRegistryEntitiesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /container-security/entities/registries/v1] DeleteRegistryEntities", response, response.Code())
	}
}

// NewDeleteRegistryEntitiesOK creates a DeleteRegistryEntitiesOK with default headers values
func NewDeleteRegistryEntitiesOK() *DeleteRegistryEntitiesOK {
	return &DeleteRegistryEntitiesOK{}
}

/*
DeleteRegistryEntitiesOK describes a response with status code 200, with default header values.

OK
*/
type DeleteRegistryEntitiesOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainExternalRegistryListResponse
}

// IsSuccess returns true when this delete registry entities o k response has a 2xx status code
func (o *DeleteRegistryEntitiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete registry entities o k response has a 3xx status code
func (o *DeleteRegistryEntitiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete registry entities o k response has a 4xx status code
func (o *DeleteRegistryEntitiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete registry entities o k response has a 5xx status code
func (o *DeleteRegistryEntitiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete registry entities o k response a status code equal to that given
func (o *DeleteRegistryEntitiesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete registry entities o k response
func (o *DeleteRegistryEntitiesOK) Code() int {
	return 200
}

func (o *DeleteRegistryEntitiesOK) Error() string {
	return fmt.Sprintf("[DELETE /container-security/entities/registries/v1][%d] deleteRegistryEntitiesOK  %+v", 200, o.Payload)
}

func (o *DeleteRegistryEntitiesOK) String() string {
	return fmt.Sprintf("[DELETE /container-security/entities/registries/v1][%d] deleteRegistryEntitiesOK  %+v", 200, o.Payload)
}

func (o *DeleteRegistryEntitiesOK) GetPayload() *models.DomainExternalRegistryListResponse {
	return o.Payload
}

func (o *DeleteRegistryEntitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainExternalRegistryListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRegistryEntitiesForbidden creates a DeleteRegistryEntitiesForbidden with default headers values
func NewDeleteRegistryEntitiesForbidden() *DeleteRegistryEntitiesForbidden {
	return &DeleteRegistryEntitiesForbidden{}
}

/*
DeleteRegistryEntitiesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteRegistryEntitiesForbidden struct {

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

// IsSuccess returns true when this delete registry entities forbidden response has a 2xx status code
func (o *DeleteRegistryEntitiesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete registry entities forbidden response has a 3xx status code
func (o *DeleteRegistryEntitiesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete registry entities forbidden response has a 4xx status code
func (o *DeleteRegistryEntitiesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete registry entities forbidden response has a 5xx status code
func (o *DeleteRegistryEntitiesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete registry entities forbidden response a status code equal to that given
func (o *DeleteRegistryEntitiesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete registry entities forbidden response
func (o *DeleteRegistryEntitiesForbidden) Code() int {
	return 403
}

func (o *DeleteRegistryEntitiesForbidden) Error() string {
	return fmt.Sprintf("[DELETE /container-security/entities/registries/v1][%d] deleteRegistryEntitiesForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRegistryEntitiesForbidden) String() string {
	return fmt.Sprintf("[DELETE /container-security/entities/registries/v1][%d] deleteRegistryEntitiesForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRegistryEntitiesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteRegistryEntitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRegistryEntitiesTooManyRequests creates a DeleteRegistryEntitiesTooManyRequests with default headers values
func NewDeleteRegistryEntitiesTooManyRequests() *DeleteRegistryEntitiesTooManyRequests {
	return &DeleteRegistryEntitiesTooManyRequests{}
}

/*
DeleteRegistryEntitiesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteRegistryEntitiesTooManyRequests struct {

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

// IsSuccess returns true when this delete registry entities too many requests response has a 2xx status code
func (o *DeleteRegistryEntitiesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete registry entities too many requests response has a 3xx status code
func (o *DeleteRegistryEntitiesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete registry entities too many requests response has a 4xx status code
func (o *DeleteRegistryEntitiesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete registry entities too many requests response has a 5xx status code
func (o *DeleteRegistryEntitiesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete registry entities too many requests response a status code equal to that given
func (o *DeleteRegistryEntitiesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the delete registry entities too many requests response
func (o *DeleteRegistryEntitiesTooManyRequests) Code() int {
	return 429
}

func (o *DeleteRegistryEntitiesTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /container-security/entities/registries/v1][%d] deleteRegistryEntitiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRegistryEntitiesTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /container-security/entities/registries/v1][%d] deleteRegistryEntitiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRegistryEntitiesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteRegistryEntitiesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
