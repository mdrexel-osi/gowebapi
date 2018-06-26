
package gowebapi

type ItemsSecurityMapping struct {
	Items []SecurityMapping `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
