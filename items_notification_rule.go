
package gowebapi

type ItemsNotificationRule struct {
	Items []NotificationRule `json:"Items,omitempty"`

	Links *PaginationLinks `json:"Links,omitempty"`
}
