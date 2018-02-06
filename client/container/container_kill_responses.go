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

// ContainerKillReader is a Reader for the ContainerKill structure.
type ContainerKillReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerKillReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewContainerKillNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewContainerKillNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewContainerKillInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewContainerKillNoContent creates a ContainerKillNoContent with default headers values
func NewContainerKillNoContent() *ContainerKillNoContent {
	return &ContainerKillNoContent{}
}

/*ContainerKillNoContent handles this case with default header values.

no error
*/
type ContainerKillNoContent struct {
}

func (o *ContainerKillNoContent) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/kill][%d] containerKillNoContent ", 204)
}

func (o *ContainerKillNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerKillNotFound creates a ContainerKillNotFound with default headers values
func NewContainerKillNotFound() *ContainerKillNotFound {
	return &ContainerKillNotFound{}
}

/*ContainerKillNotFound handles this case with default header values.

no such container
*/
type ContainerKillNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ContainerKillNotFound) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/kill][%d] containerKillNotFound  %+v", 404, o.Payload)
}

func (o *ContainerKillNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerKillInternalServerError creates a ContainerKillInternalServerError with default headers values
func NewContainerKillInternalServerError() *ContainerKillInternalServerError {
	return &ContainerKillInternalServerError{}
}

/*ContainerKillInternalServerError handles this case with default header values.

server error
*/
type ContainerKillInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ContainerKillInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/kill][%d] containerKillInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerKillInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
