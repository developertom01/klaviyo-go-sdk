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

// checks if the GetProfileResponseCompoundDocumentData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProfileResponseCompoundDocumentData{}

// GetProfileResponseCompoundDocumentData struct for GetProfileResponseCompoundDocumentData
type GetProfileResponseCompoundDocumentData struct {
	Type ProfileEnum `json:"type"`
	// Primary key that uniquely identifies this profile. Generated by Klaviyo.
	Id NullableString `json:"id,omitempty"`
	Attributes GetProfileResponseDataAllOfAttributes `json:"attributes"`
	Links ObjectLinks `json:"links"`
	Relationships *GetProfileResponseCompoundDocumentDataAllOfRelationships `json:"relationships,omitempty"`
}

type _GetProfileResponseCompoundDocumentData GetProfileResponseCompoundDocumentData

// NewGetProfileResponseCompoundDocumentData instantiates a new GetProfileResponseCompoundDocumentData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProfileResponseCompoundDocumentData(type_ ProfileEnum, attributes GetProfileResponseDataAllOfAttributes, links ObjectLinks) *GetProfileResponseCompoundDocumentData {
	this := GetProfileResponseCompoundDocumentData{}
	this.Type = type_
	this.Attributes = attributes
	this.Links = links
	return &this
}

// NewGetProfileResponseCompoundDocumentDataWithDefaults instantiates a new GetProfileResponseCompoundDocumentData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProfileResponseCompoundDocumentDataWithDefaults() *GetProfileResponseCompoundDocumentData {
	this := GetProfileResponseCompoundDocumentData{}
	return &this
}

// GetType returns the Type field value
func (o *GetProfileResponseCompoundDocumentData) GetType() ProfileEnum {
	if o == nil {
		var ret ProfileEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetProfileResponseCompoundDocumentData) GetTypeOk() (*ProfileEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetProfileResponseCompoundDocumentData) SetType(v ProfileEnum) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProfileResponseCompoundDocumentData) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProfileResponseCompoundDocumentData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *GetProfileResponseCompoundDocumentData) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *GetProfileResponseCompoundDocumentData) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *GetProfileResponseCompoundDocumentData) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *GetProfileResponseCompoundDocumentData) UnsetId() {
	o.Id.Unset()
}

// GetAttributes returns the Attributes field value
func (o *GetProfileResponseCompoundDocumentData) GetAttributes() GetProfileResponseDataAllOfAttributes {
	if o == nil {
		var ret GetProfileResponseDataAllOfAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GetProfileResponseCompoundDocumentData) GetAttributesOk() (*GetProfileResponseDataAllOfAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GetProfileResponseCompoundDocumentData) SetAttributes(v GetProfileResponseDataAllOfAttributes) {
	o.Attributes = v
}

// GetLinks returns the Links field value
func (o *GetProfileResponseCompoundDocumentData) GetLinks() ObjectLinks {
	if o == nil {
		var ret ObjectLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GetProfileResponseCompoundDocumentData) GetLinksOk() (*ObjectLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *GetProfileResponseCompoundDocumentData) SetLinks(v ObjectLinks) {
	o.Links = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GetProfileResponseCompoundDocumentData) GetRelationships() GetProfileResponseCompoundDocumentDataAllOfRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret GetProfileResponseCompoundDocumentDataAllOfRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProfileResponseCompoundDocumentData) GetRelationshipsOk() (*GetProfileResponseCompoundDocumentDataAllOfRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GetProfileResponseCompoundDocumentData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GetProfileResponseCompoundDocumentDataAllOfRelationships and assigns it to the Relationships field.
func (o *GetProfileResponseCompoundDocumentData) SetRelationships(v GetProfileResponseCompoundDocumentDataAllOfRelationships) {
	o.Relationships = &v
}

func (o GetProfileResponseCompoundDocumentData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProfileResponseCompoundDocumentData) ToMap() (map[string]interface{}, error) {
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

func (o *GetProfileResponseCompoundDocumentData) UnmarshalJSON(data []byte) (err error) {
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

	varGetProfileResponseCompoundDocumentData := _GetProfileResponseCompoundDocumentData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetProfileResponseCompoundDocumentData)

	if err != nil {
		return err
	}

	*o = GetProfileResponseCompoundDocumentData(varGetProfileResponseCompoundDocumentData)

	return err
}

type NullableGetProfileResponseCompoundDocumentData struct {
	value *GetProfileResponseCompoundDocumentData
	isSet bool
}

func (v NullableGetProfileResponseCompoundDocumentData) Get() *GetProfileResponseCompoundDocumentData {
	return v.value
}

func (v *NullableGetProfileResponseCompoundDocumentData) Set(val *GetProfileResponseCompoundDocumentData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProfileResponseCompoundDocumentData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProfileResponseCompoundDocumentData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProfileResponseCompoundDocumentData(val *GetProfileResponseCompoundDocumentData) *NullableGetProfileResponseCompoundDocumentData {
	return &NullableGetProfileResponseCompoundDocumentData{value: val, isSet: true}
}

func (v NullableGetProfileResponseCompoundDocumentData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProfileResponseCompoundDocumentData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

