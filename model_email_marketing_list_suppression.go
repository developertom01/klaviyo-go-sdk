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

// checks if the EmailMarketingListSuppression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailMarketingListSuppression{}

// EmailMarketingListSuppression struct for EmailMarketingListSuppression
type EmailMarketingListSuppression struct {
	// The ID of list to which the suppression applies.
	ListId string `json:"list_id"`
	// The reason the profile was suppressed from the list.
	Reason string `json:"reason"`
	// The timestamp when the profile was suppressed from the list, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm).
	Timestamp time.Time `json:"timestamp"`
}

type _EmailMarketingListSuppression EmailMarketingListSuppression

// NewEmailMarketingListSuppression instantiates a new EmailMarketingListSuppression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailMarketingListSuppression(listId string, reason string, timestamp time.Time) *EmailMarketingListSuppression {
	this := EmailMarketingListSuppression{}
	this.ListId = listId
	this.Reason = reason
	this.Timestamp = timestamp
	return &this
}

// NewEmailMarketingListSuppressionWithDefaults instantiates a new EmailMarketingListSuppression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailMarketingListSuppressionWithDefaults() *EmailMarketingListSuppression {
	this := EmailMarketingListSuppression{}
	return &this
}

// GetListId returns the ListId field value
func (o *EmailMarketingListSuppression) GetListId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ListId
}

// GetListIdOk returns a tuple with the ListId field value
// and a boolean to check if the value has been set.
func (o *EmailMarketingListSuppression) GetListIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ListId, true
}

// SetListId sets field value
func (o *EmailMarketingListSuppression) SetListId(v string) {
	o.ListId = v
}

// GetReason returns the Reason field value
func (o *EmailMarketingListSuppression) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *EmailMarketingListSuppression) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *EmailMarketingListSuppression) SetReason(v string) {
	o.Reason = v
}

// GetTimestamp returns the Timestamp field value
func (o *EmailMarketingListSuppression) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *EmailMarketingListSuppression) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *EmailMarketingListSuppression) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

func (o EmailMarketingListSuppression) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailMarketingListSuppression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["list_id"] = o.ListId
	toSerialize["reason"] = o.Reason
	toSerialize["timestamp"] = o.Timestamp
	return toSerialize, nil
}

func (o *EmailMarketingListSuppression) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"list_id",
		"reason",
		"timestamp",
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

	varEmailMarketingListSuppression := _EmailMarketingListSuppression{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmailMarketingListSuppression)

	if err != nil {
		return err
	}

	*o = EmailMarketingListSuppression(varEmailMarketingListSuppression)

	return err
}

type NullableEmailMarketingListSuppression struct {
	value *EmailMarketingListSuppression
	isSet bool
}

func (v NullableEmailMarketingListSuppression) Get() *EmailMarketingListSuppression {
	return v.value
}

func (v *NullableEmailMarketingListSuppression) Set(val *EmailMarketingListSuppression) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailMarketingListSuppression) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailMarketingListSuppression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailMarketingListSuppression(val *EmailMarketingListSuppression) *NullableEmailMarketingListSuppression {
	return &NullableEmailMarketingListSuppression{value: val, isSet: true}
}

func (v NullableEmailMarketingListSuppression) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailMarketingListSuppression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


