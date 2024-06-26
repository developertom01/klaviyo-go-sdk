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

// checks if the GetProfileListRelationshipsResponseCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProfileListRelationshipsResponseCollection{}

// GetProfileListRelationshipsResponseCollection struct for GetProfileListRelationshipsResponseCollection
type GetProfileListRelationshipsResponseCollection struct {
	Data []GetProfileResponseCompoundDocumentDataAllOfRelationshipsListsDataInner `json:"data"`
}

type _GetProfileListRelationshipsResponseCollection GetProfileListRelationshipsResponseCollection

// NewGetProfileListRelationshipsResponseCollection instantiates a new GetProfileListRelationshipsResponseCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProfileListRelationshipsResponseCollection(data []GetProfileResponseCompoundDocumentDataAllOfRelationshipsListsDataInner) *GetProfileListRelationshipsResponseCollection {
	this := GetProfileListRelationshipsResponseCollection{}
	this.Data = data
	return &this
}

// NewGetProfileListRelationshipsResponseCollectionWithDefaults instantiates a new GetProfileListRelationshipsResponseCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProfileListRelationshipsResponseCollectionWithDefaults() *GetProfileListRelationshipsResponseCollection {
	this := GetProfileListRelationshipsResponseCollection{}
	return &this
}

// GetData returns the Data field value
func (o *GetProfileListRelationshipsResponseCollection) GetData() []GetProfileResponseCompoundDocumentDataAllOfRelationshipsListsDataInner {
	if o == nil {
		var ret []GetProfileResponseCompoundDocumentDataAllOfRelationshipsListsDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetProfileListRelationshipsResponseCollection) GetDataOk() ([]GetProfileResponseCompoundDocumentDataAllOfRelationshipsListsDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetProfileListRelationshipsResponseCollection) SetData(v []GetProfileResponseCompoundDocumentDataAllOfRelationshipsListsDataInner) {
	o.Data = v
}

func (o GetProfileListRelationshipsResponseCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProfileListRelationshipsResponseCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetProfileListRelationshipsResponseCollection) UnmarshalJSON(data []byte) (err error) {
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

	varGetProfileListRelationshipsResponseCollection := _GetProfileListRelationshipsResponseCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetProfileListRelationshipsResponseCollection)

	if err != nil {
		return err
	}

	*o = GetProfileListRelationshipsResponseCollection(varGetProfileListRelationshipsResponseCollection)

	return err
}

type NullableGetProfileListRelationshipsResponseCollection struct {
	value *GetProfileListRelationshipsResponseCollection
	isSet bool
}

func (v NullableGetProfileListRelationshipsResponseCollection) Get() *GetProfileListRelationshipsResponseCollection {
	return v.value
}

func (v *NullableGetProfileListRelationshipsResponseCollection) Set(val *GetProfileListRelationshipsResponseCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProfileListRelationshipsResponseCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProfileListRelationshipsResponseCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProfileListRelationshipsResponseCollection(val *GetProfileListRelationshipsResponseCollection) *NullableGetProfileListRelationshipsResponseCollection {
	return &NullableGetProfileListRelationshipsResponseCollection{value: val, isSet: true}
}

func (v NullableGetProfileListRelationshipsResponseCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProfileListRelationshipsResponseCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


