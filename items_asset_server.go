
package gowebapi

type ItemsAssetServer struct {
	Items []AssetServer `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
