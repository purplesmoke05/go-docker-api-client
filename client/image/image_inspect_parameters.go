// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewImageInspectParams creates a new ImageInspectParams object
// with the default values initialized.
func NewImageInspectParams() *ImageInspectParams {
	var ()
	return &ImageInspectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewImageInspectParamsWithTimeout creates a new ImageInspectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewImageInspectParamsWithTimeout(timeout time.Duration) *ImageInspectParams {
	var ()
	return &ImageInspectParams{

		timeout: timeout,
	}
}

// NewImageInspectParamsWithContext creates a new ImageInspectParams object
// with the default values initialized, and the ability to set a context for a request
func NewImageInspectParamsWithContext(ctx context.Context) *ImageInspectParams {
	var ()
	return &ImageInspectParams{

		Context: ctx,
	}
}

// NewImageInspectParamsWithHTTPClient creates a new ImageInspectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewImageInspectParamsWithHTTPClient(client *http.Client) *ImageInspectParams {
	var ()
	return &ImageInspectParams{
		HTTPClient: client,
	}
}

/*ImageInspectParams contains all the parameters to send to the API endpoint
for the image inspect operation typically these are written to a http.Request
*/
type ImageInspectParams struct {

	/*Name
	  Image name or id

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the image inspect params
func (o *ImageInspectParams) WithTimeout(timeout time.Duration) *ImageInspectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image inspect params
func (o *ImageInspectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image inspect params
func (o *ImageInspectParams) WithContext(ctx context.Context) *ImageInspectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image inspect params
func (o *ImageInspectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image inspect params
func (o *ImageInspectParams) WithHTTPClient(client *http.Client) *ImageInspectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image inspect params
func (o *ImageInspectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the image inspect params
func (o *ImageInspectParams) WithName(name string) *ImageInspectParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the image inspect params
func (o *ImageInspectParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ImageInspectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
