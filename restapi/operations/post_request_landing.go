package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostRequestLandingHandlerFunc turns a function with the right signature into a post request landing handler
type PostRequestLandingHandlerFunc func(PostRequestLandingParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostRequestLandingHandlerFunc) Handle(params PostRequestLandingParams) middleware.Responder {
	return fn(params)
}

// PostRequestLandingHandler interface for that can handle valid post request landing params
type PostRequestLandingHandler interface {
	Handle(PostRequestLandingParams) middleware.Responder
}

// NewPostRequestLanding creates a new http.Handler for the post request landing operation
func NewPostRequestLanding(ctx *middleware.Context, handler PostRequestLandingHandler) *PostRequestLanding {
	return &PostRequestLanding{Context: ctx, Handler: handler}
}

/*PostRequestLanding swagger:route POST /request-landing postRequestLanding

Land a ship in the ship bay

*/
type PostRequestLanding struct {
	Context *middleware.Context
	Handler PostRequestLandingHandler
}

func (o *PostRequestLanding) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _, _ := o.Context.RouteInfo(r)
	var Params = NewPostRequestLandingParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
