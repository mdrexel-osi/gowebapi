
package gowebapi

type ItemsAnalysisTemplate struct {
	Items []AnalysisTemplate `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
