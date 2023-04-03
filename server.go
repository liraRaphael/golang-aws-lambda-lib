package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/liraRaphael/golang-aws-lambda-lib/dtos"
)

type Lambda struct {
	routes map[string][]*dtos.RouteDTO
}

func (l *Lambda) Init() *Lambda {
	l.routes = make(map[string][]*dtos.RouteDTO)

	return l
}

func (l *Lambda) AddRoute(route *dtos.RouteDTO) {
	l.routes[route.HttpMethod] = append(l.routes[route.HttpMethod], route)
}

func (l *Lambda) Run() {
	lambda.Start(l.execute)
}

func (l *Lambda) execute(ctx events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return events.APIGatewayProxyResponse{}, nil
}
