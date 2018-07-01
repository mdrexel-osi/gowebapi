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
	"encoding/base64"
	"fmt"
	"log"
	"strings"
)

type BaseElementType int
type AnalysisRuleOwnerType int
type EnumerationSourceMarkerType int
type TimeRuleOwnerType int

// Used by AFAnalysisRuleWebID
const (
	ANALYSIS_RULE_OWNER_ANALYSIS AnalysisRuleOwnerType = 0
	ANALYSIS_RULE_OWNER_TEMPLATE AnalysisRuleOwnerType = 1
)

// Used by AFAttribute
const (
	IS_AF_ELEMENT      BaseElementType = 0
	IS_AF_EVENTFRAME   BaseElementType = 1
	IS_AF_NOTIFICATION BaseElementType = 2
)

// Used by AFEnumerationSet
const (
	ENUM_SOURCE_MARKER_IS_PISYSTEM = 0
	ENUM_SOURCE_MARKER_IS_PISERVER = 1
)

// Used by AFTimeRule
const (
	TIME_RULE_OWNER_ANALYSIS         = 0
	TIME_RULE_OWNER_ANALYSISTEMPLATE = 1
)

// *************************************************
// Structures to represent decoded WebIDs
// *************************************************

type AFAnalysisWebID struct {
	Type    string
	Version string
	Marker  string
	Name    string
}

type AFAnalysisCategoryWebID struct {
	Type    string
	Version string
	Marker  string
	Name    string
}

type AFAnalysisTemplateWebID struct {
	Type    string
	Version string
	Marker  string
	Name    string
}

type AFAnalysisRuleWebID struct {
	Type        string
	Version     string
	Marker      string
	OwnerMarker string
	Name        string
}
type AFAnalysisRulePluginWebID struct {
	Type    string
	Version string
	Marker  string
	Name    string
}

type AFAttributeWebID struct {
	Type              string
	Version           string
	Marker            string
	BaseElementMarker string
	Name              string
}

type AFAttributeCategoryWebID struct {
	Type    string
	Version string
	Marker  string
	Name    string
}

type AFAttributeTemplateWebID struct {
	Type                  string
	Version               string
	Marker                string
	ElementTemplateMarker string
	Name                  string
}

type AFDatabaseWebID struct {
	Type    string
	Version string
	Marker  string
	Name    string
}

type AFElementWebID struct {
	Type    string
	Version string
	Marker  string
	Name    string
}

type AFElementCategoryWebID struct {
	Type    string
	Version string
	Marker  string
	Name    string
}

type AFElementTemplateWebID struct {
	Type    string
	Version string
	Marker  string
	Name    string
}

type AFEnumerationSetWebID struct {
	Type         string
	Version      string
	Marker       string
	SourceMarker string
	Name         string
}

type AFEnumerationValueWebID struct {
	Type         string
	Version      string
	Marker       string
	SourceMarker string
	Name         string
}

type AFEventFrameWebID struct {
	Type    string
	Version string
	Marker  string
	Name    string
}

type AFNotificationWebID struct {
	Type    string
	Version string
	Marker  string
	Name    string
}

type AFNotificationTemplateWebID struct {
	Type    string
	Version string
	Marker  string
	Name    string
}

type AFNotificationContactTemplateWebID struct {
	Type    string
	Version string
	Marker  string
	Name    string
}

type AFTimeRuleWebID struct {
	Type        string
	Version     string
	Marker      string
	OwnerMarker string
	Name        string
}

type AFTimeRulePluginWebID struct {
	Type    string
	Version string
	Marker  string
	Name    string
}

type AFSecurityIdentityWebId struct {
	Type    string
	Version string
	Marker  string
	Name    string
}

type AFSecurityMappingWebId struct {
	Type    string
	Version string
	Marker  string
	Name    string
}

type AFTableWebId struct {
	Type    string
	Version string
	Marker  string
	Name    string
}

type AFTableCategoryWebId struct {
	Type    string
	Version string
	Marker  string
	Name    string
}

type PIPointWebId struct {
	Type    string
	Version string
	Marker  string
	Name    string
}

type PIServerWebId struct {
	Type    string
	Version string
	Marker  string
	Name    string
}

