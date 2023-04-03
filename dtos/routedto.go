package dtos

type RouteDTO struct {
	HttpMethod string
	Endpoint   string
	Handler    interface{}
}
