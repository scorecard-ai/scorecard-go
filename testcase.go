// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package scorecard

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/scorecard-go/internal/apijson"
	"github.com/stainless-sdks/scorecard-go/internal/apiquery"
	"github.com/stainless-sdks/scorecard-go/internal/requestconfig"
	"github.com/stainless-sdks/scorecard-go/option"
	"github.com/stainless-sdks/scorecard-go/packages/pagination"
	"github.com/stainless-sdks/scorecard-go/packages/param"
	"github.com/stainless-sdks/scorecard-go/packages/respjson"
)

// TestcaseService contains methods and other services that help with interacting
// with the Scorecard API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTestcaseService] method instead.
type TestcaseService struct {
	Options []option.RequestOption
}

// NewTestcaseService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTestcaseService(opts ...option.RequestOption) (r TestcaseService) {
	r = TestcaseService{}
	r.Options = opts
	return
}

// Create multiple Testcases in the specified Testset.
func (r *TestcaseService) New(ctx context.Context, testsetID string, body TestcaseNewParams, opts ...option.RequestOption) (res *TestcaseNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if testsetID == "" {
		err = errors.New("missing required testsetId parameter")
		return
	}
	path := fmt.Sprintf("testsets/%s/testcases", testsetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Replace the data of an existing Testcase while keeping its ID.
func (r *TestcaseService) Update(ctx context.Context, testcaseID string, body TestcaseUpdateParams, opts ...option.RequestOption) (res *Testcase, err error) {
	opts = append(r.Options[:], opts...)
	if testcaseID == "" {
		err = errors.New("missing required testcaseId parameter")
		return
	}
	path := fmt.Sprintf("testcases/%s", testcaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieve a paginated list of Testcases belonging to a Testset.
func (r *TestcaseService) List(ctx context.Context, testsetID string, query TestcaseListParams, opts ...option.RequestOption) (res *pagination.PaginatedResponse[Testcase], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if testsetID == "" {
		err = errors.New("missing required testsetId parameter")
		return
	}
	path := fmt.Sprintf("testsets/%s/testcases", testsetID)
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

// Retrieve a paginated list of Testcases belonging to a Testset.
func (r *TestcaseService) ListAutoPaging(ctx context.Context, testsetID string, query TestcaseListParams, opts ...option.RequestOption) *pagination.PaginatedResponseAutoPager[Testcase] {
	return pagination.NewPaginatedResponseAutoPager(r.List(ctx, testsetID, query, opts...))
}

// Delete multiple Testcases by their IDs.
func (r *TestcaseService) Delete(ctx context.Context, body TestcaseDeleteParams, opts ...option.RequestOption) (res *TestcaseDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "testcases/bulk-delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a specific Testcase by ID.
func (r *TestcaseService) Get(ctx context.Context, testcaseID string, opts ...option.RequestOption) (res *Testcase, err error) {
	opts = append(r.Options[:], opts...)
	if testcaseID == "" {
		err = errors.New("missing required testcaseId parameter")
		return
	}
	path := fmt.Sprintf("testcases/%s", testcaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// A test case in the Scorecard system. Contains JSON data that is validated
// against the schema defined by its Testset. The `inputs` and `labels` fields are
// derived from the `data` field based on the Testset's `fieldMapping`, and include
// all mapped fields, including those with validation errors. Testcases are stored
// regardless of validation results, with any validation errors included in the
// `validationErrors` field.
type Testcase struct {
	// The ID of the Testcase.
	ID string `json:"id,required"`
	// Derived from data based on the Testset's fieldMapping. Contains all fields
	// marked as inputs, including those with validation errors.
	Inputs map[string]any `json:"inputs,required"`
	// The JSON data of the Testcase, which is validated against the Testset's schema.
	JsonData map[string]any `json:"jsonData,required"`
	// Derived from data based on the Testset's fieldMapping. Contains all fields
	// marked as labels, including those with validation errors.
	Labels map[string]any `json:"labels,required"`
	// The ID of the Testset this Testcase belongs to.
	TestsetID string `json:"testsetId,required"`
	// Validation errors found in the Testcase data. If present, the Testcase doesn't
	// fully conform to its Testset's schema.
	ValidationErrors []TestcaseValidationError `json:"validationErrors"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Inputs           respjson.Field
		JsonData         respjson.Field
		Labels           respjson.Field
		TestsetID        respjson.Field
		ValidationErrors respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Testcase) RawJSON() string { return r.JSON.raw }
func (r *Testcase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TestcaseValidationError struct {
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
func (r TestcaseValidationError) RawJSON() string { return r.JSON.raw }
func (r *TestcaseValidationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TestcaseNewResponse struct {
	Items []Testcase `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TestcaseNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TestcaseNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TestcaseDeleteResponse struct {
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
func (r TestcaseDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *TestcaseDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TestcaseNewParams struct {
	// Testcases to create (max 100).
	Items []TestcaseNewParamsItem `json:"items,omitzero,required"`
	paramObj
}

func (r TestcaseNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TestcaseNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TestcaseNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A test case in the Scorecard system. Contains JSON data that is validated
// against the schema defined by its Testset. The `inputs` and `labels` fields are
// derived from the `data` field based on the Testset's `fieldMapping`, and include
// all mapped fields, including those with validation errors. Testcases are stored
// regardless of validation results, with any validation errors included in the
// `validationErrors` field.
//
// The property JsonData is required.
type TestcaseNewParamsItem struct {
	// The JSON data of the Testcase, which is validated against the Testset's schema.
	JsonData map[string]any `json:"jsonData,omitzero,required"`
	paramObj
}

func (r TestcaseNewParamsItem) MarshalJSON() (data []byte, err error) {
	type shadow TestcaseNewParamsItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TestcaseNewParamsItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TestcaseUpdateParams struct {
	// The JSON data of the Testcase, which is validated against the Testset's schema.
	JsonData map[string]any `json:"jsonData,omitzero,required"`
	paramObj
}

func (r TestcaseUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TestcaseUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TestcaseUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TestcaseListParams struct {
	// Cursor for pagination. Pass the `nextCursor` from the previous response to get
	// the next page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return (1-100). Use with `cursor` for pagination
	// through large sets.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TestcaseListParams]'s query parameters as `url.Values`.
func (r TestcaseListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TestcaseDeleteParams struct {
	// IDs of Testcases to delete.
	IDs []string `json:"ids,omitzero,required"`
	paramObj
}

func (r TestcaseDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow TestcaseDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TestcaseDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
