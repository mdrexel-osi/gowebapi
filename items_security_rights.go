
package gowebapi

type ItemsSecurityRights struct {
	Items []SecurityRights `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
