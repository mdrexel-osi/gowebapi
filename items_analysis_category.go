
package gowebapi

type ItemsAnalysisCategory struct {
	Items []AnalysisCategory `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
