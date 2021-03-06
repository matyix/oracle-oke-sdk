// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClustersapiCloudDetails Schema description.
// swagger:model clustersapiCloudDetails
type ClustersapiCloudDetails struct {

	// bmc
	Bmc *ClustersapiBmcCloudDetails `json:"bmc,omitempty"`
}

// Validate validates this clustersapi cloud details
func (m *ClustersapiCloudDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBmc(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClustersapiCloudDetails) validateBmc(formats strfmt.Registry) error {

	if swag.IsZero(m.Bmc) { // not required
		return nil
	}

	if m.Bmc != nil {

		if err := m.Bmc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bmc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClustersapiCloudDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustersapiCloudDetails) UnmarshalBinary(b []byte) error {
	var res ClustersapiCloudDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
