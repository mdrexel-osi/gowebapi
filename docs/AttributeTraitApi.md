# \AttributeTraitApi

All URIs are relative to *https://localhost/piwebapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttributeTraitGet**](AttributeTraitApi.md#AttributeTraitGet) | **Get** /attributetraits/{name} | Retrieve an attribute trait.
[**AttributeTraitGetByCategory**](AttributeTraitApi.md#AttributeTraitGetByCategory) | **Get** /attributetraits | Retrieve all attribute traits of the specified category/categories.


# **AttributeTraitGet**
> AttributeTrait AttributeTraitGet(ctx, name, optional)
Retrieve an attribute trait.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The name or abbreviation of the attribute trait. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The name or abbreviation of the attribute trait. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 

### Return type

[**AttributeTrait**](AttributeTrait.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AttributeTraitGetByCategory**
> ItemsAttributeTrait AttributeTraitGetByCategory(ctx, category, optional)
Retrieve all attribute traits of the specified category/categories.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **category** | [**[]string**](string.md)| The category of the attribute traits. Multiple categories may be specified with multiple instances of the parameter. If the parameter is not specified, or if its value is \&quot;all\&quot;, then all attribute traits of all categories will be returned. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | [**[]string**](string.md)| The category of the attribute traits. Multiple categories may be specified with multiple instances of the parameter. If the parameter is not specified, or if its value is \&quot;all\&quot;, then all attribute traits of all categories will be returned. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 

### Return type

[**ItemsAttributeTrait**](Items[AttributeTrait].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

