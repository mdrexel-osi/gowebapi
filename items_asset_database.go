
package gowebapi

type ItemsAssetDatabase struct {
	Items []AssetDatabase `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
