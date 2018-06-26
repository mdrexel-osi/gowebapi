
package gowebapi

type AnalysisCategory struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Path string `json:"Path,omitempty"`

	Links *AnalysisCategoryLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
