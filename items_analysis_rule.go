
package gowebapi

type ItemsAnalysisRule struct {
	Items []AnalysisRule `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
