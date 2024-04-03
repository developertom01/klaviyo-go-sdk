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

// checks if the GetAccounts4XXResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccounts4XXResponse{}

// GetAccounts4XXResponse struct for GetAccounts4XXResponse
type GetAccounts4XXResponse struct {
	Errors []GetAccounts4XXResponseErrorsInner `json:"errors"`
}

type _GetAccounts4XXResponse GetAccounts4XXResponse

// NewGetAccounts4XXResponse instantiates a new GetAccounts4XXResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccounts4XXResponse(errors []GetAccounts4XXResponseErrorsInner) *GetAccounts4XXResponse {
	this := GetAccounts4XXResponse{}
	this.Errors = errors
	return &this
}

// NewGetAccounts4XXResponseWithDefaults instantiates a new GetAccounts4XXResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccounts4XXResponseWithDefaults() *GetAccounts4XXResponse {
	this := GetAccounts4XXResponse{}
	return &this
}

// GetErrors returns the Errors field value
func (o *GetAccounts4XXResponse) GetErrors() []GetAccounts4XXResponseErrorsInner {
	if o == nil {
		var ret []GetAccounts4XXResponseErrorsInner
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *GetAccounts4XXResponse) GetErrorsOk() ([]GetAccounts4XXResponseErrorsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *GetAccounts4XXResponse) SetErrors(v []GetAccounts4XXResponseErrorsInner) {
	o.Errors = v
}

func (o GetAccounts4XXResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccounts4XXResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errors"] = o.Errors
	return toSerialize, nil
}

func (o *GetAccounts4XXResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"errors",
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

	varGetAccounts4XXResponse := _GetAccounts4XXResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetAccounts4XXResponse)

	if err != nil {
		return err
	}

	*o = GetAccounts4XXResponse(varGetAccounts4XXResponse)

	return err
}

type NullableGetAccounts4XXResponse struct {
	value *GetAccounts4XXResponse
	isSet bool
}

func (v NullableGetAccounts4XXResponse) Get() *GetAccounts4XXResponse {
	return v.value
}

func (v *NullableGetAccounts4XXResponse) Set(val *GetAccounts4XXResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccounts4XXResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccounts4XXResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccounts4XXResponse(val *GetAccounts4XXResponse) *NullableGetAccounts4XXResponse {
	return &NullableGetAccounts4XXResponse{value: val, isSet: true}
}

func (v NullableGetAccounts4XXResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccounts4XXResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