type PISystemWebId struct {
	Type    string
	Version string
	Marker  string
	Name    string
}

type UOMWebId struct {
	Type    string
	Version string
	Marker  string
	Name    string
}

type UOMClassWebId struct {
	Type    string
	Version string
	Marker  string
	Name    string
}

// *************************************************
// Functions to create parseable WebID structures
// *************************************************

func NewAFAnalysisWebID(path string) AFAnalysisWebID {
	var webID AFAnalysisWebID

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "Xs"
	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewAFAnalysisCategoryWebID(path string) AFAnalysisCategoryWebID {
	var webID AFAnalysisCategoryWebID

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "XC"
	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewAFAnalysisTemplateWebID(path string) AFAnalysisTemplateWebID {
	var webID AFAnalysisTemplateWebID

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "XT"
	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewAFAnalysisRuleWebID(path string, ownerType AnalysisRuleOwnerType) AFAnalysisRuleWebID {
	var webID AFAnalysisRuleWebID

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "XR"

	switch ownerType {
	case ANALYSIS_RULE_OWNER_ANALYSIS:
		webID.OwnerMarker = "X"
	case ANALYSIS_RULE_OWNER_TEMPLATE:
		webID.OwnerMarker = "T"
	}

	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewAFAnalysisRulePluginWebID(path string) AFAnalysisRulePluginWebID {
	var webID AFAnalysisRulePluginWebID

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "XP"
	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewAFAttributeWebID(path string, baseType BaseElementType) AFAttributeWebID {
	var webID AFAttributeWebID

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "Ab"

	switch baseType {
	case 0:
		webID.BaseElementMarker = "E"
	case 1:
		webID.BaseElementMarker = "F"
	case 2:
		webID.BaseElementMarker = "N"
	}

	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewAFAttributeCategoryWebID(path string) AFAttributeCategoryWebID {
	var webID AFAttributeCategoryWebID

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "AC"
	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewAFAttributeTemplateWebID(path string) AFAttributeTemplateWebID {
	var webID AFAttributeTemplateWebID

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "AT"
	webID.ElementTemplateMarker = "E"
	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewAFDatabaseWebID(path string) AFDatabaseWebID {
	var webID AFDatabaseWebID

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "RD"
	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewAFElementWebID(path string) AFElementWebID {
	var webID AFElementWebID

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "Em"
	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewAFElementCategoryWebID(path string) AFElementCategoryWebID {
	var webID AFElementCategoryWebID

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "EC"
	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewAFElementTemplateWebID(path string) AFElementTemplateWebID {
	var webID AFElementTemplateWebID

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "ET"
	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewAFEnumerationSetWebID(path string, sourceMarker EnumerationSourceMarkerType) AFEnumerationSetWebID {
	var webID AFEnumerationSetWebID

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "MS"

	switch sourceMarker {
	case ENUM_SOURCE_MARKER_IS_PISYSTEM:
		webID.SourceMarker = "R"
	case ENUM_SOURCE_MARKER_IS_PISERVER:
		webID.SourceMarker = "D"
	}

	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewAFEnumerationValueWebID(path string, sourceMarker EnumerationSourceMarkerType) AFEnumerationValueWebID {
	var webID AFEnumerationValueWebID

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "MV"

	switch sourceMarker {
	case ENUM_SOURCE_MARKER_IS_PISYSTEM:
		webID.SourceMarker = "R"
	case ENUM_SOURCE_MARKER_IS_PISERVER:
		webID.SourceMarker = "D"
	}

	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewAFEventFrameWebID(path string) AFEventFrameWebID {
	var webID AFEventFrameWebID

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "Fm"
	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewAFNotificationWebID(path string) AFNotificationWebID {
	var webID AFNotificationWebID

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "Nf"
	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewAFNotificationTemplateWebID(path string) AFNotificationTemplateWebID {
	var webID AFNotificationTemplateWebID

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "NT"
	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewAFNotificationContactTemplateWebID(path string) AFNotificationContactTemplateWebID {
	var webID AFNotificationContactTemplateWebID

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "NC"
	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewAFTimeRuleWebID(path string, ownerType TimeRuleOwnerType) AFTimeRuleWebID {
	var webID AFTimeRuleWebID

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "TR"

	switch ownerType {
	case TIME_RULE_OWNER_ANALYSIS:
		webID.OwnerMarker = "X"
	case TIME_RULE_OWNER_ANALYSISTEMPLATE:
		webID.OwnerMarker = "T"
	}

	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewAFTimeRulePluginWebID(path string) AFTimeRulePluginWebID {
	var webID AFTimeRulePluginWebID

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "TP"

	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewAFSecurityIdentityWebID(path string) AFSecurityIdentityWebId {
	var webID AFSecurityIdentityWebId

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "SI"
	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewAFSecurityMappingWebID(path string) AFSecurityMappingWebId {
	var webID AFSecurityMappingWebId

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "SM"
	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewAFTableWebID(path string) AFTableWebId {
	var webID AFTableWebId

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "Bl"
	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewAFTableCategoryWebID(path string) AFTableCategoryWebId {
	var webID AFTableCategoryWebId

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "BC"
	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewPIPointWebID(path string) PIPointWebId {
	var webID PIPointWebId

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "DP"
	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewPIServerWebID(path string) PIServerWebId {
	var webID PIServerWebId

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "DS"
	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewPISystemWebID(path string) PISystemWebId {
	var webID PISystemWebId

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "RS"
	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewUOMWebID(path string) UOMWebId {
	var webID UOMWebId

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "Ut"
	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

func NewUOMClassWebID(path string) UOMClassWebId {
	var webID UOMClassWebId

	webID.Type = "P"
	webID.Version = "1"
	webID.Marker = "UC"
	webID.Name = strings.TrimSpace(strings.ToUpper(path))

	return webID
}

// *************************************************
// Functions that process WebIDs
// *************************************************

// Converts a string to a de-padded Base64 string
func Base64EncodeNoPadding(text string) string {
	return strings.Replace(base64.StdEncoding.EncodeToString([]byte(text)), "=", "", -1)
}

// By passing in any number of WebID structures, you get back a valid WebID to use
// when calling PI Web API endpoints or sending structure lists.
func EncodeWebID(webIdStructure interface{}) (webIDString string) {

	switch webIdStructure.(type) {
	case AFAnalysisWebID:
		webId := webIdStructure.(AFAnalysisWebID)
		return fmt.Sprintf("%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			Base64EncodeNoPadding(webId.Name),
		)
	case AFAnalysisCategoryWebID:
		webId := webIdStructure.(AFAnalysisCategoryWebID)
		return fmt.Sprintf("%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			Base64EncodeNoPadding(webId.Name),
		)
	case AFAnalysisTemplateWebID:
		webId := webIdStructure.(AFAnalysisTemplateWebID)
		return fmt.Sprintf("%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			Base64EncodeNoPadding(webId.Name),
		)
	case AFAnalysisRuleWebID:
		webId := webIdStructure.(AFAnalysisRuleWebID)
		return fmt.Sprintf("%v%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			webId.OwnerMarker,
			Base64EncodeNoPadding(webId.Name),
		)
	case AFAnalysisRulePluginWebID:
		webId := webIdStructure.(AFAnalysisRulePluginWebID)
		return fmt.Sprintf("%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			Base64EncodeNoPadding(webId.Name),
		)
	case AFAttributeWebID:
		webId := webIdStructure.(AFAttributeWebID)
		return fmt.Sprintf("%v%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			webId.BaseElementMarker,
			Base64EncodeNoPadding(webId.Name),
		)
	case AFAttributeCategoryWebID:
		webId := webIdStructure.(AFAttributeCategoryWebID)
		return fmt.Sprintf("%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			Base64EncodeNoPadding(webId.Name),
		)
	case AFAttributeTemplateWebID:
		webID := webIdStructure.(AFAttributeTemplateWebID)
		return fmt.Sprintf("%v%v%v%v%v",
			webID.Type,
			webID.Version,
			webID.Marker,
			webID.ElementTemplateMarker,
			Base64EncodeNoPadding(webID.Name),
		)
	case AFDatabaseWebID:
		webId := webIdStructure.(AFDatabaseWebID)
		return fmt.Sprintf("%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			Base64EncodeNoPadding(webId.Name),
		)
	case AFElementWebID:
		webId := webIdStructure.(AFElementWebID)
		return fmt.Sprintf("%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			Base64EncodeNoPadding(webId.Name),
		)
	case AFElementCategoryWebID:
		webId := webIdStructure.(AFElementCategoryWebID)
		return fmt.Sprintf("%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			Base64EncodeNoPadding(webId.Name),
		)
	case AFElementTemplateWebID:
		webId := webIdStructure.(AFElementTemplateWebID)
		return fmt.Sprintf("%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			Base64EncodeNoPadding(webId.Name),
		)
	case AFEnumerationSetWebID:
		webId := webIdStructure.(AFEnumerationSetWebID)
		return fmt.Sprintf("%v%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			webId.SourceMarker,
			Base64EncodeNoPadding(webId.Name),
		)
	case AFEnumerationValueWebID:
		webId := webIdStructure.(AFEnumerationValueWebID)
		return fmt.Sprintf("%v%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			webId.SourceMarker,
			Base64EncodeNoPadding(webId.Name),
		)
	case AFEventFrameWebID:
		webId := webIdStructure.(AFEventFrameWebID)
		return fmt.Sprintf("%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			Base64EncodeNoPadding(webId.Name),
		)
	case AFNotificationWebID:
		webId := webIdStructure.(AFNotificationWebID)
		return fmt.Sprintf("%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			Base64EncodeNoPadding(webId.Name),
		)
	case AFNotificationTemplateWebID:
		webId := webIdStructure.(AFNotificationTemplateWebID)
		return fmt.Sprintf("%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			Base64EncodeNoPadding(webId.Name),
		)
	case AFNotificationContactTemplateWebID:
		webId := webIdStructure.(AFNotificationContactTemplateWebID)
		return fmt.Sprintf("%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			Base64EncodeNoPadding(webId.Name),
		)
	case AFTimeRuleWebID:
		webId := webIdStructure.(AFTimeRuleWebID)
		return fmt.Sprintf("%v%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			webId.OwnerMarker,
			Base64EncodeNoPadding(webId.Name),
		)
	case AFTimeRulePluginWebID:
		webId := webIdStructure.(AFTimeRulePluginWebID)
		return fmt.Sprintf("%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			Base64EncodeNoPadding(webId.Name),
		)
	case AFSecurityIdentityWebId:
		webId := webIdStructure.(AFSecurityIdentityWebId)
		return fmt.Sprintf("%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			Base64EncodeNoPadding(webId.Name),
		)
	case AFSecurityMappingWebId:
		webId := webIdStructure.(AFSecurityMappingWebId)
		return fmt.Sprintf("%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			Base64EncodeNoPadding(webId.Name),
		)
	case AFTableWebId:
		webId := webIdStructure.(AFTableWebId)
		return fmt.Sprintf("%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			Base64EncodeNoPadding(webId.Name),
		)
	case AFTableCategoryWebId:
		webId := webIdStructure.(AFTableCategoryWebId)
		return fmt.Sprintf("%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			Base64EncodeNoPadding(webId.Name),
		)
	case PIPointWebId:
		webId := webIdStructure.(PIPointWebId)
		return fmt.Sprintf("%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			Base64EncodeNoPadding(webId.Name),
		)
	case PIServerWebId:
		webId := webIdStructure.(PIServerWebId)
		return fmt.Sprintf("%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			Base64EncodeNoPadding(webId.Name),
		)
	case PISystemWebId:
		webId := webIdStructure.(PISystemWebId)
		return fmt.Sprintf("%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			Base64EncodeNoPadding(webId.Name),
		)
	case UOMWebId:
		webId := webIdStructure.(UOMWebId)
		return fmt.Sprintf("%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			Base64EncodeNoPadding(webId.Name),
		)
	case UOMClassWebId:
		webId := webIdStructure.(UOMClassWebId)
		return fmt.Sprintf("%v%v%v%v",
			webId.Type,
			webId.Version,
			webId.Marker,
			Base64EncodeNoPadding(webId.Name),
		)
	}

	log.Fatal("The object passed to gowebapi.EncodeWebID is not a supported type.")
	return ""
}
