/*
Apicurio Registry API [v3]

Apicurio Registry is a datastore for standard event schemas and API designs. Apicurio Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.  The Apicurio Registry REST API enables client applications to manage the artifacts in the registry. This API provides create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata.   The supported artifact types include: - Apache Avro schema - AsyncAPI specification - Google protocol buffers - GraphQL schema - JSON Schema - Kafka Connect schema - OpenAPI specification - Web Services Description Language - XML Schema Definition   **Important**: The Apicurio Registry REST API is available from `https://MY-REGISTRY-URL/apis/registry/v3` by default. Therefore you must prefix all API operation paths with `../apis/registry/v3` in this case. For example: `../apis/registry/v3/ids/globalIds/{globalId}`. 

API version: 3.0.x
Contact: apicurio@lists.jboss.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registryclient

import (
	"encoding/json"
	"fmt"
)

// checks if the RoleMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleMapping{}

// RoleMapping The mapping between a user/principal and their role.
type RoleMapping struct {
	// 
	PrincipalId string `json:"principalId"`
	Role RoleType `json:"role"`
	// A friendly name for the principal.
	PrincipalName *string `json:"principalName,omitempty"`
}

type _RoleMapping RoleMapping

// NewRoleMapping instantiates a new RoleMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleMapping(principalId string, role RoleType) *RoleMapping {
	this := RoleMapping{}
	this.PrincipalId = principalId
	this.Role = role
	return &this
}

// NewRoleMappingWithDefaults instantiates a new RoleMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleMappingWithDefaults() *RoleMapping {
	this := RoleMapping{}
	return &this
}

// GetPrincipalId returns the PrincipalId field value
func (o *RoleMapping) GetPrincipalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrincipalId
}

// GetPrincipalIdOk returns a tuple with the PrincipalId field value
// and a boolean to check if the value has been set.
func (o *RoleMapping) GetPrincipalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrincipalId, true
}

// SetPrincipalId sets field value
func (o *RoleMapping) SetPrincipalId(v string) {
	o.PrincipalId = v
}

// GetRole returns the Role field value
func (o *RoleMapping) GetRole() RoleType {
	if o == nil {
		var ret RoleType
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *RoleMapping) GetRoleOk() (*RoleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *RoleMapping) SetRole(v RoleType) {
	o.Role = v
}

// GetPrincipalName returns the PrincipalName field value if set, zero value otherwise.
func (o *RoleMapping) GetPrincipalName() string {
	if o == nil || IsNil(o.PrincipalName) {
		var ret string
		return ret
	}
	return *o.PrincipalName
}

// GetPrincipalNameOk returns a tuple with the PrincipalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleMapping) GetPrincipalNameOk() (*string, bool) {
	if o == nil || IsNil(o.PrincipalName) {
		return nil, false
	}
	return o.PrincipalName, true
}

// HasPrincipalName returns a boolean if a field has been set.
func (o *RoleMapping) HasPrincipalName() bool {
	if o != nil && !IsNil(o.PrincipalName) {
		return true
	}

	return false
}

// SetPrincipalName gets a reference to the given string and assigns it to the PrincipalName field.
func (o *RoleMapping) SetPrincipalName(v string) {
	o.PrincipalName = &v
}

func (o RoleMapping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["principalId"] = o.PrincipalId
	toSerialize["role"] = o.Role
	if !IsNil(o.PrincipalName) {
		toSerialize["principalName"] = o.PrincipalName
	}
	return toSerialize, nil
}

func (o *RoleMapping) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"principalId",
		"role",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRoleMapping := _RoleMapping{}

	err = json.Unmarshal(bytes, &varRoleMapping)

	if err != nil {
		return err
	}

	*o = RoleMapping(varRoleMapping)

	return err
}

type NullableRoleMapping struct {
	value *RoleMapping
	isSet bool
}

func (v NullableRoleMapping) Get() *RoleMapping {
	return v.value
}

func (v *NullableRoleMapping) Set(val *RoleMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleMapping(val *RoleMapping) *NullableRoleMapping {
	return &NullableRoleMapping{value: val, isSet: true}
}

func (v NullableRoleMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


