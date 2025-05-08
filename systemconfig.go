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

// SystemConfigService contains methods and other services that help with
// interacting with the Scorecard API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSystemConfigService] method instead.
type SystemConfigService struct {
	Options []option.RequestOption
}

// NewSystemConfigService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSystemConfigService(opts ...option.RequestOption) (r SystemConfigService) {
	r = SystemConfigService{}
	r.Options = opts
	return
}

// Create a new configuration for a system.
//
// Each configuration contains specific parameter values that match the system's
// configSchema - things like model parameters, thresholds, or processing options.
// Once created, configurations cannot be modified, ensuring stable reference
// points for evaluations.
//
// When creating a configuration:
//
//   - The 'config' object is validated against the parent system's configSchema
//   - Configurations with validation errors are still stored, with errors included
//     in the response
//   - Validation errors indicate fields that don't match the schema but don't
//     prevent creation
//   - Having validation errors may affect how some evaluation metrics are calculated
func (r *SystemConfigService) New(ctx context.Context, systemID string, body SystemConfigNewParams, opts ...option.RequestOption) (res *SystemConfig, err error) {
	opts = append(r.Options[:], opts...)
	if systemID == "" {
		err = errors.New("missing required systemId parameter")
		return
	}
	path := fmt.Sprintf("systems/%s/configs", systemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a paginated list of configurations for a specific system.
//
// System configurations provide concrete parameter values for a System Under Test,
// defining exactly how the system should be configured during an evaluation run.
func (r *SystemConfigService) List(ctx context.Context, systemID string, query SystemConfigListParams, opts ...option.RequestOption) (res *pagination.PaginatedResponse[SystemConfig], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if systemID == "" {
		err = errors.New("missing required systemId parameter")
		return
	}
	path := fmt.Sprintf("systems/%s/configs", systemID)
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

// Retrieve a paginated list of configurations for a specific system.
//
// System configurations provide concrete parameter values for a System Under Test,
// defining exactly how the system should be configured during an evaluation run.
func (r *SystemConfigService) ListAutoPaging(ctx context.Context, systemID string, query SystemConfigListParams, opts ...option.RequestOption) *pagination.PaginatedResponseAutoPager[SystemConfig] {
	return pagination.NewPaginatedResponseAutoPager(r.List(ctx, systemID, query, opts...))
}

// Retrieve a specific system configuration by ID.
func (r *SystemConfigService) Get(ctx context.Context, systemConfigID string, query SystemConfigGetParams, opts ...option.RequestOption) (res *SystemConfig, err error) {
	opts = append(r.Options[:], opts...)
	if query.SystemID == "" {
		err = errors.New("missing required systemId parameter")
		return
	}
	if systemConfigID == "" {
		err = errors.New("missing required systemConfigId parameter")
		return
	}
	path := fmt.Sprintf("systems/%s/configs/%s", query.SystemID, systemConfigID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// A SystemConfig defines the specific settings for a System Under Test.
//
// Configurations contain parameter values that determine system behavior during
// evaluation. They are immutable snapshots - once created, they never change.
//
// When running evaluations, you reference a specific configId to establish which
// configuration to test.
//
// Configurations will be validated against the system's configSchema, with
// non-conforming values generating warnings.
type SystemConfig struct {
	// The ID of the system configuration.
	ID string `json:"id,required" format:"uuid"`
	// The configuration of the system.
	Config map[string]any `json:"config,required"`
	// The name of the system configuration.
	Name string `json:"name,required"`
	// The ID of the system the configuration belongs to.
	SystemID string `json:"systemId,required" format:"uuid"`
	// Validation errors found in the configuration. If present, the configuration
	// doesn't fully conform to its system's configSchema.
	ValidationErrors []SystemConfigValidationError `json:"validationErrors"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Config           respjson.Field
		Name             respjson.Field
		SystemID         respjson.Field
		ValidationErrors respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SystemConfig) RawJSON() string { return r.JSON.raw }
func (r *SystemConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SystemConfigValidationError struct {
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
func (r SystemConfigValidationError) RawJSON() string { return r.JSON.raw }
func (r *SystemConfigValidationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SystemConfigNewParams struct {
	// The configuration of the system.
	Config map[string]any `json:"config,omitzero,required"`
	// The name of the system configuration.
	Name string `json:"name,required"`
	// Validation errors found in the configuration. If present, the configuration
	// doesn't fully conform to its system's configSchema.
	ValidationErrors []SystemConfigNewParamsValidationError `json:"validationErrors,omitzero"`
	paramObj
}

func (r SystemConfigNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SystemConfigNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SystemConfigNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Message, Path are required.
type SystemConfigNewParamsValidationError struct {
	// Human-readable error description.
	Message string `json:"message,required"`
	// JSON Pointer to the field with the validation error.
	Path string `json:"path,required"`
	paramObj
}

func (r SystemConfigNewParamsValidationError) MarshalJSON() (data []byte, err error) {
	type shadow SystemConfigNewParamsValidationError
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SystemConfigNewParamsValidationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SystemConfigListParams struct {
	// Cursor for pagination. Pass the `nextCursor` from the previous response to get
	// the next page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return (1-100). Use with `cursor` for pagination
	// through large sets.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SystemConfigListParams]'s query parameters as `url.Values`.
func (r SystemConfigListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SystemConfigGetParams struct {
	SystemID string `path:"systemId,required" format:"uuid" json:"-"`
	paramObj
}
