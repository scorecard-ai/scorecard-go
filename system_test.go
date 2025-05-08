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

func TestSystemNew(t *testing.T) {
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
	_, err := client.Systems.New(
		context.TODO(),
		"314",
		scorecard.SystemNewParams{
			ConfigSchema: map[string]any{
				"type":       "bar",
				"properties": "bar",
				"required":   "bar",
			},
			Description: "Production chatbot powered by GPT-4",
			InputSchema: map[string]any{
				"type":       "bar",
				"properties": "bar",
				"required":   "bar",
			},
			Name: "GPT-4 Chatbot",
			OutputSchema: map[string]any{
				"type":       "bar",
				"properties": "bar",
				"required":   "bar",
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

func TestSystemUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Systems.Update(
		context.TODO(),
		"12345678-0a8b-4f66-b6f3-2ddcfa097257",
		scorecard.SystemUpdateParams{
			ConfigSchema: map[string]any{
				"foo": "bar",
			},
			Description: scorecard.String("Updated production chatbot powered by GPT-4 Turbo"),
			InputSchema: map[string]any{
				"foo": "bar",
			},
			Name: scorecard.String("GPT-4 Turbo Chatbot"),
			OutputSchema: map[string]any{
				"foo": "bar",
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

func TestSystemListWithOptionalParams(t *testing.T) {
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
	_, err := client.Systems.List(
		context.TODO(),
		"314",
		scorecard.SystemListParams{
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

func TestSystemDelete(t *testing.T) {
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
	_, err := client.Systems.Delete(context.TODO(), "12345678-0a8b-4f66-b6f3-2ddcfa097257")
	if err != nil {
		var apierr *scorecard.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSystemGet(t *testing.T) {
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
	_, err := client.Systems.Get(context.TODO(), "12345678-0a8b-4f66-b6f3-2ddcfa097257")
	if err != nil {
		var apierr *scorecard.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
