# \ArtifactTypeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListArtifactTypes**](ArtifactTypeAPI.md#ListArtifactTypes) | **Get** /admin/artifactTypes | List artifact types



## ListArtifactTypes

> []ArtifactTypeInfo ListArtifactTypes(ctx).Execute()

List artifact types



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtifactTypeAPI.ListArtifactTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactTypeAPI.ListArtifactTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListArtifactTypes`: []ArtifactTypeInfo
    fmt.Fprintf(os.Stdout, "Response from `ArtifactTypeAPI.ListArtifactTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListArtifactTypesRequest struct via the builder pattern


### Return type

[**[]ArtifactTypeInfo**](ArtifactTypeInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

