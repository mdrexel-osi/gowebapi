
package gowebapi

type WebException struct {
	StatusCode int32 `json:"StatusCode,omitempty"`

	Errors []string `json:"Errors,omitempty"`
}
