
package gowebapi

type Request struct {
	Method string `json:"Method,omitempty"`

	Resource string `json:"Resource,omitempty"`

	RequestTemplate *RequestTemplate `json:"RequestTemplate,omitempty"`

	Parameters []string `json:"Parameters,omitempty"`

	Headers map[string]string `json:"Headers,omitempty"`

	Content string `json:"Content,omitempty"`

	ParentIds []string `json:"ParentIds,omitempty"`
}
