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

// NewDevelopAdminPassAppVersionParams creates a new DevelopAdminPassAppVersionParams object
// with the default values initialized.
func NewDevelopAdminPassAppVersionParams() *DevelopAdminPassAppVersionParams {
	var ()
	return &DevelopAdminPassAppVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDevelopAdminPassAppVersionParamsWithTimeout creates a new DevelopAdminPassAppVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDevelopAdminPassAppVersionParamsWithTimeout(timeout time.Duration) *DevelopAdminPassAppVersionParams {
	var ()
	return &DevelopAdminPassAppVersionParams{

		timeout: timeout,
	}
}

// NewDevelopAdminPassAppVersionParamsWithContext creates a new DevelopAdminPassAppVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewDevelopAdminPassAppVersionParamsWithContext(ctx context.Context) *DevelopAdminPassAppVersionParams {
	var ()
	return &DevelopAdminPassAppVersionParams{

		Context: ctx,
	}
}

// NewDevelopAdminPassAppVersionParamsWithHTTPClient creates a new DevelopAdminPassAppVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDevelopAdminPassAppVersionParamsWithHTTPClient(client *http.Client) *DevelopAdminPassAppVersionParams {
	var ()
	return &DevelopAdminPassAppVersionParams{
		HTTPClient: client,
	}
}

/*DevelopAdminPassAppVersionParams contains all the parameters to send to the API endpoint
for the develop admin pass app version operation typically these are written to a http.Request
*/
type DevelopAdminPassAppVersionParams struct {

	/*Body*/
	Body *models.OpenpitrixPassAppVersionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the develop admin pass app version params
func (o *DevelopAdminPassAppVersionParams) WithTimeout(timeout time.Duration) *DevelopAdminPassAppVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the develop admin pass app version params
func (o *DevelopAdminPassAppVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the develop admin pass app version params
func (o *DevelopAdminPassAppVersionParams) WithContext(ctx context.Context) *DevelopAdminPassAppVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the develop admin pass app version params
func (o *DevelopAdminPassAppVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the develop admin pass app version params
func (o *DevelopAdminPassAppVersionParams) WithHTTPClient(client *http.Client) *DevelopAdminPassAppVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the develop admin pass app version params
func (o *DevelopAdminPassAppVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the develop admin pass app version params
func (o *DevelopAdminPassAppVersionParams) WithBody(body *models.OpenpitrixPassAppVersionRequest) *DevelopAdminPassAppVersionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the develop admin pass app version params
func (o *DevelopAdminPassAppVersionParams) SetBody(body *models.OpenpitrixPassAppVersionRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DevelopAdminPassAppVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
