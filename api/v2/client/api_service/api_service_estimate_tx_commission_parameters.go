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

// NewAPIServiceEstimateTxCommissionParams creates a new APIServiceEstimateTxCommissionParams object
// with the default values initialized.
func NewAPIServiceEstimateTxCommissionParams() *APIServiceEstimateTxCommissionParams {
	var (
		dataTypeDefault = string("unknown")
	)
	return &APIServiceEstimateTxCommissionParams{
		DataType: &dataTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIServiceEstimateTxCommissionParamsWithTimeout creates a new APIServiceEstimateTxCommissionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIServiceEstimateTxCommissionParamsWithTimeout(timeout time.Duration) *APIServiceEstimateTxCommissionParams {
	var (
		dataTypeDefault = string("unknown")
	)
	return &APIServiceEstimateTxCommissionParams{
		DataType: &dataTypeDefault,

		timeout: timeout,
	}
}

// NewAPIServiceEstimateTxCommissionParamsWithContext creates a new APIServiceEstimateTxCommissionParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIServiceEstimateTxCommissionParamsWithContext(ctx context.Context) *APIServiceEstimateTxCommissionParams {
	var (
		dataTypeDefault = string("unknown")
	)
	return &APIServiceEstimateTxCommissionParams{
		DataType: &dataTypeDefault,

		Context: ctx,
	}
}

// NewAPIServiceEstimateTxCommissionParamsWithHTTPClient creates a new APIServiceEstimateTxCommissionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIServiceEstimateTxCommissionParamsWithHTTPClient(client *http.Client) *APIServiceEstimateTxCommissionParams {
	var (
		dataTypeDefault = string("unknown")
	)
	return &APIServiceEstimateTxCommissionParams{
		DataType:   &dataTypeDefault,
		HTTPClient: client,
	}
}

/*APIServiceEstimateTxCommissionParams contains all the parameters to send to the API endpoint
for the Api service estimate tx commission operation typically these are written to a http.Request
*/
type APIServiceEstimateTxCommissionParams struct {

	/*DataGasCoinID*/
	DataGasCoinID *int64
	/*DataMtxs*/
	DataMtxs *string
	/*DataPayload*/
	DataPayload *strfmt.Base64
	/*DataType*/
	DataType *string
	/*Height*/
	Height *string
	/*Tx*/
	Tx string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Api service estimate tx commission params
func (o *APIServiceEstimateTxCommissionParams) WithTimeout(timeout time.Duration) *APIServiceEstimateTxCommissionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Api service estimate tx commission params
func (o *APIServiceEstimateTxCommissionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Api service estimate tx commission params
func (o *APIServiceEstimateTxCommissionParams) WithContext(ctx context.Context) *APIServiceEstimateTxCommissionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Api service estimate tx commission params
func (o *APIServiceEstimateTxCommissionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Api service estimate tx commission params
func (o *APIServiceEstimateTxCommissionParams) WithHTTPClient(client *http.Client) *APIServiceEstimateTxCommissionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Api service estimate tx commission params
func (o *APIServiceEstimateTxCommissionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDataGasCoinID adds the dataGasCoinID to the Api service estimate tx commission params
func (o *APIServiceEstimateTxCommissionParams) WithDataGasCoinID(dataGasCoinID *int64) *APIServiceEstimateTxCommissionParams {
	o.SetDataGasCoinID(dataGasCoinID)
	return o
}

// SetDataGasCoinID adds the dataGasCoinId to the Api service estimate tx commission params
func (o *APIServiceEstimateTxCommissionParams) SetDataGasCoinID(dataGasCoinID *int64) {
	o.DataGasCoinID = dataGasCoinID
}

// WithDataMtxs adds the dataMtxs to the Api service estimate tx commission params
func (o *APIServiceEstimateTxCommissionParams) WithDataMtxs(dataMtxs *string) *APIServiceEstimateTxCommissionParams {
	o.SetDataMtxs(dataMtxs)
	return o
}

// SetDataMtxs adds the dataMtxs to the Api service estimate tx commission params
func (o *APIServiceEstimateTxCommissionParams) SetDataMtxs(dataMtxs *string) {
	o.DataMtxs = dataMtxs
}

// WithDataPayload adds the dataPayload to the Api service estimate tx commission params
func (o *APIServiceEstimateTxCommissionParams) WithDataPayload(dataPayload *strfmt.Base64) *APIServiceEstimateTxCommissionParams {
	o.SetDataPayload(dataPayload)
	return o
}

// SetDataPayload adds the dataPayload to the Api service estimate tx commission params
func (o *APIServiceEstimateTxCommissionParams) SetDataPayload(dataPayload *strfmt.Base64) {
	o.DataPayload = dataPayload
}

// WithDataType adds the dataType to the Api service estimate tx commission params
func (o *APIServiceEstimateTxCommissionParams) WithDataType(dataType *string) *APIServiceEstimateTxCommissionParams {
	o.SetDataType(dataType)
	return o
}

// SetDataType adds the dataType to the Api service estimate tx commission params
func (o *APIServiceEstimateTxCommissionParams) SetDataType(dataType *string) {
	o.DataType = dataType
}

// WithHeight adds the height to the Api service estimate tx commission params
func (o *APIServiceEstimateTxCommissionParams) WithHeight(height *string) *APIServiceEstimateTxCommissionParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the Api service estimate tx commission params
func (o *APIServiceEstimateTxCommissionParams) SetHeight(height *string) {
	o.Height = height
}

// WithTx adds the tx to the Api service estimate tx commission params
func (o *APIServiceEstimateTxCommissionParams) WithTx(tx string) *APIServiceEstimateTxCommissionParams {
	o.SetTx(tx)
	return o
}

// SetTx adds the tx to the Api service estimate tx commission params
func (o *APIServiceEstimateTxCommissionParams) SetTx(tx string) {
	o.Tx = tx
}

// WriteToRequest writes these params to a swagger request
func (o *APIServiceEstimateTxCommissionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DataGasCoinID != nil {

		// query param data.gas_coin_id
		var qrDataGasCoinID int64
		if o.DataGasCoinID != nil {
			qrDataGasCoinID = *o.DataGasCoinID
		}
		qDataGasCoinID := swag.FormatInt64(qrDataGasCoinID)
		if qDataGasCoinID != "" {
			if err := r.SetQueryParam("data.gas_coin_id", qDataGasCoinID); err != nil {
				return err
			}
		}

	}

	if o.DataMtxs != nil {

		// query param data.mtxs
		var qrDataMtxs string
		if o.DataMtxs != nil {
			qrDataMtxs = *o.DataMtxs
		}
		qDataMtxs := qrDataMtxs
		if qDataMtxs != "" {
			if err := r.SetQueryParam("data.mtxs", qDataMtxs); err != nil {
				return err
			}
		}

	}

	if o.DataPayload != nil {

		// query param data.payload
		var qrDataPayload strfmt.Base64
		if o.DataPayload != nil {
			qrDataPayload = *o.DataPayload
		}
		qDataPayload := qrDataPayload.String()
		if qDataPayload != "" {
			if err := r.SetQueryParam("data.payload", qDataPayload); err != nil {
				return err
			}
		}

	}

	if o.DataType != nil {

		// query param data.type
		var qrDataType string
		if o.DataType != nil {
			qrDataType = *o.DataType
		}
		qDataType := qrDataType
		if qDataType != "" {
			if err := r.SetQueryParam("data.type", qDataType); err != nil {
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

	// path param tx
	if err := r.SetPathParam("tx", o.Tx); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
