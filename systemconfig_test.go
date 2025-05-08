// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package scorecard_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/scorecard-go"
	"github.com/stainless-sdks/scorecard-go/internal/testutil"
	"github.com/stainless-sdks/scorecard-go/option"
)

func TestSystemConfigNewWithOptionalParams(t *testing.T) {
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
	_, err := client.SystemConfigs.New(
		context.TODO(),
		"12345678-0a8b-4f66-b6f3-2ddcfa097257",
		scorecard.SystemConfigNewParams{
			Config: map[string]any{
				"temperature": "bar",
				"maxTokens":   "bar",
				"model":       "bar",
			},
			Name: "Production (Low Temperature)",
			ValidationErrors: []scorecard.SystemConfigNewParamsValidationError{{
				Message: "Required field missing",
				Path:    "/data/question",
			}},
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

func TestSystemConfigListWithOptionalParams(t *testing.T) {
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
	_, err := client.SystemConfigs.List(
		context.TODO(),
		"12345678-0a8b-4f66-b6f3-2ddcfa097257",
		scorecard.SystemConfigListParams{
			Cursor: scorecard.String("eyJvZmZzZXQiOjAsInBhZ2VJZCI6ImNvZGUifQ"),
			Limit:  scorecard.Int(20),
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

func TestSystemConfigGet(t *testing.T) {
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
	_, err := client.SystemConfigs.Get(
		context.TODO(),
		"87654321-4d3b-4ae4-8c7a-4b6e2a19ccf0",
		scorecard.SystemConfigGetParams{
			SystemID: "12345678-0a8b-4f66-b6f3-2ddcfa097257",
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
