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

// checks if the GetCatalogItemResponseCompoundDocument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCatalogItemResponseCompoundDocument{}

// GetCatalogItemResponseCompoundDocument struct for GetCatalogItemResponseCompoundDocument
type GetCatalogItemResponseCompoundDocument struct {
	Data GetCatalogItemResponseCollectionCompoundDocumentDataInner `json:"data"`
	Included []CatalogVariantResponseObjectResource `json:"included,omitempty"`
}

type _GetCatalogItemResponseCompoundDocument GetCatalogItemResponseCompoundDocument

// NewGetCatalogItemResponseCompoundDocument instantiates a new GetCatalogItemResponseCompoundDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCatalogItemResponseCompoundDocument(data GetCatalogItemResponseCollectionCompoundDocumentDataInner) *GetCatalogItemResponseCompoundDocument {
	this := GetCatalogItemResponseCompoundDocument{}
	this.Data = data
	return &this
}

// NewGetCatalogItemResponseCompoundDocumentWithDefaults instantiates a new GetCatalogItemResponseCompoundDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCatalogItemResponseCompoundDocumentWithDefaults() *GetCatalogItemResponseCompoundDocument {
	this := GetCatalogItemResponseCompoundDocument{}
	return &this
}

// GetData returns the Data field value
func (o *GetCatalogItemResponseCompoundDocument) GetData() GetCatalogItemResponseCollectionCompoundDocumentDataInner {
	if o == nil {
		var ret GetCatalogItemResponseCollectionCompoundDocumentDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetCatalogItemResponseCompoundDocument) GetDataOk() (*GetCatalogItemResponseCollectionCompoundDocumentDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetCatalogItemResponseCompoundDocument) SetData(v GetCatalogItemResponseCollectionCompoundDocumentDataInner) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *GetCatalogItemResponseCompoundDocument) GetIncluded() []CatalogVariantResponseObjectResource {
	if o == nil || IsNil(o.Included) {
		var ret []CatalogVariantResponseObjectResource
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCatalogItemResponseCompoundDocument) GetIncludedOk() ([]CatalogVariantResponseObjectResource, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *GetCatalogItemResponseCompoundDocument) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []CatalogVariantResponseObjectResource and assigns it to the Included field.
func (o *GetCatalogItemResponseCompoundDocument) SetIncluded(v []CatalogVariantResponseObjectResource) {
	o.Included = v
}

func (o GetCatalogItemResponseCompoundDocument) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCatalogItemResponseCompoundDocument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	return toSerialize, nil
}

func (o *GetCatalogItemResponseCompoundDocument) UnmarshalJSON(data []byte) (err error) {
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

	varGetCatalogItemResponseCompoundDocument := _GetCatalogItemResponseCompoundDocument{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCatalogItemResponseCompoundDocument)

	if err != nil {
		return err
	}

	*o = GetCatalogItemResponseCompoundDocument(varGetCatalogItemResponseCompoundDocument)

	return err
}

type NullableGetCatalogItemResponseCompoundDocument struct {
	value *GetCatalogItemResponseCompoundDocument
	isSet bool
}

func (v NullableGetCatalogItemResponseCompoundDocument) Get() *GetCatalogItemResponseCompoundDocument {
	return v.value
}

func (v *NullableGetCatalogItemResponseCompoundDocument) Set(val *GetCatalogItemResponseCompoundDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCatalogItemResponseCompoundDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCatalogItemResponseCompoundDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCatalogItemResponseCompoundDocument(val *GetCatalogItemResponseCompoundDocument) *NullableGetCatalogItemResponseCompoundDocument {
	return &NullableGetCatalogItemResponseCompoundDocument{value: val, isSet: true}
}

func (v NullableGetCatalogItemResponseCompoundDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCatalogItemResponseCompoundDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

