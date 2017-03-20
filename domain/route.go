package domain

import "net/http"

// RouteHandlerVersion type
type RouteHandlerVersion string

// Handler func
type HandlerFunc func(http.ResponseWriter, *http.Request)

// RouteHandlers is a map of route version to its handler
type RouteHandlers map[RouteHandlerVersion]http.HandlerFunc

type Route struct {
	Name           string
	Method         string
	Pattern        string
	DefaultVersion RouteHandlerVersion
	RouteHandlers  RouteHandlers
	Handler        HandlerFunc
}

// the routes type
type Routes []Route

func (r *Routes) Append(routes ...*Routes) Routes {
	result := Routes{}

	for _, route := range *r {
		result = append(result, route)
	}

	for _, _routes := range routes {
		for _, route := range *_routes {
			result = append(result, route)
		}
	}

	return result
}
