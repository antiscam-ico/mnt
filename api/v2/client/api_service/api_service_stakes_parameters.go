// Code generated by go-swagger; DO NOT EDIT.

package api_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewAPIServiceStakesParams creates a new APIServiceStakesParams object
// with the default values initialized.
func NewAPIServiceStakesParams() *APIServiceStakesParams {
	var ()
	return &APIServiceStakesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIServiceStakesParamsWithTimeout creates a new APIServiceStakesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIServiceStakesParamsWithTimeout(timeout time.Duration) *APIServiceStakesParams {
	var ()
	return &APIServiceStakesParams{

		timeout: timeout,
	}
}

// NewAPIServiceStakesParamsWithContext creates a new APIServiceStakesParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIServiceStakesParamsWithContext(ctx context.Context) *APIServiceStakesParams {
	var ()
	return &APIServiceStakesParams{

		Context: ctx,
	}
}

// NewAPIServiceStakesParamsWithHTTPClient creates a new APIServiceStakesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIServiceStakesParamsWithHTTPClient(client *http.Client) *APIServiceStakesParams {
	var ()
	return &APIServiceStakesParams{
		HTTPClient: client,
	}
}

/*APIServiceStakesParams contains all the parameters to send to the API endpoint
for the Api service stakes operation typically these are written to a http.Request
*/
type APIServiceStakesParams struct {

	/*Address*/
	Address *string
	/*Coin*/
	Coin *string
	/*Height*/
	Height *string
	/*PublicKey*/
	PublicKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Api service stakes params
func (o *APIServiceStakesParams) WithTimeout(timeout time.Duration) *APIServiceStakesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Api service stakes params
func (o *APIServiceStakesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Api service stakes params
func (o *APIServiceStakesParams) WithContext(ctx context.Context) *APIServiceStakesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Api service stakes params
func (o *APIServiceStakesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Api service stakes params
func (o *APIServiceStakesParams) WithHTTPClient(client *http.Client) *APIServiceStakesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Api service stakes params
func (o *APIServiceStakesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the Api service stakes params
func (o *APIServiceStakesParams) WithAddress(address *string) *APIServiceStakesParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the Api service stakes params
func (o *APIServiceStakesParams) SetAddress(address *string) {
	o.Address = address
}

// WithCoin adds the coin to the Api service stakes params
func (o *APIServiceStakesParams) WithCoin(coin *string) *APIServiceStakesParams {
	o.SetCoin(coin)
	return o
}

// SetCoin adds the coin to the Api service stakes params
func (o *APIServiceStakesParams) SetCoin(coin *string) {
	o.Coin = coin
}

// WithHeight adds the height to the Api service stakes params
func (o *APIServiceStakesParams) WithHeight(height *string) *APIServiceStakesParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the Api service stakes params
func (o *APIServiceStakesParams) SetHeight(height *string) {
	o.Height = height
}

// WithPublicKey adds the publicKey to the Api service stakes params
func (o *APIServiceStakesParams) WithPublicKey(publicKey string) *APIServiceStakesParams {
	o.SetPublicKey(publicKey)
	return o
}

// SetPublicKey adds the publicKey to the Api service stakes params
func (o *APIServiceStakesParams) SetPublicKey(publicKey string) {
	o.PublicKey = publicKey
}

// WriteToRequest writes these params to a swagger request
func (o *APIServiceStakesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Address != nil {

		// query param address
		var qrAddress string
		if o.Address != nil {
			qrAddress = *o.Address
		}
		qAddress := qrAddress
		if qAddress != "" {
			if err := r.SetQueryParam("address", qAddress); err != nil {
				return err
			}
		}

	}

	if o.Coin != nil {

		// query param coin
		var qrCoin string
		if o.Coin != nil {
			qrCoin = *o.Coin
		}
		qCoin := qrCoin
		if qCoin != "" {
			if err := r.SetQueryParam("coin", qCoin); err != nil {
				return err
			}
		}

	}

	if o.Height != nil {

		// query param height
		var qrHeight string
		if o.Height != nil {
			qrHeight = *o.Height
		}
		qHeight := qrHeight
		if qHeight != "" {
			if err := r.SetQueryParam("height", qHeight); err != nil {
				return err
			}
		}

	}

	// path param public_key
	if err := r.SetPathParam("public_key", o.PublicKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
