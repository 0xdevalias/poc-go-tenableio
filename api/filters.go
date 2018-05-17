package api

import (
	"fmt"
	"net/url"

	"github.com/go-resty/resty"
)

type CommonFilterHelpers struct {
	dateRange        *int
	filters          []Filter
	filterSearchType *FilterSearchType
}

type FilterSearchType string

const (
	FilterSearchTypeAnd FilterSearchType = "and"
	FilterSearchTypeOr  FilterSearchType = "or"
)

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
