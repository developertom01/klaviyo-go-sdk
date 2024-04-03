/*
Klaviyo API

The Klaviyo REST API. Please visit https://developers.klaviyo.com for more details.

API version: 2024-02-15
Contact: developers@klaviyo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package klaviyo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GetTagFlowRelationshipsResponseCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTagFlowRelationshipsResponseCollection{}

// GetTagFlowRelationshipsResponseCollection struct for GetTagFlowRelationshipsResponseCollection
type GetTagFlowRelationshipsResponseCollection struct {
	Data []GetTagFlowRelationshipsResponseCollectionDataInner `json:"data"`
}

type _GetTagFlowRelationshipsResponseCollection GetTagFlowRelationshipsResponseCollection

// NewGetTagFlowRelationshipsResponseCollection instantiates a new GetTagFlowRelationshipsResponseCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTagFlowRelationshipsResponseCollection(data []GetTagFlowRelationshipsResponseCollectionDataInner) *GetTagFlowRelationshipsResponseCollection {
	this := GetTagFlowRelationshipsResponseCollection{}
	this.Data = data
	return &this
}

// NewGetTagFlowRelationshipsResponseCollectionWithDefaults instantiates a new GetTagFlowRelationshipsResponseCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTagFlowRelationshipsResponseCollectionWithDefaults() *GetTagFlowRelationshipsResponseCollection {
	this := GetTagFlowRelationshipsResponseCollection{}
	return &this
}

// GetData returns the Data field value
func (o *GetTagFlowRelationshipsResponseCollection) GetData() []GetTagFlowRelationshipsResponseCollectionDataInner {
	if o == nil {
		var ret []GetTagFlowRelationshipsResponseCollectionDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetTagFlowRelationshipsResponseCollection) GetDataOk() ([]GetTagFlowRelationshipsResponseCollectionDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetTagFlowRelationshipsResponseCollection) SetData(v []GetTagFlowRelationshipsResponseCollectionDataInner) {
	o.Data = v
}

func (o GetTagFlowRelationshipsResponseCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTagFlowRelationshipsResponseCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetTagFlowRelationshipsResponseCollection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetTagFlowRelationshipsResponseCollection := _GetTagFlowRelationshipsResponseCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTagFlowRelationshipsResponseCollection)

	if err != nil {
		return err
	}

	*o = GetTagFlowRelationshipsResponseCollection(varGetTagFlowRelationshipsResponseCollection)

	return err
}

type NullableGetTagFlowRelationshipsResponseCollection struct {
	value *GetTagFlowRelationshipsResponseCollection
	isSet bool
}

func (v NullableGetTagFlowRelationshipsResponseCollection) Get() *GetTagFlowRelationshipsResponseCollection {
	return v.value
}

func (v *NullableGetTagFlowRelationshipsResponseCollection) Set(val *GetTagFlowRelationshipsResponseCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTagFlowRelationshipsResponseCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTagFlowRelationshipsResponseCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTagFlowRelationshipsResponseCollection(val *GetTagFlowRelationshipsResponseCollection) *NullableGetTagFlowRelationshipsResponseCollection {
	return &NullableGetTagFlowRelationshipsResponseCollection{value: val, isSet: true}
}

func (v NullableGetTagFlowRelationshipsResponseCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTagFlowRelationshipsResponseCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


