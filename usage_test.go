// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package scorecard_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/scorecard-go"
	"github.com/stainless-sdks/scorecard-go/internal/testutil"
	"github.com/stainless-sdks/scorecard-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := scorecard.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	testset, err := client.Testsets.New(
		context.TODO(),
		"314",
		scorecard.TestsetNewParams{
			Description: "Testset for long context Q&A chatbot.",
			FieldMapping: scorecard.TestsetNewParamsFieldMapping{
				Inputs:   []string{"question"},
				Labels:   []string{"idealAnswer"},
				Metadata: []string{"string"},
			},
			JsonSchema: map[string]any{
				"type":       "bar",
				"properties": "bar",
			},
			Name: "Long Context Q&A",
		},
	)
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", testset.ID)
}
