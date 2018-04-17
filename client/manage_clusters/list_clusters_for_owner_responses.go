// Code generated by go-swagger; DO NOT EDIT.

package manage_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/matyix/oracle-oke-client/models"
)

// ListClustersForOwnerReader is a Reader for the ListClustersForOwner structure.
type ListClustersForOwnerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListClustersForOwnerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListClustersForOwnerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListClustersForOwnerOK creates a ListClustersForOwnerOK with default headers values
func NewListClustersForOwnerOK() *ListClustersForOwnerOK {
	return &ListClustersForOwnerOK{}
}

/*ListClustersForOwnerOK handles this case with default header values.

List of cluster objects for the organization specified by ownerId.
*/
type ListClustersForOwnerOK struct {
	Payload *models.ClustersapiListResponse
}

func (o *ListClustersForOwnerOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/clusters/{ownerId}][%d] listClustersForOwnerOK  %+v", 200, o.Payload)
}

func (o *ListClustersForOwnerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClustersapiListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}