# \ArtifactRulesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateArtifactRule**](ArtifactRulesAPI.md#CreateArtifactRule) | **Post** /groups/{groupId}/artifacts/{artifactId}/rules | Create artifact rule
[**DeleteArtifactRule**](ArtifactRulesAPI.md#DeleteArtifactRule) | **Delete** /groups/{groupId}/artifacts/{artifactId}/rules/{rule} | Delete artifact rule
[**DeleteArtifactRules**](ArtifactRulesAPI.md#DeleteArtifactRules) | **Delete** /groups/{groupId}/artifacts/{artifactId}/rules | Delete artifact rules
[**GetArtifactRuleConfig**](ArtifactRulesAPI.md#GetArtifactRuleConfig) | **Get** /groups/{groupId}/artifacts/{artifactId}/rules/{rule} | Get artifact rule configuration
[**ListArtifactRules**](ArtifactRulesAPI.md#ListArtifactRules) | **Get** /groups/{groupId}/artifacts/{artifactId}/rules | List artifact rules
[**TestUpdateArtifact**](ArtifactRulesAPI.md#TestUpdateArtifact) | **Put** /groups/{groupId}/artifacts/{artifactId}/test | Test update artifact
[**UpdateArtifactRuleConfig**](ArtifactRulesAPI.md#UpdateArtifactRuleConfig) | **Put** /groups/{groupId}/artifacts/{artifactId}/rules/{rule} | Update artifact rule configuration



## CreateArtifactRule

> CreateArtifactRule(ctx, groupId, artifactId).Rule(rule).Execute()

Create artifact rule



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
    rule := *openapiclient.NewRule("Config_example") // Rule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ArtifactRulesAPI.CreateArtifactRule(context.Background(), groupId, artifactId).Rule(rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactRulesAPI.CreateArtifactRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateArtifactRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rule** | [**Rule**](Rule.md) |  | 

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


## DeleteArtifactRule

> DeleteArtifactRule(ctx, groupId, artifactId, rule).Execute()

Delete artifact rule



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
    rule := "rule_example" // string | The unique name/type of a rule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ArtifactRulesAPI.DeleteArtifactRule(context.Background(), groupId, artifactId, rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactRulesAPI.DeleteArtifactRule``: %v\n", err)
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
**rule** | **string** | The unique name/type of a rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArtifactRuleRequest struct via the builder pattern


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


## DeleteArtifactRules

> DeleteArtifactRules(ctx, groupId, artifactId).Execute()

Delete artifact rules



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
    r, err := apiClient.ArtifactRulesAPI.DeleteArtifactRules(context.Background(), groupId, artifactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactRulesAPI.DeleteArtifactRules``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteArtifactRulesRequest struct via the builder pattern


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


## GetArtifactRuleConfig

> Rule GetArtifactRuleConfig(ctx, groupId, artifactId, rule).Execute()

Get artifact rule configuration



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
    rule := "rule_example" // string | The unique name/type of a rule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtifactRulesAPI.GetArtifactRuleConfig(context.Background(), groupId, artifactId, rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactRulesAPI.GetArtifactRuleConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArtifactRuleConfig`: Rule
    fmt.Fprintf(os.Stdout, "Response from `ArtifactRulesAPI.GetArtifactRuleConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. | 
**artifactId** | **string** | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier. | 
**rule** | **string** | The unique name/type of a rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtifactRuleConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Rule**](Rule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListArtifactRules

> []RuleType ListArtifactRules(ctx, groupId, artifactId).Execute()

List artifact rules



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
    resp, r, err := apiClient.ArtifactRulesAPI.ListArtifactRules(context.Background(), groupId, artifactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactRulesAPI.ListArtifactRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListArtifactRules`: []RuleType
    fmt.Fprintf(os.Stdout, "Response from `ArtifactRulesAPI.ListArtifactRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. | 
**artifactId** | **string** | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListArtifactRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]RuleType**](RuleType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestUpdateArtifact

> TestUpdateArtifact(ctx, groupId, artifactId).Body(body).Execute()

Test update artifact



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
    body := os.NewFile(1234, "some_file") // *os.File | The content of the artifact being tested. This is often, but not always, JSON data representing one of the supported artifact types:  * Avro (`AVRO`) * Protobuf (`PROTOBUF`) * JSON Schema (`JSON`) * Kafka Connect (`KCONNECT`) * OpenAPI (`OPENAPI`) * AsyncAPI (`ASYNCAPI`) * GraphQL (`GRAPHQL`) * Web Services Description Language (`WSDL`) * XML Schema (`XSD`) 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ArtifactRulesAPI.TestUpdateArtifact(context.Background(), groupId, artifactId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactRulesAPI.TestUpdateArtifact``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTestUpdateArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | ***os.File** | The content of the artifact being tested. This is often, but not always, JSON data representing one of the supported artifact types:  * Avro (&#x60;AVRO&#x60;) * Protobuf (&#x60;PROTOBUF&#x60;) * JSON Schema (&#x60;JSON&#x60;) * Kafka Connect (&#x60;KCONNECT&#x60;) * OpenAPI (&#x60;OPENAPI&#x60;) * AsyncAPI (&#x60;ASYNCAPI&#x60;) * GraphQL (&#x60;GRAPHQL&#x60;) * Web Services Description Language (&#x60;WSDL&#x60;) * XML Schema (&#x60;XSD&#x60;)  | 

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


## UpdateArtifactRuleConfig

> Rule UpdateArtifactRuleConfig(ctx, groupId, artifactId, rule).Rule2(rule2).Execute()

Update artifact rule configuration



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
    rule := "rule_example" // string | The unique name/type of a rule.
    rule2 := *openapiclient.NewRule("Config_example") // Rule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ArtifactRulesAPI.UpdateArtifactRuleConfig(context.Background(), groupId, artifactId, rule).Rule2(rule2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArtifactRulesAPI.UpdateArtifactRuleConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateArtifactRuleConfig`: Rule
    fmt.Fprintf(os.Stdout, "Response from `ArtifactRulesAPI.UpdateArtifactRuleConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The artifact group ID.  Must be a string provided by the client, representing the name of the grouping of artifacts. | 
**artifactId** | **string** | The artifact ID.  Can be a string (client-provided) or UUID (server-generated), representing the unique artifact identifier. | 
**rule** | **string** | The unique name/type of a rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateArtifactRuleConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **rule2** | [**Rule**](Rule.md) |  | 

### Return type

[**Rule**](Rule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

