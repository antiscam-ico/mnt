// Code generated by go-swagger; DO NOT EDIT.

package api_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/MinterTeam/minter-go-sdk/v2/api/v2/models"
)

// APIServiceValidatorsReader is a Reader for the APIServiceValidators structure.
type APIServiceValidatorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIServiceValidatorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIServiceValidatorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAPIServiceValidatorsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAPIServiceValidatorsOK creates a APIServiceValidatorsOK with default headers values
func NewAPIServiceValidatorsOK() *APIServiceValidatorsOK {
	return &APIServiceValidatorsOK{}
}

/*APIServiceValidatorsOK handles this case with default header values.

A successful response.
*/
type APIServiceValidatorsOK struct {
	Payload *models.APIPbValidatorsResponse
}

func (o *APIServiceValidatorsOK) Error() string {
	return fmt.Sprintf("[GET /validators][%d] apiServiceValidatorsOK  %+v", 200, o.Payload)
}

func (o *APIServiceValidatorsOK) GetPayload() *models.APIPbValidatorsResponse {
	return o.Payload
}

func (o *APIServiceValidatorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPbValidatorsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIServiceValidatorsDefault creates a APIServiceValidatorsDefault with default headers values
func NewAPIServiceValidatorsDefault(code int) *APIServiceValidatorsDefault {
	return &APIServiceValidatorsDefault{
		_statusCode: code,
	}
}

/*APIServiceValidatorsDefault handles this case with default header values.

An unexpected error response
*/
type APIServiceValidatorsDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the Api service validators default response
func (o *APIServiceValidatorsDefault) Code() int {
	return o._statusCode
}

func (o *APIServiceValidatorsDefault) Error() string {
	return fmt.Sprintf("[GET /validators][%d] ApiService_Validators default  %+v", o._statusCode, o.Payload)
}

func (o *APIServiceValidatorsDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *APIServiceValidatorsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
