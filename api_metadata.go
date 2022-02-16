/*
Apicurio Registry API [v2]

Apicurio Registry is a datastore for standard event schemas and API designs. Apicurio Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.  The Apicurio Registry REST API enables client applications to manage the artifacts in the registry. This API provides create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata.   The supported artifact types include: - Apache Avro schema - AsyncAPI specification - Google protocol buffers - GraphQL schema - JSON Schema - Kafka Connect schema - OpenAPI specification - Web Services Description Language - XML Schema Definition   **Important**: The Apicurio Registry REST API is available from `https://MY-REGISTRY-URL/apis/registry/v2` by default. Therefore you must prefix all API operation paths with `../apis/registry/v2` in this case. For example: `../apis/registry/v2/ids/globalIds/{globalId}`. 

API version: 2.2.0.Final
Contact: apicurio@lists.jboss.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registryclient

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"os"
)

// Linger please
var (
	_ context.Context
)

// MetadataApiService MetadataApi service
type MetadataApiService service

type ApiDeleteArtifactVersionMetaDataRequest struct {
	ctx context.Context
	ApiService *MetadataApiService
	groupId string
	artifactId string
	version string
}


func (r ApiDeleteArtifactVersionMetaDataRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteArtifactVersionMetaDataExecute(r)
}

/*
DeleteArtifactVersionMetaData Delete artifact version metadata

Deletes the user-editable metadata properties of the artifact version.  Any properties
that are not user-editable are preserved.

This operation can fail for the following reasons:

* No artifact with this `artifactId` exists (HTTP error `404`)
* No version with this `version` exists (HTTP error `404`)
* A server error occurred (HTTP error `500`)


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.
 @param artifactId The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier.
 @param version The unique identifier of a specific version of the artifact content.
 @return ApiDeleteArtifactVersionMetaDataRequest
*/
func (a *MetadataApiService) DeleteArtifactVersionMetaData(ctx context.Context, groupId string, artifactId string, version string) ApiDeleteArtifactVersionMetaDataRequest {
	return ApiDeleteArtifactVersionMetaDataRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		artifactId: artifactId,
		version: version,
	}
}

// Execute executes the request
func (a *MetadataApiService) DeleteArtifactVersionMetaDataExecute(r ApiDeleteArtifactVersionMetaDataRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetadataApiService.DeleteArtifactVersionMetaData")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups/{groupId}/artifacts/{artifactId}/versions/{version}/meta"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"artifactId"+"}", url.PathEscape(parameterToString(r.artifactId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", url.PathEscape(parameterToString(r.version, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetArtifactMetaDataRequest struct {
	ctx context.Context
	ApiService *MetadataApiService
	groupId string
	artifactId string
}


func (r ApiGetArtifactMetaDataRequest) Execute() (*ArtifactMetaData, *http.Response, error) {
	return r.ApiService.GetArtifactMetaDataExecute(r)
}

/*
GetArtifactMetaData Get artifact metadata

Gets the metadata for an artifact in the registry.  The returned metadata includes
both generated (read-only) and editable metadata (such as name and description).

This operation can fail for the following reasons:

* No artifact with this `artifactId` exists (HTTP error `404`)
* A server error occurred (HTTP error `500`)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.
 @param artifactId The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier.
 @return ApiGetArtifactMetaDataRequest
*/
func (a *MetadataApiService) GetArtifactMetaData(ctx context.Context, groupId string, artifactId string) ApiGetArtifactMetaDataRequest {
	return ApiGetArtifactMetaDataRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		artifactId: artifactId,
	}
}

// Execute executes the request
//  @return ArtifactMetaData
func (a *MetadataApiService) GetArtifactMetaDataExecute(r ApiGetArtifactMetaDataRequest) (*ArtifactMetaData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ArtifactMetaData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetadataApiService.GetArtifactMetaData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups/{groupId}/artifacts/{artifactId}/meta"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"artifactId"+"}", url.PathEscape(parameterToString(r.artifactId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetArtifactVersionMetaDataRequest struct {
	ctx context.Context
	ApiService *MetadataApiService
	groupId string
	artifactId string
	version string
}


