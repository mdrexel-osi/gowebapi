
package gowebapi

type SearchByAttribute struct {
	SearchRoot string `json:"SearchRoot,omitempty"`

	ElementTemplate string `json:"ElementTemplate,omitempty"`

	Items []ValueQuery `json:"Items,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
