// Code generated by go-swagger; DO NOT EDIT.

package manage_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListClustersForOwnerParams creates a new ListClustersForOwnerParams object
// with the default values initialized.
func NewListClustersForOwnerParams() *ListClustersForOwnerParams {
	var ()
	return &ListClustersForOwnerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListClustersForOwnerParamsWithTimeout creates a new ListClustersForOwnerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListClustersForOwnerParamsWithTimeout(timeout time.Duration) *ListClustersForOwnerParams {
	var ()
	return &ListClustersForOwnerParams{

		timeout: timeout,
	}
}

// NewListClustersForOwnerParamsWithContext creates a new ListClustersForOwnerParams object
// with the default values initialized, and the ability to set a context for a request
func NewListClustersForOwnerParamsWithContext(ctx context.Context) *ListClustersForOwnerParams {
	var ()
	return &ListClustersForOwnerParams{

		Context: ctx,
	}
}

// NewListClustersForOwnerParamsWithHTTPClient creates a new ListClustersForOwnerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListClustersForOwnerParamsWithHTTPClient(client *http.Client) *ListClustersForOwnerParams {
	var ()
	return &ListClustersForOwnerParams{
		HTTPClient: client,
	}
}

/*ListClustersForOwnerParams contains all the parameters to send to the API endpoint
for the list clusters for owner operation typically these are written to a http.Request
*/
type ListClustersForOwnerParams struct {

	/*OwnerID
	  Wercker organization ID.

	Example: 59bbf8affb1beb01009cae7a

	*/
	OwnerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list clusters for owner params
func (o *ListClustersForOwnerParams) WithTimeout(timeout time.Duration) *ListClustersForOwnerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list clusters for owner params
func (o *ListClustersForOwnerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list clusters for owner params
func (o *ListClustersForOwnerParams) WithContext(ctx context.Context) *ListClustersForOwnerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list clusters for owner params
func (o *ListClustersForOwnerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list clusters for owner params
func (o *ListClustersForOwnerParams) WithHTTPClient(client *http.Client) *ListClustersForOwnerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list clusters for owner params
func (o *ListClustersForOwnerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwnerID adds the ownerID to the list clusters for owner params
func (o *ListClustersForOwnerParams) WithOwnerID(ownerID string) *ListClustersForOwnerParams {
	o.SetOwnerID(ownerID)
	return o
}

// SetOwnerID adds the ownerId to the list clusters for owner params
func (o *ListClustersForOwnerParams) SetOwnerID(ownerID string) {
	o.OwnerID = ownerID
}

// WriteToRequest writes these params to a swagger request
func (o *ListClustersForOwnerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ownerId
	if err := r.SetPathParam("ownerId", o.OwnerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}