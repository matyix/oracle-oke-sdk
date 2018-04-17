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

// NewGetAuthConfigParams creates a new GetAuthConfigParams object
// with the default values initialized.
func NewGetAuthConfigParams() *GetAuthConfigParams {
	var ()
	return &GetAuthConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAuthConfigParamsWithTimeout creates a new GetAuthConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAuthConfigParamsWithTimeout(timeout time.Duration) *GetAuthConfigParams {
	var ()
	return &GetAuthConfigParams{

		timeout: timeout,
	}
}

// NewGetAuthConfigParamsWithContext creates a new GetAuthConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAuthConfigParamsWithContext(ctx context.Context) *GetAuthConfigParams {
	var ()
	return &GetAuthConfigParams{

		Context: ctx,
	}
}

// NewGetAuthConfigParamsWithHTTPClient creates a new GetAuthConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAuthConfigParamsWithHTTPClient(client *http.Client) *GetAuthConfigParams {
	var ()
	return &GetAuthConfigParams{
		HTTPClient: client,
	}
}

/*GetAuthConfigParams contains all the parameters to send to the API endpoint
for the get auth config operation typically these are written to a http.Request
*/
type GetAuthConfigParams struct {

	/*ID
	  Cluster ID.

	Example: c56ead49fb8

	*/
	ID string
	/*OwnerID
	  Wercker organization ID.

	Example: 59bbf8affb1beb01009cae7a

	*/
	OwnerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get auth config params
func (o *GetAuthConfigParams) WithTimeout(timeout time.Duration) *GetAuthConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get auth config params
func (o *GetAuthConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get auth config params
func (o *GetAuthConfigParams) WithContext(ctx context.Context) *GetAuthConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get auth config params
func (o *GetAuthConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get auth config params
func (o *GetAuthConfigParams) WithHTTPClient(client *http.Client) *GetAuthConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get auth config params
func (o *GetAuthConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get auth config params
func (o *GetAuthConfigParams) WithID(id string) *GetAuthConfigParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get auth config params
func (o *GetAuthConfigParams) SetID(id string) {
	o.ID = id
}

// WithOwnerID adds the ownerID to the get auth config params
func (o *GetAuthConfigParams) WithOwnerID(ownerID string) *GetAuthConfigParams {
	o.SetOwnerID(ownerID)
	return o
}

// SetOwnerID adds the ownerId to the get auth config params
func (o *GetAuthConfigParams) SetOwnerID(ownerID string) {
	o.OwnerID = ownerID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAuthConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param ownerId
	if err := r.SetPathParam("ownerId", o.OwnerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
