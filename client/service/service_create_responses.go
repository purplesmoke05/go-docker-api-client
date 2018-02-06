// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/airking05/accepta/models"
)

// ServiceCreateReader is a Reader for the ServiceCreate structure.
type ServiceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewServiceCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewServiceCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewServiceCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewServiceCreateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewServiceCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewServiceCreateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServiceCreateCreated creates a ServiceCreateCreated with default headers values
func NewServiceCreateCreated() *ServiceCreateCreated {
	return &ServiceCreateCreated{}
}

/*ServiceCreateCreated handles this case with default header values.

no error
*/
type ServiceCreateCreated struct {
	Payload *models.ServiceCreateCreatedBody
}

func (o *ServiceCreateCreated) Error() string {
	return fmt.Sprintf("[POST /services/create][%d] serviceCreateCreated  %+v", 201, o.Payload)
}

func (o *ServiceCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceCreateCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceCreateBadRequest creates a ServiceCreateBadRequest with default headers values
func NewServiceCreateBadRequest() *ServiceCreateBadRequest {
	return &ServiceCreateBadRequest{}
}

/*ServiceCreateBadRequest handles this case with default header values.

bad parameter
*/
type ServiceCreateBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *ServiceCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/create][%d] serviceCreateBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceCreateForbidden creates a ServiceCreateForbidden with default headers values
func NewServiceCreateForbidden() *ServiceCreateForbidden {
	return &ServiceCreateForbidden{}
}

/*ServiceCreateForbidden handles this case with default header values.

network is not eligible for services
*/
type ServiceCreateForbidden struct {
	Payload *models.ErrorResponse
}

func (o *ServiceCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /services/create][%d] serviceCreateForbidden  %+v", 403, o.Payload)
}

func (o *ServiceCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceCreateConflict creates a ServiceCreateConflict with default headers values
func NewServiceCreateConflict() *ServiceCreateConflict {
	return &ServiceCreateConflict{}
}

/*ServiceCreateConflict handles this case with default header values.

name conflicts with an existing service
*/
type ServiceCreateConflict struct {
	Payload *models.ErrorResponse
}

func (o *ServiceCreateConflict) Error() string {
	return fmt.Sprintf("[POST /services/create][%d] serviceCreateConflict  %+v", 409, o.Payload)
}

func (o *ServiceCreateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceCreateInternalServerError creates a ServiceCreateInternalServerError with default headers values
func NewServiceCreateInternalServerError() *ServiceCreateInternalServerError {
	return &ServiceCreateInternalServerError{}
}

/*ServiceCreateInternalServerError handles this case with default header values.

server error
*/
type ServiceCreateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ServiceCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /services/create][%d] serviceCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *ServiceCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceCreateServiceUnavailable creates a ServiceCreateServiceUnavailable with default headers values
func NewServiceCreateServiceUnavailable() *ServiceCreateServiceUnavailable {
	return &ServiceCreateServiceUnavailable{}
}

/*ServiceCreateServiceUnavailable handles this case with default header values.

node is not part of a swarm
*/
type ServiceCreateServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *ServiceCreateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /services/create][%d] serviceCreateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ServiceCreateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
