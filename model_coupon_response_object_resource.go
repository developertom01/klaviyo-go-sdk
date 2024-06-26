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

// checks if the CouponResponseObjectResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CouponResponseObjectResource{}

// CouponResponseObjectResource struct for CouponResponseObjectResource
type CouponResponseObjectResource struct {
	Type CouponEnum `json:"type"`
	// The internal id of a Coupon is equivalent to its external id stored within an integration.
	Id string `json:"id"`
	Attributes CouponResponseObjectResourceAttributes `json:"attributes"`
	Links ObjectLinks `json:"links"`
}

type _CouponResponseObjectResource CouponResponseObjectResource

// NewCouponResponseObjectResource instantiates a new CouponResponseObjectResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponResponseObjectResource(type_ CouponEnum, id string, attributes CouponResponseObjectResourceAttributes, links ObjectLinks) *CouponResponseObjectResource {
	this := CouponResponseObjectResource{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewCouponResponseObjectResourceWithDefaults instantiates a new CouponResponseObjectResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponResponseObjectResourceWithDefaults() *CouponResponseObjectResource {
	this := CouponResponseObjectResource{}
	return &this
}

// GetType returns the Type field value
func (o *CouponResponseObjectResource) GetType() CouponEnum {
	if o == nil {
		var ret CouponEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CouponResponseObjectResource) GetTypeOk() (*CouponEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CouponResponseObjectResource) SetType(v CouponEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CouponResponseObjectResource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CouponResponseObjectResource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CouponResponseObjectResource) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *CouponResponseObjectResource) GetAttributes() CouponResponseObjectResourceAttributes {
	if o == nil {
		var ret CouponResponseObjectResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CouponResponseObjectResource) GetAttributesOk() (*CouponResponseObjectResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CouponResponseObjectResource) SetAttributes(v CouponResponseObjectResourceAttributes) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *CouponResponseObjectResource) GetLinks() ObjectLinks {
	if o == nil {
		var ret ObjectLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CouponResponseObjectResource) GetLinksOk() (*ObjectLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CouponResponseObjectResource) SetLinks(v ObjectLinks) {
	o.Links = v
}

func (o CouponResponseObjectResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CouponResponseObjectResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["attributes"] = o.Attributes
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *CouponResponseObjectResource) UnmarshalJSON(data []byte) (err error) {
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

	varCouponResponseObjectResource := _CouponResponseObjectResource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCouponResponseObjectResource)

	if err != nil {
		return err
	}

	*o = CouponResponseObjectResource(varCouponResponseObjectResource)

	return err
}

type NullableCouponResponseObjectResource struct {
	value *CouponResponseObjectResource
	isSet bool
}

func (v NullableCouponResponseObjectResource) Get() *CouponResponseObjectResource {
	return v.value
}

func (v *NullableCouponResponseObjectResource) Set(val *CouponResponseObjectResource) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponResponseObjectResource) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponResponseObjectResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponResponseObjectResource(val *CouponResponseObjectResource) *NullableCouponResponseObjectResource {
	return &NullableCouponResponseObjectResource{value: val, isSet: true}
}

func (v NullableCouponResponseObjectResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponResponseObjectResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


