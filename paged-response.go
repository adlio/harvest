package harvest

type Pageable interface {
	HasNextPage() bool
}

type PagedResponse struct {
	PerPage      int64  `json:"per_page"`
	TotalPages   int64  `json:"total_pages"`
	TotalEntries int64  `json:"total_entries"`
	NextPage     *int64 `json:"next_page"`
	PreviousPage *int64 `json:"previous_page"`
	Page         int64  `json:"page"`
}

func (r *PagedResponse) HasNextPage() bool {
	return r.NextPage != nil && r.Page < r.TotalPages
}
