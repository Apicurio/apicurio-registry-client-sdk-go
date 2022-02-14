/*
Apicurio Registry API [v2]

Apicurio Registry is a datastore for standard event schemas and API designs. Apicurio Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.  The Apicurio Registry REST API enables client applications to manage the artifacts in the registry. This API provides create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata.   The supported artifact types include: - Apache Avro schema - AsyncAPI specification - Google protocol buffers - GraphQL schema - JSON Schema - Kafka Connect schema - OpenAPI specification - Web Services Description Language - XML Schema Definition   **Important**: The Apicurio Registry REST API is available from `https://MY-REGISTRY-URL/apis/registry/v2` by default. Therefore you must prefix all API operation paths with `../apis/registry/v2` in this case. For example: `../apis/registry/v2/ids/globalIds/{globalId}`. 

API version: 2.2.0.Final
Contact: apicurio@lists.jboss.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registryclient

import (
	"encoding/json"
	"fmt"
)

// SortOrder the model 'SortOrder'
type SortOrder string

// List of SortOrder
const (
	ASC SortOrder = "asc"
	DESC SortOrder = "desc"
)

// All allowed values of SortOrder enum
var AllowedSortOrderEnumValues = []SortOrder{
	"asc",
	"desc",
}

func (v *SortOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SortOrder(value)
	for _, existing := range AllowedSortOrderEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SortOrder", value)
}

// NewSortOrderFromValue returns a pointer to a valid SortOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSortOrderFromValue(v string) (*SortOrder, error) {
	ev := SortOrder(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SortOrder: valid values are %v", v, AllowedSortOrderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SortOrder) IsValid() bool {
	for _, existing := range AllowedSortOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SortOrder value
func (v SortOrder) Ptr() *SortOrder {
	return &v
}

type NullableSortOrder struct {
	value *SortOrder
	isSet bool
}

func (v NullableSortOrder) Get() *SortOrder {
	return v.value
}

func (v *NullableSortOrder) Set(val *SortOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableSortOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableSortOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSortOrder(val *SortOrder) *NullableSortOrder {
	return &NullableSortOrder{value: val, isSet: true}
}

func (v NullableSortOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSortOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

