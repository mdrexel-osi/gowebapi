
package gowebapi

type ItemsEnumerationValue struct {
	Items []EnumerationValue `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
