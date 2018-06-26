
package gowebapi

type AnalysisRule struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Path string `json:"Path,omitempty"`

	ConfigString string `json:"ConfigString,omitempty"`

	DisplayString string `json:"DisplayString,omitempty"`

	EditorType string `json:"EditorType,omitempty"`

	HasChildren bool `json:"HasChildren,omitempty"`

	IsConfigured bool `json:"IsConfigured,omitempty"`

	IsInitializing bool `json:"IsInitializing,omitempty"`

	PlugInName string `json:"PlugInName,omitempty"`

	SupportedBehaviors []string `json:"SupportedBehaviors,omitempty"`

	VariableMapping string `json:"VariableMapping,omitempty"`

	Links *AnalysisRuleLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
