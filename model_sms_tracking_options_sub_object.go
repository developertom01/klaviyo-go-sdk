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

// checks if the SMSTrackingOptionsSubObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SMSTrackingOptionsSubObject{}

// SMSTrackingOptionsSubObject struct for SMSTrackingOptionsSubObject
type SMSTrackingOptionsSubObject struct {
	// Whether the campaign needs UTM parameters. If set to False, UTM params will not be used.
	IsAddUtm NullableBool `json:"is_add_utm,omitempty"`
	// A list of UTM parameters. If an empty list is given and is_add_utm is True, uses company defaults.
	UtmParams []UTMParamsSubObject `json:"utm_params,omitempty"`
}

// NewSMSTrackingOptionsSubObject instantiates a new SMSTrackingOptionsSubObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSMSTrackingOptionsSubObject() *SMSTrackingOptionsSubObject {
	this := SMSTrackingOptionsSubObject{}
	return &this
}

// NewSMSTrackingOptionsSubObjectWithDefaults instantiates a new SMSTrackingOptionsSubObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSMSTrackingOptionsSubObjectWithDefaults() *SMSTrackingOptionsSubObject {
	this := SMSTrackingOptionsSubObject{}
	return &this
}

// GetIsAddUtm returns the IsAddUtm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SMSTrackingOptionsSubObject) GetIsAddUtm() bool {
	if o == nil || IsNil(o.IsAddUtm.Get()) {
		var ret bool
		return ret
	}
	return *o.IsAddUtm.Get()
}

// GetIsAddUtmOk returns a tuple with the IsAddUtm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SMSTrackingOptionsSubObject) GetIsAddUtmOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsAddUtm.Get(), o.IsAddUtm.IsSet()
}

// HasIsAddUtm returns a boolean if a field has been set.
func (o *SMSTrackingOptionsSubObject) HasIsAddUtm() bool {
	if o != nil && o.IsAddUtm.IsSet() {
		return true
	}

	return false
}

// SetIsAddUtm gets a reference to the given NullableBool and assigns it to the IsAddUtm field.
func (o *SMSTrackingOptionsSubObject) SetIsAddUtm(v bool) {
	o.IsAddUtm.Set(&v)
}
// SetIsAddUtmNil sets the value for IsAddUtm to be an explicit nil
func (o *SMSTrackingOptionsSubObject) SetIsAddUtmNil() {
	o.IsAddUtm.Set(nil)
}

// UnsetIsAddUtm ensures that no value is present for IsAddUtm, not even an explicit nil
func (o *SMSTrackingOptionsSubObject) UnsetIsAddUtm() {
	o.IsAddUtm.Unset()
}

// GetUtmParams returns the UtmParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SMSTrackingOptionsSubObject) GetUtmParams() []UTMParamsSubObject {
	if o == nil {
		var ret []UTMParamsSubObject
		return ret
	}
	return o.UtmParams
}

// GetUtmParamsOk returns a tuple with the UtmParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SMSTrackingOptionsSubObject) GetUtmParamsOk() ([]UTMParamsSubObject, bool) {
	if o == nil || IsNil(o.UtmParams) {
		return nil, false
	}
	return o.UtmParams, true
}

// HasUtmParams returns a boolean if a field has been set.
func (o *SMSTrackingOptionsSubObject) HasUtmParams() bool {
	if o != nil && !IsNil(o.UtmParams) {
		return true
	}

	return false
}

// SetUtmParams gets a reference to the given []UTMParamsSubObject and assigns it to the UtmParams field.
func (o *SMSTrackingOptionsSubObject) SetUtmParams(v []UTMParamsSubObject) {
	o.UtmParams = v
}

func (o SMSTrackingOptionsSubObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SMSTrackingOptionsSubObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IsAddUtm.IsSet() {
		toSerialize["is_add_utm"] = o.IsAddUtm.Get()
	}
	if o.UtmParams != nil {
		toSerialize["utm_params"] = o.UtmParams
	}
	return toSerialize, nil
}

type NullableSMSTrackingOptionsSubObject struct {
	value *SMSTrackingOptionsSubObject
	isSet bool
}

func (v NullableSMSTrackingOptionsSubObject) Get() *SMSTrackingOptionsSubObject {
	return v.value
}

func (v *NullableSMSTrackingOptionsSubObject) Set(val *SMSTrackingOptionsSubObject) {
	v.value = val
	v.isSet = true
}

func (v NullableSMSTrackingOptionsSubObject) IsSet() bool {
	return v.isSet
}

func (v *NullableSMSTrackingOptionsSubObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSMSTrackingOptionsSubObject(val *SMSTrackingOptionsSubObject) *NullableSMSTrackingOptionsSubObject {
	return &NullableSMSTrackingOptionsSubObject{value: val, isSet: true}
}

func (v NullableSMSTrackingOptionsSubObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSMSTrackingOptionsSubObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

