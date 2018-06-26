
package gowebapi

type TableData struct {
	Columns map[string]string `json:"Columns,omitempty"`

	Rows []map[string]interface{} `json:"Rows,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
