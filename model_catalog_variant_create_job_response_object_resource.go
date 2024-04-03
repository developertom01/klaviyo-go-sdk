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

// checks if the CatalogVariantCreateJobResponseObjectResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogVariantCreateJobResponseObjectResource{}

// CatalogVariantCreateJobResponseObjectResource struct for CatalogVariantCreateJobResponseObjectResource
type CatalogVariantCreateJobResponseObjectResource struct {
	Type CatalogVariantBulkCreateJobEnum `json:"type"`
	// Unique identifier for retrieving the job. Generated by Klaviyo.
	Id string `json:"id"`
	Attributes CouponCodeCreateJobResponseObjectResourceAttributes `json:"attributes"`
	Links ObjectLinks `json:"links"`
}

type _CatalogVariantCreateJobResponseObjectResource CatalogVariantCreateJobResponseObjectResource

// NewCatalogVariantCreateJobResponseObjectResource instantiates a new CatalogVariantCreateJobResponseObjectResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogVariantCreateJobResponseObjectResource(type_ CatalogVariantBulkCreateJobEnum, id string, attributes CouponCodeCreateJobResponseObjectResourceAttributes, links ObjectLinks) *CatalogVariantCreateJobResponseObjectResource {
	this := CatalogVariantCreateJobResponseObjectResource{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewCatalogVariantCreateJobResponseObjectResourceWithDefaults instantiates a new CatalogVariantCreateJobResponseObjectResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogVariantCreateJobResponseObjectResourceWithDefaults() *CatalogVariantCreateJobResponseObjectResource {
	this := CatalogVariantCreateJobResponseObjectResource{}
	return &this
}

// GetType returns the Type field value
func (o *CatalogVariantCreateJobResponseObjectResource) GetType() CatalogVariantBulkCreateJobEnum {
	if o == nil {
		var ret CatalogVariantBulkCreateJobEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CatalogVariantCreateJobResponseObjectResource) GetTypeOk() (*CatalogVariantBulkCreateJobEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CatalogVariantCreateJobResponseObjectResource) SetType(v CatalogVariantBulkCreateJobEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CatalogVariantCreateJobResponseObjectResource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CatalogVariantCreateJobResponseObjectResource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CatalogVariantCreateJobResponseObjectResource) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *CatalogVariantCreateJobResponseObjectResource) GetAttributes() CouponCodeCreateJobResponseObjectResourceAttributes {
	if o == nil {
		var ret CouponCodeCreateJobResponseObjectResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CatalogVariantCreateJobResponseObjectResource) GetAttributesOk() (*CouponCodeCreateJobResponseObjectResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CatalogVariantCreateJobResponseObjectResource) SetAttributes(v CouponCodeCreateJobResponseObjectResourceAttributes) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *CatalogVariantCreateJobResponseObjectResource) GetLinks() ObjectLinks {
	if o == nil {
		var ret ObjectLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CatalogVariantCreateJobResponseObjectResource) GetLinksOk() (*ObjectLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CatalogVariantCreateJobResponseObjectResource) SetLinks(v ObjectLinks) {
	o.Links = v
}

func (o CatalogVariantCreateJobResponseObjectResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogVariantCreateJobResponseObjectResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["attributes"] = o.Attributes
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *CatalogVariantCreateJobResponseObjectResource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"attributes",
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

	varCatalogVariantCreateJobResponseObjectResource := _CatalogVariantCreateJobResponseObjectResource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogVariantCreateJobResponseObjectResource)

	if err != nil {
		return err
	}

	*o = CatalogVariantCreateJobResponseObjectResource(varCatalogVariantCreateJobResponseObjectResource)

	return err
}

type NullableCatalogVariantCreateJobResponseObjectResource struct {
	value *CatalogVariantCreateJobResponseObjectResource
	isSet bool
}

func (v NullableCatalogVariantCreateJobResponseObjectResource) Get() *CatalogVariantCreateJobResponseObjectResource {
	return v.value
}

func (v *NullableCatalogVariantCreateJobResponseObjectResource) Set(val *CatalogVariantCreateJobResponseObjectResource) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogVariantCreateJobResponseObjectResource) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogVariantCreateJobResponseObjectResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogVariantCreateJobResponseObjectResource(val *CatalogVariantCreateJobResponseObjectResource) *NullableCatalogVariantCreateJobResponseObjectResource {
	return &NullableCatalogVariantCreateJobResponseObjectResource{value: val, isSet: true}
}

func (v NullableCatalogVariantCreateJobResponseObjectResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogVariantCreateJobResponseObjectResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

