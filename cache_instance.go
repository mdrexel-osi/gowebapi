
package gowebapi

import (
	"time"
)

type CacheInstance struct {
	Id string `json:"Id,omitempty"`

	LastRefreshTime time.Time `json:"LastRefreshTime,omitempty"`

	WillRefreshAfter time.Time `json:"WillRefreshAfter,omitempty"`

	ScheduledExpirationTime time.Time `json:"ScheduledExpirationTime,omitempty"`

	User string `json:"User,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
