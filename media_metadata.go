// ************************************************************************
// * Copyright 2018 OSIsoft, LLC
// * Licensed under the Apache License, Version 2.0 (the "License");
// * you may not use this file except in compliance with the License.
// * You may obtain a copy of the License at
// * 
// *   <http://www.apache.org/licenses/LICENSE-2.0>
// * 
// * Unless required by applicable law or agreed to in writing, software
// * distributed under the License is distributed on an "AS IS" BASIS,
// * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// * See the License for the specific language governing permissions and
// * limitations under the License.
// ************************************************************************

package gowebapi

import (
	"time"
)

type MediaMetadata struct {
	Author string `json:"Author,omitempty"`

	ChangeDate time.Time `json:"ChangeDate,omitempty"`

	Description string `json:"Description,omitempty"`

	Name string `json:"Name,omitempty"`

	Size float32 `json:"Size,omitempty"`

	Links *MediaMetadataLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