func (r ApiGetArtifactVersionMetaDataRequest) Execute() (*VersionMetaData, *http.Response, error) {
	return r.ApiService.GetArtifactVersionMetaDataExecute(r)
}

/*
GetArtifactVersionMetaData Get artifact version metadata

Retrieves the metadata for a single version of the artifact.  The version metadata is 
a subset of the artifact metadata and only includes the metadata that is specific to
the version (for example, this doesn't include `modifiedOn`).

This operation can fail for the following reasons:

* No artifact with this `artifactId` exists (HTTP error `404`)
* No version with this `version` exists (HTTP error `404`)
* A server error occurred (HTTP error `500`)


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.
 @param artifactId The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier.
 @param version The unique identifier of a specific version of the artifact content.
 @return ApiGetArtifactVersionMetaDataRequest
*/
func (a *MetadataApiService) GetArtifactVersionMetaData(ctx context.Context, groupId string, artifactId string, version string) ApiGetArtifactVersionMetaDataRequest {
	return ApiGetArtifactVersionMetaDataRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		artifactId: artifactId,
		version: version,
	}
}

// Execute executes the request
//  @return VersionMetaData
func (a *MetadataApiService) GetArtifactVersionMetaDataExecute(r ApiGetArtifactVersionMetaDataRequest) (*VersionMetaData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VersionMetaData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetadataApiService.GetArtifactVersionMetaData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups/{groupId}/artifacts/{artifactId}/versions/{version}/meta"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"artifactId"+"}", url.PathEscape(parameterToString(r.artifactId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", url.PathEscape(parameterToString(r.version, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetArtifactVersionMetaDataByContentRequest struct {
	ctx context.Context
	ApiService *MetadataApiService
	groupId string
	artifactId string
	body **os.File
	canonical *bool
}

// The content of an artifact version.
func (r ApiGetArtifactVersionMetaDataByContentRequest) Body(body *os.File) ApiGetArtifactVersionMetaDataByContentRequest {
	r.body = &body
	return r
}
// Parameter that can be set to &#x60;true&#x60; to indicate that the server should \&quot;canonicalize\&quot; the content when searching for a matching version.  Canonicalization is unique to each artifact type, but typically involves removing any extra whitespace and formatting the content in a consistent manner.
func (r ApiGetArtifactVersionMetaDataByContentRequest) Canonical(canonical bool) ApiGetArtifactVersionMetaDataByContentRequest {
	r.canonical = &canonical
	return r
}

func (r ApiGetArtifactVersionMetaDataByContentRequest) Execute() (*VersionMetaData, *http.Response, error) {
	return r.ApiService.GetArtifactVersionMetaDataByContentExecute(r)
}

/*
GetArtifactVersionMetaDataByContent Get artifact version metadata by content

Gets the metadata for an artifact that matches the raw content.  Searches the registry
for a version of the given artifact matching the content provided in the body of the
POST.

This operation can fail for the following reasons:

* Provided content (request body) was empty (HTTP error `400`)
* No artifact with the `artifactId` exists (HTTP error `404`)
* No artifact version matching the provided content exists (HTTP error `404`)
* A server error occurred (HTTP error `500`)


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.
 @param artifactId The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier.
 @return ApiGetArtifactVersionMetaDataByContentRequest
*/
func (a *MetadataApiService) GetArtifactVersionMetaDataByContent(ctx context.Context, groupId string, artifactId string) ApiGetArtifactVersionMetaDataByContentRequest {
	return ApiGetArtifactVersionMetaDataByContentRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		artifactId: artifactId,
	}
}

// Execute executes the request
//  @return VersionMetaData
func (a *MetadataApiService) GetArtifactVersionMetaDataByContentExecute(r ApiGetArtifactVersionMetaDataByContentRequest) (*VersionMetaData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VersionMetaData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetadataApiService.GetArtifactVersionMetaDataByContent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups/{groupId}/artifacts/{artifactId}/meta"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"artifactId"+"}", url.PathEscape(parameterToString(r.artifactId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	if r.canonical != nil {
		localVarQueryParams.Add("canonical", parameterToString(*r.canonical, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateArtifactMetaDataRequest struct {
	ctx context.Context
	ApiService *MetadataApiService
	groupId string
	artifactId string
	editableMetaData *EditableMetaData
}

// Updated artifact metadata.
func (r ApiUpdateArtifactMetaDataRequest) EditableMetaData(editableMetaData EditableMetaData) ApiUpdateArtifactMetaDataRequest {
	r.editableMetaData = &editableMetaData
	return r
}

func (r ApiUpdateArtifactMetaDataRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateArtifactMetaDataExecute(r)
}

/*
UpdateArtifactMetaData Update artifact metadata

Updates the editable parts of the artifact's metadata.  Not all metadata fields can
be updated.  For example, `createdOn` and `createdBy` are both read-only properties.

This operation can fail for the following reasons:

* No artifact with the `artifactId` exists (HTTP error `404`)
* A server error occurred (HTTP error `500`)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.
 @param artifactId The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier.
 @return ApiUpdateArtifactMetaDataRequest
*/
func (a *MetadataApiService) UpdateArtifactMetaData(ctx context.Context, groupId string, artifactId string) ApiUpdateArtifactMetaDataRequest {
	return ApiUpdateArtifactMetaDataRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		artifactId: artifactId,
	}
}

// Execute executes the request
func (a *MetadataApiService) UpdateArtifactMetaDataExecute(r ApiUpdateArtifactMetaDataRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetadataApiService.UpdateArtifactMetaData")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups/{groupId}/artifacts/{artifactId}/meta"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"artifactId"+"}", url.PathEscape(parameterToString(r.artifactId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.editableMetaData == nil {
		return nil, reportError("editableMetaData is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.editableMetaData
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUpdateArtifactVersionMetaDataRequest struct {
	ctx context.Context
	ApiService *MetadataApiService
	groupId string
	artifactId string
	version string
	editableMetaData *EditableMetaData
}

func (r ApiUpdateArtifactVersionMetaDataRequest) EditableMetaData(editableMetaData EditableMetaData) ApiUpdateArtifactVersionMetaDataRequest {
	r.editableMetaData = &editableMetaData
	return r
}

func (r ApiUpdateArtifactVersionMetaDataRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateArtifactVersionMetaDataExecute(r)
}

/*
UpdateArtifactVersionMetaData Update artifact version metadata

Updates the user-editable portion of the artifact version's metadata.  Only some of 
the metadata fields are editable by the user.  For example, `description` is editable, 
but `createdOn` is not.

This operation can fail for the following reasons:

* No artifact with this `artifactId` exists (HTTP error `404`)
* No version with this `version` exists (HTTP error `404`)
* A server error occurred (HTTP error `500`)


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.
 @param artifactId The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier.
 @param version The unique identifier of a specific version of the artifact content.
 @return ApiUpdateArtifactVersionMetaDataRequest
*/
func (a *MetadataApiService) UpdateArtifactVersionMetaData(ctx context.Context, groupId string, artifactId string, version string) ApiUpdateArtifactVersionMetaDataRequest {
	return ApiUpdateArtifactVersionMetaDataRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		artifactId: artifactId,
		version: version,
	}
}

// Execute executes the request
func (a *MetadataApiService) UpdateArtifactVersionMetaDataExecute(r ApiUpdateArtifactVersionMetaDataRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetadataApiService.UpdateArtifactVersionMetaData")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/groups/{groupId}/artifacts/{artifactId}/versions/{version}/meta"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"artifactId"+"}", url.PathEscape(parameterToString(r.artifactId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"version"+"}", url.PathEscape(parameterToString(r.version, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.editableMetaData == nil {
		return nil, reportError("editableMetaData is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.editableMetaData
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
