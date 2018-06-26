
package gowebapi

type EnumerationValue struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Value int32 `json:"Value,omitempty"`

	Path string `json:"Path,omitempty"`

	Parent string `json:"Parent,omitempty"`

	Links *EnumerationValueLinks `json:"Links,omitempty"`

	SerializeWebId bool `json:"SerializeWebId,omitempty"`

	SerializeId bool `json:"SerializeId,omitempty"`

	SerializeDescription bool `json:"SerializeDescription,omitempty"`

	SerializePath bool `json:"SerializePath,omitempty"`

	SerializeLinks bool `json:"SerializeLinks,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
