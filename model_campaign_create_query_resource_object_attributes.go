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

// checks if the CampaignCreateQueryResourceObjectAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignCreateQueryResourceObjectAttributes{}

// CampaignCreateQueryResourceObjectAttributes struct for CampaignCreateQueryResourceObjectAttributes
type CampaignCreateQueryResourceObjectAttributes struct {
	// The campaign name
	Name string `json:"name"`
	Audiences AudiencesSubObject `json:"audiences"`
	SendStrategy *SendStrategySubObject `json:"send_strategy,omitempty"`
	SendOptions NullableCampaignCreateQueryResourceObjectAttributesSendOptions `json:"send_options,omitempty"`
	TrackingOptions NullableCampaignCreateQueryResourceObjectAttributesTrackingOptions `json:"tracking_options,omitempty"`
	CampaignMessages CampaignCreateQueryResourceObjectAttributesCampaignMessages `json:"campaign-messages"`
}

type _CampaignCreateQueryResourceObjectAttributes CampaignCreateQueryResourceObjectAttributes

// NewCampaignCreateQueryResourceObjectAttributes instantiates a new CampaignCreateQueryResourceObjectAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignCreateQueryResourceObjectAttributes(name string, audiences AudiencesSubObject, campaignMessages CampaignCreateQueryResourceObjectAttributesCampaignMessages) *CampaignCreateQueryResourceObjectAttributes {
	this := CampaignCreateQueryResourceObjectAttributes{}
	this.Name = name
	this.Audiences = audiences
	this.CampaignMessages = campaignMessages
	return &this
}

// NewCampaignCreateQueryResourceObjectAttributesWithDefaults instantiates a new CampaignCreateQueryResourceObjectAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignCreateQueryResourceObjectAttributesWithDefaults() *CampaignCreateQueryResourceObjectAttributes {
	this := CampaignCreateQueryResourceObjectAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *CampaignCreateQueryResourceObjectAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CampaignCreateQueryResourceObjectAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CampaignCreateQueryResourceObjectAttributes) SetName(v string) {
	o.Name = v
}

// GetAudiences returns the Audiences field value
func (o *CampaignCreateQueryResourceObjectAttributes) GetAudiences() AudiencesSubObject {
	if o == nil {
		var ret AudiencesSubObject
		return ret
	}

	return o.Audiences
}

// GetAudiencesOk returns a tuple with the Audiences field value
// and a boolean to check if the value has been set.
func (o *CampaignCreateQueryResourceObjectAttributes) GetAudiencesOk() (*AudiencesSubObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Audiences, true
}

// SetAudiences sets field value
func (o *CampaignCreateQueryResourceObjectAttributes) SetAudiences(v AudiencesSubObject) {
	o.Audiences = v
}

// GetSendStrategy returns the SendStrategy field value if set, zero value otherwise.
func (o *CampaignCreateQueryResourceObjectAttributes) GetSendStrategy() SendStrategySubObject {
	if o == nil || IsNil(o.SendStrategy) {
		var ret SendStrategySubObject
		return ret
	}
	return *o.SendStrategy
}

// GetSendStrategyOk returns a tuple with the SendStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCreateQueryResourceObjectAttributes) GetSendStrategyOk() (*SendStrategySubObject, bool) {
	if o == nil || IsNil(o.SendStrategy) {
		return nil, false
	}
	return o.SendStrategy, true
}

// HasSendStrategy returns a boolean if a field has been set.
func (o *CampaignCreateQueryResourceObjectAttributes) HasSendStrategy() bool {
	if o != nil && !IsNil(o.SendStrategy) {
		return true
	}

	return false
}

// SetSendStrategy gets a reference to the given SendStrategySubObject and assigns it to the SendStrategy field.
func (o *CampaignCreateQueryResourceObjectAttributes) SetSendStrategy(v SendStrategySubObject) {
	o.SendStrategy = &v
}

// GetSendOptions returns the SendOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignCreateQueryResourceObjectAttributes) GetSendOptions() CampaignCreateQueryResourceObjectAttributesSendOptions {
	if o == nil || IsNil(o.SendOptions.Get()) {
		var ret CampaignCreateQueryResourceObjectAttributesSendOptions
		return ret
	}
	return *o.SendOptions.Get()
}

// GetSendOptionsOk returns a tuple with the SendOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignCreateQueryResourceObjectAttributes) GetSendOptionsOk() (*CampaignCreateQueryResourceObjectAttributesSendOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.SendOptions.Get(), o.SendOptions.IsSet()
}

// HasSendOptions returns a boolean if a field has been set.
func (o *CampaignCreateQueryResourceObjectAttributes) HasSendOptions() bool {
	if o != nil && o.SendOptions.IsSet() {
		return true
	}

	return false
}

// SetSendOptions gets a reference to the given NullableCampaignCreateQueryResourceObjectAttributesSendOptions and assigns it to the SendOptions field.
func (o *CampaignCreateQueryResourceObjectAttributes) SetSendOptions(v CampaignCreateQueryResourceObjectAttributesSendOptions) {
	o.SendOptions.Set(&v)
}
// SetSendOptionsNil sets the value for SendOptions to be an explicit nil
func (o *CampaignCreateQueryResourceObjectAttributes) SetSendOptionsNil() {
	o.SendOptions.Set(nil)
}

