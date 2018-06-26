# \CalculationApi

All URIs are relative to *https://localhost/piwebapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CalculationGetAtIntervals**](CalculationApi.md#CalculationGetAtIntervals) | **Get** /calculation/intervals | Returns results of evaluating the expression over the time range from the start time to the end time at a defined interval.
[**CalculationGetAtRecorded**](CalculationApi.md#CalculationGetAtRecorded) | **Get** /calculation/recorded | Returns the result of evaluating the expression at each point in time over the time range from the start time to the end time where a recorded value exists for a member of the expression.
[**CalculationGetAtTimes**](CalculationApi.md#CalculationGetAtTimes) | **Get** /calculation/times | Returns the result of evaluating the expression at the specified timestamps.
[**CalculationGetSummary**](CalculationApi.md#CalculationGetSummary) | **Get** /calculation/summary | Returns the result of evaluating the expression over the time range from the start time to the end time. The time range is first divided into a number of summary intervals. Then the calculation is performed for the specified summaries over each interval.


# **CalculationGetAtIntervals**
> TimedValues CalculationGetAtIntervals(ctx, optional)
Returns results of evaluating the expression over the time range from the start time to the end time at a defined interval.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endTime** | **string**| An optional end time. The default is &#39;*&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s end time, or &#39;*&#39; if that is not set. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order. | 
 **expression** | **string**| A string containing the expression to be evaluated. The syntax for the expression generally follows the Performance Equation syntax as described in the PI Server documentation, with the exception that expressions which target AF objects use attribute names in place of tag names in the equation. | 
 **sampleInterval** | **string**| A time span specifies how often the filter expression is evaluated when computing the summary for an interval. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **startTime** | **string**| An optional start time. The default is &#39;*-1d&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s start time, or &#39;*-1d&#39; if that is not set. | 
 **webId** | **string**| The ID of the target object of the expression. A target object can be a Data Server, a database, an element, an event frame or an attribute. References to attributes or points are based on the target. If this parameter is not provided, the target object is set to null. | 

### Return type

[**TimedValues**](TimedValues.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CalculationGetAtRecorded**
> TimedValues CalculationGetAtRecorded(ctx, optional)
Returns the result of evaluating the expression at each point in time over the time range from the start time to the end time where a recorded value exists for a member of the expression.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endTime** | **string**| An optional end time. The default is &#39;*&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s end time, or &#39;*&#39; if that is not set. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order. | 
 **expression** | **string**| A string containing the expression to be evaluated. The syntax for the expression generally follows the Performance Equation syntax as described in the PI Server documentation, with the exception that expressions which target AF objects use attribute names in place of tag names in the equation. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **startTime** | **string**| An optional start time. The default is &#39;*-1d&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s start time, or &#39;*-1d&#39; if that is not set. | 
 **webId** | **string**| The ID of the target object of the expression. A target object can be a Data Server, a database, an element, an event frame or an attribute. References to attributes or points are based on the target. If this parameter is not provided, the target object is set to null. | 

### Return type

[**TimedValues**](TimedValues.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CalculationGetAtTimes**
> TimedValues CalculationGetAtTimes(ctx, optional)
Returns the result of evaluating the expression at the specified timestamps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expression** | **string**| A string containing the expression to be evaluated. The syntax for the expression generally follows the Performance Equation syntax as described in the PI Server documentation, with the exception that expressions which target AF objects use attribute names in place of tag names in the equation. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39;. | 
 **time** | [**[]string**](string.md)| A list of timestamps at which to calculate the expression. | 
 **webId** | **string**| The ID of the target object of the expression. A target object can be a Data Server, a database, an element, an event frame or an attribute. References to attributes or points are based on the target. If this parameter is not provided, the target object is set to null. | 

### Return type

[**TimedValues**](TimedValues.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CalculationGetSummary**
> ItemsSummaryValue CalculationGetSummary(ctx, optional)
Returns the result of evaluating the expression over the time range from the start time to the end time. The time range is first divided into a number of summary intervals. Then the calculation is performed for the specified summaries over each interval.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **calculationBasis** | **string**| Specifies the method of evaluating the data over the time range. The default is &#39;TimeWeighted&#39;. | 
 **endTime** | **string**| An optional end time. The default is &#39;*&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s end time, or &#39;*&#39; if that is not set. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order. | 
 **expression** | **string**| A string containing the expression to be evaluated. The syntax for the expression generally follows the Performance Equation syntax as described in the PI Server documentation, with the exception that expressions which target AF objects use attribute names in place of tag names in the equation. | 
 **sampleInterval** | **string**| A time span specifies how often the filter expression is evaluated when computing the summary for an interval, if the sampleType is &#39;Interval&#39;. | 
 **sampleType** | **string**| A flag which specifies one or more summaries to compute for each interval over the time range. The default is &#39;ExpressionRecordedValues&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **startTime** | **string**| An optional start time. The default is &#39;*-1d&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s start time, or &#39;*-1d&#39; if that is not set. | 
 **summaryDuration** | **string**| The duration of each summary interval. | 
 **summaryType** | [**[]string**](string.md)| Specifies the kinds of summaries to produce over the range. The default is &#39;Total&#39;. Multiple summary types may be specified by using multiple instances of summaryType. | 
 **timeType** | **string**| Specifies how to calculate the timestamp for each interval. The default is &#39;Auto&#39;. | 
 **webId** | **string**| The ID of the target object of the expression. A target object can be a Data Server, a database, an element, an event frame or an attribute. References to attributes or points are based on the target. If this parameter is not provided, the target object is set to null. | 

### Return type

[**ItemsSummaryValue**](Items[SummaryValue].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

