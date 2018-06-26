
package gowebapi

type ItemsAttributeTrait struct {
	Items []AttributeTrait `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
