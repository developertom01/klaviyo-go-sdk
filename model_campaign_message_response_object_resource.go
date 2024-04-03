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

// checks if the CampaignMessageResponseObjectResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignMessageResponseObjectResource{}

// CampaignMessageResponseObjectResource struct for CampaignMessageResponseObjectResource
type CampaignMessageResponseObjectResource struct {
	Type CampaignMessageEnum `json:"type"`
	// The message ID
	Id string `json:"id"`
	Attributes CampaignMessageResponseObjectResourceAttributes `json:"attributes"`
	Links ObjectLinks `json:"links"`
}

type _CampaignMessageResponseObjectResource CampaignMessageResponseObjectResource

// NewCampaignMessageResponseObjectResource instantiates a new CampaignMessageResponseObjectResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignMessageResponseObjectResource(type_ CampaignMessageEnum, id string, attributes CampaignMessageResponseObjectResourceAttributes, links ObjectLinks) *CampaignMessageResponseObjectResource {
	this := CampaignMessageResponseObjectResource{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewCampaignMessageResponseObjectResourceWithDefaults instantiates a new CampaignMessageResponseObjectResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignMessageResponseObjectResourceWithDefaults() *CampaignMessageResponseObjectResource {
	this := CampaignMessageResponseObjectResource{}
	return &this
}

// GetType returns the Type field value
func (o *CampaignMessageResponseObjectResource) GetType() CampaignMessageEnum {
	if o == nil {
		var ret CampaignMessageEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CampaignMessageResponseObjectResource) GetTypeOk() (*CampaignMessageEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CampaignMessageResponseObjectResource) SetType(v CampaignMessageEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CampaignMessageResponseObjectResource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CampaignMessageResponseObjectResource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CampaignMessageResponseObjectResource) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *CampaignMessageResponseObjectResource) GetAttributes() CampaignMessageResponseObjectResourceAttributes {
	if o == nil {
		var ret CampaignMessageResponseObjectResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CampaignMessageResponseObjectResource) GetAttributesOk() (*CampaignMessageResponseObjectResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CampaignMessageResponseObjectResource) SetAttributes(v CampaignMessageResponseObjectResourceAttributes) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *CampaignMessageResponseObjectResource) GetLinks() ObjectLinks {
	if o == nil {
		var ret ObjectLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CampaignMessageResponseObjectResource) GetLinksOk() (*ObjectLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CampaignMessageResponseObjectResource) SetLinks(v ObjectLinks) {
	o.Links = v
}

func (o CampaignMessageResponseObjectResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignMessageResponseObjectResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["attributes"] = o.Attributes
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *CampaignMessageResponseObjectResource) UnmarshalJSON(data []byte) (err error) {
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

	varCampaignMessageResponseObjectResource := _CampaignMessageResponseObjectResource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCampaignMessageResponseObjectResource)

	if err != nil {
		return err
	}

	*o = CampaignMessageResponseObjectResource(varCampaignMessageResponseObjectResource)

	return err
}

type NullableCampaignMessageResponseObjectResource struct {
	value *CampaignMessageResponseObjectResource
	isSet bool
}

func (v NullableCampaignMessageResponseObjectResource) Get() *CampaignMessageResponseObjectResource {
	return v.value
}

func (v *NullableCampaignMessageResponseObjectResource) Set(val *CampaignMessageResponseObjectResource) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignMessageResponseObjectResource) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignMessageResponseObjectResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignMessageResponseObjectResource(val *CampaignMessageResponseObjectResource) *NullableCampaignMessageResponseObjectResource {
	return &NullableCampaignMessageResponseObjectResource{value: val, isSet: true}
}

func (v NullableCampaignMessageResponseObjectResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignMessageResponseObjectResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

