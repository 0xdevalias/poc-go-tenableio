package api

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-resty/resty"
)

type RequestModifier func(*resty.Request) *resty.Request

type FilterConfigShort RequestModifier

type Filter struct {
	Filter  string `json:"filter"`
	Quality string `json:"quality"`
	Value   string `json:"value"`
}

type FilterShort struct {
	Field     string
	Operation string
	Value     string
}

type SortConfig struct {
	Field string
	Type  FilterSortType
}

type CommonFilterHelpers struct {
	dateRange        *int
	filters          []Filter
	filterSearchType *FilterSearchType
}

type FilterSearchType string

type FilterSortType string

const (
	FilterSearchTypeAnd FilterSearchType = "and"
	FilterSearchTypeOr  FilterSearchType = "or"

	FilterSortTypeAsc  FilterSortType = "asc"
	FilterSortTypeDesc FilterSortType = "desc"
)

func WithFilters(filters ...FilterShort) FilterConfigShort {
	return func(r *resty.Request) *resty.Request {
		params := url.Values{}
		for _, f := range filters {
			params.Add("f", fmt.Sprintf("%s:%s:%s", f.Field, f.Operation, f.Value))
		}

		r.SetMultiValueQueryParams(params)

		return r
	}
}

func WithFilterType(ft FilterSearchType) FilterConfigShort {
	return func(r *resty.Request) *resty.Request {
		r.SetQueryParam("ft", string(ft))
		return r
	}
}

func WithWildcardFilter(filterText string) FilterConfigShort {
	return func(r *resty.Request) *resty.Request {
		r.SetQueryParam("w", filterText)
		return r
	}
}

func WithWildcardFields(fields ...string) FilterConfigShort {
	return func(r *resty.Request) *resty.Request {
		r.SetQueryParam("wf", strings.Join(fields, ","))
		return r
	}
}

func WithSort(sorts ...SortConfig) FilterConfigShort {
	return func(r *resty.Request) *resty.Request {
		var ss []string
		for _, sort := range sorts {
			ss = append(ss, fmt.Sprintf("%s:%s", sort.Field, sort.Type))
		}

		r.SetQueryParam("sort", strings.Join(ss, ","))

		return r
	}
}

func WithLimit(limit int) FilterConfigShort {
	return func(r *resty.Request) *resty.Request {
		r.SetQueryParam("limit", strconv.FormatInt(int64(limit), 10))
		return r
	}
}

func WithOffset(offset int) FilterConfigShort {
	return func(r *resty.Request) *resty.Request {
		r.SetQueryParam("offset", strconv.FormatInt(int64(offset), 10))
		return r
	}
}

func (c CommonFilterHelpers) WithDateRange(numDaysIncluded int) CommonFilterHelpers {
	return CommonFilterHelpers{
		dateRange:        &numDaysIncluded,
		filters:          c.filters,
		filterSearchType: c.filterSearchType,
	}
}

func (c CommonFilterHelpers) WithFilters(filters ...Filter) CommonFilterHelpers {
	return CommonFilterHelpers{
		dateRange:        c.dateRange,
		filters:          filters,
		filterSearchType: c.filterSearchType,
	}
}

func (c CommonFilterHelpers) WithSearchType(st FilterSearchType) CommonFilterHelpers {
	return CommonFilterHelpers{
		dateRange:        c.dateRange,
		filters:          c.filters,
		filterSearchType: &st,
	}
}

func (c *WorkbenchesAPI) ApplyCommonFilters(req *resty.Request) {
	if c.dateRange != nil && *c.dateRange >= 0 {
		req.SetQueryParam("date_range", fmt.Sprintf("%v", *c.dateRange))
	}

	if len(c.filters) > 0 {
		req.SetMultiValueQueryParams(FiltersToValues(c.filters))
	}

	if c.filterSearchType != nil {
		req.SetQueryParam("filter.search_type", fmt.Sprintf("%s", *c.filterSearchType))
	}
}

func FiltersToValues(filters []Filter) url.Values {
	values := url.Values{}

	makeKey := func(num int, key string) string { return fmt.Sprintf("filter.%d.%s", num, key) }

	for i, f := range filters {
		values.Add(makeKey(i, "filter"), f.Filter)
		values.Add(makeKey(i, "quality"), f.Quality)
		values.Add(makeKey(i, "value"), f.Value)
	}

	return values
}
