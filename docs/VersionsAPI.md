# \VersionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddArtifactVersionComment**](VersionsAPI.md#AddArtifactVersionComment) | **Post** /groups/{groupId}/artifacts/{artifactId}/versions/{version}/comments | Add new comment
[**CreateArtifactVersion**](VersionsAPI.md#CreateArtifactVersion) | **Post** /groups/{groupId}/artifacts/{artifactId}/versions | Create artifact version
[**DeleteArtifactVersion**](VersionsAPI.md#DeleteArtifactVersion) | **Delete** /groups/{groupId}/artifacts/{artifactId}/versions/{version} | Delete artifact version
[**DeleteArtifactVersionComment**](VersionsAPI.md#DeleteArtifactVersionComment) | **Delete** /groups/{groupId}/artifacts/{artifactId}/versions/{version}/comments/{commentId} | Delete a single comment
[**GetArtifactVersion**](VersionsAPI.md#GetArtifactVersion) | **Get** /groups/{groupId}/artifacts/{artifactId}/versions/{version} | Get artifact version
[**GetArtifactVersionComments**](VersionsAPI.md#GetArtifactVersionComments) | **Get** /groups/{groupId}/artifacts/{artifactId}/versions/{version}/comments | Get artifact version comments
[**GetArtifactVersionReferences**](VersionsAPI.md#GetArtifactVersionReferences) | **Get** /groups/{groupId}/artifacts/{artifactId}/versions/{version}/references | Get artifact version references
[**ListArtifactVersions**](VersionsAPI.md#ListArtifactVersions) | **Get** /groups/{groupId}/artifacts/{artifactId}/versions | List artifact versions
[**UpdateArtifactVersionComment**](VersionsAPI.md#UpdateArtifactVersionComment) | **Put** /groups/{groupId}/artifacts/{artifactId}/versions/{version}/comments/{commentId} | Update a comment
[**UpdateArtifactVersionState**](VersionsAPI.md#UpdateArtifactVersionState) | **Put** /groups/{groupId}/artifacts/{artifactId}/versions/{version}/state | Update artifact version state



## AddArtifactVersionComment

> Comment AddArtifactVersionComment(ctx, groupId, artifactId, version).NewComment(newComment).Execute()

Add new comment



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
    newComment := *openapiclient.NewNewComment("Value_example") // NewComment | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VersionsAPI.AddArtifactVersionComment(context.Background(), groupId, artifactId, version).NewComment(newComment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsAPI.AddArtifactVersionComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddArtifactVersionComment`: Comment
    fmt.Fprintf(os.Stdout, "Response from `VersionsAPI.AddArtifactVersionComment`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiAddArtifactVersionCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **newComment** | [**NewComment**](NewComment.md) |  | 

### Return type

[**Comment**](Comment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateArtifactVersion

> VersionMetaData CreateArtifactVersion(ctx, groupId, artifactId).Body(body).XRegistryVersion(xRegistryVersion).XRegistryName(xRegistryName).XRegistryDescription(xRegistryDescription).XRegistryDescriptionEncoded(xRegistryDescriptionEncoded).XRegistryNameEncoded(xRegistryNameEncoded).Execute()

Create artifact version



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
    body := os.NewFile(1234, "some_file") // *os.File | The content of the artifact version being created or the content and a set of references to other artifacts. This is often, but not always, JSON data representing one of the supported artifact types:  * Avro (`AVRO`) * Protobuf (`PROTOBUF`) * JSON Schema (`JSON`) * Kafka Connect (`KCONNECT`) * OpenAPI (`OPENAPI`) * AsyncAPI (`ASYNCAPI`) * GraphQL (`GRAPHQL`) * Web Services Description Language (`WSDL`) * XML Schema (`XSD`) 
    xRegistryVersion := "xRegistryVersion_example" // string | Specifies the version number of this new version of the artifact content.  This would typically be a simple integer or a SemVer value.  It must be unique within the artifact.  If this is not provided, the server will generate a new, unique version number for this new updated content. (optional)
    xRegistryName := "xRegistryName_example" // string | Specifies the artifact name of this new version of the artifact content. Name must be ASCII-only string. If this is not provided, the server will extract the name from the artifact content. (optional)
    xRegistryDescription := "xRegistryDescription_example" // string | Specifies the artifact description of this new version of the artifact content. Description must be ASCII-only string. If this is not provided, the server will extract the description from the artifact content. (optional)
    xRegistryDescriptionEncoded := "xRegistryDescriptionEncoded_example" // string | Specifies the artifact description of this new version of the artifact content. Value of this must be Base64 encoded string. If this is not provided, the server will extract the description from the artifact content. (optional)
    xRegistryNameEncoded := "xRegistryNameEncoded_example" // string | Specifies the artifact name of this new version of the artifact content. Value of this must be Base64 encoded string. If this is not provided, the server will extract the name from the artifact content. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VersionsAPI.CreateArtifactVersion(context.Background(), groupId, artifactId).Body(body).XRegistryVersion(xRegistryVersion).XRegistryName(xRegistryName).XRegistryDescription(xRegistryDescription).XRegistryDescriptionEncoded(xRegistryDescriptionEncoded).XRegistryNameEncoded(xRegistryNameEncoded).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsAPI.CreateArtifactVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateArtifactVersion`: VersionMetaData
    fmt.Fprintf(os.Stdout, "Response from `VersionsAPI.CreateArtifactVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. | 
**artifactId** | **string** | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateArtifactVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | ***os.File** | The content of the artifact version being created or the content and a set of references to other artifacts. This is often, but not always, JSON data representing one of the supported artifact types:  * Avro (&#x60;AVRO&#x60;) * Protobuf (&#x60;PROTOBUF&#x60;) * JSON Schema (&#x60;JSON&#x60;) * Kafka Connect (&#x60;KCONNECT&#x60;) * OpenAPI (&#x60;OPENAPI&#x60;) * AsyncAPI (&#x60;ASYNCAPI&#x60;) * GraphQL (&#x60;GRAPHQL&#x60;) * Web Services Description Language (&#x60;WSDL&#x60;) * XML Schema (&#x60;XSD&#x60;)  | 
 **xRegistryVersion** | **string** | Specifies the version number of this new version of the artifact content.  This would typically be a simple integer or a SemVer value.  It must be unique within the artifact.  If this is not provided, the server will generate a new, unique version number for this new updated content. | 
 **xRegistryName** | **string** | Specifies the artifact name of this new version of the artifact content. Name must be ASCII-only string. If this is not provided, the server will extract the name from the artifact content. | 
 **xRegistryDescription** | **string** | Specifies the artifact description of this new version of the artifact content. Description must be ASCII-only string. If this is not provided, the server will extract the description from the artifact content. | 
 **xRegistryDescriptionEncoded** | **string** | Specifies the artifact description of this new version of the artifact content. Value of this must be Base64 encoded string. If this is not provided, the server will extract the description from the artifact content. | 
 **xRegistryNameEncoded** | **string** | Specifies the artifact name of this new version of the artifact content. Value of this must be Base64 encoded string. If this is not provided, the server will extract the name from the artifact content. | 

### Return type

[**VersionMetaData**](VersionMetaData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/create.extended+json, application/vnd.create.extended+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteArtifactVersion

> DeleteArtifactVersion(ctx, groupId, artifactId, version).Execute()

Delete artifact version



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
    r, err := apiClient.VersionsAPI.DeleteArtifactVersion(context.Background(), groupId, artifactId, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsAPI.DeleteArtifactVersion``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteArtifactVersionRequest struct via the builder pattern


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


## DeleteArtifactVersionComment

> DeleteArtifactVersionComment(ctx, groupId, artifactId, version, commentId).Execute()

Delete a single comment



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
    commentId := "commentId_example" // string | The unique identifier of a single comment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VersionsAPI.DeleteArtifactVersionComment(context.Background(), groupId, artifactId, version, commentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsAPI.DeleteArtifactVersionComment``: %v\n", err)
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
**commentId** | **string** | The unique identifier of a single comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArtifactVersionCommentRequest struct via the builder pattern


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


## GetArtifactVersion

> *os.File GetArtifactVersion(ctx, groupId, artifactId, version).References(references).Execute()

Get artifact version



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
    references := openapiclient.HandleReferencesType("PRESERVE") // HandleReferencesType | Allows the user to specify how references in the content should be treated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VersionsAPI.GetArtifactVersion(context.Background(), groupId, artifactId, version).References(references).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsAPI.GetArtifactVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArtifactVersion`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `VersionsAPI.GetArtifactVersion`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetArtifactVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **references** | [**HandleReferencesType**](HandleReferencesType.md) | Allows the user to specify how references in the content should be treated. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifactVersionComments

> []Comment GetArtifactVersionComments(ctx, groupId, artifactId, version).Execute()

Get artifact version comments



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
    resp, r, err := apiClient.VersionsAPI.GetArtifactVersionComments(context.Background(), groupId, artifactId, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsAPI.GetArtifactVersionComments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArtifactVersionComments`: []Comment
    fmt.Fprintf(os.Stdout, "Response from `VersionsAPI.GetArtifactVersionComments`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetArtifactVersionCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]Comment**](Comment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifactVersionReferences

> []ArtifactReference GetArtifactVersionReferences(ctx, groupId, artifactId, version).RefType(refType).Execute()

Get artifact version references



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
    refType := openapiclient.ReferenceType("OUTBOUND") // ReferenceType | Determines the type of reference to return, either INBOUND or OUTBOUND.  Defaults to OUTBOUND. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VersionsAPI.GetArtifactVersionReferences(context.Background(), groupId, artifactId, version).RefType(refType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsAPI.GetArtifactVersionReferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArtifactVersionReferences`: []ArtifactReference
    fmt.Fprintf(os.Stdout, "Response from `VersionsAPI.GetArtifactVersionReferences`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetArtifactVersionReferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **refType** | [**ReferenceType**](ReferenceType.md) | Determines the type of reference to return, either INBOUND or OUTBOUND.  Defaults to OUTBOUND. | 

### Return type

[**[]ArtifactReference**](ArtifactReference.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListArtifactVersions

> VersionSearchResults ListArtifactVersions(ctx, groupId, artifactId).Offset(offset).Limit(limit).Execute()

List artifact versions



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
    offset := int32(56) // int32 | The number of versions to skip before starting to collect the result set.  Defaults to 0. (optional)
    limit := int32(56) // int32 | The number of versions to return.  Defaults to 20. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VersionsAPI.ListArtifactVersions(context.Background(), groupId, artifactId).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsAPI.ListArtifactVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListArtifactVersions`: VersionSearchResults
    fmt.Fprintf(os.Stdout, "Response from `VersionsAPI.ListArtifactVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. | 
**artifactId** | **string** | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListArtifactVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **int32** | The number of versions to skip before starting to collect the result set.  Defaults to 0. | 
 **limit** | **int32** | The number of versions to return.  Defaults to 20. | 

### Return type

[**VersionSearchResults**](VersionSearchResults.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateArtifactVersionComment

> UpdateArtifactVersionComment(ctx, groupId, artifactId, version, commentId).NewComment(newComment).Execute()

Update a comment



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
    commentId := "commentId_example" // string | The unique identifier of a single comment.
    newComment := *openapiclient.NewNewComment("Value_example") // NewComment | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VersionsAPI.UpdateArtifactVersionComment(context.Background(), groupId, artifactId, version, commentId).NewComment(newComment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsAPI.UpdateArtifactVersionComment``: %v\n", err)
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
**commentId** | **string** | The unique identifier of a single comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateArtifactVersionCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **newComment** | [**NewComment**](NewComment.md) |  | 

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


## UpdateArtifactVersionState

> UpdateArtifactVersionState(ctx, groupId, artifactId, version).UpdateState(updateState).Execute()

Update artifact version state



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
    updateState := *openapiclient.NewUpdateState(openapiclient.ArtifactState("ENABLED")) // UpdateState | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VersionsAPI.UpdateArtifactVersionState(context.Background(), groupId, artifactId, version).UpdateState(updateState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionsAPI.UpdateArtifactVersionState``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateArtifactVersionStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateState** | [**UpdateState**](UpdateState.md) |  | 

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

