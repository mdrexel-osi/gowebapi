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

type ElementTemplateLinks struct {
	Self string `json:"Self,omitempty"`

	AnalysisTemplates string `json:"AnalysisTemplates,omitempty"`

	AttributeTemplates string `json:"AttributeTemplates,omitempty"`

	Database string `json:"Database,omitempty"`

	Categories string `json:"Categories,omitempty"`

	BaseTemplate string `json:"BaseTemplate,omitempty"`

	BaseTemplates string `json:"BaseTemplates,omitempty"`

	DerivedTemplates string `json:"DerivedTemplates,omitempty"`

	DefaultAttribute string `json:"DefaultAttribute,omitempty"`

	NotificationRuleTemplates string `json:"NotificationRuleTemplates,omitempty"`

	Security string `json:"Security,omitempty"`

	SecurityEntries string `json:"SecurityEntries,omitempty"`
}
