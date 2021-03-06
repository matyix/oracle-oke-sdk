// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClustersapiPool clustersapi pool
// swagger:model clustersapiPool
type ClustersapiPool struct {

	// ads
	Ads ClustersapiPoolAds `json:"ads"`

	// The OCID of the image to run on this node.
	ImageID string `json:"imageId,omitempty"`

	// The name of the image to run on this node.
	ImageName string `json:"imageName,omitempty"`

	// Kubernetes version.
	K8Version string `json:"k8Version,omitempty"`

	// The name of the pool.
	Name string `json:"name,omitempty"`

	// The shape, or resource profile, that determines the number of CPUs and amount of memory assigned to each worker node.
	Shape string `json:"shape,omitempty"`

	// The number of worker nodes to run per Availability Domain.
	WorkersPerAD int32 `json:"workersPerAD,omitempty"`
}

// Validate validates this clustersapi pool
func (m *ClustersapiPool) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ClustersapiPool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustersapiPool) UnmarshalBinary(b []byte) error {
	var res ClustersapiPool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
