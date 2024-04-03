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

// checks if the GetCatalogCategoryCreateJobResponseCollectionCompoundDocument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCatalogCategoryCreateJobResponseCollectionCompoundDocument{}

// GetCatalogCategoryCreateJobResponseCollectionCompoundDocument struct for GetCatalogCategoryCreateJobResponseCollectionCompoundDocument
type GetCatalogCategoryCreateJobResponseCollectionCompoundDocument struct {
	Data []GetCatalogCategoryCreateJobResponseCollectionCompoundDocumentDataInner `json:"data"`
	Links CollectionLinks `json:"links"`
}

type _GetCatalogCategoryCreateJobResponseCollectionCompoundDocument GetCatalogCategoryCreateJobResponseCollectionCompoundDocument

// NewGetCatalogCategoryCreateJobResponseCollectionCompoundDocument instantiates a new GetCatalogCategoryCreateJobResponseCollectionCompoundDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCatalogCategoryCreateJobResponseCollectionCompoundDocument(data []GetCatalogCategoryCreateJobResponseCollectionCompoundDocumentDataInner, links CollectionLinks) *GetCatalogCategoryCreateJobResponseCollectionCompoundDocument {
	this := GetCatalogCategoryCreateJobResponseCollectionCompoundDocument{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGetCatalogCategoryCreateJobResponseCollectionCompoundDocumentWithDefaults instantiates a new GetCatalogCategoryCreateJobResponseCollectionCompoundDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCatalogCategoryCreateJobResponseCollectionCompoundDocumentWithDefaults() *GetCatalogCategoryCreateJobResponseCollectionCompoundDocument {
	this := GetCatalogCategoryCreateJobResponseCollectionCompoundDocument{}
	return &this
}

// GetData returns the Data field value
func (o *GetCatalogCategoryCreateJobResponseCollectionCompoundDocument) GetData() []GetCatalogCategoryCreateJobResponseCollectionCompoundDocumentDataInner {
	if o == nil {
		var ret []GetCatalogCategoryCreateJobResponseCollectionCompoundDocumentDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetCatalogCategoryCreateJobResponseCollectionCompoundDocument) GetDataOk() ([]GetCatalogCategoryCreateJobResponseCollectionCompoundDocumentDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetCatalogCategoryCreateJobResponseCollectionCompoundDocument) SetData(v []GetCatalogCategoryCreateJobResponseCollectionCompoundDocumentDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *GetCatalogCategoryCreateJobResponseCollectionCompoundDocument) GetLinks() CollectionLinks {
	if o == nil {
		var ret CollectionLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GetCatalogCategoryCreateJobResponseCollectionCompoundDocument) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GetCatalogCategoryCreateJobResponseCollectionCompoundDocument) SetLinks(v CollectionLinks) {
	o.Links = v
}

func (o GetCatalogCategoryCreateJobResponseCollectionCompoundDocument) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCatalogCategoryCreateJobResponseCollectionCompoundDocument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *GetCatalogCategoryCreateJobResponseCollectionCompoundDocument) UnmarshalJSON(data []byte) (err error) {
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

	varGetCatalogCategoryCreateJobResponseCollectionCompoundDocument := _GetCatalogCategoryCreateJobResponseCollectionCompoundDocument{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCatalogCategoryCreateJobResponseCollectionCompoundDocument)

	if err != nil {
		return err
	}

	*o = GetCatalogCategoryCreateJobResponseCollectionCompoundDocument(varGetCatalogCategoryCreateJobResponseCollectionCompoundDocument)

	return err
}

type NullableGetCatalogCategoryCreateJobResponseCollectionCompoundDocument struct {
	value *GetCatalogCategoryCreateJobResponseCollectionCompoundDocument
	isSet bool
}

func (v NullableGetCatalogCategoryCreateJobResponseCollectionCompoundDocument) Get() *GetCatalogCategoryCreateJobResponseCollectionCompoundDocument {
	return v.value
}

func (v *NullableGetCatalogCategoryCreateJobResponseCollectionCompoundDocument) Set(val *GetCatalogCategoryCreateJobResponseCollectionCompoundDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCatalogCategoryCreateJobResponseCollectionCompoundDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCatalogCategoryCreateJobResponseCollectionCompoundDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCatalogCategoryCreateJobResponseCollectionCompoundDocument(val *GetCatalogCategoryCreateJobResponseCollectionCompoundDocument) *NullableGetCatalogCategoryCreateJobResponseCollectionCompoundDocument {
	return &NullableGetCatalogCategoryCreateJobResponseCollectionCompoundDocument{value: val, isSet: true}
}

func (v NullableGetCatalogCategoryCreateJobResponseCollectionCompoundDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCatalogCategoryCreateJobResponseCollectionCompoundDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

