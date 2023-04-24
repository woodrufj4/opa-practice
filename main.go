package main

import (
	"context"
	_ "embed"
	"log"

	"github.com/open-policy-agent/opa/rego"
)

//go:embed policy.rego
var defaultPolicy string

func main() {

	regoStruct := rego.New(
		rego.Query("data.policy.allow"),
		rego.Module("policy", defaultPolicy),
	)

	query, err := regoStruct.PrepareForEval(context.Background())

	if err != nil {
		log.Fatalf("failed to prepare for evaluation. Error: %s", err.Error())
	}

	input := map[string]interface{}{
		"user": map[string]interface{}{
			"roles": []string{
				"admin",
			},
		},
	}

	resultset, err := query.Eval(context.Background(), rego.EvalInput(input))

	if err != nil {
		log.Fatalf("failed to evaluate query. Error: %s", err.Error())
	}

	response := "allowed"

	if !resultset.Allowed() {
		response = "denied"
	}

	println(response)

}
