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

// TestsetService contains methods and other services that help with interacting
// with the Scorecard API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTestsetService] method instead.
type TestsetService struct {
	Options []option.RequestOption
}

// NewTestsetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTestsetService(opts ...option.RequestOption) (r TestsetService) {
	r = TestsetService{}
	r.Options = opts
	return
}

// Create a new Testset for a Project. The Testset will be created in the Project
// specified in the path.
func (r *TestsetService) New(ctx context.Context, projectID string, body TestsetNewParams, opts ...option.RequestOption) (res *Testset, err error) {
	opts = append(r.Options[:], opts...)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return
	}
	path := fmt.Sprintf("projects/%s/testsets", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a Testset. Only the fields provided in the request body will be updated.
// If a field is provided, the new content will replace the existing content. If a
// field is not provided, the existing content will remain unchanged.
//
// When updating the schema:
//
//   - If field mappings are not provided and existing mappings reference fields that
//     no longer exist, those mappings will be automatically removed
//   - To preserve all existing mappings, ensure all referenced fields remain in the
//     updated schema
//   - For complete control, provide both schema and fieldMapping when updating the
//     schema
func (r *TestsetService) Update(ctx context.Context, testsetID string, body TestsetUpdateParams, opts ...option.RequestOption) (res *Testset, err error) {
	opts = append(r.Options[:], opts...)
	if testsetID == "" {
		err = errors.New("missing required testsetId parameter")
		return
	}
	path := fmt.Sprintf("testsets/%s", testsetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieve a paginated list of Testsets belonging to a Project.
func (r *TestsetService) List(ctx context.Context, projectID string, query TestsetListParams, opts ...option.RequestOption) (res *pagination.PaginatedResponse[Testset], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if projectID == "" {
		err = errors.New("missing required projectId parameter")
		return
	}
	path := fmt.Sprintf("projects/%s/testsets", projectID)
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

// Retrieve a paginated list of Testsets belonging to a Project.
func (r *TestsetService) ListAutoPaging(ctx context.Context, projectID string, query TestsetListParams, opts ...option.RequestOption) *pagination.PaginatedResponseAutoPager[Testset] {
	return pagination.NewPaginatedResponseAutoPager(r.List(ctx, projectID, query, opts...))
}

// Delete Testset
func (r *TestsetService) Delete(ctx context.Context, testsetID string, opts ...option.RequestOption) (res *TestsetDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if testsetID == "" {
		err = errors.New("missing required testsetId parameter")
		return
	}
	path := fmt.Sprintf("testsets/%s", testsetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get Testset by ID
func (r *TestsetService) Get(ctx context.Context, testsetID string, opts ...option.RequestOption) (res *Testset, err error) {
	opts = append(r.Options[:], opts...)
	if testsetID == "" {
		err = errors.New("missing required testsetId parameter")
		return
	}
	path := fmt.Sprintf("testsets/%s", testsetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// A collection of Testcases that share the same schema. Each Testset defines the
// structure of its Testcases through a JSON schema. The `fieldMapping` object maps
// top-level keys of the Testcase schema to their roles (input/label). Fields not
// mentioned in the `fieldMapping` during creation or update are treated as
// metadata.
//
// ## JSON Schema validation constraints supported:
//
//   - **Required fields** - Fields listed in the schema's `required` array must be
//     present in Testcases.
//   - **Type validation** - Values must match the specified type (string, number,
//     boolean, null, integer, object, array).
//   - **Enum validation** - Values must be one of the options specified in the
//     `enum` array.
//   - **Object property validation** - Properties of objects must conform to their
//     defined schemas.
//   - **Array item validation** - Items in arrays must conform to the `items`
//     schema.
//   - **Logical composition** - Values must conform to at least one schema in the
//     `anyOf` array.
//
// Testcases that fail validation will still be stored, but will include
// `validationErrors` detailing the issues. Extra fields in the Testcase data that
// are not in the schema will be stored but are ignored during validation.
type Testset struct {
	// The ID of the Testset.
	ID string `json:"id,required"`
	// The description of the Testset.
	Description string `json:"description,required"`
	// Maps top-level keys of the Testcase schema to their roles (input/label).
	// Unmapped fields are treated as metadata.
	FieldMapping TestsetFieldMapping `json:"fieldMapping,required"`
	// The JSON schema for each Testcase in the Testset.
	JsonSchema map[string]any `json:"jsonSchema,required"`
	// The name of the Testset.
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Description  respjson.Field
		FieldMapping respjson.Field
		JsonSchema   respjson.Field
		Name         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Testset) RawJSON() string { return r.JSON.raw }
func (r *Testset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maps top-level keys of the Testcase schema to their roles (input/label).
// Unmapped fields are treated as metadata.
type TestsetFieldMapping struct {
	// Fields that represent inputs to the AI system.
	Inputs []string `json:"inputs,required"`
	// Fields that represent expected outputs/labels.
	Labels []string `json:"labels,required"`
	// Fields that are not inputs or labels.
	Metadata []string `json:"metadata,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Inputs      respjson.Field
		Labels      respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TestsetFieldMapping) RawJSON() string { return r.JSON.raw }
func (r *TestsetFieldMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TestsetDeleteResponse struct {
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
func (r TestsetDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *TestsetDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TestsetNewParams struct {
	// The description of the Testset.
	Description string `json:"description,required"`
	// Maps top-level keys of the Testcase schema to their roles (input/label).
	// Unmapped fields are treated as metadata.
	FieldMapping TestsetNewParamsFieldMapping `json:"fieldMapping,omitzero,required"`
	// The JSON schema for each Testcase in the Testset.
	JsonSchema map[string]any `json:"jsonSchema,omitzero,required"`
	// The name of the Testset.
	Name string `json:"name,required"`
	paramObj
}

func (r TestsetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TestsetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TestsetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maps top-level keys of the Testcase schema to their roles (input/label).
// Unmapped fields are treated as metadata.
//
// The properties Inputs, Labels, Metadata are required.
type TestsetNewParamsFieldMapping struct {
	// Fields that represent inputs to the AI system.
	Inputs []string `json:"inputs,omitzero,required"`
	// Fields that represent expected outputs/labels.
	Labels []string `json:"labels,omitzero,required"`
	// Fields that are not inputs or labels.
	Metadata []string `json:"metadata,omitzero,required"`
	paramObj
}

func (r TestsetNewParamsFieldMapping) MarshalJSON() (data []byte, err error) {
	type shadow TestsetNewParamsFieldMapping
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TestsetNewParamsFieldMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TestsetUpdateParams struct {
	// The description of the Testset.
	Description param.Opt[string] `json:"description,omitzero"`
	// The name of the Testset.
	Name param.Opt[string] `json:"name,omitzero"`
	// Maps top-level keys of the Testcase schema to their roles (input/label).
	// Unmapped fields are treated as metadata.
	FieldMapping TestsetUpdateParamsFieldMapping `json:"fieldMapping,omitzero"`
	// The JSON schema for each Testcase in the Testset.
	JsonSchema map[string]any `json:"jsonSchema,omitzero"`
	paramObj
}

func (r TestsetUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TestsetUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TestsetUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Maps top-level keys of the Testcase schema to their roles (input/label).
// Unmapped fields are treated as metadata.
//
// The properties Inputs, Labels, Metadata are required.
type TestsetUpdateParamsFieldMapping struct {
	// Fields that represent inputs to the AI system.
	Inputs []string `json:"inputs,omitzero,required"`
	// Fields that represent expected outputs/labels.
	Labels []string `json:"labels,omitzero,required"`
	// Fields that are not inputs or labels.
	Metadata []string `json:"metadata,omitzero,required"`
	paramObj
}

func (r TestsetUpdateParamsFieldMapping) MarshalJSON() (data []byte, err error) {
	type shadow TestsetUpdateParamsFieldMapping
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TestsetUpdateParamsFieldMapping) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TestsetListParams struct {
	// Cursor for pagination. Pass the `nextCursor` from the previous response to get
	// the next page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return (1-100). Use with `cursor` for pagination
	// through large sets.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TestsetListParams]'s query parameters as `url.Values`.
func (r TestsetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
