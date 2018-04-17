// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClustersapiUpgradeK8SMasterRequest Upgrade Kubernetes master request.
// swagger:model clustersapiUpgradeK8SMasterRequest
type ClustersapiUpgradeK8SMasterRequest struct {

	// Cluster ID.
	ClusterID string `json:"clusterId,omitempty"`

	// Kubernetes version to upgrade to.
	K8sVersion string `json:"k8sVersion,omitempty"`

	// Wercker organization ID.
	OwnerID string `json:"ownerId,omitempty"`
}

// Validate validates this clustersapi upgrade k8 s master request
func (m *ClustersapiUpgradeK8SMasterRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ClustersapiUpgradeK8SMasterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustersapiUpgradeK8SMasterRequest) UnmarshalBinary(b []byte) error {
	var res ClustersapiUpgradeK8SMasterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}