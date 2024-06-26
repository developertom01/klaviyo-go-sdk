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

// checks if the FlowMessageResponseObjectResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlowMessageResponseObjectResource{}

// FlowMessageResponseObjectResource struct for FlowMessageResponseObjectResource
type FlowMessageResponseObjectResource struct {
	Type FlowMessageEnum `json:"type"`
	Id string `json:"id"`
	Attributes FlowMessageResponseObjectResourceAttributes `json:"attributes"`
	Links ObjectLinks `json:"links"`
}

type _FlowMessageResponseObjectResource FlowMessageResponseObjectResource

// NewFlowMessageResponseObjectResource instantiates a new FlowMessageResponseObjectResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowMessageResponseObjectResource(type_ FlowMessageEnum, id string, attributes FlowMessageResponseObjectResourceAttributes, links ObjectLinks) *FlowMessageResponseObjectResource {
	this := FlowMessageResponseObjectResource{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewFlowMessageResponseObjectResourceWithDefaults instantiates a new FlowMessageResponseObjectResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowMessageResponseObjectResourceWithDefaults() *FlowMessageResponseObjectResource {
	this := FlowMessageResponseObjectResource{}
	return &this
}

// GetType returns the Type field value
func (o *FlowMessageResponseObjectResource) GetType() FlowMessageEnum {
	if o == nil {
		var ret FlowMessageEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FlowMessageResponseObjectResource) GetTypeOk() (*FlowMessageEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FlowMessageResponseObjectResource) SetType(v FlowMessageEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *FlowMessageResponseObjectResource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FlowMessageResponseObjectResource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FlowMessageResponseObjectResource) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *FlowMessageResponseObjectResource) GetAttributes() FlowMessageResponseObjectResourceAttributes {
	if o == nil {
		var ret FlowMessageResponseObjectResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *FlowMessageResponseObjectResource) GetAttributesOk() (*FlowMessageResponseObjectResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *FlowMessageResponseObjectResource) SetAttributes(v FlowMessageResponseObjectResourceAttributes) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *FlowMessageResponseObjectResource) GetLinks() ObjectLinks {
	if o == nil {
		var ret ObjectLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *FlowMessageResponseObjectResource) GetLinksOk() (*ObjectLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *FlowMessageResponseObjectResource) SetLinks(v ObjectLinks) {
	o.Links = v
}

func (o FlowMessageResponseObjectResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlowMessageResponseObjectResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["attributes"] = o.Attributes
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *FlowMessageResponseObjectResource) UnmarshalJSON(data []byte) (err error) {
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

	varFlowMessageResponseObjectResource := _FlowMessageResponseObjectResource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFlowMessageResponseObjectResource)

	if err != nil {
		return err
	}

	*o = FlowMessageResponseObjectResource(varFlowMessageResponseObjectResource)

	return err
}

type NullableFlowMessageResponseObjectResource struct {
	value *FlowMessageResponseObjectResource
	isSet bool
}

func (v NullableFlowMessageResponseObjectResource) Get() *FlowMessageResponseObjectResource {
	return v.value
}

func (v *NullableFlowMessageResponseObjectResource) Set(val *FlowMessageResponseObjectResource) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowMessageResponseObjectResource) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowMessageResponseObjectResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowMessageResponseObjectResource(val *FlowMessageResponseObjectResource) *NullableFlowMessageResponseObjectResource {
	return &NullableFlowMessageResponseObjectResource{value: val, isSet: true}
}

func (v NullableFlowMessageResponseObjectResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowMessageResponseObjectResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


