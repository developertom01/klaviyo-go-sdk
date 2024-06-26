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

// checks if the GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData{}

// GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData struct for GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData
type GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData struct {
	Type TagGroupEnum `json:"type"`
	Id string `json:"id"`
}

type _GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData

// NewGetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData instantiates a new GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData(type_ TagGroupEnum, id string) *GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData {
	this := GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupDataWithDefaults instantiates a new GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupDataWithDefaults() *GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData {
	this := GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData{}
	return &this
}

// GetType returns the Type field value
func (o *GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData) GetType() TagGroupEnum {
	if o == nil {
		var ret TagGroupEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData) GetTypeOk() (*TagGroupEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData) SetType(v TagGroupEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData) SetId(v string) {
	o.Id = v
}

func (o GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varGetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData := _GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData)

	if err != nil {
		return err
	}

	*o = GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData(varGetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData)

	return err
}

type NullableGetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData struct {
	value *GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData
	isSet bool
}

func (v NullableGetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData) Get() *GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData {
	return v.value
}

func (v *NullableGetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData) Set(val *GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData(val *GetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData) *NullableGetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData {
	return &NullableGetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData{value: val, isSet: true}
}

func (v NullableGetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTagResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsTagGroupData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


