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

type AnalysisRulePlugIn struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Path string `json:"Path,omitempty"`

	AssemblyFileName string `json:"AssemblyFileName,omitempty"`

	AssemblyID string `json:"AssemblyID,omitempty"`

	AssemblyLoadProperties []string `json:"AssemblyLoadProperties,omitempty"`

	AssemblyTime time.Time `json:"AssemblyTime,omitempty"`

	CompatibilityVersion int32 `json:"CompatibilityVersion,omitempty"`

	IsBrowsable bool `json:"IsBrowsable,omitempty"`

	IsNonEditableConfig bool `json:"IsNonEditableConfig,omitempty"`

	LoadedAssemblyTime time.Time `json:"LoadedAssemblyTime,omitempty"`

	LoadedVersion string `json:"LoadedVersion,omitempty"`

	Version string `json:"Version,omitempty"`

	Links *AnalysisRulePlugInLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
