
package gowebapi

type ItemsSecurityEntry struct {
	Items []SecurityEntry `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
