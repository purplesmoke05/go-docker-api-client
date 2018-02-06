// Code generated by go-swagger; DO NOT EDIT.

package volume

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/airking05/accepta/models"
)

// VolumeListReader is a Reader for the VolumeList structure.
type VolumeListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVolumeListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewVolumeListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVolumeListOK creates a VolumeListOK with default headers values
func NewVolumeListOK() *VolumeListOK {
	return &VolumeListOK{}
}

/*VolumeListOK handles this case with default header values.

Summary volume data that matches the query
*/
type VolumeListOK struct {
	Payload *models.VolumeListOKBody
}

func (o *VolumeListOK) Error() string {
	return fmt.Sprintf("[GET /volumes][%d] volumeListOK  %+v", 200, o.Payload)
}

func (o *VolumeListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeListInternalServerError creates a VolumeListInternalServerError with default headers values
func NewVolumeListInternalServerError() *VolumeListInternalServerError {
	return &VolumeListInternalServerError{}
}

/*VolumeListInternalServerError handles this case with default header values.

Server error
*/
type VolumeListInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *VolumeListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /volumes][%d] volumeListInternalServerError  %+v", 500, o.Payload)
}

func (o *VolumeListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
