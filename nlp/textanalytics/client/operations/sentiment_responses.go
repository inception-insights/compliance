package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/docker/compliance/nlp/textanalytics/models"
)

// SentimentReader is a Reader for the Sentiment structure.
type SentimentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SentimentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSentimentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSentimentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSentimentOK creates a SentimentOK with default headers values
func NewSentimentOK() *SentimentOK {
	return &SentimentOK{}
}

/*SentimentOK handles this case with default header values.

OK
*/
type SentimentOK struct {
	Payload *models.SentimentBatchResultV2
}

func (o *SentimentOK) Error() string {
	return fmt.Sprintf("[POST /sentiment][%d] sentimentOK  %+v", 200, o.Payload)
}

func (o *SentimentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SentimentBatchResultV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSentimentBadRequest creates a SentimentBadRequest with default headers values
func NewSentimentBadRequest() *SentimentBadRequest {
	return &SentimentBadRequest{}
}

/*SentimentBadRequest handles this case with default header values.

BadRequest
*/
type SentimentBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *SentimentBadRequest) Error() string {
	return fmt.Sprintf("[POST /sentiment][%d] sentimentBadRequest  %+v", 400, o.Payload)
}

func (o *SentimentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
