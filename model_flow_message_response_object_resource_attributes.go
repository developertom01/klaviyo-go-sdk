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

// checks if the FlowMessageResponseObjectResourceAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlowMessageResponseObjectResourceAttributes{}

// FlowMessageResponseObjectResourceAttributes struct for FlowMessageResponseObjectResourceAttributes
type FlowMessageResponseObjectResourceAttributes struct {
	Name string `json:"name"`
	Channel string `json:"channel"`
	Content FlowMessageResponseObjectResourceAttributesContent `json:"content"`
	Created NullableTime `json:"created,omitempty"`
	Updated NullableTime `json:"updated,omitempty"`
}

type _FlowMessageResponseObjectResourceAttributes FlowMessageResponseObjectResourceAttributes

// NewFlowMessageResponseObjectResourceAttributes instantiates a new FlowMessageResponseObjectResourceAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowMessageResponseObjectResourceAttributes(name string, channel string, content FlowMessageResponseObjectResourceAttributesContent) *FlowMessageResponseObjectResourceAttributes {
	this := FlowMessageResponseObjectResourceAttributes{}
	this.Name = name
	this.Channel = channel
	this.Content = content
	return &this
}

// NewFlowMessageResponseObjectResourceAttributesWithDefaults instantiates a new FlowMessageResponseObjectResourceAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowMessageResponseObjectResourceAttributesWithDefaults() *FlowMessageResponseObjectResourceAttributes {
	this := FlowMessageResponseObjectResourceAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *FlowMessageResponseObjectResourceAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FlowMessageResponseObjectResourceAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FlowMessageResponseObjectResourceAttributes) SetName(v string) {
	o.Name = v
}

// GetChannel returns the Channel field value
func (o *FlowMessageResponseObjectResourceAttributes) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *FlowMessageResponseObjectResourceAttributes) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *FlowMessageResponseObjectResourceAttributes) SetChannel(v string) {
	o.Channel = v
}

// GetContent returns the Content field value
func (o *FlowMessageResponseObjectResourceAttributes) GetContent() FlowMessageResponseObjectResourceAttributesContent {
	if o == nil {
		var ret FlowMessageResponseObjectResourceAttributesContent
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *FlowMessageResponseObjectResourceAttributes) GetContentOk() (*FlowMessageResponseObjectResourceAttributesContent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *FlowMessageResponseObjectResourceAttributes) SetContent(v FlowMessageResponseObjectResourceAttributesContent) {
	o.Content = v
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlowMessageResponseObjectResourceAttributes) GetCreated() time.Time {
	if o == nil || IsNil(o.Created.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlowMessageResponseObjectResourceAttributes) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *FlowMessageResponseObjectResourceAttributes) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *FlowMessageResponseObjectResourceAttributes) SetCreated(v time.Time) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *FlowMessageResponseObjectResourceAttributes) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *FlowMessageResponseObjectResourceAttributes) UnsetCreated() {
	o.Created.Unset()
}

// GetUpdated returns the Updated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlowMessageResponseObjectResourceAttributes) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Updated.Get()
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlowMessageResponseObjectResourceAttributes) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Updated.Get(), o.Updated.IsSet()
}

// HasUpdated returns a boolean if a field has been set.
func (o *FlowMessageResponseObjectResourceAttributes) HasUpdated() bool {
	if o != nil && o.Updated.IsSet() {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given NullableTime and assigns it to the Updated field.
func (o *FlowMessageResponseObjectResourceAttributes) SetUpdated(v time.Time) {
	o.Updated.Set(&v)
}
// SetUpdatedNil sets the value for Updated to be an explicit nil
func (o *FlowMessageResponseObjectResourceAttributes) SetUpdatedNil() {
	o.Updated.Set(nil)
}

// UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
func (o *FlowMessageResponseObjectResourceAttributes) UnsetUpdated() {
	o.Updated.Unset()
}

func (o FlowMessageResponseObjectResourceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlowMessageResponseObjectResourceAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["channel"] = o.Channel
	toSerialize["content"] = o.Content
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if o.Updated.IsSet() {
		toSerialize["updated"] = o.Updated.Get()
	}
	return toSerialize, nil
}

func (o *FlowMessageResponseObjectResourceAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"channel",
		"content",
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

	varFlowMessageResponseObjectResourceAttributes := _FlowMessageResponseObjectResourceAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFlowMessageResponseObjectResourceAttributes)

	if err != nil {
		return err
	}

	*o = FlowMessageResponseObjectResourceAttributes(varFlowMessageResponseObjectResourceAttributes)

	return err
}

type NullableFlowMessageResponseObjectResourceAttributes struct {
	value *FlowMessageResponseObjectResourceAttributes
	isSet bool
}

func (v NullableFlowMessageResponseObjectResourceAttributes) Get() *FlowMessageResponseObjectResourceAttributes {
	return v.value
}

func (v *NullableFlowMessageResponseObjectResourceAttributes) Set(val *FlowMessageResponseObjectResourceAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowMessageResponseObjectResourceAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowMessageResponseObjectResourceAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowMessageResponseObjectResourceAttributes(val *FlowMessageResponseObjectResourceAttributes) *NullableFlowMessageResponseObjectResourceAttributes {
	return &NullableFlowMessageResponseObjectResourceAttributes{value: val, isSet: true}
}

func (v NullableFlowMessageResponseObjectResourceAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowMessageResponseObjectResourceAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

