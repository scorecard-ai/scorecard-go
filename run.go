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

// RunService contains methods and other services that help with interacting with
// the Scorecard API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRunService] method instead.
type RunService struct {
	Options []option.RequestOption
}

// NewRunService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRunService(opts ...option.RequestOption) (r RunService) {
	r = RunService{}
	r.Options = opts
	return
}

// Create a new Run.
func (r *RunService) New(ctx context.Context, projectID string, body RunNewParams, opts ...option.RequestOption) (res *Run, err error) {
	opts = append(r.Options[:], opts...)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return
	}
	path := fmt.Sprintf("projects/%s/runs", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update the status of a Run.
func (r *RunService) Update(ctx context.Context, runID string, body RunUpdateParams, opts ...option.RequestOption) (res *RunUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if runID == "" {
		err = errors.New("missing required runId parameter")
		return
	}
	path := fmt.Sprintf("runs/%s", runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// A Run in the Scorecard system.
type Run struct {
	// The ID of the Run.
	ID string `json:"id,required"`
	// The IDs of the metrics this Run is using.
	MetricIDs []string `json:"metricIds,required"`
	// The status of the Run.
	//
	// Any of "pending", "awaiting_execution", "running_execution", "awaiting_scoring",
	// "running_scoring", "awaiting_human_scoring", "completed".
	Status RunStatus `json:"status,required"`
	// The ID of the Testset this Run is testing.
	TestsetID string `json:"testsetId,required"`
	// The ID of the system configuration this Run is using.
	SystemConfigID string `json:"systemConfigId" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		MetricIDs      respjson.Field
		Status         respjson.Field
		TestsetID      respjson.Field
		SystemConfigID respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Run) RawJSON() string { return r.JSON.raw }
func (r *Run) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the Run.
type RunStatus string

const (
	RunStatusPending              RunStatus = "pending"
	RunStatusAwaitingExecution    RunStatus = "awaiting_execution"
	RunStatusRunningExecution     RunStatus = "running_execution"
	RunStatusAwaitingScoring      RunStatus = "awaiting_scoring"
	RunStatusRunningScoring       RunStatus = "running_scoring"
	RunStatusAwaitingHumanScoring RunStatus = "awaiting_human_scoring"
	RunStatusCompleted            RunStatus = "completed"
)

type RunUpdateResponse struct {
	// The ID of the Run.
	ID string `json:"id,required"`
	// The status of the Run.
	//
	// Any of "pending", "awaiting_execution", "running_execution", "awaiting_scoring",
	// "running_scoring", "awaiting_human_scoring", "completed".
	Status RunUpdateResponseStatus `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RunUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *RunUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the Run.
type RunUpdateResponseStatus string

const (
	RunUpdateResponseStatusPending              RunUpdateResponseStatus = "pending"
	RunUpdateResponseStatusAwaitingExecution    RunUpdateResponseStatus = "awaiting_execution"
	RunUpdateResponseStatusRunningExecution     RunUpdateResponseStatus = "running_execution"
	RunUpdateResponseStatusAwaitingScoring      RunUpdateResponseStatus = "awaiting_scoring"
	RunUpdateResponseStatusRunningScoring       RunUpdateResponseStatus = "running_scoring"
	RunUpdateResponseStatusAwaitingHumanScoring RunUpdateResponseStatus = "awaiting_human_scoring"
	RunUpdateResponseStatusCompleted            RunUpdateResponseStatus = "completed"
)

type RunNewParams struct {
	// The IDs of the metrics this Run is using.
	MetricIDs []string `json:"metricIds,omitzero,required"`
	// The ID of the Testset this Run is testing.
	TestsetID string `json:"testsetId,required"`
	// The ID of the system configuration this Run is using.
	SystemConfigID param.Opt[string] `json:"systemConfigId,omitzero" format:"uuid"`
	paramObj
}

func (r RunNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RunNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RunNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RunUpdateParams struct {
	// The status of the Run.
	//
	// Any of "pending", "awaiting_execution", "running_execution", "awaiting_scoring",
	// "running_scoring", "awaiting_human_scoring", "completed".
	Status RunUpdateParamsStatus `json:"status,omitzero,required"`
	paramObj
}

func (r RunUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow RunUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RunUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the Run.
type RunUpdateParamsStatus string

const (
	RunUpdateParamsStatusPending              RunUpdateParamsStatus = "pending"
	RunUpdateParamsStatusAwaitingExecution    RunUpdateParamsStatus = "awaiting_execution"
	RunUpdateParamsStatusRunningExecution     RunUpdateParamsStatus = "running_execution"
	RunUpdateParamsStatusAwaitingScoring      RunUpdateParamsStatus = "awaiting_scoring"
	RunUpdateParamsStatusRunningScoring       RunUpdateParamsStatus = "running_scoring"
	RunUpdateParamsStatusAwaitingHumanScoring RunUpdateParamsStatus = "awaiting_human_scoring"
	RunUpdateParamsStatusCompleted            RunUpdateParamsStatus = "completed"
)
