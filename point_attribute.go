
package gowebapi

type PointAttribute struct {
	Name string `json:"Name,omitempty"`

	Value *interface{} `json:"Value,omitempty"`

	Links *PointAttributeLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
