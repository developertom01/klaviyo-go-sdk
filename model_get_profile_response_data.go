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

// checks if the GetProfileResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProfileResponseData{}

// GetProfileResponseData struct for GetProfileResponseData
type GetProfileResponseData struct {
	Type ProfileEnum `json:"type"`
	// Primary key that uniquely identifies this profile. Generated by Klaviyo.
	Id NullableString `json:"id,omitempty"`
	Attributes GetProfileResponseDataAllOfAttributes `json:"attributes"`
	Links ObjectLinks `json:"links"`
	Relationships *GetProfileResponseDataAllOfRelationships `json:"relationships,omitempty"`
}

type _GetProfileResponseData GetProfileResponseData

// NewGetProfileResponseData instantiates a new GetProfileResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProfileResponseData(type_ ProfileEnum, attributes GetProfileResponseDataAllOfAttributes, links ObjectLinks) *GetProfileResponseData {
	this := GetProfileResponseData{}
	this.Type = type_
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewGetProfileResponseDataWithDefaults instantiates a new GetProfileResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProfileResponseDataWithDefaults() *GetProfileResponseData {
	this := GetProfileResponseData{}
	return &this
}

// GetType returns the Type field value
func (o *GetProfileResponseData) GetType() ProfileEnum {
	if o == nil {
		var ret ProfileEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetProfileResponseData) GetTypeOk() (*ProfileEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetProfileResponseData) SetType(v ProfileEnum) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProfileResponseData) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProfileResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *GetProfileResponseData) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *GetProfileResponseData) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *GetProfileResponseData) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *GetProfileResponseData) UnsetId() {
	o.Id.Unset()
}

// GetAttributes returns the Attributes field value
func (o *GetProfileResponseData) GetAttributes() GetProfileResponseDataAllOfAttributes {
	if o == nil {
		var ret GetProfileResponseDataAllOfAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GetProfileResponseData) GetAttributesOk() (*GetProfileResponseDataAllOfAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GetProfileResponseData) SetAttributes(v GetProfileResponseDataAllOfAttributes) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *GetProfileResponseData) GetLinks() ObjectLinks {
	if o == nil {
		var ret ObjectLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GetProfileResponseData) GetLinksOk() (*ObjectLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GetProfileResponseData) SetLinks(v ObjectLinks) {
	o.Links = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GetProfileResponseData) GetRelationships() GetProfileResponseDataAllOfRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret GetProfileResponseDataAllOfRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProfileResponseData) GetRelationshipsOk() (*GetProfileResponseDataAllOfRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GetProfileResponseData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GetProfileResponseDataAllOfRelationships and assigns it to the Relationships field.
func (o *GetProfileResponseData) SetRelationships(v GetProfileResponseDataAllOfRelationships) {
	o.Relationships = &v
}

func (o GetProfileResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProfileResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["links"] = o.Links
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

func (o *GetProfileResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varGetProfileResponseData := _GetProfileResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetProfileResponseData)

	if err != nil {
		return err
	}

	*o = GetProfileResponseData(varGetProfileResponseData)

	return err
}

type NullableGetProfileResponseData struct {
	value *GetProfileResponseData
	isSet bool
}

func (v NullableGetProfileResponseData) Get() *GetProfileResponseData {
	return v.value
}

func (v *NullableGetProfileResponseData) Set(val *GetProfileResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProfileResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProfileResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProfileResponseData(val *GetProfileResponseData) *NullableGetProfileResponseData {
	return &NullableGetProfileResponseData{value: val, isSet: true}
}

func (v NullableGetProfileResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProfileResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

