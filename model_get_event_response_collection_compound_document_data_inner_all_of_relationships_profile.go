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

// checks if the GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile{}

// GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile struct for GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile
type GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile struct {
	Data GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfileData `json:"data"`
	Links *RelationshipLinks `json:"links,omitempty"`
}

type _GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile

// NewGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile instantiates a new GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile(data GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfileData) *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile {
	this := GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile{}
	this.Data = data
	return &this
}

// NewGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfileWithDefaults instantiates a new GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfileWithDefaults() *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile {
	this := GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile{}
	return &this
}

// GetData returns the Data field value
func (o *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile) GetData() GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfileData {
	if o == nil {
		var ret GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfileData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile) GetDataOk() (*GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfileData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile) SetData(v GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfileData) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

func (o GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile := _GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile)

	if err != nil {
		return err
	}

	*o = GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile(varGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile)

	return err
}

type NullableGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile struct {
	value *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile
	isSet bool
}

func (v NullableGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile) Get() *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile {
	return v.value
}

func (v *NullableGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile) Set(val *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile(val *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile) *NullableGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile {
	return &NullableGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile{value: val, isSet: true}
}

func (v NullableGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


