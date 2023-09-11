# \AdminAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGlobalRule**](AdminAPI.md#CreateGlobalRule) | **Post** /admin/rules | Create global rule
[**CreateRoleMapping**](AdminAPI.md#CreateRoleMapping) | **Post** /admin/roleMappings | Create a new role mapping
[**DeleteAllGlobalRules**](AdminAPI.md#DeleteAllGlobalRules) | **Delete** /admin/rules | Delete all global rules
[**DeleteGlobalRule**](AdminAPI.md#DeleteGlobalRule) | **Delete** /admin/rules/{rule} | Delete global rule
[**DeleteRoleMapping**](AdminAPI.md#DeleteRoleMapping) | **Delete** /admin/roleMappings/{principalId} | Delete a role mapping
[**ExportData**](AdminAPI.md#ExportData) | **Get** /admin/export | Export registry data
[**GetConfigProperty**](AdminAPI.md#GetConfigProperty) | **Get** /admin/config/properties/{propertyName} | Get configuration property value
[**GetGlobalRuleConfig**](AdminAPI.md#GetGlobalRuleConfig) | **Get** /admin/rules/{rule} | Get global rule configuration
[**GetRoleMapping**](AdminAPI.md#GetRoleMapping) | **Get** /admin/roleMappings/{principalId} | Return a single role mapping
[**ImportData**](AdminAPI.md#ImportData) | **Post** /admin/import | Import registry data
[**ListArtifactTypes**](AdminAPI.md#ListArtifactTypes) | **Get** /admin/artifactTypes | List artifact types
[**ListConfigProperties**](AdminAPI.md#ListConfigProperties) | **Get** /admin/config/properties | List all configuration properties
[**ListGlobalRules**](AdminAPI.md#ListGlobalRules) | **Get** /admin/rules | List global rules
[**ListRoleMappings**](AdminAPI.md#ListRoleMappings) | **Get** /admin/roleMappings | List all role mappings
[**ResetConfigProperty**](AdminAPI.md#ResetConfigProperty) | **Delete** /admin/config/properties/{propertyName} | Reset a configuration property
[**UpdateConfigProperty**](AdminAPI.md#UpdateConfigProperty) | **Put** /admin/config/properties/{propertyName} | Update a configuration property
[**UpdateGlobalRuleConfig**](AdminAPI.md#UpdateGlobalRuleConfig) | **Put** /admin/rules/{rule} | Update global rule configuration
[**UpdateRoleMapping**](AdminAPI.md#UpdateRoleMapping) | **Put** /admin/roleMappings/{principalId} | Update a role mapping



## CreateGlobalRule

> CreateGlobalRule(ctx).Rule(rule).Execute()

Create global rule



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
    rule := *openapiclient.NewRule("Config_example") // Rule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.CreateGlobalRule(context.Background()).Rule(rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CreateGlobalRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGlobalRuleRequest struct via the builder pattern


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


## CreateRoleMapping

> CreateRoleMapping(ctx).RoleMapping(roleMapping).Execute()

Create a new role mapping



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
    roleMapping := *openapiclient.NewRoleMapping("PrincipalId_example", openapiclient.RoleType("READ_ONLY")) // RoleMapping | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.CreateRoleMapping(context.Background()).RoleMapping(roleMapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.CreateRoleMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleMapping** | [**RoleMapping**](RoleMapping.md) |  | 

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


## DeleteAllGlobalRules

> DeleteAllGlobalRules(ctx).Execute()

Delete all global rules



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
    r, err := apiClient.AdminAPI.DeleteAllGlobalRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.DeleteAllGlobalRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllGlobalRulesRequest struct via the builder pattern


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


## DeleteGlobalRule

> DeleteGlobalRule(ctx, rule).Execute()

Delete global rule



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
    rule := openapiclient.RuleType("VALIDITY") // RuleType | The unique name/type of a rule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.DeleteGlobalRule(context.Background(), rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.DeleteGlobalRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rule** | [**RuleType**](.md) | The unique name/type of a rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGlobalRuleRequest struct via the builder pattern


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


## DeleteRoleMapping

> DeleteRoleMapping(ctx, principalId).Execute()

Delete a role mapping



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
    principalId := "principalId_example" // string | Unique id of a principal (typically either a user or service account).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.DeleteRoleMapping(context.Background(), principalId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.DeleteRoleMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalId** | **string** | Unique id of a principal (typically either a user or service account). | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleMappingRequest struct via the builder pattern


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


## ExportData

> *os.File ExportData(ctx).ForBrowser(forBrowser).Execute()

Export registry data



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
    forBrowser := true // bool | Indicates if the operation is done for a browser.  If true, the response will be a JSON payload with a property called `href`.  This `href` will be a single-use, naked download link suitable for use by a web browser to download the content. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminAPI.ExportData(context.Background()).ForBrowser(forBrowser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.ExportData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportData`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AdminAPI.ExportData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forBrowser** | **bool** | Indicates if the operation is done for a browser.  If true, the response will be a JSON payload with a property called &#x60;href&#x60;.  This &#x60;href&#x60; will be a single-use, naked download link suitable for use by a web browser to download the content. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigProperty

> ConfigurationProperty GetConfigProperty(ctx, propertyName).Execute()

Get configuration property value



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
    propertyName := "propertyName_example" // string | The name of a configuration property.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminAPI.GetConfigProperty(context.Background(), propertyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetConfigProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfigProperty`: ConfigurationProperty
    fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetConfigProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**propertyName** | **string** | The name of a configuration property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigurationProperty**](ConfigurationProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalRuleConfig

> Rule GetGlobalRuleConfig(ctx, rule).Execute()

Get global rule configuration



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
    rule := openapiclient.RuleType("VALIDITY") // RuleType | The unique name/type of a rule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminAPI.GetGlobalRuleConfig(context.Background(), rule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetGlobalRuleConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalRuleConfig`: Rule
    fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetGlobalRuleConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rule** | [**RuleType**](.md) | The unique name/type of a rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalRuleConfigRequest struct via the builder pattern


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


## GetRoleMapping

> RoleMapping GetRoleMapping(ctx, principalId).Execute()

Return a single role mapping



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
    principalId := "principalId_example" // string | Unique id of a principal (typically either a user or service account).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminAPI.GetRoleMapping(context.Background(), principalId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.GetRoleMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleMapping`: RoleMapping
    fmt.Fprintf(os.Stdout, "Response from `AdminAPI.GetRoleMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalId** | **string** | Unique id of a principal (typically either a user or service account). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RoleMapping**](RoleMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportData

> ImportData(ctx).Body(body).XRegistryPreserveGlobalId(xRegistryPreserveGlobalId).XRegistryPreserveContentId(xRegistryPreserveContentId).Execute()

Import registry data



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
    body := os.NewFile(1234, "some_file") // *os.File | The ZIP file representing the previously exported registry data.
    xRegistryPreserveGlobalId := true // bool | If this header is set to false, global ids of imported data will be ignored and replaced by next id in global id sequence. This allows to import any data even thought the global ids would cause a conflict. (optional)
    xRegistryPreserveContentId := true // bool | If this header is set to false, content ids of imported data will be ignored and replaced by next id in content id sequence. The mapping between content and artifacts will be preserved. This allows to import any data even thought the content ids would cause a conflict. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.ImportData(context.Background()).Body(body).XRegistryPreserveGlobalId(xRegistryPreserveGlobalId).XRegistryPreserveContentId(xRegistryPreserveContentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.ImportData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** | The ZIP file representing the previously exported registry data. | 
 **xRegistryPreserveGlobalId** | **bool** | If this header is set to false, global ids of imported data will be ignored and replaced by next id in global id sequence. This allows to import any data even thought the global ids would cause a conflict. | 
 **xRegistryPreserveContentId** | **bool** | If this header is set to false, content ids of imported data will be ignored and replaced by next id in content id sequence. The mapping between content and artifacts will be preserved. This allows to import any data even thought the content ids would cause a conflict. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/zip
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    resp, r, err := apiClient.AdminAPI.ListArtifactTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.ListArtifactTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListArtifactTypes`: []ArtifactTypeInfo
    fmt.Fprintf(os.Stdout, "Response from `AdminAPI.ListArtifactTypes`: %v\n", resp)
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


## ListConfigProperties

> []ConfigurationProperty ListConfigProperties(ctx).Execute()

List all configuration properties



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
    resp, r, err := apiClient.AdminAPI.ListConfigProperties(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.ListConfigProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConfigProperties`: []ConfigurationProperty
    fmt.Fprintf(os.Stdout, "Response from `AdminAPI.ListConfigProperties`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListConfigPropertiesRequest struct via the builder pattern


### Return type

[**[]ConfigurationProperty**](ConfigurationProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGlobalRules

> []RuleType ListGlobalRules(ctx).Execute()

List global rules



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
    resp, r, err := apiClient.AdminAPI.ListGlobalRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.ListGlobalRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGlobalRules`: []RuleType
    fmt.Fprintf(os.Stdout, "Response from `AdminAPI.ListGlobalRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGlobalRulesRequest struct via the builder pattern


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


## ListRoleMappings

> []RoleMapping ListRoleMappings(ctx).Execute()

List all role mappings



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
    resp, r, err := apiClient.AdminAPI.ListRoleMappings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.ListRoleMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoleMappings`: []RoleMapping
    fmt.Fprintf(os.Stdout, "Response from `AdminAPI.ListRoleMappings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRoleMappingsRequest struct via the builder pattern


### Return type

[**[]RoleMapping**](RoleMapping.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetConfigProperty

> ResetConfigProperty(ctx, propertyName).Execute()

Reset a configuration property



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
    propertyName := "propertyName_example" // string | The name of a configuration property.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.ResetConfigProperty(context.Background(), propertyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.ResetConfigProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**propertyName** | **string** | The name of a configuration property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetConfigPropertyRequest struct via the builder pattern


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


## UpdateConfigProperty

> UpdateConfigProperty(ctx, propertyName).UpdateConfigurationProperty(updateConfigurationProperty).Execute()

Update a configuration property



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
    propertyName := "propertyName_example" // string | The name of a configuration property.
    updateConfigurationProperty := *openapiclient.NewUpdateConfigurationProperty("Value_example") // UpdateConfigurationProperty | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.UpdateConfigProperty(context.Background(), propertyName).UpdateConfigurationProperty(updateConfigurationProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.UpdateConfigProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**propertyName** | **string** | The name of a configuration property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateConfigurationProperty** | [**UpdateConfigurationProperty**](UpdateConfigurationProperty.md) |  | 

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


## UpdateGlobalRuleConfig

> Rule UpdateGlobalRuleConfig(ctx, rule).Rule2(rule2).Execute()

Update global rule configuration



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
    rule := openapiclient.RuleType("VALIDITY") // RuleType | The unique name/type of a rule.
    rule2 := *openapiclient.NewRule("Config_example") // Rule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminAPI.UpdateGlobalRuleConfig(context.Background(), rule).Rule2(rule2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.UpdateGlobalRuleConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGlobalRuleConfig`: Rule
    fmt.Fprintf(os.Stdout, "Response from `AdminAPI.UpdateGlobalRuleConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rule** | [**RuleType**](.md) | The unique name/type of a rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGlobalRuleConfigRequest struct via the builder pattern


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


## UpdateRoleMapping

> UpdateRoleMapping(ctx, principalId).UpdateRole(updateRole).Execute()

Update a role mapping



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
    principalId := "principalId_example" // string | Unique id of a principal (typically either a user or service account).
    updateRole := *openapiclient.NewUpdateRole(openapiclient.RoleType("READ_ONLY")) // UpdateRole | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.UpdateRoleMapping(context.Background(), principalId).UpdateRole(updateRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.UpdateRoleMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalId** | **string** | Unique id of a principal (typically either a user or service account). | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRole** | [**UpdateRole**](UpdateRole.md) |  | 

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

