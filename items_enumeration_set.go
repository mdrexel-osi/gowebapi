
package gowebapi

type ItemsEnumerationSet struct {
	Items []EnumerationSet `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
