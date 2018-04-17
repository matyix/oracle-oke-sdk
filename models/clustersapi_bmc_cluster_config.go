// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClustersapiBmcClusterConfig This object contains details required to define the cluster.
// swagger:model clustersapiBmcClusterConfig
type ClustersapiBmcClusterConfig struct {

	// The OCID of the compartment the cluster will reside in.
	Compartment string `json:"compartment,omitempty"`

	// pools
	Pools ClustersapiBmcClusterConfigPools `json:"pools"`

	// The region in which the cluster is being created.
	Region string `json:"region,omitempty"`

	// The base64-encoded SSH public key that matches the private key to connect to the cluster.
	SSHPublicKey string `json:"sshPublicKey,omitempty"`

	// The OCID of the VCN that the cluster will reside in.
	VcnID string `json:"vcnId,omitempty"`
}

// Validate validates this clustersapi bmc cluster config
func (m *ClustersapiBmcClusterConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ClustersapiBmcClusterConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustersapiBmcClusterConfig) UnmarshalBinary(b []byte) error {
	var res ClustersapiBmcClusterConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}