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

// checks if the TagCampaignOp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagCampaignOp{}

// TagCampaignOp struct for TagCampaignOp
type TagCampaignOp struct {
	Data []TagCampaignOpDataInner `json:"data"`
}

type _TagCampaignOp TagCampaignOp

// NewTagCampaignOp instantiates a new TagCampaignOp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagCampaignOp(data []TagCampaignOpDataInner) *TagCampaignOp {
	this := TagCampaignOp{}
	this.Data = data
	return &this
}

// NewTagCampaignOpWithDefaults instantiates a new TagCampaignOp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagCampaignOpWithDefaults() *TagCampaignOp {
	this := TagCampaignOp{}
	return &this
}

// GetData returns the Data field value
func (o *TagCampaignOp) GetData() []TagCampaignOpDataInner {
	if o == nil {
		var ret []TagCampaignOpDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TagCampaignOp) GetDataOk() ([]TagCampaignOpDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *TagCampaignOp) SetData(v []TagCampaignOpDataInner) {
	o.Data = v
}

func (o TagCampaignOp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagCampaignOp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *TagCampaignOp) UnmarshalJSON(data []byte) (err error) {
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

	varTagCampaignOp := _TagCampaignOp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTagCampaignOp)

	if err != nil {
		return err
	}

	*o = TagCampaignOp(varTagCampaignOp)

	return err
}

type NullableTagCampaignOp struct {
	value *TagCampaignOp
	isSet bool
}

func (v NullableTagCampaignOp) Get() *TagCampaignOp {
	return v.value
}

func (v *NullableTagCampaignOp) Set(val *TagCampaignOp) {
	v.value = val
	v.isSet = true
}

func (v NullableTagCampaignOp) IsSet() bool {
	return v.isSet
}

func (v *NullableTagCampaignOp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagCampaignOp(val *TagCampaignOp) *NullableTagCampaignOp {
	return &NullableTagCampaignOp{value: val, isSet: true}
}

func (v NullableTagCampaignOp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagCampaignOp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


