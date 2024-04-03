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

// checks if the GetSegmentMemberResponseCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSegmentMemberResponseCollection{}

// GetSegmentMemberResponseCollection struct for GetSegmentMemberResponseCollection
type GetSegmentMemberResponseCollection struct {
	Data []GetSegmentMemberResponseCollectionDataInner `json:"data"`
	Links CollectionLinks `json:"links"`
}

type _GetSegmentMemberResponseCollection GetSegmentMemberResponseCollection

// NewGetSegmentMemberResponseCollection instantiates a new GetSegmentMemberResponseCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSegmentMemberResponseCollection(data []GetSegmentMemberResponseCollectionDataInner, links CollectionLinks) *GetSegmentMemberResponseCollection {
	this := GetSegmentMemberResponseCollection{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGetSegmentMemberResponseCollectionWithDefaults instantiates a new GetSegmentMemberResponseCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSegmentMemberResponseCollectionWithDefaults() *GetSegmentMemberResponseCollection {
	this := GetSegmentMemberResponseCollection{}
	return &this
}

// GetData returns the Data field value
func (o *GetSegmentMemberResponseCollection) GetData() []GetSegmentMemberResponseCollectionDataInner {
	if o == nil {
		var ret []GetSegmentMemberResponseCollectionDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetSegmentMemberResponseCollection) GetDataOk() ([]GetSegmentMemberResponseCollectionDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetSegmentMemberResponseCollection) SetData(v []GetSegmentMemberResponseCollectionDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *GetSegmentMemberResponseCollection) GetLinks() CollectionLinks {
	if o == nil {
		var ret CollectionLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GetSegmentMemberResponseCollection) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GetSegmentMemberResponseCollection) SetLinks(v CollectionLinks) {
	o.Links = v
}

func (o GetSegmentMemberResponseCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSegmentMemberResponseCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *GetSegmentMemberResponseCollection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"links",
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

	varGetSegmentMemberResponseCollection := _GetSegmentMemberResponseCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetSegmentMemberResponseCollection)

	if err != nil {
		return err
	}

	*o = GetSegmentMemberResponseCollection(varGetSegmentMemberResponseCollection)

	return err
}

type NullableGetSegmentMemberResponseCollection struct {
	value *GetSegmentMemberResponseCollection
	isSet bool
}

func (v NullableGetSegmentMemberResponseCollection) Get() *GetSegmentMemberResponseCollection {
	return v.value
}

func (v *NullableGetSegmentMemberResponseCollection) Set(val *GetSegmentMemberResponseCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSegmentMemberResponseCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSegmentMemberResponseCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSegmentMemberResponseCollection(val *GetSegmentMemberResponseCollection) *NullableGetSegmentMemberResponseCollection {
	return &NullableGetSegmentMemberResponseCollection{value: val, isSet: true}
}

func (v NullableGetSegmentMemberResponseCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSegmentMemberResponseCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