// UnsetSendOptions ensures that no value is present for SendOptions, not even an explicit nil
func (o *CampaignCreateQueryResourceObjectAttributes) UnsetSendOptions() {
	o.SendOptions.Unset()
}

// GetTrackingOptions returns the TrackingOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignCreateQueryResourceObjectAttributes) GetTrackingOptions() CampaignCreateQueryResourceObjectAttributesTrackingOptions {
	if o == nil || IsNil(o.TrackingOptions.Get()) {
		var ret CampaignCreateQueryResourceObjectAttributesTrackingOptions
		return ret
	}
	return *o.TrackingOptions.Get()
}

// GetTrackingOptionsOk returns a tuple with the TrackingOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignCreateQueryResourceObjectAttributes) GetTrackingOptionsOk() (*CampaignCreateQueryResourceObjectAttributesTrackingOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrackingOptions.Get(), o.TrackingOptions.IsSet()
}

// HasTrackingOptions returns a boolean if a field has been set.
func (o *CampaignCreateQueryResourceObjectAttributes) HasTrackingOptions() bool {
	if o != nil && o.TrackingOptions.IsSet() {
		return true
	}

	return false
}

// SetTrackingOptions gets a reference to the given NullableCampaignCreateQueryResourceObjectAttributesTrackingOptions and assigns it to the TrackingOptions field.
func (o *CampaignCreateQueryResourceObjectAttributes) SetTrackingOptions(v CampaignCreateQueryResourceObjectAttributesTrackingOptions) {
	o.TrackingOptions.Set(&v)
}
// SetTrackingOptionsNil sets the value for TrackingOptions to be an explicit nil
func (o *CampaignCreateQueryResourceObjectAttributes) SetTrackingOptionsNil() {
	o.TrackingOptions.Set(nil)
}

// UnsetTrackingOptions ensures that no value is present for TrackingOptions, not even an explicit nil
func (o *CampaignCreateQueryResourceObjectAttributes) UnsetTrackingOptions() {
	o.TrackingOptions.Unset()
}

// GetCampaignMessages returns the CampaignMessages field value
func (o *CampaignCreateQueryResourceObjectAttributes) GetCampaignMessages() CampaignCreateQueryResourceObjectAttributesCampaignMessages {
	if o == nil {
		var ret CampaignCreateQueryResourceObjectAttributesCampaignMessages
		return ret
	}

	return o.CampaignMessages
}

// GetCampaignMessagesOk returns a tuple with the CampaignMessages field value
// and a boolean to check if the value has been set.
func (o *CampaignCreateQueryResourceObjectAttributes) GetCampaignMessagesOk() (*CampaignCreateQueryResourceObjectAttributesCampaignMessages, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignMessages, true
}

// SetCampaignMessages sets field value
func (o *CampaignCreateQueryResourceObjectAttributes) SetCampaignMessages(v CampaignCreateQueryResourceObjectAttributesCampaignMessages) {
	o.CampaignMessages = v
}

func (o CampaignCreateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignCreateQueryResourceObjectAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["audiences"] = o.Audiences
	if !IsNil(o.SendStrategy) {
		toSerialize["send_strategy"] = o.SendStrategy
	}
	if o.SendOptions.IsSet() {
		toSerialize["send_options"] = o.SendOptions.Get()
	}
	if o.TrackingOptions.IsSet() {
		toSerialize["tracking_options"] = o.TrackingOptions.Get()
	}
	toSerialize["campaign-messages"] = o.CampaignMessages
	return toSerialize, nil
}

func (o *CampaignCreateQueryResourceObjectAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"audiences",
		"campaign-messages",
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

	varCampaignCreateQueryResourceObjectAttributes := _CampaignCreateQueryResourceObjectAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCampaignCreateQueryResourceObjectAttributes)

	if err != nil {
		return err
	}

	*o = CampaignCreateQueryResourceObjectAttributes(varCampaignCreateQueryResourceObjectAttributes)

	return err
}

type NullableCampaignCreateQueryResourceObjectAttributes struct {
	value *CampaignCreateQueryResourceObjectAttributes
	isSet bool
}

func (v NullableCampaignCreateQueryResourceObjectAttributes) Get() *CampaignCreateQueryResourceObjectAttributes {
	return v.value
}

func (v *NullableCampaignCreateQueryResourceObjectAttributes) Set(val *CampaignCreateQueryResourceObjectAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignCreateQueryResourceObjectAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignCreateQueryResourceObjectAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignCreateQueryResourceObjectAttributes(val *CampaignCreateQueryResourceObjectAttributes) *NullableCampaignCreateQueryResourceObjectAttributes {
	return &NullableCampaignCreateQueryResourceObjectAttributes{value: val, isSet: true}
}

func (v NullableCampaignCreateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignCreateQueryResourceObjectAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

