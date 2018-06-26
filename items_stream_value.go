
package gowebapi

type ItemsStreamValue struct {
	Items []StreamValue `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
