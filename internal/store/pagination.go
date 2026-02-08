package store

import (
	"net/http"
	"strconv"
)

type PaginationQuery struct {
	Limit  int    `json:"limit" validate:"gte=1,lte=20"`
	Offset int    `json:"offset" validate:"gte=0"`
	Sort   string `json:"sort" validate:"oneof=asc desc"`
}

func (fq PaginationQuery) Parse(r *http.Request) (PaginationQuery, error) {
	qs := r.URL.Query()

	limitQuery := qs.Get("limit")
	if limitQuery != "" {
		limit, err := strconv.Atoi(limitQuery)
		if err != nil {
			return fq, err
		}

		fq.Limit = limit
	}

	OffsetQuery := qs.Get("offset")
	if OffsetQuery != "" {
		offset, err := strconv.Atoi(OffsetQuery)
		if err != nil {
			return fq, err
		}

		fq.Offset = offset
	}

	SortQuery := qs.Get("sort")
	if SortQuery != "" {
		fq.Sort = SortQuery
	}

	return fq, nil
}
