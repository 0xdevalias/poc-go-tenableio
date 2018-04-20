package api

import (
	"fmt"
	"net/url"

	"github.com/go-resty/resty"
)

type commonFilterHelpers struct {
	dateRange        *int
	filters          []Filter
	filterSearchType *FilterSearchType
}

type FilterSearchType string

const (
	FilterSearchTypeAnd FilterSearchType = "and"
	FilterSearchTypeOr  FilterSearchType = "or"
)

func (c commonFilterHelpers) WithDateRange(numDaysIncluded int) commonFilterHelpers {
	return commonFilterHelpers{
		dateRange:        &numDaysIncluded,
		filters:          c.filters,
		filterSearchType: c.filterSearchType,
	}
}

func (c commonFilterHelpers) WithFilters(filters ...Filter) commonFilterHelpers {
	return commonFilterHelpers{
		dateRange:        c.dateRange,
		filters:          filters,
		filterSearchType: c.filterSearchType,
	}
}

func (c commonFilterHelpers) WithSearchType(st FilterSearchType) commonFilterHelpers {
	return commonFilterHelpers{
		dateRange:        c.dateRange,
		filters:          c.filters,
		filterSearchType: &st,
	}
}

func (c *WorkbenchesAPI) applyCommonFilters(req *resty.Request) {
	if c.dateRange != nil && *c.dateRange >= 0 {
		req.SetQueryParam("date_range", fmt.Sprintf("%v", *c.dateRange))
	}

	if len(c.filters) > 0 {
		req.SetMultiValueQueryParams(filtersToValues(c.filters))
	}

	if c.filterSearchType != nil {
		req.SetQueryParam("filter.search_type", fmt.Sprintf("%s", *c.filterSearchType))
	}
}

func filtersToValues(filters []Filter) url.Values {
	values := url.Values{}

	makeKey := func(num int, key string) string { return fmt.Sprintf("filter.%d.%s", num, key) }

	for i, f := range filters {
		values.Add(makeKey(i, "filter"), f.Filter)
		values.Add(makeKey(i, "quality"), f.Quality)
		values.Add(makeKey(i, "value"), f.Value)
	}

	return values
}
