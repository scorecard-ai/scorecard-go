// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package scorecard

import (
	"context"
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

// ProjectService contains methods and other services that help with interacting
// with the Scorecard API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectService] method instead.
type ProjectService struct {
	Options []option.RequestOption
}

// NewProjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProjectService(opts ...option.RequestOption) (r ProjectService) {
	r = ProjectService{}
	r.Options = opts
	return
}

// Create a new Project.
func (r *ProjectService) New(ctx context.Context, body ProjectNewParams, opts ...option.RequestOption) (res *Project, err error) {
	opts = append(r.Options[:], opts...)
	path := "projects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a paginated list of all Projects. Projects are ordered by creation
// date, with oldest Projects first.
func (r *ProjectService) List(ctx context.Context, query ProjectListParams, opts ...option.RequestOption) (res *pagination.PaginatedResponse[Project], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "projects"
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

// Retrieve a paginated list of all Projects. Projects are ordered by creation
// date, with oldest Projects first.
func (r *ProjectService) ListAutoPaging(ctx context.Context, query ProjectListParams, opts ...option.RequestOption) *pagination.PaginatedResponseAutoPager[Project] {
	return pagination.NewPaginatedResponseAutoPager(r.List(ctx, query, opts...))
}

// A Project in the Scorecard system.
type Project struct {
	// The ID of the Project.
	ID string `json:"id,required"`
	// The description of the Project.
	Description string `json:"description,required"`
	// The name of the Project.
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Description respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Project) RawJSON() string { return r.JSON.raw }
func (r *Project) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectNewParams struct {
	// The description of the Project.
	Description string `json:"description,required"`
	// The name of the Project.
	Name string `json:"name,required"`
	paramObj
}

func (r ProjectNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectListParams struct {
	// Cursor for pagination. Pass the `nextCursor` from the previous response to get
	// the next page of results.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return (1-100). Use with `cursor` for pagination
	// through large sets.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProjectListParams]'s query parameters as `url.Values`.
func (r ProjectListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
