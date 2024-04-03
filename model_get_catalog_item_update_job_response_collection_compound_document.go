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

// checks if the GetCatalogItemUpdateJobResponseCollectionCompoundDocument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCatalogItemUpdateJobResponseCollectionCompoundDocument{}

// GetCatalogItemUpdateJobResponseCollectionCompoundDocument struct for GetCatalogItemUpdateJobResponseCollectionCompoundDocument
type GetCatalogItemUpdateJobResponseCollectionCompoundDocument struct {
	Data []GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInner `json:"data"`
	Links CollectionLinks `json:"links"`
}

type _GetCatalogItemUpdateJobResponseCollectionCompoundDocument GetCatalogItemUpdateJobResponseCollectionCompoundDocument

// NewGetCatalogItemUpdateJobResponseCollectionCompoundDocument instantiates a new GetCatalogItemUpdateJobResponseCollectionCompoundDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCatalogItemUpdateJobResponseCollectionCompoundDocument(data []GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInner, links CollectionLinks) *GetCatalogItemUpdateJobResponseCollectionCompoundDocument {
	this := GetCatalogItemUpdateJobResponseCollectionCompoundDocument{}
	this.Data = data
	this.Links = links
	return &this
}

// NewGetCatalogItemUpdateJobResponseCollectionCompoundDocumentWithDefaults instantiates a new GetCatalogItemUpdateJobResponseCollectionCompoundDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCatalogItemUpdateJobResponseCollectionCompoundDocumentWithDefaults() *GetCatalogItemUpdateJobResponseCollectionCompoundDocument {
	this := GetCatalogItemUpdateJobResponseCollectionCompoundDocument{}
	return &this
}

// GetData returns the Data field value
func (o *GetCatalogItemUpdateJobResponseCollectionCompoundDocument) GetData() []GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInner {
	if o == nil {
		var ret []GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetCatalogItemUpdateJobResponseCollectionCompoundDocument) GetDataOk() ([]GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetCatalogItemUpdateJobResponseCollectionCompoundDocument) SetData(v []GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *GetCatalogItemUpdateJobResponseCollectionCompoundDocument) GetLinks() CollectionLinks {
	if o == nil {
		var ret CollectionLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GetCatalogItemUpdateJobResponseCollectionCompoundDocument) GetLinksOk() (*CollectionLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GetCatalogItemUpdateJobResponseCollectionCompoundDocument) SetLinks(v CollectionLinks) {
	o.Links = v
}

func (o GetCatalogItemUpdateJobResponseCollectionCompoundDocument) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCatalogItemUpdateJobResponseCollectionCompoundDocument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *GetCatalogItemUpdateJobResponseCollectionCompoundDocument) UnmarshalJSON(data []byte) (err error) {
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

	varGetCatalogItemUpdateJobResponseCollectionCompoundDocument := _GetCatalogItemUpdateJobResponseCollectionCompoundDocument{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCatalogItemUpdateJobResponseCollectionCompoundDocument)

	if err != nil {
		return err
	}

	*o = GetCatalogItemUpdateJobResponseCollectionCompoundDocument(varGetCatalogItemUpdateJobResponseCollectionCompoundDocument)

	return err
}

type NullableGetCatalogItemUpdateJobResponseCollectionCompoundDocument struct {
	value *GetCatalogItemUpdateJobResponseCollectionCompoundDocument
	isSet bool
}

func (v NullableGetCatalogItemUpdateJobResponseCollectionCompoundDocument) Get() *GetCatalogItemUpdateJobResponseCollectionCompoundDocument {
	return v.value
}

func (v *NullableGetCatalogItemUpdateJobResponseCollectionCompoundDocument) Set(val *GetCatalogItemUpdateJobResponseCollectionCompoundDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCatalogItemUpdateJobResponseCollectionCompoundDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCatalogItemUpdateJobResponseCollectionCompoundDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCatalogItemUpdateJobResponseCollectionCompoundDocument(val *GetCatalogItemUpdateJobResponseCollectionCompoundDocument) *NullableGetCatalogItemUpdateJobResponseCollectionCompoundDocument {
	return &NullableGetCatalogItemUpdateJobResponseCollectionCompoundDocument{value: val, isSet: true}
}

func (v NullableGetCatalogItemUpdateJobResponseCollectionCompoundDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCatalogItemUpdateJobResponseCollectionCompoundDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

