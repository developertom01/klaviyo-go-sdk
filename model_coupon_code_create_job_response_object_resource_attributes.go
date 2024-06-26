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

// checks if the CouponCodeCreateJobResponseObjectResourceAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CouponCodeCreateJobResponseObjectResourceAttributes{}

// CouponCodeCreateJobResponseObjectResourceAttributes struct for CouponCodeCreateJobResponseObjectResourceAttributes
type CouponCodeCreateJobResponseObjectResourceAttributes struct {
	// Status of the asynchronous job.
	Status string `json:"status"`
	// The date and time the job was created in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm).
	CreatedAt time.Time `json:"created_at"`
	// The total number of operations to be processed by the job. See `completed_count` for the job's current progress.
	TotalCount int32 `json:"total_count"`
	// The total number of operations that have been completed by the job.
	CompletedCount NullableInt32 `json:"completed_count,omitempty"`
	// The total number of operations that have failed as part of the job.
	FailedCount NullableInt32 `json:"failed_count,omitempty"`
	// Date and time the job was completed in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm).
	CompletedAt NullableTime `json:"completed_at,omitempty"`
	// Array of errors encountered during the processing of the job.
	Errors []APIJobErrorPayload `json:"errors,omitempty"`
	// Date and time the job expires in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm).
	ExpiresAt NullableTime `json:"expires_at,omitempty"`
}

type _CouponCodeCreateJobResponseObjectResourceAttributes CouponCodeCreateJobResponseObjectResourceAttributes

// NewCouponCodeCreateJobResponseObjectResourceAttributes instantiates a new CouponCodeCreateJobResponseObjectResourceAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponCodeCreateJobResponseObjectResourceAttributes(status string, createdAt time.Time, totalCount int32) *CouponCodeCreateJobResponseObjectResourceAttributes {
	this := CouponCodeCreateJobResponseObjectResourceAttributes{}
	this.Status = status
	this.CreatedAt = createdAt
	this.TotalCount = totalCount
	var completedCount int32 = 0
	this.CompletedCount = *NewNullableInt32(&completedCount)
	var failedCount int32 = 0
	this.FailedCount = *NewNullableInt32(&failedCount)
	return &this
}

// NewCouponCodeCreateJobResponseObjectResourceAttributesWithDefaults instantiates a new CouponCodeCreateJobResponseObjectResourceAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponCodeCreateJobResponseObjectResourceAttributesWithDefaults() *CouponCodeCreateJobResponseObjectResourceAttributes {
	this := CouponCodeCreateJobResponseObjectResourceAttributes{}
	var completedCount int32 = 0
	this.CompletedCount = *NewNullableInt32(&completedCount)
	var failedCount int32 = 0
	this.FailedCount = *NewNullableInt32(&failedCount)
	return &this
}

// GetStatus returns the Status field value
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) SetStatus(v string) {
	o.Status = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetTotalCount returns the TotalCount field value
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) SetTotalCount(v int32) {
	o.TotalCount = v
}

// GetCompletedCount returns the CompletedCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetCompletedCount() int32 {
	if o == nil || IsNil(o.CompletedCount.Get()) {
		var ret int32
		return ret
	}
	return *o.CompletedCount.Get()
}

// GetCompletedCountOk returns a tuple with the CompletedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetCompletedCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedCount.Get(), o.CompletedCount.IsSet()
}

// HasCompletedCount returns a boolean if a field has been set.
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) HasCompletedCount() bool {
	if o != nil && o.CompletedCount.IsSet() {
		return true
	}

	return false
}

// SetCompletedCount gets a reference to the given NullableInt32 and assigns it to the CompletedCount field.
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) SetCompletedCount(v int32) {
	o.CompletedCount.Set(&v)
}
// SetCompletedCountNil sets the value for CompletedCount to be an explicit nil
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) SetCompletedCountNil() {
	o.CompletedCount.Set(nil)
}

// UnsetCompletedCount ensures that no value is present for CompletedCount, not even an explicit nil
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) UnsetCompletedCount() {
	o.CompletedCount.Unset()
}

// GetFailedCount returns the FailedCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetFailedCount() int32 {
	if o == nil || IsNil(o.FailedCount.Get()) {
		var ret int32
		return ret
	}
	return *o.FailedCount.Get()
}

