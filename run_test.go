// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package scorecard_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/scorecard-ai/scorecard-go"
	"github.com/scorecard-ai/scorecard-go/internal/testutil"
	"github.com/scorecard-ai/scorecard-go/option"
)

func TestRunNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Runs.New(
		context.TODO(),
		"314",
		scorecard.RunNewParams{
			MetricIDs:      []string{"789", "101"},
			TestsetID:      "246",
			SystemConfigID: scorecard.String("87654321-4d3b-4ae4-8c7a-4b6e2a19ccf0"),
		},
	)
	if err != nil {
		var apierr *scorecard.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRunUpdate(t *testing.T) {
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
	_, err := client.Runs.Update(
		context.TODO(),
		"135",
		scorecard.RunUpdateParams{
			Status: scorecard.RunUpdateParamsStatusAwaitingScoring,
		},
	)
	if err != nil {
		var apierr *scorecard.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
