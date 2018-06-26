
package gowebapi

type StreamSummaries struct {
	WebId string `json:"WebId,omitempty"`

	Name string `json:"Name,omitempty"`

	Path string `json:"Path,omitempty"`

	Items []SummaryValue `json:"Items,omitempty"`

	Links *StreamSummariesLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
