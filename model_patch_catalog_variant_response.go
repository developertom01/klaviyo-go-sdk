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

// checks if the PatchCatalogVariantResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchCatalogVariantResponse{}

// PatchCatalogVariantResponse struct for PatchCatalogVariantResponse
type PatchCatalogVariantResponse struct {
	Data PostCatalogVariantResponseData `json:"data"`
}

type _PatchCatalogVariantResponse PatchCatalogVariantResponse

// NewPatchCatalogVariantResponse instantiates a new PatchCatalogVariantResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchCatalogVariantResponse(data PostCatalogVariantResponseData) *PatchCatalogVariantResponse {
	this := PatchCatalogVariantResponse{}
	this.Data = data
	return &this
}

// NewPatchCatalogVariantResponseWithDefaults instantiates a new PatchCatalogVariantResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchCatalogVariantResponseWithDefaults() *PatchCatalogVariantResponse {
	this := PatchCatalogVariantResponse{}
	return &this
}

// GetData returns the Data field value
func (o *PatchCatalogVariantResponse) GetData() PostCatalogVariantResponseData {
	if o == nil {
		var ret PostCatalogVariantResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PatchCatalogVariantResponse) GetDataOk() (*PostCatalogVariantResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PatchCatalogVariantResponse) SetData(v PostCatalogVariantResponseData) {
	o.Data = v
}

func (o PatchCatalogVariantResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchCatalogVariantResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *PatchCatalogVariantResponse) UnmarshalJSON(data []byte) (err error) {
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

	varPatchCatalogVariantResponse := _PatchCatalogVariantResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPatchCatalogVariantResponse)

	if err != nil {
		return err
	}

	*o = PatchCatalogVariantResponse(varPatchCatalogVariantResponse)

	return err
}

type NullablePatchCatalogVariantResponse struct {
	value *PatchCatalogVariantResponse
	isSet bool
}

func (v NullablePatchCatalogVariantResponse) Get() *PatchCatalogVariantResponse {
	return v.value
}

func (v *NullablePatchCatalogVariantResponse) Set(val *PatchCatalogVariantResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchCatalogVariantResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchCatalogVariantResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchCatalogVariantResponse(val *PatchCatalogVariantResponse) *NullablePatchCatalogVariantResponse {
	return &NullablePatchCatalogVariantResponse{value: val, isSet: true}
}

func (v NullablePatchCatalogVariantResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchCatalogVariantResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


