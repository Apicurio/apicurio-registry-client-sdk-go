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

// IfExists 
type IfExists string

// List of IfExists
const (
	FAIL IfExists = "FAIL"
	UPDATE IfExists = "UPDATE"
	RETURN IfExists = "RETURN"
	RETURN_OR_UPDATE IfExists = "RETURN_OR_UPDATE"
)

// All allowed values of IfExists enum
var AllowedIfExistsEnumValues = []IfExists{
	"FAIL",
	"UPDATE",
	"RETURN",
	"RETURN_OR_UPDATE",
}

func (v *IfExists) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IfExists(value)
	for _, existing := range AllowedIfExistsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IfExists", value)
}

// NewIfExistsFromValue returns a pointer to a valid IfExists
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIfExistsFromValue(v string) (*IfExists, error) {
	ev := IfExists(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IfExists: valid values are %v", v, AllowedIfExistsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IfExists) IsValid() bool {
	for _, existing := range AllowedIfExistsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IfExists value
func (v IfExists) Ptr() *IfExists {
	return &v
}

type NullableIfExists struct {
	value *IfExists
	isSet bool
}

func (v NullableIfExists) Get() *IfExists {
	return v.value
}

func (v *NullableIfExists) Set(val *IfExists) {
	v.value = val
	v.isSet = true
}

func (v NullableIfExists) IsSet() bool {
	return v.isSet
}

func (v *NullableIfExists) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIfExists(val *IfExists) *NullableIfExists {
	return &NullableIfExists{value: val, isSet: true}
}

func (v NullableIfExists) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIfExists) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

