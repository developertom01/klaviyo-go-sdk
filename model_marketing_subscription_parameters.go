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
	"time"
	"bytes"
	"fmt"
)

// checks if the MarketingSubscriptionParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarketingSubscriptionParameters{}

// MarketingSubscriptionParameters struct for MarketingSubscriptionParameters
type MarketingSubscriptionParameters struct {
	// The Consent status to subscribe to for the \"Marketing\" type. Currently supports \"SUBSCRIBED\".
	Consent string `json:"consent"`
	// The timestamp of when the profile's consent was gathered
	ConsentedAt NullableTime `json:"consented_at,omitempty"`
}

type _MarketingSubscriptionParameters MarketingSubscriptionParameters

// NewMarketingSubscriptionParameters instantiates a new MarketingSubscriptionParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketingSubscriptionParameters(consent string) *MarketingSubscriptionParameters {
	this := MarketingSubscriptionParameters{}
	this.Consent = consent
	return &this
}

// NewMarketingSubscriptionParametersWithDefaults instantiates a new MarketingSubscriptionParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketingSubscriptionParametersWithDefaults() *MarketingSubscriptionParameters {
	this := MarketingSubscriptionParameters{}
	return &this
}

// GetConsent returns the Consent field value
func (o *MarketingSubscriptionParameters) GetConsent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Consent
}

// GetConsentOk returns a tuple with the Consent field value
// and a boolean to check if the value has been set.
func (o *MarketingSubscriptionParameters) GetConsentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Consent, true
}

// SetConsent sets field value
func (o *MarketingSubscriptionParameters) SetConsent(v string) {
	o.Consent = v
}

// GetConsentedAt returns the ConsentedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MarketingSubscriptionParameters) GetConsentedAt() time.Time {
	if o == nil || IsNil(o.ConsentedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ConsentedAt.Get()
}

// GetConsentedAtOk returns a tuple with the ConsentedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MarketingSubscriptionParameters) GetConsentedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConsentedAt.Get(), o.ConsentedAt.IsSet()
}

// HasConsentedAt returns a boolean if a field has been set.
func (o *MarketingSubscriptionParameters) HasConsentedAt() bool {
	if o != nil && o.ConsentedAt.IsSet() {
		return true
	}

	return false
}

// SetConsentedAt gets a reference to the given NullableTime and assigns it to the ConsentedAt field.
func (o *MarketingSubscriptionParameters) SetConsentedAt(v time.Time) {
	o.ConsentedAt.Set(&v)
}
// SetConsentedAtNil sets the value for ConsentedAt to be an explicit nil
func (o *MarketingSubscriptionParameters) SetConsentedAtNil() {
	o.ConsentedAt.Set(nil)
}

// UnsetConsentedAt ensures that no value is present for ConsentedAt, not even an explicit nil
func (o *MarketingSubscriptionParameters) UnsetConsentedAt() {
	o.ConsentedAt.Unset()
}

func (o MarketingSubscriptionParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketingSubscriptionParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["consent"] = o.Consent
	if o.ConsentedAt.IsSet() {
		toSerialize["consented_at"] = o.ConsentedAt.Get()
	}
	return toSerialize, nil
}

func (o *MarketingSubscriptionParameters) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"consent",
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

	varMarketingSubscriptionParameters := _MarketingSubscriptionParameters{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMarketingSubscriptionParameters)

	if err != nil {
		return err
	}

	*o = MarketingSubscriptionParameters(varMarketingSubscriptionParameters)

	return err
}

type NullableMarketingSubscriptionParameters struct {
	value *MarketingSubscriptionParameters
	isSet bool
}

func (v NullableMarketingSubscriptionParameters) Get() *MarketingSubscriptionParameters {
	return v.value
}

func (v *NullableMarketingSubscriptionParameters) Set(val *MarketingSubscriptionParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketingSubscriptionParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketingSubscriptionParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketingSubscriptionParameters(val *MarketingSubscriptionParameters) *NullableMarketingSubscriptionParameters {
	return &NullableMarketingSubscriptionParameters{value: val, isSet: true}
}

func (v NullableMarketingSubscriptionParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketingSubscriptionParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


