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

// checks if the GetCampaignResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCampaignResponseData{}

// GetCampaignResponseData struct for GetCampaignResponseData
type GetCampaignResponseData struct {
	Type CampaignEnum `json:"type"`
	// The campaign ID
	Id string `json:"id"`
	Attributes CampaignResponseObjectResourceAttributes `json:"attributes"`
	Links ObjectLinks `json:"links"`
	Relationships *GetCampaignResponseDataAllOfRelationships `json:"relationships,omitempty"`
}

type _GetCampaignResponseData GetCampaignResponseData

// NewGetCampaignResponseData instantiates a new GetCampaignResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCampaignResponseData(type_ CampaignEnum, id string, attributes CampaignResponseObjectResourceAttributes, links ObjectLinks) *GetCampaignResponseData {
	this := GetCampaignResponseData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewGetCampaignResponseDataWithDefaults instantiates a new GetCampaignResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCampaignResponseDataWithDefaults() *GetCampaignResponseData {
	this := GetCampaignResponseData{}
	return &this
}

// GetType returns the Type field value
func (o *GetCampaignResponseData) GetType() CampaignEnum {
	if o == nil {
		var ret CampaignEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetCampaignResponseData) GetTypeOk() (*CampaignEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetCampaignResponseData) SetType(v CampaignEnum) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GetCampaignResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetCampaignResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetCampaignResponseData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *GetCampaignResponseData) GetAttributes() CampaignResponseObjectResourceAttributes {
	if o == nil {
		var ret CampaignResponseObjectResourceAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GetCampaignResponseData) GetAttributesOk() (*CampaignResponseObjectResourceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GetCampaignResponseData) SetAttributes(v CampaignResponseObjectResourceAttributes) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *GetCampaignResponseData) GetLinks() ObjectLinks {
	if o == nil {
		var ret ObjectLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GetCampaignResponseData) GetLinksOk() (*ObjectLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GetCampaignResponseData) SetLinks(v ObjectLinks) {
	o.Links = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GetCampaignResponseData) GetRelationships() GetCampaignResponseDataAllOfRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret GetCampaignResponseDataAllOfRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCampaignResponseData) GetRelationshipsOk() (*GetCampaignResponseDataAllOfRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GetCampaignResponseData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GetCampaignResponseDataAllOfRelationships and assigns it to the Relationships field.
func (o *GetCampaignResponseData) SetRelationships(v GetCampaignResponseDataAllOfRelationships) {
	o.Relationships = &v
}

func (o GetCampaignResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCampaignResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["attributes"] = o.Attributes
	toSerialize["links"] = o.Links
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

func (o *GetCampaignResponseData) UnmarshalJSON(data []byte) (err error) {
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

	varGetCampaignResponseData := _GetCampaignResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetCampaignResponseData)

	if err != nil {
		return err
	}

	*o = GetCampaignResponseData(varGetCampaignResponseData)

	return err
}

type NullableGetCampaignResponseData struct {
	value *GetCampaignResponseData
	isSet bool
}

func (v NullableGetCampaignResponseData) Get() *GetCampaignResponseData {
	return v.value
}

func (v *NullableGetCampaignResponseData) Set(val *GetCampaignResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCampaignResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCampaignResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCampaignResponseData(val *GetCampaignResponseData) *NullableGetCampaignResponseData {
	return &NullableGetCampaignResponseData{value: val, isSet: true}
}

func (v NullableGetCampaignResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCampaignResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


