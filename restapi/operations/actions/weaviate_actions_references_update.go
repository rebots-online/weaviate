/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */ // Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateActionsReferencesUpdateHandlerFunc turns a function with the right signature into a weaviate actions references update handler
type WeaviateActionsReferencesUpdateHandlerFunc func(WeaviateActionsReferencesUpdateParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateActionsReferencesUpdateHandlerFunc) Handle(params WeaviateActionsReferencesUpdateParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// WeaviateActionsReferencesUpdateHandler interface for that can handle valid weaviate actions references update params
type WeaviateActionsReferencesUpdateHandler interface {
	Handle(WeaviateActionsReferencesUpdateParams, *models.Principal) middleware.Responder
}

// NewWeaviateActionsReferencesUpdate creates a new http.Handler for the weaviate actions references update operation
func NewWeaviateActionsReferencesUpdate(ctx *middleware.Context, handler WeaviateActionsReferencesUpdateHandler) *WeaviateActionsReferencesUpdate {
	return &WeaviateActionsReferencesUpdate{Context: ctx, Handler: handler}
}

/*WeaviateActionsReferencesUpdate swagger:route PUT /actions/{id}/references/{propertyName} actions weaviateActionsReferencesUpdate

Replace all references to a class-property.

Replace all references to a class-property.

*/
type WeaviateActionsReferencesUpdate struct {
	Context *middleware.Context
	Handler WeaviateActionsReferencesUpdateHandler
}

func (o *WeaviateActionsReferencesUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateActionsReferencesUpdateParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
