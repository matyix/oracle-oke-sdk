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

	"github.com/matyix/oracle-oke-client/models"
)

// NewGetNodeImagesParams creates a new GetNodeImagesParams object
// with the default values initialized.
func NewGetNodeImagesParams() *GetNodeImagesParams {
	var ()
	return &GetNodeImagesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNodeImagesParamsWithTimeout creates a new GetNodeImagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNodeImagesParamsWithTimeout(timeout time.Duration) *GetNodeImagesParams {
	var ()
	return &GetNodeImagesParams{

		timeout: timeout,
	}
}

// NewGetNodeImagesParamsWithContext creates a new GetNodeImagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNodeImagesParamsWithContext(ctx context.Context) *GetNodeImagesParams {
	var ()
	return &GetNodeImagesParams{

		Context: ctx,
	}
}

// NewGetNodeImagesParamsWithHTTPClient creates a new GetNodeImagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNodeImagesParamsWithHTTPClient(client *http.Client) *GetNodeImagesParams {
	var ()
	return &GetNodeImagesParams{
		HTTPClient: client,
	}
}

/*GetNodeImagesParams contains all the parameters to send to the API endpoint
for the get node images operation typically these are written to a http.Request
*/
type GetNodeImagesParams struct {

	/*Body
	  Body parameters.

	*/
	Body *models.ClustersapiGetNodeImagesRequest
	/*OwnerID
	  Wercker organization ID.

	Example: 59bbf8affb1beb01009cae7a

	*/
	OwnerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get node images params
func (o *GetNodeImagesParams) WithTimeout(timeout time.Duration) *GetNodeImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get node images params
func (o *GetNodeImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get node images params
func (o *GetNodeImagesParams) WithContext(ctx context.Context) *GetNodeImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get node images params
func (o *GetNodeImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get node images params
func (o *GetNodeImagesParams) WithHTTPClient(client *http.Client) *GetNodeImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get node images params
func (o *GetNodeImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get node images params
func (o *GetNodeImagesParams) WithBody(body *models.ClustersapiGetNodeImagesRequest) *GetNodeImagesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get node images params
func (o *GetNodeImagesParams) SetBody(body *models.ClustersapiGetNodeImagesRequest) {
	o.Body = body
}

// WithOwnerID adds the ownerID to the get node images params
func (o *GetNodeImagesParams) WithOwnerID(ownerID string) *GetNodeImagesParams {
	o.SetOwnerID(ownerID)
	return o
}

// SetOwnerID adds the ownerId to the get node images params
func (o *GetNodeImagesParams) SetOwnerID(ownerID string) {
	o.OwnerID = ownerID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNodeImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
