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
	"github.com/go-openapi/swag"
)

// NewAPIServiceCoinIDParams creates a new APIServiceCoinIDParams object
// with the default values initialized.
func NewAPIServiceCoinIDParams() *APIServiceCoinIDParams {
	var ()
	return &APIServiceCoinIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIServiceCoinIDParamsWithTimeout creates a new APIServiceCoinIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIServiceCoinIDParamsWithTimeout(timeout time.Duration) *APIServiceCoinIDParams {
	var ()
	return &APIServiceCoinIDParams{

		timeout: timeout,
	}
}

// NewAPIServiceCoinIDParamsWithContext creates a new APIServiceCoinIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIServiceCoinIDParamsWithContext(ctx context.Context) *APIServiceCoinIDParams {
	var ()
	return &APIServiceCoinIDParams{

		Context: ctx,
	}
}

// NewAPIServiceCoinIDParamsWithHTTPClient creates a new APIServiceCoinIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIServiceCoinIDParamsWithHTTPClient(client *http.Client) *APIServiceCoinIDParams {
	var ()
	return &APIServiceCoinIDParams{
		HTTPClient: client,
	}
}

/*APIServiceCoinIDParams contains all the parameters to send to the API endpoint
for the Api service coin Id operation typically these are written to a http.Request
*/
type APIServiceCoinIDParams struct {

	/*Height*/
	Height *string
	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Api service coin Id params
func (o *APIServiceCoinIDParams) WithTimeout(timeout time.Duration) *APIServiceCoinIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Api service coin Id params
func (o *APIServiceCoinIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Api service coin Id params
func (o *APIServiceCoinIDParams) WithContext(ctx context.Context) *APIServiceCoinIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Api service coin Id params
func (o *APIServiceCoinIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Api service coin Id params
func (o *APIServiceCoinIDParams) WithHTTPClient(client *http.Client) *APIServiceCoinIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Api service coin Id params
func (o *APIServiceCoinIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHeight adds the height to the Api service coin Id params
func (o *APIServiceCoinIDParams) WithHeight(height *string) *APIServiceCoinIDParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the Api service coin Id params
func (o *APIServiceCoinIDParams) SetHeight(height *string) {
	o.Height = height
}

// WithID adds the id to the Api service coin Id params
func (o *APIServiceCoinIDParams) WithID(id int64) *APIServiceCoinIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the Api service coin Id params
func (o *APIServiceCoinIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *APIServiceCoinIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
