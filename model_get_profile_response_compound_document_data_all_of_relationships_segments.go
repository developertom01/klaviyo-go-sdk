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

// checks if the GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments{}

// GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments struct for GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments
type GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments struct {
	Data []GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegmentsDataInner `json:"data"`
	Links *RelationshipLinks `json:"links,omitempty"`
}

type _GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments

// NewGetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments instantiates a new GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments(data []GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegmentsDataInner) *GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments {
	this := GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments{}
	this.Data = data
	return &this
}

// NewGetProfileResponseCompoundDocumentDataAllOfRelationshipsSegmentsWithDefaults instantiates a new GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProfileResponseCompoundDocumentDataAllOfRelationshipsSegmentsWithDefaults() *GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments {
	this := GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments{}
	return &this
}

// GetData returns the Data field value
func (o *GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments) GetData() []GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegmentsDataInner {
	if o == nil {
		var ret []GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegmentsDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments) GetDataOk() ([]GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegmentsDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments) SetData(v []GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegmentsDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments) GetLinks() RelationshipLinks {
	if o == nil || IsNil(o.Links) {
		var ret RelationshipLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments) GetLinksOk() (*RelationshipLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RelationshipLinks and assigns it to the Links field.
func (o *GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments) SetLinks(v RelationshipLinks) {
	o.Links = &v
}

func (o GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments) UnmarshalJSON(data []byte) (err error) {
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

	varGetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments := _GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments)

	if err != nil {
		return err
	}

	*o = GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments(varGetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments)

	return err
}

type NullableGetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments struct {
	value *GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments
	isSet bool
}

func (v NullableGetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments) Get() *GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments {
	return v.value
}

func (v *NullableGetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments) Set(val *GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments(val *GetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments) *NullableGetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments {
	return &NullableGetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments{value: val, isSet: true}
}

func (v NullableGetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProfileResponseCompoundDocumentDataAllOfRelationshipsSegments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

