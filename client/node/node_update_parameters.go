// Code generated by go-swagger; DO NOT EDIT.

package node

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

	models "github.com/airking05/accepta/models"
)

// NewNodeUpdateParams creates a new NodeUpdateParams object
// with the default values initialized.
func NewNodeUpdateParams() *NodeUpdateParams {
	var ()
	return &NodeUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNodeUpdateParamsWithTimeout creates a new NodeUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNodeUpdateParamsWithTimeout(timeout time.Duration) *NodeUpdateParams {
	var ()
	return &NodeUpdateParams{

		timeout: timeout,
	}
}

// NewNodeUpdateParamsWithContext creates a new NodeUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewNodeUpdateParamsWithContext(ctx context.Context) *NodeUpdateParams {
	var ()
	return &NodeUpdateParams{

		Context: ctx,
	}
}

// NewNodeUpdateParamsWithHTTPClient creates a new NodeUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNodeUpdateParamsWithHTTPClient(client *http.Client) *NodeUpdateParams {
	var ()
	return &NodeUpdateParams{
		HTTPClient: client,
	}
}

/*NodeUpdateParams contains all the parameters to send to the API endpoint
for the node update operation typically these are written to a http.Request
*/
type NodeUpdateParams struct {

	/*Body*/
	Body *models.NodeSpec
	/*ID
	  The ID of the node

	*/
	ID string
	/*Version
	  The version number of the node object being updated. This is required to avoid conflicting writes.

	*/
	Version int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the node update params
func (o *NodeUpdateParams) WithTimeout(timeout time.Duration) *NodeUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the node update params
func (o *NodeUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the node update params
func (o *NodeUpdateParams) WithContext(ctx context.Context) *NodeUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the node update params
func (o *NodeUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the node update params
func (o *NodeUpdateParams) WithHTTPClient(client *http.Client) *NodeUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the node update params
func (o *NodeUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the node update params
func (o *NodeUpdateParams) WithBody(body *models.NodeSpec) *NodeUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the node update params
func (o *NodeUpdateParams) SetBody(body *models.NodeSpec) {
	o.Body = body
}

// WithID adds the id to the node update params
func (o *NodeUpdateParams) WithID(id string) *NodeUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the node update params
func (o *NodeUpdateParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the node update params
func (o *NodeUpdateParams) WithVersion(version int64) *NodeUpdateParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the node update params
func (o *NodeUpdateParams) SetVersion(version int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *NodeUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := swag.FormatInt64(qrVersion)
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
