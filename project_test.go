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

func TestProjectNew(t *testing.T) {
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
	_, err := client.Projects.New(context.TODO(), scorecard.ProjectNewParams{
		Description: "This is a test project",
		Name:        "My Project",
	})
	if err != nil {
		var apierr *scorecard.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProjectListWithOptionalParams(t *testing.T) {
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
	_, err := client.Projects.List(context.TODO(), scorecard.ProjectListParams{
		Cursor: scorecard.String("123"),
		Limit:  scorecard.Int(20),
	})
	if err != nil {
		var apierr *scorecard.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
