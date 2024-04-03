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

// checks if the GetCatalogItemCreateJobResponseCollectionCompoundDocument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCatalogItemCreateJobResponseCollectionCompoundDocument{}

// GetCatalogItemCreateJobResponseCollectionCompoundDocument struct for GetCatalogItemCreateJobResponseCollectionCompoundDocument
type GetCatalogItemCreateJobResponseCollectionCompoundDocument struct {
	Data []GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner `json:"data"`
	Links CollectionLinks `json:"links"`
}

type _GetCatalogItemCreateJobResponseCollectionCompoundDocument GetCatalogItemCreateJobResponseCollectionCompoundDocument

// NewGetCatalogItemCreateJobResponseCollectionCompoundDocument instantiates a new GetCatalogItemCreateJobResponseCollectionCompoundDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCatalogItemCreateJobResponseCollectionCompoundDocument(data []GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner, links CollectionLinks) *GetCatalogItemCreateJobResponseCollectionCompoundDocument {
	this := GetCatalogItemCreateJobResponseCollectionCompoundDocument{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGetCatalogItemCreateJobResponseCollectionCompoundDocumentWithDefaults instantiates a new GetCatalogItemCreateJobResponseCollectionCompoundDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCatalogItemCreateJobResponseCollectionCompoundDocumentWithDefaults() *GetCatalogItemCreateJobResponseCollectionCompoundDocument {
	this := GetCatalogItemCreateJobResponseCollectionCompoundDocument{}
	return &this
}

// GetData returns the Data field value
func (o *GetCatalogItemCreateJobResponseCollectionCompoundDocument) GetData() []GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner {
	if o == nil {
		var ret []GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetCatalogItemCreateJobResponseCollectionCompoundDocument) GetDataOk() ([]GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetCatalogItemCreateJobResponseCollectionCompoundDocument) SetData(v []GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *GetCatalogItemCreateJobResponseCollectionCompoundDocument) GetLinks() CollectionLinks {
	if o == nil {
		var ret CollectionLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GetCatalogItemCreateJobResponseCollectionCompoundDocument) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GetCatalogItemCreateJobResponseCollectionCompoundDocument) SetLinks(v CollectionLinks) {
	o.Links = v
}

func (o GetCatalogItemCreateJobResponseCollectionCompoundDocument) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCatalogItemCreateJobResponseCollectionCompoundDocument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *GetCatalogItemCreateJobResponseCollectionCompoundDocument) UnmarshalJSON(data []byte) (err error) {
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

	varGetCatalogItemCreateJobResponseCollectionCompoundDocument := _GetCatalogItemCreateJobResponseCollectionCompoundDocument{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCatalogItemCreateJobResponseCollectionCompoundDocument)

	if err != nil {
		return err
	}

	*o = GetCatalogItemCreateJobResponseCollectionCompoundDocument(varGetCatalogItemCreateJobResponseCollectionCompoundDocument)

	return err
}

type NullableGetCatalogItemCreateJobResponseCollectionCompoundDocument struct {
	value *GetCatalogItemCreateJobResponseCollectionCompoundDocument
	isSet bool
}

func (v NullableGetCatalogItemCreateJobResponseCollectionCompoundDocument) Get() *GetCatalogItemCreateJobResponseCollectionCompoundDocument {
	return v.value
}

func (v *NullableGetCatalogItemCreateJobResponseCollectionCompoundDocument) Set(val *GetCatalogItemCreateJobResponseCollectionCompoundDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCatalogItemCreateJobResponseCollectionCompoundDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCatalogItemCreateJobResponseCollectionCompoundDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCatalogItemCreateJobResponseCollectionCompoundDocument(val *GetCatalogItemCreateJobResponseCollectionCompoundDocument) *NullableGetCatalogItemCreateJobResponseCollectionCompoundDocument {
	return &NullableGetCatalogItemCreateJobResponseCollectionCompoundDocument{value: val, isSet: true}
}

func (v NullableGetCatalogItemCreateJobResponseCollectionCompoundDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCatalogItemCreateJobResponseCollectionCompoundDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

