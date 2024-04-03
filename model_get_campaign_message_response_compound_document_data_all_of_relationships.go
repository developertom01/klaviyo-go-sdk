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
)

// checks if the GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships{}

// GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships struct for GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships
type GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships struct {
	Campaign *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaign `json:"campaign,omitempty"`
	Template *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplate `json:"template,omitempty"`
}

// NewGetCampaignMessageResponseCompoundDocumentDataAllOfRelationships instantiates a new GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCampaignMessageResponseCompoundDocumentDataAllOfRelationships() *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships {
	this := GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships{}
	return &this
}

// NewGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsWithDefaults instantiates a new GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsWithDefaults() *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships {
	this := GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships{}
	return &this
}

// GetCampaign returns the Campaign field value if set, zero value otherwise.
func (o *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships) GetCampaign() GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaign {
	if o == nil || IsNil(o.Campaign) {
		var ret GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaign
		return ret
	}
	return *o.Campaign
}

// GetCampaignOk returns a tuple with the Campaign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships) GetCampaignOk() (*GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaign, bool) {
	if o == nil || IsNil(o.Campaign) {
		return nil, false
	}
	return o.Campaign, true
}

// HasCampaign returns a boolean if a field has been set.
func (o *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships) HasCampaign() bool {
	if o != nil && !IsNil(o.Campaign) {
		return true
	}

	return false
}

// SetCampaign gets a reference to the given GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaign and assigns it to the Campaign field.
func (o *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships) SetCampaign(v GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsCampaign) {
	o.Campaign = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships) GetTemplate() GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplate {
	if o == nil || IsNil(o.Template) {
		var ret GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplate
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships) GetTemplateOk() (*GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplate, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplate and assigns it to the Template field.
func (o *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships) SetTemplate(v GetCampaignMessageResponseCompoundDocumentDataAllOfRelationshipsTemplate) {
	o.Template = &v
}

func (o GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Campaign) {
		toSerialize["campaign"] = o.Campaign
	}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	return toSerialize, nil
}

type NullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationships struct {
	value *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships
	isSet bool
}

func (v NullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationships) Get() *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships {
	return v.value
}

func (v *NullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationships) Set(val *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationships(val *GetCampaignMessageResponseCompoundDocumentDataAllOfRelationships) *NullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationships {
	return &NullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationships{value: val, isSet: true}
}

func (v NullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCampaignMessageResponseCompoundDocumentDataAllOfRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


