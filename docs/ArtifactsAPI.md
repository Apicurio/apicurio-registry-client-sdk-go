# \ArtifactsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateArtifact**](ArtifactsAPI.md#CreateArtifact) | **Post** /groups/{groupId}/artifacts | Create artifact
[**DeleteArtifact**](ArtifactsAPI.md#DeleteArtifact) | **Delete** /groups/{groupId}/artifacts/{artifactId} | Delete artifact
[**DeleteArtifactsInGroup**](ArtifactsAPI.md#DeleteArtifactsInGroup) | **Delete** /groups/{groupId}/artifacts | Delete artifacts in group
[**GetContentByGlobalId**](ArtifactsAPI.md#GetContentByGlobalId) | **Get** /ids/globalIds/{globalId} | Get artifact by global ID
[**GetContentByHash**](ArtifactsAPI.md#GetContentByHash) | **Get** /ids/contentHashes/{contentHash}/ | Get artifact content by SHA-256 hash
[**GetContentById**](ArtifactsAPI.md#GetContentById) | **Get** /ids/contentIds/{contentId}/ | Get artifact content by ID
[**GetLatestArtifact**](ArtifactsAPI.md#GetLatestArtifact) | **Get** /groups/{groupId}/artifacts/{artifactId} | Get latest artifact
[**ListArtifactsInGroup**](ArtifactsAPI.md#ListArtifactsInGroup) | **Get** /groups/{groupId}/artifacts | List artifacts in group
[**ReferencesByContentHash**](ArtifactsAPI.md#ReferencesByContentHash) | **Get** /ids/contentHashes/{contentHash}/references | List artifact references by hash
[**ReferencesByContentId**](ArtifactsAPI.md#ReferencesByContentId) | **Get** /ids/contentIds/{contentId}/references | List artifact references by content ID
[**ReferencesByGlobalId**](ArtifactsAPI.md#ReferencesByGlobalId) | **Get** /ids/globalIds/{globalId}/references | List artifact references by global ID
[**SearchArtifacts**](ArtifactsAPI.md#SearchArtifacts) | **Get** /search/artifacts | Search for artifacts
[**SearchArtifactsByContent**](ArtifactsAPI.md#SearchArtifactsByContent) | **Post** /search/artifacts | Search for artifacts by content
[**UpdateArtifact**](ArtifactsAPI.md#UpdateArtifact) | **Put** /groups/{groupId}/artifacts/{artifactId} | Update artifact
[**UpdateArtifactState**](ArtifactsAPI.md#UpdateArtifactState) | **Put** /groups/{groupId}/artifacts/{artifactId}/state | Update artifact state



## CreateArtifact

> ArtifactMetaData CreateArtifact(ctx, groupId).Body(body).XRegistryArtifactType(xRegistryArtifactType).XRegistryArtifactId(xRegistryArtifactId).XRegistryVersion(xRegistryVersion).IfExists(ifExists).Canonical(canonical).XRegistryDescription(xRegistryDescription).XRegistryDescriptionEncoded(xRegistryDescriptionEncoded).XRegistryName(xRegistryName).XRegistryNameEncoded(xRegistryNameEncoded).XRegistryContentHash(xRegistryContentHash).XRegistryHashAlgorithm(xRegistryHashAlgorithm).Execute()

Create artifact



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
    body := os.NewFile(1234, "some_file") // *os.File | The content of the artifact being created. This is often, but not always, JSON data representing one of the supported artifact types:  * Avro (`AVRO`) * Protobuf (`PROTOBUF`) * JSON Schema (`JSON`) * Kafka Connect (`KCONNECT`) * OpenAPI (`OPENAPI`) * AsyncAPI (`ASYNCAPI`) * GraphQL (`GRAPHQL`) * Web Services Description Language (`WSDL`) * XML Schema (`XSD`) 
    xRegistryArtifactType := "xRegistryArtifactType_example" // string | Specifies the type of the artifact being added. Possible values include:  * Avro (`AVRO`) * Protobuf (`PROTOBUF`) * JSON Schema (`JSON`) * Kafka Connect (`KCONNECT`) * OpenAPI (`OPENAPI`) * AsyncAPI (`ASYNCAPI`) * GraphQL (`GRAPHQL`) * Web Services Description Language (`WSDL`) * XML Schema (`XSD`) (optional)
    xRegistryArtifactId := "xRegistryArtifactId_example" // string | A client-provided, globally unique identifier for the new artifact. (optional)
    xRegistryVersion := "xRegistryVersion_example" // string | Specifies the version number of this initial version of the artifact content.  This would typically be a simple integer or a SemVer value.  If not provided, the server will assign a version number automatically (starting with version `1`). (optional)
    ifExists := openapiclient.IfExists("FAIL") // IfExists | Set this option to instruct the server on what to do if the artifact already exists. (optional)
    canonical := true // bool | Used only when the `ifExists` query parameter is set to `RETURN_OR_UPDATE`, this parameter can be set to `true` to indicate that the server should \"canonicalize\" the content when searching for a matching version.  The canonicalization algorithm is unique to each artifact type, but typically involves removing extra whitespace and formatting the content in a consistent manner. (optional)
    xRegistryDescription := "xRegistryDescription_example" // string | Specifies the description of artifact being added. Description must be ASCII-only string. If this is not provided, the server will extract the description from the artifact content. (optional)
    xRegistryDescriptionEncoded := "xRegistryDescriptionEncoded_example" // string | Specifies the description of artifact being added. Value of this must be Base64 encoded string. If this is not provided, the server will extract the description from the artifact content. (optional)
    xRegistryName := "xRegistryName_example" // string | Specifies the name of artifact being added. Name must be ASCII-only string. If this is not provided, the server will extract the name from the artifact content. (optional)
    xRegistryNameEncoded := "xRegistryNameEncoded_example" // string | Specifies the name of artifact being added. Value of this must be Base64 encoded string. If this is not provided, the server will extract the name from the artifact content. (optional)
    xRegistryContentHash := "xRegistryContentHash_example" // string | Specifies the (optional) hash of the artifact to be verified. (optional)
    xRegistryHashAlgorithm := "xRegistryHashAlgorithm_example" // string | The algorithm to use when checking the content validity. (available: SHA256, MD5; default: SHA256) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtifactsAPI.CreateArtifact(context.Background(), groupId).Body(body).XRegistryArtifactType(xRegistryArtifactType).XRegistryArtifactId(xRegistryArtifactId).XRegistryVersion(xRegistryVersion).IfExists(ifExists).Canonical(canonical).XRegistryDescription(xRegistryDescription).XRegistryDescriptionEncoded(xRegistryDescriptionEncoded).XRegistryName(xRegistryName).XRegistryNameEncoded(xRegistryNameEncoded).XRegistryContentHash(xRegistryContentHash).XRegistryHashAlgorithm(xRegistryHashAlgorithm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.CreateArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateArtifact`: ArtifactMetaData
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.CreateArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** | The content of the artifact being created. This is often, but not always, JSON data representing one of the supported artifact types:  * Avro (&#x60;AVRO&#x60;) * Protobuf (&#x60;PROTOBUF&#x60;) * JSON Schema (&#x60;JSON&#x60;) * Kafka Connect (&#x60;KCONNECT&#x60;) * OpenAPI (&#x60;OPENAPI&#x60;) * AsyncAPI (&#x60;ASYNCAPI&#x60;) * GraphQL (&#x60;GRAPHQL&#x60;) * Web Services Description Language (&#x60;WSDL&#x60;) * XML Schema (&#x60;XSD&#x60;)  | 
 **xRegistryArtifactType** | **string** | Specifies the type of the artifact being added. Possible values include:  * Avro (&#x60;AVRO&#x60;) * Protobuf (&#x60;PROTOBUF&#x60;) * JSON Schema (&#x60;JSON&#x60;) * Kafka Connect (&#x60;KCONNECT&#x60;) * OpenAPI (&#x60;OPENAPI&#x60;) * AsyncAPI (&#x60;ASYNCAPI&#x60;) * GraphQL (&#x60;GRAPHQL&#x60;) * Web Services Description Language (&#x60;WSDL&#x60;) * XML Schema (&#x60;XSD&#x60;) | 
 **xRegistryArtifactId** | **string** | A client-provided, globally unique identifier for the new artifact. | 
 **xRegistryVersion** | **string** | Specifies the version number of this initial version of the artifact content.  This would typically be a simple integer or a SemVer value.  If not provided, the server will assign a version number automatically (starting with version &#x60;1&#x60;). | 
 **ifExists** | [**IfExists**](IfExists.md) | Set this option to instruct the server on what to do if the artifact already exists. | 
 **canonical** | **bool** | Used only when the &#x60;ifExists&#x60; query parameter is set to &#x60;RETURN_OR_UPDATE&#x60;, this parameter can be set to &#x60;true&#x60; to indicate that the server should \&quot;canonicalize\&quot; the content when searching for a matching version.  The canonicalization algorithm is unique to each artifact type, but typically involves removing extra whitespace and formatting the content in a consistent manner. | 
 **xRegistryDescription** | **string** | Specifies the description of artifact being added. Description must be ASCII-only string. If this is not provided, the server will extract the description from the artifact content. | 
 **xRegistryDescriptionEncoded** | **string** | Specifies the description of artifact being added. Value of this must be Base64 encoded string. If this is not provided, the server will extract the description from the artifact content. | 
 **xRegistryName** | **string** | Specifies the name of artifact being added. Name must be ASCII-only string. If this is not provided, the server will extract the name from the artifact content. | 
 **xRegistryNameEncoded** | **string** | Specifies the name of artifact being added. Value of this must be Base64 encoded string. If this is not provided, the server will extract the name from the artifact content. | 
 **xRegistryContentHash** | **string** | Specifies the (optional) hash of the artifact to be verified. | 
 **xRegistryHashAlgorithm** | **string** | The algorithm to use when checking the content validity. (available: SHA256, MD5; default: SHA256) | 

### Return type

[**ArtifactMetaData**](ArtifactMetaData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/create.extended+json, application/vnd.create.extended+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteArtifact

> DeleteArtifact(ctx, groupId, artifactId).Execute()

Delete artifact



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
    r, err := apiClient.ArtifactsAPI.DeleteArtifact(context.Background(), groupId, artifactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.DeleteArtifact``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteArtifactRequest struct via the builder pattern


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


## DeleteArtifactsInGroup

> DeleteArtifactsInGroup(ctx, groupId).Execute()

Delete artifacts in group



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ArtifactsAPI.DeleteArtifactsInGroup(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.DeleteArtifactsInGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArtifactsInGroupRequest struct via the builder pattern


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


## GetContentByGlobalId

> *os.File GetContentByGlobalId(ctx, globalId).References(references).Execute()

Get artifact by global ID



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
    globalId := int64(789) // int64 | Global identifier for an artifact version.
    references := openapiclient.HandleReferencesType("PRESERVE") // HandleReferencesType | Allows the user to specify how references in the content should be treated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtifactsAPI.GetContentByGlobalId(context.Background(), globalId).References(references).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.GetContentByGlobalId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContentByGlobalId`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.GetContentByGlobalId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**globalId** | **int64** | Global identifier for an artifact version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContentByGlobalIdRequest struct via the builder pattern


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


## GetContentByHash

> *os.File GetContentByHash(ctx, contentHash).Execute()

Get artifact content by SHA-256 hash



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
    contentHash := "contentHash_example" // string | SHA-256 content hash for a single artifact content.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtifactsAPI.GetContentByHash(context.Background(), contentHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.GetContentByHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContentByHash`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.GetContentByHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentHash** | **string** | SHA-256 content hash for a single artifact content. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContentByHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetContentById

> *os.File GetContentById(ctx, contentId).Execute()

Get artifact content by ID



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
    contentId := int64(789) // int64 | Global identifier for a single artifact content.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtifactsAPI.GetContentById(context.Background(), contentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.GetContentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContentById`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.GetContentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentId** | **int64** | Global identifier for a single artifact content. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetLatestArtifact

> *os.File GetLatestArtifact(ctx, groupId, artifactId).References(references).Execute()

Get latest artifact



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
    references := openapiclient.HandleReferencesType("PRESERVE") // HandleReferencesType | Allows the user to specify how references in the content should be treated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtifactsAPI.GetLatestArtifact(context.Background(), groupId, artifactId).References(references).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.GetLatestArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLatestArtifact`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.GetLatestArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. | 
**artifactId** | **string** | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestArtifactRequest struct via the builder pattern


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


## ListArtifactsInGroup

> ArtifactSearchResults ListArtifactsInGroup(ctx, groupId).Limit(limit).Offset(offset).Order(order).Orderby(orderby).Execute()

List artifacts in group



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
    limit := int32(56) // int32 | The number of artifacts to return.  Defaults to 20. (optional)
    offset := int32(56) // int32 | The number of artifacts to skip before starting the result set.  Defaults to 0. (optional)
    order := openapiclient.SortOrder("asc") // SortOrder | Sort order, ascending (`asc`) or descending (`desc`). (optional)
    orderby := openapiclient.SortBy("name") // SortBy | The field to sort by.  Can be one of:  * `name` * `createdOn`  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtifactsAPI.ListArtifactsInGroup(context.Background(), groupId).Limit(limit).Offset(offset).Order(order).Orderby(orderby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.ListArtifactsInGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListArtifactsInGroup`: ArtifactSearchResults
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.ListArtifactsInGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListArtifactsInGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The number of artifacts to return.  Defaults to 20. | 
 **offset** | **int32** | The number of artifacts to skip before starting the result set.  Defaults to 0. | 
 **order** | [**SortOrder**](SortOrder.md) | Sort order, ascending (&#x60;asc&#x60;) or descending (&#x60;desc&#x60;). | 
 **orderby** | [**SortBy**](SortBy.md) | The field to sort by.  Can be one of:  * &#x60;name&#x60; * &#x60;createdOn&#x60;  | 

### Return type

[**ArtifactSearchResults**](ArtifactSearchResults.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferencesByContentHash

> []ArtifactReference ReferencesByContentHash(ctx, contentHash).Execute()

List artifact references by hash



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
    contentHash := "contentHash_example" // string | SHA-256 content hash for a single artifact content.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtifactsAPI.ReferencesByContentHash(context.Background(), contentHash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.ReferencesByContentHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReferencesByContentHash`: []ArtifactReference
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.ReferencesByContentHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentHash** | **string** | SHA-256 content hash for a single artifact content. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReferencesByContentHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ReferencesByContentId

> []ArtifactReference ReferencesByContentId(ctx, contentId).Execute()

List artifact references by content ID



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
    contentId := int64(789) // int64 | Global identifier for a single artifact content.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtifactsAPI.ReferencesByContentId(context.Background(), contentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.ReferencesByContentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReferencesByContentId`: []ArtifactReference
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.ReferencesByContentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentId** | **int64** | Global identifier for a single artifact content. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReferencesByContentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ReferencesByGlobalId

> []ArtifactReference ReferencesByGlobalId(ctx, globalId).RefType(refType).Execute()

List artifact references by global ID



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
    globalId := int64(789) // int64 | Global identifier for an artifact version.
    refType := openapiclient.ReferenceType("OUTBOUND") // ReferenceType | Determines the type of reference to return, either INBOUND or OUTBOUND.  Defaults to OUTBOUND. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtifactsAPI.ReferencesByGlobalId(context.Background(), globalId).RefType(refType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.ReferencesByGlobalId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReferencesByGlobalId`: []ArtifactReference
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.ReferencesByGlobalId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**globalId** | **int64** | Global identifier for an artifact version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReferencesByGlobalIdRequest struct via the builder pattern


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


## SearchArtifacts

> ArtifactSearchResults SearchArtifacts(ctx).Name(name).Offset(offset).Limit(limit).Order(order).Orderby(orderby).Labels(labels).Properties(properties).Description(description).Group(group).GlobalId(globalId).ContentId(contentId).Execute()

Search for artifacts



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
    name := "name_example" // string | Filter by artifact name. (optional)
    offset := int32(56) // int32 | The number of artifacts to skip before starting to collect the result set.  Defaults to 0. (optional) (default to 0)
    limit := int32(56) // int32 | The number of artifacts to return.  Defaults to 20. (optional) (default to 20)
    order := openapiclient.SortOrder("asc") // SortOrder | Sort order, ascending (`asc`) or descending (`desc`). (optional)
    orderby := openapiclient.SortBy("name") // SortBy | The field to sort by.  Can be one of:  * `name` * `createdOn`  (optional)
    labels := []string{"Inner_example"} // []string | Filter by label.  Include one or more label to only return artifacts containing all of the specified labels. (optional)
    properties := []string{"Inner_example"} // []string | Filter by one or more name/value property.  Separate each name/value pair using a colon.  For example `properties=foo:bar` will return only artifacts with a custom property named `foo` and value `bar`. (optional)
    description := "description_example" // string | Filter by description. (optional)
    group := "group_example" // string | Filter by artifact group. (optional)
    globalId := int64(789) // int64 | Filter by globalId. (optional)
    contentId := int64(789) // int64 | Filter by contentId. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtifactsAPI.SearchArtifacts(context.Background()).Name(name).Offset(offset).Limit(limit).Order(order).Orderby(orderby).Labels(labels).Properties(properties).Description(description).Group(group).GlobalId(globalId).ContentId(contentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.SearchArtifacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchArtifacts`: ArtifactSearchResults
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.SearchArtifacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by artifact name. | 
 **offset** | **int32** | The number of artifacts to skip before starting to collect the result set.  Defaults to 0. | [default to 0]
 **limit** | **int32** | The number of artifacts to return.  Defaults to 20. | [default to 20]
 **order** | [**SortOrder**](SortOrder.md) | Sort order, ascending (&#x60;asc&#x60;) or descending (&#x60;desc&#x60;). | 
 **orderby** | [**SortBy**](SortBy.md) | The field to sort by.  Can be one of:  * &#x60;name&#x60; * &#x60;createdOn&#x60;  | 
 **labels** | **[]string** | Filter by label.  Include one or more label to only return artifacts containing all of the specified labels. | 
 **properties** | **[]string** | Filter by one or more name/value property.  Separate each name/value pair using a colon.  For example &#x60;properties&#x3D;foo:bar&#x60; will return only artifacts with a custom property named &#x60;foo&#x60; and value &#x60;bar&#x60;. | 
 **description** | **string** | Filter by description. | 
 **group** | **string** | Filter by artifact group. | 
 **globalId** | **int64** | Filter by globalId. | 
 **contentId** | **int64** | Filter by contentId. | 

### Return type

[**ArtifactSearchResults**](ArtifactSearchResults.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchArtifactsByContent

> ArtifactSearchResults SearchArtifactsByContent(ctx).Body(body).Canonical(canonical).ArtifactType(artifactType).Offset(offset).Limit(limit).Order(order).Orderby(orderby).Execute()

Search for artifacts by content



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
    body := os.NewFile(1234, "some_file") // *os.File | The content to search for.
    canonical := true // bool | Parameter that can be set to `true` to indicate that the server should \"canonicalize\" the content when searching for matching artifacts.  Canonicalization is unique to each artifact type, but typically involves removing any extra whitespace and formatting the content in a consistent manner.  Must be used along with the `artifactType` query parameter. (optional)
    artifactType := "artifactType_example" // string | Indicates the type of artifact represented by the content being used for the search.  This is only needed when using the `canonical` query parameter, so that the server knows how to canonicalize the content prior to searching for matching artifacts. (optional)
    offset := int32(56) // int32 | The number of artifacts to skip before starting to collect the result set.  Defaults to 0. (optional) (default to 0)
    limit := int32(56) // int32 | The number of artifacts to return.  Defaults to 20. (optional) (default to 20)
    order := "order_example" // string | Sort order, ascending (`asc`) or descending (`desc`). (optional)
    orderby := "orderby_example" // string | The field to sort by.  Can be one of:  * `name` * `createdOn`  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtifactsAPI.SearchArtifactsByContent(context.Background()).Body(body).Canonical(canonical).ArtifactType(artifactType).Offset(offset).Limit(limit).Order(order).Orderby(orderby).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.SearchArtifactsByContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchArtifactsByContent`: ArtifactSearchResults
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.SearchArtifactsByContent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchArtifactsByContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** | The content to search for. | 
 **canonical** | **bool** | Parameter that can be set to &#x60;true&#x60; to indicate that the server should \&quot;canonicalize\&quot; the content when searching for matching artifacts.  Canonicalization is unique to each artifact type, but typically involves removing any extra whitespace and formatting the content in a consistent manner.  Must be used along with the &#x60;artifactType&#x60; query parameter. | 
 **artifactType** | **string** | Indicates the type of artifact represented by the content being used for the search.  This is only needed when using the &#x60;canonical&#x60; query parameter, so that the server knows how to canonicalize the content prior to searching for matching artifacts. | 
 **offset** | **int32** | The number of artifacts to skip before starting to collect the result set.  Defaults to 0. | [default to 0]
 **limit** | **int32** | The number of artifacts to return.  Defaults to 20. | [default to 20]
 **order** | **string** | Sort order, ascending (&#x60;asc&#x60;) or descending (&#x60;desc&#x60;). | 
 **orderby** | **string** | The field to sort by.  Can be one of:  * &#x60;name&#x60; * &#x60;createdOn&#x60;  | 

### Return type

[**ArtifactSearchResults**](ArtifactSearchResults.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateArtifact

> ArtifactMetaData UpdateArtifact(ctx, groupId, artifactId).Body(body).XRegistryVersion(xRegistryVersion).XRegistryName(xRegistryName).XRegistryNameEncoded(xRegistryNameEncoded).XRegistryDescription(xRegistryDescription).XRegistryDescriptionEncoded(xRegistryDescriptionEncoded).Execute()

Update artifact



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
    body := os.NewFile(1234, "some_file") // *os.File | The new content of the artifact being updated. This is often, but not always, JSON data representing one of the supported artifact types:  * Avro (`AVRO`) * Protobuf (`PROTOBUF`) * JSON Schema (`JSON`) * Kafka Connect (`KCONNECT`) * OpenAPI (`OPENAPI`) * AsyncAPI (`ASYNCAPI`) * GraphQL (`GRAPHQL`) * Web Services Description Language (`WSDL`) * XML Schema (`XSD`) 
    xRegistryVersion := "xRegistryVersion_example" // string | Specifies the version number of this new version of the artifact content.  This would typically be a simple integer or a SemVer value.  If not provided, the server will assign a version number automatically. (optional)
    xRegistryName := "xRegistryName_example" // string | Specifies the artifact name of this new version of the artifact content. Name must be ASCII-only string. If this is not provided, the server will extract the name from the artifact content. (optional)
    xRegistryNameEncoded := "xRegistryNameEncoded_example" // string | Specifies the artifact name of this new version of the artifact content. Value of this must be Base64 encoded string. If this is not provided, the server will extract the name from the artifact content. (optional)
    xRegistryDescription := "xRegistryDescription_example" // string | Specifies the artifact description of this new version of the artifact content. Description must be ASCII-only string. If this is not provided, the server will extract the description from the artifact content. (optional)
    xRegistryDescriptionEncoded := "xRegistryDescriptionEncoded_example" // string | Specifies the artifact description of this new version of the artifact content. Value of this must be Base64 encoded string. If this is not provided, the server will extract the description from the artifact content. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtifactsAPI.UpdateArtifact(context.Background(), groupId, artifactId).Body(body).XRegistryVersion(xRegistryVersion).XRegistryName(xRegistryName).XRegistryNameEncoded(xRegistryNameEncoded).XRegistryDescription(xRegistryDescription).XRegistryDescriptionEncoded(xRegistryDescriptionEncoded).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.UpdateArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateArtifact`: ArtifactMetaData
    fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.UpdateArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. | 
**artifactId** | **string** | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | ***os.File** | The new content of the artifact being updated. This is often, but not always, JSON data representing one of the supported artifact types:  * Avro (&#x60;AVRO&#x60;) * Protobuf (&#x60;PROTOBUF&#x60;) * JSON Schema (&#x60;JSON&#x60;) * Kafka Connect (&#x60;KCONNECT&#x60;) * OpenAPI (&#x60;OPENAPI&#x60;) * AsyncAPI (&#x60;ASYNCAPI&#x60;) * GraphQL (&#x60;GRAPHQL&#x60;) * Web Services Description Language (&#x60;WSDL&#x60;) * XML Schema (&#x60;XSD&#x60;)  | 
 **xRegistryVersion** | **string** | Specifies the version number of this new version of the artifact content.  This would typically be a simple integer or a SemVer value.  If not provided, the server will assign a version number automatically. | 
 **xRegistryName** | **string** | Specifies the artifact name of this new version of the artifact content. Name must be ASCII-only string. If this is not provided, the server will extract the name from the artifact content. | 
 **xRegistryNameEncoded** | **string** | Specifies the artifact name of this new version of the artifact content. Value of this must be Base64 encoded string. If this is not provided, the server will extract the name from the artifact content. | 
 **xRegistryDescription** | **string** | Specifies the artifact description of this new version of the artifact content. Description must be ASCII-only string. If this is not provided, the server will extract the description from the artifact content. | 
 **xRegistryDescriptionEncoded** | **string** | Specifies the artifact description of this new version of the artifact content. Value of this must be Base64 encoded string. If this is not provided, the server will extract the description from the artifact content. | 

### Return type

[**ArtifactMetaData**](ArtifactMetaData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/create.extended+json, application/vnd.create.extended+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateArtifactState

> UpdateArtifactState(ctx, groupId, artifactId).UpdateState(updateState).Execute()

Update artifact state



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
    updateState := *openapiclient.NewUpdateState(openapiclient.ArtifactState("ENABLED")) // UpdateState | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ArtifactsAPI.UpdateArtifactState(context.Background(), groupId, artifactId).UpdateState(updateState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.UpdateArtifactState``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateArtifactStateRequest struct via the builder pattern


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

