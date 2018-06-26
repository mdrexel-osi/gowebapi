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

type Analysis struct {
	WebId string `json:"WebId,omitempty"`

	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	Description string `json:"Description,omitempty"`

	Path string `json:"Path,omitempty"`

	AnalysisRulePlugInName string `json:"AnalysisRulePlugInName,omitempty"`

	AutoCreated bool `json:"AutoCreated,omitempty"`

	CategoryNames []string `json:"CategoryNames,omitempty"`

	GroupId int32 `json:"GroupId,omitempty"`

	HasNotification bool `json:"HasNotification,omitempty"`

	HasTarget bool `json:"HasTarget,omitempty"`

	HasTemplate bool `json:"HasTemplate,omitempty"`

	IsConfigured bool `json:"IsConfigured,omitempty"`

	IsTimeRuleDefinedByTemplate bool `json:"IsTimeRuleDefinedByTemplate,omitempty"`

	MaximumQueueSize int32 `json:"MaximumQueueSize,omitempty"`

	OutputTime string `json:"OutputTime,omitempty"`

	Priority string `json:"Priority,omitempty"`

	PublishResults bool `json:"PublishResults,omitempty"`

	Status string `json:"Status,omitempty"`

	TargetWebId string `json:"TargetWebId,omitempty"`

	TemplateName string `json:"TemplateName,omitempty"`

	TimeRulePlugInName string `json:"TimeRulePlugInName,omitempty"`

	Links *AnalysisLinks `json:"Links,omitempty"`

	WebException *WebException `json:"WebException,omitempty"`
}
