package helper

import (
	"net/url"
	"strconv"
)

var (
	SortVar      = "sort"
	SortNameAsc  = "name_asc"
	SortNameDesc = "name_desc"
	SortAsc      = "asc"
	SortDesc     = "desc"

	PageVar     = "page"
	PageSizeVar = "pageSize"
)

func SetPaginationOptions(values *url.Values) (string, int, int) {
	var sortOpt string
	sort := values.Get(SortVar)
	switch sort {
	case SortNameAsc:
		sortOpt = "product_name asc"
	case SortNameDesc:
		sortOpt = "product_name desc"
	case SortAsc:
		sortOpt = "created_at asc"
	case SortDesc:
		sortOpt = "created_at desc"
	default:
		sortOpt = "created_at asc"
	}
	values.Del(SortVar)

	pageNumber, _ := strconv.Atoi(values.Get(PageVar))
	if pageNumber <= 0 {
		pageNumber = 1
	}
	values.Del(PageVar)

	pageSize, _ := strconv.Atoi(values.Get(PageSizeVar))
	if pageSize <= 0 {
		pageSize = 4
	}
	values.Del(PageSizeVar)

	return sortOpt, pageNumber, pageSize
}
