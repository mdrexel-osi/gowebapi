
package gowebapi

type ItemsUnitClass struct {
	Items []UnitClass `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
