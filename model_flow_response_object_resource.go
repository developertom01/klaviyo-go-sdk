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

// checks if the FlowResponseObjectResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlowResponseObjectResource{}

// FlowResponseObjectResource struct for FlowResponseObjectResource
type FlowResponseObjectResource struct {
	Type FlowEnum `json:"type"`
	Id string `json:"id"`
	Attributes FlowResponseObjectResourceAttributes `json:"attributes"`
	Links ObjectLinks `json:"links"`
}

type _FlowResponseObjectResource FlowResponseObjectResource

// NewFlowResponseObjectResource instantiates a new FlowResponseObjectResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowResponseObjectResource(type_ FlowEnum, id string, attributes FlowResponseObjectResourceAttributes, links ObjectLinks) *FlowResponseObjectResource {
	this := FlowResponseObjectResource{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewFlowResponseObjectResourceWithDefaults instantiates a new FlowResponseObjectResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowResponseObjectResourceWithDefaults() *FlowResponseObjectResource {
	this := FlowResponseObjectResource{}
	return &this
}

// GetType returns the Type field value
func (o *FlowResponseObjectResource) GetType() FlowEnum {
	if o == nil {
		var ret FlowEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FlowResponseObjectResource) GetTypeOk() (*FlowEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FlowResponseObjectResource) SetType(v FlowEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *FlowResponseObjectResource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FlowResponseObjectResource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FlowResponseObjectResource) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *FlowResponseObjectResource) GetAttributes() FlowResponseObjectResourceAttributes {
	if o == nil {
		var ret FlowResponseObjectResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *FlowResponseObjectResource) GetAttributesOk() (*FlowResponseObjectResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *FlowResponseObjectResource) SetAttributes(v FlowResponseObjectResourceAttributes) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *FlowResponseObjectResource) GetLinks() ObjectLinks {
	if o == nil {
		var ret ObjectLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *FlowResponseObjectResource) GetLinksOk() (*ObjectLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *FlowResponseObjectResource) SetLinks(v ObjectLinks) {
	o.Links = v
}

func (o FlowResponseObjectResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlowResponseObjectResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["attributes"] = o.Attributes
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *FlowResponseObjectResource) UnmarshalJSON(data []byte) (err error) {
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

	varFlowResponseObjectResource := _FlowResponseObjectResource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFlowResponseObjectResource)

	if err != nil {
		return err
	}

	*o = FlowResponseObjectResource(varFlowResponseObjectResource)

	return err
}

type NullableFlowResponseObjectResource struct {
	value *FlowResponseObjectResource
	isSet bool
}

func (v NullableFlowResponseObjectResource) Get() *FlowResponseObjectResource {
	return v.value
}

func (v *NullableFlowResponseObjectResource) Set(val *FlowResponseObjectResource) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowResponseObjectResource) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowResponseObjectResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowResponseObjectResource(val *FlowResponseObjectResource) *NullableFlowResponseObjectResource {
	return &NullableFlowResponseObjectResource{value: val, isSet: true}
}

func (v NullableFlowResponseObjectResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowResponseObjectResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

