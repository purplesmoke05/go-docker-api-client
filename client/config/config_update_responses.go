// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/airking05/go-docker-api-client/models"
)

// ConfigUpdateReader is a Reader for the ConfigUpdate structure.
type ConfigUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConfigUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewConfigUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewConfigUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewConfigUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewConfigUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewConfigUpdateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewConfigUpdateOK creates a ConfigUpdateOK with default headers values
func NewConfigUpdateOK() *ConfigUpdateOK {
	return &ConfigUpdateOK{}
}

/*ConfigUpdateOK handles this case with default header values.

no error
*/
type ConfigUpdateOK struct {
}

func (o *ConfigUpdateOK) Error() string {
	return fmt.Sprintf("[POST /configs/{id}/update][%d] configUpdateOK ", 200)
}

func (o *ConfigUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConfigUpdateBadRequest creates a ConfigUpdateBadRequest with default headers values
func NewConfigUpdateBadRequest() *ConfigUpdateBadRequest {
	return &ConfigUpdateBadRequest{}
}

/*ConfigUpdateBadRequest handles this case with default header values.

bad parameter
*/
type ConfigUpdateBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *ConfigUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /configs/{id}/update][%d] configUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ConfigUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigUpdateNotFound creates a ConfigUpdateNotFound with default headers values
func NewConfigUpdateNotFound() *ConfigUpdateNotFound {
	return &ConfigUpdateNotFound{}
}

/*ConfigUpdateNotFound handles this case with default header values.

no such config
*/
type ConfigUpdateNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ConfigUpdateNotFound) Error() string {
	return fmt.Sprintf("[POST /configs/{id}/update][%d] configUpdateNotFound  %+v", 404, o.Payload)
}

func (o *ConfigUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigUpdateInternalServerError creates a ConfigUpdateInternalServerError with default headers values
func NewConfigUpdateInternalServerError() *ConfigUpdateInternalServerError {
	return &ConfigUpdateInternalServerError{}
}

/*ConfigUpdateInternalServerError handles this case with default header values.

server error
*/
type ConfigUpdateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ConfigUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /configs/{id}/update][%d] configUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *ConfigUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigUpdateServiceUnavailable creates a ConfigUpdateServiceUnavailable with default headers values
func NewConfigUpdateServiceUnavailable() *ConfigUpdateServiceUnavailable {
	return &ConfigUpdateServiceUnavailable{}
}

/*ConfigUpdateServiceUnavailable handles this case with default header values.

node is not part of a swarm
*/
type ConfigUpdateServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *ConfigUpdateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /configs/{id}/update][%d] configUpdateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *ConfigUpdateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}