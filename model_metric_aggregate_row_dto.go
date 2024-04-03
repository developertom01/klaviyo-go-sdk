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

// checks if the MetricAggregateRowDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricAggregateRowDTO{}

// MetricAggregateRowDTO struct for MetricAggregateRowDTO
type MetricAggregateRowDTO struct {
	// List of dimensions associated with this set of measurements
	Dimensions []string `json:"dimensions"`
	// Dictionary of measurement_key, values
	Measurements map[string]interface{} `json:"measurements"`
}

type _MetricAggregateRowDTO MetricAggregateRowDTO

// NewMetricAggregateRowDTO instantiates a new MetricAggregateRowDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricAggregateRowDTO(dimensions []string, measurements map[string]interface{}) *MetricAggregateRowDTO {
	this := MetricAggregateRowDTO{}
	this.Dimensions = dimensions
	this.Measurements = measurements
	return &this
}

// NewMetricAggregateRowDTOWithDefaults instantiates a new MetricAggregateRowDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricAggregateRowDTOWithDefaults() *MetricAggregateRowDTO {
	this := MetricAggregateRowDTO{}
	return &this
}

// GetDimensions returns the Dimensions field value
func (o *MetricAggregateRowDTO) GetDimensions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value
// and a boolean to check if the value has been set.
func (o *MetricAggregateRowDTO) GetDimensionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dimensions, true
}

// SetDimensions sets field value
func (o *MetricAggregateRowDTO) SetDimensions(v []string) {
	o.Dimensions = v
}

// GetMeasurements returns the Measurements field value
func (o *MetricAggregateRowDTO) GetMeasurements() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Measurements
}

// GetMeasurementsOk returns a tuple with the Measurements field value
// and a boolean to check if the value has been set.
func (o *MetricAggregateRowDTO) GetMeasurementsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Measurements, true
}

// SetMeasurements sets field value
func (o *MetricAggregateRowDTO) SetMeasurements(v map[string]interface{}) {
	o.Measurements = v
}

func (o MetricAggregateRowDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricAggregateRowDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dimensions"] = o.Dimensions
	toSerialize["measurements"] = o.Measurements
	return toSerialize, nil
}

func (o *MetricAggregateRowDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dimensions",
		"measurements",
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

	varMetricAggregateRowDTO := _MetricAggregateRowDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMetricAggregateRowDTO)

	if err != nil {
		return err
	}

	*o = MetricAggregateRowDTO(varMetricAggregateRowDTO)

	return err
}

type NullableMetricAggregateRowDTO struct {
	value *MetricAggregateRowDTO
	isSet bool
}

func (v NullableMetricAggregateRowDTO) Get() *MetricAggregateRowDTO {
	return v.value
}

func (v *NullableMetricAggregateRowDTO) Set(val *MetricAggregateRowDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricAggregateRowDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricAggregateRowDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricAggregateRowDTO(val *MetricAggregateRowDTO) *NullableMetricAggregateRowDTO {
	return &NullableMetricAggregateRowDTO{value: val, isSet: true}
}

func (v NullableMetricAggregateRowDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricAggregateRowDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


