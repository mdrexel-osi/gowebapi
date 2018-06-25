# \ConfigurationApi

All URIs are relative to *https://localhost/piwebapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigurationDelete**](ConfigurationApi.md#ConfigurationDelete) | **Delete** /system/configuration/{key} | Delete a configuration item.
[**ConfigurationGet**](ConfigurationApi.md#ConfigurationGet) | **Get** /system/configuration/{key} | Get the value of a configuration item.
[**ConfigurationList**](ConfigurationApi.md#ConfigurationList) | **Get** /system/configuration | Get the current system configuration.


# **ConfigurationDelete**
> ConfigurationDelete(ctx, key)
Delete a configuration item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **key** | **string**| The key of the configuration item to remove. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigurationGet**
> interface{} ConfigurationGet(ctx, key)
Get the value of a configuration item.

The response content may vary based on the configuration item's data type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **key** | **string**| The key of the configuration item. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigurationList**
> map[string]interface{} ConfigurationList(ctx, )
Get the current system configuration.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

