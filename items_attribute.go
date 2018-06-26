
package gowebapi

type ItemsAttribute struct {
	Items []Attribute `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
