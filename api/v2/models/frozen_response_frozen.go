// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FrozenResponseFrozen frozen response frozen
//
// swagger:model FrozenResponseFrozen
type FrozenResponseFrozen struct {

	// address
	Address string `json:"address,omitempty"`

	// candidate key
	CandidateKey string `json:"candidate_key,omitempty"`

	// coin
	Coin string `json:"coin,omitempty"`

	// height
	Height string `json:"height,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this frozen response frozen
func (m *FrozenResponseFrozen) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FrozenResponseFrozen) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FrozenResponseFrozen) UnmarshalBinary(b []byte) error {
	var res FrozenResponseFrozen
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
