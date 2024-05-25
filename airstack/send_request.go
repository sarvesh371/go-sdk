package airstack

import (
	"context"

	"github.com/machinebox/graphql"
)

type GraphQLRequest struct {
	client *graphql.Client
	apiKey string
}

func AirstackClient(apiKey string) (*GraphQLRequest, error) {
	if apiKey == "" {
		panic("API key cannot be empty")
	}

	client := graphql.NewClient(API_ENDPOINT_PROD)

	return &GraphQLRequest{
		client: client,
		apiKey: apiKey,
	}, nil
}

func (g *GraphQLRequest) executeQuery(query string, variables map[string]interface{}) (interface{}, error) {
	graphqlRequest := graphql.NewRequest(query)
	graphqlRequest.Header.Add("Authorization", g.apiKey)
	setVariables(graphqlRequest, variables)

	var graphqlResponse interface{}
	if err := g.client.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		return graphqlResponse, err
	}
	return graphqlResponse, nil

}
