
package gowebapi

type ItemsElement struct {
	Items []Element `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