// GetFailedCountOk returns a tuple with the FailedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetFailedCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailedCount.Get(), o.FailedCount.IsSet()
}

// HasFailedCount returns a boolean if a field has been set.
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) HasFailedCount() bool {
	if o != nil && o.FailedCount.IsSet() {
		return true
	}

	return false
}

// SetFailedCount gets a reference to the given NullableInt32 and assigns it to the FailedCount field.
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) SetFailedCount(v int32) {
	o.FailedCount.Set(&v)
}
// SetFailedCountNil sets the value for FailedCount to be an explicit nil
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) SetFailedCountNil() {
	o.FailedCount.Set(nil)
}

// UnsetFailedCount ensures that no value is present for FailedCount, not even an explicit nil
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) UnsetFailedCount() {
	o.FailedCount.Unset()
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetCompletedAt() time.Time {
	if o == nil || IsNil(o.CompletedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) HasCompletedAt() bool {
	if o != nil && o.CompletedAt.IsSet() {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given NullableTime and assigns it to the CompletedAt field.
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) SetCompletedAt(v time.Time) {
	o.CompletedAt.Set(&v)
}
// SetCompletedAtNil sets the value for CompletedAt to be an explicit nil
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) SetCompletedAtNil() {
	o.CompletedAt.Set(nil)
}

// UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) UnsetCompletedAt() {
	o.CompletedAt.Unset()
}

// GetErrors returns the Errors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetErrors() []APIJobErrorPayload {
	if o == nil {
		var ret []APIJobErrorPayload
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetErrorsOk() ([]APIJobErrorPayload, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []APIJobErrorPayload and assigns it to the Errors field.
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) SetErrors(v []APIJobErrorPayload) {
	o.Errors = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given NullableTime and assigns it to the ExpiresAt field.
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) SetExpiresAt(v time.Time) {
	o.ExpiresAt.Set(&v)
}
// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
func (o *CouponCodeCreateJobResponseObjectResourceAttributes) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

func (o CouponCodeCreateJobResponseObjectResourceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CouponCodeCreateJobResponseObjectResourceAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["total_count"] = o.TotalCount
	if o.CompletedCount.IsSet() {
		toSerialize["completed_count"] = o.CompletedCount.Get()
	}
	if o.FailedCount.IsSet() {
		toSerialize["failed_count"] = o.FailedCount.Get()
	}
	if o.CompletedAt.IsSet() {
		toSerialize["completed_at"] = o.CompletedAt.Get()
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.ExpiresAt.IsSet() {
		toSerialize["expires_at"] = o.ExpiresAt.Get()
	}
	return toSerialize, nil
}

func (o *CouponCodeCreateJobResponseObjectResourceAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"created_at",
		"total_count",
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

	varCouponCodeCreateJobResponseObjectResourceAttributes := _CouponCodeCreateJobResponseObjectResourceAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCouponCodeCreateJobResponseObjectResourceAttributes)

	if err != nil {
		return err
	}

	*o = CouponCodeCreateJobResponseObjectResourceAttributes(varCouponCodeCreateJobResponseObjectResourceAttributes)

	return err
}

type NullableCouponCodeCreateJobResponseObjectResourceAttributes struct {
	value *CouponCodeCreateJobResponseObjectResourceAttributes
	isSet bool
}

func (v NullableCouponCodeCreateJobResponseObjectResourceAttributes) Get() *CouponCodeCreateJobResponseObjectResourceAttributes {
	return v.value
}

func (v *NullableCouponCodeCreateJobResponseObjectResourceAttributes) Set(val *CouponCodeCreateJobResponseObjectResourceAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponCodeCreateJobResponseObjectResourceAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponCodeCreateJobResponseObjectResourceAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponCodeCreateJobResponseObjectResourceAttributes(val *CouponCodeCreateJobResponseObjectResourceAttributes) *NullableCouponCodeCreateJobResponseObjectResourceAttributes {
	return &NullableCouponCodeCreateJobResponseObjectResourceAttributes{value: val, isSet: true}
}

func (v NullableCouponCodeCreateJobResponseObjectResourceAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponCodeCreateJobResponseObjectResourceAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


