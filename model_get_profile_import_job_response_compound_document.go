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

// checks if the GetProfileImportJobResponseCompoundDocument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProfileImportJobResponseCompoundDocument{}

// GetProfileImportJobResponseCompoundDocument struct for GetProfileImportJobResponseCompoundDocument
type GetProfileImportJobResponseCompoundDocument struct {
	Data GetProfileImportJobResponseCollectionCompoundDocumentDataInner `json:"data"`
	Included []ListResponseObjectResource `json:"included,omitempty"`
}

type _GetProfileImportJobResponseCompoundDocument GetProfileImportJobResponseCompoundDocument

// NewGetProfileImportJobResponseCompoundDocument instantiates a new GetProfileImportJobResponseCompoundDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProfileImportJobResponseCompoundDocument(data GetProfileImportJobResponseCollectionCompoundDocumentDataInner) *GetProfileImportJobResponseCompoundDocument {
	this := GetProfileImportJobResponseCompoundDocument{}
	this.Data = data
	return &this
}

// NewGetProfileImportJobResponseCompoundDocumentWithDefaults instantiates a new GetProfileImportJobResponseCompoundDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProfileImportJobResponseCompoundDocumentWithDefaults() *GetProfileImportJobResponseCompoundDocument {
	this := GetProfileImportJobResponseCompoundDocument{}
	return &this
}

// GetData returns the Data field value
func (o *GetProfileImportJobResponseCompoundDocument) GetData() GetProfileImportJobResponseCollectionCompoundDocumentDataInner {
	if o == nil {
		var ret GetProfileImportJobResponseCollectionCompoundDocumentDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetProfileImportJobResponseCompoundDocument) GetDataOk() (*GetProfileImportJobResponseCollectionCompoundDocumentDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetProfileImportJobResponseCompoundDocument) SetData(v GetProfileImportJobResponseCollectionCompoundDocumentDataInner) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *GetProfileImportJobResponseCompoundDocument) GetIncluded() []ListResponseObjectResource {
	if o == nil || IsNil(o.Included) {
		var ret []ListResponseObjectResource
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProfileImportJobResponseCompoundDocument) GetIncludedOk() ([]ListResponseObjectResource, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *GetProfileImportJobResponseCompoundDocument) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []ListResponseObjectResource and assigns it to the Included field.
func (o *GetProfileImportJobResponseCompoundDocument) SetIncluded(v []ListResponseObjectResource) {
	o.Included = v
}

func (o GetProfileImportJobResponseCompoundDocument) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProfileImportJobResponseCompoundDocument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	return toSerialize, nil
}

func (o *GetProfileImportJobResponseCompoundDocument) UnmarshalJSON(data []byte) (err error) {
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

	varGetProfileImportJobResponseCompoundDocument := _GetProfileImportJobResponseCompoundDocument{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetProfileImportJobResponseCompoundDocument)

	if err != nil {
		return err
	}

	*o = GetProfileImportJobResponseCompoundDocument(varGetProfileImportJobResponseCompoundDocument)

	return err
}

type NullableGetProfileImportJobResponseCompoundDocument struct {
	value *GetProfileImportJobResponseCompoundDocument
	isSet bool
}

func (v NullableGetProfileImportJobResponseCompoundDocument) Get() *GetProfileImportJobResponseCompoundDocument {
	return v.value
}

func (v *NullableGetProfileImportJobResponseCompoundDocument) Set(val *GetProfileImportJobResponseCompoundDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProfileImportJobResponseCompoundDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProfileImportJobResponseCompoundDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProfileImportJobResponseCompoundDocument(val *GetProfileImportJobResponseCompoundDocument) *NullableGetProfileImportJobResponseCompoundDocument {
	return &NullableGetProfileImportJobResponseCompoundDocument{value: val, isSet: true}
}

func (v NullableGetProfileImportJobResponseCompoundDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProfileImportJobResponseCompoundDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

