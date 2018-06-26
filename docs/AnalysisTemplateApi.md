# \AnalysisTemplateApi

All URIs are relative to *https://localhost/piwebapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalysisTemplateCreateFromAnalysis**](AnalysisTemplateApi.md#AnalysisTemplateCreateFromAnalysis) | **Post** /analysistemplates | Create an Analysis template based upon a specified Analysis.
[**AnalysisTemplateCreateSecurityEntry**](AnalysisTemplateApi.md#AnalysisTemplateCreateSecurityEntry) | **Post** /analysistemplates/{webId}/securityentries | Create a security entry owned by the analysis template.
[**AnalysisTemplateDelete**](AnalysisTemplateApi.md#AnalysisTemplateDelete) | **Delete** /analysistemplates/{webId} | Delete an analysis template.
[**AnalysisTemplateDeleteSecurityEntry**](AnalysisTemplateApi.md#AnalysisTemplateDeleteSecurityEntry) | **Delete** /analysistemplates/{webId}/securityentries/{name} | Delete a security entry owned by the analysis template.
[**AnalysisTemplateGet**](AnalysisTemplateApi.md#AnalysisTemplateGet) | **Get** /analysistemplates/{webId} | Retrieve an analysis template.
[**AnalysisTemplateGetAnalysisTemplatesQuery**](AnalysisTemplateApi.md#AnalysisTemplateGetAnalysisTemplatesQuery) | **Get** /analysistemplates/search | Retrieve analysis templates based on the specified conditions. By default, returns all analysis templates.
[**AnalysisTemplateGetByPath**](AnalysisTemplateApi.md#AnalysisTemplateGetByPath) | **Get** /analysistemplates | Retrieve an analysis template by path.
[**AnalysisTemplateGetCategories**](AnalysisTemplateApi.md#AnalysisTemplateGetCategories) | **Get** /analysistemplates/{webId}/categories | Get an analysis template&#39;s categories.
[**AnalysisTemplateGetSecurity**](AnalysisTemplateApi.md#AnalysisTemplateGetSecurity) | **Get** /analysistemplates/{webId}/security | Get the security information of the specified security item associated with the analysis template for a specified user.
[**AnalysisTemplateGetSecurityEntries**](AnalysisTemplateApi.md#AnalysisTemplateGetSecurityEntries) | **Get** /analysistemplates/{webId}/securityentries | Retrieve the security entries associated with the analysis template based on the specified criteria. By default, all security entries for this analysis template are returned.
[**AnalysisTemplateGetSecurityEntryByName**](AnalysisTemplateApi.md#AnalysisTemplateGetSecurityEntryByName) | **Get** /analysistemplates/{webId}/securityentries/{name} | Retrieve the security entry associated with the analysis template with the specified name.
[**AnalysisTemplateUpdate**](AnalysisTemplateApi.md#AnalysisTemplateUpdate) | **Patch** /analysistemplates/{webId} | Update an analysis template by replacing items in its definition.
[**AnalysisTemplateUpdateSecurityEntry**](AnalysisTemplateApi.md#AnalysisTemplateUpdateSecurityEntry) | **Put** /analysistemplates/{webId}/securityentries/{name} | Update a security entry owned by the analysis template.


# **AnalysisTemplateCreateFromAnalysis**
> AnalysisTemplateCreateFromAnalysis(ctx, analysisWebId, optional)
Create an Analysis template based upon a specified Analysis.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **analysisWebId** | **string**| The ID of the Analysis, on which the template is created. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analysisWebId** | **string**| The ID of the Analysis, on which the template is created. | 
 **name** | **string**| The name for the created template, which must be unique within the database&#39;s AnalysisTemplate collection. If the name ends with an asterisk (*), then a unique name will be generated based on the supplied name. The default is the specified Analysis&#39; name suffixed with an asterisk (*). | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AnalysisTemplateCreateSecurityEntry**
> AnalysisTemplateCreateSecurityEntry(ctx, webId, securityEntry, optional)
Create a security entry owned by the analysis template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the analysis template, where the security entry will be created. | 
  **securityEntry** | [**SecurityEntry**](SecurityEntry.md)| The new security entry definition. The full list of allow and deny rights must be supplied. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the analysis template, where the security entry will be created. | 
 **securityEntry** | [**SecurityEntry**](SecurityEntry.md)| The new security entry definition. The full list of allow and deny rights must be supplied. | 
 **applyToChildren** | **bool**| If false, the new access permissions are only applied to the associated object. If true, the access permissions of children with any parent-child reference types will change when the permissions on the primary parent change. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AnalysisTemplateDelete**
> AnalysisTemplateDelete(ctx, webId)
Delete an analysis template.

Deleting an analysis template will delete any analysis which was created from it, unless the analysis is tied to a notification.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the analysis template to update. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AnalysisTemplateDeleteSecurityEntry**
> AnalysisTemplateDeleteSecurityEntry(ctx, name, webId, optional)
Delete a security entry owned by the analysis template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The name of the security entry. For every backslash character (\\) in the security entry name, replace with asterisk (*). As an example, use domain*username instead of domain\\username. | 
  **webId** | **string**| The ID of the analysis template, where the security entry will be deleted. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The name of the security entry. For every backslash character (\\) in the security entry name, replace with asterisk (*). As an example, use domain*username instead of domain\\username. | 
 **webId** | **string**| The ID of the analysis template, where the security entry will be deleted. | 
 **applyToChildren** | **bool**| If false, the new access permissions are only applied to the associated object. If true, the access permissions of children with any parent-child reference types will change when the permissions on the primary parent change. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AnalysisTemplateGet**
> AnalysisTemplate AnalysisTemplateGet(ctx, webId, optional)
Retrieve an analysis template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the analysis template. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the analysis template. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**AnalysisTemplate**](AnalysisTemplate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AnalysisTemplateGetAnalysisTemplatesQuery**
> ItemsAnalysisTemplate AnalysisTemplateGetAnalysisTemplatesQuery(ctx, optional)
Retrieve analysis templates based on the specified conditions. By default, returns all analysis templates.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **databaseWebId** | **string**| The ID of the asset database to use as the root of the query. | 
 **maxCount** | **int32**| The maximum number of objects to be returned per call (page size). The default is 1000. | 
 **query** | **string**| The query string is a list of filters used to perform an AFSearch for the analyses in the asset database. An example would be: \&quot;query&#x3D; Name:&#x3D;MyAnalysisTemplate1*\&quot;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **startIndex** | **int32**| The starting index (zero based) of the items to be returned. The default is 0. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsAnalysisTemplate**](Items[AnalysisTemplate].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AnalysisTemplateGetByPath**
> AnalysisTemplate AnalysisTemplateGetByPath(ctx, path, optional)
Retrieve an analysis template by path.

This method returns an analysis template based on the hierarchical path associated with it, and should be used when a path has been received from a separate part of the PI System for use in the PI Web API. Users should primarily search with the WebID when available.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **path** | **string**| The path to the analysis template. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| The path to the analysis template. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**AnalysisTemplate**](AnalysisTemplate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AnalysisTemplateGetCategories**
> ItemsAnalysisCategory AnalysisTemplateGetCategories(ctx, webId, optional)
Get an analysis template's categories.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the analysis template. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the analysis template. | 
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

# **AnalysisTemplateGetSecurity**
> ItemsSecurityRights AnalysisTemplateGetSecurity(ctx, webId, userIdentity, optional)
Get the security information of the specified security item associated with the analysis template for a specified user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the analysis template for the security to be checked. | 
  **userIdentity** | [**[]string**](string.md)| The user identity for the security information to be checked. Multiple security identities may be specified with multiple instances of the parameter. If the parameter is not specified, only the current user&#39;s security rights will be returned. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the analysis template for the security to be checked. | 
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

# **AnalysisTemplateGetSecurityEntries**
> ItemsSecurityEntry AnalysisTemplateGetSecurityEntries(ctx, webId, optional)
Retrieve the security entries associated with the analysis template based on the specified criteria. By default, all security entries for this analysis template are returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the analysis template. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the analysis template. | 
 **nameFilter** | **string**| The name query string used for filtering security entries. The default is no filter. | 
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

# **AnalysisTemplateGetSecurityEntryByName**
> SecurityEntry AnalysisTemplateGetSecurityEntryByName(ctx, name, webId, optional)
Retrieve the security entry associated with the analysis template with the specified name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The name of the security entry. For every backslash character (\\) in the security entry name, replace with asterisk (*). As an example, use domain*username instead of domain\\username. | 
  **webId** | **string**| The ID of the analysis template. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The name of the security entry. For every backslash character (\\) in the security entry name, replace with asterisk (*). As an example, use domain*username instead of domain\\username. | 
 **webId** | **string**| The ID of the analysis template. | 
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

# **AnalysisTemplateUpdate**
> AnalysisTemplateUpdate(ctx, webId, template)
Update an analysis template by replacing items in its definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the analysis template to update. | 
  **template** | [**AnalysisTemplate**](AnalysisTemplate.md)| A partial analysis template containing the desired changes. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AnalysisTemplateUpdateSecurityEntry**
> AnalysisTemplateUpdateSecurityEntry(ctx, name, webId, securityEntry, optional)
Update a security entry owned by the analysis template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The name of the security entry. | 
  **webId** | **string**| The ID of the analysis template, where the security entry will be updated. | 
  **securityEntry** | [**SecurityEntry**](SecurityEntry.md)| The new security entry definition. The full list of allow and deny rights must be supplied or they will be removed. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The name of the security entry. | 
 **webId** | **string**| The ID of the analysis template, where the security entry will be updated. | 
 **securityEntry** | [**SecurityEntry**](SecurityEntry.md)| The new security entry definition. The full list of allow and deny rights must be supplied or they will be removed. | 
 **applyToChildren** | **bool**| If false, the new access permissions are only applied to the associated object. If true, the access permissions of children with any parent-child reference types will change when the permissions on the primary parent change. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

