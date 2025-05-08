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

func TestTestsetNew(t *testing.T) {
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
	_, err := client.Testsets.New(
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
		var apierr *scorecard.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTestsetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Testsets.Update(
		context.TODO(),
		"246",
		scorecard.TestsetUpdateParams{
			Description: scorecard.String("Updated description for the Q&A Testset."),
			FieldMapping: scorecard.TestsetUpdateParamsFieldMapping{
				Inputs:   []string{"string"},
				Labels:   []string{"string"},
				Metadata: []string{"string"},
			},
			JsonSchema: map[string]any{
				"foo": "bar",
			},
			Name: scorecard.String("Updated Q&A Testset"),
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

func TestTestsetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Testsets.List(
		context.TODO(),
		"314",
		scorecard.TestsetListParams{
			Cursor: scorecard.String("123"),
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

func TestTestsetDelete(t *testing.T) {
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
	_, err := client.Testsets.Delete(context.TODO(), "246")
	if err != nil {
		var apierr *scorecard.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTestsetGet(t *testing.T) {
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
	_, err := client.Testsets.Get(context.TODO(), "246")
	if err != nil {
		var apierr *scorecard.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
