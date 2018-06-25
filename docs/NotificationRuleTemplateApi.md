# \NotificationRuleTemplateApi

All URIs are relative to *https://localhost/piwebapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotificationRuleTemplateGet**](NotificationRuleTemplateApi.md#NotificationRuleTemplateGet) | **Get** /notificationruletemplates/{webId} | Get the specified notification rule template.
[**NotificationRuleTemplateGetNotificationRuleTemplateSubscribers**](NotificationRuleTemplateApi.md#NotificationRuleTemplateGetNotificationRuleTemplateSubscribers) | **Get** /notificationruletemplates/{webId}/notificationrulesubscribers | Retrieve notification rule template subscribers.
[**NotificationRuleTemplateGetNotificationRuleTemplatesQuery**](NotificationRuleTemplateApi.md#NotificationRuleTemplateGetNotificationRuleTemplatesQuery) | **Get** /notificationruletemplates/search | Retrieve Notification rule templates based on the specified conditions. Returns Notification rule templates using the specified search query string.


# **NotificationRuleTemplateGet**
> NotificationRuleTemplate NotificationRuleTemplateGet(ctx, webId, optional)
Get the specified notification rule template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the notification rule template. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the notification rule template. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**NotificationRuleTemplate**](NotificationRuleTemplate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NotificationRuleTemplateGetNotificationRuleTemplateSubscribers**
> ItemsNotificationRuleSubscriber NotificationRuleTemplateGetNotificationRuleTemplateSubscribers(ctx, webId, optional)
Retrieve notification rule template subscribers.

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
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsNotificationRuleSubscriber**](Items[NotificationRuleSubscriber].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NotificationRuleTemplateGetNotificationRuleTemplatesQuery**
> ItemsNotificationRuleTemplate NotificationRuleTemplateGetNotificationRuleTemplatesQuery(ctx, optional)
Retrieve Notification rule templates based on the specified conditions. Returns Notification rule templates using the specified search query string.

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
 **query** | **string**| The query string is a list of filters used to perform an AFSearch for the Notification rule templates in the asset database. An example would be: \&quot;query&#x3D;NotificationRuleTemplate:{ Name:&#x3D;&#39;MyNotificationRuleTemplate&#39; } Type:&#x3D;Int32\&quot;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **startIndex** | **int32**| The starting index (zero based) of the items to be returned. The default is 0. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsNotificationRuleTemplate**](Items[NotificationRuleTemplate].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

