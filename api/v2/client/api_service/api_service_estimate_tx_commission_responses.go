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

// APIServiceEstimateTxCommissionReader is a Reader for the APIServiceEstimateTxCommission structure.
type APIServiceEstimateTxCommissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *APIServiceEstimateTxCommissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAPIServiceEstimateTxCommissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAPIServiceEstimateTxCommissionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAPIServiceEstimateTxCommissionOK creates a APIServiceEstimateTxCommissionOK with default headers values
func NewAPIServiceEstimateTxCommissionOK() *APIServiceEstimateTxCommissionOK {
	return &APIServiceEstimateTxCommissionOK{}
}

/*APIServiceEstimateTxCommissionOK handles this case with default header values.

A successful response.
*/
type APIServiceEstimateTxCommissionOK struct {
	Payload *models.APIPbEstimateTxCommissionResponse
}

func (o *APIServiceEstimateTxCommissionOK) Error() string {
	return fmt.Sprintf("[GET /estimate_tx_commission/{tx}][%d] apiServiceEstimateTxCommissionOK  %+v", 200, o.Payload)
}

func (o *APIServiceEstimateTxCommissionOK) GetPayload() *models.APIPbEstimateTxCommissionResponse {
	return o.Payload
}

func (o *APIServiceEstimateTxCommissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIPbEstimateTxCommissionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAPIServiceEstimateTxCommissionDefault creates a APIServiceEstimateTxCommissionDefault with default headers values
func NewAPIServiceEstimateTxCommissionDefault(code int) *APIServiceEstimateTxCommissionDefault {
	return &APIServiceEstimateTxCommissionDefault{
		_statusCode: code,
	}
}

/*APIServiceEstimateTxCommissionDefault handles this case with default header values.

An unexpected error response
*/
type APIServiceEstimateTxCommissionDefault struct {
	_statusCode int

	Payload *models.RuntimeError
}

// Code gets the status code for the Api service estimate tx commission default response
func (o *APIServiceEstimateTxCommissionDefault) Code() int {
	return o._statusCode
}

func (o *APIServiceEstimateTxCommissionDefault) Error() string {
	return fmt.Sprintf("[GET /estimate_tx_commission/{tx}][%d] ApiService_EstimateTxCommission default  %+v", o._statusCode, o.Payload)
}

func (o *APIServiceEstimateTxCommissionDefault) GetPayload() *models.RuntimeError {
	return o.Payload
}

func (o *APIServiceEstimateTxCommissionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
