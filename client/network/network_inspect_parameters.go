// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewNetworkInspectParams creates a new NetworkInspectParams object
// with the default values initialized.
func NewNetworkInspectParams() *NetworkInspectParams {
	var (
		verboseDefault = bool(false)
	)
	return &NetworkInspectParams{
		Verbose: &verboseDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkInspectParamsWithTimeout creates a new NetworkInspectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNetworkInspectParamsWithTimeout(timeout time.Duration) *NetworkInspectParams {
	var (
		verboseDefault = bool(false)
	)
	return &NetworkInspectParams{
		Verbose: &verboseDefault,

		timeout: timeout,
	}
}

// NewNetworkInspectParamsWithContext creates a new NetworkInspectParams object
// with the default values initialized, and the ability to set a context for a request
func NewNetworkInspectParamsWithContext(ctx context.Context) *NetworkInspectParams {
	var (
		verboseDefault = bool(false)
	)
	return &NetworkInspectParams{
		Verbose: &verboseDefault,

		Context: ctx,
	}
}

// NewNetworkInspectParamsWithHTTPClient creates a new NetworkInspectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNetworkInspectParamsWithHTTPClient(client *http.Client) *NetworkInspectParams {
	var (
		verboseDefault = bool(false)
	)
	return &NetworkInspectParams{
		Verbose:    &verboseDefault,
		HTTPClient: client,
	}
}

/*NetworkInspectParams contains all the parameters to send to the API endpoint
for the network inspect operation typically these are written to a http.Request
*/
type NetworkInspectParams struct {

	/*ID
	  Network ID or name

	*/
	ID string
	/*Scope
	  Filter the network by scope (swarm, global, or local)

	*/
	Scope *string
	/*Verbose
	  Detailed inspect output for troubleshooting

	*/
	Verbose *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the network inspect params
func (o *NetworkInspectParams) WithTimeout(timeout time.Duration) *NetworkInspectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network inspect params
func (o *NetworkInspectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network inspect params
func (o *NetworkInspectParams) WithContext(ctx context.Context) *NetworkInspectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network inspect params
func (o *NetworkInspectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the network inspect params
func (o *NetworkInspectParams) WithHTTPClient(client *http.Client) *NetworkInspectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the network inspect params
func (o *NetworkInspectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the network inspect params
func (o *NetworkInspectParams) WithID(id string) *NetworkInspectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the network inspect params
func (o *NetworkInspectParams) SetID(id string) {
	o.ID = id
}

// WithScope adds the scope to the network inspect params
func (o *NetworkInspectParams) WithScope(scope *string) *NetworkInspectParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the network inspect params
func (o *NetworkInspectParams) SetScope(scope *string) {
	o.Scope = scope
}

// WithVerbose adds the verbose to the network inspect params
func (o *NetworkInspectParams) WithVerbose(verbose *bool) *NetworkInspectParams {
	o.SetVerbose(verbose)
	return o
}

// SetVerbose adds the verbose to the network inspect params
func (o *NetworkInspectParams) SetVerbose(verbose *bool) {
	o.Verbose = verbose
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkInspectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Scope != nil {

		// query param scope
		var qrScope string
		if o.Scope != nil {
			qrScope = *o.Scope
		}
		qScope := qrScope
		if qScope != "" {
			if err := r.SetQueryParam("scope", qScope); err != nil {
				return err
			}
		}

	}

	if o.Verbose != nil {

		// query param verbose
		var qrVerbose bool
		if o.Verbose != nil {
			qrVerbose = *o.Verbose
		}
		qVerbose := swag.FormatBool(qrVerbose)
		if qVerbose != "" {
			if err := r.SetQueryParam("verbose", qVerbose); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
