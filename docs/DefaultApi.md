# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReferencesByContentHash**](DefaultApi.md#ReferencesByContentHash) | **Get** /ids/contentHashes/{contentHash}/references | Returns a list with all the references for the artifact with the given hash
[**ReferencesByContentId**](DefaultApi.md#ReferencesByContentId) | **Get** /ids/contentIds/{contentId}/references | Returns a list with all the references for the artifact with the given content id.
[**ReferencesByGlobalId**](DefaultApi.md#ReferencesByGlobalId) | **Get** /ids/globalIds/{globalId}/references | Returns a list with all the references for the artifact with the given global id.



## ReferencesByContentHash

> []ArtifactReference ReferencesByContentHash(ctx, UNKNOWN_PARAMETER_NAME).Execute()

Returns a list with all the references for the artifact with the given hash



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    UNKNOWN_PARAMETER_NAME := TODO //  | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ReferencesByContentHash(context.Background(), UNKNOWN_PARAMETER_NAME).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReferencesByContentHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReferencesByContentHash`: []ArtifactReference
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ReferencesByContentHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**UNKNOWN_PARAMETER_NAME** | [****](.md) |  | 

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

> []ArtifactReference ReferencesByContentId(ctx, UNKNOWN_PARAMETER_NAME).Execute()

Returns a list with all the references for the artifact with the given content id.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    UNKNOWN_PARAMETER_NAME := TODO //  | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ReferencesByContentId(context.Background(), UNKNOWN_PARAMETER_NAME).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReferencesByContentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReferencesByContentId`: []ArtifactReference
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ReferencesByContentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**UNKNOWN_PARAMETER_NAME** | [****](.md) |  | 

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

> []ArtifactReference ReferencesByGlobalId(ctx, UNKNOWN_PARAMETER_NAME).Execute()

Returns a list with all the references for the artifact with the given global id.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    UNKNOWN_PARAMETER_NAME := TODO //  | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ReferencesByGlobalId(context.Background(), UNKNOWN_PARAMETER_NAME).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ReferencesByGlobalId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReferencesByGlobalId`: []ArtifactReference
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ReferencesByGlobalId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**UNKNOWN_PARAMETER_NAME** | [****](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReferencesByGlobalIdRequest struct via the builder pattern


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

