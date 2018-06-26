
package gowebapi

type Value struct {
	Value *interface{} `json:"Value,omitempty"`

	Exception *Errors `json:"Exception,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
