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

// checks if the CampaignMessageCreateQueryResourceObjectAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignMessageCreateQueryResourceObjectAttributes{}

// CampaignMessageCreateQueryResourceObjectAttributes struct for CampaignMessageCreateQueryResourceObjectAttributes
type CampaignMessageCreateQueryResourceObjectAttributes struct {
	// The channel the message is to be sent on (email or sms, for example)
	Channel string `json:"channel"`
	// The label or name on the message
	Label NullableString `json:"label,omitempty"`
	Content NullableCampaignMessageCreateQueryResourceObjectAttributesContent `json:"content,omitempty"`
	RenderOptions *RenderOptionsSubObject `json:"render_options,omitempty"`
}

type _CampaignMessageCreateQueryResourceObjectAttributes CampaignMessageCreateQueryResourceObjectAttributes

// NewCampaignMessageCreateQueryResourceObjectAttributes instantiates a new CampaignMessageCreateQueryResourceObjectAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignMessageCreateQueryResourceObjectAttributes(channel string) *CampaignMessageCreateQueryResourceObjectAttributes {
	this := CampaignMessageCreateQueryResourceObjectAttributes{}
	this.Channel = channel
	return &this
}

// NewCampaignMessageCreateQueryResourceObjectAttributesWithDefaults instantiates a new CampaignMessageCreateQueryResourceObjectAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignMessageCreateQueryResourceObjectAttributesWithDefaults() *CampaignMessageCreateQueryResourceObjectAttributes {
	this := CampaignMessageCreateQueryResourceObjectAttributes{}
	return &this
}

// GetChannel returns the Channel field value
func (o *CampaignMessageCreateQueryResourceObjectAttributes) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *CampaignMessageCreateQueryResourceObjectAttributes) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *CampaignMessageCreateQueryResourceObjectAttributes) SetChannel(v string) {
	o.Channel = v
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignMessageCreateQueryResourceObjectAttributes) GetLabel() string {
	if o == nil || IsNil(o.Label.Get()) {
		var ret string
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignMessageCreateQueryResourceObjectAttributes) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *CampaignMessageCreateQueryResourceObjectAttributes) HasLabel() bool {
	if o != nil && o.Label.IsSet() {
		return true
	}

	return false
}

// SetLabel gets a reference to the given NullableString and assigns it to the Label field.
func (o *CampaignMessageCreateQueryResourceObjectAttributes) SetLabel(v string) {
	o.Label.Set(&v)
}
// SetLabelNil sets the value for Label to be an explicit nil
func (o *CampaignMessageCreateQueryResourceObjectAttributes) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil
func (o *CampaignMessageCreateQueryResourceObjectAttributes) UnsetLabel() {
	o.Label.Unset()
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignMessageCreateQueryResourceObjectAttributes) GetContent() CampaignMessageCreateQueryResourceObjectAttributesContent {
	if o == nil || IsNil(o.Content.Get()) {
		var ret CampaignMessageCreateQueryResourceObjectAttributesContent
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignMessageCreateQueryResourceObjectAttributes) GetContentOk() (*CampaignMessageCreateQueryResourceObjectAttributesContent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *CampaignMessageCreateQueryResourceObjectAttributes) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableCampaignMessageCreateQueryResourceObjectAttributesContent and assigns it to the Content field.
func (o *CampaignMessageCreateQueryResourceObjectAttributes) SetContent(v CampaignMessageCreateQueryResourceObjectAttributesContent) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *CampaignMessageCreateQueryResourceObjectAttributes) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *CampaignMessageCreateQueryResourceObjectAttributes) UnsetContent() {
	o.Content.Unset()
}

// GetRenderOptions returns the RenderOptions field value if set, zero value otherwise.
func (o *CampaignMessageCreateQueryResourceObjectAttributes) GetRenderOptions() RenderOptionsSubObject {
	if o == nil || IsNil(o.RenderOptions) {
		var ret RenderOptionsSubObject
		return ret
	}
	return *o.RenderOptions
}

// GetRenderOptionsOk returns a tuple with the RenderOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignMessageCreateQueryResourceObjectAttributes) GetRenderOptionsOk() (*RenderOptionsSubObject, bool) {
	if o == nil || IsNil(o.RenderOptions) {
		return nil, false
	}
	return o.RenderOptions, true
}

// HasRenderOptions returns a boolean if a field has been set.
func (o *CampaignMessageCreateQueryResourceObjectAttributes) HasRenderOptions() bool {
	if o != nil && !IsNil(o.RenderOptions) {
		return true
	}

	return false
}

// SetRenderOptions gets a reference to the given RenderOptionsSubObject and assigns it to the RenderOptions field.
func (o *CampaignMessageCreateQueryResourceObjectAttributes) SetRenderOptions(v RenderOptionsSubObject) {
	o.RenderOptions = &v
}

func (o CampaignMessageCreateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignMessageCreateQueryResourceObjectAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel"] = o.Channel
	if o.Label.IsSet() {
		toSerialize["label"] = o.Label.Get()
	}
	if o.Content.IsSet() {
		toSerialize["content"] = o.Content.Get()
	}
	if !IsNil(o.RenderOptions) {
		toSerialize["render_options"] = o.RenderOptions
	}
	return toSerialize, nil
}

func (o *CampaignMessageCreateQueryResourceObjectAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"channel",
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

	varCampaignMessageCreateQueryResourceObjectAttributes := _CampaignMessageCreateQueryResourceObjectAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCampaignMessageCreateQueryResourceObjectAttributes)

	if err != nil {
		return err
	}

	*o = CampaignMessageCreateQueryResourceObjectAttributes(varCampaignMessageCreateQueryResourceObjectAttributes)

	return err
}

type NullableCampaignMessageCreateQueryResourceObjectAttributes struct {
	value *CampaignMessageCreateQueryResourceObjectAttributes
	isSet bool
}

func (v NullableCampaignMessageCreateQueryResourceObjectAttributes) Get() *CampaignMessageCreateQueryResourceObjectAttributes {
	return v.value
}

func (v *NullableCampaignMessageCreateQueryResourceObjectAttributes) Set(val *CampaignMessageCreateQueryResourceObjectAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignMessageCreateQueryResourceObjectAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignMessageCreateQueryResourceObjectAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignMessageCreateQueryResourceObjectAttributes(val *CampaignMessageCreateQueryResourceObjectAttributes) *NullableCampaignMessageCreateQueryResourceObjectAttributes {
	return &NullableCampaignMessageCreateQueryResourceObjectAttributes{value: val, isSet: true}
}

func (v NullableCampaignMessageCreateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignMessageCreateQueryResourceObjectAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

