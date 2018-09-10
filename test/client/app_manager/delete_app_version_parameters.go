// Code generated by go-swagger; DO NOT EDIT.

package app_manager

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

	"openpitrix.io/openpitrix/test/models"
)

// NewDeleteAppVersionParams creates a new DeleteAppVersionParams object
// with the default values initialized.
func NewDeleteAppVersionParams() *DeleteAppVersionParams {
	var ()
	return &DeleteAppVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAppVersionParamsWithTimeout creates a new DeleteAppVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAppVersionParamsWithTimeout(timeout time.Duration) *DeleteAppVersionParams {
	var ()
	return &DeleteAppVersionParams{

		timeout: timeout,
	}
}

// NewDeleteAppVersionParamsWithContext creates a new DeleteAppVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAppVersionParamsWithContext(ctx context.Context) *DeleteAppVersionParams {
	var ()
	return &DeleteAppVersionParams{

		Context: ctx,
	}
}

// NewDeleteAppVersionParamsWithHTTPClient creates a new DeleteAppVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAppVersionParamsWithHTTPClient(client *http.Client) *DeleteAppVersionParams {
	var ()
	return &DeleteAppVersionParams{
		HTTPClient: client,
	}
}

/*DeleteAppVersionParams contains all the parameters to send to the API endpoint
for the delete app version operation typically these are written to a http.Request
*/
type DeleteAppVersionParams struct {

	/*Body*/
	Body *models.OpenpitrixDeleteAppVersionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete app version params
func (o *DeleteAppVersionParams) WithTimeout(timeout time.Duration) *DeleteAppVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete app version params
func (o *DeleteAppVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete app version params
func (o *DeleteAppVersionParams) WithContext(ctx context.Context) *DeleteAppVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete app version params
func (o *DeleteAppVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete app version params
func (o *DeleteAppVersionParams) WithHTTPClient(client *http.Client) *DeleteAppVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete app version params
func (o *DeleteAppVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete app version params
func (o *DeleteAppVersionParams) WithBody(body *models.OpenpitrixDeleteAppVersionRequest) *DeleteAppVersionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete app version params
func (o *DeleteAppVersionParams) SetBody(body *models.OpenpitrixDeleteAppVersionRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAppVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
