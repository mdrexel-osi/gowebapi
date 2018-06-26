
package gowebapi

type ItemsChannelInstance struct {
	Items []ChannelInstance `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
