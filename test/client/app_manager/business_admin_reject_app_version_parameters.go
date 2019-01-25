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

// NewBusinessAdminRejectAppVersionParams creates a new BusinessAdminRejectAppVersionParams object
// with the default values initialized.
func NewBusinessAdminRejectAppVersionParams() *BusinessAdminRejectAppVersionParams {
	var ()
	return &BusinessAdminRejectAppVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBusinessAdminRejectAppVersionParamsWithTimeout creates a new BusinessAdminRejectAppVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBusinessAdminRejectAppVersionParamsWithTimeout(timeout time.Duration) *BusinessAdminRejectAppVersionParams {
	var ()
	return &BusinessAdminRejectAppVersionParams{

		timeout: timeout,
	}
}

// NewBusinessAdminRejectAppVersionParamsWithContext creates a new BusinessAdminRejectAppVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewBusinessAdminRejectAppVersionParamsWithContext(ctx context.Context) *BusinessAdminRejectAppVersionParams {
	var ()
	return &BusinessAdminRejectAppVersionParams{

		Context: ctx,
	}
}

// NewBusinessAdminRejectAppVersionParamsWithHTTPClient creates a new BusinessAdminRejectAppVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBusinessAdminRejectAppVersionParamsWithHTTPClient(client *http.Client) *BusinessAdminRejectAppVersionParams {
	var ()
	return &BusinessAdminRejectAppVersionParams{
		HTTPClient: client,
	}
}

/*BusinessAdminRejectAppVersionParams contains all the parameters to send to the API endpoint
for the business admin reject app version operation typically these are written to a http.Request
*/
type BusinessAdminRejectAppVersionParams struct {

	/*Body*/
	Body *models.OpenpitrixRejectAppVersionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the business admin reject app version params
func (o *BusinessAdminRejectAppVersionParams) WithTimeout(timeout time.Duration) *BusinessAdminRejectAppVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the business admin reject app version params
func (o *BusinessAdminRejectAppVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the business admin reject app version params
func (o *BusinessAdminRejectAppVersionParams) WithContext(ctx context.Context) *BusinessAdminRejectAppVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the business admin reject app version params
func (o *BusinessAdminRejectAppVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the business admin reject app version params
func (o *BusinessAdminRejectAppVersionParams) WithHTTPClient(client *http.Client) *BusinessAdminRejectAppVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the business admin reject app version params
func (o *BusinessAdminRejectAppVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the business admin reject app version params
func (o *BusinessAdminRejectAppVersionParams) WithBody(body *models.OpenpitrixRejectAppVersionRequest) *BusinessAdminRejectAppVersionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the business admin reject app version params
func (o *BusinessAdminRejectAppVersionParams) SetBody(body *models.OpenpitrixRejectAppVersionRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *BusinessAdminRejectAppVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
