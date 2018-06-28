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

type Point struct {
	WebId string `json:"WebId,omitempty"`

	Id int32 `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Path string `json:"Path,omitempty"`

	Descriptor string `json:"Descriptor,omitempty"`

	PointClass string `json:"PointClass,omitempty"`

	PointType string `json:"PointType,omitempty"`

	DigitalSetName string `json:"DigitalSetName,omitempty"`

	Span float32 `json:"Span,omitempty"`

	Zero float32 `json:"Zero,omitempty"`

	EngineeringUnits string `json:"EngineeringUnits,omitempty"`

	Step bool `json:"Step,omitempty"`

	Future bool `json:"Future,omitempty"`

	DisplayDigits int32 `json:"DisplayDigits,omitempty"`

	Links *PointLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
