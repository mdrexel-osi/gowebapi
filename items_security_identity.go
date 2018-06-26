
package gowebapi

type ItemsSecurityIdentity struct {
	Items []SecurityIdentity `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
