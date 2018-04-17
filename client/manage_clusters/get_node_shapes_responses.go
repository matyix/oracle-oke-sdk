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

// GetNodeShapesReader is a Reader for the GetNodeShapes structure.
type GetNodeShapesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNodeShapesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNodeShapesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNodeShapesOK creates a GetNodeShapesOK with default headers values
func NewGetNodeShapesOK() *GetNodeShapesOK {
	return &GetNodeShapesOK{}
}

/*GetNodeShapesOK handles this case with default header values.

List of supported VM shapes.
*/
type GetNodeShapesOK struct {
	Payload *models.ClustersapiGetNodeShapesResponse
}

func (o *GetNodeShapesOK) Error() string {
	return fmt.Sprintf("[POST /api/v3/clusters/{ownerId}/node/shapes][%d] getNodeShapesOK  %+v", 200, o.Payload)
}

func (o *GetNodeShapesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClustersapiGetNodeShapesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
