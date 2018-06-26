
package gowebapi

type ItemsStreamSummaries struct {
	Items []StreamSummaries `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
