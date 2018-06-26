# \StreamApi

All URIs are relative to *https://localhost/piwebapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StreamGetChannel**](StreamApi.md#StreamGetChannel) | **Get** /streams/{webId}/channel | Opens a channel that will send messages about any value changes for the specified stream.
[**StreamGetEnd**](StreamApi.md#StreamGetEnd) | **Get** /streams/{webId}/end | Returns the end-of-stream value of the stream.
[**StreamGetInterpolated**](StreamApi.md#StreamGetInterpolated) | **Get** /streams/{webId}/interpolated | Retrieves interpolated values over the specified time range at the specified sampling interval.
[**StreamGetInterpolatedAtTimes**](StreamApi.md#StreamGetInterpolatedAtTimes) | **Get** /streams/{webId}/interpolatedattimes | Retrieves interpolated values over the specified time range at the specified sampling interval.
[**StreamGetPlot**](StreamApi.md#StreamGetPlot) | **Get** /streams/{webId}/plot | Retrieves values over the specified time range suitable for plotting over the number of intervals (typically represents pixels).
[**StreamGetRecorded**](StreamApi.md#StreamGetRecorded) | **Get** /streams/{webId}/recorded | Returns a list of compressed values for the requested time range from the source provider.
[**StreamGetRecordedAtTime**](StreamApi.md#StreamGetRecordedAtTime) | **Get** /streams/{webId}/recordedattime | Returns a single recorded value based on the passed time and retrieval mode from the stream.
[**StreamGetRecordedAtTimes**](StreamApi.md#StreamGetRecordedAtTimes) | **Get** /streams/{webId}/recordedattimes | Retrieves recorded values at the specified times.
[**StreamGetSummary**](StreamApi.md#StreamGetSummary) | **Get** /streams/{webId}/summary | Returns a summary over the specified time range for the stream.
[**StreamGetValue**](StreamApi.md#StreamGetValue) | **Get** /streams/{webId}/value | Returns the value of the stream at the specified time. By default, this is usually the current value.
[**StreamUpdateValue**](StreamApi.md#StreamUpdateValue) | **Post** /streams/{webId}/value | Updates a value for the specified stream.
[**StreamUpdateValues**](StreamApi.md#StreamUpdateValues) | **Post** /streams/{webId}/recorded | Updates multiple values for the specified stream.


# **StreamGetChannel**
> StreamGetChannel(ctx, webId, optional)
Opens a channel that will send messages about any value changes for the specified stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the stream. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the stream. | 
 **heartbeatRate** | **int32**| HeartbeatRate is an integer multiple of the Polling Interval. It specifies the rate at which a client will receive an empty message if there are no data updates. It can be used to check that the connection is still alive. Zero/negative values correspond to no heartbeat. By default, no empty messages will be sent to the user. | 
 **includeInitialValues** | **bool**| Specified if the channel should send a message with the current value of the stream after the connection is opened. The default is &#39;false&#39;. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamGetEnd**
> TimedValue StreamGetEnd(ctx, webId, optional)
Returns the end-of-stream value of the stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the stream. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the stream. | 
 **desiredUnits** | **string**| The name or abbreviation of the desired units of measure for the returned value, as found in the UOM database associated with the attribute. If not specified for an attribute, the attribute&#39;s default unit of measure is used. If the underlying stream is a point, this value may not be specified, as points are not associated with a unit of measure. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 

### Return type

[**TimedValue**](TimedValue.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamGetInterpolated**
> TimedValues StreamGetInterpolated(ctx, webId, optional)
Retrieves interpolated values over the specified time range at the specified sampling interval.

Any time series value in the response that contains an 'Errors' property indicates PI Web API encountered a handled error during the transfer of the response stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the stream. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the stream. | 
 **desiredUnits** | **string**| The name or abbreviation of the desired units of measure for the returned value, as found in the UOM database associated with the attribute. If not specified for an attribute, the attribute&#39;s default unit of measure is used. If the underlying stream is a point, this value may not be specified, as points are not associated with a unit of measure. | 
 **endTime** | **string**| An optional end time. The default is &#39;*&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s end time, or &#39;*&#39; if that is not set. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order. | 
 **filterExpression** | **string**| An optional string containing a filter expression. Expression variables are relative to the data point. Use &#39;.&#39; to reference the containing attribute. If the attribute does not support filtering, the filter will be ignored. The default is no filtering. | 
 **includeFilteredValues** | **bool**| Specify &#39;true&#39; to indicate that values which fail the filter criteria are present in the returned data at the times where they occurred with a value set to a &#39;Filtered&#39; enumeration value with bad status. Repeated consecutive failures are omitted. | 
 **interval** | **string**| The sampling interval, in AFTimeSpan format. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **startTime** | **string**| An optional start time. The default is &#39;*-1d&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s start time, or &#39;*-1d&#39; if that is not set. | 
 **syncTime** | **string**| An optional start time anchor, in AFTime format. When specified, interpolated data retrieval will use the sync time as the origin for calculating the interval times. | 
 **syncTimeBoundaryType** | **string**| An optional string specifying the boundary type to use when applying a syncTime. The allowed values are &#39;Inside&#39; and &#39;Outside&#39;. The default is &#39;Inside&#39;. | 
 **timeZone** | **string**| The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used. | 

### Return type

[**TimedValues**](TimedValues.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamGetInterpolatedAtTimes**
> TimedValues StreamGetInterpolatedAtTimes(ctx, webId, optional)
Retrieves interpolated values over the specified time range at the specified sampling interval.

Any time series value in the response that contains an 'Errors' property indicates PI Web API encountered a handled error during the transfer of the response stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the stream. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the stream. | 
 **desiredUnits** | **string**| The name or abbreviation of the desired units of measure for the returned value, as found in the UOM database associated with the attribute. If not specified for an attribute, the attribute&#39;s default unit of measure is used. If the underlying stream is a point, this value may not be specified, as points are not associated with a unit of measure. | 
 **filterExpression** | **string**| An optional string containing a filter expression. Expression variables are relative to the data point. Use &#39;.&#39; to reference the containing attribute. If the attribute does not support filtering, the filter will be ignored. The default is no filtering. | 
 **includeFilteredValues** | **bool**| Specify &#39;true&#39; to indicate that values which fail the filter criteria are present in the returned data at the times where they occurred with a value set to a &#39;Filtered&#39; enumeration value with bad status. Repeated consecutive failures are omitted. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39;. | 
 **time** | [**[]string**](string.md)| The timestamp at which to retrieve an interpolated value. Multiple timestamps may be specified with multiple instances of the parameter. | 
 **timeZone** | **string**| The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used. | 

### Return type

[**TimedValues**](TimedValues.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamGetPlot**
> TimedValues StreamGetPlot(ctx, webId, optional)
Retrieves values over the specified time range suitable for plotting over the number of intervals (typically represents pixels).

For each interval, the data available is examined and significant values are returned. Each interval can produce up to 5 values if they are unique, the first value in the interval, the last value, the highest value, the lowest value and at most one exceptional point (bad status or digital state). Any time series value in the response that contains an 'Errors' property indicates PI Web API encountered a handled error during the transfer of the response stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the stream. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the stream. | 
 **desiredUnits** | **string**| The name or abbreviation of the desired units of measure for the returned value, as found in the UOM database associated with the attribute. If not specified for an attribute, the attribute&#39;s default unit of measure is used. If the underlying stream is a point, this value may not be specified, as points are not associated with a unit of measure. | 
 **endTime** | **string**| An optional end time. The default is &#39;*&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s end time, or &#39;*&#39; if that is not set. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order. | 
 **intervals** | **int32**| The number of intervals to plot over. Typically, this would be the number of horizontal pixels in the trend. The default is &#39;24&#39;. For each interval, the data available is examined and significant values are returned. Each interval can produce up to 5 values if they are unique, the first value in the interval, the last value, the highest value, the lowest value and at most one exceptional point (bad status or digital state). | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **startTime** | **string**| An optional start time. The default is &#39;*-1d&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s start time, or &#39;*-1d&#39; if that is not set. | 
 **timeZone** | **string**| The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used. | 

### Return type

[**TimedValues**](TimedValues.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamGetRecorded**
> ExtendedTimedValues StreamGetRecorded(ctx, webId, optional)
Returns a list of compressed values for the requested time range from the source provider.

Returned times are affected by the specified boundary type. If no values are found for the time range and conditions specified then the HTTP response will be success, with a body containing an empty array of Items. When specifying true for the includeFilteredValues parameter, consecutive filtered events are not returned. The first value that would be filtered out is returned with its time and the enumeration value \"Filtered\". The next value in the collection will be the next compressed value in the specified direction that passes the filter criteria - if any. When both boundaryType and a filterExpression are specified, the events returned for the boundary condition specified are passed through the filter. If the includeFilteredValues parameter is true, the boundary values will be reported at the proper timestamps with the enumeration value \"Filtered\" when the filter conditions are not met at the boundary time. If the includeFilteredValues parameter is false for this case, no event is returned for the boundary time. Any time series value in the response that contains an 'Errors' property indicates PI Web API encountered a handled error during the transfer of the response stream.   If only recorded values with annotations are desired, the filterExpression parameter should include the filter IsSet('.', \"a\").

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the stream. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the stream. | 
 **associations** | **string**| Associated values to return in the response, separated by semicolons (;). This call supports Annotations to return events with annotation values. If this parameter is not specified, annotation values are not returned. | 
 **boundaryType** | **string**| An optional value that determines how the times and values of the returned end points are determined. The default is &#39;Inside&#39;. | 
 **desiredUnits** | **string**| The name or abbreviation of the desired units of measure for the returned value, as found in the UOM database associated with the attribute. If not specified for an attribute, the attribute&#39;s default unit of measure is used. If the underlying stream is a point, this value may not be specified, as points are not associated with a unit of measure. | 
 **endTime** | **string**| An optional end time. The default is &#39;*&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s end time, or &#39;*&#39; if that is not set. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order. | 
 **filterExpression** | **string**| An optional string containing a filter expression. Expression variables are relative to the data point. Use &#39;.&#39; to reference the containing attribute. The default is no filtering. | 
 **includeFilteredValues** | **bool**| Specify &#39;true&#39; to indicate that values which fail the filter criteria are present in the returned data at the times where they occurred with a value set to a &#39;Filtered&#39; enumeration value with bad status. Repeated consecutive failures are omitted. | 
 **maxCount** | **int32**| The maximum number of values to be returned. The default is 1000. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **startTime** | **string**| An optional start time. The default is &#39;*-1d&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s start time, or &#39;*-1d&#39; if that is not set. | 
 **timeZone** | **string**| The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used. | 

### Return type

[**ExtendedTimedValues**](ExtendedTimedValues.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamGetRecordedAtTime**
> ExtendedTimedValue StreamGetRecordedAtTime(ctx, webId, time, optional)
Returns a single recorded value based on the passed time and retrieval mode from the stream.

If only recorded values with annotations are desired, the filterExpression parameter should include the filter IsSet('.', \"a\").

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the stream. | 
  **time** | **string**| The timestamp at which the value is desired. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the stream. | 
 **time** | **string**| The timestamp at which the value is desired. | 
 **associations** | **string**| Associated values to return in the response, separated by semicolons (;). This call supports Annotations to return events with annotation values. If this parameter is not specified, annotation values are not returned. | 
 **desiredUnits** | **string**| The name or abbreviation of the desired units of measure for the returned value, as found in the UOM database associated with the attribute. If not specified for an attribute, the attribute&#39;s default unit of measure is used. If the underlying stream is a point, this value may not be specified, as points are not associated with a unit of measure. | 
 **retrievalMode** | **string**| An optional value that determines the value to return when a value doesn&#39;t exist at the exact time specified. The default is &#39;Auto&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **timeZone** | **string**| The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used. | 

### Return type

[**ExtendedTimedValue**](ExtendedTimedValue.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamGetRecordedAtTimes**
> ExtendedTimedValues StreamGetRecordedAtTimes(ctx, webId, optional)
Retrieves recorded values at the specified times.

Any time series value in the response that contains an 'Errors' property indicates PI Web API encountered a handled error during the transfer of the response stream.   If only recorded values with annotations are desired, the filterExpression parameter should include the filter IsSet('.', \"a\").

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the stream. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the stream. | 
 **associations** | **string**| Associated values to return in the response, separated by semicolons (;). This call supports Annotations to return events with annotation values. If this parameter is not specified, annotation values are not returned. | 
 **desiredUnits** | **string**| The name or abbreviation of the desired units of measure for the returned value, as found in the UOM database associated with the attribute. If not specified for an attribute, the attribute&#39;s default unit of measure is used. If the underlying stream is a point, this value may not be specified, as points are not associated with a unit of measure. | 
 **retrievalMode** | **string**| An optional value that determines the value to return when a value doesn&#39;t exist at the exact time specified. The default is &#39;Auto&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39;. | 
 **time** | [**[]string**](string.md)| The timestamp at which to retrieve a recorded value. Multiple timestamps may be specified with multiple instances of the parameter. | 
 **timeZone** | **string**| The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used. | 

### Return type

[**ExtendedTimedValues**](ExtendedTimedValues.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamGetSummary**
> ItemsSummaryValue StreamGetSummary(ctx, webId, optional)
Returns a summary over the specified time range for the stream.

Count is the only summary type supported on non-numeric streams. Requesting a summary for any other type will generate an error. Time-weighted totals are computed by integrating the rate tag values over the requested time range. If some of the data are bad in the time range, the calculated total is divided by the fraction of the time period for which there are good values. This approach is equivalent to assuming that during the period of bad data, the tag takes on the average values for the entire calculation time range. The PercentGood summary may be used to determine if the calculation results are suitable for the application's purposes. For time-weighted totals, if the time unit rate of the stream cannot be determined, then the value will be totaled assuming a unit of \"per day\" and no unit of measure will be assigned to the value. If the measured time component of the tag is not based on a day, the user of the data must convert the totalized value to the correct units.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the stream. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the stream. | 
 **calculationBasis** | **string**| Specifies the method of evaluating the data over the time range. The default is &#39;TimeWeighted&#39;. | 
 **endTime** | **string**| An optional end time. The default is &#39;*&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s end time, or &#39;*&#39; if that is not set. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order. | 
 **filterExpression** | **string**| A string containing a filter expression. Expression variables are relative to the attribute. Use &#39;.&#39; to reference the containing attribute. | 
 **sampleInterval** | **string**| When the sampleType is Interval, sampleInterval specifies how often the filter expression is evaluated when computing the summary for an interval. | 
 **sampleType** | **string**| Defines the evaluation of an expression over a time range. The default is &#39;ExpressionRecordedValues&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **startTime** | **string**| An optional start time. The default is &#39;*-1d&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s start time, or &#39;*-1d&#39; if that is not set. | 
 **summaryDuration** | **string**| The duration of each summary interval. If specified in hours, minutes, seconds, or milliseconds, the summary durations will be evenly spaced UTC time intervals. Longer interval types are interpreted using wall clock rules and are time zone dependent. | 
 **summaryType** | [**[]string**](string.md)| Specifies the kinds of summaries to produce over the range. The default is &#39;Total&#39;. Multiple summary types may be specified by using multiple instances of summaryType. | 
 **timeType** | **string**| Specifies how to calculate the timestamp for each interval. The default is &#39;Auto&#39;. | 
 **timeZone** | **string**| The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used. | 

### Return type

[**ItemsSummaryValue**](Items[SummaryValue].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamGetValue**
> TimedValue StreamGetValue(ctx, webId, optional)
Returns the value of the stream at the specified time. By default, this is usually the current value.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the stream. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the stream. | 
 **desiredUnits** | **string**| The name or abbreviation of the desired units of measure for the returned value, as found in the UOM database associated with the attribute. If not specified for an attribute, the attribute&#39;s default unit of measure is used. If the underlying stream is a point, this value may not be specified, as points are not associated with a unit of measure. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **time** | **string**| An optional time. The default time context is determined from the owning object - for example, the time range of the event frame or transfer which holds this attribute. Otherwise, the implementation of the Data Reference determines the meaning of no context. For Points or simply configured PI Point Data References, this means the snapshot value of the PI Point on the Data Server. | 
 **timeZone** | **string**| The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used. | 

### Return type

[**TimedValue**](TimedValue.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamUpdateValue**
> StreamUpdateValue(ctx, webId, value, optional)
Updates a value for the specified stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the stream. | 
  **value** | [**TimedValue**](TimedValue.md)| The value to add or update. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the stream. | 
 **value** | [**TimedValue**](TimedValue.md)| The value to add or update. | 
 **bufferOption** | **string**| The desired AFBufferOption. The default is &#39;BufferIfPossible&#39;. | 
 **updateOption** | **string**| The desired AFUpdateOption. The default is &#39;Replace&#39;. This parameter is ignored if the attribute is a configuration item. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StreamUpdateValues**
> ItemsSubstatus StreamUpdateValues(ctx, webId, values, optional)
Updates multiple values for the specified stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the stream. | 
  **values** | [**[]TimedValue**](TimedValue.md)| The values to add or update. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the stream. | 
 **values** | [**[]TimedValue**](TimedValue.md)| The values to add or update. | 
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

