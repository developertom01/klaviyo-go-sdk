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

// checks if the PostCampaignRecipientEstimationJobResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCampaignRecipientEstimationJobResponseData{}

// PostCampaignRecipientEstimationJobResponseData struct for PostCampaignRecipientEstimationJobResponseData
type PostCampaignRecipientEstimationJobResponseData struct {
	Type CampaignRecipientEstimationJobEnum `json:"type"`
	// The ID of the campaign used for estimating recipients
	Id string `json:"id"`
	Attributes CampaignRecipientEstimationJobResponseObjectResourceAttributes `json:"attributes"`
	Links ObjectLinks `json:"links"`
}

type _PostCampaignRecipientEstimationJobResponseData PostCampaignRecipientEstimationJobResponseData

// NewPostCampaignRecipientEstimationJobResponseData instantiates a new PostCampaignRecipientEstimationJobResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCampaignRecipientEstimationJobResponseData(type_ CampaignRecipientEstimationJobEnum, id string, attributes CampaignRecipientEstimationJobResponseObjectResourceAttributes, links ObjectLinks) *PostCampaignRecipientEstimationJobResponseData {
	this := PostCampaignRecipientEstimationJobResponseData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewPostCampaignRecipientEstimationJobResponseDataWithDefaults instantiates a new PostCampaignRecipientEstimationJobResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCampaignRecipientEstimationJobResponseDataWithDefaults() *PostCampaignRecipientEstimationJobResponseData {
	this := PostCampaignRecipientEstimationJobResponseData{}
	return &this
}

// GetType returns the Type field value
func (o *PostCampaignRecipientEstimationJobResponseData) GetType() CampaignRecipientEstimationJobEnum {
	if o == nil {
		var ret CampaignRecipientEstimationJobEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PostCampaignRecipientEstimationJobResponseData) GetTypeOk() (*CampaignRecipientEstimationJobEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PostCampaignRecipientEstimationJobResponseData) SetType(v CampaignRecipientEstimationJobEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *PostCampaignRecipientEstimationJobResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PostCampaignRecipientEstimationJobResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PostCampaignRecipientEstimationJobResponseData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *PostCampaignRecipientEstimationJobResponseData) GetAttributes() CampaignRecipientEstimationJobResponseObjectResourceAttributes {
	if o == nil {
		var ret CampaignRecipientEstimationJobResponseObjectResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PostCampaignRecipientEstimationJobResponseData) GetAttributesOk() (*CampaignRecipientEstimationJobResponseObjectResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PostCampaignRecipientEstimationJobResponseData) SetAttributes(v CampaignRecipientEstimationJobResponseObjectResourceAttributes) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *PostCampaignRecipientEstimationJobResponseData) GetLinks() ObjectLinks {
	if o == nil {
		var ret ObjectLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PostCampaignRecipientEstimationJobResponseData) GetLinksOk() (*ObjectLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *PostCampaignRecipientEstimationJobResponseData) SetLinks(v ObjectLinks) {
	o.Links = v
}

func (o PostCampaignRecipientEstimationJobResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCampaignRecipientEstimationJobResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["attributes"] = o.Attributes
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *PostCampaignRecipientEstimationJobResponseData) UnmarshalJSON(data []byte) (err error) {
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

	varPostCampaignRecipientEstimationJobResponseData := _PostCampaignRecipientEstimationJobResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostCampaignRecipientEstimationJobResponseData)

	if err != nil {
		return err
	}

	*o = PostCampaignRecipientEstimationJobResponseData(varPostCampaignRecipientEstimationJobResponseData)

	return err
}

type NullablePostCampaignRecipientEstimationJobResponseData struct {
	value *PostCampaignRecipientEstimationJobResponseData
	isSet bool
}

func (v NullablePostCampaignRecipientEstimationJobResponseData) Get() *PostCampaignRecipientEstimationJobResponseData {
	return v.value
}

func (v *NullablePostCampaignRecipientEstimationJobResponseData) Set(val *PostCampaignRecipientEstimationJobResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCampaignRecipientEstimationJobResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCampaignRecipientEstimationJobResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCampaignRecipientEstimationJobResponseData(val *PostCampaignRecipientEstimationJobResponseData) *NullablePostCampaignRecipientEstimationJobResponseData {
	return &NullablePostCampaignRecipientEstimationJobResponseData{value: val, isSet: true}
}

func (v NullablePostCampaignRecipientEstimationJobResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCampaignRecipientEstimationJobResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


