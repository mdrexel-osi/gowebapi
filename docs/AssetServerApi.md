# \AssetServerApi

All URIs are relative to *https://localhost/piwebapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetServerCreateAssetDatabase**](AssetServerApi.md#AssetServerCreateAssetDatabase) | **Post** /assetservers/{webId}/assetdatabases | Create an asset database.
[**AssetServerCreateSecurityEntry**](AssetServerApi.md#AssetServerCreateSecurityEntry) | **Post** /assetservers/{webId}/securityentries | Create a security entry owned by the asset server.
[**AssetServerCreateSecurityIdentity**](AssetServerApi.md#AssetServerCreateSecurityIdentity) | **Post** /assetservers/{webId}/securityidentities | Create a security identity.
[**AssetServerCreateSecurityMapping**](AssetServerApi.md#AssetServerCreateSecurityMapping) | **Post** /assetservers/{webId}/securitymappings | Create a security mapping.
[**AssetServerCreateUnitClass**](AssetServerApi.md#AssetServerCreateUnitClass) | **Post** /assetservers/{webId}/unitclasses | Create a unit class in the specified Asset Server.
[**AssetServerDeleteSecurityEntry**](AssetServerApi.md#AssetServerDeleteSecurityEntry) | **Delete** /assetservers/{webId}/securityentries/{name} | Delete a security entry owned by the asset server.
[**AssetServerGet**](AssetServerApi.md#AssetServerGet) | **Get** /assetservers/{webId} | Retrieve an Asset Server.
[**AssetServerGetAnalysisRulePlugIns**](AssetServerApi.md#AssetServerGetAnalysisRulePlugIns) | **Get** /assetservers/{webId}/analysisruleplugins | Retrieve a list of all Analysis Rule Plug-in&#39;s.
[**AssetServerGetByName**](AssetServerApi.md#AssetServerGetByName) | **Get** /assetservers#name | Retrieve an Asset Server by name.
[**AssetServerGetByPath**](AssetServerApi.md#AssetServerGetByPath) | **Get** /assetservers#path | Retrieve an Asset Server by path.
[**AssetServerGetDatabases**](AssetServerApi.md#AssetServerGetDatabases) | **Get** /assetservers/{webId}/assetdatabases | Retrieve a list of all Asset Databases on the specified Asset Server.
[**AssetServerGetNotificationContactTemplates**](AssetServerApi.md#AssetServerGetNotificationContactTemplates) | **Get** /assetservers/{webId}/notificationcontacttemplates | Retrieve a list of all notification contact templates on the specified Asset Server.
[**AssetServerGetSecurity**](AssetServerApi.md#AssetServerGetSecurity) | **Get** /assetservers/{webId}/security | Get the security information of the specified security item associated with the asset server for a specified user.
[**AssetServerGetSecurityEntries**](AssetServerApi.md#AssetServerGetSecurityEntries) | **Get** /assetservers/{webId}/securityentries | Retrieve the security entries of the specified security item associated with the asset server based on the specified criteria. By default, all security entries for this asset server are returned.
[**AssetServerGetSecurityEntryByName**](AssetServerApi.md#AssetServerGetSecurityEntryByName) | **Get** /assetservers/{webId}/securityentries/{name} | Retrieve the security entry of the specified security item associated with the asset server with the specified name.
[**AssetServerGetSecurityIdentities**](AssetServerApi.md#AssetServerGetSecurityIdentities) | **Get** /assetservers/{webId}/securityidentities | Retrieve security identities based on the specified criteria. By default, all security identities in the specified Asset Server are returned.
[**AssetServerGetSecurityIdentitiesForUser**](AssetServerApi.md#AssetServerGetSecurityIdentitiesForUser) | **Get** /assetservers/{webId}/securityidentities#userIdentity | Retrieve security identities for a specific user.
[**AssetServerGetSecurityMappings**](AssetServerApi.md#AssetServerGetSecurityMappings) | **Get** /assetservers/{webId}/securitymappings | Retrieve security mappings based on the specified criteria. By default, all security mappings in the specified Asset Server are returned.
[**AssetServerGetTimeRulePlugIns**](AssetServerApi.md#AssetServerGetTimeRulePlugIns) | **Get** /assetservers/{webId}/timeruleplugins | Retrieve a list of all Time Rule Plug-in&#39;s.
[**AssetServerGetUnitClasses**](AssetServerApi.md#AssetServerGetUnitClasses) | **Get** /assetservers/{webId}/unitclasses | Retrieve a list of all unit classes on the specified Asset Server.
[**AssetServerList**](AssetServerApi.md#AssetServerList) | **Get** /assetservers | Retrieve a list of all Asset Servers known to this service.
[**AssetServerUpdateSecurityEntry**](AssetServerApi.md#AssetServerUpdateSecurityEntry) | **Put** /assetservers/{webId}/securityentries/{name} | Update a security entry owned by the asset server.


# **AssetServerCreateAssetDatabase**
> AssetServerCreateAssetDatabase(ctx, webId, database, optional)
Create an asset database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the asset server on which to create the database. | 
  **database** | [**AssetDatabase**](AssetDatabase.md)| The new database definition. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the asset server on which to create the database. | 
 **database** | [**AssetDatabase**](AssetDatabase.md)| The new database definition. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetServerCreateSecurityEntry**
> AssetServerCreateSecurityEntry(ctx, webId, securityEntry, optional)
Create a security entry owned by the asset server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the asset server where the security entry will be created. | 
  **securityEntry** | [**SecurityEntry**](SecurityEntry.md)| The new security entry definition. The full list of allow and deny rights must be supplied. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the asset server where the security entry will be created. | 
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

# **AssetServerCreateSecurityIdentity**
> AssetServerCreateSecurityIdentity(ctx, webId, securityIdentity, optional)
Create a security identity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the asset server on which to create the security identity. | 
  **securityIdentity** | [**SecurityIdentity**](SecurityIdentity.md)| The new security identity definition. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the asset server on which to create the security identity. | 
 **securityIdentity** | [**SecurityIdentity**](SecurityIdentity.md)| The new security identity definition. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetServerCreateSecurityMapping**
> AssetServerCreateSecurityMapping(ctx, webId, securityMapping, optional)
Create a security mapping.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the asset server on which to create the security mapping. | 
  **securityMapping** | [**SecurityMapping**](SecurityMapping.md)| The new security mapping definition. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the asset server on which to create the security mapping. | 
 **securityMapping** | [**SecurityMapping**](SecurityMapping.md)| The new security mapping definition. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetServerCreateUnitClass**
> AssetServerCreateUnitClass(ctx, webId, unitClass, optional)
Create a unit class in the specified Asset Server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the server. | 
  **unitClass** | [**UnitClass**](UnitClass.md)| The new unit class definition. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the server. | 
 **unitClass** | [**UnitClass**](UnitClass.md)| The new unit class definition. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetServerDeleteSecurityEntry**
> AssetServerDeleteSecurityEntry(ctx, name, webId, optional)
Delete a security entry owned by the asset server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The name of the security entry. For every backslash character (\\) in the security entry name, replace with asterisk (*). As an example, use domain*username instead of domain\\username. | 
  **webId** | **string**| The ID of the asset server where the security entry will be deleted. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The name of the security entry. For every backslash character (\\) in the security entry name, replace with asterisk (*). As an example, use domain*username instead of domain\\username. | 
 **webId** | **string**| The ID of the asset server where the security entry will be deleted. | 
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

# **AssetServerGet**
> AssetServer AssetServerGet(ctx, webId, optional)
Retrieve an Asset Server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the server. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the server. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**AssetServer**](AssetServer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetServerGetAnalysisRulePlugIns**
> ItemsAnalysisRulePlugIn AssetServerGetAnalysisRulePlugIns(ctx, webId, optional)
Retrieve a list of all Analysis Rule Plug-in's.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the asset server, where the Analysis Rule Plug-in&#39;s are installed. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the asset server, where the Analysis Rule Plug-in&#39;s are installed. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsAnalysisRulePlugIn**](Items[AnalysisRulePlugIn].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetServerGetByName**
> AssetServer AssetServerGetByName(ctx, name, optional)
Retrieve an Asset Server by name.

This method returns an asset server based on the name associated with it. Users should primarily search with the WebID when available.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The name of the server. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The name of the server. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**AssetServer**](AssetServer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetServerGetByPath**
> AssetServer AssetServerGetByPath(ctx, path, optional)
Retrieve an Asset Server by path.

This method returns an asset server based on the hierarchical path associated with it, and should be used when a path has been received from a separate part of the PI System for use in the PI Web API. Users should primarily search with the WebID when available.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **path** | **string**| The path to the server. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| The path to the server. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**AssetServer**](AssetServer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetServerGetDatabases**
> ItemsAssetDatabase AssetServerGetDatabases(ctx, webId, optional)
Retrieve a list of all Asset Databases on the specified Asset Server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the server. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the server. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsAssetDatabase**](Items[AssetDatabase].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetServerGetNotificationContactTemplates**
> ItemsNotificationContactTemplate AssetServerGetNotificationContactTemplates(ctx, webId, optional)
Retrieve a list of all notification contact templates on the specified Asset Server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the server. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the server. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsNotificationContactTemplate**](Items[NotificationContactTemplate].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetServerGetSecurity**
> ItemsSecurityRights AssetServerGetSecurity(ctx, webId, securityItem, userIdentity, optional)
Get the security information of the specified security item associated with the asset server for a specified user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the asset server for the security to be checked. | 
  **securityItem** | [**[]string**](string.md)| The security item of the desired security information to be returned. Multiple security items may be specified with multiple instances of the parameter. If the parameter is not specified, only &#39;Default&#39; security item of the security information will be returned. | 
  **userIdentity** | [**[]string**](string.md)| The user identity for the security information to be checked. Multiple security identities may be specified with multiple instances of the parameter. If the parameter is not specified, only the current user&#39;s security rights will be returned. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the asset server for the security to be checked. | 
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

# **AssetServerGetSecurityEntries**
> ItemsSecurityEntry AssetServerGetSecurityEntries(ctx, webId, optional)
Retrieve the security entries of the specified security item associated with the asset server based on the specified criteria. By default, all security entries for this asset server are returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the asset server. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the asset server. | 
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

# **AssetServerGetSecurityEntryByName**
> SecurityEntry AssetServerGetSecurityEntryByName(ctx, name, webId, optional)
Retrieve the security entry of the specified security item associated with the asset server with the specified name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The name of the security entry. For every backslash character (\\) in the security entry name, replace with asterisk (*). As an example, use domain*username instead of domain\\username. | 
  **webId** | **string**| The ID of the asset server. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The name of the security entry. For every backslash character (\\) in the security entry name, replace with asterisk (*). As an example, use domain*username instead of domain\\username. | 
 **webId** | **string**| The ID of the asset server. | 
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

# **AssetServerGetSecurityIdentities**
> ItemsSecurityIdentity AssetServerGetSecurityIdentities(ctx, webId, optional)
Retrieve security identities based on the specified criteria. By default, all security identities in the specified Asset Server are returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the asset server to search. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the asset server to search. | 
 **field** | **string**| Specifies which of the object&#39;s properties are searched. The default is &#39;Name&#39;. | 
 **maxCount** | **int32**| The maximum number of objects to be returned. The default is 1000. | 
 **query** | **string**| The query string used for finding objects. The default is no query string. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **sortField** | **string**| The field or property of the object used to sort the returned collection. The default is &#39;Name&#39;. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39;. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsSecurityIdentity**](Items[SecurityIdentity].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetServerGetSecurityIdentitiesForUser**
> ItemsSecurityIdentity AssetServerGetSecurityIdentitiesForUser(ctx, webId, userIdentity, optional)
Retrieve security identities for a specific user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the server. | 
  **userIdentity** | **string**| The user identity to search for. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the server. | 
 **userIdentity** | **string**| The user identity to search for. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsSecurityIdentity**](Items[SecurityIdentity].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetServerGetSecurityMappings**
> ItemsSecurityMapping AssetServerGetSecurityMappings(ctx, webId, optional)
Retrieve security mappings based on the specified criteria. By default, all security mappings in the specified Asset Server are returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the asset server to search. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the asset server to search. | 
 **field** | **string**| Specifies which of the object&#39;s properties are searched. The default is &#39;Name&#39;. | 
 **maxCount** | **int32**| The maximum number of objects to be returned. The default is 1000. | 
 **query** | **string**| The query string used for finding objects. The default is no query string. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **sortField** | **string**| The field or property of the object used to sort the returned collection. The default is &#39;Name&#39;. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39;. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsSecurityMapping**](Items[SecurityMapping].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetServerGetTimeRulePlugIns**
> ItemsTimeRulePlugIn AssetServerGetTimeRulePlugIns(ctx, webId, optional)
Retrieve a list of all Time Rule Plug-in's.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the asset server, where the Time Rule Plug-in&#39;s are installed. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the asset server, where the Time Rule Plug-in&#39;s are installed. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsTimeRulePlugIn**](Items[TimeRulePlugIn].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetServerGetUnitClasses**
> ItemsUnitClass AssetServerGetUnitClasses(ctx, webId, optional)
Retrieve a list of all unit classes on the specified Asset Server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the server. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the server. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsUnitClass**](Items[UnitClass].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetServerList**
> ItemsAssetServer AssetServerList(ctx, optional)
Retrieve a list of all Asset Servers known to this service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsAssetServer**](Items[AssetServer].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssetServerUpdateSecurityEntry**
> AssetServerUpdateSecurityEntry(ctx, name, webId, securityEntry, optional)
Update a security entry owned by the asset server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The name of the security entry. | 
  **webId** | **string**| The ID of the asset server where the security entry will be updated. | 
  **securityEntry** | [**SecurityEntry**](SecurityEntry.md)| The new security entry definition. The full list of allow and deny rights must be supplied or they will be removed. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The name of the security entry. | 
 **webId** | **string**| The ID of the asset server where the security entry will be updated. | 
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

