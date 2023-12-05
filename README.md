# Apicurio Registry Client SDK for Go

Apicurio Registry is a datastore for standard event schemas and API designs. Apicurio Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.

The Apicurio Registry REST API enables client applications to manage the artifacts in the registry. This API provides create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata. 

The supported artifact types include:
- Apache Avro schema
- AsyncAPI specification
- Google protocol buffers
- GraphQL schema
- JSON Schema
- Kafka Connect schema
- OpenAPI specification
- Web Services Description Language
- XML Schema Definition


**Important**: The Apicurio Registry REST API is available from `https://MY-REGISTRY-URL/apis/registry/v3` by default. Therefore you must prefix all API operation paths with `../apis/registry/v3` in this case. For example: `../apis/registry/v3/ids/globalIds/{globalId}`.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.0.x
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://github.com/apicurio/apicurio-registry](https://github.com/apicurio/apicurio-registry)

## Installation

Install the client SDK library to your Go project:

```shell
go get github.com/Apicurio/apicurio-registry-client-sdk-go
```

Add the following import to use it:

```golang
import "github.com/Apicurio/apicurio-registry-client-sdk-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Creating an API client

To create an API client using the default configuration options:

```golang
cfg := registryclient.NewConfiguration()
registryClient := registryclient.NewAPIClient(&cfg)
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), registryclient.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), registryclient.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), registryclient.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), registryclient.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AdminAPI* | [**CreateGlobalRule**](docs/AdminAPI.md#createglobalrule) | **Post** /admin/rules | Create global rule
*AdminAPI* | [**CreateRoleMapping**](docs/AdminAPI.md#createrolemapping) | **Post** /admin/roleMappings | Create a new role mapping
*AdminAPI* | [**DeleteAllGlobalRules**](docs/AdminAPI.md#deleteallglobalrules) | **Delete** /admin/rules | Delete all global rules
*AdminAPI* | [**DeleteGlobalRule**](docs/AdminAPI.md#deleteglobalrule) | **Delete** /admin/rules/{rule} | Delete global rule
*AdminAPI* | [**DeleteRoleMapping**](docs/AdminAPI.md#deleterolemapping) | **Delete** /admin/roleMappings/{principalId} | Delete a role mapping
*AdminAPI* | [**ExportData**](docs/AdminAPI.md#exportdata) | **Get** /admin/export | Export registry data
*AdminAPI* | [**GetConfigProperty**](docs/AdminAPI.md#getconfigproperty) | **Get** /admin/config/properties/{propertyName} | Get configuration property value
*AdminAPI* | [**GetGlobalRuleConfig**](docs/AdminAPI.md#getglobalruleconfig) | **Get** /admin/rules/{rule} | Get global rule configuration
*AdminAPI* | [**GetRoleMapping**](docs/AdminAPI.md#getrolemapping) | **Get** /admin/roleMappings/{principalId} | Return a single role mapping
*AdminAPI* | [**ImportData**](docs/AdminAPI.md#importdata) | **Post** /admin/import | Import registry data
*AdminAPI* | [**ListArtifactTypes**](docs/AdminAPI.md#listartifacttypes) | **Get** /admin/artifactTypes | List artifact types
*AdminAPI* | [**ListConfigProperties**](docs/AdminAPI.md#listconfigproperties) | **Get** /admin/config/properties | List all configuration properties
*AdminAPI* | [**ListGlobalRules**](docs/AdminAPI.md#listglobalrules) | **Get** /admin/rules | List global rules
*AdminAPI* | [**ListRoleMappings**](docs/AdminAPI.md#listrolemappings) | **Get** /admin/roleMappings | List all role mappings
*AdminAPI* | [**ResetConfigProperty**](docs/AdminAPI.md#resetconfigproperty) | **Delete** /admin/config/properties/{propertyName} | Reset a configuration property
*AdminAPI* | [**UpdateConfigProperty**](docs/AdminAPI.md#updateconfigproperty) | **Put** /admin/config/properties/{propertyName} | Update a configuration property
*AdminAPI* | [**UpdateGlobalRuleConfig**](docs/AdminAPI.md#updateglobalruleconfig) | **Put** /admin/rules/{rule} | Update global rule configuration
*AdminAPI* | [**UpdateRoleMapping**](docs/AdminAPI.md#updaterolemapping) | **Put** /admin/roleMappings/{principalId} | Update a role mapping
*ArtifactRulesAPI* | [**CreateArtifactRule**](docs/ArtifactRulesAPI.md#createartifactrule) | **Post** /groups/{groupId}/artifacts/{artifactId}/rules | Create artifact rule
*ArtifactRulesAPI* | [**DeleteArtifactRule**](docs/ArtifactRulesAPI.md#deleteartifactrule) | **Delete** /groups/{groupId}/artifacts/{artifactId}/rules/{rule} | Delete artifact rule
*ArtifactRulesAPI* | [**DeleteArtifactRules**](docs/ArtifactRulesAPI.md#deleteartifactrules) | **Delete** /groups/{groupId}/artifacts/{artifactId}/rules | Delete artifact rules
*ArtifactRulesAPI* | [**GetArtifactRuleConfig**](docs/ArtifactRulesAPI.md#getartifactruleconfig) | **Get** /groups/{groupId}/artifacts/{artifactId}/rules/{rule} | Get artifact rule configuration
*ArtifactRulesAPI* | [**ListArtifactRules**](docs/ArtifactRulesAPI.md#listartifactrules) | **Get** /groups/{groupId}/artifacts/{artifactId}/rules | List artifact rules
*ArtifactRulesAPI* | [**TestUpdateArtifact**](docs/ArtifactRulesAPI.md#testupdateartifact) | **Put** /groups/{groupId}/artifacts/{artifactId}/test | Test update artifact
*ArtifactRulesAPI* | [**UpdateArtifactRuleConfig**](docs/ArtifactRulesAPI.md#updateartifactruleconfig) | **Put** /groups/{groupId}/artifacts/{artifactId}/rules/{rule} | Update artifact rule configuration
*ArtifactTypeAPI* | [**ListArtifactTypes**](docs/ArtifactTypeAPI.md#listartifacttypes) | **Get** /admin/artifactTypes | List artifact types
*ArtifactsAPI* | [**CreateArtifact**](docs/ArtifactsAPI.md#createartifact) | **Post** /groups/{groupId}/artifacts | Create artifact
*ArtifactsAPI* | [**DeleteArtifact**](docs/ArtifactsAPI.md#deleteartifact) | **Delete** /groups/{groupId}/artifacts/{artifactId} | Delete artifact
*ArtifactsAPI* | [**DeleteArtifactsInGroup**](docs/ArtifactsAPI.md#deleteartifactsingroup) | **Delete** /groups/{groupId}/artifacts | Delete artifacts in group
*ArtifactsAPI* | [**GetContentByGlobalId**](docs/ArtifactsAPI.md#getcontentbyglobalid) | **Get** /ids/globalIds/{globalId} | Get artifact by global ID
*ArtifactsAPI* | [**GetContentByHash**](docs/ArtifactsAPI.md#getcontentbyhash) | **Get** /ids/contentHashes/{contentHash}/ | Get artifact content by SHA-256 hash
*ArtifactsAPI* | [**GetContentById**](docs/ArtifactsAPI.md#getcontentbyid) | **Get** /ids/contentIds/{contentId}/ | Get artifact content by ID
*ArtifactsAPI* | [**GetLatestArtifact**](docs/ArtifactsAPI.md#getlatestartifact) | **Get** /groups/{groupId}/artifacts/{artifactId} | Get latest artifact
*ArtifactsAPI* | [**ListArtifactsInGroup**](docs/ArtifactsAPI.md#listartifactsingroup) | **Get** /groups/{groupId}/artifacts | List artifacts in group
*ArtifactsAPI* | [**ReferencesByContentHash**](docs/ArtifactsAPI.md#referencesbycontenthash) | **Get** /ids/contentHashes/{contentHash}/references | List artifact references by hash
*ArtifactsAPI* | [**ReferencesByContentId**](docs/ArtifactsAPI.md#referencesbycontentid) | **Get** /ids/contentIds/{contentId}/references | List artifact references by content ID
*ArtifactsAPI* | [**ReferencesByGlobalId**](docs/ArtifactsAPI.md#referencesbyglobalid) | **Get** /ids/globalIds/{globalId}/references | List artifact references by global ID
*ArtifactsAPI* | [**SearchArtifacts**](docs/ArtifactsAPI.md#searchartifacts) | **Get** /search/artifacts | Search for artifacts
*ArtifactsAPI* | [**SearchArtifactsByContent**](docs/ArtifactsAPI.md#searchartifactsbycontent) | **Post** /search/artifacts | Search for artifacts by content
*ArtifactsAPI* | [**UpdateArtifact**](docs/ArtifactsAPI.md#updateartifact) | **Put** /groups/{groupId}/artifacts/{artifactId} | Update artifact
*ArtifactsAPI* | [**UpdateArtifactState**](docs/ArtifactsAPI.md#updateartifactstate) | **Put** /groups/{groupId}/artifacts/{artifactId}/state | Update artifact state
*GroupsAPI* | [**CreateGroup**](docs/GroupsAPI.md#creategroup) | **Post** /groups | Create a new group
*GroupsAPI* | [**DeleteGroupById**](docs/GroupsAPI.md#deletegroupbyid) | **Delete** /groups/{groupId} | Delete a group by the specified ID.
*GroupsAPI* | [**GetGroupById**](docs/GroupsAPI.md#getgroupbyid) | **Get** /groups/{groupId} | Get a group by the specified ID.
*GroupsAPI* | [**ListGroups**](docs/GroupsAPI.md#listgroups) | **Get** /groups | List groups
*MetadataAPI* | [**DeleteArtifactVersionMetaData**](docs/MetadataAPI.md#deleteartifactversionmetadata) | **Delete** /groups/{groupId}/artifacts/{artifactId}/versions/{version}/meta | Delete artifact version metadata
*MetadataAPI* | [**GetArtifactMetaData**](docs/MetadataAPI.md#getartifactmetadata) | **Get** /groups/{groupId}/artifacts/{artifactId}/meta | Get artifact metadata
*MetadataAPI* | [**GetArtifactOwner**](docs/MetadataAPI.md#getartifactowner) | **Get** /groups/{groupId}/artifacts/{artifactId}/owner | Get artifact owner
*MetadataAPI* | [**GetArtifactVersionMetaData**](docs/MetadataAPI.md#getartifactversionmetadata) | **Get** /groups/{groupId}/artifacts/{artifactId}/versions/{version}/meta | Get artifact version metadata
*MetadataAPI* | [**GetArtifactVersionMetaDataByContent**](docs/MetadataAPI.md#getartifactversionmetadatabycontent) | **Post** /groups/{groupId}/artifacts/{artifactId}/meta | Get artifact version metadata by content
*MetadataAPI* | [**UpdateArtifactMetaData**](docs/MetadataAPI.md#updateartifactmetadata) | **Put** /groups/{groupId}/artifacts/{artifactId}/meta | Update artifact metadata
*MetadataAPI* | [**UpdateArtifactOwner**](docs/MetadataAPI.md#updateartifactowner) | **Put** /groups/{groupId}/artifacts/{artifactId}/owner | Update artifact owner
*MetadataAPI* | [**UpdateArtifactVersionMetaData**](docs/MetadataAPI.md#updateartifactversionmetadata) | **Put** /groups/{groupId}/artifacts/{artifactId}/versions/{version}/meta | Update artifact version metadata
*SystemAPI* | [**GetResourceLimits**](docs/SystemAPI.md#getresourcelimits) | **Get** /system/limits | Get resource limits information
*SystemAPI* | [**GetSystemInfo**](docs/SystemAPI.md#getsysteminfo) | **Get** /system/info | Get system information
*UsersAPI* | [**GetCurrentUserInfo**](docs/UsersAPI.md#getcurrentuserinfo) | **Get** /users/me | Get current user
*VersionsAPI* | [**AddArtifactVersionComment**](docs/VersionsAPI.md#addartifactversioncomment) | **Post** /groups/{groupId}/artifacts/{artifactId}/versions/{version}/comments | Add new comment
*VersionsAPI* | [**CreateArtifactVersion**](docs/VersionsAPI.md#createartifactversion) | **Post** /groups/{groupId}/artifacts/{artifactId}/versions | Create artifact version
*VersionsAPI* | [**DeleteArtifactVersion**](docs/VersionsAPI.md#deleteartifactversion) | **Delete** /groups/{groupId}/artifacts/{artifactId}/versions/{version} | Delete artifact version
*VersionsAPI* | [**DeleteArtifactVersionComment**](docs/VersionsAPI.md#deleteartifactversioncomment) | **Delete** /groups/{groupId}/artifacts/{artifactId}/versions/{version}/comments/{commentId} | Delete a single comment
*VersionsAPI* | [**GetArtifactVersion**](docs/VersionsAPI.md#getartifactversion) | **Get** /groups/{groupId}/artifacts/{artifactId}/versions/{version} | Get artifact version
*VersionsAPI* | [**GetArtifactVersionComments**](docs/VersionsAPI.md#getartifactversioncomments) | **Get** /groups/{groupId}/artifacts/{artifactId}/versions/{version}/comments | Get artifact version comments
*VersionsAPI* | [**GetArtifactVersionReferences**](docs/VersionsAPI.md#getartifactversionreferences) | **Get** /groups/{groupId}/artifacts/{artifactId}/versions/{version}/references | Get artifact version references
*VersionsAPI* | [**ListArtifactVersions**](docs/VersionsAPI.md#listartifactversions) | **Get** /groups/{groupId}/artifacts/{artifactId}/versions | List artifact versions
*VersionsAPI* | [**UpdateArtifactVersionComment**](docs/VersionsAPI.md#updateartifactversioncomment) | **Put** /groups/{groupId}/artifacts/{artifactId}/versions/{version}/comments/{commentId} | Update a comment
*VersionsAPI* | [**UpdateArtifactVersionState**](docs/VersionsAPI.md#updateartifactversionstate) | **Put** /groups/{groupId}/artifacts/{artifactId}/versions/{version}/state | Update artifact version state


## Documentation For Models

 - [ArtifactContent](docs/ArtifactContent.md)
 - [ArtifactMetaData](docs/ArtifactMetaData.md)
 - [ArtifactOwner](docs/ArtifactOwner.md)
 - [ArtifactReference](docs/ArtifactReference.md)
 - [ArtifactSearchResults](docs/ArtifactSearchResults.md)
 - [ArtifactState](docs/ArtifactState.md)
 - [ArtifactTypeInfo](docs/ArtifactTypeInfo.md)
 - [Comment](docs/Comment.md)
 - [ConfigurationProperty](docs/ConfigurationProperty.md)
 - [CreateGroupMetaData](docs/CreateGroupMetaData.md)
 - [DownloadRef](docs/DownloadRef.md)
 - [EditableMetaData](docs/EditableMetaData.md)
 - [Error](docs/Error.md)
 - [GroupMetaData](docs/GroupMetaData.md)
 - [GroupSearchResults](docs/GroupSearchResults.md)
 - [HandleReferencesType](docs/HandleReferencesType.md)
 - [IfExists](docs/IfExists.md)
 - [Limits](docs/Limits.md)
 - [NewComment](docs/NewComment.md)
 - [ReferenceType](docs/ReferenceType.md)
 - [RoleMapping](docs/RoleMapping.md)
 - [RoleType](docs/RoleType.md)
 - [Rule](docs/Rule.md)
 - [RuleType](docs/RuleType.md)
 - [RuleViolationCause](docs/RuleViolationCause.md)
 - [RuleViolationError](docs/RuleViolationError.md)
 - [SearchedArtifact](docs/SearchedArtifact.md)
 - [SearchedGroup](docs/SearchedGroup.md)
 - [SearchedVersion](docs/SearchedVersion.md)
 - [SortBy](docs/SortBy.md)
 - [SortOrder](docs/SortOrder.md)
 - [SystemInfo](docs/SystemInfo.md)
 - [UpdateConfigurationProperty](docs/UpdateConfigurationProperty.md)
 - [UpdateRole](docs/UpdateRole.md)
 - [UpdateState](docs/UpdateState.md)
 - [UserInfo](docs/UserInfo.md)
 - [VersionMetaData](docs/VersionMetaData.md)
 - [VersionSearchResults](docs/VersionSearchResults.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

apicurio@lists.jboss.org
