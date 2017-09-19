// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteV1ClustersClusterIDHandlerFunc turns a function with the right signature into a delete v1 clusters cluster ID handler
type DeleteV1ClustersClusterIDHandlerFunc func(DeleteV1ClustersClusterIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteV1ClustersClusterIDHandlerFunc) Handle(params DeleteV1ClustersClusterIDParams) middleware.Responder {
	return fn(params)
}

// DeleteV1ClustersClusterIDHandler interface for that can handle valid delete v1 clusters cluster ID params
type DeleteV1ClustersClusterIDHandler interface {
	Handle(DeleteV1ClustersClusterIDParams) middleware.Responder
}

// NewDeleteV1ClustersClusterID creates a new http.Handler for the delete v1 clusters cluster ID operation
func NewDeleteV1ClustersClusterID(ctx *middleware.Context, handler DeleteV1ClustersClusterIDHandler) *DeleteV1ClustersClusterID {
	return &DeleteV1ClustersClusterID{Context: ctx, Handler: handler}
}

/*DeleteV1ClustersClusterID swagger:route DELETE /v1/clusters/{clusterId} clusters deleteV1ClustersClusterId

Deletes a cluster

Delete a single cluster identified via its id

*/
type DeleteV1ClustersClusterID struct {
	Context *middleware.Context
	Handler DeleteV1ClustersClusterIDHandler
}

func (o *DeleteV1ClustersClusterID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteV1ClustersClusterIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}