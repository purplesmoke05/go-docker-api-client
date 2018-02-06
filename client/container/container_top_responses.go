// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/airking05/accepta/models"
)

// ContainerTopReader is a Reader for the ContainerTop structure.
type ContainerTopReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerTopReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewContainerTopOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewContainerTopNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewContainerTopInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewContainerTopOK creates a ContainerTopOK with default headers values
func NewContainerTopOK() *ContainerTopOK {
	return &ContainerTopOK{}
}

/*ContainerTopOK handles this case with default header values.

no error
*/
type ContainerTopOK struct {
	Payload *models.ContainerTopOKBody
}

func (o *ContainerTopOK) Error() string {
	return fmt.Sprintf("[GET /containers/{id}/top][%d] containerTopOK  %+v", 200, o.Payload)
}

func (o *ContainerTopOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContainerTopOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerTopNotFound creates a ContainerTopNotFound with default headers values
func NewContainerTopNotFound() *ContainerTopNotFound {
	return &ContainerTopNotFound{}
}

/*ContainerTopNotFound handles this case with default header values.

no such container
*/
type ContainerTopNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ContainerTopNotFound) Error() string {
	return fmt.Sprintf("[GET /containers/{id}/top][%d] containerTopNotFound  %+v", 404, o.Payload)
}

func (o *ContainerTopNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerTopInternalServerError creates a ContainerTopInternalServerError with default headers values
func NewContainerTopInternalServerError() *ContainerTopInternalServerError {
	return &ContainerTopInternalServerError{}
}

/*ContainerTopInternalServerError handles this case with default header values.

server error
*/
type ContainerTopInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ContainerTopInternalServerError) Error() string {
	return fmt.Sprintf("[GET /containers/{id}/top][%d] containerTopInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerTopInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
