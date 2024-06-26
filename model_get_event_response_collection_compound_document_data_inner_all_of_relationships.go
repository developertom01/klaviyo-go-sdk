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

// checks if the GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships{}

// GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships struct for GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships
type GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships struct {
	Profile *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile `json:"profile,omitempty"`
	Metric *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsMetric `json:"metric,omitempty"`
	Attributions *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsAttributions `json:"attributions,omitempty"`
}

// NewGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships instantiates a new GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships() *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships {
	this := GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships{}
	return &this
}

// NewGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsWithDefaults instantiates a new GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsWithDefaults() *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships {
	this := GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships) GetProfile() GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile {
	if o == nil || IsNil(o.Profile) {
		var ret GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships) GetProfileOk() (*GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile and assigns it to the Profile field.
func (o *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships) SetProfile(v GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsProfile) {
	o.Profile = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships) GetMetric() GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsMetric {
	if o == nil || IsNil(o.Metric) {
		var ret GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsMetric
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships) GetMetricOk() (*GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsMetric, bool) {
	if o == nil || IsNil(o.Metric) {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships) HasMetric() bool {
	if o != nil && !IsNil(o.Metric) {
		return true
	}

	return false
}

// SetMetric gets a reference to the given GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsMetric and assigns it to the Metric field.
func (o *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships) SetMetric(v GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsMetric) {
	o.Metric = &v
}

// GetAttributions returns the Attributions field value if set, zero value otherwise.
func (o *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships) GetAttributions() GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsAttributions {
	if o == nil || IsNil(o.Attributions) {
		var ret GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsAttributions
		return ret
	}
	return *o.Attributions
}

// GetAttributionsOk returns a tuple with the Attributions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships) GetAttributionsOk() (*GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsAttributions, bool) {
	if o == nil || IsNil(o.Attributions) {
		return nil, false
	}
	return o.Attributions, true
}

// HasAttributions returns a boolean if a field has been set.
func (o *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships) HasAttributions() bool {
	if o != nil && !IsNil(o.Attributions) {
		return true
	}

	return false
}

// SetAttributions gets a reference to the given GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsAttributions and assigns it to the Attributions field.
func (o *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships) SetAttributions(v GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationshipsAttributions) {
	o.Attributions = &v
}

func (o GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.Metric) {
		toSerialize["metric"] = o.Metric
	}
	if !IsNil(o.Attributions) {
		toSerialize["attributions"] = o.Attributions
	}
	return toSerialize, nil
}

type NullableGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships struct {
	value *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships
	isSet bool
}

func (v NullableGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships) Get() *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships {
	return v.value
}

func (v *NullableGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships) Set(val *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships(val *GetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships) *NullableGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships {
	return &NullableGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships{value: val, isSet: true}
}

func (v NullableGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEventResponseCollectionCompoundDocumentDataInnerAllOfRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


