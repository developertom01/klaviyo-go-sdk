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

// checks if the GetCampaignMessagesRelationshipListResponseCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCampaignMessagesRelationshipListResponseCollection{}

// GetCampaignMessagesRelationshipListResponseCollection struct for GetCampaignMessagesRelationshipListResponseCollection
type GetCampaignMessagesRelationshipListResponseCollection struct {
	Data []GetCampaignMessagesRelationshipListResponseCollectionDataInner `json:"data"`
}

type _GetCampaignMessagesRelationshipListResponseCollection GetCampaignMessagesRelationshipListResponseCollection

// NewGetCampaignMessagesRelationshipListResponseCollection instantiates a new GetCampaignMessagesRelationshipListResponseCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCampaignMessagesRelationshipListResponseCollection(data []GetCampaignMessagesRelationshipListResponseCollectionDataInner) *GetCampaignMessagesRelationshipListResponseCollection {
	this := GetCampaignMessagesRelationshipListResponseCollection{}
	this.Data = data
	return &this
}

// NewGetCampaignMessagesRelationshipListResponseCollectionWithDefaults instantiates a new GetCampaignMessagesRelationshipListResponseCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCampaignMessagesRelationshipListResponseCollectionWithDefaults() *GetCampaignMessagesRelationshipListResponseCollection {
	this := GetCampaignMessagesRelationshipListResponseCollection{}
	return &this
}

// GetData returns the Data field value
func (o *GetCampaignMessagesRelationshipListResponseCollection) GetData() []GetCampaignMessagesRelationshipListResponseCollectionDataInner {
	if o == nil {
		var ret []GetCampaignMessagesRelationshipListResponseCollectionDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetCampaignMessagesRelationshipListResponseCollection) GetDataOk() ([]GetCampaignMessagesRelationshipListResponseCollectionDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetCampaignMessagesRelationshipListResponseCollection) SetData(v []GetCampaignMessagesRelationshipListResponseCollectionDataInner) {
	o.Data = v
}

func (o GetCampaignMessagesRelationshipListResponseCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCampaignMessagesRelationshipListResponseCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetCampaignMessagesRelationshipListResponseCollection) UnmarshalJSON(data []byte) (err error) {
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

	varGetCampaignMessagesRelationshipListResponseCollection := _GetCampaignMessagesRelationshipListResponseCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCampaignMessagesRelationshipListResponseCollection)

	if err != nil {
		return err
	}

	*o = GetCampaignMessagesRelationshipListResponseCollection(varGetCampaignMessagesRelationshipListResponseCollection)

	return err
}

type NullableGetCampaignMessagesRelationshipListResponseCollection struct {
	value *GetCampaignMessagesRelationshipListResponseCollection
	isSet bool
}

func (v NullableGetCampaignMessagesRelationshipListResponseCollection) Get() *GetCampaignMessagesRelationshipListResponseCollection {
	return v.value
}

func (v *NullableGetCampaignMessagesRelationshipListResponseCollection) Set(val *GetCampaignMessagesRelationshipListResponseCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCampaignMessagesRelationshipListResponseCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCampaignMessagesRelationshipListResponseCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCampaignMessagesRelationshipListResponseCollection(val *GetCampaignMessagesRelationshipListResponseCollection) *NullableGetCampaignMessagesRelationshipListResponseCollection {
	return &NullableGetCampaignMessagesRelationshipListResponseCollection{value: val, isSet: true}
}

func (v NullableGetCampaignMessagesRelationshipListResponseCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCampaignMessagesRelationshipListResponseCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

