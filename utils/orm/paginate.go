package orm

import (
	"context"
	"fmt"
	"math"
	"strings"
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

		sort = fmt.Sprintf("%s %s %s,", sort, sortKey, sortDirection)
	}

	return sort
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

		sortDirection := "asc"
		if sort[0] == '-' {
			sortDirection = "desc"
		}

		sortFields, ok := fieldSortMap[sort]
		if !ok {
			continue
		}

		sortQuery = fmt.Sprintf("%s %s,", sortQuery, o.setSortField(sortFields, sortDirection))

		validSorts = append(validSorts, sort)
	}

	return sortQuery, validSorts
}

func (o *orm) Paginate(ctx context.Context, data interface{}, pagingResp *BasePagingResponse, options PaginateOptions) {
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

	pagingResp = &BasePagingResponse{
		Count:       int(count),
		CurrentPage: options.Paging.Page,
		TotalPage:   calculateTotalPages(int(count), options.Paging.Limit),
		Limit:       options.Paging.Limit,
		Sorts:       arrSort,
	}
}
