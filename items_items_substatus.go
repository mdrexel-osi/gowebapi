
package gowebapi

type ItemsItemsSubstatus struct {
	Items []ItemsSubstatus `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
