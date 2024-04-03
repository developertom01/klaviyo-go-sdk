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

// checks if the GetFlowActionResponseCompoundDocument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFlowActionResponseCompoundDocument{}

// GetFlowActionResponseCompoundDocument struct for GetFlowActionResponseCompoundDocument
type GetFlowActionResponseCompoundDocument struct {
	Data GetFlowActionResponseCompoundDocumentData `json:"data"`
	Included []GetFlowActionResponseCompoundDocumentIncludedInner `json:"included,omitempty"`
}

type _GetFlowActionResponseCompoundDocument GetFlowActionResponseCompoundDocument

// NewGetFlowActionResponseCompoundDocument instantiates a new GetFlowActionResponseCompoundDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlowActionResponseCompoundDocument(data GetFlowActionResponseCompoundDocumentData) *GetFlowActionResponseCompoundDocument {
	this := GetFlowActionResponseCompoundDocument{}
	this.Data = data
	return &this
}

// NewGetFlowActionResponseCompoundDocumentWithDefaults instantiates a new GetFlowActionResponseCompoundDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlowActionResponseCompoundDocumentWithDefaults() *GetFlowActionResponseCompoundDocument {
	this := GetFlowActionResponseCompoundDocument{}
	return &this
}

// GetData returns the Data field value
func (o *GetFlowActionResponseCompoundDocument) GetData() GetFlowActionResponseCompoundDocumentData {
	if o == nil {
		var ret GetFlowActionResponseCompoundDocumentData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetFlowActionResponseCompoundDocument) GetDataOk() (*GetFlowActionResponseCompoundDocumentData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetFlowActionResponseCompoundDocument) SetData(v GetFlowActionResponseCompoundDocumentData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *GetFlowActionResponseCompoundDocument) GetIncluded() []GetFlowActionResponseCompoundDocumentIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []GetFlowActionResponseCompoundDocumentIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlowActionResponseCompoundDocument) GetIncludedOk() ([]GetFlowActionResponseCompoundDocumentIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *GetFlowActionResponseCompoundDocument) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []GetFlowActionResponseCompoundDocumentIncludedInner and assigns it to the Included field.
func (o *GetFlowActionResponseCompoundDocument) SetIncluded(v []GetFlowActionResponseCompoundDocumentIncludedInner) {
	o.Included = v
}

func (o GetFlowActionResponseCompoundDocument) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlowActionResponseCompoundDocument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	return toSerialize, nil
}

func (o *GetFlowActionResponseCompoundDocument) UnmarshalJSON(data []byte) (err error) {
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

	varGetFlowActionResponseCompoundDocument := _GetFlowActionResponseCompoundDocument{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetFlowActionResponseCompoundDocument)

	if err != nil {
		return err
	}

	*o = GetFlowActionResponseCompoundDocument(varGetFlowActionResponseCompoundDocument)

	return err
}

type NullableGetFlowActionResponseCompoundDocument struct {
	value *GetFlowActionResponseCompoundDocument
	isSet bool
}

func (v NullableGetFlowActionResponseCompoundDocument) Get() *GetFlowActionResponseCompoundDocument {
	return v.value
}

func (v *NullableGetFlowActionResponseCompoundDocument) Set(val *GetFlowActionResponseCompoundDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlowActionResponseCompoundDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlowActionResponseCompoundDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlowActionResponseCompoundDocument(val *GetFlowActionResponseCompoundDocument) *NullableGetFlowActionResponseCompoundDocument {
	return &NullableGetFlowActionResponseCompoundDocument{value: val, isSet: true}
}

func (v NullableGetFlowActionResponseCompoundDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlowActionResponseCompoundDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


