
package gowebapi

type Substatus struct {
	Substatus int32 `json:"Substatus,omitempty"`

	Message string `json:"Message,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
