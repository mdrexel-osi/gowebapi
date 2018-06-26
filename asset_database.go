
package gowebapi

type AssetDatabase struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Path string `json:"Path,omitempty"`

	ExtendedProperties map[string]Value `json:"ExtendedProperties,omitempty"`

	Links *AssetDatabaseLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
