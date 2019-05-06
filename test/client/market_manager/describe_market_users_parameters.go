// Code generated by go-swagger; DO NOT EDIT.

package market_manager

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

// NewDescribeMarketUsersParams creates a new DescribeMarketUsersParams object
// with the default values initialized.
func NewDescribeMarketUsersParams() *DescribeMarketUsersParams {
	var ()
	return &DescribeMarketUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeMarketUsersParamsWithTimeout creates a new DescribeMarketUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeMarketUsersParamsWithTimeout(timeout time.Duration) *DescribeMarketUsersParams {
	var ()
	return &DescribeMarketUsersParams{

		timeout: timeout,
	}
}

// NewDescribeMarketUsersParamsWithContext creates a new DescribeMarketUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeMarketUsersParamsWithContext(ctx context.Context) *DescribeMarketUsersParams {
	var ()
	return &DescribeMarketUsersParams{

		Context: ctx,
	}
}

// NewDescribeMarketUsersParamsWithHTTPClient creates a new DescribeMarketUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeMarketUsersParamsWithHTTPClient(client *http.Client) *DescribeMarketUsersParams {
	var ()
	return &DescribeMarketUsersParams{
		HTTPClient: client,
	}
}

/*DescribeMarketUsersParams contains all the parameters to send to the API endpoint
for the describe market users operation typically these are written to a http.Request
*/
type DescribeMarketUsersParams struct {

	/*Limit
	  default is 20, max value is 200.

	*/
	Limit *int64
	/*MarketID*/
	MarketID []string
	/*Offset
	  default is 0.

	*/
	Offset *int64
	/*Owner*/
	Owner []string
	/*Reverse
	  value = 0 sort ASC, value = 1 sort DESC.

	*/
	Reverse *bool
	/*SearchWord*/
	SearchWord *string
	/*SortKey
	  sort key, order by sort_key, default create_time.

	*/
	SortKey *string
	/*UserID*/
	UserID []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe market users params
func (o *DescribeMarketUsersParams) WithTimeout(timeout time.Duration) *DescribeMarketUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe market users params
func (o *DescribeMarketUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe market users params
func (o *DescribeMarketUsersParams) WithContext(ctx context.Context) *DescribeMarketUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe market users params
func (o *DescribeMarketUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe market users params
func (o *DescribeMarketUsersParams) WithHTTPClient(client *http.Client) *DescribeMarketUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe market users params
func (o *DescribeMarketUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the describe market users params
func (o *DescribeMarketUsersParams) WithLimit(limit *int64) *DescribeMarketUsersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the describe market users params
func (o *DescribeMarketUsersParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMarketID adds the marketID to the describe market users params
func (o *DescribeMarketUsersParams) WithMarketID(marketID []string) *DescribeMarketUsersParams {
	o.SetMarketID(marketID)
	return o
}

// SetMarketID adds the marketId to the describe market users params
func (o *DescribeMarketUsersParams) SetMarketID(marketID []string) {
	o.MarketID = marketID
}

// WithOffset adds the offset to the describe market users params
func (o *DescribeMarketUsersParams) WithOffset(offset *int64) *DescribeMarketUsersParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the describe market users params
func (o *DescribeMarketUsersParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOwner adds the owner to the describe market users params
func (o *DescribeMarketUsersParams) WithOwner(owner []string) *DescribeMarketUsersParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the describe market users params
func (o *DescribeMarketUsersParams) SetOwner(owner []string) {
	o.Owner = owner
}

// WithReverse adds the reverse to the describe market users params
func (o *DescribeMarketUsersParams) WithReverse(reverse *bool) *DescribeMarketUsersParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the describe market users params
func (o *DescribeMarketUsersParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithSearchWord adds the searchWord to the describe market users params
func (o *DescribeMarketUsersParams) WithSearchWord(searchWord *string) *DescribeMarketUsersParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the describe market users params
func (o *DescribeMarketUsersParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WithSortKey adds the sortKey to the describe market users params
func (o *DescribeMarketUsersParams) WithSortKey(sortKey *string) *DescribeMarketUsersParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the describe market users params
func (o *DescribeMarketUsersParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WithUserID adds the userID to the describe market users params
func (o *DescribeMarketUsersParams) WithUserID(userID []string) *DescribeMarketUsersParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the describe market users params
func (o *DescribeMarketUsersParams) SetUserID(userID []string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeMarketUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	valuesMarketID := o.MarketID

	joinedMarketID := swag.JoinByFormat(valuesMarketID, "multi")
	// query array param market_id
	if err := r.SetQueryParam("market_id", joinedMarketID...); err != nil {
		return err
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	valuesOwner := o.Owner

	joinedOwner := swag.JoinByFormat(valuesOwner, "multi")
	// query array param owner
	if err := r.SetQueryParam("owner", joinedOwner...); err != nil {
		return err
	}

	if o.Reverse != nil {

		// query param reverse
		var qrReverse bool
		if o.Reverse != nil {
			qrReverse = *o.Reverse
		}
		qReverse := swag.FormatBool(qrReverse)
		if qReverse != "" {
			if err := r.SetQueryParam("reverse", qReverse); err != nil {
				return err
			}
		}

	}

	if o.SearchWord != nil {

		// query param search_word
		var qrSearchWord string
		if o.SearchWord != nil {
			qrSearchWord = *o.SearchWord
		}
		qSearchWord := qrSearchWord
		if qSearchWord != "" {
			if err := r.SetQueryParam("search_word", qSearchWord); err != nil {
				return err
			}
		}

	}

	if o.SortKey != nil {

		// query param sort_key
		var qrSortKey string
		if o.SortKey != nil {
			qrSortKey = *o.SortKey
		}
		qSortKey := qrSortKey
		if qSortKey != "" {
			if err := r.SetQueryParam("sort_key", qSortKey); err != nil {
				return err
			}
		}

	}

	valuesUserID := o.UserID

	joinedUserID := swag.JoinByFormat(valuesUserID, "multi")
	// query array param user_id
	if err := r.SetQueryParam("user_id", joinedUserID...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
