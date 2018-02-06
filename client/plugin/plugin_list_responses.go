// Code generated by go-swagger; DO NOT EDIT.

package plugin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/airking05/accepta/models"
)

// PluginListReader is a Reader for the PluginList structure.
type PluginListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPluginListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewPluginListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPluginListOK creates a PluginListOK with default headers values
func NewPluginListOK() *PluginListOK {
	return &PluginListOK{}
}

/*PluginListOK handles this case with default header values.

No error
*/
type PluginListOK struct {
	Payload models.PluginListOKBody
}

func (o *PluginListOK) Error() string {
	return fmt.Sprintf("[GET /plugins][%d] pluginListOK  %+v", 200, o.Payload)
}

func (o *PluginListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPluginListInternalServerError creates a PluginListInternalServerError with default headers values
func NewPluginListInternalServerError() *PluginListInternalServerError {
	return &PluginListInternalServerError{}
}

/*PluginListInternalServerError handles this case with default header values.

Server error
*/
type PluginListInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *PluginListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /plugins][%d] pluginListInternalServerError  %+v", 500, o.Payload)
}

func (o *PluginListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
