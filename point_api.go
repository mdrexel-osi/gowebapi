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

type PointApiService service

/* PointApiService Delete a point.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the point.
@return */
func (a *PointApiService) PointDelete(ctx context.Context, webId string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/points/{webId}"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

/* PointApiService Get a point.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the point.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
    @param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.
@return Point*/
func (a *PointApiService) PointGet(ctx context.Context, webId string, localVarOptionals map[string]interface{}) (Point, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     Point
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/points/{webId}"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
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

/* PointApiService Get a point attribute by name.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param name The name of the attribute.
@param webId The ID of the point.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
    @param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.
@return PointAttribute*/
func (a *PointApiService) PointGetAttributeByName(ctx context.Context, name string, webId string, localVarOptionals map[string]interface{}) (PointAttribute, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     PointAttribute
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/points/{webId}/attributes/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
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

/* PointApiService Get point attributes.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the point.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "name" ([]string) The name of a point attribute to be returned. Multiple attributes may be specified with multiple instances of the parameter.
    @param "nameFilter" (string) The filter to the names of the list of point attributes to be returned. The default is no filter.
    @param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
    @param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.
@return ItemsPointAttribute*/
func (a *PointApiService) PointGetAttributes(ctx context.Context, webId string, localVarOptionals map[string]interface{}) (ItemsPointAttribute, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsPointAttribute
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/points/{webId}/attributes"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["nameFilter"], "string", "nameFilter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["name"].([]string); localVarOk {
		localVarQueryParams.Add("name", parameterToString(localVarTempParam, "multi"))
	}
	if localVarTempParam, localVarOk := localVarOptionals["nameFilter"].(string); localVarOk {
		localVarQueryParams.Add("nameFilter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
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

/* PointApiService Get a point by path.
This method returns a PI Point based on the hierarchical path associated with it, and should be used when a path has been received from a separate part of the PI System for use in the PI Web API. Users should primarily search with the WebID when available.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param path The path to the point.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
    @param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.
@return Point*/
func (a *PointApiService) PointGetByPath(ctx context.Context, path string, localVarOptionals map[string]interface{}) (Point, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     Point
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/points"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("path", parameterToString(path, ""))
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
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

/* PointApiService Retrieve multiple points by web id or path.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param optional (nil or map[string]interface{}) with one or more of:
    @param "asParallel" (bool) Specifies if the retrieval processes should be run in parallel on the server. This may improve the response time for large amounts of requested points. The default is &#39;false&#39;.
    @param "includeMode" (string) The include mode for the return list. The default is &#39;All&#39;.
    @param "path" ([]string) The path of a point. Multiple points may be specified with multiple instances of the parameter.
    @param "selectedFields" (string) List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned.
    @param "webId" ([]string) The ID of a point. Multiple points may be specified with multiple instances of the parameter.
    @param "webIdType" (string) Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;.
@return ItemsItemPoint*/
func (a *PointApiService) PointGetMultiple(ctx context.Context, localVarOptionals map[string]interface{}) (ItemsItemPoint, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     ItemsItemPoint
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/points/multiple"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["asParallel"], "bool", "asParallel"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["includeMode"], "string", "includeMode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["selectedFields"], "string", "selectedFields"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["webIdType"], "string", "webIdType"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["asParallel"].(bool); localVarOk {
		localVarQueryParams.Add("asParallel", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["includeMode"].(string); localVarOk {
		localVarQueryParams.Add("includeMode", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["path"].([]string); localVarOk {
		localVarQueryParams.Add("path", parameterToString(localVarTempParam, "multi"))
	}
	if localVarTempParam, localVarOk := localVarOptionals["selectedFields"].(string); localVarOk {
		localVarQueryParams.Add("selectedFields", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["webId"].([]string); localVarOk {
		localVarQueryParams.Add("webId", parameterToString(localVarTempParam, "multi"))
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

/* PointApiService Update a point.
* @param ctx context.Context for authentication, logging, tracing, etc.
@param webId The ID of the point.
@param pointDTO A partial point containing the desired changes.
@return */
func (a *PointApiService) PointUpdate(ctx context.Context, webId string, pointDTO Point) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Patch")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/points/{webId}"
	localVarPath = strings.Replace(localVarPath, "{"+"webId"+"}", fmt.Sprintf("%v", webId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = &pointDTO
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
