// Code generated by go-swagger; DO NOT EDIT.

package secret

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/airking05/accepta/models"
)

// SecretUpdateReader is a Reader for the SecretUpdate structure.
type SecretUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSecretUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSecretUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSecretUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSecretUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewSecretUpdateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSecretUpdateOK creates a SecretUpdateOK with default headers values
func NewSecretUpdateOK() *SecretUpdateOK {
	return &SecretUpdateOK{}
}

/*SecretUpdateOK handles this case with default header values.

no error
*/
type SecretUpdateOK struct {
}

func (o *SecretUpdateOK) Error() string {
	return fmt.Sprintf("[POST /secrets/{id}/update][%d] secretUpdateOK ", 200)
}

func (o *SecretUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSecretUpdateBadRequest creates a SecretUpdateBadRequest with default headers values
func NewSecretUpdateBadRequest() *SecretUpdateBadRequest {
	return &SecretUpdateBadRequest{}
}

/*SecretUpdateBadRequest handles this case with default header values.

bad parameter
*/
type SecretUpdateBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *SecretUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /secrets/{id}/update][%d] secretUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *SecretUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretUpdateNotFound creates a SecretUpdateNotFound with default headers values
func NewSecretUpdateNotFound() *SecretUpdateNotFound {
	return &SecretUpdateNotFound{}
}

/*SecretUpdateNotFound handles this case with default header values.

no such secret
*/
type SecretUpdateNotFound struct {
	Payload *models.ErrorResponse
}

func (o *SecretUpdateNotFound) Error() string {
	return fmt.Sprintf("[POST /secrets/{id}/update][%d] secretUpdateNotFound  %+v", 404, o.Payload)
}

func (o *SecretUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretUpdateInternalServerError creates a SecretUpdateInternalServerError with default headers values
func NewSecretUpdateInternalServerError() *SecretUpdateInternalServerError {
	return &SecretUpdateInternalServerError{}
}

/*SecretUpdateInternalServerError handles this case with default header values.

server error
*/
type SecretUpdateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SecretUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /secrets/{id}/update][%d] secretUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *SecretUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretUpdateServiceUnavailable creates a SecretUpdateServiceUnavailable with default headers values
func NewSecretUpdateServiceUnavailable() *SecretUpdateServiceUnavailable {
	return &SecretUpdateServiceUnavailable{}
}

/*SecretUpdateServiceUnavailable handles this case with default header values.

node is not part of a swarm
*/
type SecretUpdateServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *SecretUpdateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /secrets/{id}/update][%d] secretUpdateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *SecretUpdateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
