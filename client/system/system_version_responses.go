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

// SystemVersionReader is a Reader for the SystemVersion structure.
type SystemVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSystemVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewSystemVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSystemVersionOK creates a SystemVersionOK with default headers values
func NewSystemVersionOK() *SystemVersionOK {
	return &SystemVersionOK{}
}

/*SystemVersionOK handles this case with default header values.

no error
*/
type SystemVersionOK struct {
	Payload *models.SystemVersionOKBody
}

func (o *SystemVersionOK) Error() string {
	return fmt.Sprintf("[GET /version][%d] systemVersionOK  %+v", 200, o.Payload)
}

func (o *SystemVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemVersionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSystemVersionInternalServerError creates a SystemVersionInternalServerError with default headers values
func NewSystemVersionInternalServerError() *SystemVersionInternalServerError {
	return &SystemVersionInternalServerError{}
}

/*SystemVersionInternalServerError handles this case with default header values.

server error
*/
type SystemVersionInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SystemVersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /version][%d] systemVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *SystemVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
