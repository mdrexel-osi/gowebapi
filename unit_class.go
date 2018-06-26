
package gowebapi

type UnitClass struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	CanonicalUnitName string `json:"CanonicalUnitName,omitempty"`

	CanonicalUnitAbbreviation string `json:"CanonicalUnitAbbreviation,omitempty"`

	Path string `json:"Path,omitempty"`

	Links *UnitClassLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
