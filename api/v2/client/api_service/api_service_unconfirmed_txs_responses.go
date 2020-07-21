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

// APIServiceUnconfirmedTxsReader is a Reader for the APIServiceUnconfirmedTxs structure.
type APIServiceUnconfirmedTxsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIServiceUnconfirmedTxsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIServiceUnconfirmedTxsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAPIServiceUnconfirmedTxsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAPIServiceUnconfirmedTxsOK creates a APIServiceUnconfirmedTxsOK with default headers values
func NewAPIServiceUnconfirmedTxsOK() *APIServiceUnconfirmedTxsOK {
	return &APIServiceUnconfirmedTxsOK{}
}

/*APIServiceUnconfirmedTxsOK handles this case with default header values.

A successful response.
*/
type APIServiceUnconfirmedTxsOK struct {
	Payload *models.APIPbUnconfirmedTxsResponse
}

func (o *APIServiceUnconfirmedTxsOK) Error() string {
	return fmt.Sprintf("[GET /unconfirmed_txs][%d] apiServiceUnconfirmedTxsOK  %+v", 200, o.Payload)
}

func (o *APIServiceUnconfirmedTxsOK) GetPayload() *models.APIPbUnconfirmedTxsResponse {
	return o.Payload
}

func (o *APIServiceUnconfirmedTxsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPbUnconfirmedTxsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIServiceUnconfirmedTxsDefault creates a APIServiceUnconfirmedTxsDefault with default headers values
func NewAPIServiceUnconfirmedTxsDefault(code int) *APIServiceUnconfirmedTxsDefault {
	return &APIServiceUnconfirmedTxsDefault{
		_statusCode: code,
	}
}

/*APIServiceUnconfirmedTxsDefault handles this case with default header values.

An unexpected error response
*/
type APIServiceUnconfirmedTxsDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the Api service unconfirmed txs default response
func (o *APIServiceUnconfirmedTxsDefault) Code() int {
	return o._statusCode
}

func (o *APIServiceUnconfirmedTxsDefault) Error() string {
	return fmt.Sprintf("[GET /unconfirmed_txs][%d] ApiService_UnconfirmedTxs default  %+v", o._statusCode, o.Payload)
}

func (o *APIServiceUnconfirmedTxsDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *APIServiceUnconfirmedTxsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
