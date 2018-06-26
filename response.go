
package gowebapi

type Response struct {
	Status int32 `json:"Status,omitempty"`

	Headers map[string]string `json:"Headers,omitempty"`

	Content *interface{} `json:"Content,omitempty"`
}
