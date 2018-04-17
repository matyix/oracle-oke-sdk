// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClustersapiListAuthConfigsResponse The list of credential sets.
// swagger:model clustersapiListAuthConfigsResponse
type ClustersapiListAuthConfigsResponse struct {

	// auth configs
	AuthConfigs ClustersapiListAuthConfigsResponseAuthConfigs `json:"authConfigs"`
}

// Validate validates this clustersapi list auth configs response
func (m *ClustersapiListAuthConfigsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ClustersapiListAuthConfigsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustersapiListAuthConfigsResponse) UnmarshalBinary(b []byte) error {
	var res ClustersapiListAuthConfigsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}