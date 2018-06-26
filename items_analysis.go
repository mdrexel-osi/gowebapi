
package gowebapi

type ItemsAnalysis struct {
	Items []Analysis `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
