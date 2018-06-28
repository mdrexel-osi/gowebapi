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

type AnalysisTemplate struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Path string `json:"Path,omitempty"`

	AnalysisRulePlugInName string `json:"AnalysisRulePlugInName,omitempty"`

	CategoryNames []string `json:"CategoryNames,omitempty"`

	CreateEnabled bool `json:"CreateEnabled,omitempty"`

	GroupId int32 `json:"GroupId,omitempty"`

	HasNotificationTemplate bool `json:"HasNotificationTemplate,omitempty"`

	HasTarget bool `json:"HasTarget,omitempty"`

	OutputTime string `json:"OutputTime,omitempty"`

	TargetName string `json:"TargetName,omitempty"`

	TimeRulePlugInName string `json:"TimeRulePlugInName,omitempty"`

	Links *AnalysisTemplateLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
