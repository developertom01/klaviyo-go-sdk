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

// checks if the CampaignRecipientEstimationResponseObjectResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignRecipientEstimationResponseObjectResource{}

// CampaignRecipientEstimationResponseObjectResource struct for CampaignRecipientEstimationResponseObjectResource
type CampaignRecipientEstimationResponseObjectResource struct {
	Type CampaignRecipientEstimationEnum `json:"type"`
	// The ID of the campaign for which to get the estimated number of recipients
	Id string `json:"id"`
	Attributes CampaignRecipientEstimationResponseObjectResourceAttributes `json:"attributes"`
	Links ObjectLinks `json:"links"`
}

type _CampaignRecipientEstimationResponseObjectResource CampaignRecipientEstimationResponseObjectResource

// NewCampaignRecipientEstimationResponseObjectResource instantiates a new CampaignRecipientEstimationResponseObjectResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignRecipientEstimationResponseObjectResource(type_ CampaignRecipientEstimationEnum, id string, attributes CampaignRecipientEstimationResponseObjectResourceAttributes, links ObjectLinks) *CampaignRecipientEstimationResponseObjectResource {
	this := CampaignRecipientEstimationResponseObjectResource{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewCampaignRecipientEstimationResponseObjectResourceWithDefaults instantiates a new CampaignRecipientEstimationResponseObjectResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignRecipientEstimationResponseObjectResourceWithDefaults() *CampaignRecipientEstimationResponseObjectResource {
	this := CampaignRecipientEstimationResponseObjectResource{}
	return &this
}

// GetType returns the Type field value
func (o *CampaignRecipientEstimationResponseObjectResource) GetType() CampaignRecipientEstimationEnum {
	if o == nil {
		var ret CampaignRecipientEstimationEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CampaignRecipientEstimationResponseObjectResource) GetTypeOk() (*CampaignRecipientEstimationEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CampaignRecipientEstimationResponseObjectResource) SetType(v CampaignRecipientEstimationEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CampaignRecipientEstimationResponseObjectResource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CampaignRecipientEstimationResponseObjectResource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CampaignRecipientEstimationResponseObjectResource) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *CampaignRecipientEstimationResponseObjectResource) GetAttributes() CampaignRecipientEstimationResponseObjectResourceAttributes {
	if o == nil {
		var ret CampaignRecipientEstimationResponseObjectResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CampaignRecipientEstimationResponseObjectResource) GetAttributesOk() (*CampaignRecipientEstimationResponseObjectResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CampaignRecipientEstimationResponseObjectResource) SetAttributes(v CampaignRecipientEstimationResponseObjectResourceAttributes) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *CampaignRecipientEstimationResponseObjectResource) GetLinks() ObjectLinks {
	if o == nil {
		var ret ObjectLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CampaignRecipientEstimationResponseObjectResource) GetLinksOk() (*ObjectLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CampaignRecipientEstimationResponseObjectResource) SetLinks(v ObjectLinks) {
	o.Links = v
}

func (o CampaignRecipientEstimationResponseObjectResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignRecipientEstimationResponseObjectResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["attributes"] = o.Attributes
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *CampaignRecipientEstimationResponseObjectResource) UnmarshalJSON(data []byte) (err error) {
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

	varCampaignRecipientEstimationResponseObjectResource := _CampaignRecipientEstimationResponseObjectResource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCampaignRecipientEstimationResponseObjectResource)

	if err != nil {
		return err
	}

	*o = CampaignRecipientEstimationResponseObjectResource(varCampaignRecipientEstimationResponseObjectResource)

	return err
}

type NullableCampaignRecipientEstimationResponseObjectResource struct {
	value *CampaignRecipientEstimationResponseObjectResource
	isSet bool
}

func (v NullableCampaignRecipientEstimationResponseObjectResource) Get() *CampaignRecipientEstimationResponseObjectResource {
	return v.value
}

func (v *NullableCampaignRecipientEstimationResponseObjectResource) Set(val *CampaignRecipientEstimationResponseObjectResource) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignRecipientEstimationResponseObjectResource) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignRecipientEstimationResponseObjectResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignRecipientEstimationResponseObjectResource(val *CampaignRecipientEstimationResponseObjectResource) *NullableCampaignRecipientEstimationResponseObjectResource {
	return &NullableCampaignRecipientEstimationResponseObjectResource{value: val, isSet: true}
}

func (v NullableCampaignRecipientEstimationResponseObjectResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignRecipientEstimationResponseObjectResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


