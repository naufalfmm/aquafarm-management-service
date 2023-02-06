package orm

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type (
	PagingRequest struct {
		Page  int
		Limit int
		Sorts []string
	}

	PaginateOptions struct {
		Paging       PagingRequest
		FieldSortMap map[string][]string
	}

	BasePagingResponse struct {
		Count       int      `json:"count"`
		CurrentPage int      `json:"currentPage"`
		TotalPage   int      `json:"totalPage"`
		Limit       int      `json:"limit"`
		Sorts       []string `json:"sorts"`
	}
)

func getQueryParamWithDefault(ec echo.Context, paramKey string, defaultVal int) int {
	var (
		queryValInt int
	)

	queryVal := ec.QueryParam(paramKey)

	queryValInt, err := strconv.Atoi(queryVal)
	if err != nil {
		queryValInt = defaultVal
	}

	return queryValInt
}

func getSortQueryParam(ec echo.Context, defaultSorts []string) []string {
	queryVals := strings.Split(ec.QueryParam("sorts"), ",")
	if len(queryVals) == 0 {
		return defaultSorts
	}

	for i, queryVal := range queryVals {
		queryVals[i] = strings.TrimSpace(queryVal)
	}

	return queryVals
}

func NewPagingRequest(ec echo.Context, defaultSorts []string) PagingRequest {
	return PagingRequest{
		Page:  getQueryParamWithDefault(ec, "page", 1),
		Limit: getQueryParamWithDefault(ec, "limit", 100),
		Sorts: getSortQueryParam(ec, defaultSorts),
	}
}

func calculateTotalPages(dataTotalCount, limit int) int {
	totalPagesInFloat := float64(dataTotalCount) / float64(limit)
	totalPagesInFloat = math.Ceil(totalPagesInFloat)

	return int(totalPagesInFloat)
}

func (o orm) setSortField(sortFields []string, direction string) string {
	var sort string

	for _, sortField := range sortFields {
		sortKey := sortField
		sortDirection := "asc"
		if sortField[0] == '-' {
			sortKey = sortField[1:]
			sortDirection = "desc"
		}

		if direction == "desc" {
			if sortDirection == "desc" {
				sortDirection = "asc"
			} else {
				sortDirection = "desc"
			}
		}

		sort = fmt.Sprintf("%s,%s %s", sort, sortKey, sortDirection)
	}

	return sort[1:]
}

func (o orm) sort(sorts []string, fieldSortMap map[string][]string) (string, []string) {
	var (
		sortQuery  string
		validSorts []string
	)

	for _, sort := range sorts {
		sort = strings.TrimSpace(sort)
		if len(sort) < 1 {
			continue
		}

		sortKey := sort
		sortDirection := "asc"
		if sort[0] == '-' {
			sortDirection = "desc"
			sortKey = sort[1:]
		}

		sortFields, ok := fieldSortMap[sortKey]
		if !ok {
			continue
		}

		sortQuery = fmt.Sprintf("%s,%s", sortQuery, o.setSortField(sortFields, sortDirection))

		validSorts = append(validSorts, sort)
	}

	return sortQuery[1:], validSorts
}

func (o *orm) Paginate(ctx context.Context, options PaginateOptions, pagingResp *BasePagingResponse, data interface{}) Orm {
	var (
		ormCount, orm Orm = o.clone(), o.newImpl(o.g)
		count             = int64(pagingResp.Count)
	)

	if err := ormCount.
		WithContext(ctx).
		Count(&count).
		Error(); err != nil {
		orm.AddError(err)
	}

	sortQuery, arrSort := o.sort(options.Paging.Sorts, options.FieldSortMap)
	if len(sortQuery) > 0 {
		orm = orm.Order(sortQuery)
	}

	if err := orm.
		WithContext(ctx).
		Offset((options.Paging.Page - 1) * options.Paging.Limit).
		Limit(options.Paging.Limit).
		Find(data).
		Error(); err != nil {
		orm.AddError(err)
	}

	pagingResp.Count = int(count)
	pagingResp.CurrentPage = options.Paging.Page
	pagingResp.TotalPage = calculateTotalPages(int(count), options.Paging.Limit)
	pagingResp.Limit = options.Paging.Limit
	pagingResp.Sorts = arrSort

	return orm
}
