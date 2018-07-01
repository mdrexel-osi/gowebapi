package gowebapi

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	// AF Analysis
	analysisWebId := NewAFAnalysisWebID("CLSAF\\Wells\\West Texas\\Clear Fork\\Rig1\\Analyses[Completion Stage]")
	fmt.Println("Analysis:\t\t", "P1XsQ0xTQUZcV0VMTFNcV0VTVCBURVhBU1xDTEVBUiBGT1JLXFJJRzFcQU5BTFlTRVNbQ09NUExFVElPTiBTVEFHRV0" == EncodeWebID(analysisWebId))
	if "P1XsQ0xTQUZcV0VMTFNcV0VTVCBURVhBU1xDTEVBUiBGT1JLXFJJRzFcQU5BTFlTRVNbQ09NUExFVElPTiBTVEFHRV0" != EncodeWebID(analysisWebId) {
		t.Fail()
	}

	// AF AnalysisCategory
	analysisCategoryWebId := NewAFAnalysisCategoryWebID("CLSAF\\Wells\\CategoriesAnalysis[DrillingCompletionStages]")
	fmt.Println("Analysis Category:\t", "P1XCQ0xTQUZcV0VMTFNcQ0FURUdPUklFU0FOQUxZU0lTW0RSSUxMSU5HQ09NUExFVElPTlNUQUdFU10" == EncodeWebID(analysisCategoryWebId))
	if "P1XCQ0xTQUZcV0VMTFNcQ0FURUdPUklFU0FOQUxZU0lTW0RSSUxMSU5HQ09NUExFVElPTlNUQUdFU10" != EncodeWebID(analysisCategoryWebId) {
		t.Fail()
	}

	// AF AnalysisTemplate
	analysisTeplateWebId := NewAFAnalysisTemplateWebID("CLSAF\\Wells\\AnalysisTemplates[Completion Stage]")
	fmt.Println("Analysis Template:\t", "P1XTQ0xTQUZcV0VMTFNcQU5BTFlTSVNURU1QTEFURVNbQ09NUExFVElPTiBTVEFHRV0" == EncodeWebID(analysisTeplateWebId))
	if "P1XTQ0xTQUZcV0VMTFNcQU5BTFlTSVNURU1QTEFURVNbQ09NUExFVElPTiBTVEFHRV0" != EncodeWebID(analysisTeplateWebId) {
		t.Fail()
	}

	// AF AnalysisRule
	analysisRuleWebId := NewAFAnalysisRuleWebID("CLSAF\\Wells\\West Texas\\Clear Fork\\Rig1\\Analyses[Completion Stage]\\AnalysisRule", ANALYSIS_RULE_OWNER_ANALYSIS)
	fmt.Println("Analysis Rule:\t\t", "P1XRXQ0xTQUZcV0VMTFNcV0VTVCBURVhBU1xDTEVBUiBGT1JLXFJJRzFcQU5BTFlTRVNbQ09NUExFVElPTiBTVEFHRV1cQU5BTFlTSVNSVUxF" == EncodeWebID(analysisRuleWebId))
	if "P1XRXQ0xTQUZcV0VMTFNcV0VTVCBURVhBU1xDTEVBUiBGT1JLXFJJRzFcQU5BTFlTRVNbQ09NUExFVElPTiBTVEFHRV1cQU5BTFlTSVNSVUxF" != EncodeWebID(analysisRuleWebId) {
		t.Fail()
	}

	// AF AnalysisRulePlugin
	analysisRulePluginWebId := NewAFAnalysisRulePluginWebID("CLSAF\\PlugInsAnalysisRule[Comparison]")
	fmt.Println("Analysis Rule Plugin\t", "P1XPQ0xTQUZcUExVR0lOU0FOQUxZU0lTUlVMRVtDT01QQVJJU09OXQ" == EncodeWebID(analysisRulePluginWebId))
	if "P1XPQ0xTQUZcUExVR0lOU0FOQUxZU0lTUlVMRVtDT01QQVJJU09OXQ" != EncodeWebID(analysisRulePluginWebId) {
		t.Fail()
	}

	// AF Attribute
	webID := NewAFAttributeWebID("CLSAF\\Wells\\West Texas\\Clear Fork\\Rig1|API Well Number", IS_AF_ELEMENT)
	fmt.Println("Attribute:\t\t", "P1AbEQ0xTQUZcV0VMTFNcV0VTVCBURVhBU1xDTEVBUiBGT1JLXFJJRzF8QVBJIFdFTEwgTlVNQkVS" == EncodeWebID(webID))
	if "P1AbEQ0xTQUZcV0VMTFNcV0VTVCBURVhBU1xDTEVBUiBGT1JLXFJJRzF8QVBJIFdFTEwgTlVNQkVS" != EncodeWebID(webID) {
		t.Fail()
	}

	// AF AttributeCategory
	attributeCategoryWebId := NewAFAttributeCategoryWebID("CLSAF\\WellsDb\\CategoriesAttribute[Configuration]")
	fmt.Println("Attribute Category:\t", "P1ACQ0xTQUZcV0VMTFNEQlxDQVRFR09SSUVTQVRUUklCVVRFW0NPTkZJR1VSQVRJT05d" == EncodeWebID(attributeCategoryWebId))
	if "P1ACQ0xTQUZcV0VMTFNEQlxDQVRFR09SSUVTQVRUUklCVVRFW0NPTkZJR1VSQVRJT05d" != EncodeWebID(attributeCategoryWebId) {
		t.Fail()
	}

	// AF AttributeTemplateWebID
	attributeTemplateWebId := NewAFAttributeTemplateWebID("CLSAF\\AMI_Meters\\ElementTemplates[UIQMeter]|ContractorID")
	fmt.Println("Attribute Template:\t", "P1ATEQ0xTQUZcQU1JX01FVEVSU1xFTEVNRU5UVEVNUExBVEVTW1VJUU1FVEVSXXxDT05UUkFDVE9SSUQ" == EncodeWebID(attributeTemplateWebId))
	if "P1ATEQ0xTQUZcQU1JX01FVEVSU1xFTEVNRU5UVEVNUExBVEVTW1VJUU1FVEVSXXxDT05UUkFDVE9SSUQ" != EncodeWebID(attributeTemplateWebId) {
		t.Fail()
	}

	// AF DatabaseWebID
	databaseWebId := NewAFDatabaseWebID("CLSAF\\AMI_Meters")
	fmt.Println("Database:\t\t", "P1RDQ0xTQUZcQU1JX01FVEVSUw" == EncodeWebID(databaseWebId))
	if "P1RDQ0xTQUZcQU1JX01FVEVSUw" != EncodeWebID(databaseWebId) {
		t.Fail()
	}

	// AF ElementCategory
	elementCategoryWebId := NewAFElementCategoryWebID("CLSAF\\Wells\\CategoriesElement[Drilling Events]")
	fmt.Println("Element Category:\t", "P1ECQ0xTQUZcV0VMTFNcQ0FURUdPUklFU0VMRU1FTlRbRFJJTExJTkcgRVZFTlRTXQ" == EncodeWebID(elementCategoryWebId))
	if "P1ECQ0xTQUZcV0VMTFNcQ0FURUdPUklFU0VMRU1FTlRbRFJJTExJTkcgRVZFTlRTXQ" != EncodeWebID(elementCategoryWebId) {
		t.Fail()
	}

	// AF Element
	elementWebID := NewAFElementWebID("CLSAF\\Wells\\West Texas\\Clear Fork")
	fmt.Println("Element:\t\t", "P1EmQ0xTQUZcV0VMTFNcV0VTVCBURVhBU1xDTEVBUiBGT1JL" == EncodeWebID(elementWebID))
	if "P1EmQ0xTQUZcV0VMTFNcV0VTVCBURVhBU1xDTEVBUiBGT1JL" != EncodeWebID(elementWebID) {
		t.Fail()
	}

	// AF Element Template
	elementTemplateWebId := NewAFElementTemplateWebID("CLSAF\\Wells\\ElementTemplates[Completion Stage]")
	fmt.Println("Element Template:\t", "P1ETQ0xTQUZcV0VMTFNcRUxFTUVOVFRFTVBMQVRFU1tDT01QTEVUSU9OIFNUQUdFXQ" == EncodeWebID(elementTemplateWebId))
	if "P1ETQ0xTQUZcV0VMTFNcRUxFTUVOVFRFTVBMQVRFU1tDT01QTEVUSU9OIFNUQUdFXQ" != EncodeWebID(elementTemplateWebId) {
		t.Fail()
	}

	// AF EnumerationSet
	enumerationSetWebId := NewAFEnumerationSetWebID("CLSAF\\Wells\\EnumerationSets[Well States]", ENUM_SOURCE_MARKER_IS_PISYSTEM)
	fmt.Println("Enumeration Set:\t", "P1MSRQ0xTQUZcV0VMTFNcRU5VTUVSQVRJT05TRVRTW1dFTEwgU1RBVEVTXQ" == EncodeWebID(enumerationSetWebId))
	if "P1MSRQ0xTQUZcV0VMTFNcRU5VTUVSQVRJT05TRVRTW1dFTEwgU1RBVEVTXQ" != EncodeWebID(enumerationSetWebId) {
		t.Fail()
	}

	// AF EnumerationValue
	enumerationValueWebId := NewAFEnumerationValueWebID("CLSAF\\Wells\\EnumerationSets[Well States]\\Drilling", ENUM_SOURCE_MARKER_IS_PISYSTEM)
	fmt.Println("Enumeration Value:\t", "P1MVRQ0xTQUZcV0VMTFNcRU5VTUVSQVRJT05TRVRTW1dFTEwgU1RBVEVTXVxEUklMTElORw" == EncodeWebID(enumerationValueWebId))
	if "P1MVRQ0xTQUZcV0VMTFNcRU5VTUVSQVRJT05TRVRTW1dFTEwgU1RBVEVTXVxEUklMTElORw" != EncodeWebID(enumerationValueWebId) {
		t.Fail()
	}

	// AF EventFrame
	eventFrameWebId := NewAFEventFrameWebID("CLSAF\\Wells\\EventFrames[OSIDEMO_Rig1 Rotary Drilling 2018-06-06 10:00:30.000]")
	fmt.Println("Event Frame:\t\t", "P1FmQ0xTQUZcV0VMTFNcRVZFTlRGUkFNRVNbT1NJREVNT19SSUcxIFJPVEFSWSBEUklMTElORyAyMDE4LTA2LTA2IDEwOjAwOjMwLjAwMF0" == EncodeWebID(eventFrameWebId))
	if "P1FmQ0xTQUZcV0VMTFNcRVZFTlRGUkFNRVNbT1NJREVNT19SSUcxIFJPVEFSWSBEUklMTElORyAyMDE4LTA2LTA2IDEwOjAwOjMwLjAwMF0" != EncodeWebID(eventFrameWebId) {
		t.Fail()
	}

	// AF Notification
	notificationWebId := NewAFNotificationWebID("THIS_ISNT_USED")
	fmt.Println("Notification:\t\t", EncodeWebID(notificationWebId) == EncodeWebID(notificationWebId))
	if "P1NfVEhJU19JU05UX1VTRUQ" != EncodeWebID(notificationWebId) {
		t.Fail()
	}

	// AF NotificationTemplate
	notificationTemplateWebId := NewAFNotificationTemplateWebID("CLSAF\\Wells\\ElementTemplates[Rig]\\NotificationRuleTemplates[Notification Rule Template]")
	fmt.Println("Notification Template:\t", "P1NTQ0xTQUZcV0VMTFNcRUxFTUVOVFRFTVBMQVRFU1tSSUddXE5PVElGSUNBVElPTlJVTEVURU1QTEFURVNbTk9USUZJQ0FUSU9OIFJVTEUgVEVNUExBVEVd" == EncodeWebID(notificationTemplateWebId))
	if "P1NTQ0xTQUZcV0VMTFNcRUxFTUVOVFRFTVBMQVRFU1tSSUddXE5PVElGSUNBVElPTlJVTEVURU1QTEFURVNbTk9USUZJQ0FUSU9OIFJVTEUgVEVNUExBVEVd" != EncodeWebID(notificationTemplateWebId) {
		t.Fail()
	}

	notificationContactTemplateWebId := NewAFNotificationContactTemplateWebID("CLSAF\\NotificationContactTemplates[Christopher Sawyer - Email]")
	fmt.Println("Notification Contact Template:\t", "P1NCQ0xTQUZcTk9USUZJQ0FUSU9OQ09OVEFDVFRFTVBMQVRFU1tDSFJJU1RPUEhFUiBTQVdZRVIgLSBFTUFJTF0" == EncodeWebID(notificationContactTemplateWebId))
	if "P1NCQ0xTQUZcTk9USUZJQ0FUSU9OQ09OVEFDVFRFTVBMQVRFU1tDSFJJU1RPUEhFUiBTQVdZRVIgLSBFTUFJTF0" != EncodeWebID(notificationContactTemplateWebId) {
		t.Fail()
	}

	// AF TimeRule
	timeRuleWebId := NewAFTimeRuleWebID("CLSAF\\Wells\\West Texas\\Clear Fork\\Rig1\\Analyses[Completion Stage]\\TimeRule", TIME_RULE_OWNER_ANALYSIS)
	fmt.Println("Time Rule:\t\t", "P1TRXQ0xTQUZcV0VMTFNcV0VTVCBURVhBU1xDTEVBUiBGT1JLXFJJRzFcQU5BTFlTRVNbQ09NUExFVElPTiBTVEFHRV1cVElNRVJVTEU" == EncodeWebID(timeRuleWebId))
	if "P1TRXQ0xTQUZcV0VMTFNcV0VTVCBURVhBU1xDTEVBUiBGT1JLXFJJRzFcQU5BTFlTRVNbQ09NUExFVElPTiBTVEFHRV1cVElNRVJVTEU" != EncodeWebID(timeRuleWebId) {
		t.Fail()
	}

	// AF TimeRulePlugin
	timeRulePluginWebId := NewAFTimeRulePluginWebID("CLSAF\\PlugInsTimeRule[Natural]")
	fmt.Println("Time Rule Plugin:\t", "P1TPQ0xTQUZcUExVR0lOU1RJTUVSVUxFW05BVFVSQUxd" == EncodeWebID(timeRulePluginWebId))
	if "P1TPQ0xTQUZcUExVR0lOU1RJTUVSVUxFW05BVFVSQUxd" != EncodeWebID(timeRulePluginWebId) {
		t.Fail()
	}

	// AF SecurityIdentity
	securityIdentitiesWebId := NewAFSecurityIdentityWebID("CLSAF\\SecurityIdentities[Administrators]")
	fmt.Println("Security Identity:\t", "P1SIQ0xTQUZcU0VDVVJJVFlJREVOVElUSUVTW0FETUlOSVNUUkFUT1JTXQ" == EncodeWebID(securityIdentitiesWebId))
	if "P1SIQ0xTQUZcU0VDVVJJVFlJREVOVElUSUVTW0FETUlOSVNUUkFUT1JTXQ" != EncodeWebID(securityIdentitiesWebId) {
		t.Fail()
	}

	// AF SecurityMapping
	securityMappingWebId := NewAFSecurityMappingWebID("CLSAF\\SecurityMappings['BUILTIN\\Administrators']")
	fmt.Println("Security Mapping:\t", "P1SMQ0xTQUZcU0VDVVJJVFlNQVBQSU5HU1snQlVJTFRJTlxBRE1JTklTVFJBVE9SUydd" == EncodeWebID(securityMappingWebId))
	if "P1SMQ0xTQUZcU0VDVVJJVFlNQVBQSU5HU1snQlVJTFRJTlxBRE1JTklTVFJBVE9SUydd" != EncodeWebID(securityMappingWebId) {
		t.Fail()
	}

	// AF Table
	tableWebId := NewAFTableWebID("CLSAF\\WellsDb\\Tables[Contractors]")
	fmt.Println("Table:\t\t\t", "P1BlQ0xTQUZcV0VMTFNEQlxUQUJMRVNbQ09OVFJBQ1RPUlNd" == EncodeWebID(tableWebId))
	if "P1BlQ0xTQUZcV0VMTFNEQlxUQUJMRVNbQ09OVFJBQ1RPUlNd" != EncodeWebID(tableWebId) {
		t.Fail()
	}

	// AF TableCategory
	tableCategoryWebId := NewAFTableCategoryWebID("CLSAF\\WellsDb\\CategoriesTable[NewTableCategory]")
	fmt.Println("Table Category:\t\t", "P1BCQ0xTQUZcV0VMTFNEQlxDQVRFR09SSUVTVEFCTEVbTkVXVEFCTEVDQVRFR09SWV0" == EncodeWebID(tableCategoryWebId))
	if "P1BCQ0xTQUZcV0VMTFNEQlxDQVRFR09SSUVTVEFCTEVbTkVXVEFCTEVDQVRFR09SWV0" != EncodeWebID(tableCategoryWebId) {
		t.Fail()
	}

	// PI Point
	piPointWebId := NewPIPointWebID("CLSAF\\SINUSOID")
	fmt.Println("PI Point:\t\t", "P1DPQ0xTQUZcU0lOVVNPSUQ" == EncodeWebID(piPointWebId))
	if "P1DPQ0xTQUZcU0lOVVNPSUQ" != EncodeWebID(piPointWebId) {
		t.Fail()
	}

	// PI Server
	piServerWebId := NewPIServerWebID("CLSAF")
	fmt.Println("PI Server:\t\t", "P1DSQ0xTQUY" == EncodeWebID(piServerWebId))
	if "P1DSQ0xTQUY" != EncodeWebID(piServerWebId) {
		t.Fail()
	}

	// PI System
	piSystemWebId := NewPISystemWebID("CLSAF")
	fmt.Println("PI System:\t\t", "P1RSQ0xTQUY" == EncodeWebID(piSystemWebId))
	if "P1RSQ0xTQUY" != EncodeWebID(piSystemWebId) {
		t.Fail()
	}

	// UOM
	unitWebId := NewUOMWebID("CLSAF\\UOMDatabase\\count")
	fmt.Println("UOM:\t\t\t", "P1UtQ0xTQUZcVU9NREFUQUJBU0VcQ09VTlQ" == EncodeWebID(unitWebId))
	if "P1UtQ0xTQUZcVU9NREFUQUJBU0VcQ09VTlQ" != EncodeWebID(unitWebId) {
		t.Fail()
	}

	// UOM Class
	unitClassWebId := NewUOMClassWebID("CLSAF\\UOMDatabase\\UOMClasses[Quantity]")
	fmt.Println("UOM Class:\t\t", "P1UCQ0xTQUZcVU9NREFUQUJBU0VcVU9NQ0xBU1NFU1tRVUFOVElUWV0" == EncodeWebID(unitClassWebId))
	if "P1UCQ0xTQUZcVU9NREFUQUJBU0VcVU9NQ0xBU1NFU1tRVUFOVElUWV0" != EncodeWebID(unitClassWebId) {
		t.Fail()
	}
}
