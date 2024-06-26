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

// checks if the GetCatalogCategoryResponseCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCatalogCategoryResponseCollection{}

// GetCatalogCategoryResponseCollection struct for GetCatalogCategoryResponseCollection
type GetCatalogCategoryResponseCollection struct {
	Data []GetCatalogCategoryResponseCollectionDataInner `json:"data"`
	Links CollectionLinks `json:"links"`
}

type _GetCatalogCategoryResponseCollection GetCatalogCategoryResponseCollection

// NewGetCatalogCategoryResponseCollection instantiates a new GetCatalogCategoryResponseCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCatalogCategoryResponseCollection(data []GetCatalogCategoryResponseCollectionDataInner, links CollectionLinks) *GetCatalogCategoryResponseCollection {
	this := GetCatalogCategoryResponseCollection{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGetCatalogCategoryResponseCollectionWithDefaults instantiates a new GetCatalogCategoryResponseCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCatalogCategoryResponseCollectionWithDefaults() *GetCatalogCategoryResponseCollection {
	this := GetCatalogCategoryResponseCollection{}
	return &this
}

// GetData returns the Data field value
func (o *GetCatalogCategoryResponseCollection) GetData() []GetCatalogCategoryResponseCollectionDataInner {
	if o == nil {
		var ret []GetCatalogCategoryResponseCollectionDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetCatalogCategoryResponseCollection) GetDataOk() ([]GetCatalogCategoryResponseCollectionDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetCatalogCategoryResponseCollection) SetData(v []GetCatalogCategoryResponseCollectionDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *GetCatalogCategoryResponseCollection) GetLinks() CollectionLinks {
	if o == nil {
		var ret CollectionLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GetCatalogCategoryResponseCollection) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GetCatalogCategoryResponseCollection) SetLinks(v CollectionLinks) {
	o.Links = v
}

func (o GetCatalogCategoryResponseCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCatalogCategoryResponseCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *GetCatalogCategoryResponseCollection) UnmarshalJSON(data []byte) (err error) {
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

	varGetCatalogCategoryResponseCollection := _GetCatalogCategoryResponseCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCatalogCategoryResponseCollection)

	if err != nil {
		return err
	}

	*o = GetCatalogCategoryResponseCollection(varGetCatalogCategoryResponseCollection)

	return err
}

type NullableGetCatalogCategoryResponseCollection struct {
	value *GetCatalogCategoryResponseCollection
	isSet bool
}

func (v NullableGetCatalogCategoryResponseCollection) Get() *GetCatalogCategoryResponseCollection {
	return v.value
}

func (v *NullableGetCatalogCategoryResponseCollection) Set(val *GetCatalogCategoryResponseCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCatalogCategoryResponseCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCatalogCategoryResponseCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCatalogCategoryResponseCollection(val *GetCatalogCategoryResponseCollection) *NullableGetCatalogCategoryResponseCollection {
	return &NullableGetCatalogCategoryResponseCollection{value: val, isSet: true}
}

func (v NullableGetCatalogCategoryResponseCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCatalogCategoryResponseCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


