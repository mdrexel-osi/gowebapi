
package gowebapi

type ItemsNotificationContactTemplate struct {
	Items []NotificationContactTemplate `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
