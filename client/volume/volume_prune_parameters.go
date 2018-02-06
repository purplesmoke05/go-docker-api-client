// Code generated by go-swagger; DO NOT EDIT.

package volume

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

// NewVolumePruneParams creates a new VolumePruneParams object
// with the default values initialized.
func NewVolumePruneParams() *VolumePruneParams {
	var ()
	return &VolumePruneParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVolumePruneParamsWithTimeout creates a new VolumePruneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVolumePruneParamsWithTimeout(timeout time.Duration) *VolumePruneParams {
	var ()
	return &VolumePruneParams{

		timeout: timeout,
	}
}

// NewVolumePruneParamsWithContext creates a new VolumePruneParams object
// with the default values initialized, and the ability to set a context for a request
func NewVolumePruneParamsWithContext(ctx context.Context) *VolumePruneParams {
	var ()
	return &VolumePruneParams{

		Context: ctx,
	}
}

// NewVolumePruneParamsWithHTTPClient creates a new VolumePruneParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVolumePruneParamsWithHTTPClient(client *http.Client) *VolumePruneParams {
	var ()
	return &VolumePruneParams{
		HTTPClient: client,
	}
}

/*VolumePruneParams contains all the parameters to send to the API endpoint
for the volume prune operation typically these are written to a http.Request
*/
type VolumePruneParams struct {

	/*Filters
	  Filters to process on the prune list, encoded as JSON (a `map[string][]string`).

	Available filters:
	- `label` (`label=<key>`, `label=<key>=<value>`, `label!=<key>`, or `label!=<key>=<value>`) Prune volumes with (or without, in case `label!=...` is used) the specified labels.


	*/
	Filters *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the volume prune params
func (o *VolumePruneParams) WithTimeout(timeout time.Duration) *VolumePruneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the volume prune params
func (o *VolumePruneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the volume prune params
func (o *VolumePruneParams) WithContext(ctx context.Context) *VolumePruneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the volume prune params
func (o *VolumePruneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the volume prune params
func (o *VolumePruneParams) WithHTTPClient(client *http.Client) *VolumePruneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the volume prune params
func (o *VolumePruneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilters adds the filters to the volume prune params
func (o *VolumePruneParams) WithFilters(filters *string) *VolumePruneParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the volume prune params
func (o *VolumePruneParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WriteToRequest writes these params to a swagger request
func (o *VolumePruneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filters != nil {

		// query param filters
		var qrFilters string
		if o.Filters != nil {
			qrFilters = *o.Filters
		}
		qFilters := qrFilters
		if qFilters != "" {
			if err := r.SetQueryParam("filters", qFilters); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
