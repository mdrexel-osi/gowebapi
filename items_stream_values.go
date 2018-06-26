
package gowebapi

type ItemsStreamValues struct {
	Items []StreamValues `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
