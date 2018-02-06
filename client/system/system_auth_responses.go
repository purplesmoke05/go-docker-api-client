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

// SystemAuthReader is a Reader for the SystemAuth structure.
type SystemAuthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemAuthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSystemAuthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewSystemAuthNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewSystemAuthInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSystemAuthOK creates a SystemAuthOK with default headers values
func NewSystemAuthOK() *SystemAuthOK {
	return &SystemAuthOK{}
}

/*SystemAuthOK handles this case with default header values.

An identity token was generated successfully.
*/
type SystemAuthOK struct {
	Payload *models.SystemAuthOKBody
}

func (o *SystemAuthOK) Error() string {
	return fmt.Sprintf("[POST /auth][%d] systemAuthOK  %+v", 200, o.Payload)
}

func (o *SystemAuthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemAuthOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSystemAuthNoContent creates a SystemAuthNoContent with default headers values
func NewSystemAuthNoContent() *SystemAuthNoContent {
	return &SystemAuthNoContent{}
}

/*SystemAuthNoContent handles this case with default header values.

No error
*/
type SystemAuthNoContent struct {
}

func (o *SystemAuthNoContent) Error() string {
	return fmt.Sprintf("[POST /auth][%d] systemAuthNoContent ", 204)
}

func (o *SystemAuthNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSystemAuthInternalServerError creates a SystemAuthInternalServerError with default headers values
func NewSystemAuthInternalServerError() *SystemAuthInternalServerError {
	return &SystemAuthInternalServerError{}
}

/*SystemAuthInternalServerError handles this case with default header values.

Server error
*/
type SystemAuthInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SystemAuthInternalServerError) Error() string {
	return fmt.Sprintf("[POST /auth][%d] systemAuthInternalServerError  %+v", 500, o.Payload)
}

func (o *SystemAuthInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
