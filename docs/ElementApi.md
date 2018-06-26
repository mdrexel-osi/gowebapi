# \ElementApi

All URIs are relative to *https://localhost/piwebapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ElementAddReferencedElement**](ElementApi.md#ElementAddReferencedElement) | **Post** /elements/{webId}/referencedelements | Add a reference to an existing element to the child elements collection.
[**ElementCreateAnalysis**](ElementApi.md#ElementCreateAnalysis) | **Post** /elements/{webId}/analyses | Create an Analysis.
[**ElementCreateAttribute**](ElementApi.md#ElementCreateAttribute) | **Post** /elements/{webId}/attributes | Create a new attribute of the specified element.
[**ElementCreateConfig**](ElementApi.md#ElementCreateConfig) | **Post** /elements/{webId}/config | Executes the create configuration function of the data references found within the attributes of the element, and optionally, its children.
[**ElementCreateElement**](ElementApi.md#ElementCreateElement) | **Post** /elements/{webId}/elements | Create a child element.
[**ElementCreateSearchByAttribute**](ElementApi.md#ElementCreateSearchByAttribute) | **Post** /elements/searchbyattribute | Create a link for a \&quot;Search Elements By Attribute Value\&quot; operation, whose queries are specified in the request content. The SearchRoot is specified by the Web Id of the root Element. If the SearchRoot is not specified, then the search starts at the Asset Database. ElementTemplate must be provided as the Web ID of the ElementTemplate, which are used to create the Elements. All the attributes in the queries must be defined as AttributeTemplates on the ElementTemplate. An array of attribute value queries are ANDed together to find the desired Element objects. At least one value query must be specified. There are limitations on SearchOperators.
[**ElementCreateSecurityEntry**](ElementApi.md#ElementCreateSecurityEntry) | **Post** /elements/{webId}/securityentries | Create a security entry owned by the element.
[**ElementDelete**](ElementApi.md#ElementDelete) | **Delete** /elements/{webId} | Delete an element.
[**ElementDeleteSecurityEntry**](ElementApi.md#ElementDeleteSecurityEntry) | **Delete** /elements/{webId}/securityentries/{name} | Delete a security entry owned by the element.
[**ElementExecuteSearchByAttribute**](ElementApi.md#ElementExecuteSearchByAttribute) | **Get** /elements/searchbyattribute/{searchId} | Execute a \&quot;Search Elements By Attribute Value\&quot; operation.
[**ElementFindElementAttributes**](ElementApi.md#ElementFindElementAttributes) | **Get** /elements/{webId}/elementattributes | Retrieves a list of element attributes matching the specified filters from the specified element.
[**ElementGet**](ElementApi.md#ElementGet) | **Get** /elements/{webId} | Retrieve an element.
[**ElementGetAnalyses**](ElementApi.md#ElementGetAnalyses) | **Get** /elements/{webId}/analyses | Retrieve analyses based on the specified conditions.
[**ElementGetAttributes**](ElementApi.md#ElementGetAttributes) | **Get** /elements/{webId}/attributes | Get the attributes of the specified element.
[**ElementGetByPath**](ElementApi.md#ElementGetByPath) | **Get** /elements | Retrieve an element by path.
[**ElementGetCategories**](ElementApi.md#ElementGetCategories) | **Get** /elements/{webId}/categories | Get an element&#39;s categories.
[**ElementGetElements**](ElementApi.md#ElementGetElements) | **Get** /elements/{webId}/elements | Retrieve elements based on the specified conditions. By default, this method selects immediate children of the specified element.
[**ElementGetElementsQuery**](ElementApi.md#ElementGetElementsQuery) | **Get** /elements/search | Retrieve elements based on the specified conditions. By default, returns all the elements.
[**ElementGetEventFrames**](ElementApi.md#ElementGetEventFrames) | **Get** /elements/{webId}/eventframes | Retrieve event frames that reference this element based on the specified conditions. By default, returns all event frames that reference this element that have been active in the past 8 hours.
[**ElementGetMultiple**](ElementApi.md#ElementGetMultiple) | **Get** /elements/multiple | Retrieve multiple elements by web id or path.
[**ElementGetNotificationRules**](ElementApi.md#ElementGetNotificationRules) | **Get** /elements/{webId}/notificationrules | Retrieve notification rules for an element
[**ElementGetPaths**](ElementApi.md#ElementGetPaths) | **Get** /elements/{webId}/paths | Get a list of the full or relative paths to this element.
[**ElementGetReferencedElements**](ElementApi.md#ElementGetReferencedElements) | **Get** /elements/{webId}/referencedelements | Retrieve referenced elements based on the specified conditions. By default, this method selects all referenced elements of the current resource.
[**ElementGetSecurity**](ElementApi.md#ElementGetSecurity) | **Get** /elements/{webId}/security | Get the security information of the specified security item associated with the element for a specified user.
[**ElementGetSecurityEntries**](ElementApi.md#ElementGetSecurityEntries) | **Get** /elements/{webId}/securityentries | Retrieve the security entries associated with the element based on the specified criteria. By default, all security entries for this element are returned.
[**ElementGetSecurityEntryByName**](ElementApi.md#ElementGetSecurityEntryByName) | **Get** /elements/{webId}/securityentries/{name} | Retrieve the security entry associated with the element with the specified name.
[**ElementRemoveReferencedElement**](ElementApi.md#ElementRemoveReferencedElement) | **Delete** /elements/{webId}/referencedelements | Remove a reference to an existing element from the child elements collection.
[**ElementUpdate**](ElementApi.md#ElementUpdate) | **Patch** /elements/{webId} | Update an element by replacing items in its definition.
[**ElementUpdateSecurityEntry**](ElementApi.md#ElementUpdateSecurityEntry) | **Put** /elements/{webId}/securityentries/{name} | Update a security entry owned by the element.


# **ElementAddReferencedElement**
> ElementAddReferencedElement(ctx, webId, referencedElementWebId, optional)
Add a reference to an existing element to the child elements collection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the element which the referenced element will be added to. | 
  **referencedElementWebId** | [**[]string**](string.md)| The ID of the referenced element. Multiple referenced elements may be specified with multiple instances of the parameter. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the element which the referenced element will be added to. | 
 **referencedElementWebId** | [**[]string**](string.md)| The ID of the referenced element. Multiple referenced elements may be specified with multiple instances of the parameter. | 
 **referenceType** | **string**| The name of the reference type between the parent and the referenced element. The default is \&quot;parent-child\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementCreateAnalysis**
> ElementCreateAnalysis(ctx, webId, analysis, optional)
Create an Analysis.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the element on which to create the Analysis. | 
  **analysis** | [**Analysis**](Analysis.md)| The new Analysis definition. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the element on which to create the Analysis. | 
 **analysis** | [**Analysis**](Analysis.md)| The new Analysis definition. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementCreateAttribute**
> ElementCreateAttribute(ctx, webId, attribute, optional)
Create a new attribute of the specified element.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the element on which to create the attribute. | 
  **attribute** | [**Attribute**](Attribute.md)| The definition of the new attribute. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the element on which to create the attribute. | 
 **attribute** | [**Attribute**](Attribute.md)| The definition of the new attribute. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementCreateConfig**
> ElementCreateConfig(ctx, webId, optional)
Executes the create configuration function of the data references found within the attributes of the element, and optionally, its children.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the element. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the element. | 
 **includeChildElements** | **bool**| If true, includes the child elements of the specified element. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementCreateElement**
> ElementCreateElement(ctx, webId, element, optional)
Create a child element.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the parent element on which to create the element. | 
  **element** | [**Element**](Element.md)| The new element definition. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the parent element on which to create the element. | 
 **element** | [**Element**](Element.md)| The new element definition. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementCreateSearchByAttribute**
> ItemsElement ElementCreateSearchByAttribute(ctx, query, optional)
Create a link for a \"Search Elements By Attribute Value\" operation, whose queries are specified in the request content. The SearchRoot is specified by the Web Id of the root Element. If the SearchRoot is not specified, then the search starts at the Asset Database. ElementTemplate must be provided as the Web ID of the ElementTemplate, which are used to create the Elements. All the attributes in the queries must be defined as AttributeTemplates on the ElementTemplate. An array of attribute value queries are ANDed together to find the desired Element objects. At least one value query must be specified. There are limitations on SearchOperators.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **query** | [**SearchByAttribute**](SearchByAttribute.md)| The query of search by attribute. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | [**SearchByAttribute**](SearchByAttribute.md)| The query of search by attribute. | 
 **noResults** | **bool**| If false, the response content will contain the first page of the search results. If true, the response content will be empty. The default is false. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsElement**](Items[Element].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementCreateSecurityEntry**
> ElementCreateSecurityEntry(ctx, webId, securityEntry, optional)
Create a security entry owned by the element.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the element where the security entry will be created. | 
  **securityEntry** | [**SecurityEntry**](SecurityEntry.md)| The new security entry definition. The full list of allow and deny rights must be supplied. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the element where the security entry will be created. | 
 **securityEntry** | [**SecurityEntry**](SecurityEntry.md)| The new security entry definition. The full list of allow and deny rights must be supplied. | 
 **applyToChildren** | **bool**| If false, the new access permissions are only applied to the associated object. If true, the access permissions of children with any parent-child reference types will change when the permissions on the primary parent change. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementDelete**
> ElementDelete(ctx, webId)
Delete an element.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the element. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementDeleteSecurityEntry**
> ElementDeleteSecurityEntry(ctx, name, webId, optional)
Delete a security entry owned by the element.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The name of the security entry. For every backslash character (\\) in the security entry name, replace with asterisk (*). As an example, use domain*username instead of domain\\username. | 
  **webId** | **string**| The ID of the element where the security entry will be deleted. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The name of the security entry. For every backslash character (\\) in the security entry name, replace with asterisk (*). As an example, use domain*username instead of domain\\username. | 
 **webId** | **string**| The ID of the element where the security entry will be deleted. | 
 **applyToChildren** | **bool**| If false, the new access permissions are only applied to the associated object. If true, the access permissions of children with any parent-child reference types will change when the permissions on the primary parent change. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementExecuteSearchByAttribute**
> ItemsElement ElementExecuteSearchByAttribute(ctx, searchId, optional)
Execute a \"Search Elements By Attribute Value\" operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **searchId** | **string**| The encoded search Id of the \&quot;Search Elements By Attribute Value\&quot; operation. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchId** | **string**| The encoded search Id of the \&quot;Search Elements By Attribute Value\&quot; operation. | 
 **categoryName** | **string**| Specify that the owner of the returned attributes must have this category. For Asset Servers older than 2.7, a 400 status code (Bad Request) will be returned if this parameter is specified. The default is no filter. | 
 **descriptionFilter** | **string**| The element description filter string used for finding objects. Only the first 440 characters of the description will be searched. For Asset Servers older than 2.7, a 400 status code (Bad Request) will be returned if this parameter is specified. The default is no filter. | 
 **maxCount** | **int32**| The maximum number of objects to be returned. The default is 1000. | 
 **nameFilter** | **string**| The name query string used for finding objects. The default is no filter. | 
 **searchFullHierarchy** | **bool**| Specifies if the search should include objects nested further than the immediate children of the searchRoot. The default is &#39;false&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **sortField** | **string**| The field or property of the object used to sort the returned collection. The default is &#39;Name&#39;. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39;. | 
 **startIndex** | **int32**| The starting index (zero based) of the items to be returned. The default is 0. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsElement**](Items[Element].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementFindElementAttributes**
> ItemsAttribute ElementFindElementAttributes(ctx, webId, optional)
Retrieves a list of element attributes matching the specified filters from the specified element.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the element to use as the root of the search. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the element to use as the root of the search. | 
 **attributeCategory** | **string**| Specify that returned attributes must have this category. The default is no filter. | 
 **attributeDescriptionFilter** | **string**| The attribute description filter string used for finding objects. Only the first 440 characters of the description will be searched. For Asset Servers older than 2.7, a 400 status code (Bad Request) will be returned if this parameter is specified. The default is no filter. | 
 **attributeNameFilter** | **string**| The attribute name filter string used for finding objects. The default is no filter. | 
 **attributeType** | **string**| Specify that returned attributes&#39; value type must be this value type. The default is no filter. | 
 **elementCategory** | **string**| Specify that the owner of the returned attributes must have this category. The default is no filter. | 
 **elementDescriptionFilter** | **string**| The element description filter string used for finding objects. Only the first 440 characters of the description will be searched. For Asset Servers older than 2.7, a 400 status code (Bad Request) will be returned if this parameter is specified. The default is no filter. | 
 **elementNameFilter** | **string**| The element name filter string used for finding objects. The default is no filter. | 
 **elementTemplate** | **string**| Specify that the owner of the returned attributes must have this template or a template derived from this template. The default is no filter. | 
 **elementType** | **string**| Specify that the element of the returned attributes must have this AFElementType. The default is no filter. | 
 **maxCount** | **int32**| The maximum number of objects to be returned (the page size). The default is 1000. | 
 **searchFullHierarchy** | **bool**| Specifies if the search should include objects nested further than immediate children of the given resource. The default is &#39;false&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **sortField** | **string**| The field or property of the object used to sort the returned collection. The default is &#39;Name&#39;. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39;. | 
 **startIndex** | **int32**| The starting index (zero based) of the items to be returned. The default is 0. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsAttribute**](Items[Attribute].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementGet**
> Element ElementGet(ctx, webId, optional)
Retrieve an element.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the element. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the element. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**Element**](Element.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementGetAnalyses**
> ItemsAnalysis ElementGetAnalyses(ctx, webId, optional)
Retrieve analyses based on the specified conditions.

Users can search for the analyses based on specific search parameters. If no parameters are specified in the search, the default values for each parameter will be used and will return the analyses that match the default search.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the element, which is the Target of the analyses. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the element, which is the Target of the analyses. | 
 **maxCount** | **int32**| The maximum number of objects to be returned per call (page size). The default is 1000. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **sortField** | **string**| The field or property of the object used to sort the returned collection. The default is &#39;Name&#39;. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39;. | 
 **startIndex** | **int32**| The starting index (zero based) of the items to be returned. The default is 0. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsAnalysis**](Items[Analysis].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementGetAttributes**
> ItemsAttribute ElementGetAttributes(ctx, webId, optional)
Get the attributes of the specified element.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the element. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the element. | 
 **categoryName** | **string**| Specify that returned attributes must have this category. The default is no category filter. | 
 **maxCount** | **int32**| The maximum number of objects to be returned per call (page size). The default is 1000. | 
 **nameFilter** | **string**| The name query string used for finding attributes. The default is no filter. | 
 **searchFullHierarchy** | **bool**| Specifies if the search should include attributes nested further than the immediate attributes of the searchRoot. The default is &#39;false&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **showExcluded** | **bool**| Specified if the search should include attributes with the Excluded property set. The default is &#39;false&#39;. | 
 **showHidden** | **bool**| Specified if the search should include attributes with the Hidden property set. The default is &#39;false&#39;. | 
 **sortField** | **string**| The field or property of the object used to sort the returned collection. The default is &#39;Name&#39;. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39;. | 
 **startIndex** | **int32**| The starting index (zero based) of the items to be returned. The default is 0. | 
 **templateName** | **string**| Specify that returned attributes must be members of this template. The default is no template filter. | 
 **trait** | [**[]string**](string.md)| The name of the attribute trait. Multiple traits may be specified with multiple instances of the parameter. | 
 **traitCategory** | [**[]string**](string.md)| The category of the attribute traits. Multiple categories may be specified with multiple instances of the parameter. If the parameter is not specified, or if its value is \&quot;all\&quot;, then all attribute traits of all categories will be returned. | 
 **valueType** | **string**| Specify that returned attributes&#39; value type must be the given value type. The default is no value type filter. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsAttribute**](Items[Attribute].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementGetByPath**
> Element ElementGetByPath(ctx, path, optional)
Retrieve an element by path.

This method returns an element based on the hierarchical path associated with it, and should be used when a path has been received from a separate part of the PI System for use in the PI Web API. Users should primarily search with the WebID when available.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **path** | **string**| The path to the element. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string**| The path to the element. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**Element**](Element.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementGetCategories**
> ItemsElementCategory ElementGetCategories(ctx, webId, optional)
Get an element's categories.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the element. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the element. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsElementCategory**](Items[ElementCategory].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementGetElements**
> ItemsElement ElementGetElements(ctx, webId, optional)
Retrieve elements based on the specified conditions. By default, this method selects immediate children of the specified element.

Users can search for the elements based on specific search parameters. If no parameters are specified in the search, the default values for each parameter will be used and will return the elements that match the default search.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the element to use as the root of the search. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the element to use as the root of the search. | 
 **categoryName** | **string**| Specify that returned elements must have this category. The default is no category filter. | 
 **descriptionFilter** | **string**| Specify that returned elements must have this description. The default is no description filter. | 
 **elementType** | **string**| Specify that returned elements must have this type. The default type is &#39;Any&#39;. | 
 **maxCount** | **int32**| The maximum number of objects to be returned per call (page size). The default is 1000. | 
 **nameFilter** | **string**| The name query string used for finding objects. The default is no filter. | 
 **searchFullHierarchy** | **bool**| Specifies if the search should include objects nested further than the immediate children of the searchRoot. The default is &#39;false&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **sortField** | **string**| The field or property of the object used to sort the returned collection. The default is &#39;Name&#39;. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39;. | 
 **startIndex** | **int32**| The starting index (zero based) of the items to be returned. The default is 0. | 
 **templateName** | **string**| Specify that returned elements must have this template or a template derived from this template. The default is no template filter. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsElement**](Items[Element].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementGetElementsQuery**
> ItemsElement ElementGetElementsQuery(ctx, optional)
Retrieve elements based on the specified conditions. By default, returns all the elements.

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
 **query** | **string**| The query string is a list of filters used to perform an AFSearch for the elements in the asset database. An example would be: \&quot;query&#x3D;Name:&#x3D;MyElement* Template:&#x3D;ElementTemplate*\&quot;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **startIndex** | **int32**| The starting index (zero based) of the items to be returned. The default is 0. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsElement**](Items[Element].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementGetEventFrames**
> ItemsEventFrame ElementGetEventFrames(ctx, webId, optional)
Retrieve event frames that reference this element based on the specified conditions. By default, returns all event frames that reference this element that have been active in the past 8 hours.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the element whose related event frames are sought. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the element whose related event frames are sought. | 
 **canBeAcknowledged** | **bool**| Specify the returned event frames&#39; canBeAcknowledged property. The default is no canBeAcknowledged filter. | 
 **categoryName** | **string**| Specify that returned event frames must have this category. The default is no category filter. | 
 **endTime** | **string**| The ending time for the search. The endTime must be greater than or equal to the startTime. The searchMode parameter will control whether the comparison will be performed against the event frame&#39;s startTime or endTime. The default is &#39;*&#39; if searchMode is not one of the &#39;Backward*&#39; or &#39;Forward*&#39; values. | 
 **isAcknowledged** | **bool**| Specify the returned event frames&#39; isAcknowledged property. The default no isAcknowledged filter. | 
 **maxCount** | **int32**| The maximum number of objects to be returned per call (page size). The default is 1000. | 
 **nameFilter** | **string**| The name query string used for finding event frames. The default is no filter. | 
 **searchMode** | **string**| Determines how the startTime and endTime parameters are treated when searching for event frame objects to be included in the returned collection. If this parameter is one of the &#39;Backward*&#39; or &#39;Forward*&#39; values, none of endTime, sortField, or sortOrder may be specified. The default is &#39;Overlapped&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **severity** | [**[]string**](string.md)| Specify that returned event frames must have this severity. Multiple severity values may be specified with multiple instances of the parameter. The default is no severity filter. | 
 **sortField** | **string**| The field or property of the object used to sort the returned collection. The default is &#39;Name&#39; if searchMode is not one of the &#39;Backward*&#39; or &#39;Forward*&#39; values. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39; if searchMode is not one of the &#39;Backward*&#39; or &#39;Forward*&#39; values. | 
 **startIndex** | **int32**| The starting index (zero based) of the items to be returned. The default is 0. | 
 **startTime** | **string**| The starting time for the search. startTime must be less than or equal to the endTime. The searchMode parameter will control whether the comparison will be performed against the event frame&#39;s startTime or endTime. The default is &#39;*-8h&#39;. | 
 **templateName** | **string**| Specify that returned event frames must have this template or a template derived from this template. The default is no template filter. Specify this parameter by name. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsEventFrame**](Items[EventFrame].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementGetMultiple**
> ItemsItemElement ElementGetMultiple(ctx, optional)
Retrieve multiple elements by web id or path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asParallel** | **bool**| Specifies if the retrieval processes should be run in parallel on the server. This may improve the response time for large amounts of requested attributes. The default is &#39;false&#39;. | 
 **includeMode** | **string**| The include mode for the return list. The default is &#39;All&#39;. | 
 **path** | [**[]string**](string.md)| The path of an element. Multiple elements may be specified with multiple instances of the parameter. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webId** | [**[]string**](string.md)| The ID of an element. Multiple elements may be specified with multiple instances of the parameter. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsItemElement**](Items[Item[Element]].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementGetNotificationRules**
> ItemsNotificationRule ElementGetNotificationRules(ctx, webId, optional)
Retrieve notification rules for an element

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

[**ItemsNotificationRule**](Items[NotificationRule].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementGetPaths**
> ItemsString ElementGetPaths(ctx, webId, optional)
Get a list of the full or relative paths to this element.

This method will return paths with the primary path at the first index. If there is no primary path, then null will be at the first index. If relative path is specified but does not exist, null will be returned at the first index.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the element. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the element. | 
 **relativePath** | **string**| The full path in ShortName format to the parent object that the returned paths should be relative. For example, \&quot;\\\\Server1\\Database2\&quot; would return all the paths to the element relative to the database. A path of \&quot;\\\\Server1\\Database2\\RootElement\&quot; would return all paths to the element relative to \&quot;RootElement\&quot;. If null, then all the full paths to the element will be returned. | 

### Return type

[**ItemsString**](Items[string].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementGetReferencedElements**
> ItemsElement ElementGetReferencedElements(ctx, webId, optional)
Retrieve referenced elements based on the specified conditions. By default, this method selects all referenced elements of the current resource.

Users can search for the referenced elements based on specific search parameters. If no parameters are specified in the search, the default values for each parameter will be used and will return the elements that match the default search.

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
 **categoryName** | **string**| Specify that returned elements must have this category. The default is no category filter. | 
 **descriptionFilter** | **string**| Specify that returned elements must have this description. The default is no description filter. | 
 **elementType** | **string**| Specify that returned elements must have this type. The default type is &#39;Any&#39;. | 
 **maxCount** | **int32**| The maximum number of objects to be returned per call (page size). The default is 1000. | 
 **nameFilter** | **string**| The name query string used for finding objects. The default is no filter. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **sortField** | **string**| The field or property of the object used to sort the returned collection. The default is &#39;Name&#39;. | 
 **sortOrder** | **string**| The order that the returned collection is sorted. The default is &#39;Ascending&#39;. | 
 **startIndex** | **int32**| The starting index (zero based) of the items to be returned. The default is 0. | 
 **templateName** | **string**| Specify that returned elements must have this template or a template derived from this template. The default is no template filter. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsElement**](Items[Element].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementGetSecurity**
> ItemsSecurityRights ElementGetSecurity(ctx, webId, userIdentity, optional)
Get the security information of the specified security item associated with the element for a specified user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the element for the security to be checked. | 
  **userIdentity** | [**[]string**](string.md)| The user identity for the security information to be checked. Multiple security identities may be specified with multiple instances of the parameter. If the parameter is not specified, only the current user&#39;s security rights will be returned. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the element for the security to be checked. | 
 **userIdentity** | [**[]string**](string.md)| The user identity for the security information to be checked. Multiple security identities may be specified with multiple instances of the parameter. If the parameter is not specified, only the current user&#39;s security rights will be returned. | 
 **forceRefresh** | **bool**| Indicates if the security cache should be refreshed before getting security information. The default is &#39;false&#39;. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsSecurityRights**](Items[SecurityRights].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementGetSecurityEntries**
> ItemsSecurityEntry ElementGetSecurityEntries(ctx, webId, optional)
Retrieve the security entries associated with the element based on the specified criteria. By default, all security entries for this element are returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the element. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webId** | **string**| The ID of the element. | 
 **nameFilter** | **string**| The name query string used for filtering security entries. The default is no filter. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**ItemsSecurityEntry**](Items[SecurityEntry].md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementGetSecurityEntryByName**
> SecurityEntry ElementGetSecurityEntryByName(ctx, name, webId, optional)
Retrieve the security entry associated with the element with the specified name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The name of the security entry. For every backslash character (\\) in the security entry name, replace with asterisk (*). As an example, use domain*username instead of domain\\username. | 
  **webId** | **string**| The ID of the element. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The name of the security entry. For every backslash character (\\) in the security entry name, replace with asterisk (*). As an example, use domain*username instead of domain\\username. | 
 **webId** | **string**| The ID of the element. | 
 **selectedFields** | **string**| List of fields to be returned in the response, separated by semicolons (;). If this parameter is not specified, all available fields will be returned. | 
 **webIdType** | **string**| Optional parameter. Used to specify the type of WebID. Useful for URL brevity and other special cases. Default is the value of the configuration item \&quot;WebIDType\&quot;. | 

### Return type

[**SecurityEntry**](SecurityEntry.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementRemoveReferencedElement**
> ElementRemoveReferencedElement(ctx, webId, referencedElementWebId)
Remove a reference to an existing element from the child elements collection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the element which the referenced element will be removed from. | 
  **referencedElementWebId** | [**[]string**](string.md)| The ID of the referenced element. Multiple referenced elements may be specified with multiple instances of the parameter. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementUpdate**
> ElementUpdate(ctx, webId, element)
Update an element by replacing items in its definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **webId** | **string**| The ID of the element. | 
  **element** | [**Element**](Element.md)| A partial element containing the desired changes. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementUpdateSecurityEntry**
> ElementUpdateSecurityEntry(ctx, name, webId, securityEntry, optional)
Update a security entry owned by the element.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The name of the security entry. | 
  **webId** | **string**| The ID of the element where the security entry will be updated. | 
  **securityEntry** | [**SecurityEntry**](SecurityEntry.md)| The new security entry definition. The full list of allow and deny rights must be supplied or they will be removed. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| The name of the security entry. | 
 **webId** | **string**| The ID of the element where the security entry will be updated. | 
 **securityEntry** | [**SecurityEntry**](SecurityEntry.md)| The new security entry definition. The full list of allow and deny rights must be supplied or they will be removed. | 
 **applyToChildren** | **bool**| If false, the new access permissions are only applied to the associated object. If true, the access permissions of children with any parent-child reference types will change when the permissions on the primary parent change. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, text/html, application/x-ms-application

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

