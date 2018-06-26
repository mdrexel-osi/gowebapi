
package gowebapi

type Table struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Path string `json:"Path,omitempty"`

	CategoryNames []string `json:"CategoryNames,omitempty"`

	TimeZone string `json:"TimeZone,omitempty"`

	ConvertToLocalTime bool `json:"ConvertToLocalTime,omitempty"`

	Links *TableLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
