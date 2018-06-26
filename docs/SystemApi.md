# \SystemApi

All URIs are relative to *https://localhost/piwebapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SystemCacheInstances**](SystemApi.md#SystemCacheInstances) | **Get** /system/cacheinstances | Get AF cache instances currently in use by the system. These are caches from which user requests are serviced. The number of instances depends on the number of users connected to the service, the service&#39;s authentication method, and the cache instance configuration.
[**SystemLanding**](SystemApi.md#SystemLanding) | **Get** /system | Get system links for this PI System Web API instance.
[**SystemStatus**](SystemApi.md#SystemStatus) | **Get** /system/status | Get information about this PI Web API instance. Examples of information returned include the system uptime, the number of cache instances for this PI System Web API instance, and the system run state.
[**SystemUserInfo**](SystemApi.md#SystemUserInfo) | **Get** /system/userinfo | Get information about the Windows identity used to fulfill the request. This depends on the service&#39;s authentication method and the credentials passed by the client. The impersonation level of the Windows identity is included.
[**SystemVersions**](SystemApi.md#SystemVersions) | **Get** /system/versions | Get the current versions of the PI Web API instance and all external plugins.


# **SystemCacheInstances**
> ItemsCacheInstance SystemCacheInstances(ctx, )
Get AF cache instances currently in use by the system. These are caches from which user requests are serviced. The number of instances depends on the number of users connected to the service, the service's authentication method, and the cache instance configuration.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ItemsCacheInstance**](Items[CacheInstance].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemLanding**
> SystemLanding SystemLanding(ctx, )
Get system links for this PI System Web API instance.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SystemLanding**](SystemLanding.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemStatus**
> SystemStatus SystemStatus(ctx, )
Get information about this PI Web API instance. Examples of information returned include the system uptime, the number of cache instances for this PI System Web API instance, and the system run state.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SystemStatus**](SystemStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemUserInfo**
> UserInfo SystemUserInfo(ctx, )
Get information about the Windows identity used to fulfill the request. This depends on the service's authentication method and the credentials passed by the client. The impersonation level of the Windows identity is included.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UserInfo**](UserInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemVersions**
> map[string]Version SystemVersions(ctx, )
Get the current versions of the PI Web API instance and all external plugins.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**map[string]Version**](Version.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

