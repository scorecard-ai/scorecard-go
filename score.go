// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package scorecard

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/scorecard-ai/scorecard-go/internal/apijson"
	"github.com/scorecard-ai/scorecard-go/internal/requestconfig"
	"github.com/scorecard-ai/scorecard-go/option"
	"github.com/scorecard-ai/scorecard-go/packages/param"
	"github.com/scorecard-ai/scorecard-go/packages/respjson"
)

// ScoreService contains methods and other services that help with interacting with
// the Scorecard API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScoreService] method instead.
type ScoreService struct {
	Options []option.RequestOption
}

// NewScoreService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewScoreService(opts ...option.RequestOption) (r ScoreService) {
	r = ScoreService{}
	r.Options = opts
	return
}

// Create or update a Score for a given Record and MetricConfig. If a Score with
// the specified Record ID and MetricConfig ID already exists, it will be updated.
// Otherwise, a new Score will be created. The score provided should conform to the
// schema defined by the MetricConfig; otherwise, validation errors will be
// reported.
func (r *ScoreService) Upsert(ctx context.Context, metricConfigID string, params ScoreUpsertParams, opts ...option.RequestOption) (res *Score, err error) {
	opts = append(r.Options[:], opts...)
	if params.RecordID == "" {
		err = errors.New("missing required recordId parameter")
		return
	}
	if metricConfigID == "" {
		err = errors.New("missing required metricConfigId parameter")
		return
	}
	path := fmt.Sprintf("records/%s/scores/%s", params.RecordID, metricConfigID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// A Score represents the evaluation of a Record against a specific MetricConfig.
// The actual `score` is stored as flexible JSON. While any JSON is accepted, it is
// expected to conform to the output schema defined by the MetricConfig. Any
// discrepancies will be noted in the `validationErrors` field, but the Score will
// still be stored.
type Score struct {
	// The ID of the MetricConfig this Score is for.
	MetricConfigID string `json:"metricConfigId,required" format:"uuid"`
	// The ID of the Record this Score is for.
	RecordID string `json:"recordId,required"`
	// The score of the Record, as arbitrary JSON. This data should ideally conform to
	// the output schema defined by the associated MetricConfig. If it doesn't,
	// validation errors will be captured in the `validationErrors` field.
	Score map[string]any `json:"score,required"`
	// Validation errors found in the Score data. If present, the Score doesn't fully
	// conform to its MetricConfig's schema.
	ValidationErrors []ScoreValidationError `json:"validationErrors"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MetricConfigID   respjson.Field
		RecordID         respjson.Field
		Score            respjson.Field
		ValidationErrors respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Score) RawJSON() string { return r.JSON.raw }
func (r *Score) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScoreValidationError struct {
	// Human-readable error description.
	Message string `json:"message,required"`
	// JSON Pointer to the field with the validation error.
	Path string `json:"path,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoreValidationError) RawJSON() string { return r.JSON.raw }
func (r *ScoreValidationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScoreUpsertParams struct {
	RecordID string `path:"recordId,required" json:"-"`
	// The score of the Record, as arbitrary JSON. This data should ideally conform to
	// the output schema defined by the associated MetricConfig. If it doesn't,
	// validation errors will be captured in the `validationErrors` field.
	Score map[string]any `json:"score,omitzero,required"`
	paramObj
}

func (r ScoreUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow ScoreUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScoreUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
