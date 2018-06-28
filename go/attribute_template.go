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

type AttributeTemplate struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Path string `json:"Path,omitempty"`

	Type_ string `json:"Type,omitempty"`

	TypeQualifier string `json:"TypeQualifier,omitempty"`

	DefaultUnitsName string `json:"DefaultUnitsName,omitempty"`

	DefaultValue *interface{} `json:"DefaultValue,omitempty"`

	DataReferencePlugIn string `json:"DataReferencePlugIn,omitempty"`

	ConfigString string `json:"ConfigString,omitempty"`

	IsConfigurationItem bool `json:"IsConfigurationItem,omitempty"`

	IsExcluded bool `json:"IsExcluded,omitempty"`

	IsHidden bool `json:"IsHidden,omitempty"`

	IsManualDataEntry bool `json:"IsManualDataEntry,omitempty"`

	HasChildren bool `json:"HasChildren,omitempty"`

	CategoryNames []string `json:"CategoryNames,omitempty"`

	TraitName string `json:"TraitName,omitempty"`

	DefaultUnitsNameAbbreviation string `json:"DefaultUnitsNameAbbreviation,omitempty"`

	Links *AttributeTemplateLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
