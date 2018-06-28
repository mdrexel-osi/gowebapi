// ************************************************************************
// * Copyright 2018 OSIsoft, LLC
// * Licensed under the Apache License, Version 2.0 (the "License");
// * you may not use this file except in compliance with the License.
// * You may obtain a copy of the License at
// * 
// *   <http://www.apache.org/licenses/LICENSE-2.0>
// * 
// * Unless required by applicable law or agreed to in writing, software
// * distributed under the License is distributed on an "AS IS" BASIS,
// * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// * See the License for the specific language governing permissions and
// * limitations under the License.
// ************************************************************************

package gowebapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/context"
)

// Linger please
var (
	_ context.Context
)

type StreamSetApiService service

/* StreamSetApiService Opens a channel that will send messages about any value changes for the attributes of an Element, Event Frame, or Attribute.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of an Element, Event Frame or Attribute, which is the base element or parent of all the stream attributes.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "categoryName" (string) Specify that included attributes must have this category. The default is no category filter.
    @param "heartbeatRate" (int32) Specifies the maximum number of consecutive empty messages that can be elapsed with no new data updates from the PI System, after which the client receives an empty payload. It helps to check if the connection is still alive. Zero/negative values correspond to no heartbeat, and the default value is no heartbeat.
    @param "includeInitialValues" (bool) Specified if the channel should send a message with the current values of all the streams after the connection is opened. The default is &#39;false&#39;.
    @param "nameFilter" (string) The name query string used for filtering attributes. The default is no filter.
    @param "searchFullHierarchy" (bool) Specifies if the search should include attributes nested further than the immediate attributes of the searchRoot. The default is &#39;false&#39;.
    @param "showExcluded" (bool) Specified if the search should include attributes with the Excluded property set. The default is &#39;false&#39;.
    @param "showHidden" (bool) Specified if the search should include attributes with the Hidden property set. The default is &#39;false&#39;.
    @param "templateName" (string) Specify that included attributes must be members of this template. The default is no template filter.
    @param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.
@return */
func (a *StreamSetApiService) StreamSetGetChannel(ctx context.Context, webId string, localVarOptionals map[string]interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streamsets/{webId}/channel"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["categoryName"], "string", "categoryName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["heartbeatRate"], "int32", "heartbeatRate"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["includeInitialValues"], "bool", "includeInitialValues"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["nameFilter"], "string", "nameFilter"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["searchFullHierarchy"], "bool", "searchFullHierarchy"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["showExcluded"], "bool", "showExcluded"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["showHidden"], "bool", "showHidden"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["templateName"], "string", "templateName"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["categoryName"].(string); localVarOk {
		localVarQueryParams.Add("categoryName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["heartbeatRate"].(int32); localVarOk {
		localVarQueryParams.Add("heartbeatRate", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["includeInitialValues"].(bool); localVarOk {
		localVarQueryParams.Add("includeInitialValues", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["nameFilter"].(string); localVarOk {
		localVarQueryParams.Add("nameFilter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["searchFullHierarchy"].(bool); localVarOk {
		localVarQueryParams.Add("searchFullHierarchy", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["showExcluded"].(bool); localVarOk {
		localVarQueryParams.Add("showExcluded", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["showHidden"].(bool); localVarOk {
		localVarQueryParams.Add("showHidden", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["templateName"].(string); localVarOk {
		localVarQueryParams.Add("templateName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["webIdType"].(string); localVarOk {
		localVarQueryParams.Add("webIdType", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/json",
		"text/html",
		"application/x-ms-application",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* StreamSetApiService Opens a channel that will send messages about any value changes for the specified streams.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of a stream. Multiple streams may be specified with multiple instances of the parameter.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "heartbeatRate" (int32) Specifies the maximum number of consecutive empty messages that can be elapsed with no new data updates from the PI System, after which the client receives an empty payload. It helps to check if the connection is still alive. Zero/negative values correspond to no heartbeat, and the default value is no heartbeat.
    @param "includeInitialValues" (bool) Specified if the channel should send a message with the current values of all the streams after the connection is opened. The default is &#39;false&#39;.
    @param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.
@return */
func (a *StreamSetApiService) StreamSetGetChannelAdHoc(ctx context.Context, webId []string, localVarOptionals map[string]interface{}) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streamsets/channel"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["heartbeatRate"], "int32", "heartbeatRate"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["includeInitialValues"], "bool", "includeInitialValues"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return nil, err
	}

	localVarQueryParams.Add("webId", parameterToString(webId, "multi"))
	if localVarTempParam, localVarOk := localVarOptionals["heartbeatRate"].(int32); localVarOk {
		localVarQueryParams.Add("heartbeatRate", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["includeInitialValues"].(bool); localVarOk {
		localVarQueryParams.Add("includeInitialValues", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["webIdType"].(string); localVarOk {
		localVarQueryParams.Add("webIdType", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/json",
		"text/html",
		"application/x-ms-application",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* StreamSetApiService Returns End of stream values of the attributes for an Element, Event Frame or Attribute
Any time series value in the response that contains an &#39;Errors&#39; property indicates PI Web API encountered a handled error during the transfer of the response stream.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of an Element, Event Frame or Attribute, which is the base element or parent of all the stream attributes.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "categoryName" (string) Specify that included attributes must have this category. The default is no category filter.
    @param "nameFilter" (string) The name query string used for filtering attributes. The default is no filter.
    @param "searchFullHierarchy" (bool) Specifies if the search should include attributes nested further than the immediate attributes of the searchRoot. The default is &#39;false&#39;.
    @param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
    @param "showExcluded" (bool) Specified if the search should include attributes with the Excluded property set. The default is &#39;false&#39;.
    @param "showHidden" (bool) Specified if the search should include attributes with the Hidden property set. The default is &#39;false&#39;.
    @param "sortField" (string) The field or property of the object used to sort the returned collection. For better performance, by default no sorting is applied. &#39;Name&#39; is the only supported field by which to sort.
    @param "sortOrder" (string) The order that the returned collection is sorted. The default is &#39;Ascending&#39;
    @param "templateName" (string) Specify that included attributes must be members of this template. The default is no template filter.
    @param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.
@return ItemsStreamValue*/
func (a *StreamSetApiService) StreamSetGetEnd(ctx context.Context, webId string, localVarOptionals map[string]interface{}) (ItemsStreamValue, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsStreamValue
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streamsets/{webId}/end"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["categoryName"], "string", "categoryName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["nameFilter"], "string", "nameFilter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["searchFullHierarchy"], "bool", "searchFullHierarchy"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["showExcluded"], "bool", "showExcluded"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["showHidden"], "bool", "showHidden"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortField"], "string", "sortField"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortOrder"], "string", "sortOrder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["templateName"], "string", "templateName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["categoryName"].(string); localVarOk {
		localVarQueryParams.Add("categoryName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["nameFilter"].(string); localVarOk {
		localVarQueryParams.Add("nameFilter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["searchFullHierarchy"].(bool); localVarOk {
		localVarQueryParams.Add("searchFullHierarchy", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["showExcluded"].(bool); localVarOk {
		localVarQueryParams.Add("showExcluded", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["showHidden"].(bool); localVarOk {
		localVarQueryParams.Add("showHidden", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortField"].(string); localVarOk {
		localVarQueryParams.Add("sortField", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortOrder"].(string); localVarOk {
		localVarQueryParams.Add("sortOrder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["templateName"].(string); localVarOk {
		localVarQueryParams.Add("templateName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["webIdType"].(string); localVarOk {
		localVarQueryParams.Add("webIdType", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/json",
		"text/html",
		"application/x-ms-application",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* StreamSetApiService Returns End Of Stream values for attributes of the specified streams
Any time series value in the response that contains an &#39;Errors&#39; property indicates PI Web API encountered a handled error during the transfer of the response stream.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of a stream. Multiple streams may be specified with multiple instances of the parameter.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
    @param "sortField" (string) The field or property of the object used to sort the returned collection. For better performance, by default no sorting is applied. &#39;Name&#39; is the only supported field by which to sort.
    @param "sortOrder" (string) The order that the returned collection is sorted. The default is &#39;Ascending&#39;
    @param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.
@return ItemsStreamValues*/
func (a *StreamSetApiService) StreamSetGetEndAdHoc(ctx context.Context, webId []string, localVarOptionals map[string]interface{}) (ItemsStreamValues, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsStreamValues
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streamsets/end"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortField"], "string", "sortField"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortOrder"], "string", "sortOrder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("webId", parameterToString(webId, "multi"))
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortField"].(string); localVarOk {
		localVarQueryParams.Add("sortField", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortOrder"].(string); localVarOk {
		localVarQueryParams.Add("sortOrder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["webIdType"].(string); localVarOk {
		localVarQueryParams.Add("webIdType", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/json",
		"text/html",
		"application/x-ms-application",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* StreamSetApiService Returns interpolated values of attributes for an element, event frame or attribute over the specified time range at the specified sampling interval.
Any time series value in the response that contains an &#39;Errors&#39; property indicates PI Web API encountered a handled error during the transfer of the response stream.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of an element, event frame or attribute, which is the base element or parent of all the stream attributes.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "categoryName" (string) Specify that included attributes must have this category. The default is no category filter.
    @param "endTime" (string) An optional end time. The default is &#39;*&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s end time, or &#39;*&#39; if that is not set. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order.
    @param "filterExpression" (string) An optional string containing a filter expression. Expression variables are relative to the data point. Use &#39;.&#39; to reference the containing attribute. If the attribute does not support filtering, the filter will be ignored. The default is no filtering.
    @param "includeFilteredValues" (bool) Specify &#39;true&#39; to indicate that values which fail the filter criteria are present in the returned data at the times where they occurred with a value set to a &#39;Filtered&#39; enumeration value with bad status. Repeated consecutive failures are omitted.
    @param "interval" (string) The sampling interval, in AFTimeSpan format.
    @param "nameFilter" (string) The name query string used for filtering attributes. The default is no filter.
    @param "searchFullHierarchy" (bool) Specifies if the search should include attributes nested further than the immediate attributes of the searchRoot. The default is &#39;false&#39;.
    @param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
    @param "showExcluded" (bool) Specified if the search should include attributes with the Excluded property set. The default is &#39;false&#39;.
    @param "showHidden" (bool) Specified if the search should include attributes with the Hidden property set. The default is &#39;false&#39;.
    @param "sortField" (string) The field or property of the object used to sort the returned collection. For better performance, by default no sorting is applied. &#39;Name&#39; is the only supported field by which to sort.
    @param "sortOrder" (string) The order that the returned collection is sorted. The default is &#39;Ascending&#39;.
    @param "startTime" (string) An optional start time. The default is &#39;*-1d&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s start time, or &#39;*-1d&#39; if that is not set.
    @param "syncTime" (string) An optional start time anchor, in AFTime format. When specified, interpolated data retrieval will use the sync time as the origin for calculating the interval times.
    @param "syncTimeBoundaryType" (string) An optional string specifying the boundary type to use when applying a syncTime. The allowed values are &#39;Inside&#39; and &#39;Outside&#39;. The default is &#39;Inside&#39;.
    @param "templateName" (string) Specify that included attributes must be members of this template. The default is no template filter.
    @param "timeZone" (string) The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used.
    @param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.
@return ItemsStreamValues*/
func (a *StreamSetApiService) StreamSetGetInterpolated(ctx context.Context, webId string, localVarOptionals map[string]interface{}) (ItemsStreamValues, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsStreamValues
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streamsets/{webId}/interpolated"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["categoryName"], "string", "categoryName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["endTime"], "string", "endTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["filterExpression"], "string", "filterExpression"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["includeFilteredValues"], "bool", "includeFilteredValues"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["interval"], "string", "interval"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["nameFilter"], "string", "nameFilter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["searchFullHierarchy"], "bool", "searchFullHierarchy"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["showExcluded"], "bool", "showExcluded"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["showHidden"], "bool", "showHidden"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortField"], "string", "sortField"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortOrder"], "string", "sortOrder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startTime"], "string", "startTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["syncTime"], "string", "syncTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["syncTimeBoundaryType"], "string", "syncTimeBoundaryType"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["templateName"], "string", "templateName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["categoryName"].(string); localVarOk {
		localVarQueryParams.Add("categoryName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["endTime"].(string); localVarOk {
		localVarQueryParams.Add("endTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filterExpression"].(string); localVarOk {
		localVarQueryParams.Add("filterExpression", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["includeFilteredValues"].(bool); localVarOk {
		localVarQueryParams.Add("includeFilteredValues", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["interval"].(string); localVarOk {
		localVarQueryParams.Add("interval", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["nameFilter"].(string); localVarOk {
		localVarQueryParams.Add("nameFilter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["searchFullHierarchy"].(bool); localVarOk {
		localVarQueryParams.Add("searchFullHierarchy", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["showExcluded"].(bool); localVarOk {
		localVarQueryParams.Add("showExcluded", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["showHidden"].(bool); localVarOk {
		localVarQueryParams.Add("showHidden", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortField"].(string); localVarOk {
		localVarQueryParams.Add("sortField", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortOrder"].(string); localVarOk {
		localVarQueryParams.Add("sortOrder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startTime"].(string); localVarOk {
		localVarQueryParams.Add("startTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["syncTime"].(string); localVarOk {
		localVarQueryParams.Add("syncTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["syncTimeBoundaryType"].(string); localVarOk {
		localVarQueryParams.Add("syncTimeBoundaryType", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["templateName"].(string); localVarOk {
		localVarQueryParams.Add("templateName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeZone"].(string); localVarOk {
		localVarQueryParams.Add("timeZone", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["webIdType"].(string); localVarOk {
		localVarQueryParams.Add("webIdType", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/json",
		"text/html",
		"application/x-ms-application",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* StreamSetApiService Returns interpolated values of the specified streams over the specified time range at the specified sampling interval.
Any time series value in the response that contains an &#39;Errors&#39; property indicates PI Web API encountered a handled error during the transfer of the response stream.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of a stream. Multiple streams may be specified with multiple instances of the parameter.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "endTime" (string) An optional end time. The default is &#39;*&#39;. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order.
    @param "filterExpression" (string) An optional string containing a filter expression. Expression variables are relative to the data point. Use &#39;.&#39; to reference the containing attribute. If the attribute does not support filtering, the filter will be ignored. The default is no filtering.
    @param "includeFilteredValues" (bool) Specify &#39;true&#39; to indicate that values which fail the filter criteria are present in the returned data at the times where they occurred with a value set to a &#39;Filtered&#39; enumeration value with bad status. Repeated consecutive failures are omitted.
    @param "interval" (string) The sampling interval, in AFTimeSpan format.
    @param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
    @param "sortField" (string) The field or property of the object used to sort the returned collection. For better performance, by default no sorting is applied. &#39;Name&#39; is the only supported field by which to sort.
    @param "sortOrder" (string) The order that the returned collection is sorted. The default is &#39;Ascending&#39;
    @param "startTime" (string) An optional start time. The default is &#39;*-1d&#39;.
    @param "syncTime" (string) An optional start time anchor, in AFTime format. When specified, interpolated data retrieval will use the sync time as the origin for calculating the interval times.
    @param "syncTimeBoundaryType" (string) An optional string specifying the boundary type to use when applying a syncTime. The allowed values are &#39;Inside&#39; and &#39;Outside&#39;. The default is &#39;Inside&#39;.
    @param "timeZone" (string) The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used.
    @param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.
@return ItemsStreamValues*/
func (a *StreamSetApiService) StreamSetGetInterpolatedAdHoc(ctx context.Context, webId []string, localVarOptionals map[string]interface{}) (ItemsStreamValues, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsStreamValues
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streamsets/interpolated"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["endTime"], "string", "endTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["filterExpression"], "string", "filterExpression"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["includeFilteredValues"], "bool", "includeFilteredValues"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["interval"], "string", "interval"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortField"], "string", "sortField"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortOrder"], "string", "sortOrder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startTime"], "string", "startTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["syncTime"], "string", "syncTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["syncTimeBoundaryType"], "string", "syncTimeBoundaryType"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("webId", parameterToString(webId, "multi"))
	if localVarTempParam, localVarOk := localVarOptionals["endTime"].(string); localVarOk {
		localVarQueryParams.Add("endTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filterExpression"].(string); localVarOk {
		localVarQueryParams.Add("filterExpression", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["includeFilteredValues"].(bool); localVarOk {
		localVarQueryParams.Add("includeFilteredValues", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["interval"].(string); localVarOk {
		localVarQueryParams.Add("interval", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortField"].(string); localVarOk {
		localVarQueryParams.Add("sortField", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortOrder"].(string); localVarOk {
		localVarQueryParams.Add("sortOrder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startTime"].(string); localVarOk {
		localVarQueryParams.Add("startTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["syncTime"].(string); localVarOk {
		localVarQueryParams.Add("syncTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["syncTimeBoundaryType"].(string); localVarOk {
		localVarQueryParams.Add("syncTimeBoundaryType", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeZone"].(string); localVarOk {
		localVarQueryParams.Add("timeZone", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["webIdType"].(string); localVarOk {
		localVarQueryParams.Add("webIdType", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/json",
		"text/html",
		"application/x-ms-application",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* StreamSetApiService Returns interpolated values of attributes for an element, event frame or attribute at the specified times.
Any time series value in the response that contains an &#39;Errors&#39; property indicates PI Web API encountered a handled error during the transfer of the response stream.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of an element, event frame or attribute, which is the base element or parent of all the stream attributes.
@param time The timestamp at which to retrieve a interpolated value. Multiple timestamps may be specified with multiple instances of the parameter.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "categoryName" (string) Specify that included attributes must have this category. The default is no category filter.
    @param "filterExpression" (string) An optional string containing a filter expression. Expression variables are relative to the data point. Use &#39;.&#39; to reference the containing attribute. If the attribute does not support filtering, the filter will be ignored. The default is no filtering.
    @param "includeFilteredValues" (bool) Specify &#39;true&#39; to indicate that values which fail the filter criteria are present in the returned data at the times where they occurred with a value set to a &#39;Filtered&#39; enumeration value with bad status. Repeated consecutive failures are omitted.
    @param "nameFilter" (string) The name query string used for filtering attributes. The default is no filter.
    @param "searchFullHierarchy" (bool) Specifies if the search should include attributes nested further than the immediate attributes of the searchRoot. The default is &#39;false&#39;.
    @param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
    @param "showExcluded" (bool) Specified if the search should include attributes with the Excluded property set. The default is &#39;false&#39;.
    @param "showHidden" (bool) Specified if the search should include attributes with the Hidden property set. The default is &#39;false&#39;.
    @param "sortOrder" (string) The order that the returned collection is sorted. The default is &#39;Ascending&#39;.
    @param "templateName" (string) Specify that included attributes must be members of this template. The default is no template filter.
    @param "timeZone" (string) The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used.
    @param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.
@return ItemsStreamValues*/
func (a *StreamSetApiService) StreamSetGetInterpolatedAtTimes(ctx context.Context, webId string, time []string, localVarOptionals map[string]interface{}) (ItemsStreamValues, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsStreamValues
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streamsets/{webId}/interpolatedattimes"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["categoryName"], "string", "categoryName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["filterExpression"], "string", "filterExpression"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["includeFilteredValues"], "bool", "includeFilteredValues"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["nameFilter"], "string", "nameFilter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["searchFullHierarchy"], "bool", "searchFullHierarchy"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["showExcluded"], "bool", "showExcluded"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["showHidden"], "bool", "showHidden"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortOrder"], "string", "sortOrder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["templateName"], "string", "templateName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("time", parameterToString(time, "multi"))
	if localVarTempParam, localVarOk := localVarOptionals["categoryName"].(string); localVarOk {
		localVarQueryParams.Add("categoryName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filterExpression"].(string); localVarOk {
		localVarQueryParams.Add("filterExpression", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["includeFilteredValues"].(bool); localVarOk {
		localVarQueryParams.Add("includeFilteredValues", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["nameFilter"].(string); localVarOk {
		localVarQueryParams.Add("nameFilter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["searchFullHierarchy"].(bool); localVarOk {
		localVarQueryParams.Add("searchFullHierarchy", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["showExcluded"].(bool); localVarOk {
		localVarQueryParams.Add("showExcluded", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["showHidden"].(bool); localVarOk {
		localVarQueryParams.Add("showHidden", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortOrder"].(string); localVarOk {
		localVarQueryParams.Add("sortOrder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["templateName"].(string); localVarOk {
		localVarQueryParams.Add("templateName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeZone"].(string); localVarOk {
		localVarQueryParams.Add("timeZone", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["webIdType"].(string); localVarOk {
		localVarQueryParams.Add("webIdType", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/json",
		"text/html",
		"application/x-ms-application",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* StreamSetApiService Returns interpolated values of the specified streams at the specified times.
Any time series value in the response that contains an &#39;Errors&#39; property indicates PI Web API encountered a handled error during the transfer of the response stream.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param time The timestamp at which to retrieve a interpolated value. Multiple timestamps may be specified with multiple instances of the parameter.
@param webId The ID of a stream. Multiple streams may be specified with multiple instances of the parameter.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "filterExpression" (string) An optional string containing a filter expression. Expression variables are relative to the data point. Use &#39;.&#39; to reference the containing attribute. If the attribute does not support filtering, the filter will be ignored. The default is no filtering.
    @param "includeFilteredValues" (bool) Specify &#39;true&#39; to indicate that values which fail the filter criteria are present in the returned data at the times where they occurred with a value set to a &#39;Filtered&#39; enumeration value with bad status. Repeated consecutive failures are omitted.
    @param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
    @param "sortOrder" (string) The order that the returned collection is sorted. The default is &#39;Ascending&#39;.
    @param "timeZone" (string) The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used.
    @param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.
@return ItemsStreamValues*/
func (a *StreamSetApiService) StreamSetGetInterpolatedAtTimesAdHoc(ctx context.Context, time []string, webId []string, localVarOptionals map[string]interface{}) (ItemsStreamValues, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsStreamValues
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streamsets/interpolatedattimes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["filterExpression"], "string", "filterExpression"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["includeFilteredValues"], "bool", "includeFilteredValues"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortOrder"], "string", "sortOrder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("time", parameterToString(time, "multi"))
	localVarQueryParams.Add("webId", parameterToString(webId, "multi"))
	if localVarTempParam, localVarOk := localVarOptionals["filterExpression"].(string); localVarOk {
		localVarQueryParams.Add("filterExpression", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["includeFilteredValues"].(bool); localVarOk {
		localVarQueryParams.Add("includeFilteredValues", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortOrder"].(string); localVarOk {
		localVarQueryParams.Add("sortOrder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeZone"].(string); localVarOk {
		localVarQueryParams.Add("timeZone", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["webIdType"].(string); localVarOk {
		localVarQueryParams.Add("webIdType", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/json",
		"text/html",
		"application/x-ms-application",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* StreamSetApiService Returns values of attributes for an element, event frame or attribute over the specified time range suitable for plotting over the number of intervals (typically represents pixels).
For each interval, the data available is examined and significant values are returned. Each interval can produce up to 5 values if they are unique, the first value in the interval, the last value, the highest value, the lowest value and at most one exceptional point (bad status or digital state). Any time series value in the response that contains an &#39;Errors&#39; property indicates PI Web API encountered a handled error during the transfer of the response stream.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of an element, event frame or attribute, which is the base element or parent of all the stream attributes.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "categoryName" (string) Specify that included attributes must have this category. The default is no category filter.
    @param "endTime" (string) An optional end time. The default is &#39;*&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s end time, or &#39;*&#39; if that is not set. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order.
    @param "intervals" (int32) The number of intervals to plot over. Typically, this would be the number of horizontal pixels in the trend. The default is &#39;24&#39;. For each interval, the data available is examined and significant values are returned. Each interval can produce up to 5 values if they are unique, the first value in the interval, the last value, the highest value, the lowest value and at most one exceptional point (bad status or digital state).
    @param "nameFilter" (string) The name query string used for filtering attributes. The default is no filter.
    @param "searchFullHierarchy" (bool) Specifies if the search should include attributes nested further than the immediate attributes of the searchRoot. The default is &#39;false&#39;.
    @param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
    @param "showExcluded" (bool) Specified if the search should include attributes with the Excluded property set. The default is &#39;false&#39;.
    @param "showHidden" (bool) Specified if the search should include attributes with the Hidden property set. The default is &#39;false&#39;.
    @param "sortField" (string) The field or property of the object used to sort the returned collection. For better performance, by default no sorting is applied. &#39;Name&#39; is the only supported field by which to sort.
    @param "sortOrder" (string) The order that the returned collection is sorted. The default is &#39;Ascending&#39;
    @param "startTime" (string) An optional start time. The default is &#39;*-1d&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s start time, or &#39;*-1d&#39; if that is not set.
    @param "templateName" (string) Specify that included attributes must be members of this template. The default is no template filter.
    @param "timeZone" (string) The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used.
    @param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.
@return ItemsStreamValues*/
func (a *StreamSetApiService) StreamSetGetPlot(ctx context.Context, webId string, localVarOptionals map[string]interface{}) (ItemsStreamValues, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsStreamValues
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streamsets/{webId}/plot"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["categoryName"], "string", "categoryName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["endTime"], "string", "endTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["intervals"], "int32", "intervals"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["nameFilter"], "string", "nameFilter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["searchFullHierarchy"], "bool", "searchFullHierarchy"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["showExcluded"], "bool", "showExcluded"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["showHidden"], "bool", "showHidden"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortField"], "string", "sortField"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortOrder"], "string", "sortOrder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startTime"], "string", "startTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["templateName"], "string", "templateName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["categoryName"].(string); localVarOk {
		localVarQueryParams.Add("categoryName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["endTime"].(string); localVarOk {
		localVarQueryParams.Add("endTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["intervals"].(int32); localVarOk {
		localVarQueryParams.Add("intervals", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["nameFilter"].(string); localVarOk {
		localVarQueryParams.Add("nameFilter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["searchFullHierarchy"].(bool); localVarOk {
		localVarQueryParams.Add("searchFullHierarchy", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["showExcluded"].(bool); localVarOk {
		localVarQueryParams.Add("showExcluded", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["showHidden"].(bool); localVarOk {
		localVarQueryParams.Add("showHidden", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortField"].(string); localVarOk {
		localVarQueryParams.Add("sortField", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortOrder"].(string); localVarOk {
		localVarQueryParams.Add("sortOrder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startTime"].(string); localVarOk {
		localVarQueryParams.Add("startTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["templateName"].(string); localVarOk {
		localVarQueryParams.Add("templateName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeZone"].(string); localVarOk {
		localVarQueryParams.Add("timeZone", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["webIdType"].(string); localVarOk {
		localVarQueryParams.Add("webIdType", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/json",
		"text/html",
		"application/x-ms-application",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* StreamSetApiService Returns values of attributes for the specified streams over the specified time range suitable for plotting over the number of intervals (typically represents pixels).
For each interval, the data available is examined and significant values are returned. Each interval can produce up to 5 values if they are unique, the first value in the interval, the last value, the highest value, the lowest value and at most one exceptional point (bad status or digital state). Any time series value in the response that contains an &#39;Errors&#39; property indicates PI Web API encountered a handled error during the transfer of the response stream.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of a stream. Multiple streams may be specified with multiple instances of the parameter.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "endTime" (string) An optional end time. The default is &#39;*&#39;. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order.
    @param "intervals" (int32) The number of intervals to plot over. Typically, this would be the number of horizontal pixels in the trend. The default is &#39;24&#39;. For each interval, the data available is examined and significant values are returned. Each interval can produce up to 5 values if they are unique, the first value in the interval, the last value, the highest value, the lowest value and at most one exceptional point (bad status or digital state).
    @param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
    @param "sortField" (string) The field or property of the object used to sort the returned collection. For better performance, by default no sorting is applied. &#39;Name&#39; is the only supported field by which to sort.
    @param "sortOrder" (string) The order that the returned collection is sorted. The default is &#39;Ascending&#39;
    @param "startTime" (string) An optional start time. The default is &#39;*-1d&#39;.
    @param "timeZone" (string) The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used.
    @param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.
@return ItemsStreamValues*/
func (a *StreamSetApiService) StreamSetGetPlotAdHoc(ctx context.Context, webId []string, localVarOptionals map[string]interface{}) (ItemsStreamValues, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsStreamValues
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streamsets/plot"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["endTime"], "string", "endTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["intervals"], "int32", "intervals"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortField"], "string", "sortField"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortOrder"], "string", "sortOrder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startTime"], "string", "startTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("webId", parameterToString(webId, "multi"))
	if localVarTempParam, localVarOk := localVarOptionals["endTime"].(string); localVarOk {
		localVarQueryParams.Add("endTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["intervals"].(int32); localVarOk {
		localVarQueryParams.Add("intervals", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortField"].(string); localVarOk {
		localVarQueryParams.Add("sortField", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortOrder"].(string); localVarOk {
		localVarQueryParams.Add("sortOrder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startTime"].(string); localVarOk {
		localVarQueryParams.Add("startTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeZone"].(string); localVarOk {
		localVarQueryParams.Add("timeZone", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["webIdType"].(string); localVarOk {
		localVarQueryParams.Add("webIdType", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/json",
		"text/html",
		"application/x-ms-application",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* StreamSetApiService Returns recorded values of the attributes for an element, event frame, or attribute.
Any time series value in the response that contains an &#39;Errors&#39; property indicates PI Web API encountered a handled error during the transfer of the response stream.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of an element, event frame or attribute, which is the base element or parent of all the stream attributes.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "boundaryType" (string) An optional value that determines how the times and values of the returned end points are determined. The default is &#39;Inside&#39;.
    @param "categoryName" (string) Specify that included attributes must have this category. The default is no category filter.
    @param "endTime" (string) An optional end time. The default is &#39;*&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s end time, or &#39;*&#39; if that is not set. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order.
    @param "filterExpression" (string) An optional string containing a filter expression. Expression variables are relative to the data point. Use &#39;.&#39; to reference the containing attribute. The default is no filtering.
    @param "includeFilteredValues" (bool) Specify &#39;true&#39; to indicate that values which fail the filter criteria are present in the returned data at the times where they occurred with a value set to a &#39;Filtered&#39; enumeration value with bad status. Repeated consecutive failures are omitted.
    @param "maxCount" (int32) The maximum number of values to be returned. The default is 1000.
    @param "nameFilter" (string) The name query string used for filtering attributes. The default is no filter.
    @param "searchFullHierarchy" (bool) Specifies if the search should include attributes nested further than the immediate attributes of the searchRoot. The default is &#39;false&#39;.
    @param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
    @param "showExcluded" (bool) Specified if the search should include attributes with the Excluded property set. The default is &#39;false&#39;.
    @param "showHidden" (bool) Specified if the search should include attributes with the Hidden property set. The default is &#39;false&#39;.
    @param "sortField" (string) The field or property of the object used to sort the returned collection. For better performance, by default no sorting is applied. &#39;Name&#39; is the only supported field by which to sort.
    @param "sortOrder" (string) The order that the returned collection is sorted. The default is &#39;Ascending&#39;
    @param "startTime" (string) An optional start time. The default is &#39;*-1d&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s start time, or &#39;*-1d&#39; if that is not set.
    @param "templateName" (string) Specify that included attributes must be members of this template. The default is no template filter.
    @param "timeZone" (string) The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used.
    @param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.
@return ItemsStreamValues*/
func (a *StreamSetApiService) StreamSetGetRecorded(ctx context.Context, webId string, localVarOptionals map[string]interface{}) (ItemsStreamValues, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsStreamValues
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streamsets/{webId}/recorded"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["boundaryType"], "string", "boundaryType"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["categoryName"], "string", "categoryName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["endTime"], "string", "endTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["filterExpression"], "string", "filterExpression"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["includeFilteredValues"], "bool", "includeFilteredValues"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["maxCount"], "int32", "maxCount"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["nameFilter"], "string", "nameFilter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["searchFullHierarchy"], "bool", "searchFullHierarchy"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["showExcluded"], "bool", "showExcluded"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["showHidden"], "bool", "showHidden"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortField"], "string", "sortField"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortOrder"], "string", "sortOrder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startTime"], "string", "startTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["templateName"], "string", "templateName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["boundaryType"].(string); localVarOk {
		localVarQueryParams.Add("boundaryType", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["categoryName"].(string); localVarOk {
		localVarQueryParams.Add("categoryName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["endTime"].(string); localVarOk {
		localVarQueryParams.Add("endTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filterExpression"].(string); localVarOk {
		localVarQueryParams.Add("filterExpression", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["includeFilteredValues"].(bool); localVarOk {
		localVarQueryParams.Add("includeFilteredValues", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["maxCount"].(int32); localVarOk {
		localVarQueryParams.Add("maxCount", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["nameFilter"].(string); localVarOk {
		localVarQueryParams.Add("nameFilter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["searchFullHierarchy"].(bool); localVarOk {
		localVarQueryParams.Add("searchFullHierarchy", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["showExcluded"].(bool); localVarOk {
		localVarQueryParams.Add("showExcluded", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["showHidden"].(bool); localVarOk {
		localVarQueryParams.Add("showHidden", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortField"].(string); localVarOk {
		localVarQueryParams.Add("sortField", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortOrder"].(string); localVarOk {
		localVarQueryParams.Add("sortOrder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startTime"].(string); localVarOk {
		localVarQueryParams.Add("startTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["templateName"].(string); localVarOk {
		localVarQueryParams.Add("templateName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeZone"].(string); localVarOk {
		localVarQueryParams.Add("timeZone", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["webIdType"].(string); localVarOk {
		localVarQueryParams.Add("webIdType", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/json",
		"text/html",
		"application/x-ms-application",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* StreamSetApiService Returns recorded values of the specified streams.
Any time series value in the response that contains an &#39;Errors&#39; property indicates PI Web API encountered a handled error during the transfer of the response stream.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of a stream. Multiple streams may be specified with multiple instances of the parameter.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "boundaryType" (string) An optional value that determines how the times and values of the returned end points are determined. The default is &#39;Inside&#39;.
    @param "endTime" (string) An optional end time. The default is &#39;*&#39;. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order.
    @param "filterExpression" (string) An optional string containing a filter expression. Expression variables are relative to the data point. Use &#39;.&#39; to reference the containing attribute. The default is no filtering.
    @param "includeFilteredValues" (bool) Specify &#39;true&#39; to indicate that values which fail the filter criteria are present in the returned data at the times where they occurred with a value set to a &#39;Filtered&#39; enumeration value with bad status. Repeated consecutive failures are omitted.
    @param "maxCount" (int32) The maximum number of values to be returned. The default is 1000.
    @param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
    @param "sortField" (string) The field or property of the object used to sort the returned collection. For better performance, by default no sorting is applied. &#39;Name&#39; is the only supported field by which to sort.
    @param "sortOrder" (string) The order that the returned collection is sorted. The default is &#39;Ascending&#39;
    @param "startTime" (string) An optional start time. The default is &#39;*-1d&#39;.
    @param "timeZone" (string) The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used.
    @param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.
@return ItemsStreamValues*/
func (a *StreamSetApiService) StreamSetGetRecordedAdHoc(ctx context.Context, webId []string, localVarOptionals map[string]interface{}) (ItemsStreamValues, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsStreamValues
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streamsets/recorded"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["boundaryType"], "string", "boundaryType"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["endTime"], "string", "endTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["filterExpression"], "string", "filterExpression"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["includeFilteredValues"], "bool", "includeFilteredValues"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["maxCount"], "int32", "maxCount"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortField"], "string", "sortField"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortOrder"], "string", "sortOrder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startTime"], "string", "startTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("webId", parameterToString(webId, "multi"))
	if localVarTempParam, localVarOk := localVarOptionals["boundaryType"].(string); localVarOk {
		localVarQueryParams.Add("boundaryType", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["endTime"].(string); localVarOk {
		localVarQueryParams.Add("endTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filterExpression"].(string); localVarOk {
		localVarQueryParams.Add("filterExpression", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["includeFilteredValues"].(bool); localVarOk {
		localVarQueryParams.Add("includeFilteredValues", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["maxCount"].(int32); localVarOk {
		localVarQueryParams.Add("maxCount", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortField"].(string); localVarOk {
		localVarQueryParams.Add("sortField", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortOrder"].(string); localVarOk {
		localVarQueryParams.Add("sortOrder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startTime"].(string); localVarOk {
		localVarQueryParams.Add("startTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeZone"].(string); localVarOk {
		localVarQueryParams.Add("timeZone", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["webIdType"].(string); localVarOk {
		localVarQueryParams.Add("webIdType", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/json",
		"text/html",
		"application/x-ms-application",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* StreamSetApiService Returns recorded values of the attributes for an element, event frame, or attribute.
Any time series value in the response that contains an &#39;Errors&#39; property indicates PI Web API encountered a handled error during the transfer of the response stream.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of an element, event frame or attribute, which is the base element or parent of all the stream attributes.
@param time The timestamp at which the values are desired.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "categoryName" (string) Specify that included attributes must have this category. The default is no category filter.
    @param "nameFilter" (string) The name query string used for filtering attributes. The default is no filter.
    @param "retrievalMode" (string) An optional value that determines the values to return when values don&#39;t exist at the exact time specified. The default is &#39;Auto&#39;.
    @param "searchFullHierarchy" (bool) Specifies if the search should include attributes nested further than the immediate attributes of the searchRoot. The default is &#39;false&#39;.
    @param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
    @param "showExcluded" (bool) Specified if the search should include attributes with the Excluded property set. The default is &#39;false&#39;.
    @param "showHidden" (bool) Specified if the search should include attributes with the Hidden property set. The default is &#39;false&#39;.
    @param "templateName" (string) Specify that included attributes must be members of this template. The default is no template filter.
    @param "timeZone" (string) The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used.
    @param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.
@return ItemsStreamValue*/
func (a *StreamSetApiService) StreamSetGetRecordedAtTime(ctx context.Context, webId string, time string, localVarOptionals map[string]interface{}) (ItemsStreamValue, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsStreamValue
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streamsets/{webId}/recordedattime"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["categoryName"], "string", "categoryName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["nameFilter"], "string", "nameFilter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["retrievalMode"], "string", "retrievalMode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["searchFullHierarchy"], "bool", "searchFullHierarchy"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["showExcluded"], "bool", "showExcluded"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["showHidden"], "bool", "showHidden"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["templateName"], "string", "templateName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("time", parameterToString(time, ""))
	if localVarTempParam, localVarOk := localVarOptionals["categoryName"].(string); localVarOk {
		localVarQueryParams.Add("categoryName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["nameFilter"].(string); localVarOk {
		localVarQueryParams.Add("nameFilter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["retrievalMode"].(string); localVarOk {
		localVarQueryParams.Add("retrievalMode", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["searchFullHierarchy"].(bool); localVarOk {
		localVarQueryParams.Add("searchFullHierarchy", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["showExcluded"].(bool); localVarOk {
		localVarQueryParams.Add("showExcluded", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["showHidden"].(bool); localVarOk {
		localVarQueryParams.Add("showHidden", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["templateName"].(string); localVarOk {
		localVarQueryParams.Add("templateName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeZone"].(string); localVarOk {
		localVarQueryParams.Add("timeZone", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["webIdType"].(string); localVarOk {
		localVarQueryParams.Add("webIdType", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/json",
		"text/html",
		"application/x-ms-application",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* StreamSetApiService Returns recorded values based on the passed time and retrieval mode.
Any time series value in the response that contains an &#39;Errors&#39; property indicates PI Web API encountered a handled error during the transfer of the response stream.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param time The timestamp at which the values are desired.
@param webId The ID of a stream. Multiple streams may be specified with multiple instances of the parameter.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "retrievalMode" (string) An optional value that determines the values to return when values don&#39;t exist at the exact time specified. The default is &#39;Auto&#39;.
    @param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
    @param "timeZone" (string) The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used.
    @param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.
@return ItemsStreamValue*/
func (a *StreamSetApiService) StreamSetGetRecordedAtTimeAdHoc(ctx context.Context, time string, webId []string, localVarOptionals map[string]interface{}) (ItemsStreamValue, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsStreamValue
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streamsets/recordedattime"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["retrievalMode"], "string", "retrievalMode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("time", parameterToString(time, ""))
	localVarQueryParams.Add("webId", parameterToString(webId, "multi"))
	if localVarTempParam, localVarOk := localVarOptionals["retrievalMode"].(string); localVarOk {
		localVarQueryParams.Add("retrievalMode", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeZone"].(string); localVarOk {
		localVarQueryParams.Add("timeZone", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["webIdType"].(string); localVarOk {
		localVarQueryParams.Add("webIdType", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/json",
		"text/html",
		"application/x-ms-application",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* StreamSetApiService Returns recorded values of attributes for an element, event frame or attribute at the specified times.
Any time series value in the response that contains an &#39;Errors&#39; property indicates PI Web API encountered a handled error during the transfer of the response stream.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of an element, event frame or attribute, which is the base element or parent of all the stream attributes.
@param time The timestamp at which to retrieve a recorded value. Multiple timestamps may be specified with multiple instances of the parameter.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "categoryName" (string) Specify that included attributes must have this category. The default is no category filter.
    @param "nameFilter" (string) The name query string used for filtering attributes. The default is no filter.
    @param "retrievalMode" (string) An optional value that determines the values to return when values don&#39;t exist at the exact time specified. The default is &#39;Auto&#39;.
    @param "searchFullHierarchy" (bool) Specifies if the search should include attributes nested further than the immediate attributes of the searchRoot. The default is &#39;false&#39;.
    @param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
    @param "showExcluded" (bool) Specified if the search should include attributes with the Excluded property set. The default is &#39;false&#39;.
    @param "showHidden" (bool) Specified if the search should include attributes with the Hidden property set. The default is &#39;false&#39;.
    @param "sortOrder" (string) The order that the returned collection is sorted. The default is &#39;Ascending&#39;.
    @param "templateName" (string) Specify that included attributes must be members of this template. The default is no template filter.
    @param "timeZone" (string) The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used.
    @param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.
@return ItemsStreamValues*/
func (a *StreamSetApiService) StreamSetGetRecordedAtTimes(ctx context.Context, webId string, time []string, localVarOptionals map[string]interface{}) (ItemsStreamValues, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsStreamValues
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streamsets/{webId}/recordedattimes"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["categoryName"], "string", "categoryName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["nameFilter"], "string", "nameFilter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["retrievalMode"], "string", "retrievalMode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["searchFullHierarchy"], "bool", "searchFullHierarchy"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["showExcluded"], "bool", "showExcluded"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["showHidden"], "bool", "showHidden"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortOrder"], "string", "sortOrder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["templateName"], "string", "templateName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("time", parameterToString(time, "multi"))
	if localVarTempParam, localVarOk := localVarOptionals["categoryName"].(string); localVarOk {
		localVarQueryParams.Add("categoryName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["nameFilter"].(string); localVarOk {
		localVarQueryParams.Add("nameFilter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["retrievalMode"].(string); localVarOk {
		localVarQueryParams.Add("retrievalMode", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["searchFullHierarchy"].(bool); localVarOk {
		localVarQueryParams.Add("searchFullHierarchy", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["showExcluded"].(bool); localVarOk {
		localVarQueryParams.Add("showExcluded", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["showHidden"].(bool); localVarOk {
		localVarQueryParams.Add("showHidden", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortOrder"].(string); localVarOk {
		localVarQueryParams.Add("sortOrder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["templateName"].(string); localVarOk {
		localVarQueryParams.Add("templateName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeZone"].(string); localVarOk {
		localVarQueryParams.Add("timeZone", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["webIdType"].(string); localVarOk {
		localVarQueryParams.Add("webIdType", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/json",
		"text/html",
		"application/x-ms-application",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* StreamSetApiService Returns recorded values of the specified streams at the specified times.
Any time series value in the response that contains an &#39;Errors&#39; property indicates PI Web API encountered a handled error during the transfer of the response stream.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param time The timestamp at which to retrieve a recorded value. Multiple timestamps may be specified with multiple instances of the parameter.
@param webId The ID of a stream. Multiple streams may be specified with multiple instances of the parameter.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "retrievalMode" (string) An optional value that determines the values to return when values don&#39;t exist at the exact time specified. The default is &#39;Auto&#39;.
    @param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
    @param "sortOrder" (string) The order that the returned collection is sorted. The default is &#39;Ascending&#39;.
    @param "timeZone" (string) The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used.
    @param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.
@return ItemsStreamValues*/
func (a *StreamSetApiService) StreamSetGetRecordedAtTimesAdHoc(ctx context.Context, time []string, webId []string, localVarOptionals map[string]interface{}) (ItemsStreamValues, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsStreamValues
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streamsets/recordedattimes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["retrievalMode"], "string", "retrievalMode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortOrder"], "string", "sortOrder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("time", parameterToString(time, "multi"))
	localVarQueryParams.Add("webId", parameterToString(webId, "multi"))
	if localVarTempParam, localVarOk := localVarOptionals["retrievalMode"].(string); localVarOk {
		localVarQueryParams.Add("retrievalMode", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortOrder"].(string); localVarOk {
		localVarQueryParams.Add("sortOrder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeZone"].(string); localVarOk {
		localVarQueryParams.Add("timeZone", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["webIdType"].(string); localVarOk {
		localVarQueryParams.Add("webIdType", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/json",
		"text/html",
		"application/x-ms-application",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* StreamSetApiService Returns summary values of the attributes for an element, event frame or attribute.
Any time series value in the response that contains an &#39;Errors&#39; property indicates PI Web API encountered a handled error during the transfer of the response stream.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of an element, event frame or attribute, which is the base element or parent of all the stream attributes.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "calculationBasis" (string) Specifies the method of evaluating the data over the time range. The default is &#39;TimeWeighted&#39;.
    @param "categoryName" (string) Specify that included attributes must have this category. The default is no category filter.
    @param "endTime" (string) An optional end time. The default is &#39;*&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s end time, or &#39;*&#39; if that is not set. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order.
    @param "filterExpression" (string) A string containing a filter expression. Expression variables are relative to the attribute. Use &#39;.&#39; to reference the containing attribute. The default is no filtering.
    @param "nameFilter" (string) The name query string used for filtering attributes. The default is no filter.
    @param "sampleInterval" (string) A time span specifies how often the filter expression is evaluated when computing the summary for an interval, if the sampleType is &#39;Interval&#39;.
    @param "sampleType" (string) A flag which specifies one or more summaries to compute for each interval over the time range. The default is &#39;ExpressionRecordedValues&#39;.
    @param "searchFullHierarchy" (bool) Specifies if the search should include attributes nested further than the immediate attributes of the searchRoot. The default is &#39;false&#39;.
    @param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
    @param "showExcluded" (bool) Specified if the search should include attributes with the Excluded property set. The default is &#39;false&#39;.
    @param "showHidden" (bool) Specified if the search should include attributes with the Hidden property set. The default is &#39;false&#39;.
    @param "startTime" (string) An optional start time. The default is &#39;*-1d&#39; for element attributes and points. For event frame attributes, the default is the event frame&#39;s start time, or &#39;*-1d&#39; if that is not set.
    @param "summaryDuration" (string) The duration of each summary interval.
    @param "summaryType" ([]string) Specifies the kinds of summaries to produce over the range. The default is &#39;Total&#39;. Multiple summary types may be specified by using multiple instances of summaryType.
    @param "templateName" (string) Specify that included attributes must be members of this template. The default is no template filter.
    @param "timeType" (string) Specifies how to calculate the timestamp for each interval. The default is &#39;Auto&#39;.
    @param "timeZone" (string) The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used.
    @param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.
@return ItemsStreamSummaries*/
func (a *StreamSetApiService) StreamSetGetSummaries(ctx context.Context, webId string, localVarOptionals map[string]interface{}) (ItemsStreamSummaries, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsStreamSummaries
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streamsets/{webId}/summary"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["calculationBasis"], "string", "calculationBasis"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["categoryName"], "string", "categoryName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["endTime"], "string", "endTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["filterExpression"], "string", "filterExpression"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["nameFilter"], "string", "nameFilter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sampleInterval"], "string", "sampleInterval"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sampleType"], "string", "sampleType"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["searchFullHierarchy"], "bool", "searchFullHierarchy"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["showExcluded"], "bool", "showExcluded"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["showHidden"], "bool", "showHidden"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startTime"], "string", "startTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["summaryDuration"], "string", "summaryDuration"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["templateName"], "string", "templateName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["timeType"], "string", "timeType"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["calculationBasis"].(string); localVarOk {
		localVarQueryParams.Add("calculationBasis", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["categoryName"].(string); localVarOk {
		localVarQueryParams.Add("categoryName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["endTime"].(string); localVarOk {
		localVarQueryParams.Add("endTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filterExpression"].(string); localVarOk {
		localVarQueryParams.Add("filterExpression", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["nameFilter"].(string); localVarOk {
		localVarQueryParams.Add("nameFilter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sampleInterval"].(string); localVarOk {
		localVarQueryParams.Add("sampleInterval", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sampleType"].(string); localVarOk {
		localVarQueryParams.Add("sampleType", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["searchFullHierarchy"].(bool); localVarOk {
		localVarQueryParams.Add("searchFullHierarchy", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["showExcluded"].(bool); localVarOk {
		localVarQueryParams.Add("showExcluded", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["showHidden"].(bool); localVarOk {
		localVarQueryParams.Add("showHidden", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startTime"].(string); localVarOk {
		localVarQueryParams.Add("startTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["summaryDuration"].(string); localVarOk {
		localVarQueryParams.Add("summaryDuration", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["summaryType"].([]string); localVarOk {
		localVarQueryParams.Add("summaryType", parameterToString(localVarTempParam, "multi"))
	}
	if localVarTempParam, localVarOk := localVarOptionals["templateName"].(string); localVarOk {
		localVarQueryParams.Add("templateName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeType"].(string); localVarOk {
		localVarQueryParams.Add("timeType", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeZone"].(string); localVarOk {
		localVarQueryParams.Add("timeZone", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["webIdType"].(string); localVarOk {
		localVarQueryParams.Add("webIdType", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/json",
		"text/html",
		"application/x-ms-application",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* StreamSetApiService Returns summary values of the specified streams.
Any time series value in the response that contains an &#39;Errors&#39; property indicates PI Web API encountered a handled error during the transfer of the response stream.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of a stream. Multiple streams may be specified with multiple instances of the parameter.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "calculationBasis" (string) Specifies the method of evaluating the data over the time range. The default is &#39;TimeWeighted&#39;.
    @param "endTime" (string) An optional end time. The default is &#39;*&#39;. Note that if endTime is earlier than startTime, the resulting values will be in time-descending order.
    @param "filterExpression" (string) A string containing a filter expression. Expression variables are relative to the attribute. Use &#39;.&#39; to reference the containing attribute. The default is no filtering.
    @param "sampleInterval" (string) A time span specifies how often the filter expression is evaluated when computing the summary for an interval, if the sampleType is &#39;Interval&#39;.
    @param "sampleType" (string) A flag which specifies one or more summaries to compute for each interval over the time range. The default is &#39;ExpressionRecordedValues&#39;.
    @param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
    @param "startTime" (string) An optional start time. The default is &#39;*-1d&#39;.
    @param "summaryDuration" (string) The duration of each summary interval.
    @param "summaryType" ([]string) Specifies the kinds of summaries to produce over the range. The default is &#39;Total&#39;. Multiple summary types may be specified by using multiple instances of summaryType.
    @param "timeType" (string) Specifies how to calculate the timestamp for each interval. The default is &#39;Auto&#39;.
    @param "timeZone" (string) The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used.
    @param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.
@return ItemsStreamSummaries*/
func (a *StreamSetApiService) StreamSetGetSummariesAdHoc(ctx context.Context, webId []string, localVarOptionals map[string]interface{}) (ItemsStreamSummaries, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsStreamSummaries
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streamsets/summary"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["calculationBasis"], "string", "calculationBasis"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["endTime"], "string", "endTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["filterExpression"], "string", "filterExpression"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sampleInterval"], "string", "sampleInterval"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sampleType"], "string", "sampleType"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startTime"], "string", "startTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["summaryDuration"], "string", "summaryDuration"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["timeType"], "string", "timeType"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("webId", parameterToString(webId, "multi"))
	if localVarTempParam, localVarOk := localVarOptionals["calculationBasis"].(string); localVarOk {
		localVarQueryParams.Add("calculationBasis", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["endTime"].(string); localVarOk {
		localVarQueryParams.Add("endTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filterExpression"].(string); localVarOk {
		localVarQueryParams.Add("filterExpression", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sampleInterval"].(string); localVarOk {
		localVarQueryParams.Add("sampleInterval", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sampleType"].(string); localVarOk {
		localVarQueryParams.Add("sampleType", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startTime"].(string); localVarOk {
		localVarQueryParams.Add("startTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["summaryDuration"].(string); localVarOk {
		localVarQueryParams.Add("summaryDuration", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["summaryType"].([]string); localVarOk {
		localVarQueryParams.Add("summaryType", parameterToString(localVarTempParam, "multi"))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeType"].(string); localVarOk {
		localVarQueryParams.Add("timeType", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeZone"].(string); localVarOk {
		localVarQueryParams.Add("timeZone", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["webIdType"].(string); localVarOk {
		localVarQueryParams.Add("webIdType", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/json",
		"text/html",
		"application/x-ms-application",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* StreamSetApiService Returns values of the attributes for an Element, Event Frame or Attribute at the specified time.
Any time series value in the response that contains an &#39;Errors&#39; property indicates PI Web API encountered a handled error during the transfer of the response stream.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of an Element, Event Frame or Attribute, which is the base element or parent of all the stream attributes.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "categoryName" (string) Specify that included attributes must have this category. The default is no category filter.
    @param "nameFilter" (string) The name query string used for filtering attributes. The default is no filter.
    @param "searchFullHierarchy" (bool) Specifies if the search should include attributes nested further than the immediate attributes of the searchRoot. The default is &#39;false&#39;.
    @param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
    @param "showExcluded" (bool) Specified if the search should include attributes with the Excluded property set. The default is &#39;false&#39;.
    @param "showHidden" (bool) Specified if the search should include attributes with the Hidden property set. The default is &#39;false&#39;.
    @param "sortField" (string) The field or property of the object used to sort the returned collection. For better performance, by default no sorting is applied. &#39;Name&#39; is the only supported field by which to sort.
    @param "sortOrder" (string) The order that the returned collection is sorted. The default is &#39;Ascending&#39;
    @param "templateName" (string) Specify that included attributes must be members of this template. The default is no template filter.
    @param "time" (string) An AF time string, which is used as the time context to get stream values if it is provided. By default, it is not specified, and the default time context of the AF object will be used.
    @param "timeZone" (string) The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used.
    @param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.
@return ItemsStreamValue*/
func (a *StreamSetApiService) StreamSetGetValues(ctx context.Context, webId string, localVarOptionals map[string]interface{}) (ItemsStreamValue, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsStreamValue
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streamsets/{webId}/value"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["categoryName"], "string", "categoryName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["nameFilter"], "string", "nameFilter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["searchFullHierarchy"], "bool", "searchFullHierarchy"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["showExcluded"], "bool", "showExcluded"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["showHidden"], "bool", "showHidden"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortField"], "string", "sortField"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortOrder"], "string", "sortOrder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["templateName"], "string", "templateName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["time"], "string", "time"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["categoryName"].(string); localVarOk {
		localVarQueryParams.Add("categoryName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["nameFilter"].(string); localVarOk {
		localVarQueryParams.Add("nameFilter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["searchFullHierarchy"].(bool); localVarOk {
		localVarQueryParams.Add("searchFullHierarchy", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["showExcluded"].(bool); localVarOk {
		localVarQueryParams.Add("showExcluded", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["showHidden"].(bool); localVarOk {
		localVarQueryParams.Add("showHidden", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortField"].(string); localVarOk {
		localVarQueryParams.Add("sortField", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortOrder"].(string); localVarOk {
		localVarQueryParams.Add("sortOrder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["templateName"].(string); localVarOk {
		localVarQueryParams.Add("templateName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["time"].(string); localVarOk {
		localVarQueryParams.Add("time", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeZone"].(string); localVarOk {
		localVarQueryParams.Add("timeZone", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["webIdType"].(string); localVarOk {
		localVarQueryParams.Add("webIdType", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/json",
		"text/html",
		"application/x-ms-application",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* StreamSetApiService Returns values of the specified streams.
Any time series value in the response that contains an &#39;Errors&#39; property indicates PI Web API encountered a handled error during the transfer of the response stream.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of a stream. Multiple streams may be specified with multiple instances of the parameter.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
    @param "sortField" (string) The field or property of the object used to sort the returned collection. For better performance, by default no sorting is applied. &#39;Name&#39; is the only supported field by which to sort.
    @param "sortOrder" (string) The order that the returned collection is sorted. The default is &#39;Ascending&#39;
    @param "time" (string) An AF time string, which is used as the time context to get stream values if it is provided. By default, it is not specified, and the default time context of the AF object will be used.
    @param "timeZone" (string) The time zone in which the time string will be interpreted. This parameter will be ignored if a time zone is specified in the time string. If no time zone is specified in either places, the PI Web API server time zone will be used.
    @param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.
@return ItemsStreamValue*/
func (a *StreamSetApiService) StreamSetGetValuesAdHoc(ctx context.Context, webId []string, localVarOptionals map[string]interface{}) (ItemsStreamValue, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsStreamValue
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streamsets/value"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortField"], "string", "sortField"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sortOrder"], "string", "sortOrder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["time"], "string", "time"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["timeZone"], "string", "timeZone"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("webId", parameterToString(webId, "multi"))
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortField"].(string); localVarOk {
		localVarQueryParams.Add("sortField", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortOrder"].(string); localVarOk {
		localVarQueryParams.Add("sortOrder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["time"].(string); localVarOk {
		localVarQueryParams.Add("time", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["timeZone"].(string); localVarOk {
		localVarQueryParams.Add("timeZone", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["webIdType"].(string); localVarOk {
		localVarQueryParams.Add("webIdType", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/json",
		"text/html",
		"application/x-ms-application",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* StreamSetApiService Updates a single value for the specified streams.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the parent element, event frame, or attribute. Attributes specified in the body must be descendants of the specified object.
@param values The values to add or update.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "bufferOption" (string) The desired AFBufferOption. The default is &#39;BufferIfPossible&#39;.
    @param "updateOption" (string) The desired AFUpdateOption. The default is &#39;Replace&#39;.
@return ItemsSubstatus*/
func (a *StreamSetApiService) StreamSetUpdateValue(ctx context.Context, webId string, values []StreamValue, localVarOptionals map[string]interface{}) (ItemsSubstatus, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsSubstatus
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streamsets/{webId}/value"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["bufferOption"], "string", "bufferOption"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["updateOption"], "string", "updateOption"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["bufferOption"].(string); localVarOk {
		localVarQueryParams.Add("bufferOption", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["updateOption"].(string); localVarOk {
		localVarQueryParams.Add("updateOption", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "text/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/json",
		"text/html",
		"application/x-ms-application",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &values
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* StreamSetApiService Updates a single value for the specified streams.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param values The values to add or update.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "bufferOption" (string) The desired AFBufferOption. The default is &#39;BufferIfPossible&#39;.
    @param "updateOption" (string) The desired AFUpdateOption. The default is &#39;Replace&#39;.
@return ItemsSubstatus*/
func (a *StreamSetApiService) StreamSetUpdateValueAdHoc(ctx context.Context, values []StreamValue, localVarOptionals map[string]interface{}) (ItemsSubstatus, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsSubstatus
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streamsets/value"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["bufferOption"], "string", "bufferOption"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["updateOption"], "string", "updateOption"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["bufferOption"].(string); localVarOk {
		localVarQueryParams.Add("bufferOption", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["updateOption"].(string); localVarOk {
		localVarQueryParams.Add("updateOption", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "text/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/json",
		"text/html",
		"application/x-ms-application",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &values
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* StreamSetApiService Updates multiple values for the specified streams.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the parent element, event frame, or attribute. Attributes specified in the body must be descendants of the specified object.
@param values The values to add or update.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "bufferOption" (string) The desired AFBufferOption. The default is &#39;BufferIfPossible&#39;.
    @param "updateOption" (string) The desired AFUpdateOption. The default is &#39;Replace&#39;.
@return ItemsItemsSubstatus*/
func (a *StreamSetApiService) StreamSetUpdateValues(ctx context.Context, webId string, values []StreamValues, localVarOptionals map[string]interface{}) (ItemsItemsSubstatus, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsItemsSubstatus
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streamsets/{webId}/recorded"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["bufferOption"], "string", "bufferOption"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["updateOption"], "string", "updateOption"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["bufferOption"].(string); localVarOk {
		localVarQueryParams.Add("bufferOption", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["updateOption"].(string); localVarOk {
		localVarQueryParams.Add("updateOption", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "text/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/json",
		"text/html",
		"application/x-ms-application",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &values
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* StreamSetApiService Updates multiple values for the specified streams.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param values The values to add or update.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "bufferOption" (string) The desired AFBufferOption. The default is &#39;BufferIfPossible&#39;.
    @param "updateOption" (string) The desired AFUpdateOption. The default is &#39;Replace&#39;.
@return ItemsItemsSubstatus*/
func (a *StreamSetApiService) StreamSetUpdateValuesAdHoc(ctx context.Context, values []StreamValues, localVarOptionals map[string]interface{}) (ItemsItemsSubstatus, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsItemsSubstatus
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/streamsets/recorded"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["bufferOption"], "string", "bufferOption"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["updateOption"], "string", "updateOption"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["bufferOption"].(string); localVarOk {
		localVarQueryParams.Add("bufferOption", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["updateOption"].(string); localVarOk {
		localVarQueryParams.Add("updateOption", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "text/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"text/json",
		"text/html",
		"application/x-ms-application",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &values
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}
