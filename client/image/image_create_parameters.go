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

// NewImageCreateParams creates a new ImageCreateParams object
// with the default values initialized.
func NewImageCreateParams() *ImageCreateParams {
	var ()
	return &ImageCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewImageCreateParamsWithTimeout creates a new ImageCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewImageCreateParamsWithTimeout(timeout time.Duration) *ImageCreateParams {
	var ()
	return &ImageCreateParams{

		timeout: timeout,
	}
}

// NewImageCreateParamsWithContext creates a new ImageCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewImageCreateParamsWithContext(ctx context.Context) *ImageCreateParams {
	var ()
	return &ImageCreateParams{

		Context: ctx,
	}
}

// NewImageCreateParamsWithHTTPClient creates a new ImageCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewImageCreateParamsWithHTTPClient(client *http.Client) *ImageCreateParams {
	var ()
	return &ImageCreateParams{
		HTTPClient: client,
	}
}

/*ImageCreateParams contains all the parameters to send to the API endpoint
for the image create operation typically these are written to a http.Request
*/
type ImageCreateParams struct {

	/*XRegistryAuth
	  A base64-encoded auth configuration. [See the authentication section for details.](#section/Authentication)

	*/
	XRegistryAuth *string
	/*FromImage
	  Name of the image to pull. The name may include a tag or digest. This parameter may only be used when pulling an image. The pull is cancelled if the HTTP connection is closed.

	*/
	FromImage *string
	/*FromSrc
	  Source to import. The value may be a URL from which the image can be retrieved or `-` to read the image from the request body. This parameter may only be used when importing an image.

	*/
	FromSrc *string
	/*InputImage
	  Image content if the value `-` has been specified in fromSrc query parameter

	*/
	InputImage string
	/*Repo
	  Repository name given to an image when it is imported. The repo may include a tag. This parameter may only be used when importing an image.

	*/
	Repo *string
	/*Tag
	  Tag or digest. If empty when pulling an image, this causes all tags for the given image to be pulled.

	*/
	Tag *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the image create params
func (o *ImageCreateParams) WithTimeout(timeout time.Duration) *ImageCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image create params
func (o *ImageCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image create params
func (o *ImageCreateParams) WithContext(ctx context.Context) *ImageCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image create params
func (o *ImageCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image create params
func (o *ImageCreateParams) WithHTTPClient(client *http.Client) *ImageCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image create params
func (o *ImageCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRegistryAuth adds the xRegistryAuth to the image create params
func (o *ImageCreateParams) WithXRegistryAuth(xRegistryAuth *string) *ImageCreateParams {
	o.SetXRegistryAuth(xRegistryAuth)
	return o
}

// SetXRegistryAuth adds the xRegistryAuth to the image create params
func (o *ImageCreateParams) SetXRegistryAuth(xRegistryAuth *string) {
	o.XRegistryAuth = xRegistryAuth
}

// WithFromImage adds the fromImage to the image create params
func (o *ImageCreateParams) WithFromImage(fromImage *string) *ImageCreateParams {
	o.SetFromImage(fromImage)
	return o
}

// SetFromImage adds the fromImage to the image create params
func (o *ImageCreateParams) SetFromImage(fromImage *string) {
	o.FromImage = fromImage
}

// WithFromSrc adds the fromSrc to the image create params
func (o *ImageCreateParams) WithFromSrc(fromSrc *string) *ImageCreateParams {
	o.SetFromSrc(fromSrc)
	return o
}

// SetFromSrc adds the fromSrc to the image create params
func (o *ImageCreateParams) SetFromSrc(fromSrc *string) {
	o.FromSrc = fromSrc
}

// WithInputImage adds the inputImage to the image create params
func (o *ImageCreateParams) WithInputImage(inputImage string) *ImageCreateParams {
	o.SetInputImage(inputImage)
	return o
}

// SetInputImage adds the inputImage to the image create params
func (o *ImageCreateParams) SetInputImage(inputImage string) {
	o.InputImage = inputImage
}

// WithRepo adds the repo to the image create params
func (o *ImageCreateParams) WithRepo(repo *string) *ImageCreateParams {
	o.SetRepo(repo)
	return o
}

// SetRepo adds the repo to the image create params
func (o *ImageCreateParams) SetRepo(repo *string) {
	o.Repo = repo
}

// WithTag adds the tag to the image create params
func (o *ImageCreateParams) WithTag(tag *string) *ImageCreateParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the image create params
func (o *ImageCreateParams) SetTag(tag *string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *ImageCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRegistryAuth != nil {

		// header param X-Registry-Auth
		if err := r.SetHeaderParam("X-Registry-Auth", *o.XRegistryAuth); err != nil {
			return err
		}

	}

	if o.FromImage != nil {

		// query param fromImage
		var qrFromImage string
		if o.FromImage != nil {
			qrFromImage = *o.FromImage
		}
		qFromImage := qrFromImage
		if qFromImage != "" {
			if err := r.SetQueryParam("fromImage", qFromImage); err != nil {
				return err
			}
		}

	}

	if o.FromSrc != nil {

		// query param fromSrc
		var qrFromSrc string
		if o.FromSrc != nil {
			qrFromSrc = *o.FromSrc
		}
		qFromSrc := qrFromSrc
		if qFromSrc != "" {
			if err := r.SetQueryParam("fromSrc", qFromSrc); err != nil {
				return err
			}
		}

	}

	if o.Repo != nil {

		// query param repo
		var qrRepo string
		if o.Repo != nil {
			qrRepo = *o.Repo
		}
		qRepo := qrRepo
		if qRepo != "" {
			if err := r.SetQueryParam("repo", qRepo); err != nil {
				return err
			}
		}

	}

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}