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

// RecordService contains methods and other services that help with interacting
// with the Scorecard API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRecordService] method instead.
type RecordService struct {
	Options []option.RequestOption
}

// NewRecordService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRecordService(opts ...option.RequestOption) (r RecordService) {
	r = RecordService{}
	r.Options = opts
	return
}

// Create a new Record in a Run.
func (r *RecordService) New(ctx context.Context, runID string, body RecordNewParams, opts ...option.RequestOption) (res *Record, err error) {
	opts = append(r.Options[:], opts...)
	if runID == "" {
		err = errors.New("missing required runId parameter")
		return
	}
	path := fmt.Sprintf("runs/%s/records", runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// A record of a system execution in the Scorecard system.
type Record struct {
	// The ID of the Record.
	ID string `json:"id,required"`
	// The actual inputs sent to the system, which should match the system's input
	// schema.
	Inputs map[string]any `json:"inputs,required"`
	// The expected outputs for the Testcase.
	Labels map[string]any `json:"labels,required"`
	// The actual outputs from the system.
	Outputs map[string]any `json:"outputs,required"`
	// The ID of the Run containing this Record.
	RunID string `json:"runId,required"`
	// The ID of the Testcase.
	TestcaseID string `json:"testcaseId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Inputs      respjson.Field
		Labels      respjson.Field
		Outputs     respjson.Field
		RunID       respjson.Field
		TestcaseID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Record) RawJSON() string { return r.JSON.raw }
func (r *Record) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecordNewParams struct {
	// The actual inputs sent to the system, which should match the system's input
	// schema.
	Inputs map[string]any `json:"inputs,omitzero,required"`
	// The expected outputs for the Testcase.
	Labels map[string]any `json:"labels,omitzero,required"`
	// The actual outputs from the system.
	Outputs map[string]any `json:"outputs,omitzero,required"`
	// The ID of the Testcase.
	TestcaseID param.Opt[string] `json:"testcaseId,omitzero"`
	paramObj
}

func (r RecordNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RecordNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecordNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
