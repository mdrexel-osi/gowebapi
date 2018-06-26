
package gowebapi

type PointLinks struct {
	Self string `json:"Self,omitempty"`

	DataServer string `json:"DataServer,omitempty"`

	Attributes string `json:"Attributes,omitempty"`

	InterpolatedData string `json:"InterpolatedData,omitempty"`

	RecordedData string `json:"RecordedData,omitempty"`

	PlotData string `json:"PlotData,omitempty"`

	SummaryData string `json:"SummaryData,omitempty"`

	Value string `json:"Value,omitempty"`

	EndValue string `json:"EndValue,omitempty"`
}
