package airstack

import (
	"encoding/json"
	"github.com/machinebox/graphql"
)

func jsonStrToMap(jsonStr string) (map[string]interface{}, error) {
	var resultMap map[string]interface{}

	err := json.Unmarshal([]byte(jsonStr), &resultMap)
	if err != nil {
		return nil, err
	}

	return resultMap, nil
}

func setVariables(request *graphql.Request, variables map[string]interface{}) {
	for key, value := range variables {
		request.Var(key, value)
	}
}
