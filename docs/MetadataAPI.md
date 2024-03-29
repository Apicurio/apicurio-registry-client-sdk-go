# \MetadataAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteArtifactVersionMetaData**](MetadataAPI.md#DeleteArtifactVersionMetaData) | **Delete** /groups/{groupId}/artifacts/{artifactId}/versions/{version}/meta | Delete artifact version metadata
[**GetArtifactMetaData**](MetadataAPI.md#GetArtifactMetaData) | **Get** /groups/{groupId}/artifacts/{artifactId}/meta | Get artifact metadata
[**GetArtifactOwner**](MetadataAPI.md#GetArtifactOwner) | **Get** /groups/{groupId}/artifacts/{artifactId}/owner | Get artifact owner
[**GetArtifactVersionMetaData**](MetadataAPI.md#GetArtifactVersionMetaData) | **Get** /groups/{groupId}/artifacts/{artifactId}/versions/{version}/meta | Get artifact version metadata
[**GetArtifactVersionMetaDataByContent**](MetadataAPI.md#GetArtifactVersionMetaDataByContent) | **Post** /groups/{groupId}/artifacts/{artifactId}/meta | Get artifact version metadata by content
[**UpdateArtifactMetaData**](MetadataAPI.md#UpdateArtifactMetaData) | **Put** /groups/{groupId}/artifacts/{artifactId}/meta | Update artifact metadata
[**UpdateArtifactOwner**](MetadataAPI.md#UpdateArtifactOwner) | **Put** /groups/{groupId}/artifacts/{artifactId}/owner | Update artifact owner
[**UpdateArtifactVersionMetaData**](MetadataAPI.md#UpdateArtifactVersionMetaData) | **Put** /groups/{groupId}/artifacts/{artifactId}/versions/{version}/meta | Update artifact version metadata



## DeleteArtifactVersionMetaData

> DeleteArtifactVersionMetaData(ctx, groupId, artifactId, version).Execute()

Delete artifact version metadata



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Apicurio/apicurio-registry-client-sdk-go"
)

func main() {
    groupId := "groupId_example" // string | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.
    artifactId := "artifactId_example" // string | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier.
    version := "version_example" // string | The unique identifier of a specific version of the artifact content.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MetadataAPI.DeleteArtifactVersionMetaData(context.Background(), groupId, artifactId, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.DeleteArtifactVersionMetaData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. | 
**artifactId** | **string** | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier. | 
**version** | **string** | The unique identifier of a specific version of the artifact content. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArtifactVersionMetaDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifactMetaData

> ArtifactMetaData GetArtifactMetaData(ctx, groupId, artifactId).Execute()

Get artifact metadata



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Apicurio/apicurio-registry-client-sdk-go"
)

func main() {
    groupId := "groupId_example" // string | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.
    artifactId := "artifactId_example" // string | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataAPI.GetArtifactMetaData(context.Background(), groupId, artifactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.GetArtifactMetaData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArtifactMetaData`: ArtifactMetaData
    fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.GetArtifactMetaData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. | 
**artifactId** | **string** | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtifactMetaDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ArtifactMetaData**](ArtifactMetaData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifactOwner

> ArtifactOwner GetArtifactOwner(ctx, groupId, artifactId).Execute()

Get artifact owner



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Apicurio/apicurio-registry-client-sdk-go"
)

func main() {
    groupId := "groupId_example" // string | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.
    artifactId := "artifactId_example" // string | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataAPI.GetArtifactOwner(context.Background(), groupId, artifactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.GetArtifactOwner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArtifactOwner`: ArtifactOwner
    fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.GetArtifactOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. | 
**artifactId** | **string** | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtifactOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ArtifactOwner**](ArtifactOwner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifactVersionMetaData

> VersionMetaData GetArtifactVersionMetaData(ctx, groupId, artifactId, version).Execute()

Get artifact version metadata



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Apicurio/apicurio-registry-client-sdk-go"
)

func main() {
    groupId := "groupId_example" // string | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.
    artifactId := "artifactId_example" // string | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier.
    version := "version_example" // string | The unique identifier of a specific version of the artifact content.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataAPI.GetArtifactVersionMetaData(context.Background(), groupId, artifactId, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.GetArtifactVersionMetaData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArtifactVersionMetaData`: VersionMetaData
    fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.GetArtifactVersionMetaData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. | 
**artifactId** | **string** | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier. | 
**version** | **string** | The unique identifier of a specific version of the artifact content. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtifactVersionMetaDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**VersionMetaData**](VersionMetaData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifactVersionMetaDataByContent

> VersionMetaData GetArtifactVersionMetaDataByContent(ctx, groupId, artifactId).Body(body).Canonical(canonical).Execute()

Get artifact version metadata by content



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Apicurio/apicurio-registry-client-sdk-go"
)

func main() {
    groupId := "groupId_example" // string | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.
    artifactId := "artifactId_example" // string | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier.
    body := os.NewFile(1234, "some_file") // *os.File | The content of an artifact version.
    canonical := true // bool | Parameter that can be set to `true` to indicate that the server should \"canonicalize\" the content when searching for a matching version.  Canonicalization is unique to each artifact type, but typically involves removing any extra whitespace and formatting the content in a consistent manner. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataAPI.GetArtifactVersionMetaDataByContent(context.Background(), groupId, artifactId).Body(body).Canonical(canonical).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.GetArtifactVersionMetaDataByContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArtifactVersionMetaDataByContent`: VersionMetaData
    fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.GetArtifactVersionMetaDataByContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. | 
**artifactId** | **string** | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtifactVersionMetaDataByContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | ***os.File** | The content of an artifact version. | 
 **canonical** | **bool** | Parameter that can be set to &#x60;true&#x60; to indicate that the server should \&quot;canonicalize\&quot; the content when searching for a matching version.  Canonicalization is unique to each artifact type, but typically involves removing any extra whitespace and formatting the content in a consistent manner. | 

### Return type

[**VersionMetaData**](VersionMetaData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/get.extended+json, application/vnd.get.extended+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateArtifactMetaData

> UpdateArtifactMetaData(ctx, groupId, artifactId).EditableMetaData(editableMetaData).Execute()

Update artifact metadata



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Apicurio/apicurio-registry-client-sdk-go"
)

func main() {
    groupId := "groupId_example" // string | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.
    artifactId := "artifactId_example" // string | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier.
    editableMetaData := *openapiclient.NewEditableMetaData() // EditableMetaData | Updated artifact metadata.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MetadataAPI.UpdateArtifactMetaData(context.Background(), groupId, artifactId).EditableMetaData(editableMetaData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.UpdateArtifactMetaData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. | 
**artifactId** | **string** | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateArtifactMetaDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **editableMetaData** | [**EditableMetaData**](EditableMetaData.md) | Updated artifact metadata. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateArtifactOwner

> UpdateArtifactOwner(ctx, groupId, artifactId).ArtifactOwner(artifactOwner).Execute()

Update artifact owner



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Apicurio/apicurio-registry-client-sdk-go"
)

func main() {
    groupId := "groupId_example" // string | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.
    artifactId := "artifactId_example" // string | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier.
    artifactOwner := *openapiclient.NewArtifactOwner() // ArtifactOwner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MetadataAPI.UpdateArtifactOwner(context.Background(), groupId, artifactId).ArtifactOwner(artifactOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.UpdateArtifactOwner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. | 
**artifactId** | **string** | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateArtifactOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **artifactOwner** | [**ArtifactOwner**](ArtifactOwner.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateArtifactVersionMetaData

> UpdateArtifactVersionMetaData(ctx, groupId, artifactId, version).EditableMetaData(editableMetaData).Execute()

Update artifact version metadata



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Apicurio/apicurio-registry-client-sdk-go"
)

func main() {
    groupId := "groupId_example" // string | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts.
    artifactId := "artifactId_example" // string | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier.
    version := "version_example" // string | The unique identifier of a specific version of the artifact content.
    editableMetaData := *openapiclient.NewEditableMetaData() // EditableMetaData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MetadataAPI.UpdateArtifactVersionMetaData(context.Background(), groupId, artifactId, version).EditableMetaData(editableMetaData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.UpdateArtifactVersionMetaData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. | 
**artifactId** | **string** | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier. | 
**version** | **string** | The unique identifier of a specific version of the artifact content. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateArtifactVersionMetaDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **editableMetaData** | [**EditableMetaData**](EditableMetaData.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

