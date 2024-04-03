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

// checks if the GetFlowMessageResponseCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFlowMessageResponseCollection{}

// GetFlowMessageResponseCollection struct for GetFlowMessageResponseCollection
type GetFlowMessageResponseCollection struct {
	Data []GetFlowMessageResponseCollectionDataInner `json:"data"`
	Links CollectionLinks `json:"links"`
}

type _GetFlowMessageResponseCollection GetFlowMessageResponseCollection

// NewGetFlowMessageResponseCollection instantiates a new GetFlowMessageResponseCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlowMessageResponseCollection(data []GetFlowMessageResponseCollectionDataInner, links CollectionLinks) *GetFlowMessageResponseCollection {
	this := GetFlowMessageResponseCollection{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGetFlowMessageResponseCollectionWithDefaults instantiates a new GetFlowMessageResponseCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlowMessageResponseCollectionWithDefaults() *GetFlowMessageResponseCollection {
	this := GetFlowMessageResponseCollection{}
	return &this
}

// GetData returns the Data field value
func (o *GetFlowMessageResponseCollection) GetData() []GetFlowMessageResponseCollectionDataInner {
	if o == nil {
		var ret []GetFlowMessageResponseCollectionDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetFlowMessageResponseCollection) GetDataOk() ([]GetFlowMessageResponseCollectionDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetFlowMessageResponseCollection) SetData(v []GetFlowMessageResponseCollectionDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *GetFlowMessageResponseCollection) GetLinks() CollectionLinks {
	if o == nil {
		var ret CollectionLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GetFlowMessageResponseCollection) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GetFlowMessageResponseCollection) SetLinks(v CollectionLinks) {
	o.Links = v
}

func (o GetFlowMessageResponseCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlowMessageResponseCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *GetFlowMessageResponseCollection) UnmarshalJSON(data []byte) (err error) {
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

	varGetFlowMessageResponseCollection := _GetFlowMessageResponseCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetFlowMessageResponseCollection)

	if err != nil {
		return err
	}

	*o = GetFlowMessageResponseCollection(varGetFlowMessageResponseCollection)

	return err
}

type NullableGetFlowMessageResponseCollection struct {
	value *GetFlowMessageResponseCollection
	isSet bool
}

func (v NullableGetFlowMessageResponseCollection) Get() *GetFlowMessageResponseCollection {
	return v.value
}

func (v *NullableGetFlowMessageResponseCollection) Set(val *GetFlowMessageResponseCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlowMessageResponseCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlowMessageResponseCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlowMessageResponseCollection(val *GetFlowMessageResponseCollection) *NullableGetFlowMessageResponseCollection {
	return &NullableGetFlowMessageResponseCollection{value: val, isSet: true}
}

func (v NullableGetFlowMessageResponseCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlowMessageResponseCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


