
package gowebapi

type TimeRule struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Path string `json:"Path,omitempty"`

	ConfigString string `json:"ConfigString,omitempty"`

	ConfigStringStored string `json:"ConfigStringStored,omitempty"`

	DisplayString string `json:"DisplayString,omitempty"`

	EditorType string `json:"EditorType,omitempty"`

	IsConfigured bool `json:"IsConfigured,omitempty"`

	IsInitializing bool `json:"IsInitializing,omitempty"`

	MergeDuplicatedItems bool `json:"MergeDuplicatedItems,omitempty"`

	PlugInName string `json:"PlugInName,omitempty"`

	Links *TimeRuleLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
