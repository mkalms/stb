/*
 * Device admin API
 *
 * This is the API that is used to manage a device
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
)



// DefaultApiRouter defines the required methods for binding the api requests to a responses for the DefaultApi
// The DefaultApiRouter implementation should parse necessary information from the http request,
// pass the data to a DefaultApiServicer to perform the required actions, then write the service results to the http response.
type DefaultApiRouter interface { 
	GetDeviceConfig(http.ResponseWriter, *http.Request)
	Health(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
	SetDeviceConfig(http.ResponseWriter, *http.Request)
}


// DefaultApiServicer defines the api actions for the DefaultApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DefaultApiServicer interface { 
	GetDeviceConfig(context.Context) (ImplResponse, error)
	Health(context.Context) (ImplResponse, error)
	Login(context.Context, LoginRequest) (ImplResponse, error)
	SetDeviceConfig(context.Context, SetDeviceConfigRequest) (ImplResponse, error)
}
