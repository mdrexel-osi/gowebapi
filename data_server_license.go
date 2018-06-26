
package gowebapi

type DataServerLicense struct {
	AmountLeft string `json:"AmountLeft,omitempty"`

	AmountUsed string `json:"AmountUsed,omitempty"`

	Name string `json:"Name,omitempty"`

	TotalAmount string `json:"TotalAmount,omitempty"`

	Links *DataServerLicenseLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
