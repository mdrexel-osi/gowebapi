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

type SecurityRights struct {
	OwnerWebId string `json:"OwnerWebId,omitempty"`

	SecurityItem string `json:"SecurityItem,omitempty"`

	UserIdentity string `json:"UserIdentity,omitempty"`

	Links *SecurityRightsLinks `json:"Links,omitempty"`

	CanAnnotate bool `json:"CanAnnotate,omitempty"`

	CanDelete bool `json:"CanDelete,omitempty"`

	CanExecute bool `json:"CanExecute,omitempty"`

	CanRead bool `json:"CanRead,omitempty"`

	CanReadData bool `json:"CanReadData,omitempty"`

	CanSubscribe bool `json:"CanSubscribe,omitempty"`

	CanSubscribeOthers bool `json:"CanSubscribeOthers,omitempty"`

	CanWrite bool `json:"CanWrite,omitempty"`

	CanWriteData bool `json:"CanWriteData,omitempty"`

	HasAdmin bool `json:"HasAdmin,omitempty"`

	Rights []string `json:"Rights,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
