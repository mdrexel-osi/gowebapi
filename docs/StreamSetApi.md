# \StreamSetApi

All URIs are relative to *https://localhost/piwebapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StreamSetGetChannel**](StreamSetApi.md#StreamSetGetChannel) | **Get** /streamsets/{webId}/channel | Opens a channel that will send messages about any value changes for the attributes of an Element, Event Frame, or Attribute.
[**StreamSetGetChannelAdHoc**](StreamSetApi.md#StreamSetGetChannelAdHoc) | **Get** /streamsets/channel | Opens a channel that will send messages about any value changes for the specified streams.
[**StreamSetGetEnd**](StreamSetApi.md#StreamSetGetEnd) | **Get** /streamsets/{webId}/end | Returns End of stream values of the attributes for an Element, Event Frame or Attribute
[**StreamSetGetEndAdHoc**](StreamSetApi.md#StreamSetGetEndAdHoc) | **Get** /streamsets/end | Returns End Of Stream values for attributes of the specified streams
[**StreamSetGetInterpolated**](StreamSetApi.md#StreamSetGetInterpolated) | **Get** /streamsets/{webId}/interpolated | Returns interpolated values of attributes for an element, event frame or attribute over the specified time range at the specified sampling interval.
[**StreamSetGetInterpolatedAdHoc**](StreamSetApi.md#StreamSetGetInterpolatedAdHoc) | **Get** /streamsets/interpolated | Returns interpolated values of the specified streams over the specified time range at the specified sampling interval.
[**StreamSetGetInterpolatedAtTimes**](StreamSetApi.md#StreamSetGetInterpolatedAtTimes) | **Get** /streamsets/{webId}/interpolatedattimes | Returns interpolated values of attributes for an element, event frame or attribute at the specified times.
[**StreamSetGetInterpolatedAtTimesAdHoc**](StreamSetApi.md#StreamSetGetInterpolatedAtTimesAdHoc) | **Get** /streamsets/interpolatedattimes | Returns interpolated values of the specified streams at the specified times.
[**StreamSetGetPlot**](StreamSetApi.md#StreamSetGetPlot) | **Get** /streamsets/{webId}/plot | Returns values of attributes for an element, event frame or attribute over the specified time range suitable for plotting over the number of intervals (typically represents pixels).
[**StreamSetGetPlotAdHoc**](StreamSetApi.md#StreamSetGetPlotAdHoc) | **Get** /streamsets/plot | Returns values of attributes for the specified streams over the specified time range suitable for plotting over the number of intervals (typically represents pixels).
[**StreamSetGetRecorded**](StreamSetApi.md#StreamSetGetRecorded) | **Get** /streamsets/{webId}/recorded | Returns recorded values of the attributes for an element, event frame, or attribute.
[**StreamSetGetRecordedAdHoc**](StreamSetApi.md#StreamSetGetRecordedAdHoc) | **Get** /streamsets/recorded | Returns recorded values of the specified streams.
[**StreamSetGetRecordedAtTime**](StreamSetApi.md#StreamSetGetRecordedAtTime) | **Get** /streamsets/{webId}/recordedattime | Returns recorded values of the attributes for an element, event frame, or attribute.
[**StreamSetGetRecordedAtTimeAdHoc**](StreamSetApi.md#StreamSetGetRecordedAtTimeAdHoc) | **Get** /streamsets/recordedattime | Returns recorded values based on the passed time and retrieval mode.
[**StreamSetGetRecordedAtTimes**](StreamSetApi.md#StreamSetGetRecordedAtTimes) | **Get** /streamsets/{webId}/recordedattimes | Returns recorded values of attributes for an element, event frame or attribute at the specified times.
[**StreamSetGetRecordedAtTimesAdHoc**](StreamSetApi.md#StreamSetGetRecordedAtTimesAdHoc) | **Get** /streamsets/recordedattimes | Returns recorded values of the specified streams at the specified times.
[**StreamSetGetSummaries**](StreamSetApi.md#StreamSetGetSummaries) | **Get** /streamsets/{webId}/summary | Returns summary values of the attributes for an element, event frame or attribute.
[**StreamSetGetSummariesAdHoc**](StreamSetApi.md#StreamSetGetSummariesAdHoc) | **Get** /streamsets/summary | Returns summary values of the specified streams.
[**StreamSetGetValues**](StreamSetApi.md#StreamSetGetValues) | **Get** /streamsets/{webId}/value | Returns values of the attributes for an Element, Event Frame or Attribute at the specified time.
[**StreamSetGetValuesAdHoc**](StreamSetApi.md#StreamSetGetValuesAdHoc) | **Get** /streamsets/value | Returns values of the specified streams.
[**StreamSetUpdateValue**](StreamSetApi.md#StreamSetUpdateValue) | **Post** /streamsets/{webId}/value | Updates a single value for the specified streams.
[**StreamSetUpdateValueAdHoc**](StreamSetApi.md#StreamSetUpdateValueAdHoc) | **Post** /streamsets/value | Updates a single value for the specified streams.
[**StreamSetUpdateValues**](StreamSetApi.md#StreamSetUpdateValues) | **Post** /streamsets/{webId}/recorded | Updates multiple values for the specified streams.
[**StreamSetUpdateValuesAdHoc**](StreamSetApi.md#StreamSetUpdateValuesAdHoc) | **Post** /streamsets/recorded | Updates multiple values for the specified streams.


# **StreamSetGetChannel**
> StreamSetGetChannel(ctx, webId, optional)
Opens a channel that will send messages about any value changes for the attributes of an Element, Event Frame, or Attribute.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of an Element, Event Frame or Attribute, which is the base element or parent of all the stream attributes. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of an Element, Event Frame or Attribute, which is the base element or parent of all the stream attributes. | 
 **categoryName** | **string**| Specify that included attributes must have this category. The default is no category filter. | 
 **heartbeatRate** | **int32**| Specifies the maximum number of consecutive empty messages that can be elapsed with no new data updates from the PI System, after which the client receives an empty payload. It helps to check if the connection is still alive. Zero/negative values correspond to no heartbeat, and the default value is no heartbeat. | 
 **includeInitialValues** | **bool**| Specified if the channel should send a message with the current values of all the streams after the connection is opened. The default is &#39;false&#39;. | 
 **nameFilter** | **string**| The name query string used for filtering attributes. The default is no filter. | 
 **searchFullHierarchy** | **bool**| Specifies if the search should include attributes nested further than the immediate attributes of the searchRoot. The default is &#39;false&#39;. | 
 **showExcluded** | **bool**| Specified if the search should include attributes with the Excluded property set. The default is &#39;false&#39;. | 
 **showHidden** | **bool**| Specified if the search should include attributes with the Hidden property set. The default is &#39;false&#39;. | 
 **templateName** | **string**| Specify that included attributes must be members of this template. The default is no template filter. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamSetGetChannelAdHoc**
> StreamSetGetChannelAdHoc(ctx, webId, optional)
Opens a channel that will send messages about any value changes for the specified streams.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | [**[]string**](string.md)| The ID of a stream. Multiple streams may be specified with multiple instances of the parameter. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | [**[]string**](string.md)| The ID of a stream. Multiple streams may be specified with multiple instances of the parameter. | 
 **heartbeatRate** | **int32**| Specifies the maximum number of consecutive empty messages that can be elapsed with no new data updates from the PI System, after which the client receives an empty payload. It helps to check if the connection is still alive. Zero/negative values correspond to no heartbeat, and the default value is no heartbeat. | 
 **includeInitialValues** | **bool**| Specified if the channel should send a message with the current values of all the streams after the connection is opened. The default is &#39;false&#39;. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamSetGetEnd**
> ItemsStreamValue StreamSetGetEnd(ctx, webId, optional)
Returns End of stream values of the attributes for an Element, Event Frame or Attribute

Any time series value in the response that contains an 'Errors' property indicates PI Web API encountered a handled error during the transfer of the response stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of an Element, Event Frame or Attribute, which is the base element or parent of all the stream attributes. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of an Element, Event Frame or Attribute, which is the base element or parent of all the stream attributes. | 
 **categoryName** | **string**| Specify that included attributes must have this category. The default is no category filter. | 
 **nameFilter** | **string**| The name query string used for filtering attributes. The default is no filter. | 
 **searchFullHierarchy** | **bool**| Specifies if the search should include attributes nested further than the immediate attributes of the searchRoot. The default is &#39;false&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **showExcluded** | **bool**| Specified if the search should include attributes with the Excluded property set. The default is &#39;false&#39;. | 
 **showHidden** | **bool**| Specified if the search should include attributes with the Hidden property set. The default is &#39;false&#39;. | 
 **sortField** | **string**| The field or property of the object used to sort the returned collection. For better performance, by default no sorting is applied. &#39;Name&#39; is the only supported field by which to sort. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39; | 
 **templateName** | **string**| Specify that included attributes must be members of this template. The default is no template filter. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsStreamValue**](Items[StreamValue].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamSetGetEndAdHoc**
> ItemsStreamValues StreamSetGetEndAdHoc(ctx, webId, optional)
Returns End Of Stream values for attributes of the specified streams

Any time series value in the response that contains an 'Errors' property indicates PI Web API encountered a handled error during the transfer of the response stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | [**[]string**](string.md)| The ID of a stream. Multiple streams may be specified with multiple instances of the parameter. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | [**[]string**](string.md)| The ID of a stream. Multiple streams may be specified with multiple instances of the parameter. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **sortField** | **string**| The field or property of the object used to sort the returned collection. For better performance, by default no sorting is applied. &#39;Name&#39; is the only supported field by which to sort. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39; | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsStreamValues**](Items[StreamValues].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamSetGetInterpolated**
> ItemsStreamValues StreamSetGetInterpolated(ctx, webId, optional)
Returns interpolated values of attributes for an element, event frame or attribute over the specified time range at the specified sampling interval.

Any time series value in the response that contains an 'Errors' property indicates PI Web API encountered a handled error during the transfer of the response stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of an element, event frame or attribute, which is the base element or parent of all the stream attributes. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of an element, event frame or attribute, which is the base element or parent of all the stream attributes. | 
 **categoryName** | **string**| Specify that included attributes must have this category. The default is no category filter. | 
 **endTime** | **string**| An optional end time. The default is &#39;*&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s end time, or &#39;*&#39; if that is not set. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order. | 
 **filterExpression** | **string**| An optional string containing a filter expression. Expression variables are relative to the data point. Use &#39;.&#39; to reference the containing attribute. If the attribute does not support filtering, the filter will be ignored. The default is no filtering. | 
 **includeFilteredValues** | **bool**| Specify &#39;true&#39; to indicate that values which fail the filter criteria are present in the returned data at the times where they occurred with a value set to a &#39;Filtered&#39; enumeration value with bad status. Repeated consecutive failures are omitted. | 
 **interval** | **string**| The sampling interval, in AFTimeSpan format. | 
 **nameFilter** | **string**| The name query string used for filtering attributes. The default is no filter. | 
 **searchFullHierarchy** | **bool**| Specifies if the search should include attributes nested further than the immediate attributes of the searchRoot. The default is &#39;false&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **showExcluded** | **bool**| Specified if the search should include attributes with the Excluded property set. The default is &#39;false&#39;. | 
 **showHidden** | **bool**| Specified if the search should include attributes with the Hidden property set. The default is &#39;false&#39;. | 
 **sortField** | **string**| The field or property of the object used to sort the returned collection. For better performance, by default no sorting is applied. &#39;Name&#39; is the only supported field by which to sort. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39;. | 
 **startTime** | **string**| An optional start time. The default is &#39;*-1d&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s start time, or &#39;*-1d&#39; if that is not set. | 
 **syncTime** | **string**| An optional start time anchor, in AFTime format. When specified, interpolated data retrieval will use the sync time as the origin for calculating the interval times. | 
 **syncTimeBoundaryType** | **string**| An optional string specifying the boundary type to use when applying a syncTime. The allowed values are &#39;Inside&#39; and &#39;Outside&#39;. The default is &#39;Inside&#39;. | 
 **templateName** | **string**| Specify that included attributes must be members of this template. The default is no template filter. | 
 **timeZone** | **string**| The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsStreamValues**](Items[StreamValues].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamSetGetInterpolatedAdHoc**
> ItemsStreamValues StreamSetGetInterpolatedAdHoc(ctx, webId, optional)
Returns interpolated values of the specified streams over the specified time range at the specified sampling interval.

Any time series value in the response that contains an 'Errors' property indicates PI Web API encountered a handled error during the transfer of the response stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | [**[]string**](string.md)| The ID of a stream. Multiple streams may be specified with multiple instances of the parameter. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | [**[]string**](string.md)| The ID of a stream. Multiple streams may be specified with multiple instances of the parameter. | 
 **endTime** | **string**| An optional end time. The default is &#39;*&#39;. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order. | 
 **filterExpression** | **string**| An optional string containing a filter expression. Expression variables are relative to the data point. Use &#39;.&#39; to reference the containing attribute. If the attribute does not support filtering, the filter will be ignored. The default is no filtering. | 
 **includeFilteredValues** | **bool**| Specify &#39;true&#39; to indicate that values which fail the filter criteria are present in the returned data at the times where they occurred with a value set to a &#39;Filtered&#39; enumeration value with bad status. Repeated consecutive failures are omitted. | 
 **interval** | **string**| The sampling interval, in AFTimeSpan format. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **sortField** | **string**| The field or property of the object used to sort the returned collection. For better performance, by default no sorting is applied. &#39;Name&#39; is the only supported field by which to sort. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39; | 
 **startTime** | **string**| An optional start time. The default is &#39;*-1d&#39;. | 
 **syncTime** | **string**| An optional start time anchor, in AFTime format. When specified, interpolated data retrieval will use the sync time as the origin for calculating the interval times. | 
 **syncTimeBoundaryType** | **string**| An optional string specifying the boundary type to use when applying a syncTime. The allowed values are &#39;Inside&#39; and &#39;Outside&#39;. The default is &#39;Inside&#39;. | 
 **timeZone** | **string**| The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsStreamValues**](Items[StreamValues].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamSetGetInterpolatedAtTimes**
> ItemsStreamValues StreamSetGetInterpolatedAtTimes(ctx, webId, time, optional)
Returns interpolated values of attributes for an element, event frame or attribute at the specified times.

Any time series value in the response that contains an 'Errors' property indicates PI Web API encountered a handled error during the transfer of the response stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of an element, event frame or attribute, which is the base element or parent of all the stream attributes. | 
  **time** | [**[]string**](string.md)| The timestamp at which to retrieve a interpolated value. Multiple timestamps may be specified with multiple instances of the parameter. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of an element, event frame or attribute, which is the base element or parent of all the stream attributes. | 
 **time** | [**[]string**](string.md)| The timestamp at which to retrieve a interpolated value. Multiple timestamps may be specified with multiple instances of the parameter. | 
 **categoryName** | **string**| Specify that included attributes must have this category. The default is no category filter. | 
 **filterExpression** | **string**| An optional string containing a filter expression. Expression variables are relative to the data point. Use &#39;.&#39; to reference the containing attribute. If the attribute does not support filtering, the filter will be ignored. The default is no filtering. | 
 **includeFilteredValues** | **bool**| Specify &#39;true&#39; to indicate that values which fail the filter criteria are present in the returned data at the times where they occurred with a value set to a &#39;Filtered&#39; enumeration value with bad status. Repeated consecutive failures are omitted. | 
 **nameFilter** | **string**| The name query string used for filtering attributes. The default is no filter. | 
 **searchFullHierarchy** | **bool**| Specifies if the search should include attributes nested further than the immediate attributes of the searchRoot. The default is &#39;false&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **showExcluded** | **bool**| Specified if the search should include attributes with the Excluded property set. The default is &#39;false&#39;. | 
 **showHidden** | **bool**| Specified if the search should include attributes with the Hidden property set. The default is &#39;false&#39;. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39;. | 
 **templateName** | **string**| Specify that included attributes must be members of this template. The default is no template filter. | 
 **timeZone** | **string**| The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsStreamValues**](Items[StreamValues].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamSetGetInterpolatedAtTimesAdHoc**
> ItemsStreamValues StreamSetGetInterpolatedAtTimesAdHoc(ctx, time, webId, optional)
Returns interpolated values of the specified streams at the specified times.

Any time series value in the response that contains an 'Errors' property indicates PI Web API encountered a handled error during the transfer of the response stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **time** | [**[]string**](string.md)| The timestamp at which to retrieve a interpolated value. Multiple timestamps may be specified with multiple instances of the parameter. | 
  **webId** | [**[]string**](string.md)| The ID of a stream. Multiple streams may be specified with multiple instances of the parameter. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **time** | [**[]string**](string.md)| The timestamp at which to retrieve a interpolated value. Multiple timestamps may be specified with multiple instances of the parameter. | 
 **webId** | [**[]string**](string.md)| The ID of a stream. Multiple streams may be specified with multiple instances of the parameter. | 
 **filterExpression** | **string**| An optional string containing a filter expression. Expression variables are relative to the data point. Use &#39;.&#39; to reference the containing attribute. If the attribute does not support filtering, the filter will be ignored. The default is no filtering. | 
 **includeFilteredValues** | **bool**| Specify &#39;true&#39; to indicate that values which fail the filter criteria are present in the returned data at the times where they occurred with a value set to a &#39;Filtered&#39; enumeration value with bad status. Repeated consecutive failures are omitted. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39;. | 
 **timeZone** | **string**| The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsStreamValues**](Items[StreamValues].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamSetGetPlot**
> ItemsStreamValues StreamSetGetPlot(ctx, webId, optional)
Returns values of attributes for an element, event frame or attribute over the specified time range suitable for plotting over the number of intervals (typically represents pixels).

For each interval, the data available is examined and significant values are returned. Each interval can produce up to 5 values if they are unique, the first value in the interval, the last value, the highest value, the lowest value and at most one exceptional point (bad status or digital state). Any time series value in the response that contains an 'Errors' property indicates PI Web API encountered a handled error during the transfer of the response stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of an element, event frame or attribute, which is the base element or parent of all the stream attributes. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of an element, event frame or attribute, which is the base element or parent of all the stream attributes. | 
 **categoryName** | **string**| Specify that included attributes must have this category. The default is no category filter. | 
 **endTime** | **string**| An optional end time. The default is &#39;*&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s end time, or &#39;*&#39; if that is not set. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order. | 
 **intervals** | **int32**| The number of intervals to plot over. Typically, this would be the number of horizontal pixels in the trend. The default is &#39;24&#39;. For each interval, the data available is examined and significant values are returned. Each interval can produce up to 5 values if they are unique, the first value in the interval, the last value, the highest value, the lowest value and at most one exceptional point (bad status or digital state). | 
 **nameFilter** | **string**| The name query string used for filtering attributes. The default is no filter. | 
 **searchFullHierarchy** | **bool**| Specifies if the search should include attributes nested further than the immediate attributes of the searchRoot. The default is &#39;false&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **showExcluded** | **bool**| Specified if the search should include attributes with the Excluded property set. The default is &#39;false&#39;. | 
 **showHidden** | **bool**| Specified if the search should include attributes with the Hidden property set. The default is &#39;false&#39;. | 
 **sortField** | **string**| The field or property of the object used to sort the returned collection. For better performance, by default no sorting is applied. &#39;Name&#39; is the only supported field by which to sort. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39; | 
 **startTime** | **string**| An optional start time. The default is &#39;*-1d&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s start time, or &#39;*-1d&#39; if that is not set. | 
 **templateName** | **string**| Specify that included attributes must be members of this template. The default is no template filter. | 
 **timeZone** | **string**| The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsStreamValues**](Items[StreamValues].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamSetGetPlotAdHoc**
> ItemsStreamValues StreamSetGetPlotAdHoc(ctx, webId, optional)
Returns values of attributes for the specified streams over the specified time range suitable for plotting over the number of intervals (typically represents pixels).

For each interval, the data available is examined and significant values are returned. Each interval can produce up to 5 values if they are unique, the first value in the interval, the last value, the highest value, the lowest value and at most one exceptional point (bad status or digital state). Any time series value in the response that contains an 'Errors' property indicates PI Web API encountered a handled error during the transfer of the response stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | [**[]string**](string.md)| The ID of a stream. Multiple streams may be specified with multiple instances of the parameter. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | [**[]string**](string.md)| The ID of a stream. Multiple streams may be specified with multiple instances of the parameter. | 
 **endTime** | **string**| An optional end time. The default is &#39;*&#39;. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order. | 
 **intervals** | **int32**| The number of intervals to plot over. Typically, this would be the number of horizontal pixels in the trend. The default is &#39;24&#39;. For each interval, the data available is examined and significant values are returned. Each interval can produce up to 5 values if they are unique, the first value in the interval, the last value, the highest value, the lowest value and at most one exceptional point (bad status or digital state). | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **sortField** | **string**| The field or property of the object used to sort the returned collection. For better performance, by default no sorting is applied. &#39;Name&#39; is the only supported field by which to sort. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39; | 
 **startTime** | **string**| An optional start time. The default is &#39;*-1d&#39;. | 
 **timeZone** | **string**| The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsStreamValues**](Items[StreamValues].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamSetGetRecorded**
> ItemsStreamValues StreamSetGetRecorded(ctx, webId, optional)
Returns recorded values of the attributes for an element, event frame, or attribute.

Any time series value in the response that contains an 'Errors' property indicates PI Web API encountered a handled error during the transfer of the response stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of an element, event frame or attribute, which is the base element or parent of all the stream attributes. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of an element, event frame or attribute, which is the base element or parent of all the stream attributes. | 
 **boundaryType** | **string**| An optional value that determines how the times and values of the returned end points are determined. The default is &#39;Inside&#39;. | 
 **categoryName** | **string**| Specify that included attributes must have this category. The default is no category filter. | 
 **endTime** | **string**| An optional end time. The default is &#39;*&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s end time, or &#39;*&#39; if that is not set. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order. | 
 **filterExpression** | **string**| An optional string containing a filter expression. Expression variables are relative to the data point. Use &#39;.&#39; to reference the containing attribute. The default is no filtering. | 
 **includeFilteredValues** | **bool**| Specify &#39;true&#39; to indicate that values which fail the filter criteria are present in the returned data at the times where they occurred with a value set to a &#39;Filtered&#39; enumeration value with bad status. Repeated consecutive failures are omitted. | 
 **maxCount** | **int32**| The maximum number of values to be returned. The default is 1000. | 
 **nameFilter** | **string**| The name query string used for filtering attributes. The default is no filter. | 
 **searchFullHierarchy** | **bool**| Specifies if the search should include attributes nested further than the immediate attributes of the searchRoot. The default is &#39;false&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **showExcluded** | **bool**| Specified if the search should include attributes with the Excluded property set. The default is &#39;false&#39;. | 
 **showHidden** | **bool**| Specified if the search should include attributes with the Hidden property set. The default is &#39;false&#39;. | 
 **sortField** | **string**| The field or property of the object used to sort the returned collection. For better performance, by default no sorting is applied. &#39;Name&#39; is the only supported field by which to sort. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39; | 
 **startTime** | **string**| An optional start time. The default is &#39;*-1d&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s start time, or &#39;*-1d&#39; if that is not set. | 
 **templateName** | **string**| Specify that included attributes must be members of this template. The default is no template filter. | 
 **timeZone** | **string**| The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsStreamValues**](Items[StreamValues].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamSetGetRecordedAdHoc**
> ItemsStreamValues StreamSetGetRecordedAdHoc(ctx, webId, optional)
Returns recorded values of the specified streams.

Any time series value in the response that contains an 'Errors' property indicates PI Web API encountered a handled error during the transfer of the response stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | [**[]string**](string.md)| The ID of a stream. Multiple streams may be specified with multiple instances of the parameter. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | [**[]string**](string.md)| The ID of a stream. Multiple streams may be specified with multiple instances of the parameter. | 
 **boundaryType** | **string**| An optional value that determines how the times and values of the returned end points are determined. The default is &#39;Inside&#39;. | 
 **endTime** | **string**| An optional end time. The default is &#39;*&#39;. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order. | 
 **filterExpression** | **string**| An optional string containing a filter expression. Expression variables are relative to the data point. Use &#39;.&#39; to reference the containing attribute. The default is no filtering. | 
 **includeFilteredValues** | **bool**| Specify &#39;true&#39; to indicate that values which fail the filter criteria are present in the returned data at the times where they occurred with a value set to a &#39;Filtered&#39; enumeration value with bad status. Repeated consecutive failures are omitted. | 
 **maxCount** | **int32**| The maximum number of values to be returned. The default is 1000. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **sortField** | **string**| The field or property of the object used to sort the returned collection. For better performance, by default no sorting is applied. &#39;Name&#39; is the only supported field by which to sort. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39; | 
 **startTime** | **string**| An optional start time. The default is &#39;*-1d&#39;. | 
 **timeZone** | **string**| The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsStreamValues**](Items[StreamValues].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamSetGetRecordedAtTime**
> ItemsStreamValue StreamSetGetRecordedAtTime(ctx, webId, time, optional)
Returns recorded values of the attributes for an element, event frame, or attribute.

Any time series value in the response that contains an 'Errors' property indicates PI Web API encountered a handled error during the transfer of the response stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of an element, event frame or attribute, which is the base element or parent of all the stream attributes. | 
  **time** | **string**| The timestamp at which the values are desired. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of an element, event frame or attribute, which is the base element or parent of all the stream attributes. | 
 **time** | **string**| The timestamp at which the values are desired. | 
 **categoryName** | **string**| Specify that included attributes must have this category. The default is no category filter. | 
 **nameFilter** | **string**| The name query string used for filtering attributes. The default is no filter. | 
 **retrievalMode** | **string**| An optional value that determines the values to return when values don&#39;t exist at the exact time specified. The default is &#39;Auto&#39;. | 
 **searchFullHierarchy** | **bool**| Specifies if the search should include attributes nested further than the immediate attributes of the searchRoot. The default is &#39;false&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **showExcluded** | **bool**| Specified if the search should include attributes with the Excluded property set. The default is &#39;false&#39;. | 
 **showHidden** | **bool**| Specified if the search should include attributes with the Hidden property set. The default is &#39;false&#39;. | 
 **templateName** | **string**| Specify that included attributes must be members of this template. The default is no template filter. | 
 **timeZone** | **string**| The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsStreamValue**](Items[StreamValue].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamSetGetRecordedAtTimeAdHoc**
> ItemsStreamValue StreamSetGetRecordedAtTimeAdHoc(ctx, time, webId, optional)
Returns recorded values based on the passed time and retrieval mode.

Any time series value in the response that contains an 'Errors' property indicates PI Web API encountered a handled error during the transfer of the response stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **time** | **string**| The timestamp at which the values are desired. | 
  **webId** | [**[]string**](string.md)| The ID of a stream. Multiple streams may be specified with multiple instances of the parameter. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **time** | **string**| The timestamp at which the values are desired. | 
 **webId** | [**[]string**](string.md)| The ID of a stream. Multiple streams may be specified with multiple instances of the parameter. | 
 **retrievalMode** | **string**| An optional value that determines the values to return when values don&#39;t exist at the exact time specified. The default is &#39;Auto&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **timeZone** | **string**| The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsStreamValue**](Items[StreamValue].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamSetGetRecordedAtTimes**
> ItemsStreamValues StreamSetGetRecordedAtTimes(ctx, webId, time, optional)
Returns recorded values of attributes for an element, event frame or attribute at the specified times.

Any time series value in the response that contains an 'Errors' property indicates PI Web API encountered a handled error during the transfer of the response stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of an element, event frame or attribute, which is the base element or parent of all the stream attributes. | 
  **time** | [**[]string**](string.md)| The timestamp at which to retrieve a recorded value. Multiple timestamps may be specified with multiple instances of the parameter. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of an element, event frame or attribute, which is the base element or parent of all the stream attributes. | 
 **time** | [**[]string**](string.md)| The timestamp at which to retrieve a recorded value. Multiple timestamps may be specified with multiple instances of the parameter. | 
 **categoryName** | **string**| Specify that included attributes must have this category. The default is no category filter. | 
 **nameFilter** | **string**| The name query string used for filtering attributes. The default is no filter. | 
 **retrievalMode** | **string**| An optional value that determines the values to return when values don&#39;t exist at the exact time specified. The default is &#39;Auto&#39;. | 
 **searchFullHierarchy** | **bool**| Specifies if the search should include attributes nested further than the immediate attributes of the searchRoot. The default is &#39;false&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **showExcluded** | **bool**| Specified if the search should include attributes with the Excluded property set. The default is &#39;false&#39;. | 
 **showHidden** | **bool**| Specified if the search should include attributes with the Hidden property set. The default is &#39;false&#39;. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39;. | 
 **templateName** | **string**| Specify that included attributes must be members of this template. The default is no template filter. | 
 **timeZone** | **string**| The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsStreamValues**](Items[StreamValues].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamSetGetRecordedAtTimesAdHoc**
> ItemsStreamValues StreamSetGetRecordedAtTimesAdHoc(ctx, time, webId, optional)
Returns recorded values of the specified streams at the specified times.

Any time series value in the response that contains an 'Errors' property indicates PI Web API encountered a handled error during the transfer of the response stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **time** | [**[]string**](string.md)| The timestamp at which to retrieve a recorded value. Multiple timestamps may be specified with multiple instances of the parameter. | 
  **webId** | [**[]string**](string.md)| The ID of a stream. Multiple streams may be specified with multiple instances of the parameter. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **time** | [**[]string**](string.md)| The timestamp at which to retrieve a recorded value. Multiple timestamps may be specified with multiple instances of the parameter. | 
 **webId** | [**[]string**](string.md)| The ID of a stream. Multiple streams may be specified with multiple instances of the parameter. | 
 **retrievalMode** | **string**| An optional value that determines the values to return when values don&#39;t exist at the exact time specified. The default is &#39;Auto&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39;. | 
 **timeZone** | **string**| The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsStreamValues**](Items[StreamValues].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamSetGetSummaries**
> ItemsStreamSummaries StreamSetGetSummaries(ctx, webId, optional)
Returns summary values of the attributes for an element, event frame or attribute.

Any time series value in the response that contains an 'Errors' property indicates PI Web API encountered a handled error during the transfer of the response stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of an element, event frame or attribute, which is the base element or parent of all the stream attributes. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of an element, event frame or attribute, which is the base element or parent of all the stream attributes. | 
 **calculationBasis** | **string**| Specifies the method of evaluating the data over the time range. The default is &#39;TimeWeighted&#39;. | 
 **categoryName** | **string**| Specify that included attributes must have this category. The default is no category filter. | 
 **endTime** | **string**| An optional end time. The default is &#39;*&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s end time, or &#39;*&#39; if that is not set. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order. | 
 **filterExpression** | **string**| A string containing a filter expression. Expression variables are relative to the attribute. Use &#39;.&#39; to reference the containing attribute. The default is no filtering. | 
 **nameFilter** | **string**| The name query string used for filtering attributes. The default is no filter. | 
 **sampleInterval** | **string**| A time span specifies how often the filter expression is evaluated when computing the summary for an interval, if the sampleType is &#39;Interval&#39;. | 
 **sampleType** | **string**| A flag which specifies one or more summaries to compute for each interval over the time range. The default is &#39;ExpressionRecordedValues&#39;. | 
 **searchFullHierarchy** | **bool**| Specifies if the search should include attributes nested further than the immediate attributes of the searchRoot. The default is &#39;false&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **showExcluded** | **bool**| Specified if the search should include attributes with the Excluded property set. The default is &#39;false&#39;. | 
 **showHidden** | **bool**| Specified if the search should include attributes with the Hidden property set. The default is &#39;false&#39;. | 
 **startTime** | **string**| An optional start time. The default is &#39;*-1d&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s start time, or &#39;*-1d&#39; if that is not set. | 
 **summaryDuration** | **string**| The duration of each summary interval. | 
 **summaryType** | [**[]string**](string.md)| Specifies the kinds of summaries to produce over the range. The default is &#39;Total&#39;. Multiple summary types may be specified by using multiple instances of summaryType. | 
 **templateName** | **string**| Specify that included attributes must be members of this template. The default is no template filter. | 
 **timeType** | **string**| Specifies how to calculate the timestamp for each interval. The default is &#39;Auto&#39;. | 
 **timeZone** | **string**| The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsStreamSummaries**](Items[StreamSummaries].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamSetGetSummariesAdHoc**
> ItemsStreamSummaries StreamSetGetSummariesAdHoc(ctx, webId, optional)
Returns summary values of the specified streams.

Any time series value in the response that contains an 'Errors' property indicates PI Web API encountered a handled error during the transfer of the response stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | [**[]string**](string.md)| The ID of a stream. Multiple streams may be specified with multiple instances of the parameter. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | [**[]string**](string.md)| The ID of a stream. Multiple streams may be specified with multiple instances of the parameter. | 
 **calculationBasis** | **string**| Specifies the method of evaluating the data over the time range. The default is &#39;TimeWeighted&#39;. | 
 **endTime** | **string**| An optional end time. The default is &#39;*&#39;. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order. | 
 **filterExpression** | **string**| A string containing a filter expression. Expression variables are relative to the attribute. Use &#39;.&#39; to reference the containing attribute. The default is no filtering. | 
 **sampleInterval** | **string**| A time span specifies how often the filter expression is evaluated when computing the summary for an interval, if the sampleType is &#39;Interval&#39;. | 
 **sampleType** | **string**| A flag which specifies one or more summaries to compute for each interval over the time range. The default is &#39;ExpressionRecordedValues&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **startTime** | **string**| An optional start time. The default is &#39;*-1d&#39;. | 
 **summaryDuration** | **string**| The duration of each summary interval. | 
 **summaryType** | [**[]string**](string.md)| Specifies the kinds of summaries to produce over the range. The default is &#39;Total&#39;. Multiple summary types may be specified by using multiple instances of summaryType. | 
 **timeType** | **string**| Specifies how to calculate the timestamp for each interval. The default is &#39;Auto&#39;. | 
 **timeZone** | **string**| The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsStreamSummaries**](Items[StreamSummaries].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamSetGetValues**
> ItemsStreamValue StreamSetGetValues(ctx, webId, optional)
Returns values of the attributes for an Element, Event Frame or Attribute at the specified time.

Any time series value in the response that contains an 'Errors' property indicates PI Web API encountered a handled error during the transfer of the response stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of an Element, Event Frame or Attribute, which is the base element or parent of all the stream attributes. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of an Element, Event Frame or Attribute, which is the base element or parent of all the stream attributes. | 
 **categoryName** | **string**| Specify that included attributes must have this category. The default is no category filter. | 
 **nameFilter** | **string**| The name query string used for filtering attributes. The default is no filter. | 
 **searchFullHierarchy** | **bool**| Specifies if the search should include attributes nested further than the immediate attributes of the searchRoot. The default is &#39;false&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **showExcluded** | **bool**| Specified if the search should include attributes with the Excluded property set. The default is &#39;false&#39;. | 
 **showHidden** | **bool**| Specified if the search should include attributes with the Hidden property set. The default is &#39;false&#39;. | 
 **sortField** | **string**| The field or property of the object used to sort the returned collection. For better performance, by default no sorting is applied. &#39;Name&#39; is the only supported field by which to sort. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39; | 
 **templateName** | **string**| Specify that included attributes must be members of this template. The default is no template filter. | 
 **time** | **string**| An AF time string, which is used as the time context to get stream values if it is provided. By default, it is not specified, and the default time context of the AF object will be used. | 
 **timeZone** | **string**| The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsStreamValue**](Items[StreamValue].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamSetGetValuesAdHoc**
> ItemsStreamValue StreamSetGetValuesAdHoc(ctx, webId, optional)
Returns values of the specified streams.

Any time series value in the response that contains an 'Errors' property indicates PI Web API encountered a handled error during the transfer of the response stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | [**[]string**](string.md)| The ID of a stream. Multiple streams may be specified with multiple instances of the parameter. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | [**[]string**](string.md)| The ID of a stream. Multiple streams may be specified with multiple instances of the parameter. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **sortField** | **string**| The field or property of the object used to sort the returned collection. For better performance, by default no sorting is applied. &#39;Name&#39; is the only supported field by which to sort. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39; | 
 **time** | **string**| An AF time string, which is used as the time context to get stream values if it is provided. By default, it is not specified, and the default time context of the AF object will be used. | 
 **timeZone** | **string**| The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsStreamValue**](Items[StreamValue].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamSetUpdateValue**
> ItemsSubstatus StreamSetUpdateValue(ctx, webId, values, optional)
Updates a single value for the specified streams.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the parent element, event frame, or attribute. Attributes specified in the body must be descendants of the specified object. | 
  **values** | [**[]StreamValue**](StreamValue.md)| The values to add or update. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the parent element, event frame, or attribute. Attributes specified in the body must be descendants of the specified object. | 
 **values** | [**[]StreamValue**](StreamValue.md)| The values to add or update. | 
 **bufferOption** | **string**| The desired AFBufferOption. The default is &#39;BufferIfPossible&#39;. | 
 **updateOption** | **string**| The desired AFUpdateOption. The default is &#39;Replace&#39;. | 

### Return type

[**ItemsSubstatus**](Items[Substatus].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamSetUpdateValueAdHoc**
> ItemsSubstatus StreamSetUpdateValueAdHoc(ctx, values, optional)
Updates a single value for the specified streams.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **values** | [**[]StreamValue**](StreamValue.md)| The values to add or update. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **values** | [**[]StreamValue**](StreamValue.md)| The values to add or update. | 
 **bufferOption** | **string**| The desired AFBufferOption. The default is &#39;BufferIfPossible&#39;. | 
 **updateOption** | **string**| The desired AFUpdateOption. The default is &#39;Replace&#39;. | 

### Return type

[**ItemsSubstatus**](Items[Substatus].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamSetUpdateValues**
> ItemsItemsSubstatus StreamSetUpdateValues(ctx, webId, values, optional)
Updates multiple values for the specified streams.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the parent element, event frame, or attribute. Attributes specified in the body must be descendants of the specified object. | 
  **values** | [**[]StreamValues**](StreamValues.md)| The values to add or update. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the parent element, event frame, or attribute. Attributes specified in the body must be descendants of the specified object. | 
 **values** | [**[]StreamValues**](StreamValues.md)| The values to add or update. | 
 **bufferOption** | **string**| The desired AFBufferOption. The default is &#39;BufferIfPossible&#39;. | 
 **updateOption** | **string**| The desired AFUpdateOption. The default is &#39;Replace&#39;. | 

### Return type

[**ItemsItemsSubstatus**](Items[Items[Substatus]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamSetUpdateValuesAdHoc**
> ItemsItemsSubstatus StreamSetUpdateValuesAdHoc(ctx, values, optional)
Updates multiple values for the specified streams.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **values** | [**[]StreamValues**](StreamValues.md)| The values to add or update. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **values** | [**[]StreamValues**](StreamValues.md)| The values to add or update. | 
 **bufferOption** | **string**| The desired AFBufferOption. The default is &#39;BufferIfPossible&#39;. | 
 **updateOption** | **string**| The desired AFUpdateOption. The default is &#39;Replace&#39;. | 

### Return type

[**ItemsItemsSubstatus**](Items[Items[Substatus]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

