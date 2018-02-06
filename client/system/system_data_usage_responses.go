// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/airking05/accepta/models"
)

// SystemDataUsageReader is a Reader for the SystemDataUsage structure.
type SystemDataUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemDataUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSystemDataUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewSystemDataUsageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSystemDataUsageOK creates a SystemDataUsageOK with default headers values
func NewSystemDataUsageOK() *SystemDataUsageOK {
	return &SystemDataUsageOK{}
}

/*SystemDataUsageOK handles this case with default header values.

no error
*/
type SystemDataUsageOK struct {
	Payload *models.SystemDataUsageOKBody
}

func (o *SystemDataUsageOK) Error() string {
	return fmt.Sprintf("[GET /system/df][%d] systemDataUsageOK  %+v", 200, o.Payload)
}

func (o *SystemDataUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemDataUsageOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSystemDataUsageInternalServerError creates a SystemDataUsageInternalServerError with default headers values
func NewSystemDataUsageInternalServerError() *SystemDataUsageInternalServerError {
	return &SystemDataUsageInternalServerError{}
}

/*SystemDataUsageInternalServerError handles this case with default header values.

server error
*/
type SystemDataUsageInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SystemDataUsageInternalServerError) Error() string {
	return fmt.Sprintf("[GET /system/df][%d] systemDataUsageInternalServerError  %+v", 500, o.Payload)
}

func (o *SystemDataUsageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
