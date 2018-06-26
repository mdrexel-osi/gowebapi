
package gowebapi

type ItemsTimeRulePlugIn struct {
	Items []TimeRulePlugIn `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
