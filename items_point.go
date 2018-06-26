
package gowebapi

type ItemsPoint struct {
	Items []Point `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
