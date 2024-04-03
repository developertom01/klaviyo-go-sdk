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

// checks if the GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants{}

// GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants struct for GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants
type GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants struct {
	Data []GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariantsDataInner `json:"data"`
	Links *RelationshipLinks `json:"links,omitempty"`
}

type _GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants

// NewGetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants instantiates a new GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants(data []GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariantsDataInner) *GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants {
	this := GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants{}
	this.Data = data
	return &this
}

// NewGetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariantsWithDefaults instantiates a new GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariantsWithDefaults() *GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants {
	this := GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants{}
	return &this
}

// GetData returns the Data field value
func (o *GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants) GetData() []GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariantsDataInner {
	if o == nil {
		var ret []GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariantsDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants) GetDataOk() ([]GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariantsDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants) SetData(v []GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariantsDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

func (o GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants) UnmarshalJSON(data []byte) (err error) {
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

	varGetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants := _GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants)

	if err != nil {
		return err
	}

	*o = GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants(varGetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants)

	return err
}

type NullableGetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants struct {
	value *GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants
	isSet bool
}

func (v NullableGetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants) Get() *GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants {
	return v.value
}

func (v *NullableGetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants) Set(val *GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants(val *GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants) *NullableGetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants {
	return &NullableGetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants{value: val, isSet: true}
}

func (v NullableGetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsVariants) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

