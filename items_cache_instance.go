
package gowebapi

type ItemsCacheInstance struct {
	Items []CacheInstance `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
