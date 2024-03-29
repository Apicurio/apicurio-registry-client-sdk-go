/*
Apicurio Registry API [v2]

Apicurio Registry is a datastore for standard event schemas and API designs. Apicurio Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.  The Apicurio Registry REST API enables client applications to manage the artifacts in the registry. This API provides create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata.   The supported artifact types include: - Apache Avro schema - AsyncAPI specification - Google protocol buffers - GraphQL schema - JSON Schema - Kafka Connect schema - OpenAPI specification - Web Services Description Language - XML Schema Definition   **Important**: The Apicurio Registry REST API is available from `https://MY-REGISTRY-URL/apis/registry/v2` by default. Therefore you must prefix all API operation paths with `../apis/registry/v2` in this case. For example: `../apis/registry/v2/ids/globalIds/{globalId}`. 

API version: 2.4.x
Contact: apicurio@lists.jboss.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registryclient

import (
	"encoding/json"
)

// checks if the DownloadRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DownloadRef{}

// DownloadRef Models a download \"link\".  Useful for browser use-cases.
type DownloadRef struct {
	DownloadId string `json:"downloadId"`
	Href *string `json:"href,omitempty"`
}

// NewDownloadRef instantiates a new DownloadRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDownloadRef(downloadId string) *DownloadRef {
	this := DownloadRef{}
	this.DownloadId = downloadId
	return &this
}

// NewDownloadRefWithDefaults instantiates a new DownloadRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDownloadRefWithDefaults() *DownloadRef {
	this := DownloadRef{}
	return &this
}

// GetDownloadId returns the DownloadId field value
func (o *DownloadRef) GetDownloadId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DownloadId
}

// GetDownloadIdOk returns a tuple with the DownloadId field value
// and a boolean to check if the value has been set.
func (o *DownloadRef) GetDownloadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DownloadId, true
}

// SetDownloadId sets field value
func (o *DownloadRef) SetDownloadId(v string) {
	o.DownloadId = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *DownloadRef) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DownloadRef) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *DownloadRef) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *DownloadRef) SetHref(v string) {
	o.Href = &v
}

func (o DownloadRef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DownloadRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["downloadId"] = o.DownloadId
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	return toSerialize, nil
}

type NullableDownloadRef struct {
	value *DownloadRef
	isSet bool
}

func (v NullableDownloadRef) Get() *DownloadRef {
	return v.value
}

func (v *NullableDownloadRef) Set(val *DownloadRef) {
	v.value = val
	v.isSet = true
}

func (v NullableDownloadRef) IsSet() bool {
	return v.isSet
}

func (v *NullableDownloadRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDownloadRef(val *DownloadRef) *NullableDownloadRef {
	return &NullableDownloadRef{value: val, isSet: true}
}

func (v NullableDownloadRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDownloadRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


