// Code generated by go-swagger; DO NOT EDIT.

package api_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/MinterTeam/minter-go-sdk/api/v2/models"
)

// APIServiceFrozenReader is a Reader for the APIServiceFrozen structure.
type APIServiceFrozenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIServiceFrozenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIServiceFrozenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAPIServiceFrozenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAPIServiceFrozenOK creates a APIServiceFrozenOK with default headers values
func NewAPIServiceFrozenOK() *APIServiceFrozenOK {
	return &APIServiceFrozenOK{}
}

/*APIServiceFrozenOK handles this case with default header values.

A successful response.
*/
type APIServiceFrozenOK struct {
	Payload *models.APIPbFrozenResponse
}

func (o *APIServiceFrozenOK) Error() string {
	return fmt.Sprintf("[GET /frozen][%d] apiServiceFrozenOK  %+v", 200, o.Payload)
}

func (o *APIServiceFrozenOK) GetPayload() *models.APIPbFrozenResponse {
	return o.Payload
}

func (o *APIServiceFrozenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPbFrozenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIServiceFrozenDefault creates a APIServiceFrozenDefault with default headers values
func NewAPIServiceFrozenDefault(code int) *APIServiceFrozenDefault {
	return &APIServiceFrozenDefault{
		_statusCode: code,
	}
}

/*APIServiceFrozenDefault handles this case with default header values.

An unexpected error response
*/
type APIServiceFrozenDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the Api service frozen default response
func (o *APIServiceFrozenDefault) Code() int {
	return o._statusCode
}

func (o *APIServiceFrozenDefault) Error() string {
	return fmt.Sprintf("[GET /frozen][%d] ApiService_Frozen default  %+v", o._statusCode, o.Payload)
}

func (o *APIServiceFrozenDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *APIServiceFrozenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
