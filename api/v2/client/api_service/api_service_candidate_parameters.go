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

// NewAPIServiceCandidateParams creates a new APIServiceCandidateParams object
// with the default values initialized.
func NewAPIServiceCandidateParams() *APIServiceCandidateParams {
	var ()
	return &APIServiceCandidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIServiceCandidateParamsWithTimeout creates a new APIServiceCandidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIServiceCandidateParamsWithTimeout(timeout time.Duration) *APIServiceCandidateParams {
	var ()
	return &APIServiceCandidateParams{

		timeout: timeout,
	}
}

// NewAPIServiceCandidateParamsWithContext creates a new APIServiceCandidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIServiceCandidateParamsWithContext(ctx context.Context) *APIServiceCandidateParams {
	var ()
	return &APIServiceCandidateParams{

		Context: ctx,
	}
}

// NewAPIServiceCandidateParamsWithHTTPClient creates a new APIServiceCandidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIServiceCandidateParamsWithHTTPClient(client *http.Client) *APIServiceCandidateParams {
	var ()
	return &APIServiceCandidateParams{
		HTTPClient: client,
	}
}

/*APIServiceCandidateParams contains all the parameters to send to the API endpoint
for the Api service candidate operation typically these are written to a http.Request
*/
type APIServiceCandidateParams struct {

	/*Height*/
	Height *string
	/*PublicKey*/
	PublicKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Api service candidate params
func (o *APIServiceCandidateParams) WithTimeout(timeout time.Duration) *APIServiceCandidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Api service candidate params
func (o *APIServiceCandidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Api service candidate params
func (o *APIServiceCandidateParams) WithContext(ctx context.Context) *APIServiceCandidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Api service candidate params
func (o *APIServiceCandidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Api service candidate params
func (o *APIServiceCandidateParams) WithHTTPClient(client *http.Client) *APIServiceCandidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Api service candidate params
func (o *APIServiceCandidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHeight adds the height to the Api service candidate params
func (o *APIServiceCandidateParams) WithHeight(height *string) *APIServiceCandidateParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the Api service candidate params
func (o *APIServiceCandidateParams) SetHeight(height *string) {
	o.Height = height
}

// WithPublicKey adds the publicKey to the Api service candidate params
func (o *APIServiceCandidateParams) WithPublicKey(publicKey string) *APIServiceCandidateParams {
	o.SetPublicKey(publicKey)
	return o
}

// SetPublicKey adds the publicKey to the Api service candidate params
func (o *APIServiceCandidateParams) SetPublicKey(publicKey string) {
	o.PublicKey = publicKey
}

// WriteToRequest writes these params to a swagger request
func (o *APIServiceCandidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param public_key
	if err := r.SetPathParam("public_key", o.PublicKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
