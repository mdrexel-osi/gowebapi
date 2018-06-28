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

type AssetDatabaseLinks struct {
	Self string `json:"Self,omitempty"`

	Elements string `json:"Elements,omitempty"`

	ElementTemplates string `json:"ElementTemplates,omitempty"`

	EventFrames string `json:"EventFrames,omitempty"`

	AssetServer string `json:"AssetServer,omitempty"`

	ElementCategories string `json:"ElementCategories,omitempty"`

	AttributeCategories string `json:"AttributeCategories,omitempty"`

	TableCategories string `json:"TableCategories,omitempty"`

	AnalysisCategories string `json:"AnalysisCategories,omitempty"`

	AnalysisTemplates string `json:"AnalysisTemplates,omitempty"`

	EnumerationSets string `json:"EnumerationSets,omitempty"`

	Tables string `json:"Tables,omitempty"`

	Security string `json:"Security,omitempty"`

	SecurityEntries string `json:"SecurityEntries,omitempty"`
}
