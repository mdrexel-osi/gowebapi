
package gowebapi

type AnalysisLinks struct {
	Self string `json:"Self,omitempty"`

	Target string `json:"Target,omitempty"`

	Database string `json:"Database,omitempty"`

	Categories string `json:"Categories,omitempty"`

	Template string `json:"Template,omitempty"`

	AnalysisRule string `json:"AnalysisRule,omitempty"`

	AnalysisRulePlugIn string `json:"AnalysisRulePlugIn,omitempty"`

	TimeRule string `json:"TimeRule,omitempty"`

	TimeRulePlugIn string `json:"TimeRulePlugIn,omitempty"`

	Security string `json:"Security,omitempty"`

	SecurityEntries string `json:"SecurityEntries,omitempty"`
}
