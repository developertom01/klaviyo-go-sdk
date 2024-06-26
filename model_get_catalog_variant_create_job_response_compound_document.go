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

// checks if the GetCatalogVariantCreateJobResponseCompoundDocument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCatalogVariantCreateJobResponseCompoundDocument{}

// GetCatalogVariantCreateJobResponseCompoundDocument struct for GetCatalogVariantCreateJobResponseCompoundDocument
type GetCatalogVariantCreateJobResponseCompoundDocument struct {
	Data GetCatalogVariantCreateJobResponseCollectionCompoundDocumentDataInner `json:"data"`
	Included []CatalogVariantResponseObjectResource `json:"included,omitempty"`
}

type _GetCatalogVariantCreateJobResponseCompoundDocument GetCatalogVariantCreateJobResponseCompoundDocument

// NewGetCatalogVariantCreateJobResponseCompoundDocument instantiates a new GetCatalogVariantCreateJobResponseCompoundDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCatalogVariantCreateJobResponseCompoundDocument(data GetCatalogVariantCreateJobResponseCollectionCompoundDocumentDataInner) *GetCatalogVariantCreateJobResponseCompoundDocument {
	this := GetCatalogVariantCreateJobResponseCompoundDocument{}
	this.Data = data
	return &this
}

// NewGetCatalogVariantCreateJobResponseCompoundDocumentWithDefaults instantiates a new GetCatalogVariantCreateJobResponseCompoundDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCatalogVariantCreateJobResponseCompoundDocumentWithDefaults() *GetCatalogVariantCreateJobResponseCompoundDocument {
	this := GetCatalogVariantCreateJobResponseCompoundDocument{}
	return &this
}

// GetData returns the Data field value
func (o *GetCatalogVariantCreateJobResponseCompoundDocument) GetData() GetCatalogVariantCreateJobResponseCollectionCompoundDocumentDataInner {
	if o == nil {
		var ret GetCatalogVariantCreateJobResponseCollectionCompoundDocumentDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetCatalogVariantCreateJobResponseCompoundDocument) GetDataOk() (*GetCatalogVariantCreateJobResponseCollectionCompoundDocumentDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetCatalogVariantCreateJobResponseCompoundDocument) SetData(v GetCatalogVariantCreateJobResponseCollectionCompoundDocumentDataInner) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *GetCatalogVariantCreateJobResponseCompoundDocument) GetIncluded() []CatalogVariantResponseObjectResource {
	if o == nil || IsNil(o.Included) {
		var ret []CatalogVariantResponseObjectResource
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCatalogVariantCreateJobResponseCompoundDocument) GetIncludedOk() ([]CatalogVariantResponseObjectResource, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *GetCatalogVariantCreateJobResponseCompoundDocument) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []CatalogVariantResponseObjectResource and assigns it to the Included field.
func (o *GetCatalogVariantCreateJobResponseCompoundDocument) SetIncluded(v []CatalogVariantResponseObjectResource) {
	o.Included = v
}

func (o GetCatalogVariantCreateJobResponseCompoundDocument) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCatalogVariantCreateJobResponseCompoundDocument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	return toSerialize, nil
}

func (o *GetCatalogVariantCreateJobResponseCompoundDocument) UnmarshalJSON(data []byte) (err error) {
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

	varGetCatalogVariantCreateJobResponseCompoundDocument := _GetCatalogVariantCreateJobResponseCompoundDocument{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCatalogVariantCreateJobResponseCompoundDocument)

	if err != nil {
		return err
	}

	*o = GetCatalogVariantCreateJobResponseCompoundDocument(varGetCatalogVariantCreateJobResponseCompoundDocument)

	return err
}

type NullableGetCatalogVariantCreateJobResponseCompoundDocument struct {
	value *GetCatalogVariantCreateJobResponseCompoundDocument
	isSet bool
}

func (v NullableGetCatalogVariantCreateJobResponseCompoundDocument) Get() *GetCatalogVariantCreateJobResponseCompoundDocument {
	return v.value
}

func (v *NullableGetCatalogVariantCreateJobResponseCompoundDocument) Set(val *GetCatalogVariantCreateJobResponseCompoundDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCatalogVariantCreateJobResponseCompoundDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCatalogVariantCreateJobResponseCompoundDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCatalogVariantCreateJobResponseCompoundDocument(val *GetCatalogVariantCreateJobResponseCompoundDocument) *NullableGetCatalogVariantCreateJobResponseCompoundDocument {
	return &NullableGetCatalogVariantCreateJobResponseCompoundDocument{value: val, isSet: true}
}

func (v NullableGetCatalogVariantCreateJobResponseCompoundDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCatalogVariantCreateJobResponseCompoundDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


