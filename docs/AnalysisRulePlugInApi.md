# \AnalysisRulePlugInApi

All URIs are relative to *https://localhost/piwebapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalysisRulePlugInGet**](AnalysisRulePlugInApi.md#AnalysisRulePlugInGet) | **Get** /analysisruleplugins/{webId} | Retrieve an Analysis Rule Plug-in.
[**AnalysisRulePlugInGetByPath**](AnalysisRulePlugInApi.md#AnalysisRulePlugInGetByPath) | **Get** /analysisruleplugins | Retrieve an Analysis Rule Plug-in by path.


# **AnalysisRulePlugInGet**
> AnalysisRulePlugIn AnalysisRulePlugInGet(ctx, webId, optional)
Retrieve an Analysis Rule Plug-in.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the Analysis Rule Plug-in. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the Analysis Rule Plug-in. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**AnalysisRulePlugIn**](AnalysisRulePlugIn.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AnalysisRulePlugInGetByPath**
> AnalysisRulePlugIn AnalysisRulePlugInGetByPath(ctx, path, optional)
Retrieve an Analysis Rule Plug-in by path.

This method returns an Analysis Rule Plug-in based on the hierarchical path associated with it, and should be used when a path has been received from a separate part of the PI System for use in the PI Web API. Users should primarily search with the WebID when available.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **path** | **string**| The path to the Analysis Rule Plug-in. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| The path to the Analysis Rule Plug-in. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**AnalysisRulePlugIn**](AnalysisRulePlugIn.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

