
package gowebapi

type ItemsSummaryValue struct {
	Items []SummaryValue `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
