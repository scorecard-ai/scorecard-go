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

func TestScoreUpsert(t *testing.T) {
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
	_, err := client.Scores.Upsert(
		context.TODO(),
		"a1b2c3d4-e5f6-7890-1234-567890abcdef",
		scorecard.ScoreUpsertParams{
			RecordID: "777",
			Score: map[string]any{
				"value":     "bar",
				"reasoning": "bar",
			},
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
