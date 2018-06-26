
package gowebapi

type Point struct {
	WebId string `json:"WebId,omitempty"`

	Id int32 `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Path string `json:"Path,omitempty"`

	Descriptor string `json:"Descriptor,omitempty"`

	PointClass string `json:"PointClass,omitempty"`

	PointType string `json:"PointType,omitempty"`

	DigitalSetName string `json:"DigitalSetName,omitempty"`

	Span float32 `json:"Span,omitempty"`

	Zero float32 `json:"Zero,omitempty"`

	EngineeringUnits string `json:"EngineeringUnits,omitempty"`

	Step bool `json:"Step,omitempty"`

	Future bool `json:"Future,omitempty"`

	DisplayDigits int32 `json:"DisplayDigits,omitempty"`

	Links *PointLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
