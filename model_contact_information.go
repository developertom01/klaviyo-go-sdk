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

// checks if the ContactInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactInformation{}

// ContactInformation struct for ContactInformation
type ContactInformation struct {
	// This field is used to auto-populate the default sender name on flow and campaign emails.
	DefaultSenderName string `json:"default_sender_name"`
	// This field is used to auto-populate the default sender email address on flow and campaign emails.
	DefaultSenderEmail string `json:"default_sender_email"`
	WebsiteUrl string `json:"website_url"`
	OrganizationName string `json:"organization_name"`
	StreetAddress StreetAddress `json:"street_address"`
}

type _ContactInformation ContactInformation

// NewContactInformation instantiates a new ContactInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactInformation(defaultSenderName string, defaultSenderEmail string, websiteUrl string, organizationName string, streetAddress StreetAddress) *ContactInformation {
	this := ContactInformation{}
	this.DefaultSenderName = defaultSenderName
	this.DefaultSenderEmail = defaultSenderEmail
	this.WebsiteUrl = websiteUrl
	this.OrganizationName = organizationName
	this.StreetAddress = streetAddress
	return &this
}

// NewContactInformationWithDefaults instantiates a new ContactInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactInformationWithDefaults() *ContactInformation {
	this := ContactInformation{}
	return &this
}

// GetDefaultSenderName returns the DefaultSenderName field value
func (o *ContactInformation) GetDefaultSenderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultSenderName
}

// GetDefaultSenderNameOk returns a tuple with the DefaultSenderName field value
// and a boolean to check if the value has been set.
func (o *ContactInformation) GetDefaultSenderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultSenderName, true
}

// SetDefaultSenderName sets field value
func (o *ContactInformation) SetDefaultSenderName(v string) {
	o.DefaultSenderName = v
}

// GetDefaultSenderEmail returns the DefaultSenderEmail field value
func (o *ContactInformation) GetDefaultSenderEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultSenderEmail
}

// GetDefaultSenderEmailOk returns a tuple with the DefaultSenderEmail field value
// and a boolean to check if the value has been set.
func (o *ContactInformation) GetDefaultSenderEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultSenderEmail, true
}

// SetDefaultSenderEmail sets field value
func (o *ContactInformation) SetDefaultSenderEmail(v string) {
	o.DefaultSenderEmail = v
}

// GetWebsiteUrl returns the WebsiteUrl field value
func (o *ContactInformation) GetWebsiteUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebsiteUrl
}

// GetWebsiteUrlOk returns a tuple with the WebsiteUrl field value
// and a boolean to check if the value has been set.
func (o *ContactInformation) GetWebsiteUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebsiteUrl, true
}

// SetWebsiteUrl sets field value
func (o *ContactInformation) SetWebsiteUrl(v string) {
	o.WebsiteUrl = v
}

// GetOrganizationName returns the OrganizationName field value
func (o *ContactInformation) GetOrganizationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value
// and a boolean to check if the value has been set.
func (o *ContactInformation) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationName, true
}

// SetOrganizationName sets field value
func (o *ContactInformation) SetOrganizationName(v string) {
	o.OrganizationName = v
}

// GetStreetAddress returns the StreetAddress field value
func (o *ContactInformation) GetStreetAddress() StreetAddress {
	if o == nil {
		var ret StreetAddress
		return ret
	}

	return o.StreetAddress
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value
// and a boolean to check if the value has been set.
func (o *ContactInformation) GetStreetAddressOk() (*StreetAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StreetAddress, true
}

// SetStreetAddress sets field value
func (o *ContactInformation) SetStreetAddress(v StreetAddress) {
	o.StreetAddress = v
}

func (o ContactInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["default_sender_name"] = o.DefaultSenderName
	toSerialize["default_sender_email"] = o.DefaultSenderEmail
	toSerialize["website_url"] = o.WebsiteUrl
	toSerialize["organization_name"] = o.OrganizationName
	toSerialize["street_address"] = o.StreetAddress
	return toSerialize, nil
}

func (o *ContactInformation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"default_sender_name",
		"default_sender_email",
		"website_url",
		"organization_name",
		"street_address",
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

	varContactInformation := _ContactInformation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContactInformation)

	if err != nil {
		return err
	}

	*o = ContactInformation(varContactInformation)

	return err
}

type NullableContactInformation struct {
	value *ContactInformation
	isSet bool
}

func (v NullableContactInformation) Get() *ContactInformation {
	return v.value
}

func (v *NullableContactInformation) Set(val *ContactInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableContactInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableContactInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactInformation(val *ContactInformation) *NullableContactInformation {
	return &NullableContactInformation{value: val, isSet: true}
}

func (v NullableContactInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

