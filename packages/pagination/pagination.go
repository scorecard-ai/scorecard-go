// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"

	"github.com/scorecard-ai/scorecard-go/internal/apijson"
	"github.com/scorecard-ai/scorecard-go/internal/requestconfig"
	"github.com/scorecard-ai/scorecard-go/option"
	"github.com/scorecard-ai/scorecard-go/packages/param"
	"github.com/scorecard-ai/scorecard-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type PaginatedResponse[T any] struct {
	Data       []T    `json:"data"`
	NextCursor string `json:"nextCursor"`
	HasMore    bool   `json:"hasMore"`
	Total      int64  `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		NextCursor  respjson.Field
		HasMore     respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r PaginatedResponse[T]) RawJSON() string { return r.JSON.raw }
func (r *PaginatedResponse[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *PaginatedResponse[T]) GetNextPage() (res *PaginatedResponse[T], err error) {
	if r.JSON.HasMore.Valid() && r.HasMore == false {
		return nil, nil
	}
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *PaginatedResponse[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &PaginatedResponse[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type PaginatedResponseAutoPager[T any] struct {
	page *PaginatedResponse[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewPaginatedResponseAutoPager[T any](page *PaginatedResponse[T], err error) *PaginatedResponseAutoPager[T] {
	return &PaginatedResponseAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PaginatedResponseAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PaginatedResponseAutoPager[T]) Current() T {
	return r.cur
}

func (r *PaginatedResponseAutoPager[T]) Err() error {
	return r.err
}

func (r *PaginatedResponseAutoPager[T]) Index() int {
	return r.run
}
