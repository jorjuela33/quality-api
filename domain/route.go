package domain

import "github.com/gin-gonic/gin"

// RouteHandlerVersion type
type RouteHandlerVersion string

// Handler func
type HandlerFunc gin.HandlerFunc

// RouteHandlers is a map of route version to its handler
type RouteHandlers map[RouteHandlerVersion]HandlerFunc

type Route struct {
	Name           string
	Method         string
	Pattern        string
	DefaultVersion RouteHandlerVersion
	RouteHandlers  RouteHandlers
	Handler        gin.HandlerFunc // this should be replaced for this HandlerFunc but not working
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
