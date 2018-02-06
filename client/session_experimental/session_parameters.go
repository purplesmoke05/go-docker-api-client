// Code generated by go-swagger; DO NOT EDIT.

package session_experimental

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

// NewSessionParams creates a new SessionParams object
// with the default values initialized.
func NewSessionParams() *SessionParams {

	return &SessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSessionParamsWithTimeout creates a new SessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSessionParamsWithTimeout(timeout time.Duration) *SessionParams {

	return &SessionParams{

		timeout: timeout,
	}
}

// NewSessionParamsWithContext creates a new SessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewSessionParamsWithContext(ctx context.Context) *SessionParams {

	return &SessionParams{

		Context: ctx,
	}
}

// NewSessionParamsWithHTTPClient creates a new SessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSessionParamsWithHTTPClient(client *http.Client) *SessionParams {

	return &SessionParams{
		HTTPClient: client,
	}
}

/*SessionParams contains all the parameters to send to the API endpoint
for the session operation typically these are written to a http.Request
*/
type SessionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the session params
func (o *SessionParams) WithTimeout(timeout time.Duration) *SessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the session params
func (o *SessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the session params
func (o *SessionParams) WithContext(ctx context.Context) *SessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the session params
func (o *SessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the session params
func (o *SessionParams) WithHTTPClient(client *http.Client) *SessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the session params
func (o *SessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
