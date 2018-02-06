// Code generated by go-swagger; DO NOT EDIT.

package exec

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

// NewExecStartParams creates a new ExecStartParams object
// with the default values initialized.
func NewExecStartParams() *ExecStartParams {
	var ()
	return &ExecStartParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExecStartParamsWithTimeout creates a new ExecStartParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExecStartParamsWithTimeout(timeout time.Duration) *ExecStartParams {
	var ()
	return &ExecStartParams{

		timeout: timeout,
	}
}

// NewExecStartParamsWithContext creates a new ExecStartParams object
// with the default values initialized, and the ability to set a context for a request
func NewExecStartParamsWithContext(ctx context.Context) *ExecStartParams {
	var ()
	return &ExecStartParams{

		Context: ctx,
	}
}

// NewExecStartParamsWithHTTPClient creates a new ExecStartParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExecStartParamsWithHTTPClient(client *http.Client) *ExecStartParams {
	var ()
	return &ExecStartParams{
		HTTPClient: client,
	}
}

/*ExecStartParams contains all the parameters to send to the API endpoint
for the exec start operation typically these are written to a http.Request
*/
type ExecStartParams struct {

	/*ExecStartConfig*/
	ExecStartConfig *models.ExecStartParamsBody
	/*ID
	  Exec instance ID

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the exec start params
func (o *ExecStartParams) WithTimeout(timeout time.Duration) *ExecStartParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the exec start params
func (o *ExecStartParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the exec start params
func (o *ExecStartParams) WithContext(ctx context.Context) *ExecStartParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the exec start params
func (o *ExecStartParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the exec start params
func (o *ExecStartParams) WithHTTPClient(client *http.Client) *ExecStartParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the exec start params
func (o *ExecStartParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExecStartConfig adds the execStartConfig to the exec start params
func (o *ExecStartParams) WithExecStartConfig(execStartConfig *models.ExecStartParamsBody) *ExecStartParams {
	o.SetExecStartConfig(execStartConfig)
	return o
}

// SetExecStartConfig adds the execStartConfig to the exec start params
func (o *ExecStartParams) SetExecStartConfig(execStartConfig *models.ExecStartParamsBody) {
	o.ExecStartConfig = execStartConfig
}

// WithID adds the id to the exec start params
func (o *ExecStartParams) WithID(id string) *ExecStartParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the exec start params
func (o *ExecStartParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExecStartParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ExecStartConfig != nil {
		if err := r.SetBodyParam(o.ExecStartConfig); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
