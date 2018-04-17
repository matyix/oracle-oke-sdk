// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClustersapiWorkerErrors Schema description.
// swagger:model clustersapiWorkerErrors
type ClustersapiWorkerErrors struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this clustersapi worker errors
func (m *ClustersapiWorkerErrors) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ClustersapiWorkerErrors) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClustersapiWorkerErrors) UnmarshalBinary(b []byte) error {
	var res ClustersapiWorkerErrors
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
