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

// checks if the MetricResponseObjectResourceAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricResponseObjectResourceAttributes{}

// MetricResponseObjectResourceAttributes struct for MetricResponseObjectResourceAttributes
type MetricResponseObjectResourceAttributes struct {
	// The name of the metric
	Name NullableString `json:"name,omitempty"`
	// Creation time in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm)
	Created NullableString `json:"created,omitempty"`
	// Last updated time in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm)
	Updated NullableString `json:"updated,omitempty"`
	// The integration associated with the event
	Integration map[string]interface{} `json:"integration,omitempty"`
}

// NewMetricResponseObjectResourceAttributes instantiates a new MetricResponseObjectResourceAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricResponseObjectResourceAttributes() *MetricResponseObjectResourceAttributes {
	this := MetricResponseObjectResourceAttributes{}
	return &this
}

// NewMetricResponseObjectResourceAttributesWithDefaults instantiates a new MetricResponseObjectResourceAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricResponseObjectResourceAttributesWithDefaults() *MetricResponseObjectResourceAttributes {
	this := MetricResponseObjectResourceAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricResponseObjectResourceAttributes) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetricResponseObjectResourceAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MetricResponseObjectResourceAttributes) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MetricResponseObjectResourceAttributes) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MetricResponseObjectResourceAttributes) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MetricResponseObjectResourceAttributes) UnsetName() {
	o.Name.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricResponseObjectResourceAttributes) GetCreated() string {
	if o == nil || IsNil(o.Created.Get()) {
		var ret string
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetricResponseObjectResourceAttributes) GetCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *MetricResponseObjectResourceAttributes) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableString and assigns it to the Created field.
func (o *MetricResponseObjectResourceAttributes) SetCreated(v string) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *MetricResponseObjectResourceAttributes) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *MetricResponseObjectResourceAttributes) UnsetCreated() {
	o.Created.Unset()
}

// GetUpdated returns the Updated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricResponseObjectResourceAttributes) GetUpdated() string {
	if o == nil || IsNil(o.Updated.Get()) {
		var ret string
		return ret
	}
	return *o.Updated.Get()
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetricResponseObjectResourceAttributes) GetUpdatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Updated.Get(), o.Updated.IsSet()
}

// HasUpdated returns a boolean if a field has been set.
func (o *MetricResponseObjectResourceAttributes) HasUpdated() bool {
	if o != nil && o.Updated.IsSet() {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given NullableString and assigns it to the Updated field.
func (o *MetricResponseObjectResourceAttributes) SetUpdated(v string) {
	o.Updated.Set(&v)
}
// SetUpdatedNil sets the value for Updated to be an explicit nil
func (o *MetricResponseObjectResourceAttributes) SetUpdatedNil() {
	o.Updated.Set(nil)
}

// UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
func (o *MetricResponseObjectResourceAttributes) UnsetUpdated() {
	o.Updated.Unset()
}

// GetIntegration returns the Integration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetricResponseObjectResourceAttributes) GetIntegration() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetricResponseObjectResourceAttributes) GetIntegrationOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Integration) {
		return map[string]interface{}{}, false
	}
	return o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *MetricResponseObjectResourceAttributes) HasIntegration() bool {
	if o != nil && !IsNil(o.Integration) {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given map[string]interface{} and assigns it to the Integration field.
func (o *MetricResponseObjectResourceAttributes) SetIntegration(v map[string]interface{}) {
	o.Integration = v
}

func (o MetricResponseObjectResourceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricResponseObjectResourceAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if o.Updated.IsSet() {
		toSerialize["updated"] = o.Updated.Get()
	}
	if o.Integration != nil {
		toSerialize["integration"] = o.Integration
	}
	return toSerialize, nil
}

type NullableMetricResponseObjectResourceAttributes struct {
	value *MetricResponseObjectResourceAttributes
	isSet bool
}

func (v NullableMetricResponseObjectResourceAttributes) Get() *MetricResponseObjectResourceAttributes {
	return v.value
}

func (v *NullableMetricResponseObjectResourceAttributes) Set(val *MetricResponseObjectResourceAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricResponseObjectResourceAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricResponseObjectResourceAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricResponseObjectResourceAttributes(val *MetricResponseObjectResourceAttributes) *NullableMetricResponseObjectResourceAttributes {
	return &NullableMetricResponseObjectResourceAttributes{value: val, isSet: true}
}

func (v NullableMetricResponseObjectResourceAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricResponseObjectResourceAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


