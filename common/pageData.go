package common

type PageData[T any] struct {
	Data     T
	PageInfo PageInfo
}

type PageInfo struct {
	Page          int
	pageCount     int
	totalPages    int
	totalContents int
	prevPage      int
	nextPage      int
}

func NewPageData[T any](d T, p PageInfo) PageData[any] {
	return PageData[any]{
		Data:     d,
		PageInfo: p,
	}
}
