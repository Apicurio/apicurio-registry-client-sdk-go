# ArtifactContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | Raw content of the artifact or a valid (and accessible) URL where the content can be found. | 
**References** | [**[]ArtifactReference**](ArtifactReference.md) | Collection of references to other artifacts. | 

## Methods

### NewArtifactContent

`func NewArtifactContent(content string, references []ArtifactReference, ) *ArtifactContent`

NewArtifactContent instantiates a new ArtifactContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactContentWithDefaults

`func NewArtifactContentWithDefaults() *ArtifactContent`

NewArtifactContentWithDefaults instantiates a new ArtifactContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ArtifactContent) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ArtifactContent) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ArtifactContent) SetContent(v string)`

SetContent sets Content field to given value.


### GetReferences

`func (o *ArtifactContent) GetReferences() []ArtifactReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ArtifactContent) GetReferencesOk() (*[]ArtifactReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ArtifactContent) SetReferences(v []ArtifactReference)`

SetReferences sets References field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


