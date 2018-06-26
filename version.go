
package gowebapi

type Version struct {
	FullVersion string `json:"FullVersion,omitempty"`

	MajorMinorRevision string `json:"MajorMinorRevision,omitempty"`

	Build string `json:"Build,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
