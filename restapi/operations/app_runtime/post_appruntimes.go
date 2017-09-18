// Code generated by go-swagger; DO NOT EDIT.

package app_runtime

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostAppruntimesHandlerFunc turns a function with the right signature into a post appruntimes handler
type PostAppruntimesHandlerFunc func(PostAppruntimesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostAppruntimesHandlerFunc) Handle(params PostAppruntimesParams) middleware.Responder {
	return fn(params)
}

// PostAppruntimesHandler interface for that can handle valid post appruntimes params
type PostAppruntimesHandler interface {
	Handle(PostAppruntimesParams) middleware.Responder
}

// NewPostAppruntimes creates a new http.Handler for the post appruntimes operation
func NewPostAppruntimes(ctx *middleware.Context, handler PostAppruntimesHandler) *PostAppruntimes {
	return &PostAppruntimes{Context: ctx, Handler: handler}
}

/*PostAppruntimes swagger:route POST /appruntimes app-runtime postAppruntimes

Creates an app runtime

Adds a new app runtime to the runtimes list.

*/
type PostAppruntimes struct {
	Context *middleware.Context
	Handler PostAppruntimesHandler
}

func (o *PostAppruntimes) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostAppruntimesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
