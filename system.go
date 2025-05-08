// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package scorecard

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/scorecard-ai/scorecard-go/internal/apijson"
	"github.com/scorecard-ai/scorecard-go/internal/apiquery"
	"github.com/scorecard-ai/scorecard-go/internal/requestconfig"
	"github.com/scorecard-ai/scorecard-go/option"
	"github.com/scorecard-ai/scorecard-go/packages/pagination"
	"github.com/scorecard-ai/scorecard-go/packages/param"
	"github.com/scorecard-ai/scorecard-go/packages/respjson"
)

// SystemService contains methods and other services that help with interacting
// with the Scorecard API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSystemService] method instead.
type SystemService struct {
	Options []option.RequestOption
}

// NewSystemService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSystemService(opts ...option.RequestOption) (r SystemService) {
	r = SystemService{}
	r.Options = opts
	return
}

// Create a new system definition that specifies the interface contracts for a
// component you want to evaluate.
//
// A system acts as a template that defines three key contracts through JSON
// Schemas:
//
//  1. Input Schema: What data your system accepts (e.g., user queries, context
//     documents)
//  2. Output Schema: What data your system produces (e.g., responses, confidence
//     scores)
//  3. Config Schema: What parameters can be adjusted (e.g., model selection,
//     temperature)
//
// This separation lets you evaluate any system as a black box, focusing on its
// interface rather than implementation details.
func (r *SystemService) New(ctx context.Context, projectID string, body SystemNewParams, opts ...option.RequestOption) (res *System, err error) {
	opts = append(r.Options[:], opts...)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return
	}
	path := fmt.Sprintf("projects/%s/systems", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update an existing system definition. Only the fields provided in the request
// body will be updated. If a field is provided, the new content will replace the
// existing content. If a field is not provided, the existing content will remain
// unchanged.
//
// When updating schemas:
//
//   - The system will accept your changes regardless of compatibility with existing
//     configurations
//   - Schema updates won't invalidate existing evaluations or configurations
//   - For significant redesigns, creating a new system definition provides a cleaner
//     separation
func (r *SystemService) Update(ctx context.Context, systemID string, body SystemUpdateParams, opts ...option.RequestOption) (res *System, err error) {
	opts = append(r.Options[:], opts...)
	if systemID == "" {
		err = errors.New("missing required systemId parameter")
		return
	}
	path := fmt.Sprintf("systems/%s", systemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieve a paginated list of all systems. Systems are ordered by creation date.
func (r *SystemService) List(ctx context.Context, projectID string, query SystemListParams, opts ...option.RequestOption) (res *pagination.PaginatedResponse[System], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return
	}
	path := fmt.Sprintf("projects/%s/systems", projectID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve a paginated list of all systems. Systems are ordered by creation date.
func (r *SystemService) ListAutoPaging(ctx context.Context, projectID string, query SystemListParams, opts ...option.RequestOption) *pagination.PaginatedResponseAutoPager[System] {
	return pagination.NewPaginatedResponseAutoPager(r.List(ctx, projectID, query, opts...))
}

// Delete a system definition by ID. This will not delete associated system
// configurations.
func (r *SystemService) Delete(ctx context.Context, systemID string, opts ...option.RequestOption) (res *SystemDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if systemID == "" {
		err = errors.New("missing required systemId parameter")
		return
	}
	path := fmt.Sprintf("systems/%s", systemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Retrieve a specific system by ID.
func (r *SystemService) Get(ctx context.Context, systemID string, opts ...option.RequestOption) (res *System, err error) {
	opts = append(r.Options[:], opts...)
	if systemID == "" {
		err = errors.New("missing required systemId parameter")
		return
	}
	path := fmt.Sprintf("systems/%s", systemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// A System Under Test (SUT) defines the interface to a component or service you
// want to evaluate.
//
// It specifies three contracts through schemas:
//
// - inputSchema: The structure of data the system accepts.
// - outputSchema: The structure of data the system produces.
// - configSchema: The parameters that modify system behavior.
//
// This abstraction lets you evaluate any system as a black box, focusing on its
// interface rather than implementation details. It's particularly useful for
// systems with variable outputs or complex internal state.
//
// Systems are templates - to run evaluations, pair them with a SystemConfig that
// provides specific parameter values.
type System struct {
	// The ID of the system.
	ID string `json:"id,required" format:"uuid"`
	// The schema of the system's configuration.
	ConfigSchema map[string]any `json:"configSchema,required"`
	// The description of the system.
	Description string `json:"description,required"`
	// The schema of the system's inputs.
	InputSchema map[string]any `json:"inputSchema,required"`
	// The name of the system.
	Name string `json:"name,required"`
	// The schema of the system's outputs.
	OutputSchema map[string]any `json:"outputSchema,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		ConfigSchema respjson.Field
		Description  respjson.Field
		InputSchema  respjson.Field
		Name         respjson.Field
		OutputSchema respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r System) RawJSON() string { return r.JSON.raw }
func (r *System) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SystemDeleteResponse struct {
	// Whether the deletion was successful.
	Success bool `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SystemDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *SystemDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SystemNewParams struct {
	// The schema of the system's configuration.
	ConfigSchema map[string]any `json:"configSchema,omitzero,required"`
	// The description of the system.
	Description string `json:"description,required"`
	// The schema of the system's inputs.
	InputSchema map[string]any `json:"inputSchema,omitzero,required"`
	// The name of the system.
	Name string `json:"name,required"`
	// The schema of the system's outputs.
	OutputSchema map[string]any `json:"outputSchema,omitzero,required"`
	paramObj
}

func (r SystemNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SystemNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SystemNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SystemUpdateParams struct {
	// The description of the system.
	Description param.Opt[string] `json:"description,omitzero"`
	// The name of the system.
	Name param.Opt[string] `json:"name,omitzero"`
	// The schema of the system's configuration.
	ConfigSchema map[string]any `json:"configSchema,omitzero"`
	// The schema of the system's inputs.
	InputSchema map[string]any `json:"inputSchema,omitzero"`
	// The schema of the system's outputs.
	OutputSchema map[string]any `json:"outputSchema,omitzero"`
	paramObj
}

func (r SystemUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SystemUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SystemUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SystemListParams struct {
	// Cursor for pagination. Pass the `nextCursor` from the previous response to get
	// the next page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return (1-100). Use with `cursor` for pagination
	// through large sets.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SystemListParams]'s query parameters as `url.Values`.
func (r SystemListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
