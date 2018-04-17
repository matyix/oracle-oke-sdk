// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClustersapiGetRegionRequest The response to a get available node shapes.
// swagger:model clustersapiGetRegionRequest
type ClustersapiGetRegionRequest struct {

	// Credential set ID.
	CloudAuthID string `json:"cloudAuthId,omitempty"`

	// Wercker organization ID.
	OwnerID string `json:"ownerId,omitempty"`
}

// Validate validates this clustersapi get region request
func (m *ClustersapiGetRegionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ClustersapiGetRegionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustersapiGetRegionRequest) UnmarshalBinary(b []byte) error {
	var res ClustersapiGetRegionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
