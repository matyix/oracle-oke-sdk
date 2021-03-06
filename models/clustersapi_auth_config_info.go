// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClustersapiAuthConfigInfo The credential set details.
// swagger:model clustersapiAuthConfigInfo
type ClustersapiAuthConfigInfo struct {

	// The type of cloud infrastructure.
	CloudType string `json:"cloudType,omitempty"`

	// The ID of the credential set.
	ID string `json:"id,omitempty"`

	// Credential set name.
	Name string `json:"name,omitempty"`

	// Wercker organization ID.
	ResourceOwnerID string `json:"resourceOwnerId,omitempty"`

	// The OCID of the tenancy.
	TenantID string `json:"tenantId,omitempty"`
}

// Validate validates this clustersapi auth config info
func (m *ClustersapiAuthConfigInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ClustersapiAuthConfigInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustersapiAuthConfigInfo) UnmarshalBinary(b []byte) error {
	var res ClustersapiAuthConfigInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
