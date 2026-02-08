package store

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

type PaginationQuery struct {
	Limit     int      `json:"limit" validate:"gte=1,lte=20"`
	Offset    int      `json:"offset" validate:"gte=0"`
	Sort      string   `json:"sort" validate:"oneof=asc desc"`
	Tags      []string `json:"tags" validate:"max=5"`
	Search    string   `json:"search" validate:"max=100"`
	StartDate string
	EndDate   string
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

	offsetQuery := qs.Get("offset")
	if offsetQuery != "" {
		offset, err := strconv.Atoi(offsetQuery)
		if err != nil {
			return fq, err
		}

		fq.Offset = offset
	}

	sortQuery := qs.Get("sort")
	if sortQuery != "" {
		fq.Sort = sortQuery
	}

	tagsQuery := qs.Get("tags")
	if tagsQuery != "" {
		fq.Tags = strings.Split(tagsQuery, ",")
	}

	searchQuery := qs.Get("search")
	if searchQuery != "" {
		fq.Search = searchQuery
	}

	startDateQuery := qs.Get("startDate")
	if startDateQuery != "" {
		fq.StartDate = parseTime(startDateQuery)
	}

	endDateQuery := qs.Get("endDate")
	if endDateQuery != "" {
		fq.EndDate = parseTime(endDateQuery)
	}

	return fq, nil
}

func parseTime(str string) string {
	t, err := time.Parse(time.DateTime, str)
	if err != nil {
		return ""
	}

	return t.Format(time.DateTime)
}
