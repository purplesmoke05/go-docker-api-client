// Code generated by go-swagger; DO NOT EDIT.

package config

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

	models "github.com/airking05/accepta/models"
)

// NewConfigCreateParams creates a new ConfigCreateParams object
// with the default values initialized.
func NewConfigCreateParams() *ConfigCreateParams {
	var ()
	return &ConfigCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConfigCreateParamsWithTimeout creates a new ConfigCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConfigCreateParamsWithTimeout(timeout time.Duration) *ConfigCreateParams {
	var ()
	return &ConfigCreateParams{

		timeout: timeout,
	}
}

// NewConfigCreateParamsWithContext creates a new ConfigCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewConfigCreateParamsWithContext(ctx context.Context) *ConfigCreateParams {
	var ()
	return &ConfigCreateParams{

		Context: ctx,
	}
}

// NewConfigCreateParamsWithHTTPClient creates a new ConfigCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewConfigCreateParamsWithHTTPClient(client *http.Client) *ConfigCreateParams {
	var ()
	return &ConfigCreateParams{
		HTTPClient: client,
	}
}

/*ConfigCreateParams contains all the parameters to send to the API endpoint
for the config create operation typically these are written to a http.Request
*/
type ConfigCreateParams struct {

	/*Body*/
	Body *models.ConfigCreateParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the config create params
func (o *ConfigCreateParams) WithTimeout(timeout time.Duration) *ConfigCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the config create params
func (o *ConfigCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the config create params
func (o *ConfigCreateParams) WithContext(ctx context.Context) *ConfigCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the config create params
func (o *ConfigCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the config create params
func (o *ConfigCreateParams) WithHTTPClient(client *http.Client) *ConfigCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the config create params
func (o *ConfigCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the config create params
func (o *ConfigCreateParams) WithBody(body *models.ConfigCreateParamsBody) *ConfigCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the config create params
func (o *ConfigCreateParams) SetBody(body *models.ConfigCreateParamsBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ConfigCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
