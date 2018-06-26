
package gowebapi

type AttributeTemplate struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Path string `json:"Path,omitempty"`

	Type_ string `json:"Type,omitempty"`

	TypeQualifier string `json:"TypeQualifier,omitempty"`

	DefaultUnitsName string `json:"DefaultUnitsName,omitempty"`

	DefaultValue *interface{} `json:"DefaultValue,omitempty"`

	DataReferencePlugIn string `json:"DataReferencePlugIn,omitempty"`

	ConfigString string `json:"ConfigString,omitempty"`

	IsConfigurationItem bool `json:"IsConfigurationItem,omitempty"`

	IsExcluded bool `json:"IsExcluded,omitempty"`

	IsHidden bool `json:"IsHidden,omitempty"`

	IsManualDataEntry bool `json:"IsManualDataEntry,omitempty"`

	HasChildren bool `json:"HasChildren,omitempty"`

	CategoryNames []string `json:"CategoryNames,omitempty"`

	TraitName string `json:"TraitName,omitempty"`

	DefaultUnitsNameAbbreviation string `json:"DefaultUnitsNameAbbreviation,omitempty"`

	Links *AttributeTemplateLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
