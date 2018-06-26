# Documentation for API Endpoints

All URIs are relative to *https://localhost/piwebapi*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AnalysisApi* | [**AnalysisCreateSecurityEntry**](docs/AnalysisApi.md#analysiscreatesecurityentry) | **Post** /analyses/{webId}/securityentries | Create a security entry owned by the analysis.
*AnalysisApi* | [**AnalysisDelete**](docs/AnalysisApi.md#analysisdelete) | **Delete** /analyses/{webId} | Delete an Analysis.
*AnalysisApi* | [**AnalysisDeleteSecurityEntry**](docs/AnalysisApi.md#analysisdeletesecurityentry) | **Delete** /analyses/{webId}/securityentries/{name} | Delete a security entry owned by the analysis.
*AnalysisApi* | [**AnalysisGet**](docs/AnalysisApi.md#analysisget) | **Get** /analyses/{webId} | Retrieve an Analysis.
*AnalysisApi* | [**AnalysisGetAnalysesQuery**](docs/AnalysisApi.md#analysisgetanalysesquery) | **Get** /analyses/search | Retrieve analyses based on the specified conditions. By default, returns all analyses.
*AnalysisApi* | [**AnalysisGetByPath**](docs/AnalysisApi.md#analysisgetbypath) | **Get** /analyses | Retrieve an Analysis by path.
*AnalysisApi* | [**AnalysisGetCategories**](docs/AnalysisApi.md#analysisgetcategories) | **Get** /analyses/{webId}/categories | Get an Analysis&#39; categories.
*AnalysisApi* | [**AnalysisGetSecurity**](docs/AnalysisApi.md#analysisgetsecurity) | **Get** /analyses/{webId}/security | Get the security information of the specified security item associated with the Analysis for a specified user.
*AnalysisApi* | [**AnalysisGetSecurityEntries**](docs/AnalysisApi.md#analysisgetsecurityentries) | **Get** /analyses/{webId}/securityentries | Retrieve the security entries associated with the analysis based on the specified criteria. By default, all security entries for this analysis are returned.
*AnalysisApi* | [**AnalysisGetSecurityEntryByName**](docs/AnalysisApi.md#analysisgetsecurityentrybyname) | **Get** /analyses/{webId}/securityentries/{name} | Retrieve the security entry associated with the analysis with the specified name.
*AnalysisApi* | [**AnalysisUpdate**](docs/AnalysisApi.md#analysisupdate) | **Patch** /analyses/{webId} | Update an Analysis.
*AnalysisApi* | [**AnalysisUpdateSecurityEntry**](docs/AnalysisApi.md#analysisupdatesecurityentry) | **Put** /analyses/{webId}/securityentries/{name} | Update a security entry owned by the analysis.
*AnalysisCategoryApi* | [**AnalysisCategoryCreateSecurityEntry**](docs/AnalysisCategoryApi.md#analysiscategorycreatesecurityentry) | **Post** /analysiscategories/{webId}/securityentries | Create a security entry owned by the analysis category.
*AnalysisCategoryApi* | [**AnalysisCategoryDelete**](docs/AnalysisCategoryApi.md#analysiscategorydelete) | **Delete** /analysiscategories/{webId} | Delete an analysis category.
*AnalysisCategoryApi* | [**AnalysisCategoryDeleteSecurityEntry**](docs/AnalysisCategoryApi.md#analysiscategorydeletesecurityentry) | **Delete** /analysiscategories/{webId}/securityentries/{name} | Delete a security entry owned by the analysis category.
*AnalysisCategoryApi* | [**AnalysisCategoryGet**](docs/AnalysisCategoryApi.md#analysiscategoryget) | **Get** /analysiscategories/{webId} | Retrieve an analysis category.
*AnalysisCategoryApi* | [**AnalysisCategoryGetByPath**](docs/AnalysisCategoryApi.md#analysiscategorygetbypath) | **Get** /analysiscategories | Retrieve an analysis category by path.
*AnalysisCategoryApi* | [**AnalysisCategoryGetSecurity**](docs/AnalysisCategoryApi.md#analysiscategorygetsecurity) | **Get** /analysiscategories/{webId}/security | Get the security information of the specified security item associated with the analysis category for a specified user.
*AnalysisCategoryApi* | [**AnalysisCategoryGetSecurityEntries**](docs/AnalysisCategoryApi.md#analysiscategorygetsecurityentries) | **Get** /analysiscategories/{webId}/securityentries | Retrieve the security entries associated with the analysis category based on the specified criteria. By default, all security entries for this analysis category are returned.
*AnalysisCategoryApi* | [**AnalysisCategoryGetSecurityEntryByName**](docs/AnalysisCategoryApi.md#analysiscategorygetsecurityentrybyname) | **Get** /analysiscategories/{webId}/securityentries/{name} | Retrieve the security entry associated with the analysis category with the specified name.
*AnalysisCategoryApi* | [**AnalysisCategoryUpdate**](docs/AnalysisCategoryApi.md#analysiscategoryupdate) | **Patch** /analysiscategories/{webId} | Update an analysis category by replacing items in its definition.
*AnalysisCategoryApi* | [**AnalysisCategoryUpdateSecurityEntry**](docs/AnalysisCategoryApi.md#analysiscategoryupdatesecurityentry) | **Put** /analysiscategories/{webId}/securityentries/{name} | Update a security entry owned by the analysis category.
*AnalysisRuleApi* | [**AnalysisRuleCreateAnalysisRule**](docs/AnalysisRuleApi.md#analysisrulecreateanalysisrule) | **Post** /analysisrules/{webId}/analysisrules | Create a new Analysis Rule as a child of an existing Analysis Rule.
*AnalysisRuleApi* | [**AnalysisRuleDelete**](docs/AnalysisRuleApi.md#analysisruledelete) | **Delete** /analysisrules/{webId} | Delete an Analysis Rule.
*AnalysisRuleApi* | [**AnalysisRuleGet**](docs/AnalysisRuleApi.md#analysisruleget) | **Get** /analysisrules/{webId} | Retrieve an Analysis Rule.
*AnalysisRuleApi* | [**AnalysisRuleGetAnalysisRules**](docs/AnalysisRuleApi.md#analysisrulegetanalysisrules) | **Get** /analysisrules/{webId}/analysisrules | Get the child Analysis Rules of the Analysis Rule.
*AnalysisRuleApi* | [**AnalysisRuleGetByPath**](docs/AnalysisRuleApi.md#analysisrulegetbypath) | **Get** /analysisrules | Retrieve an Analysis Rule by path.
*AnalysisRuleApi* | [**AnalysisRuleUpdate**](docs/AnalysisRuleApi.md#analysisruleupdate) | **Patch** /analysisrules/{webId} | Update an Analysis Rule by replacing items in its definition.
*AnalysisRulePlugInApi* | [**AnalysisRulePlugInGet**](docs/AnalysisRulePlugInApi.md#analysisrulepluginget) | **Get** /analysisruleplugins/{webId} | Retrieve an Analysis Rule Plug-in.
*AnalysisRulePlugInApi* | [**AnalysisRulePlugInGetByPath**](docs/AnalysisRulePlugInApi.md#analysisruleplugingetbypath) | **Get** /analysisruleplugins | Retrieve an Analysis Rule Plug-in by path.
*AnalysisTemplateApi* | [**AnalysisTemplateCreateFromAnalysis**](docs/AnalysisTemplateApi.md#analysistemplatecreatefromanalysis) | **Post** /analysistemplates | Create an Analysis template based upon a specified Analysis.
*AnalysisTemplateApi* | [**AnalysisTemplateCreateSecurityEntry**](docs/AnalysisTemplateApi.md#analysistemplatecreatesecurityentry) | **Post** /analysistemplates/{webId}/securityentries | Create a security entry owned by the analysis template.
*AnalysisTemplateApi* | [**AnalysisTemplateDelete**](docs/AnalysisTemplateApi.md#analysistemplatedelete) | **Delete** /analysistemplates/{webId} | Delete an analysis template.
*AnalysisTemplateApi* | [**AnalysisTemplateDeleteSecurityEntry**](docs/AnalysisTemplateApi.md#analysistemplatedeletesecurityentry) | **Delete** /analysistemplates/{webId}/securityentries/{name} | Delete a security entry owned by the analysis template.
*AnalysisTemplateApi* | [**AnalysisTemplateGet**](docs/AnalysisTemplateApi.md#analysistemplateget) | **Get** /analysistemplates/{webId} | Retrieve an analysis template.
*AnalysisTemplateApi* | [**AnalysisTemplateGetAnalysisTemplatesQuery**](docs/AnalysisTemplateApi.md#analysistemplategetanalysistemplatesquery) | **Get** /analysistemplates/search | Retrieve analysis templates based on the specified conditions. By default, returns all analysis templates.
*AnalysisTemplateApi* | [**AnalysisTemplateGetByPath**](docs/AnalysisTemplateApi.md#analysistemplategetbypath) | **Get** /analysistemplates | Retrieve an analysis template by path.
*AnalysisTemplateApi* | [**AnalysisTemplateGetCategories**](docs/AnalysisTemplateApi.md#analysistemplategetcategories) | **Get** /analysistemplates/{webId}/categories | Get an analysis template&#39;s categories.
*AnalysisTemplateApi* | [**AnalysisTemplateGetSecurity**](docs/AnalysisTemplateApi.md#analysistemplategetsecurity) | **Get** /analysistemplates/{webId}/security | Get the security information of the specified security item associated with the analysis template for a specified user.
*AnalysisTemplateApi* | [**AnalysisTemplateGetSecurityEntries**](docs/AnalysisTemplateApi.md#analysistemplategetsecurityentries) | **Get** /analysistemplates/{webId}/securityentries | Retrieve the security entries associated with the analysis template based on the specified criteria. By default, all security entries for this analysis template are returned.
*AnalysisTemplateApi* | [**AnalysisTemplateGetSecurityEntryByName**](docs/AnalysisTemplateApi.md#analysistemplategetsecurityentrybyname) | **Get** /analysistemplates/{webId}/securityentries/{name} | Retrieve the security entry associated with the analysis template with the specified name.
*AnalysisTemplateApi* | [**AnalysisTemplateUpdate**](docs/AnalysisTemplateApi.md#analysistemplateupdate) | **Patch** /analysistemplates/{webId} | Update an analysis template by replacing items in its definition.
*AnalysisTemplateApi* | [**AnalysisTemplateUpdateSecurityEntry**](docs/AnalysisTemplateApi.md#analysistemplateupdatesecurityentry) | **Put** /analysistemplates/{webId}/securityentries/{name} | Update a security entry owned by the analysis template.
*AssetDatabaseApi* | [**AssetDatabaseAddReferencedElement**](docs/AssetDatabaseApi.md#assetdatabaseaddreferencedelement) | **Post** /assetdatabases/{webId}/referencedelements | Add a reference to an existing element to the specified database.
*AssetDatabaseApi* | [**AssetDatabaseCreateAnalysisCategory**](docs/AssetDatabaseApi.md#assetdatabasecreateanalysiscategory) | **Post** /assetdatabases/{webId}/analysiscategories | Create an analysis category at the Asset Database root.
*AssetDatabaseApi* | [**AssetDatabaseCreateAnalysisTemplate**](docs/AssetDatabaseApi.md#assetdatabasecreateanalysistemplate) | **Post** /assetdatabases/{webId}/analysistemplates | Create an analysis template at the Asset Database root.
*AssetDatabaseApi* | [**AssetDatabaseCreateAttributeCategory**](docs/AssetDatabaseApi.md#assetdatabasecreateattributecategory) | **Post** /assetdatabases/{webId}/attributecategories | Create an attribute category at the Asset Database root.
*AssetDatabaseApi* | [**AssetDatabaseCreateElement**](docs/AssetDatabaseApi.md#assetdatabasecreateelement) | **Post** /assetdatabases/{webId}/elements | Create a child element.
*AssetDatabaseApi* | [**AssetDatabaseCreateElementCategory**](docs/AssetDatabaseApi.md#assetdatabasecreateelementcategory) | **Post** /assetdatabases/{webId}/elementcategories | Create an element category at the Asset Database root.
*AssetDatabaseApi* | [**AssetDatabaseCreateElementTemplate**](docs/AssetDatabaseApi.md#assetdatabasecreateelementtemplate) | **Post** /assetdatabases/{webId}/elementtemplates | Create a template at the Asset Database root. Specify InstanceType of \&quot;Element\&quot; or \&quot;EventFrame\&quot; to create element or event frame template respectively. Only these two types of templates can be created.
*AssetDatabaseApi* | [**AssetDatabaseCreateEnumerationSet**](docs/AssetDatabaseApi.md#assetdatabasecreateenumerationset) | **Post** /assetdatabases/{webId}/enumerationsets | Create an enumeration set at the Asset Database.
*AssetDatabaseApi* | [**AssetDatabaseCreateEventFrame**](docs/AssetDatabaseApi.md#assetdatabasecreateeventframe) | **Post** /assetdatabases/{webId}/eventframes | Create an event frame.
*AssetDatabaseApi* | [**AssetDatabaseCreateSecurityEntry**](docs/AssetDatabaseApi.md#assetdatabasecreatesecurityentry) | **Post** /assetdatabases/{webId}/securityentries | Create a security entry owned by the asset database.
*AssetDatabaseApi* | [**AssetDatabaseCreateTable**](docs/AssetDatabaseApi.md#assetdatabasecreatetable) | **Post** /assetdatabases/{webId}/tables | Create a table on the Asset Database.
*AssetDatabaseApi* | [**AssetDatabaseCreateTableCategory**](docs/AssetDatabaseApi.md#assetdatabasecreatetablecategory) | **Post** /assetdatabases/{webId}/tablecategories | Create a table category on the Asset Database.
*AssetDatabaseApi* | [**AssetDatabaseDelete**](docs/AssetDatabaseApi.md#assetdatabasedelete) | **Delete** /assetdatabases/{webId} | Delete an asset database.
*AssetDatabaseApi* | [**AssetDatabaseDeleteSecurityEntry**](docs/AssetDatabaseApi.md#assetdatabasedeletesecurityentry) | **Delete** /assetdatabases/{webId}/securityentries/{name} | Delete a security entry owned by the asset database.
*AssetDatabaseApi* | [**AssetDatabaseExport**](docs/AssetDatabaseApi.md#assetdatabaseexport) | **Get** /assetdatabases/{webId}/export | Export the asset database.
*AssetDatabaseApi* | [**AssetDatabaseFindAnalyses**](docs/AssetDatabaseApi.md#assetdatabasefindanalyses) | **Get** /assetdatabases/{webId}/analyses | Retrieve analyses based on the specified conditions.
*AssetDatabaseApi* | [**AssetDatabaseFindElementAttributes**](docs/AssetDatabaseApi.md#assetdatabasefindelementattributes) | **Get** /assetdatabases/{webId}/elementattributes | Retrieves a list of element attributes matching the specified filters from the specified asset database.
*AssetDatabaseApi* | [**AssetDatabaseFindEventFrameAttributes**](docs/AssetDatabaseApi.md#assetdatabasefindeventframeattributes) | **Get** /assetdatabases/{webId}/eventframeattributes | Retrieves a list of event frame attributes matching the specified filters from the specified asset database.
*AssetDatabaseApi* | [**AssetDatabaseGet**](docs/AssetDatabaseApi.md#assetdatabaseget) | **Get** /assetdatabases/{webId} | Retrieve an Asset Database.
*AssetDatabaseApi* | [**AssetDatabaseGetAnalysisCategories**](docs/AssetDatabaseApi.md#assetdatabasegetanalysiscategories) | **Get** /assetdatabases/{webId}/analysiscategories | Retrieve analysis categories for a given Asset Database.
*AssetDatabaseApi* | [**AssetDatabaseGetAnalysisTemplates**](docs/AssetDatabaseApi.md#assetdatabasegetanalysistemplates) | **Get** /assetdatabases/{webId}/analysistemplates | Retrieve analysis templates based on the specified criteria. By default, all analysis templates in the specified Asset Database are returned.
*AssetDatabaseApi* | [**AssetDatabaseGetAttributeCategories**](docs/AssetDatabaseApi.md#assetdatabasegetattributecategories) | **Get** /assetdatabases/{webId}/attributecategories | Retrieve attribute categories for a given Asset Database.
*AssetDatabaseApi* | [**AssetDatabaseGetByPath**](docs/AssetDatabaseApi.md#assetdatabasegetbypath) | **Get** /assetdatabases | Retrieve an Asset Database by path.
*AssetDatabaseApi* | [**AssetDatabaseGetElementCategories**](docs/AssetDatabaseApi.md#assetdatabasegetelementcategories) | **Get** /assetdatabases/{webId}/elementcategories | Retrieve element categories for a given Asset Database.
*AssetDatabaseApi* | [**AssetDatabaseGetElementTemplates**](docs/AssetDatabaseApi.md#assetdatabasegetelementtemplates) | **Get** /assetdatabases/{webId}/elementtemplates | Retrieve element templates based on the specified criteria. Only templates of instance type \&quot;Element\&quot; and \&quot;EventFrame\&quot; are returned. By default, all element and event frame templates in the specified Asset Database are returned.
*AssetDatabaseApi* | [**AssetDatabaseGetElements**](docs/AssetDatabaseApi.md#assetdatabasegetelements) | **Get** /assetdatabases/{webId}/elements | Retrieve elements based on the specified conditions. By default, this method selects immediate children of the specified asset database.
*AssetDatabaseApi* | [**AssetDatabaseGetEnumerationSets**](docs/AssetDatabaseApi.md#assetdatabasegetenumerationsets) | **Get** /assetdatabases/{webId}/enumerationsets | Retrieve enumeration sets for given asset database.
*AssetDatabaseApi* | [**AssetDatabaseGetEventFrames**](docs/AssetDatabaseApi.md#assetdatabasegeteventframes) | **Get** /assetdatabases/{webId}/eventframes | Retrieve event frames based on the specified conditions. By default, returns all children of the specified root resource that have been active in the past 8 hours.
*AssetDatabaseApi* | [**AssetDatabaseGetReferencedElements**](docs/AssetDatabaseApi.md#assetdatabasegetreferencedelements) | **Get** /assetdatabases/{webId}/referencedelements | Retrieve referenced elements based on the specified conditions. By default, this method selects all referenced elements at the root level of the asset database.
*AssetDatabaseApi* | [**AssetDatabaseGetSecurity**](docs/AssetDatabaseApi.md#assetdatabasegetsecurity) | **Get** /assetdatabases/{webId}/security | Get the security information of the specified security item associated with the asset database for a specified user.
*AssetDatabaseApi* | [**AssetDatabaseGetSecurityEntries**](docs/AssetDatabaseApi.md#assetdatabasegetsecurityentries) | **Get** /assetdatabases/{webId}/securityentries | Retrieve the security entries of the specified security item associated with the asset database based on the specified criteria. By default, all security entries for this asset database are returned.
*AssetDatabaseApi* | [**AssetDatabaseGetSecurityEntryByName**](docs/AssetDatabaseApi.md#assetdatabasegetsecurityentrybyname) | **Get** /assetdatabases/{webId}/securityentries/{name} | Retrieve the security entry of the specified security item associated with the asset database with the specified name.
*AssetDatabaseApi* | [**AssetDatabaseGetTableCategories**](docs/AssetDatabaseApi.md#assetdatabasegettablecategories) | **Get** /assetdatabases/{webId}/tablecategories | Retrieve table categories for a given Asset Database.
*AssetDatabaseApi* | [**AssetDatabaseGetTables**](docs/AssetDatabaseApi.md#assetdatabasegettables) | **Get** /assetdatabases/{webId}/tables | Retrieve tables for given Asset Database.
*AssetDatabaseApi* | [**AssetDatabaseImport**](docs/AssetDatabaseApi.md#assetdatabaseimport) | **Post** /assetdatabases/{webId}/import | Import an asset database.
*AssetDatabaseApi* | [**AssetDatabaseRemoveReferencedElement**](docs/AssetDatabaseApi.md#assetdatabaseremovereferencedelement) | **Delete** /assetdatabases/{webId}/referencedelements | Remove a reference to an existing element from the specified database.
*AssetDatabaseApi* | [**AssetDatabaseUpdate**](docs/AssetDatabaseApi.md#assetdatabaseupdate) | **Patch** /assetdatabases/{webId} | Update an asset database by replacing items in its definition.
*AssetDatabaseApi* | [**AssetDatabaseUpdateSecurityEntry**](docs/AssetDatabaseApi.md#assetdatabaseupdatesecurityentry) | **Put** /assetdatabases/{webId}/securityentries/{name} | Update a security entry owned by the asset database.
*AssetServerApi* | [**AssetServerCreateAssetDatabase**](docs/AssetServerApi.md#assetservercreateassetdatabase) | **Post** /assetservers/{webId}/assetdatabases | Create an asset database.
*AssetServerApi* | [**AssetServerCreateSecurityEntry**](docs/AssetServerApi.md#assetservercreatesecurityentry) | **Post** /assetservers/{webId}/securityentries | Create a security entry owned by the asset server.
*AssetServerApi* | [**AssetServerCreateSecurityIdentity**](docs/AssetServerApi.md#assetservercreatesecurityidentity) | **Post** /assetservers/{webId}/securityidentities | Create a security identity.
*AssetServerApi* | [**AssetServerCreateSecurityMapping**](docs/AssetServerApi.md#assetservercreatesecuritymapping) | **Post** /assetservers/{webId}/securitymappings | Create a security mapping.
*AssetServerApi* | [**AssetServerCreateUnitClass**](docs/AssetServerApi.md#assetservercreateunitclass) | **Post** /assetservers/{webId}/unitclasses | Create a unit class in the specified Asset Server.
*AssetServerApi* | [**AssetServerDeleteSecurityEntry**](docs/AssetServerApi.md#assetserverdeletesecurityentry) | **Delete** /assetservers/{webId}/securityentries/{name} | Delete a security entry owned by the asset server.
*AssetServerApi* | [**AssetServerGet**](docs/AssetServerApi.md#assetserverget) | **Get** /assetservers/{webId} | Retrieve an Asset Server.
*AssetServerApi* | [**AssetServerGetAnalysisRulePlugIns**](docs/AssetServerApi.md#assetservergetanalysisruleplugins) | **Get** /assetservers/{webId}/analysisruleplugins | Retrieve a list of all Analysis Rule Plug-in&#39;s.
*AssetServerApi* | [**AssetServerGetByName**](docs/AssetServerApi.md#assetservergetbyname) | **Get** /assetservers#name | Retrieve an Asset Server by name.
*AssetServerApi* | [**AssetServerGetByPath**](docs/AssetServerApi.md#assetservergetbypath) | **Get** /assetservers#path | Retrieve an Asset Server by path.
*AssetServerApi* | [**AssetServerGetDatabases**](docs/AssetServerApi.md#assetservergetdatabases) | **Get** /assetservers/{webId}/assetdatabases | Retrieve a list of all Asset Databases on the specified Asset Server.
*AssetServerApi* | [**AssetServerGetNotificationContactTemplates**](docs/AssetServerApi.md#assetservergetnotificationcontacttemplates) | **Get** /assetservers/{webId}/notificationcontacttemplates | Retrieve a list of all notification contact templates on the specified Asset Server.
*AssetServerApi* | [**AssetServerGetSecurity**](docs/AssetServerApi.md#assetservergetsecurity) | **Get** /assetservers/{webId}/security | Get the security information of the specified security item associated with the asset server for a specified user.
*AssetServerApi* | [**AssetServerGetSecurityEntries**](docs/AssetServerApi.md#assetservergetsecurityentries) | **Get** /assetservers/{webId}/securityentries | Retrieve the security entries of the specified security item associated with the asset server based on the specified criteria. By default, all security entries for this asset server are returned.
*AssetServerApi* | [**AssetServerGetSecurityEntryByName**](docs/AssetServerApi.md#assetservergetsecurityentrybyname) | **Get** /assetservers/{webId}/securityentries/{name} | Retrieve the security entry of the specified security item associated with the asset server with the specified name.
*AssetServerApi* | [**AssetServerGetSecurityIdentities**](docs/AssetServerApi.md#assetservergetsecurityidentities) | **Get** /assetservers/{webId}/securityidentities | Retrieve security identities based on the specified criteria. By default, all security identities in the specified Asset Server are returned.
*AssetServerApi* | [**AssetServerGetSecurityIdentitiesForUser**](docs/AssetServerApi.md#assetservergetsecurityidentitiesforuser) | **Get** /assetservers/{webId}/securityidentities#userIdentity | Retrieve security identities for a specific user.
*AssetServerApi* | [**AssetServerGetSecurityMappings**](docs/AssetServerApi.md#assetservergetsecuritymappings) | **Get** /assetservers/{webId}/securitymappings | Retrieve security mappings based on the specified criteria. By default, all security mappings in the specified Asset Server are returned.
*AssetServerApi* | [**AssetServerGetTimeRulePlugIns**](docs/AssetServerApi.md#assetservergettimeruleplugins) | **Get** /assetservers/{webId}/timeruleplugins | Retrieve a list of all Time Rule Plug-in&#39;s.
*AssetServerApi* | [**AssetServerGetUnitClasses**](docs/AssetServerApi.md#assetservergetunitclasses) | **Get** /assetservers/{webId}/unitclasses | Retrieve a list of all unit classes on the specified Asset Server.
*AssetServerApi* | [**AssetServerList**](docs/AssetServerApi.md#assetserverlist) | **Get** /assetservers | Retrieve a list of all Asset Servers known to this service.
*AssetServerApi* | [**AssetServerUpdateSecurityEntry**](docs/AssetServerApi.md#assetserverupdatesecurityentry) | **Put** /assetservers/{webId}/securityentries/{name} | Update a security entry owned by the asset server.
*AttributeApi* | [**AttributeCreateAttribute**](docs/AttributeApi.md#attributecreateattribute) | **Post** /attributes/{webId}/attributes | Create a new attribute as a child of the specified attribute.
*AttributeApi* | [**AttributeCreateConfig**](docs/AttributeApi.md#attributecreateconfig) | **Post** /attributes/{webId}/config | Create or update an attribute&#39;s DataReference configuration (Create/Update PI point for PI Point DataReference).
*AttributeApi* | [**AttributeDelete**](docs/AttributeApi.md#attributedelete) | **Delete** /attributes/{webId} | Delete an attribute.
*AttributeApi* | [**AttributeGet**](docs/AttributeApi.md#attributeget) | **Get** /attributes/{webId} | Retrieve an attribute.
*AttributeApi* | [**AttributeGetAttributes**](docs/AttributeApi.md#attributegetattributes) | **Get** /attributes/{webId}/attributes | Get the child attributes of the specified attribute.
*AttributeApi* | [**AttributeGetAttributesQuery**](docs/AttributeApi.md#attributegetattributesquery) | **Get** /attributes/search | Retrieve attributes based on the specified conditions. Returns attributes using the specified search query string.
*AttributeApi* | [**AttributeGetByPath**](docs/AttributeApi.md#attributegetbypath) | **Get** /attributes | Retrieve an attribute by path.
*AttributeApi* | [**AttributeGetCategories**](docs/AttributeApi.md#attributegetcategories) | **Get** /attributes/{webId}/categories | Get an attribute&#39;s categories.
*AttributeApi* | [**AttributeGetMultiple**](docs/AttributeApi.md#attributegetmultiple) | **Get** /attributes/multiple | Retrieve multiple attributes by web id or path.
*AttributeApi* | [**AttributeGetValue**](docs/AttributeApi.md#attributegetvalue) | **Get** /attributes/{webId}/value | Get the attribute&#39;s value. This call is intended for use with attributes that have no data reference only. For attributes with a data reference, consult the documentation for Streams.
*AttributeApi* | [**AttributeSetValue**](docs/AttributeApi.md#attributesetvalue) | **Put** /attributes/{webId}/value | Set the value of a configuration item attribute. For attributes with a data reference or non-configuration item attributes, consult the documentation for streams.
*AttributeApi* | [**AttributeUpdate**](docs/AttributeApi.md#attributeupdate) | **Patch** /attributes/{webId} | Update an attribute by replacing items in its definition.
*AttributeCategoryApi* | [**AttributeCategoryCreateSecurityEntry**](docs/AttributeCategoryApi.md#attributecategorycreatesecurityentry) | **Post** /attributecategories/{webId}/securityentries | Create a security entry owned by the attribute category.
*AttributeCategoryApi* | [**AttributeCategoryDelete**](docs/AttributeCategoryApi.md#attributecategorydelete) | **Delete** /attributecategories/{webId} | Delete an attribute category.
*AttributeCategoryApi* | [**AttributeCategoryDeleteSecurityEntry**](docs/AttributeCategoryApi.md#attributecategorydeletesecurityentry) | **Delete** /attributecategories/{webId}/securityentries/{name} | Delete a security entry owned by the attribute category.
*AttributeCategoryApi* | [**AttributeCategoryGet**](docs/AttributeCategoryApi.md#attributecategoryget) | **Get** /attributecategories/{webId} | Retrieve an attribute category.
*AttributeCategoryApi* | [**AttributeCategoryGetByPath**](docs/AttributeCategoryApi.md#attributecategorygetbypath) | **Get** /attributecategories | Retrieve an attribute category by path.
*AttributeCategoryApi* | [**AttributeCategoryGetSecurity**](docs/AttributeCategoryApi.md#attributecategorygetsecurity) | **Get** /attributecategories/{webId}/security | Get the security information of the specified security item associated with the attribute category for a specified user.
*AttributeCategoryApi* | [**AttributeCategoryGetSecurityEntries**](docs/AttributeCategoryApi.md#attributecategorygetsecurityentries) | **Get** /attributecategories/{webId}/securityentries | Retrieve the security entries associated with the attribute category based on the specified criteria. By default, all security entries for this attribute category are returned.
*AttributeCategoryApi* | [**AttributeCategoryGetSecurityEntryByName**](docs/AttributeCategoryApi.md#attributecategorygetsecurityentrybyname) | **Get** /attributecategories/{webId}/securityentries/{name} | Retrieve the security entry associated with the attribute category with the specified name.
*AttributeCategoryApi* | [**AttributeCategoryUpdate**](docs/AttributeCategoryApi.md#attributecategoryupdate) | **Patch** /attributecategories/{webId} | Update an attribute category by replacing items in its definition.
*AttributeCategoryApi* | [**AttributeCategoryUpdateSecurityEntry**](docs/AttributeCategoryApi.md#attributecategoryupdatesecurityentry) | **Put** /attributecategories/{webId}/securityentries/{name} | Update a security entry owned by the attribute category.
*AttributeTemplateApi* | [**AttributeTemplateCreateAttributeTemplate**](docs/AttributeTemplateApi.md#attributetemplatecreateattributetemplate) | **Post** /attributetemplates/{webId}/attributetemplates | Create an attribute template as a child of another attribute template.
*AttributeTemplateApi* | [**AttributeTemplateDelete**](docs/AttributeTemplateApi.md#attributetemplatedelete) | **Delete** /attributetemplates/{webId} | Delete an attribute template.
*AttributeTemplateApi* | [**AttributeTemplateGet**](docs/AttributeTemplateApi.md#attributetemplateget) | **Get** /attributetemplates/{webId} | Retrieve an attribute template.
*AttributeTemplateApi* | [**AttributeTemplateGetAttributeTemplates**](docs/AttributeTemplateApi.md#attributetemplategetattributetemplates) | **Get** /attributetemplates/{webId}/attributetemplates | Retrieve an attribute template&#39;s child attribute templates.
*AttributeTemplateApi* | [**AttributeTemplateGetByPath**](docs/AttributeTemplateApi.md#attributetemplategetbypath) | **Get** /attributetemplates | Retrieve an attribute template by path.
*AttributeTemplateApi* | [**AttributeTemplateGetCategories**](docs/AttributeTemplateApi.md#attributetemplategetcategories) | **Get** /attributetemplates/{webId}/categories | Get an attribute template&#39;s categories.
*AttributeTemplateApi* | [**AttributeTemplateUpdate**](docs/AttributeTemplateApi.md#attributetemplateupdate) | **Patch** /attributetemplates/{webId} | Update an existing attribute template by replacing items in its definition.
*AttributeTraitApi* | [**AttributeTraitGet**](docs/AttributeTraitApi.md#attributetraitget) | **Get** /attributetraits/{name} | Retrieve an attribute trait.
*AttributeTraitApi* | [**AttributeTraitGetByCategory**](docs/AttributeTraitApi.md#attributetraitgetbycategory) | **Get** /attributetraits | Retrieve all attribute traits of the specified category/categories.
*BatchApi* | [**BatchExecute**](docs/BatchApi.md#batchexecute) | **Post** /batch | Execute a batch of requests against the service. As shown in the Sample Request, the input is a dictionary with IDs as keys and request objects as values. Each request object specifies the HTTP method and the resource and, optionally, the content and a list of parent IDs. The list of parent IDs specifies which other requests must complete before the given request will be executed. The example first creates an element, then gets the element by the response&#39;s Location header, then creates an attribute for the element. Note that the resource can be an absolute URL or a JsonPath that references the response to the parent request. The batch&#39;s response is a dictionary uses keys corresponding those provided in the request, with response objects containing a status code, response headers, and the response body. A request can alternatively specify a request template in place of a resource. In this case, a single JsonPath may select multiple tokens, and a separate subrequest will be made from the template for each token. The responses of these subrequests will returned as the content of a single outer response.
*CalculationApi* | [**CalculationGetAtIntervals**](docs/CalculationApi.md#calculationgetatintervals) | **Get** /calculation/intervals | Returns results of evaluating the expression over the time range from the start time to the end time at a defined interval.
*CalculationApi* | [**CalculationGetAtRecorded**](docs/CalculationApi.md#calculationgetatrecorded) | **Get** /calculation/recorded | Returns the result of evaluating the expression at each point in time over the time range from the start time to the end time where a recorded value exists for a member of the expression.
*CalculationApi* | [**CalculationGetAtTimes**](docs/CalculationApi.md#calculationgetattimes) | **Get** /calculation/times | Returns the result of evaluating the expression at the specified timestamps.
*CalculationApi* | [**CalculationGetSummary**](docs/CalculationApi.md#calculationgetsummary) | **Get** /calculation/summary | Returns the result of evaluating the expression over the time range from the start time to the end time. The time range is first divided into a number of summary intervals. Then the calculation is performed for the specified summaries over each interval.
*ChannelApi* | [**ChannelInstances**](docs/ChannelApi.md#channelinstances) | **Get** /channels/instances | Retrieves a list of currently running channel instances.
*ConfigurationApi* | [**ConfigurationDelete**](docs/ConfigurationApi.md#configurationdelete) | **Delete** /system/configuration/{key} | Delete a configuration item.
*ConfigurationApi* | [**ConfigurationGet**](docs/ConfigurationApi.md#configurationget) | **Get** /system/configuration/{key} | Get the value of a configuration item.
*ConfigurationApi* | [**ConfigurationList**](docs/ConfigurationApi.md#configurationlist) | **Get** /system/configuration | Get the current system configuration.
*DataServerApi* | [**DataServerCreateEnumerationSet**](docs/DataServerApi.md#dataservercreateenumerationset) | **Post** /dataservers/{webId}/enumerationsets | Create an enumeration set on the Data Server.
*DataServerApi* | [**DataServerCreatePoint**](docs/DataServerApi.md#dataservercreatepoint) | **Post** /dataservers/{webId}/points | Create a point in the specified Data Server.
*DataServerApi* | [**DataServerGet**](docs/DataServerApi.md#dataserverget) | **Get** /dataservers/{webId} | Retrieve a Data Server.
*DataServerApi* | [**DataServerGetByName**](docs/DataServerApi.md#dataservergetbyname) | **Get** /dataservers#name | Retrieve a Data Server by name.
*DataServerApi* | [**DataServerGetByPath**](docs/DataServerApi.md#dataservergetbypath) | **Get** /dataservers#path | Retrieve a Data Server by path.
*DataServerApi* | [**DataServerGetEnumerationSets**](docs/DataServerApi.md#dataservergetenumerationsets) | **Get** /dataservers/{webId}/enumerationsets | Retrieve enumeration sets for given Data Server.
*DataServerApi* | [**DataServerGetLicense**](docs/DataServerApi.md#dataservergetlicense) | **Get** /dataservers/{webId}/license | Retrieves the specified license for the given Data Server. The fields of the response object are string representations of the numerical values reported by the Data Server, with \&quot;Infinity\&quot; representing a license field with no limit.
*DataServerApi* | [**DataServerGetPoints**](docs/DataServerApi.md#dataservergetpoints) | **Get** /dataservers/{webId}/points | Retrieve a list of points on a specified Data Server.
*DataServerApi* | [**DataServerList**](docs/DataServerApi.md#dataserverlist) | **Get** /dataservers | Retrieve a list of Data Servers known to this service.
*ElementApi* | [**ElementAddReferencedElement**](docs/ElementApi.md#elementaddreferencedelement) | **Post** /elements/{webId}/referencedelements | Add a reference to an existing element to the child elements collection.
*ElementApi* | [**ElementCreateAnalysis**](docs/ElementApi.md#elementcreateanalysis) | **Post** /elements/{webId}/analyses | Create an Analysis.
*ElementApi* | [**ElementCreateAttribute**](docs/ElementApi.md#elementcreateattribute) | **Post** /elements/{webId}/attributes | Create a new attribute of the specified element.
*ElementApi* | [**ElementCreateConfig**](docs/ElementApi.md#elementcreateconfig) | **Post** /elements/{webId}/config | Executes the create configuration function of the data references found within the attributes of the element, and optionally, its children.
*ElementApi* | [**ElementCreateElement**](docs/ElementApi.md#elementcreateelement) | **Post** /elements/{webId}/elements | Create a child element.
*ElementApi* | [**ElementCreateSearchByAttribute**](docs/ElementApi.md#elementcreatesearchbyattribute) | **Post** /elements/searchbyattribute | Create a link for a \&quot;Search Elements By Attribute Value\&quot; operation, whose queries are specified in the request content. The SearchRoot is specified by the Web Id of the root Element. If the SearchRoot is not specified, then the search starts at the Asset Database. ElementTemplate must be provided as the Web ID of the ElementTemplate, which are used to create the Elements. All the attributes in the queries must be defined as AttributeTemplates on the ElementTemplate. An array of attribute value queries are ANDed together to find the desired Element objects. At least one value query must be specified. There are limitations on SearchOperators.
*ElementApi* | [**ElementCreateSecurityEntry**](docs/ElementApi.md#elementcreatesecurityentry) | **Post** /elements/{webId}/securityentries | Create a security entry owned by the element.
*ElementApi* | [**ElementDelete**](docs/ElementApi.md#elementdelete) | **Delete** /elements/{webId} | Delete an element.
*ElementApi* | [**ElementDeleteSecurityEntry**](docs/ElementApi.md#elementdeletesecurityentry) | **Delete** /elements/{webId}/securityentries/{name} | Delete a security entry owned by the element.
*ElementApi* | [**ElementExecuteSearchByAttribute**](docs/ElementApi.md#elementexecutesearchbyattribute) | **Get** /elements/searchbyattribute/{searchId} | Execute a \&quot;Search Elements By Attribute Value\&quot; operation.
*ElementApi* | [**ElementFindElementAttributes**](docs/ElementApi.md#elementfindelementattributes) | **Get** /elements/{webId}/elementattributes | Retrieves a list of element attributes matching the specified filters from the specified element.
*ElementApi* | [**ElementGet**](docs/ElementApi.md#elementget) | **Get** /elements/{webId} | Retrieve an element.
*ElementApi* | [**ElementGetAnalyses**](docs/ElementApi.md#elementgetanalyses) | **Get** /elements/{webId}/analyses | Retrieve analyses based on the specified conditions.
*ElementApi* | [**ElementGetAttributes**](docs/ElementApi.md#elementgetattributes) | **Get** /elements/{webId}/attributes | Get the attributes of the specified element.
*ElementApi* | [**ElementGetByPath**](docs/ElementApi.md#elementgetbypath) | **Get** /elements | Retrieve an element by path.
*ElementApi* | [**ElementGetCategories**](docs/ElementApi.md#elementgetcategories) | **Get** /elements/{webId}/categories | Get an element&#39;s categories.
*ElementApi* | [**ElementGetElements**](docs/ElementApi.md#elementgetelements) | **Get** /elements/{webId}/elements | Retrieve elements based on the specified conditions. By default, this method selects immediate children of the specified element.
*ElementApi* | [**ElementGetElementsQuery**](docs/ElementApi.md#elementgetelementsquery) | **Get** /elements/search | Retrieve elements based on the specified conditions. By default, returns all the elements.
*ElementApi* | [**ElementGetEventFrames**](docs/ElementApi.md#elementgeteventframes) | **Get** /elements/{webId}/eventframes | Retrieve event frames that reference this element based on the specified conditions. By default, returns all event frames that reference this element that have been active in the past 8 hours.
*ElementApi* | [**ElementGetMultiple**](docs/ElementApi.md#elementgetmultiple) | **Get** /elements/multiple | Retrieve multiple elements by web id or path.
*ElementApi* | [**ElementGetNotificationRules**](docs/ElementApi.md#elementgetnotificationrules) | **Get** /elements/{webId}/notificationrules | Retrieve notification rules for an element
*ElementApi* | [**ElementGetPaths**](docs/ElementApi.md#elementgetpaths) | **Get** /elements/{webId}/paths | Get a list of the full or relative paths to this element.
*ElementApi* | [**ElementGetReferencedElements**](docs/ElementApi.md#elementgetreferencedelements) | **Get** /elements/{webId}/referencedelements | Retrieve referenced elements based on the specified conditions. By default, this method selects all referenced elements of the current resource.
*ElementApi* | [**ElementGetSecurity**](docs/ElementApi.md#elementgetsecurity) | **Get** /elements/{webId}/security | Get the security information of the specified security item associated with the element for a specified user.
*ElementApi* | [**ElementGetSecurityEntries**](docs/ElementApi.md#elementgetsecurityentries) | **Get** /elements/{webId}/securityentries | Retrieve the security entries associated with the element based on the specified criteria. By default, all security entries for this element are returned.
*ElementApi* | [**ElementGetSecurityEntryByName**](docs/ElementApi.md#elementgetsecurityentrybyname) | **Get** /elements/{webId}/securityentries/{name} | Retrieve the security entry associated with the element with the specified name.
*ElementApi* | [**ElementRemoveReferencedElement**](docs/ElementApi.md#elementremovereferencedelement) | **Delete** /elements/{webId}/referencedelements | Remove a reference to an existing element from the child elements collection.
*ElementApi* | [**ElementUpdate**](docs/ElementApi.md#elementupdate) | **Patch** /elements/{webId} | Update an element by replacing items in its definition.
*ElementApi* | [**ElementUpdateSecurityEntry**](docs/ElementApi.md#elementupdatesecurityentry) | **Put** /elements/{webId}/securityentries/{name} | Update a security entry owned by the element.
*ElementCategoryApi* | [**ElementCategoryCreateSecurityEntry**](docs/ElementCategoryApi.md#elementcategorycreatesecurityentry) | **Post** /elementcategories/{webId}/securityentries | Create a security entry owned by the element category.
*ElementCategoryApi* | [**ElementCategoryDelete**](docs/ElementCategoryApi.md#elementcategorydelete) | **Delete** /elementcategories/{webId} | Delete an element category.
*ElementCategoryApi* | [**ElementCategoryDeleteSecurityEntry**](docs/ElementCategoryApi.md#elementcategorydeletesecurityentry) | **Delete** /elementcategories/{webId}/securityentries/{name} | Delete a security entry owned by the element category.
*ElementCategoryApi* | [**ElementCategoryGet**](docs/ElementCategoryApi.md#elementcategoryget) | **Get** /elementcategories/{webId} | Retrieve an element category.
*ElementCategoryApi* | [**ElementCategoryGetByPath**](docs/ElementCategoryApi.md#elementcategorygetbypath) | **Get** /elementcategories | Retrieve an element category by path.
*ElementCategoryApi* | [**ElementCategoryGetSecurity**](docs/ElementCategoryApi.md#elementcategorygetsecurity) | **Get** /elementcategories/{webId}/security | Get the security information of the specified security item associated with the element category for a specified user.
*ElementCategoryApi* | [**ElementCategoryGetSecurityEntries**](docs/ElementCategoryApi.md#elementcategorygetsecurityentries) | **Get** /elementcategories/{webId}/securityentries | Retrieve the security entries associated with the element category based on the specified criteria. By default, all security entries for this element category are returned.
*ElementCategoryApi* | [**ElementCategoryGetSecurityEntryByName**](docs/ElementCategoryApi.md#elementcategorygetsecurityentrybyname) | **Get** /elementcategories/{webId}/securityentries/{name} | Retrieve the security entry associated with the element category with the specified name.
*ElementCategoryApi* | [**ElementCategoryUpdate**](docs/ElementCategoryApi.md#elementcategoryupdate) | **Patch** /elementcategories/{webId} | Update an element category by replacing items in its definition.
*ElementCategoryApi* | [**ElementCategoryUpdateSecurityEntry**](docs/ElementCategoryApi.md#elementcategoryupdatesecurityentry) | **Put** /elementcategories/{webId}/securityentries/{name} | Update a security entry owned by the element category.
*ElementTemplateApi* | [**ElementTemplateCreateAttributeTemplate**](docs/ElementTemplateApi.md#elementtemplatecreateattributetemplate) | **Post** /elementtemplates/{webId}/attributetemplates | Create an attribute template.
*ElementTemplateApi* | [**ElementTemplateCreateSecurityEntry**](docs/ElementTemplateApi.md#elementtemplatecreatesecurityentry) | **Post** /elementtemplates/{webId}/securityentries | Create a security entry owned by the element template.
*ElementTemplateApi* | [**ElementTemplateDelete**](docs/ElementTemplateApi.md#elementtemplatedelete) | **Delete** /elementtemplates/{webId} | Delete an element template.
*ElementTemplateApi* | [**ElementTemplateDeleteSecurityEntry**](docs/ElementTemplateApi.md#elementtemplatedeletesecurityentry) | **Delete** /elementtemplates/{webId}/securityentries/{name} | Delete a security entry owned by the element template.
*ElementTemplateApi* | [**ElementTemplateGet**](docs/ElementTemplateApi.md#elementtemplateget) | **Get** /elementtemplates/{webId} | Retrieve an element template.
*ElementTemplateApi* | [**ElementTemplateGetAnalysisTemplates**](docs/ElementTemplateApi.md#elementtemplategetanalysistemplates) | **Get** /elementtemplates/{webId}/analysistemplates | Get analysis templates for an element template.
*ElementTemplateApi* | [**ElementTemplateGetAttributeTemplates**](docs/ElementTemplateApi.md#elementtemplategetattributetemplates) | **Get** /elementtemplates/{webId}/attributetemplates | Get child attribute templates for an element template.
*ElementTemplateApi* | [**ElementTemplateGetBaseElementTemplates**](docs/ElementTemplateApi.md#elementtemplategetbaseelementtemplates) | **Get** /elementtemplates/{webId}/baseelementtemplates | Get base element templates for an element template.
*ElementTemplateApi* | [**ElementTemplateGetByPath**](docs/ElementTemplateApi.md#elementtemplategetbypath) | **Get** /elementtemplates | Retrieve an element template by path.
*ElementTemplateApi* | [**ElementTemplateGetCategories**](docs/ElementTemplateApi.md#elementtemplategetcategories) | **Get** /elementtemplates/{webId}/categories | Get an element template&#39;s categories.
*ElementTemplateApi* | [**ElementTemplateGetDerivedElementTemplates**](docs/ElementTemplateApi.md#elementtemplategetderivedelementtemplates) | **Get** /elementtemplates/{webId}/derivedelementtemplates | Get derived element templates for an element template.
*ElementTemplateApi* | [**ElementTemplateGetNotificationRuleTemplates**](docs/ElementTemplateApi.md#elementtemplategetnotificationruletemplates) | **Get** /elementtemplates/{webId}/notificationruletemplates | Get notification rule templates for an element template
*ElementTemplateApi* | [**ElementTemplateGetSecurity**](docs/ElementTemplateApi.md#elementtemplategetsecurity) | **Get** /elementtemplates/{webId}/security | Get the security information of the specified security item associated with the element template for a specified user.
*ElementTemplateApi* | [**ElementTemplateGetSecurityEntries**](docs/ElementTemplateApi.md#elementtemplategetsecurityentries) | **Get** /elementtemplates/{webId}/securityentries | Retrieve the security entries associated with the element template based on the specified criteria. By default, all security entries for this element template are returned.
*ElementTemplateApi* | [**ElementTemplateGetSecurityEntryByName**](docs/ElementTemplateApi.md#elementtemplategetsecurityentrybyname) | **Get** /elementtemplates/{webId}/securityentries/{name} | Retrieve the security entry associated with the element template with the specified name.
*ElementTemplateApi* | [**ElementTemplateUpdate**](docs/ElementTemplateApi.md#elementtemplateupdate) | **Patch** /elementtemplates/{webId} | Update an element template by replacing items in its definition.
*ElementTemplateApi* | [**ElementTemplateUpdateSecurityEntry**](docs/ElementTemplateApi.md#elementtemplateupdatesecurityentry) | **Put** /elementtemplates/{webId}/securityentries/{name} | Update a security entry owned by the element template.
*EnumerationSetApi* | [**EnumerationSetCreateSecurityEntry**](docs/EnumerationSetApi.md#enumerationsetcreatesecurityentry) | **Post** /enumerationsets/{webId}/securityentries | Create a security entry owned by the enumeration set.
*EnumerationSetApi* | [**EnumerationSetCreateValue**](docs/EnumerationSetApi.md#enumerationsetcreatevalue) | **Post** /enumerationsets/{webId}/enumerationvalues | Create an enumeration value for a enumeration set.
*EnumerationSetApi* | [**EnumerationSetDelete**](docs/EnumerationSetApi.md#enumerationsetdelete) | **Delete** /enumerationsets/{webId} | Delete an enumeration set.
*EnumerationSetApi* | [**EnumerationSetDeleteSecurityEntry**](docs/EnumerationSetApi.md#enumerationsetdeletesecurityentry) | **Delete** /enumerationsets/{webId}/securityentries/{name} | Delete a security entry owned by the enumeration set.
*EnumerationSetApi* | [**EnumerationSetGet**](docs/EnumerationSetApi.md#enumerationsetget) | **Get** /enumerationsets/{webId} | Retrieve an enumeration set.
*EnumerationSetApi* | [**EnumerationSetGetByPath**](docs/EnumerationSetApi.md#enumerationsetgetbypath) | **Get** /enumerationsets | Retrieve an enumeration set by path.
*EnumerationSetApi* | [**EnumerationSetGetSecurity**](docs/EnumerationSetApi.md#enumerationsetgetsecurity) | **Get** /enumerationsets/{webId}/security | Get the security information of the specified security item associated with the enumeration set for a specified user.
*EnumerationSetApi* | [**EnumerationSetGetSecurityEntries**](docs/EnumerationSetApi.md#enumerationsetgetsecurityentries) | **Get** /enumerationsets/{webId}/securityentries | Retrieve the security entries associated with the enumeration set based on the specified criteria. By default, all security entries for this enumeration set are returned.
*EnumerationSetApi* | [**EnumerationSetGetSecurityEntryByName**](docs/EnumerationSetApi.md#enumerationsetgetsecurityentrybyname) | **Get** /enumerationsets/{webId}/securityentries/{name} | Retrieve the security entry associated with the enumeration set with the specified name.
*EnumerationSetApi* | [**EnumerationSetGetValues**](docs/EnumerationSetApi.md#enumerationsetgetvalues) | **Get** /enumerationsets/{webId}/enumerationvalues | Retrieve an enumeration set&#39;s values.
*EnumerationSetApi* | [**EnumerationSetUpdate**](docs/EnumerationSetApi.md#enumerationsetupdate) | **Patch** /enumerationsets/{webId} | Update an enumeration set by replacing items in its definition.
*EnumerationSetApi* | [**EnumerationSetUpdateSecurityEntry**](docs/EnumerationSetApi.md#enumerationsetupdatesecurityentry) | **Put** /enumerationsets/{webId}/securityentries/{name} | Update a security entry owned by the enumeration set.
*EnumerationValueApi* | [**EnumerationValueDeleteEnumerationValue**](docs/EnumerationValueApi.md#enumerationvaluedeleteenumerationvalue) | **Delete** /enumerationvalues/{webId} | Delete an enumeration value from an enumeration set.
*EnumerationValueApi* | [**EnumerationValueGet**](docs/EnumerationValueApi.md#enumerationvalueget) | **Get** /enumerationvalues/{webId} | Retrieve an enumeration value mapping
*EnumerationValueApi* | [**EnumerationValueGetByPath**](docs/EnumerationValueApi.md#enumerationvaluegetbypath) | **Get** /enumerationvalues | Retrieve an enumeration value by path.
*EnumerationValueApi* | [**EnumerationValueUpdateEnumerationValue**](docs/EnumerationValueApi.md#enumerationvalueupdateenumerationvalue) | **Patch** /enumerationvalues/{webId} | Update an enumeration value by replacing items in its definition.
*EventFrameApi* | [**EventFrameAcknowledge**](docs/EventFrameApi.md#eventframeacknowledge) | **Patch** /eventframes/{webId}/acknowledge | Calls the EventFrame&#39;s Acknowledge method.
*EventFrameApi* | [**EventFrameCaptureValues**](docs/EventFrameApi.md#eventframecapturevalues) | **Post** /eventframes/{webId}/attributes/capture | Calls the EventFrame&#39;s CaptureValues method.
*EventFrameApi* | [**EventFrameCreateAnnotation**](docs/EventFrameApi.md#eventframecreateannotation) | **Post** /eventframes/{webId}/annotations | Create an annotation on an event frame.
*EventFrameApi* | [**EventFrameCreateAttribute**](docs/EventFrameApi.md#eventframecreateattribute) | **Post** /eventframes/{webId}/attributes | Create a new attribute of the specified event frame.
*EventFrameApi* | [**EventFrameCreateConfig**](docs/EventFrameApi.md#eventframecreateconfig) | **Post** /eventframes/{webId}/config | Executes the create configuration function of the data references found within the attributes of the event frame, and optionally, its children.
*EventFrameApi* | [**EventFrameCreateEventFrame**](docs/EventFrameApi.md#eventframecreateeventframe) | **Post** /eventframes/{webId}/eventframes | Create an event frame as a child of the specified event frame.
*EventFrameApi* | [**EventFrameCreateSearchByAttribute**](docs/EventFrameApi.md#eventframecreatesearchbyattribute) | **Post** /eventframes/searchbyattribute | Create a link for a \&quot;Search EventFrames By Attribute Value\&quot; operation, whose queries are specified in the request content. The SearchRoot is specified by the Web Id of the root EventFrame. If the SearchRoot is not specified, then the search starts at the Asset Database. ElementTemplate must be provided as the Web ID of the ElementTemplate, which are used to create the EventFrames. All the attributes in the queries must be defined as AttributeTemplates on the ElementTemplate. An array of attribute value queries are ANDed together to find the desired Element objects. At least one value query must be specified. There are limitations on SearchOperators.
*EventFrameApi* | [**EventFrameCreateSecurityEntry**](docs/EventFrameApi.md#eventframecreatesecurityentry) | **Post** /eventframes/{webId}/securityentries | Create a security entry owned by the event frame.
*EventFrameApi* | [**EventFrameDelete**](docs/EventFrameApi.md#eventframedelete) | **Delete** /eventframes/{webId} | Delete an event frame.
*EventFrameApi* | [**EventFrameDeleteAnnotation**](docs/EventFrameApi.md#eventframedeleteannotation) | **Delete** /eventframes/{webId}/annotations/{id} | Delete an annotation on an event frame. If the annotation has attached media, the attached media will also be deleted.
*EventFrameApi* | [**EventFrameDeleteAnnotationAttachmentMediaById**](docs/EventFrameApi.md#eventframedeleteannotationattachmentmediabyid) | **Delete** /eventframes/{webId}/annotations/{id}/attachment/media | Delete attached media from an annotation on an event frame.
*EventFrameApi* | [**EventFrameDeleteSecurityEntry**](docs/EventFrameApi.md#eventframedeletesecurityentry) | **Delete** /eventframes/{webId}/securityentries/{name} | Delete a security entry owned by the event frame.
*EventFrameApi* | [**EventFrameExecuteSearchByAttribute**](docs/EventFrameApi.md#eventframeexecutesearchbyattribute) | **Get** /eventframes/searchbyattribute/{searchId} | Execute a \&quot;Search EventFrames By Attribute Value\&quot; operation.
*EventFrameApi* | [**EventFrameFindEventFrameAttributes**](docs/EventFrameApi.md#eventframefindeventframeattributes) | **Get** /eventframes/{webId}/eventframeattributes | Retrieves a list of event frame attributes matching the specified filters from the specified event frame.
*EventFrameApi* | [**EventFrameGet**](docs/EventFrameApi.md#eventframeget) | **Get** /eventframes/{webId} | Retrieve an event frame.
*EventFrameApi* | [**EventFrameGetAnnotationAttachmentMediaMetadataById**](docs/EventFrameApi.md#eventframegetannotationattachmentmediametadatabyid) | **Get** /eventframes/{webId}/annotations/{id}/attachment/media/metadata | Gets the metadata of the media attached to the specified annotation.
*EventFrameApi* | [**EventFrameGetAnnotationById**](docs/EventFrameApi.md#eventframegetannotationbyid) | **Get** /eventframes/{webId}/annotations/{id} | Get a specific annotation on an event frame.
*EventFrameApi* | [**EventFrameGetAnnotations**](docs/EventFrameApi.md#eventframegetannotations) | **Get** /eventframes/{webId}/annotations | Get an event frame&#39;s annotations.
*EventFrameApi* | [**EventFrameGetAttributes**](docs/EventFrameApi.md#eventframegetattributes) | **Get** /eventframes/{webId}/attributes | Get the attributes of the specified event frame.
*EventFrameApi* | [**EventFrameGetByPath**](docs/EventFrameApi.md#eventframegetbypath) | **Get** /eventframes | Retrieve an event frame by path.
*EventFrameApi* | [**EventFrameGetCategories**](docs/EventFrameApi.md#eventframegetcategories) | **Get** /eventframes/{webId}/categories | Get an event frame&#39;s categories.
*EventFrameApi* | [**EventFrameGetEventFrames**](docs/EventFrameApi.md#eventframegeteventframes) | **Get** /eventframes/{webId}/eventframes | Retrieve event frames based on the specified conditions. By default, returns all children of the specified root event frame that have been active in the past 8 hours.
*EventFrameApi* | [**EventFrameGetEventFramesQuery**](docs/EventFrameApi.md#eventframegeteventframesquery) | **Get** /eventframes/search | Retrieve event frames based on the specified conditions. Returns event frames using the specified search query string.
*EventFrameApi* | [**EventFrameGetMultiple**](docs/EventFrameApi.md#eventframegetmultiple) | **Get** /eventframes/multiple | Retrieve multiple event frames by web ids or paths.
*EventFrameApi* | [**EventFrameGetReferencedElements**](docs/EventFrameApi.md#eventframegetreferencedelements) | **Get** /eventframes/{webId}/referencedelements | Retrieve the event frame&#39;s referenced elements.
*EventFrameApi* | [**EventFrameGetSecurity**](docs/EventFrameApi.md#eventframegetsecurity) | **Get** /eventframes/{webId}/security | Get the security information of the specified security item associated with the event frame for a specified user.
*EventFrameApi* | [**EventFrameGetSecurityEntries**](docs/EventFrameApi.md#eventframegetsecurityentries) | **Get** /eventframes/{webId}/securityentries | Retrieve the security entries associated with the event frame based on the specified criteria. By default, all security entries for this event frame are returned.
*EventFrameApi* | [**EventFrameGetSecurityEntryByName**](docs/EventFrameApi.md#eventframegetsecurityentrybyname) | **Get** /eventframes/{webId}/securityentries/{name} | Retrieve the security entry associated with the event frame with the specified name.
*EventFrameApi* | [**EventFrameUpdate**](docs/EventFrameApi.md#eventframeupdate) | **Patch** /eventframes/{webId} | Update an event frame by replacing items in its definition.
*EventFrameApi* | [**EventFrameUpdateAnnotation**](docs/EventFrameApi.md#eventframeupdateannotation) | **Patch** /eventframes/{webId}/annotations/{id} | Update an annotation on an event frame by replacing items in its definition.
*EventFrameApi* | [**EventFrameUpdateSecurityEntry**](docs/EventFrameApi.md#eventframeupdatesecurityentry) | **Put** /eventframes/{webId}/securityentries/{name} | Update a security entry owned by the event frame.
*HomeApi* | [**HomeGet**](docs/HomeApi.md#homeget) | **Get** / | Get top level links for this PI System Web API instance.
*NotificationContactTemplateApi* | [**NotificationContactTemplateGet**](docs/NotificationContactTemplateApi.md#notificationcontacttemplateget) | **Get** /notificationcontacttemplates/{webId} | Retrieve a notification contact template.
*NotificationRuleApi* | [**NotificationRuleGetNotificationRuleSubscribers**](docs/NotificationRuleApi.md#notificationrulegetnotificationrulesubscribers) | **Get** /notificationrules/{webId}/notificationrulesubscribers | Retrieve notification rule subscribers.
*NotificationRuleApi* | [**NotificationRuleGetNotificationRules**](docs/NotificationRuleApi.md#notificationrulegetnotificationrules) | **Get** /notificationrules/{webId} | Retrieve a notification rule.
*NotificationRuleApi* | [**NotificationRuleGetNotificationRulesQuery**](docs/NotificationRuleApi.md#notificationrulegetnotificationrulesquery) | **Get** /notificationrules/search | Retrieve notification rules based on the specified conditions. Returns notification rules using the specified search query string.
*NotificationRuleSubscriberApi* | [**NotificationRuleSubscriberGetNotificationRuleSubscriber**](docs/NotificationRuleSubscriberApi.md#notificationrulesubscribergetnotificationrulesubscriber) | **Get** /notificationrulesubscribers/{webId} | Retrieve a notification rule subscriber.
*NotificationRuleSubscriberApi* | [**NotificationRuleSubscriberGetNotificationRuleSubscriberByPath**](docs/NotificationRuleSubscriberApi.md#notificationrulesubscribergetnotificationrulesubscriberbypath) | **Get** /notificationrulesubscribers | Retrieve a notification rule subscriber by path.
*NotificationRuleSubscriberApi* | [**NotificationRuleSubscriberGetNotificationRuleSubscribers**](docs/NotificationRuleSubscriberApi.md#notificationrulesubscribergetnotificationrulesubscribers) | **Get** /notificationrulesubscribers/{webId}/notificationrulesubscribers | Retrieve notification rule subscriber subscribers.
*NotificationRuleTemplateApi* | [**NotificationRuleTemplateGet**](docs/NotificationRuleTemplateApi.md#notificationruletemplateget) | **Get** /notificationruletemplates/{webId} | Get the specified notification rule template.
*NotificationRuleTemplateApi* | [**NotificationRuleTemplateGetNotificationRuleTemplateSubscribers**](docs/NotificationRuleTemplateApi.md#notificationruletemplategetnotificationruletemplatesubscribers) | **Get** /notificationruletemplates/{webId}/notificationrulesubscribers | Retrieve notification rule template subscribers.
*NotificationRuleTemplateApi* | [**NotificationRuleTemplateGetNotificationRuleTemplatesQuery**](docs/NotificationRuleTemplateApi.md#notificationruletemplategetnotificationruletemplatesquery) | **Get** /notificationruletemplates/search | Retrieve Notification rule templates based on the specified conditions. Returns Notification rule templates using the specified search query string.
*PointApi* | [**PointDelete**](docs/PointApi.md#pointdelete) | **Delete** /points/{webId} | Delete a point.
*PointApi* | [**PointGet**](docs/PointApi.md#pointget) | **Get** /points/{webId} | Get a point.
*PointApi* | [**PointGetAttributeByName**](docs/PointApi.md#pointgetattributebyname) | **Get** /points/{webId}/attributes/{name} | Get a point attribute by name.
*PointApi* | [**PointGetAttributes**](docs/PointApi.md#pointgetattributes) | **Get** /points/{webId}/attributes | Get point attributes.
*PointApi* | [**PointGetByPath**](docs/PointApi.md#pointgetbypath) | **Get** /points | Get a point by path.
*PointApi* | [**PointGetMultiple**](docs/PointApi.md#pointgetmultiple) | **Get** /points/multiple | Retrieve multiple points by web id or path.
*PointApi* | [**PointUpdate**](docs/PointApi.md#pointupdate) | **Patch** /points/{webId} | Update a point.
*SecurityIdentityApi* | [**SecurityIdentityDelete**](docs/SecurityIdentityApi.md#securityidentitydelete) | **Delete** /securityidentities/{webId} | Delete a security identity.
*SecurityIdentityApi* | [**SecurityIdentityGet**](docs/SecurityIdentityApi.md#securityidentityget) | **Get** /securityidentities/{webId} | Retrieve a security identity.
*SecurityIdentityApi* | [**SecurityIdentityGetByPath**](docs/SecurityIdentityApi.md#securityidentitygetbypath) | **Get** /securityidentities | Retrieve a security identity by path.
*SecurityIdentityApi* | [**SecurityIdentityGetSecurity**](docs/SecurityIdentityApi.md#securityidentitygetsecurity) | **Get** /securityidentities/{webId}/security | Get the security information of the specified security item associated with the security identity for a specified user.
*SecurityIdentityApi* | [**SecurityIdentityGetSecurityEntries**](docs/SecurityIdentityApi.md#securityidentitygetsecurityentries) | **Get** /securityidentities/{webId}/securityentries | Retrieve the security entries associated with the security identity based on the specified criteria. By default, all security entries for this security identity are returned.
*SecurityIdentityApi* | [**SecurityIdentityGetSecurityEntryByName**](docs/SecurityIdentityApi.md#securityidentitygetsecurityentrybyname) | **Get** /securityidentities/{webId}/securityentries/{name} | Retrieve the security entry associated with the security identity with the specified name.
*SecurityIdentityApi* | [**SecurityIdentityGetSecurityMappings**](docs/SecurityIdentityApi.md#securityidentitygetsecuritymappings) | **Get** /securityidentities/{webId}/securitymappings | Get security mappings for the specified security identity.
*SecurityIdentityApi* | [**SecurityIdentityUpdate**](docs/SecurityIdentityApi.md#securityidentityupdate) | **Patch** /securityidentities/{webId} | Update a security identity by replacing items in its definition.
*SecurityMappingApi* | [**SecurityMappingDelete**](docs/SecurityMappingApi.md#securitymappingdelete) | **Delete** /securitymappings/{webId} | Delete a security mapping.
*SecurityMappingApi* | [**SecurityMappingGet**](docs/SecurityMappingApi.md#securitymappingget) | **Get** /securitymappings/{webId} | Retrieve a security mapping.
*SecurityMappingApi* | [**SecurityMappingGetByPath**](docs/SecurityMappingApi.md#securitymappinggetbypath) | **Get** /securitymappings | Retrieve a security mapping by path.
*SecurityMappingApi* | [**SecurityMappingGetSecurity**](docs/SecurityMappingApi.md#securitymappinggetsecurity) | **Get** /securitymappings/{webId}/security | Get the security information of the specified security item associated with the security mapping for a specified user.
*SecurityMappingApi* | [**SecurityMappingGetSecurityEntries**](docs/SecurityMappingApi.md#securitymappinggetsecurityentries) | **Get** /securitymappings/{webId}/securityentries | Retrieve the security entries associated with the security mapping based on the specified criteria. By default, all security entries for this security mapping are returned.
*SecurityMappingApi* | [**SecurityMappingGetSecurityEntryByName**](docs/SecurityMappingApi.md#securitymappinggetsecurityentrybyname) | **Get** /securitymappings/{webId}/securityentries/{name} | Retrieve the security entry associated with the security mapping with the specified name.
*SecurityMappingApi* | [**SecurityMappingUpdate**](docs/SecurityMappingApi.md#securitymappingupdate) | **Patch** /securitymappings/{webId} | Update a security mapping by replacing items in its definition.
*StreamApi* | [**StreamGetChannel**](docs/StreamApi.md#streamgetchannel) | **Get** /streams/{webId}/channel | Opens a channel that will send messages about any value changes for the specified stream.
*StreamApi* | [**StreamGetEnd**](docs/StreamApi.md#streamgetend) | **Get** /streams/{webId}/end | Returns the end-of-stream value of the stream.
*StreamApi* | [**StreamGetInterpolated**](docs/StreamApi.md#streamgetinterpolated) | **Get** /streams/{webId}/interpolated | Retrieves interpolated values over the specified time range at the specified sampling interval.
*StreamApi* | [**StreamGetInterpolatedAtTimes**](docs/StreamApi.md#streamgetinterpolatedattimes) | **Get** /streams/{webId}/interpolatedattimes | Retrieves interpolated values over the specified time range at the specified sampling interval.
*StreamApi* | [**StreamGetPlot**](docs/StreamApi.md#streamgetplot) | **Get** /streams/{webId}/plot | Retrieves values over the specified time range suitable for plotting over the number of intervals (typically represents pixels).
*StreamApi* | [**StreamGetRecorded**](docs/StreamApi.md#streamgetrecorded) | **Get** /streams/{webId}/recorded | Returns a list of compressed values for the requested time range from the source provider.
*StreamApi* | [**StreamGetRecordedAtTime**](docs/StreamApi.md#streamgetrecordedattime) | **Get** /streams/{webId}/recordedattime | Returns a single recorded value based on the passed time and retrieval mode from the stream.
*StreamApi* | [**StreamGetRecordedAtTimes**](docs/StreamApi.md#streamgetrecordedattimes) | **Get** /streams/{webId}/recordedattimes | Retrieves recorded values at the specified times.
*StreamApi* | [**StreamGetSummary**](docs/StreamApi.md#streamgetsummary) | **Get** /streams/{webId}/summary | Returns a summary over the specified time range for the stream.
*StreamApi* | [**StreamGetValue**](docs/StreamApi.md#streamgetvalue) | **Get** /streams/{webId}/value | Returns the value of the stream at the specified time. By default, this is usually the current value.
*StreamApi* | [**StreamUpdateValue**](docs/StreamApi.md#streamupdatevalue) | **Post** /streams/{webId}/value | Updates a value for the specified stream.
*StreamApi* | [**StreamUpdateValues**](docs/StreamApi.md#streamupdatevalues) | **Post** /streams/{webId}/recorded | Updates multiple values for the specified stream.
*StreamSetApi* | [**StreamSetGetChannel**](docs/StreamSetApi.md#streamsetgetchannel) | **Get** /streamsets/{webId}/channel | Opens a channel that will send messages about any value changes for the attributes of an Element, Event Frame, or Attribute.
*StreamSetApi* | [**StreamSetGetChannelAdHoc**](docs/StreamSetApi.md#streamsetgetchanneladhoc) | **Get** /streamsets/channel | Opens a channel that will send messages about any value changes for the specified streams.
*StreamSetApi* | [**StreamSetGetEnd**](docs/StreamSetApi.md#streamsetgetend) | **Get** /streamsets/{webId}/end | Returns End of stream values of the attributes for an Element, Event Frame or Attribute
*StreamSetApi* | [**StreamSetGetEndAdHoc**](docs/StreamSetApi.md#streamsetgetendadhoc) | **Get** /streamsets/end | Returns End Of Stream values for attributes of the specified streams
*StreamSetApi* | [**StreamSetGetInterpolated**](docs/StreamSetApi.md#streamsetgetinterpolated) | **Get** /streamsets/{webId}/interpolated | Returns interpolated values of attributes for an element, event frame or attribute over the specified time range at the specified sampling interval.
*StreamSetApi* | [**StreamSetGetInterpolatedAdHoc**](docs/StreamSetApi.md#streamsetgetinterpolatedadhoc) | **Get** /streamsets/interpolated | Returns interpolated values of the specified streams over the specified time range at the specified sampling interval.
*StreamSetApi* | [**StreamSetGetInterpolatedAtTimes**](docs/StreamSetApi.md#streamsetgetinterpolatedattimes) | **Get** /streamsets/{webId}/interpolatedattimes | Returns interpolated values of attributes for an element, event frame or attribute at the specified times.
*StreamSetApi* | [**StreamSetGetInterpolatedAtTimesAdHoc**](docs/StreamSetApi.md#streamsetgetinterpolatedattimesadhoc) | **Get** /streamsets/interpolatedattimes | Returns interpolated values of the specified streams at the specified times.
*StreamSetApi* | [**StreamSetGetPlot**](docs/StreamSetApi.md#streamsetgetplot) | **Get** /streamsets/{webId}/plot | Returns values of attributes for an element, event frame or attribute over the specified time range suitable for plotting over the number of intervals (typically represents pixels).
*StreamSetApi* | [**StreamSetGetPlotAdHoc**](docs/StreamSetApi.md#streamsetgetplotadhoc) | **Get** /streamsets/plot | Returns values of attributes for the specified streams over the specified time range suitable for plotting over the number of intervals (typically represents pixels).
*StreamSetApi* | [**StreamSetGetRecorded**](docs/StreamSetApi.md#streamsetgetrecorded) | **Get** /streamsets/{webId}/recorded | Returns recorded values of the attributes for an element, event frame, or attribute.
*StreamSetApi* | [**StreamSetGetRecordedAdHoc**](docs/StreamSetApi.md#streamsetgetrecordedadhoc) | **Get** /streamsets/recorded | Returns recorded values of the specified streams.
*StreamSetApi* | [**StreamSetGetRecordedAtTime**](docs/StreamSetApi.md#streamsetgetrecordedattime) | **Get** /streamsets/{webId}/recordedattime | Returns recorded values of the attributes for an element, event frame, or attribute.
*StreamSetApi* | [**StreamSetGetRecordedAtTimeAdHoc**](docs/StreamSetApi.md#streamsetgetrecordedattimeadhoc) | **Get** /streamsets/recordedattime | Returns recorded values based on the passed time and retrieval mode.
*StreamSetApi* | [**StreamSetGetRecordedAtTimes**](docs/StreamSetApi.md#streamsetgetrecordedattimes) | **Get** /streamsets/{webId}/recordedattimes | Returns recorded values of attributes for an element, event frame or attribute at the specified times.
*StreamSetApi* | [**StreamSetGetRecordedAtTimesAdHoc**](docs/StreamSetApi.md#streamsetgetrecordedattimesadhoc) | **Get** /streamsets/recordedattimes | Returns recorded values of the specified streams at the specified times.
*StreamSetApi* | [**StreamSetGetSummaries**](docs/StreamSetApi.md#streamsetgetsummaries) | **Get** /streamsets/{webId}/summary | Returns summary values of the attributes for an element, event frame or attribute.
*StreamSetApi* | [**StreamSetGetSummariesAdHoc**](docs/StreamSetApi.md#streamsetgetsummariesadhoc) | **Get** /streamsets/summary | Returns summary values of the specified streams.
*StreamSetApi* | [**StreamSetGetValues**](docs/StreamSetApi.md#streamsetgetvalues) | **Get** /streamsets/{webId}/value | Returns values of the attributes for an Element, Event Frame or Attribute at the specified time.
*StreamSetApi* | [**StreamSetGetValuesAdHoc**](docs/StreamSetApi.md#streamsetgetvaluesadhoc) | **Get** /streamsets/value | Returns values of the specified streams.
*StreamSetApi* | [**StreamSetUpdateValue**](docs/StreamSetApi.md#streamsetupdatevalue) | **Post** /streamsets/{webId}/value | Updates a single value for the specified streams.
*StreamSetApi* | [**StreamSetUpdateValueAdHoc**](docs/StreamSetApi.md#streamsetupdatevalueadhoc) | **Post** /streamsets/value | Updates a single value for the specified streams.
*StreamSetApi* | [**StreamSetUpdateValues**](docs/StreamSetApi.md#streamsetupdatevalues) | **Post** /streamsets/{webId}/recorded | Updates multiple values for the specified streams.
*StreamSetApi* | [**StreamSetUpdateValuesAdHoc**](docs/StreamSetApi.md#streamsetupdatevaluesadhoc) | **Post** /streamsets/recorded | Updates multiple values for the specified streams.
*SystemApi* | [**SystemCacheInstances**](docs/SystemApi.md#systemcacheinstances) | **Get** /system/cacheinstances | Get AF cache instances currently in use by the system. These are caches from which user requests are serviced. The number of instances depends on the number of users connected to the service, the service&#39;s authentication method, and the cache instance configuration.
*SystemApi* | [**SystemLanding**](docs/SystemApi.md#systemlanding) | **Get** /system | Get system links for this PI System Web API instance.
*SystemApi* | [**SystemStatus**](docs/SystemApi.md#systemstatus) | **Get** /system/status | Get information about this PI Web API instance. Examples of information returned include the system uptime, the number of cache instances for this PI System Web API instance, and the system run state.
*SystemApi* | [**SystemUserInfo**](docs/SystemApi.md#systemuserinfo) | **Get** /system/userinfo | Get information about the Windows identity used to fulfill the request. This depends on the service&#39;s authentication method and the credentials passed by the client. The impersonation level of the Windows identity is included.
*SystemApi* | [**SystemVersions**](docs/SystemApi.md#systemversions) | **Get** /system/versions | Get the current versions of the PI Web API instance and all external plugins.
*TableApi* | [**TableCreateSecurityEntry**](docs/TableApi.md#tablecreatesecurityentry) | **Post** /tables/{webId}/securityentries | Create a security entry owned by the table.
*TableApi* | [**TableDelete**](docs/TableApi.md#tabledelete) | **Delete** /tables/{webId} | Delete a table.
*TableApi* | [**TableDeleteSecurityEntry**](docs/TableApi.md#tabledeletesecurityentry) | **Delete** /tables/{webId}/securityentries/{name} | Delete a security entry owned by the table.
*TableApi* | [**TableGet**](docs/TableApi.md#tableget) | **Get** /tables/{webId} | Retrieve a table.
*TableApi* | [**TableGetByPath**](docs/TableApi.md#tablegetbypath) | **Get** /tables | Retrieve a table by path.
*TableApi* | [**TableGetCategories**](docs/TableApi.md#tablegetcategories) | **Get** /tables/{webId}/categories | Get a table&#39;s categories.
*TableApi* | [**TableGetData**](docs/TableApi.md#tablegetdata) | **Get** /tables/{webId}/data | Get the table&#39;s data.
*TableApi* | [**TableGetSecurity**](docs/TableApi.md#tablegetsecurity) | **Get** /tables/{webId}/security | Get the security information of the specified security item associated with the table for a specified user.
*TableApi* | [**TableGetSecurityEntries**](docs/TableApi.md#tablegetsecurityentries) | **Get** /tables/{webId}/securityentries | Retrieve the security entries associated with the table based on the specified criteria. By default, all security entries for this table are returned.
*TableApi* | [**TableGetSecurityEntryByName**](docs/TableApi.md#tablegetsecurityentrybyname) | **Get** /tables/{webId}/securityentries/{name} | Retrieve the security entry associated with the table with the specified name.
*TableApi* | [**TableUpdate**](docs/TableApi.md#tableupdate) | **Patch** /tables/{webId} | Update a table by replacing items in its definition.
*TableApi* | [**TableUpdateData**](docs/TableApi.md#tableupdatedata) | **Put** /tables/{webId}/data | Update the table&#39;s data.
*TableApi* | [**TableUpdateSecurityEntry**](docs/TableApi.md#tableupdatesecurityentry) | **Put** /tables/{webId}/securityentries/{name} | Update a security entry owned by the table.
*TableCategoryApi* | [**TableCategoryCreateSecurityEntry**](docs/TableCategoryApi.md#tablecategorycreatesecurityentry) | **Post** /tablecategories/{webId}/securityentries | Create a security entry owned by the table category.
*TableCategoryApi* | [**TableCategoryDelete**](docs/TableCategoryApi.md#tablecategorydelete) | **Delete** /tablecategories/{webId} | Delete a table category.
*TableCategoryApi* | [**TableCategoryDeleteSecurityEntry**](docs/TableCategoryApi.md#tablecategorydeletesecurityentry) | **Delete** /tablecategories/{webId}/securityentries/{name} | Delete a security entry owned by the table category.
*TableCategoryApi* | [**TableCategoryGet**](docs/TableCategoryApi.md#tablecategoryget) | **Get** /tablecategories/{webId} | Retrieve a table category.
*TableCategoryApi* | [**TableCategoryGetByPath**](docs/TableCategoryApi.md#tablecategorygetbypath) | **Get** /tablecategories | Retrieve a table category by path.
*TableCategoryApi* | [**TableCategoryGetSecurity**](docs/TableCategoryApi.md#tablecategorygetsecurity) | **Get** /tablecategories/{webId}/security | Get the security information of the specified security item associated with the table category for a specified user.
*TableCategoryApi* | [**TableCategoryGetSecurityEntries**](docs/TableCategoryApi.md#tablecategorygetsecurityentries) | **Get** /tablecategories/{webId}/securityentries | Retrieve the security entries associated with the table category based on the specified criteria. By default, all security entries for this table category are returned.
*TableCategoryApi* | [**TableCategoryGetSecurityEntryByName**](docs/TableCategoryApi.md#tablecategorygetsecurityentrybyname) | **Get** /tablecategories/{webId}/securityentries/{name} | Retrieve the security entry associated with the table category with the specified name.
*TableCategoryApi* | [**TableCategoryUpdate**](docs/TableCategoryApi.md#tablecategoryupdate) | **Patch** /tablecategories/{webId} | Update a table category by replacing items in its definition.
*TableCategoryApi* | [**TableCategoryUpdateSecurityEntry**](docs/TableCategoryApi.md#tablecategoryupdatesecurityentry) | **Put** /tablecategories/{webId}/securityentries/{name} | Update a security entry owned by the table category.
*TimeRuleApi* | [**TimeRuleDelete**](docs/TimeRuleApi.md#timeruledelete) | **Delete** /timerules/{webId} | Delete a Time Rule.
*TimeRuleApi* | [**TimeRuleGet**](docs/TimeRuleApi.md#timeruleget) | **Get** /timerules/{webId} | Retrieve a Time Rule.
*TimeRuleApi* | [**TimeRuleGetByPath**](docs/TimeRuleApi.md#timerulegetbypath) | **Get** /timerules | Retrieve a Time Rule by path.
*TimeRuleApi* | [**TimeRuleUpdate**](docs/TimeRuleApi.md#timeruleupdate) | **Patch** /timerules/{webId} | Update a Time Rule by replacing items in its definition.
*TimeRulePlugInApi* | [**TimeRulePlugInGet**](docs/TimeRulePlugInApi.md#timerulepluginget) | **Get** /timeruleplugins/{webId} | Retrieve a Time Rule Plug-in.
*TimeRulePlugInApi* | [**TimeRulePlugInGetByPath**](docs/TimeRulePlugInApi.md#timeruleplugingetbypath) | **Get** /timeruleplugins | Retrieve a Time Rule Plug-in by path.
*UnitApi* | [**UnitDelete**](docs/UnitApi.md#unitdelete) | **Delete** /units/{webId} | Delete a unit.
*UnitApi* | [**UnitGet**](docs/UnitApi.md#unitget) | **Get** /units/{webId} | Retrieve a unit.
*UnitApi* | [**UnitGetByPath**](docs/UnitApi.md#unitgetbypath) | **Get** /units | Retrieve a unit by path.
*UnitApi* | [**UnitUpdate**](docs/UnitApi.md#unitupdate) | **Patch** /units/{webId} | Update a unit.
*UnitClassApi* | [**UnitClassCreateUnit**](docs/UnitClassApi.md#unitclasscreateunit) | **Post** /unitclasses/{webId}/units | Create a unit in the specified Unit Class.
*UnitClassApi* | [**UnitClassDelete**](docs/UnitClassApi.md#unitclassdelete) | **Delete** /unitclasses/{webId} | Delete a unit class.
*UnitClassApi* | [**UnitClassGet**](docs/UnitClassApi.md#unitclassget) | **Get** /unitclasses/{webId} | Retrieve a unit class.
*UnitClassApi* | [**UnitClassGetByPath**](docs/UnitClassApi.md#unitclassgetbypath) | **Get** /unitclasses | Retrieve a unit class by path.
*UnitClassApi* | [**UnitClassGetCanonicalUnit**](docs/UnitClassApi.md#unitclassgetcanonicalunit) | **Get** /unitclasses/{webId}/canonicalunit | Get the canonical unit of a unit class.
*UnitClassApi* | [**UnitClassGetUnits**](docs/UnitClassApi.md#unitclassgetunits) | **Get** /unitclasses/{webId}/units | Get a list of all units belonging to the unit class.
*UnitClassApi* | [**UnitClassUpdate**](docs/UnitClassApi.md#unitclassupdate) | **Patch** /unitclasses/{webId} | Update a unit class.


## Documentation For Models

 - [Ambiguous](docs/Ambiguous.md)
 - [Analysis](docs/Analysis.md)
 - [AnalysisCategory](docs/AnalysisCategory.md)
 - [AnalysisCategoryLinks](docs/AnalysisCategoryLinks.md)
 - [AnalysisLinks](docs/AnalysisLinks.md)
 - [AnalysisRule](docs/AnalysisRule.md)
 - [AnalysisRuleLinks](docs/AnalysisRuleLinks.md)
 - [AnalysisRulePlugIn](docs/AnalysisRulePlugIn.md)
 - [AnalysisRulePlugInLinks](docs/AnalysisRulePlugInLinks.md)
 - [AnalysisTemplate](docs/AnalysisTemplate.md)
 - [AnalysisTemplateLinks](docs/AnalysisTemplateLinks.md)
 - [Annotation](docs/Annotation.md)
 - [AnnotationLinks](docs/AnnotationLinks.md)
 - [AssetDatabase](docs/AssetDatabase.md)
 - [AssetDatabaseLinks](docs/AssetDatabaseLinks.md)
 - [AssetServer](docs/AssetServer.md)
 - [AssetServerLinks](docs/AssetServerLinks.md)
 - [Attribute](docs/Attribute.md)
 - [AttributeCategory](docs/AttributeCategory.md)
 - [AttributeCategoryLinks](docs/AttributeCategoryLinks.md)
 - [AttributeLinks](docs/AttributeLinks.md)
 - [AttributeTemplate](docs/AttributeTemplate.md)
 - [AttributeTemplateLinks](docs/AttributeTemplateLinks.md)
 - [AttributeTrait](docs/AttributeTrait.md)
 - [AttributeTraitLinks](docs/AttributeTraitLinks.md)
 - [CacheInstance](docs/CacheInstance.md)
 - [ChannelInstance](docs/ChannelInstance.md)
 - [DataServer](docs/DataServer.md)
 - [DataServerLicense](docs/DataServerLicense.md)
 - [DataServerLicenseLinks](docs/DataServerLicenseLinks.md)
 - [DataServerLinks](docs/DataServerLinks.md)
 - [Element](docs/Element.md)
 - [ElementCategory](docs/ElementCategory.md)
 - [ElementCategoryLinks](docs/ElementCategoryLinks.md)
 - [ElementLinks](docs/ElementLinks.md)
 - [ElementTemplate](docs/ElementTemplate.md)
 - [ElementTemplateLinks](docs/ElementTemplateLinks.md)
 - [EnumerationSet](docs/EnumerationSet.md)
 - [EnumerationSetLinks](docs/EnumerationSetLinks.md)
 - [EnumerationValue](docs/EnumerationValue.md)
 - [EnumerationValueLinks](docs/EnumerationValueLinks.md)
 - [Errors](docs/Errors.md)
 - [EventFrame](docs/EventFrame.md)
 - [EventFrameLinks](docs/EventFrameLinks.md)
 - [ExtendedTimedValue](docs/ExtendedTimedValue.md)
 - [ExtendedTimedValues](docs/ExtendedTimedValues.md)
 - [ItemAttribute](docs/ItemAttribute.md)
 - [ItemElement](docs/ItemElement.md)
 - [ItemEventFrame](docs/ItemEventFrame.md)
 - [ItemPoint](docs/ItemPoint.md)
 - [ItemsAnalysis](docs/ItemsAnalysis.md)
 - [ItemsAnalysisCategory](docs/ItemsAnalysisCategory.md)
 - [ItemsAnalysisRule](docs/ItemsAnalysisRule.md)
 - [ItemsAnalysisRulePlugIn](docs/ItemsAnalysisRulePlugIn.md)
 - [ItemsAnalysisTemplate](docs/ItemsAnalysisTemplate.md)
 - [ItemsAnnotation](docs/ItemsAnnotation.md)
 - [ItemsAssetDatabase](docs/ItemsAssetDatabase.md)
 - [ItemsAssetServer](docs/ItemsAssetServer.md)
 - [ItemsAttribute](docs/ItemsAttribute.md)
 - [ItemsAttributeCategory](docs/ItemsAttributeCategory.md)
 - [ItemsAttributeTemplate](docs/ItemsAttributeTemplate.md)
 - [ItemsAttributeTrait](docs/ItemsAttributeTrait.md)
 - [ItemsCacheInstance](docs/ItemsCacheInstance.md)
 - [ItemsChannelInstance](docs/ItemsChannelInstance.md)
 - [ItemsDataServer](docs/ItemsDataServer.md)
 - [ItemsElement](docs/ItemsElement.md)
 - [ItemsElementCategory](docs/ItemsElementCategory.md)
 - [ItemsElementTemplate](docs/ItemsElementTemplate.md)
 - [ItemsEnumerationSet](docs/ItemsEnumerationSet.md)
 - [ItemsEnumerationValue](docs/ItemsEnumerationValue.md)
 - [ItemsEventFrame](docs/ItemsEventFrame.md)
 - [ItemsItemAttribute](docs/ItemsItemAttribute.md)
 - [ItemsItemElement](docs/ItemsItemElement.md)
 - [ItemsItemEventFrame](docs/ItemsItemEventFrame.md)
 - [ItemsItemPoint](docs/ItemsItemPoint.md)
 - [ItemsItemsSubstatus](docs/ItemsItemsSubstatus.md)
 - [ItemsNotificationContactTemplate](docs/ItemsNotificationContactTemplate.md)
 - [ItemsNotificationRule](docs/ItemsNotificationRule.md)
 - [ItemsNotificationRuleSubscriber](docs/ItemsNotificationRuleSubscriber.md)
 - [ItemsNotificationRuleTemplate](docs/ItemsNotificationRuleTemplate.md)
 - [ItemsPoint](docs/ItemsPoint.md)
 - [ItemsPointAttribute](docs/ItemsPointAttribute.md)
 - [ItemsSecurityEntry](docs/ItemsSecurityEntry.md)
 - [ItemsSecurityIdentity](docs/ItemsSecurityIdentity.md)
 - [ItemsSecurityMapping](docs/ItemsSecurityMapping.md)
 - [ItemsSecurityRights](docs/ItemsSecurityRights.md)
 - [ItemsStreamSummaries](docs/ItemsStreamSummaries.md)
 - [ItemsStreamValue](docs/ItemsStreamValue.md)
 - [ItemsStreamValues](docs/ItemsStreamValues.md)
 - [ItemsString](docs/ItemsString.md)
 - [ItemsSubstatus](docs/ItemsSubstatus.md)
 - [ItemsSummaryValue](docs/ItemsSummaryValue.md)
 - [ItemsTable](docs/ItemsTable.md)
 - [ItemsTableCategory](docs/ItemsTableCategory.md)
 - [ItemsTimeRulePlugIn](docs/ItemsTimeRulePlugIn.md)
 - [ItemsUnitClass](docs/ItemsUnitClass.md)
 - [Landing](docs/Landing.md)
 - [LandingLinks](docs/LandingLinks.md)
 - [MediaMetadata](docs/MediaMetadata.md)
 - [MediaMetadataLinks](docs/MediaMetadataLinks.md)
 - [NotificationContactTemplate](docs/NotificationContactTemplate.md)
 - [NotificationContactTemplateLinks](docs/NotificationContactTemplateLinks.md)
 - [NotificationRule](docs/NotificationRule.md)
 - [NotificationRuleSubscriber](docs/NotificationRuleSubscriber.md)
 - [NotificationRuleTemplate](docs/NotificationRuleTemplate.md)
 - [PaginationLinks](docs/PaginationLinks.md)
 - [Point](docs/Point.md)
 - [PointAttribute](docs/PointAttribute.md)
 - [PointAttributeLinks](docs/PointAttributeLinks.md)
 - [PointLinks](docs/PointLinks.md)
 - [PropertyError](docs/PropertyError.md)
 - [Request](docs/Request.md)
 - [RequestTemplate](docs/RequestTemplate.md)
 - [Response](docs/Response.md)
 - [SearchByAttribute](docs/SearchByAttribute.md)
 - [Security](docs/Security.md)
 - [SecurityEntry](docs/SecurityEntry.md)
 - [SecurityEntryLinks](docs/SecurityEntryLinks.md)
 - [SecurityIdentity](docs/SecurityIdentity.md)
 - [SecurityIdentityLinks](docs/SecurityIdentityLinks.md)
 - [SecurityMapping](docs/SecurityMapping.md)
 - [SecurityMappingLinks](docs/SecurityMappingLinks.md)
 - [SecurityRights](docs/SecurityRights.md)
 - [SecurityRightsLinks](docs/SecurityRightsLinks.md)
 - [StreamAnnotation](docs/StreamAnnotation.md)
 - [StreamSummaries](docs/StreamSummaries.md)
 - [StreamSummariesLinks](docs/StreamSummariesLinks.md)
 - [StreamValue](docs/StreamValue.md)
 - [StreamValueLinks](docs/StreamValueLinks.md)
 - [StreamValues](docs/StreamValues.md)
 - [StreamValuesLinks](docs/StreamValuesLinks.md)
 - [Substatus](docs/Substatus.md)
 - [SummaryValue](docs/SummaryValue.md)
 - [SystemLanding](docs/SystemLanding.md)
 - [SystemLandingLinks](docs/SystemLandingLinks.md)
 - [SystemStatus](docs/SystemStatus.md)
 - [Table](docs/Table.md)
 - [TableCategory](docs/TableCategory.md)
 - [TableCategoryLinks](docs/TableCategoryLinks.md)
 - [TableData](docs/TableData.md)
 - [TableLinks](docs/TableLinks.md)
 - [TimeRule](docs/TimeRule.md)
 - [TimeRuleLinks](docs/TimeRuleLinks.md)
 - [TimeRulePlugIn](docs/TimeRulePlugIn.md)
 - [TimeRulePlugInLinks](docs/TimeRulePlugInLinks.md)
 - [TimedValue](docs/TimedValue.md)
 - [TimedValues](docs/TimedValues.md)
 - [Unit](docs/Unit.md)
 - [UnitClass](docs/UnitClass.md)
 - [UnitClassLinks](docs/UnitClassLinks.md)
 - [UnitLinks](docs/UnitLinks.md)
 - [UserInfo](docs/UserInfo.md)
 - [Value](docs/Value.md)
 - [ValueQuery](docs/ValueQuery.md)
 - [Version](docs/Version.md)
 - [WebException](docs/WebException.md)


## Documentation For Authorization
 Endpoints do not require authorization.


## Author

techsupport@osisoft.com

