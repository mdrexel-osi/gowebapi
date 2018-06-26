
package gowebapi

type AttributeTrait struct {
	Name string `json:"Name,omitempty"`

	Abbreviation string `json:"Abbreviation,omitempty"`

	AllowChildAttributes bool `json:"AllowChildAttributes,omitempty"`

	AllowDuplicates bool `json:"AllowDuplicates,omitempty"`

	IsAllowedOnRootAttribute bool `json:"IsAllowedOnRootAttribute,omitempty"`

	IsTypeInherited bool `json:"IsTypeInherited,omitempty"`

	IsUOMInherited bool `json:"IsUOMInherited,omitempty"`

	RequireNumeric bool `json:"RequireNumeric,omitempty"`

	RequireString bool `json:"RequireString,omitempty"`

	Links *AttributeTraitLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
