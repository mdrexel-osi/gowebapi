
package gowebapi

type ItemsAnalysisRulePlugIn struct {
	Items []AnalysisRulePlugIn `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
