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

type EnumerationValue struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Value int32 `json:"Value,omitempty"`

	Path string `json:"Path,omitempty"`

	Parent string `json:"Parent,omitempty"`

	Links *EnumerationValueLinks `json:"Links,omitempty"`

	SerializeWebId bool `json:"SerializeWebId,omitempty"`

	SerializeId bool `json:"SerializeId,omitempty"`

	SerializeDescription bool `json:"SerializeDescription,omitempty"`

	SerializePath bool `json:"SerializePath,omitempty"`

	SerializeLinks bool `json:"SerializeLinks,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
