
package gowebapi

type ItemsItemPoint struct {
	Items []ItemPoint `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
