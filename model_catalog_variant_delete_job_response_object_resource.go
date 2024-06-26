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

// checks if the CatalogVariantDeleteJobResponseObjectResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogVariantDeleteJobResponseObjectResource{}

// CatalogVariantDeleteJobResponseObjectResource struct for CatalogVariantDeleteJobResponseObjectResource
type CatalogVariantDeleteJobResponseObjectResource struct {
	Type CatalogVariantBulkDeleteJobEnum `json:"type"`
	// Unique identifier for retrieving the job. Generated by Klaviyo.
	Id string `json:"id"`
	Attributes CouponCodeCreateJobResponseObjectResourceAttributes `json:"attributes"`
	Links ObjectLinks `json:"links"`
}

type _CatalogVariantDeleteJobResponseObjectResource CatalogVariantDeleteJobResponseObjectResource

// NewCatalogVariantDeleteJobResponseObjectResource instantiates a new CatalogVariantDeleteJobResponseObjectResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogVariantDeleteJobResponseObjectResource(type_ CatalogVariantBulkDeleteJobEnum, id string, attributes CouponCodeCreateJobResponseObjectResourceAttributes, links ObjectLinks) *CatalogVariantDeleteJobResponseObjectResource {
	this := CatalogVariantDeleteJobResponseObjectResource{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewCatalogVariantDeleteJobResponseObjectResourceWithDefaults instantiates a new CatalogVariantDeleteJobResponseObjectResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogVariantDeleteJobResponseObjectResourceWithDefaults() *CatalogVariantDeleteJobResponseObjectResource {
	this := CatalogVariantDeleteJobResponseObjectResource{}
	return &this
}

// GetType returns the Type field value
func (o *CatalogVariantDeleteJobResponseObjectResource) GetType() CatalogVariantBulkDeleteJobEnum {
	if o == nil {
		var ret CatalogVariantBulkDeleteJobEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CatalogVariantDeleteJobResponseObjectResource) GetTypeOk() (*CatalogVariantBulkDeleteJobEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CatalogVariantDeleteJobResponseObjectResource) SetType(v CatalogVariantBulkDeleteJobEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CatalogVariantDeleteJobResponseObjectResource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CatalogVariantDeleteJobResponseObjectResource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CatalogVariantDeleteJobResponseObjectResource) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *CatalogVariantDeleteJobResponseObjectResource) GetAttributes() CouponCodeCreateJobResponseObjectResourceAttributes {
	if o == nil {
		var ret CouponCodeCreateJobResponseObjectResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CatalogVariantDeleteJobResponseObjectResource) GetAttributesOk() (*CouponCodeCreateJobResponseObjectResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CatalogVariantDeleteJobResponseObjectResource) SetAttributes(v CouponCodeCreateJobResponseObjectResourceAttributes) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *CatalogVariantDeleteJobResponseObjectResource) GetLinks() ObjectLinks {
	if o == nil {
		var ret ObjectLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CatalogVariantDeleteJobResponseObjectResource) GetLinksOk() (*ObjectLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CatalogVariantDeleteJobResponseObjectResource) SetLinks(v ObjectLinks) {
	o.Links = v
}

func (o CatalogVariantDeleteJobResponseObjectResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogVariantDeleteJobResponseObjectResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["attributes"] = o.Attributes
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *CatalogVariantDeleteJobResponseObjectResource) UnmarshalJSON(data []byte) (err error) {
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

	varCatalogVariantDeleteJobResponseObjectResource := _CatalogVariantDeleteJobResponseObjectResource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogVariantDeleteJobResponseObjectResource)

	if err != nil {
		return err
	}

	*o = CatalogVariantDeleteJobResponseObjectResource(varCatalogVariantDeleteJobResponseObjectResource)

	return err
}

type NullableCatalogVariantDeleteJobResponseObjectResource struct {
	value *CatalogVariantDeleteJobResponseObjectResource
	isSet bool
}

func (v NullableCatalogVariantDeleteJobResponseObjectResource) Get() *CatalogVariantDeleteJobResponseObjectResource {
	return v.value
}

func (v *NullableCatalogVariantDeleteJobResponseObjectResource) Set(val *CatalogVariantDeleteJobResponseObjectResource) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogVariantDeleteJobResponseObjectResource) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogVariantDeleteJobResponseObjectResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogVariantDeleteJobResponseObjectResource(val *CatalogVariantDeleteJobResponseObjectResource) *NullableCatalogVariantDeleteJobResponseObjectResource {
	return &NullableCatalogVariantDeleteJobResponseObjectResource{value: val, isSet: true}
}

func (v NullableCatalogVariantDeleteJobResponseObjectResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogVariantDeleteJobResponseObjectResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


