// Code generated by go-swagger; DO NOT EDIT.

package overwatch_dashboard

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

// AggregatesEventsCollectionsReader is a Reader for the AggregatesEventsCollections structure.
type AggregatesEventsCollectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AggregatesEventsCollectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAggregatesEventsCollectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAggregatesEventsCollectionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAggregatesEventsCollectionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /overwatch-dashboards/aggregates/events-collections/GET/v1] AggregatesEventsCollections", response, response.Code())
	}
}

// NewAggregatesEventsCollectionsOK creates a AggregatesEventsCollectionsOK with default headers values
func NewAggregatesEventsCollectionsOK() *AggregatesEventsCollectionsOK {
	return &AggregatesEventsCollectionsOK{}
}

/*
AggregatesEventsCollectionsOK describes a response with status code 200, with default header values.

OK
*/
type AggregatesEventsCollectionsOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaAggregatesResponse
}

// IsSuccess returns true when this aggregates events collections o k response has a 2xx status code
func (o *AggregatesEventsCollectionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aggregates events collections o k response has a 3xx status code
func (o *AggregatesEventsCollectionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregates events collections o k response has a 4xx status code
func (o *AggregatesEventsCollectionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregates events collections o k response has a 5xx status code
func (o *AggregatesEventsCollectionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregates events collections o k response a status code equal to that given
func (o *AggregatesEventsCollectionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the aggregates events collections o k response
func (o *AggregatesEventsCollectionsOK) Code() int {
	return 200
}

func (o *AggregatesEventsCollectionsOK) Error() string {
	return fmt.Sprintf("[POST /overwatch-dashboards/aggregates/events-collections/GET/v1][%d] aggregatesEventsCollectionsOK  %+v", 200, o.Payload)
}

func (o *AggregatesEventsCollectionsOK) String() string {
	return fmt.Sprintf("[POST /overwatch-dashboards/aggregates/events-collections/GET/v1][%d] aggregatesEventsCollectionsOK  %+v", 200, o.Payload)
}

func (o *AggregatesEventsCollectionsOK) GetPayload() *models.MsaAggregatesResponse {
	return o.Payload
}

func (o *AggregatesEventsCollectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregatesEventsCollectionsForbidden creates a AggregatesEventsCollectionsForbidden with default headers values
func NewAggregatesEventsCollectionsForbidden() *AggregatesEventsCollectionsForbidden {
	return &AggregatesEventsCollectionsForbidden{}
}

/*
AggregatesEventsCollectionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AggregatesEventsCollectionsForbidden struct {

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

// IsSuccess returns true when this aggregates events collections forbidden response has a 2xx status code
func (o *AggregatesEventsCollectionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregates events collections forbidden response has a 3xx status code
func (o *AggregatesEventsCollectionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregates events collections forbidden response has a 4xx status code
func (o *AggregatesEventsCollectionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregates events collections forbidden response has a 5xx status code
func (o *AggregatesEventsCollectionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregates events collections forbidden response a status code equal to that given
func (o *AggregatesEventsCollectionsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the aggregates events collections forbidden response
func (o *AggregatesEventsCollectionsForbidden) Code() int {
	return 403
}

func (o *AggregatesEventsCollectionsForbidden) Error() string {
	return fmt.Sprintf("[POST /overwatch-dashboards/aggregates/events-collections/GET/v1][%d] aggregatesEventsCollectionsForbidden  %+v", 403, o.Payload)
}

func (o *AggregatesEventsCollectionsForbidden) String() string {
	return fmt.Sprintf("[POST /overwatch-dashboards/aggregates/events-collections/GET/v1][%d] aggregatesEventsCollectionsForbidden  %+v", 403, o.Payload)
}

func (o *AggregatesEventsCollectionsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregatesEventsCollectionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregatesEventsCollectionsTooManyRequests creates a AggregatesEventsCollectionsTooManyRequests with default headers values
func NewAggregatesEventsCollectionsTooManyRequests() *AggregatesEventsCollectionsTooManyRequests {
	return &AggregatesEventsCollectionsTooManyRequests{}
}

/*
AggregatesEventsCollectionsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AggregatesEventsCollectionsTooManyRequests struct {

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

// IsSuccess returns true when this aggregates events collections too many requests response has a 2xx status code
func (o *AggregatesEventsCollectionsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregates events collections too many requests response has a 3xx status code
func (o *AggregatesEventsCollectionsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregates events collections too many requests response has a 4xx status code
func (o *AggregatesEventsCollectionsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregates events collections too many requests response has a 5xx status code
func (o *AggregatesEventsCollectionsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregates events collections too many requests response a status code equal to that given
func (o *AggregatesEventsCollectionsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the aggregates events collections too many requests response
func (o *AggregatesEventsCollectionsTooManyRequests) Code() int {
	return 429
}

func (o *AggregatesEventsCollectionsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /overwatch-dashboards/aggregates/events-collections/GET/v1][%d] aggregatesEventsCollectionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregatesEventsCollectionsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /overwatch-dashboards/aggregates/events-collections/GET/v1][%d] aggregatesEventsCollectionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregatesEventsCollectionsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregatesEventsCollectionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
