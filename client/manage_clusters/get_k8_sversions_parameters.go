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

// NewGetK8SVersionsParams creates a new GetK8SVersionsParams object
// with the default values initialized.
func NewGetK8SVersionsParams() *GetK8SVersionsParams {
	var ()
	return &GetK8SVersionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetK8SVersionsParamsWithTimeout creates a new GetK8SVersionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetK8SVersionsParamsWithTimeout(timeout time.Duration) *GetK8SVersionsParams {
	var ()
	return &GetK8SVersionsParams{

		timeout: timeout,
	}
}

// NewGetK8SVersionsParamsWithContext creates a new GetK8SVersionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetK8SVersionsParamsWithContext(ctx context.Context) *GetK8SVersionsParams {
	var ()
	return &GetK8SVersionsParams{

		Context: ctx,
	}
}

// NewGetK8SVersionsParamsWithHTTPClient creates a new GetK8SVersionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetK8SVersionsParamsWithHTTPClient(client *http.Client) *GetK8SVersionsParams {
	var ()
	return &GetK8SVersionsParams{
		HTTPClient: client,
	}
}

/*GetK8SVersionsParams contains all the parameters to send to the API endpoint
for the get k8 s versions operation typically these are written to a http.Request
*/
type GetK8SVersionsParams struct {

	/*Body
	  Body parameters.

	*/
	Body *models.ClustersapiGetK8SVersionsRequest
	/*OwnerID
	  Wercker organization ID.

	Example: 59bbf8affb1beb01009cae7a

	*/
	OwnerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get k8 s versions params
func (o *GetK8SVersionsParams) WithTimeout(timeout time.Duration) *GetK8SVersionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get k8 s versions params
func (o *GetK8SVersionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get k8 s versions params
func (o *GetK8SVersionsParams) WithContext(ctx context.Context) *GetK8SVersionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get k8 s versions params
func (o *GetK8SVersionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get k8 s versions params
func (o *GetK8SVersionsParams) WithHTTPClient(client *http.Client) *GetK8SVersionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get k8 s versions params
func (o *GetK8SVersionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get k8 s versions params
func (o *GetK8SVersionsParams) WithBody(body *models.ClustersapiGetK8SVersionsRequest) *GetK8SVersionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get k8 s versions params
func (o *GetK8SVersionsParams) SetBody(body *models.ClustersapiGetK8SVersionsRequest) {
	o.Body = body
}

// WithOwnerID adds the ownerID to the get k8 s versions params
func (o *GetK8SVersionsParams) WithOwnerID(ownerID string) *GetK8SVersionsParams {
	o.SetOwnerID(ownerID)
	return o
}

// SetOwnerID adds the ownerId to the get k8 s versions params
func (o *GetK8SVersionsParams) SetOwnerID(ownerID string) {
	o.OwnerID = ownerID
}

// WriteToRequest writes these params to a swagger request
func (o *GetK8SVersionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
