# \AssetDatabaseApi

All URIs are relative to *https://localhost/piwebapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetDatabaseAddReferencedElement**](AssetDatabaseApi.md#AssetDatabaseAddReferencedElement) | **Post** /assetdatabases/{webId}/referencedelements | Add a reference to an existing element to the specified database.
[**AssetDatabaseCreateAnalysisCategory**](AssetDatabaseApi.md#AssetDatabaseCreateAnalysisCategory) | **Post** /assetdatabases/{webId}/analysiscategories | Create an analysis category at the Asset Database root.
[**AssetDatabaseCreateAnalysisTemplate**](AssetDatabaseApi.md#AssetDatabaseCreateAnalysisTemplate) | **Post** /assetdatabases/{webId}/analysistemplates | Create an analysis template at the Asset Database root.
[**AssetDatabaseCreateAttributeCategory**](AssetDatabaseApi.md#AssetDatabaseCreateAttributeCategory) | **Post** /assetdatabases/{webId}/attributecategories | Create an attribute category at the Asset Database root.
[**AssetDatabaseCreateElement**](AssetDatabaseApi.md#AssetDatabaseCreateElement) | **Post** /assetdatabases/{webId}/elements | Create a child element.
[**AssetDatabaseCreateElementCategory**](AssetDatabaseApi.md#AssetDatabaseCreateElementCategory) | **Post** /assetdatabases/{webId}/elementcategories | Create an element category at the Asset Database root.
[**AssetDatabaseCreateElementTemplate**](AssetDatabaseApi.md#AssetDatabaseCreateElementTemplate) | **Post** /assetdatabases/{webId}/elementtemplates | Create a template at the Asset Database root. Specify InstanceType of \&quot;Element\&quot; or \&quot;EventFrame\&quot; to create element or event frame template respectively. Only these two types of templates can be created.
[**AssetDatabaseCreateEnumerationSet**](AssetDatabaseApi.md#AssetDatabaseCreateEnumerationSet) | **Post** /assetdatabases/{webId}/enumerationsets | Create an enumeration set at the Asset Database.
[**AssetDatabaseCreateEventFrame**](AssetDatabaseApi.md#AssetDatabaseCreateEventFrame) | **Post** /assetdatabases/{webId}/eventframes | Create an event frame.
[**AssetDatabaseCreateSecurityEntry**](AssetDatabaseApi.md#AssetDatabaseCreateSecurityEntry) | **Post** /assetdatabases/{webId}/securityentries | Create a security entry owned by the asset database.
[**AssetDatabaseCreateTable**](AssetDatabaseApi.md#AssetDatabaseCreateTable) | **Post** /assetdatabases/{webId}/tables | Create a table on the Asset Database.
[**AssetDatabaseCreateTableCategory**](AssetDatabaseApi.md#AssetDatabaseCreateTableCategory) | **Post** /assetdatabases/{webId}/tablecategories | Create a table category on the Asset Database.
[**AssetDatabaseDelete**](AssetDatabaseApi.md#AssetDatabaseDelete) | **Delete** /assetdatabases/{webId} | Delete an asset database.
[**AssetDatabaseDeleteSecurityEntry**](AssetDatabaseApi.md#AssetDatabaseDeleteSecurityEntry) | **Delete** /assetdatabases/{webId}/securityentries/{name} | Delete a security entry owned by the asset database.
[**AssetDatabaseExport**](AssetDatabaseApi.md#AssetDatabaseExport) | **Get** /assetdatabases/{webId}/export | Export the asset database.
[**AssetDatabaseFindAnalyses**](AssetDatabaseApi.md#AssetDatabaseFindAnalyses) | **Get** /assetdatabases/{webId}/analyses | Retrieve analyses based on the specified conditions.
[**AssetDatabaseFindElementAttributes**](AssetDatabaseApi.md#AssetDatabaseFindElementAttributes) | **Get** /assetdatabases/{webId}/elementattributes | Retrieves a list of element attributes matching the specified filters from the specified asset database.
[**AssetDatabaseFindEventFrameAttributes**](AssetDatabaseApi.md#AssetDatabaseFindEventFrameAttributes) | **Get** /assetdatabases/{webId}/eventframeattributes | Retrieves a list of event frame attributes matching the specified filters from the specified asset database.
[**AssetDatabaseGet**](AssetDatabaseApi.md#AssetDatabaseGet) | **Get** /assetdatabases/{webId} | Retrieve an Asset Database.
[**AssetDatabaseGetAnalysisCategories**](AssetDatabaseApi.md#AssetDatabaseGetAnalysisCategories) | **Get** /assetdatabases/{webId}/analysiscategories | Retrieve analysis categories for a given Asset Database.
[**AssetDatabaseGetAnalysisTemplates**](AssetDatabaseApi.md#AssetDatabaseGetAnalysisTemplates) | **Get** /assetdatabases/{webId}/analysistemplates | Retrieve analysis templates based on the specified criteria. By default, all analysis templates in the specified Asset Database are returned.
[**AssetDatabaseGetAttributeCategories**](AssetDatabaseApi.md#AssetDatabaseGetAttributeCategories) | **Get** /assetdatabases/{webId}/attributecategories | Retrieve attribute categories for a given Asset Database.
[**AssetDatabaseGetByPath**](AssetDatabaseApi.md#AssetDatabaseGetByPath) | **Get** /assetdatabases | Retrieve an Asset Database by path.
[**AssetDatabaseGetElementCategories**](AssetDatabaseApi.md#AssetDatabaseGetElementCategories) | **Get** /assetdatabases/{webId}/elementcategories | Retrieve element categories for a given Asset Database.
[**AssetDatabaseGetElementTemplates**](AssetDatabaseApi.md#AssetDatabaseGetElementTemplates) | **Get** /assetdatabases/{webId}/elementtemplates | Retrieve element templates based on the specified criteria. Only templates of instance type \&quot;Element\&quot; and \&quot;EventFrame\&quot; are returned. By default, all element and event frame templates in the specified Asset Database are returned.
[**AssetDatabaseGetElements**](AssetDatabaseApi.md#AssetDatabaseGetElements) | **Get** /assetdatabases/{webId}/elements | Retrieve elements based on the specified conditions. By default, this method selects immediate children of the specified asset database.
[**AssetDatabaseGetEnumerationSets**](AssetDatabaseApi.md#AssetDatabaseGetEnumerationSets) | **Get** /assetdatabases/{webId}/enumerationsets | Retrieve enumeration sets for given asset database.
[**AssetDatabaseGetEventFrames**](AssetDatabaseApi.md#AssetDatabaseGetEventFrames) | **Get** /assetdatabases/{webId}/eventframes | Retrieve event frames based on the specified conditions. By default, returns all children of the specified root resource that have been active in the past 8 hours.
[**AssetDatabaseGetReferencedElements**](AssetDatabaseApi.md#AssetDatabaseGetReferencedElements) | **Get** /assetdatabases/{webId}/referencedelements | Retrieve referenced elements based on the specified conditions. By default, this method selects all referenced elements at the root level of the asset database.
[**AssetDatabaseGetSecurity**](AssetDatabaseApi.md#AssetDatabaseGetSecurity) | **Get** /assetdatabases/{webId}/security | Get the security information of the specified security item associated with the asset database for a specified user.
[**AssetDatabaseGetSecurityEntries**](AssetDatabaseApi.md#AssetDatabaseGetSecurityEntries) | **Get** /assetdatabases/{webId}/securityentries | Retrieve the security entries of the specified security item associated with the asset database based on the specified criteria. By default, all security entries for this asset database are returned.
[**AssetDatabaseGetSecurityEntryByName**](AssetDatabaseApi.md#AssetDatabaseGetSecurityEntryByName) | **Get** /assetdatabases/{webId}/securityentries/{name} | Retrieve the security entry of the specified security item associated with the asset database with the specified name.
[**AssetDatabaseGetTableCategories**](AssetDatabaseApi.md#AssetDatabaseGetTableCategories) | **Get** /assetdatabases/{webId}/tablecategories | Retrieve table categories for a given Asset Database.
[**AssetDatabaseGetTables**](AssetDatabaseApi.md#AssetDatabaseGetTables) | **Get** /assetdatabases/{webId}/tables | Retrieve tables for given Asset Database.
[**AssetDatabaseImport**](AssetDatabaseApi.md#AssetDatabaseImport) | **Post** /assetdatabases/{webId}/import | Import an asset database.
[**AssetDatabaseRemoveReferencedElement**](AssetDatabaseApi.md#AssetDatabaseRemoveReferencedElement) | **Delete** /assetdatabases/{webId}/referencedelements | Remove a reference to an existing element from the specified database.
[**AssetDatabaseUpdate**](AssetDatabaseApi.md#AssetDatabaseUpdate) | **Patch** /assetdatabases/{webId} | Update an asset database by replacing items in its definition.
[**AssetDatabaseUpdateSecurityEntry**](AssetDatabaseApi.md#AssetDatabaseUpdateSecurityEntry) | **Put** /assetdatabases/{webId}/securityentries/{name} | Update a security entry owned by the asset database.


# **AssetDatabaseAddReferencedElement**
> AssetDatabaseAddReferencedElement(ctx, webId, referencedElementWebId, optional)
Add a reference to an existing element to the specified database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the database which the referenced element will be added to. | 
  **referencedElementWebId** | [**[]string**](string.md)| The ID of the referenced element. Multiple referenced elements may be specified with multiple instances of the parameter. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the database which the referenced element will be added to. | 
 **referencedElementWebId** | [**[]string**](string.md)| The ID of the referenced element. Multiple referenced elements may be specified with multiple instances of the parameter. | 
 **referenceType** | **string**| The name of the reference type between the parent and the referenced element. This must be a \&quot;strong\&quot; reference type. The default is \&quot;parent-child\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseCreateAnalysisCategory**
> AssetDatabaseCreateAnalysisCategory(ctx, webId, analysisCategory, optional)
Create an analysis category at the Asset Database root.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the database in which to create the analysis category. | 
  **analysisCategory** | [**AnalysisCategory**](AnalysisCategory.md)| The new analysis category definition. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the database in which to create the analysis category. | 
 **analysisCategory** | [**AnalysisCategory**](AnalysisCategory.md)| The new analysis category definition. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseCreateAnalysisTemplate**
> AssetDatabaseCreateAnalysisTemplate(ctx, webId, template, optional)
Create an analysis template at the Asset Database root.

Analyses that are based on an analysis template will inherit characteristics defined in the template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the database in which to create the analysis template. | 
  **template** | [**AnalysisTemplate**](AnalysisTemplate.md)| The new analysis template definition. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the database in which to create the analysis template. | 
 **template** | [**AnalysisTemplate**](AnalysisTemplate.md)| The new analysis template definition. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseCreateAttributeCategory**
> AssetDatabaseCreateAttributeCategory(ctx, webId, attributeCategory, optional)
Create an attribute category at the Asset Database root.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the database in which to create the attribute category. | 
  **attributeCategory** | [**AttributeCategory**](AttributeCategory.md)| The new attribute category definition. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the database in which to create the attribute category. | 
 **attributeCategory** | [**AttributeCategory**](AttributeCategory.md)| The new attribute category definition. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseCreateElement**
> AssetDatabaseCreateElement(ctx, webId, element, optional)
Create a child element.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the asset database on which to create the element. | 
  **element** | [**Element**](Element.md)| The new element definition. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the asset database on which to create the element. | 
 **element** | [**Element**](Element.md)| The new element definition. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseCreateElementCategory**
> AssetDatabaseCreateElementCategory(ctx, webId, elementCategory, optional)
Create an element category at the Asset Database root.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the database in which to create the element category. | 
  **elementCategory** | [**ElementCategory**](ElementCategory.md)| The new element category definition. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the database in which to create the element category. | 
 **elementCategory** | [**ElementCategory**](ElementCategory.md)| The new element category definition. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseCreateElementTemplate**
> AssetDatabaseCreateElementTemplate(ctx, webId, template, optional)
Create a template at the Asset Database root. Specify InstanceType of \"Element\" or \"EventFrame\" to create element or event frame template respectively. Only these two types of templates can be created.

Elements and event frames that are based on an element template will inherit characteristics defined in the template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the database in which to create the element template. | 
  **template** | [**ElementTemplate**](ElementTemplate.md)| The new element template definition. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the database in which to create the element template. | 
 **template** | [**ElementTemplate**](ElementTemplate.md)| The new element template definition. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseCreateEnumerationSet**
> AssetDatabaseCreateEnumerationSet(ctx, webId, enumerationSet, optional)
Create an enumeration set at the Asset Database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the database in which to create the enumeration set. | 
  **enumerationSet** | [**EnumerationSet**](EnumerationSet.md)| The new enumeration set definition. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the database in which to create the enumeration set. | 
 **enumerationSet** | [**EnumerationSet**](EnumerationSet.md)| The new enumeration set definition. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseCreateEventFrame**
> AssetDatabaseCreateEventFrame(ctx, webId, eventFrame, optional)
Create an event frame.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the database on which to create the event frame. | 
  **eventFrame** | [**EventFrame**](EventFrame.md)| The new event frame definition. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the database on which to create the event frame. | 
 **eventFrame** | [**EventFrame**](EventFrame.md)| The new event frame definition. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseCreateSecurityEntry**
> AssetDatabaseCreateSecurityEntry(ctx, webId, securityEntry, optional)
Create a security entry owned by the asset database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the asset database where the security entry will be created. | 
  **securityEntry** | [**SecurityEntry**](SecurityEntry.md)| The new security entry definition. The full list of allow and deny rights must be supplied. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the asset database where the security entry will be created. | 
 **securityEntry** | [**SecurityEntry**](SecurityEntry.md)| The new security entry definition. The full list of allow and deny rights must be supplied. | 
 **applyToChildren** | **bool**| If false, the new access permissions are only applied to the associated object. If true, the access permissions of children with any parent-child reference types will change when the permissions on the primary parent change. | 
 **securityItem** | **string**| The security item of the desired security entries to be created. If the parameter is not specified, security entries of the &#39;Default&#39; security item will be created. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseCreateTable**
> AssetDatabaseCreateTable(ctx, webId, table, optional)
Create a table on the Asset Database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the database in which to create the table. | 
  **table** | [**Table**](Table.md)| The new table definition. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the database in which to create the table. | 
 **table** | [**Table**](Table.md)| The new table definition. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseCreateTableCategory**
> AssetDatabaseCreateTableCategory(ctx, webId, tableCategory, optional)
Create a table category on the Asset Database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the database in which to create the table category. | 
  **tableCategory** | [**TableCategory**](TableCategory.md)| The new table category definition. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the database in which to create the table category. | 
 **tableCategory** | [**TableCategory**](TableCategory.md)| The new table category definition. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseDelete**
> AssetDatabaseDelete(ctx, webId)
Delete an asset database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the database. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseDeleteSecurityEntry**
> AssetDatabaseDeleteSecurityEntry(ctx, name, webId, optional)
Delete a security entry owned by the asset database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The name of the security entry. For every backslash character (\\) in the security entry name, replace with asterisk (*). As an example, use domain*username instead of domain\\username. | 
  **webId** | **string**| The ID of the asset database where the security entry will be deleted. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The name of the security entry. For every backslash character (\\) in the security entry name, replace with asterisk (*). As an example, use domain*username instead of domain\\username. | 
 **webId** | **string**| The ID of the asset database where the security entry will be deleted. | 
 **applyToChildren** | **bool**| If false, the new access permissions are only applied to the associated object. If true, the access permissions of children with any parent-child reference types will change when the permissions on the primary parent change. | 
 **securityItem** | **string**| The security item of the desired security entries to be deleted. If the parameter is not specified, security entries of the &#39;Default&#39; security item will be deleted. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseExport**
> AssetDatabaseExport(ctx, webId, optional)
Export the asset database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the database. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the database. | 
 **endTime** | **string**| The latest ending time for AFEventFrame, AFTransfer, and AFCase objects that may be part of the export. The default is &#39;*&#39;. | 
 **exportMode** | [**[]string**](string.md)| Indicates the type of export to perform. The default is &#39;StrongReferences&#39;. Multiple export modes may be specified by using multiple instances of exportMode. | 
 **startTime** | **string**| The earliest starting time for AFEventFrame, AFTransfer, and AFCase objects that may be part of the export. The default is &#39;*-30d&#39;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseFindAnalyses**
> ItemsAnalysis AssetDatabaseFindAnalyses(ctx, webId, field, optional)
Retrieve analyses based on the specified conditions.

Users can search for the analyses based on specific search parameters. If no parameters are specified in the search, the default values for each parameter will be used and will return the analyses that match the default search.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the database to search for the analyses. | 
  **field** | [**[]string**](string.md)| Specifies which of the object&#39;s properties are searched. Multiple search fields may be specified with multiple instances of the parameter. The default is &#39;Name&#39;. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the database to search for the analyses. | 
 **field** | [**[]string**](string.md)| Specifies which of the object&#39;s properties are searched. Multiple search fields may be specified with multiple instances of the parameter. The default is &#39;Name&#39;. | 
 **maxCount** | **int32**| The maximum number of objects to be returned per call (page size). The default is 1000. | 
 **query** | **string**| The query string used for finding analyses. The default is null. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **sortField** | **string**| The field or property of the object used to sort the returned collection. The default is &#39;Name&#39;. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39;. | 
 **startIndex** | **int32**| The starting index (zero based) of the items to be returned. The default is 0. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsAnalysis**](Items[Analysis].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseFindElementAttributes**
> ItemsAttribute AssetDatabaseFindElementAttributes(ctx, webId, optional)
Retrieves a list of element attributes matching the specified filters from the specified asset database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the asset database to use as the root of the search. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the asset database to use as the root of the search. | 
 **attributeCategory** | **string**| Specify that returned attributes must have this category. The default is no filter. | 
 **attributeDescriptionFilter** | **string**| The attribute description filter string used for finding objects. Only the first 440 characters of the description will be searched. For Asset Servers older than 2.7, a 400 status code (Bad Request) will be returned if this parameter is specified. The default is no filter. | 
 **attributeNameFilter** | **string**| The attribute name filter string used for finding objects. The default is no filter. | 
 **attributeType** | **string**| Specify that returned attributes&#39; value type must be this value type. The default is no filter. | 
 **elementCategory** | **string**| Specify that the owner of the returned attributes must have this category. The default is no filter. | 
 **elementDescriptionFilter** | **string**| The element description filter string used for finding objects. Only the first 440 characters of the description will be searched. For Asset Servers older than 2.7, a 400 status code (Bad Request) will be returned if this parameter is specified. The default is no filter. | 
 **elementNameFilter** | **string**| The element name filter string used for finding objects. The default is no filter. | 
 **elementTemplate** | **string**| Specify that the owner of the returned attributes must have this template or a template derived from this template. The default is no filter. | 
 **elementType** | **string**| Specify that the element of the returned attributes must have this AFElementType. The default is no filter. | 
 **maxCount** | **int32**| The maximum number of objects to be returned (the page size). The default is 1000. | 
 **searchFullHierarchy** | **bool**| Specifies if the search should include objects nested further than immediate children of the given resource. The default is &#39;false&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **sortField** | **string**| The field or property of the object used to sort the returned collection. The default is &#39;Name&#39;. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39;. | 
 **startIndex** | **int32**| The starting index (zero based) of the items to be returned. The default is 0. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsAttribute**](Items[Attribute].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseFindEventFrameAttributes**
> ItemsAttribute AssetDatabaseFindEventFrameAttributes(ctx, webId, optional)
Retrieves a list of event frame attributes matching the specified filters from the specified asset database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the asset database to use as the root of the search. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the asset database to use as the root of the search. | 
 **attributeCategory** | **string**| Specify that returned attributes must have this category. The default is no filter. | 
 **attributeDescriptionFilter** | **string**| The attribute description filter string used for finding objects. Only the first 440 characters of the description will be searched. For Asset Servers older than 2.7, a 400 status code (Bad Request) will be returned if this parameter is specified. The default is no filter. | 
 **attributeNameFilter** | **string**| The attribute name filter string used for finding objects. The default is no filter. | 
 **attributeType** | **string**| Specify that returned attributes&#39; value type must be this value type. The default is no filter. | 
 **endTime** | **string**| A string representing the latest ending time for the event frames to be matched. The endTime must be greater than or equal to the startTime. The default is &#39;*&#39;. | 
 **eventFrameCategory** | **string**| Specify that the owner of the returned attributes must have this category. The default is no filter. | 
 **eventFrameDescriptionFilter** | **string**| The event frame description filter string used for finding objects. Only the first 440 characters of the description will be searched. For Asset Servers older than 2.7, a 400 status code (Bad Request) will be returned if this parameter is specified. The default is no filter. | 
 **eventFrameNameFilter** | **string**| The event frame name filter string used for finding objects. The default is no filter. | 
 **eventFrameTemplate** | **string**| Specify that the owner of the returned attributes must have this template or a template derived from this template. The default is no filter. | 
 **maxCount** | **int32**| The maximum number of objects to be returned (the page size). The default is 1000. | 
 **referencedElementNameFilter** | **string**| The name query string which must match the name of a referenced element. The default is no filter. | 
 **searchFullHierarchy** | **bool**| Specifies if the search should include objects nested further than immediate children of the given resource. The default is &#39;false&#39;. | 
 **searchMode** | **string**| Determines how the startTime and endTime parameters are treated when searching for event frames. The default is &#39;Overlapped&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **sortField** | **string**| The field or property of the object used to sort the returned collection. The default is &#39;Name&#39;. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39;. | 
 **startIndex** | **int32**| The starting index (zero based) of the items to be returned. The default is 0. | 
 **startTime** | **string**| A string representing the earliest starting time for the event frames to be matched. startTime must be less than or equal to the endTime. The default is &#39;*-8h&#39;. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsAttribute**](Items[Attribute].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseGet**
> AssetDatabase AssetDatabaseGet(ctx, webId, optional)
Retrieve an Asset Database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the database. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the database. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**AssetDatabase**](AssetDatabase.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseGetAnalysisCategories**
> ItemsAnalysisCategory AssetDatabaseGetAnalysisCategories(ctx, webId, optional)
Retrieve analysis categories for a given Asset Database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the database. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the database. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsAnalysisCategory**](Items[AnalysisCategory].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseGetAnalysisTemplates**
> ItemsAnalysisTemplate AssetDatabaseGetAnalysisTemplates(ctx, webId, field, optional)
Retrieve analysis templates based on the specified criteria. By default, all analysis templates in the specified Asset Database are returned.

Users can search for the analysis templates based on specific search parameters. If no parameters are specified in the search, the default values for each parameter will be used and will return the templates that match the default search.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the database to search. | 
  **field** | [**[]string**](string.md)| Specifies which of the object&#39;s properties are searched. Multiple search fields may be specified with multiple instances of the parameter. The default is &#39;Name&#39;. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the database to search. | 
 **field** | [**[]string**](string.md)| Specifies which of the object&#39;s properties are searched. Multiple search fields may be specified with multiple instances of the parameter. The default is &#39;Name&#39;. | 
 **maxCount** | **int32**| The maximum number of objects to be returned per call (page size). The default is 1000. | 
 **query** | **string**| The query string used for finding objects. The default is no query string. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **sortField** | **string**| The field or property of the object used to sort the returned collection. The default is &#39;Name&#39;. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39;. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsAnalysisTemplate**](Items[AnalysisTemplate].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseGetAttributeCategories**
> ItemsAttributeCategory AssetDatabaseGetAttributeCategories(ctx, webId, optional)
Retrieve attribute categories for a given Asset Database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the database. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the database. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsAttributeCategory**](Items[AttributeCategory].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseGetByPath**
> AssetDatabase AssetDatabaseGetByPath(ctx, path, optional)
Retrieve an Asset Database by path.

This method returns an asset database based on the hierarchical path associated with it, and should be used when a path has been received from a separate part of the PI System for use in the PI Web API. Users should primarily search with the WebID when available.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **path** | **string**| The path to the database. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| The path to the database. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**AssetDatabase**](AssetDatabase.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseGetElementCategories**
> ItemsElementCategory AssetDatabaseGetElementCategories(ctx, webId, optional)
Retrieve element categories for a given Asset Database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the database. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the database. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsElementCategory**](Items[ElementCategory].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseGetElementTemplates**
> ItemsElementTemplate AssetDatabaseGetElementTemplates(ctx, webId, field, optional)
Retrieve element templates based on the specified criteria. Only templates of instance type \"Element\" and \"EventFrame\" are returned. By default, all element and event frame templates in the specified Asset Database are returned.

Users can search for the element and event frame template based on specific search parameters. If no parameters are specified in the search, the default values for each parameter will be used and will return the templates that match the default search.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the database to search. | 
  **field** | [**[]string**](string.md)| Specifies which of the object&#39;s properties are searched. Multiple search fields may be specified with multiple instances of the parameter. The default is &#39;Name&#39;. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the database to search. | 
 **field** | [**[]string**](string.md)| Specifies which of the object&#39;s properties are searched. Multiple search fields may be specified with multiple instances of the parameter. The default is &#39;Name&#39;. | 
 **maxCount** | **int32**| The maximum number of objects to be returned per call (page size). The default is 1000. | 
 **query** | **string**| The query string used for finding objects. The default is no query string. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **sortField** | **string**| The field or property of the object used to sort the returned collection. The default is &#39;Name&#39;. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39;. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsElementTemplate**](Items[ElementTemplate].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseGetElements**
> ItemsElement AssetDatabaseGetElements(ctx, webId, optional)
Retrieve elements based on the specified conditions. By default, this method selects immediate children of the specified asset database.

Users can search for the elements based on specific search parameters. If no parameters are specified in the search, the default values for each parameter will be used and will return the elements that match the default search.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the database to use as the root of the search. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the database to use as the root of the search. | 
 **categoryName** | **string**| Specify that returned elements must have this category. The default is no category filter. | 
 **descriptionFilter** | **string**| Specify that returned elements must have this description. The default is no description filter. | 
 **elementType** | **string**| Specify that returned elements must have this type. The default type is &#39;Any&#39;. | 
 **maxCount** | **int32**| The maximum number of objects to be returned per call (page size). The default is 1000. | 
 **nameFilter** | **string**| The name query string used for finding objects. The default is no filter. | 
 **searchFullHierarchy** | **bool**| Specifies if the search should include objects nested further than the immediate children of the searchRoot. The default is &#39;false&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **sortField** | **string**| The field or property of the object used to sort the returned collection. The default is &#39;Name&#39;. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39;. | 
 **startIndex** | **int32**| The starting index (zero based) of the items to be returned. The default is 0. | 
 **templateName** | **string**| Specify that returned elements must have this template or a template derived from this template. The default is no template filter. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsElement**](Items[Element].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseGetEnumerationSets**
> ItemsEnumerationSet AssetDatabaseGetEnumerationSets(ctx, webId, optional)
Retrieve enumeration sets for given asset database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the database. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the database. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsEnumerationSet**](Items[EnumerationSet].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseGetEventFrames**
> ItemsEventFrame AssetDatabaseGetEventFrames(ctx, webId, optional)
Retrieve event frames based on the specified conditions. By default, returns all children of the specified root resource that have been active in the past 8 hours.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the asset database to use as the root of the search. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the asset database to use as the root of the search. | 
 **canBeAcknowledged** | **bool**| Specify the returned event frames&#39; canBeAcknowledged property. The default is no canBeAcknowledged filter. | 
 **categoryName** | **string**| Specify that returned event frames must have this category. The default is no category filter. | 
 **endTime** | **string**| The ending time for the search. The endTime must be greater than or equal to the startTime. The searchMode parameter will control whether the comparison will be performed against the event frame&#39;s startTime or endTime. The default is &#39;*&#39; if searchMode is not one of the &#39;Backward*&#39; or &#39;Forward*&#39; values. | 
 **isAcknowledged** | **bool**| Specify the returned event frames&#39; isAcknowledged property. The default no isAcknowledged filter. | 
 **maxCount** | **int32**| The maximum number of objects to be returned per call (page size). The default is 1000. | 
 **nameFilter** | **string**| The name query string used for finding event frames. The default is no filter. | 
 **referencedElementNameFilter** | **string**| The name query string which must match the name of a referenced element. The default is no filter. | 
 **referencedElementTemplateName** | **string**| Specify that returned event frames must have an element in the event frame&#39;s referenced elements collection that derives from the template. Specify this parameter by name. | 
 **searchFullHierarchy** | **bool**| Specifies whether the search should include objects nested further than the immediate children of the search root. The default is &#39;false&#39;. | 
 **searchMode** | **string**| Determines how the startTime and endTime parameters are treated when searching for event frame objects to be included in the returned collection. If this parameter is one of the &#39;Backward*&#39; or &#39;Forward*&#39; values, none of endTime, sortField, or sortOrder may be specified. The default is &#39;Overlapped&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **severity** | [**[]string**](string.md)| Specify that returned event frames must have this severity. Multiple severity values may be specified with multiple instances of the parameter. The default is no severity filter. | 
 **sortField** | **string**| The field or property of the object used to sort the returned collection. The default is &#39;Name&#39; if searchMode is not one of the &#39;Backward*&#39; or &#39;Forward*&#39; values. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39; if searchMode is not one of the &#39;Backward*&#39; or &#39;Forward*&#39; values. | 
 **startIndex** | **int32**| The starting index (zero based) of the items to be returned. The default is 0. | 
 **startTime** | **string**| The starting time for the search. startTime must be less than or equal to the endTime. The searchMode parameter will control whether the comparison will be performed against the event frame&#39;s startTime or endTime. The default is &#39;*-8h&#39;. | 
 **templateName** | **string**| Specify that returned event frames must have this template or a template derived from this template. The default is no template filter. Specify this parameter by name. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsEventFrame**](Items[EventFrame].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseGetReferencedElements**
> ItemsElement AssetDatabaseGetReferencedElements(ctx, webId, optional)
Retrieve referenced elements based on the specified conditions. By default, this method selects all referenced elements at the root level of the asset database.

Users can search for the referenced elements based on specific search parameters. If no parameters are specified in the search, the default values for each parameter will be used and will return the elements that match the default search.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the resource to use as the root of the search. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the resource to use as the root of the search. | 
 **categoryName** | **string**| Specify that returned elements must have this category. The default is no category filter. | 
 **descriptionFilter** | **string**| Specify that returned elements must have this description. The default is no description filter. | 
 **elementType** | **string**| Specify that returned elements must have this type. The default type is &#39;Any&#39;. | 
 **maxCount** | **int32**| The maximum number of objects to be returned per call (page size). The default is 1000. | 
 **nameFilter** | **string**| The name query string used for finding objects. The default is no filter. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **sortField** | **string**| The field or property of the object used to sort the returned collection. The default is &#39;Name&#39;. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39;. | 
 **startIndex** | **int32**| The starting index (zero based) of the items to be returned. The default is 0. | 
 **templateName** | **string**| Specify that returned elements must have this template or a template derived from this template. The default is no template filter. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsElement**](Items[Element].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseGetSecurity**
> ItemsSecurityRights AssetDatabaseGetSecurity(ctx, webId, securityItem, userIdentity, optional)
Get the security information of the specified security item associated with the asset database for a specified user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the asset database for the security to be checked. | 
  **securityItem** | [**[]string**](string.md)| The security item of the desired security information to be returned. Multiple security items may be specified with multiple instances of the parameter. If the parameter is not specified, only &#39;Default&#39; security item of the security information will be returned. | 
  **userIdentity** | [**[]string**](string.md)| The user identity for the security information to be checked. Multiple security identities may be specified with multiple instances of the parameter. If the parameter is not specified, only the current user&#39;s security rights will be returned. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the asset database for the security to be checked. | 
 **securityItem** | [**[]string**](string.md)| The security item of the desired security information to be returned. Multiple security items may be specified with multiple instances of the parameter. If the parameter is not specified, only &#39;Default&#39; security item of the security information will be returned. | 
 **userIdentity** | [**[]string**](string.md)| The user identity for the security information to be checked. Multiple security identities may be specified with multiple instances of the parameter. If the parameter is not specified, only the current user&#39;s security rights will be returned. | 
 **forceRefresh** | **bool**| Indicates if the security cache should be refreshed before getting security information. The default is &#39;false&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsSecurityRights**](Items[SecurityRights].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseGetSecurityEntries**
> ItemsSecurityEntry AssetDatabaseGetSecurityEntries(ctx, webId, optional)
Retrieve the security entries of the specified security item associated with the asset database based on the specified criteria. By default, all security entries for this asset database are returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the asset database. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the asset database. | 
 **nameFilter** | **string**| The name query string used for filtering security entries. The default is no filter. | 
 **securityItem** | **string**| The security item of the desired security entries information to be returned. If the parameter is not specified, security entries of the &#39;Default&#39; security item will be returned. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsSecurityEntry**](Items[SecurityEntry].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseGetSecurityEntryByName**
> SecurityEntry AssetDatabaseGetSecurityEntryByName(ctx, name, webId, optional)
Retrieve the security entry of the specified security item associated with the asset database with the specified name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The name of the security entry. For every backslash character (\\) in the security entry name, replace with asterisk (*). As an example, use domain*username instead of domain\\username. | 
  **webId** | **string**| The ID of the asset database. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The name of the security entry. For every backslash character (\\) in the security entry name, replace with asterisk (*). As an example, use domain*username instead of domain\\username. | 
 **webId** | **string**| The ID of the asset database. | 
 **securityItem** | **string**| The security item of the desired security entries information to be returned. If the parameter is not specified, security entries of the &#39;Default&#39; security item will be returned. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**SecurityEntry**](SecurityEntry.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseGetTableCategories**
> ItemsTableCategory AssetDatabaseGetTableCategories(ctx, webId, optional)
Retrieve table categories for a given Asset Database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the database. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the database. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsTableCategory**](Items[TableCategory].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseGetTables**
> ItemsTable AssetDatabaseGetTables(ctx, webId, optional)
Retrieve tables for given Asset Database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the database. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the database. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsTable**](Items[Table].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseImport**
> AssetDatabaseImport(ctx, webId, optional)
Import an asset database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the asset database. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the asset database. | 
 **importMode** | [**[]string**](string.md)| Indicates the type of import to perform. The default is &#39;AllowCreate | AllowUpdate | AutoCheckIn&#39;. Multiple import modes may be specified by using multiple instances of importMode. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseRemoveReferencedElement**
> AssetDatabaseRemoveReferencedElement(ctx, webId, referencedElementWebId)
Remove a reference to an existing element from the specified database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the database which the referenced element will be removed from. | 
  **referencedElementWebId** | [**[]string**](string.md)| The ID of the referenced element. Multiple referenced elements may be specified with multiple instances of the parameter. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseUpdate**
> AssetDatabaseUpdate(ctx, webId, database)
Update an asset database by replacing items in its definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the database. | 
  **database** | [**AssetDatabase**](AssetDatabase.md)| A partial database containing the desired changes. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetDatabaseUpdateSecurityEntry**
> AssetDatabaseUpdateSecurityEntry(ctx, name, webId, securityEntry, optional)
Update a security entry owned by the asset database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The name of the security entry. | 
  **webId** | **string**| The ID of the asset database where the security entry will be updated. | 
  **securityEntry** | [**SecurityEntry**](SecurityEntry.md)| The new security entry definition. The full list of allow and deny rights must be supplied or they will be removed. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The name of the security entry. | 
 **webId** | **string**| The ID of the asset database where the security entry will be updated. | 
 **securityEntry** | [**SecurityEntry**](SecurityEntry.md)| The new security entry definition. The full list of allow and deny rights must be supplied or they will be removed. | 
 **applyToChildren** | **bool**| If false, the new access permissions are only applied to the associated object. If true, the access permissions of children with any parent-child reference types will change when the permissions on the primary parent change. | 
 **securityItem** | **string**| The security item of the desired security entries to be updated. If the parameter is not specified, security entries of the &#39;Default&#39; security item will be updated. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

