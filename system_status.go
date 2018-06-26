
package gowebapi

import (
	"time"
)

type SystemStatus struct {
	UpTimeInMinutes float64 `json:"UpTimeInMinutes,omitempty"`

	State string `json:"State,omitempty"`

	CacheInstances int32 `json:"CacheInstances,omitempty"`

	ServerTime time.Time `json:"ServerTime,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
