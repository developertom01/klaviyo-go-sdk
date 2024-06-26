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

// checks if the GetCatalogCategoryDeleteJobResponseCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCatalogCategoryDeleteJobResponseCollection{}

// GetCatalogCategoryDeleteJobResponseCollection struct for GetCatalogCategoryDeleteJobResponseCollection
type GetCatalogCategoryDeleteJobResponseCollection struct {
	Data []GetCatalogCategoryDeleteJobResponseCollectionDataInner `json:"data"`
	Links CollectionLinks `json:"links"`
}

type _GetCatalogCategoryDeleteJobResponseCollection GetCatalogCategoryDeleteJobResponseCollection

// NewGetCatalogCategoryDeleteJobResponseCollection instantiates a new GetCatalogCategoryDeleteJobResponseCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCatalogCategoryDeleteJobResponseCollection(data []GetCatalogCategoryDeleteJobResponseCollectionDataInner, links CollectionLinks) *GetCatalogCategoryDeleteJobResponseCollection {
	this := GetCatalogCategoryDeleteJobResponseCollection{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGetCatalogCategoryDeleteJobResponseCollectionWithDefaults instantiates a new GetCatalogCategoryDeleteJobResponseCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCatalogCategoryDeleteJobResponseCollectionWithDefaults() *GetCatalogCategoryDeleteJobResponseCollection {
	this := GetCatalogCategoryDeleteJobResponseCollection{}
	return &this
}

// GetData returns the Data field value
func (o *GetCatalogCategoryDeleteJobResponseCollection) GetData() []GetCatalogCategoryDeleteJobResponseCollectionDataInner {
	if o == nil {
		var ret []GetCatalogCategoryDeleteJobResponseCollectionDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetCatalogCategoryDeleteJobResponseCollection) GetDataOk() ([]GetCatalogCategoryDeleteJobResponseCollectionDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetCatalogCategoryDeleteJobResponseCollection) SetData(v []GetCatalogCategoryDeleteJobResponseCollectionDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *GetCatalogCategoryDeleteJobResponseCollection) GetLinks() CollectionLinks {
	if o == nil {
		var ret CollectionLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GetCatalogCategoryDeleteJobResponseCollection) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GetCatalogCategoryDeleteJobResponseCollection) SetLinks(v CollectionLinks) {
	o.Links = v
}

func (o GetCatalogCategoryDeleteJobResponseCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCatalogCategoryDeleteJobResponseCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *GetCatalogCategoryDeleteJobResponseCollection) UnmarshalJSON(data []byte) (err error) {
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

	varGetCatalogCategoryDeleteJobResponseCollection := _GetCatalogCategoryDeleteJobResponseCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCatalogCategoryDeleteJobResponseCollection)

	if err != nil {
		return err
	}

	*o = GetCatalogCategoryDeleteJobResponseCollection(varGetCatalogCategoryDeleteJobResponseCollection)

	return err
}

type NullableGetCatalogCategoryDeleteJobResponseCollection struct {
	value *GetCatalogCategoryDeleteJobResponseCollection
	isSet bool
}

func (v NullableGetCatalogCategoryDeleteJobResponseCollection) Get() *GetCatalogCategoryDeleteJobResponseCollection {
	return v.value
}

func (v *NullableGetCatalogCategoryDeleteJobResponseCollection) Set(val *GetCatalogCategoryDeleteJobResponseCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCatalogCategoryDeleteJobResponseCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCatalogCategoryDeleteJobResponseCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCatalogCategoryDeleteJobResponseCollection(val *GetCatalogCategoryDeleteJobResponseCollection) *NullableGetCatalogCategoryDeleteJobResponseCollection {
	return &NullableGetCatalogCategoryDeleteJobResponseCollection{value: val, isSet: true}
}

func (v NullableGetCatalogCategoryDeleteJobResponseCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCatalogCategoryDeleteJobResponseCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


