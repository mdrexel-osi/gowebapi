# \DataServerApi

All URIs are relative to *https://localhost/piwebapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataServerCreateEnumerationSet**](DataServerApi.md#DataServerCreateEnumerationSet) | **Post** /dataservers/{webId}/enumerationsets | Create an enumeration set on the Data Server.
[**DataServerCreatePoint**](DataServerApi.md#DataServerCreatePoint) | **Post** /dataservers/{webId}/points | Create a point in the specified Data Server.
[**DataServerGet**](DataServerApi.md#DataServerGet) | **Get** /dataservers/{webId} | Retrieve a Data Server.
[**DataServerGetByName**](DataServerApi.md#DataServerGetByName) | **Get** /dataservers#name | Retrieve a Data Server by name.
[**DataServerGetByPath**](DataServerApi.md#DataServerGetByPath) | **Get** /dataservers#path | Retrieve a Data Server by path.
[**DataServerGetEnumerationSets**](DataServerApi.md#DataServerGetEnumerationSets) | **Get** /dataservers/{webId}/enumerationsets | Retrieve enumeration sets for given Data Server.
[**DataServerGetLicense**](DataServerApi.md#DataServerGetLicense) | **Get** /dataservers/{webId}/license | Retrieves the specified license for the given Data Server. The fields of the response object are string representations of the numerical values reported by the Data Server, with \&quot;Infinity\&quot; representing a license field with no limit.
[**DataServerGetPoints**](DataServerApi.md#DataServerGetPoints) | **Get** /dataservers/{webId}/points | Retrieve a list of points on a specified Data Server.
[**DataServerList**](DataServerApi.md#DataServerList) | **Get** /dataservers | Retrieve a list of Data Servers known to this service.


# **DataServerCreateEnumerationSet**
> DataServerCreateEnumerationSet(ctx, webId, enumerationSet, optional)
Create an enumeration set on the Data Server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the server on which to create the enumeration set. | 
  **enumerationSet** | [**EnumerationSet**](EnumerationSet.md)| The new enumeration set definition. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the server on which to create the enumeration set. | 
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

# **DataServerCreatePoint**
> DataServerCreatePoint(ctx, webId, pointDTO, optional)
Create a point in the specified Data Server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the server. | 
  **pointDTO** | [**Point**](Point.md)| The new point definition. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the server. | 
 **pointDTO** | [**Point**](Point.md)| The new point definition. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DataServerGet**
> DataServer DataServerGet(ctx, webId, optional)
Retrieve a Data Server.

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

[**DataServer**](DataServer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DataServerGetByName**
> DataServer DataServerGetByName(ctx, name, optional)
Retrieve a Data Server by name.

This method returns a data server based on the name. Users should primarily search with the WebID when available.

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

[**DataServer**](DataServer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DataServerGetByPath**
> DataServer DataServerGetByPath(ctx, path, optional)
Retrieve a Data Server by path.

This method returns a data server based on the hierarchical path associated with it, and should be used when a path has been received from a separate part of the PI System for use in the PI Web API. Users should primarily search with the WebID when available.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **path** | **string**| The path to the server. Note that the path supplied to this method must be of the form &#39;\\\\servername&#39;. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| The path to the server. Note that the path supplied to this method must be of the form &#39;\\\\servername&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**DataServer**](DataServer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DataServerGetEnumerationSets**
> ItemsEnumerationSet DataServerGetEnumerationSets(ctx, webId, optional)
Retrieve enumeration sets for given Data Server.

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

[**ItemsEnumerationSet**](Items[EnumerationSet].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DataServerGetLicense**
> DataServerLicense DataServerGetLicense(ctx, webId, optional)
Retrieves the specified license for the given Data Server. The fields of the response object are string representations of the numerical values reported by the Data Server, with \"Infinity\" representing a license field with no limit.

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
 **module** | **string**| The case-sensitive name of the module. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**DataServerLicense**](DataServerLicense.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DataServerGetPoints**
> ItemsPoint DataServerGetPoints(ctx, webId, optional)
Retrieve a list of points on a specified Data Server.

Users can search for the data servers based on specific search parameters. If no parameters are specified in the search, the default values for each parameter will be used and will return the data servers that match the default search.

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
 **maxCount** | **int32**| The maximum number of objects to be returned per call (page size). The default is 1000. | 
 **nameFilter** | **string**| A query string for filtering by point name. The default is no filter. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **startIndex** | **int32**| The starting index (zero based) of the items to be returned. The default is &#39;0&#39;. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsPoint**](Items[Point].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DataServerList**
> ItemsDataServer DataServerList(ctx, optional)
Retrieve a list of Data Servers known to this service.

This method returns a list of all available known Data Servers that the user can connect to. Even though a server may be returned in the list, the user may not have permission to access it.

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

[**ItemsDataServer**](Items[DataServer].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

