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

// checks if the ProfilePartialUpdateQueryResourceObjectAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfilePartialUpdateQueryResourceObjectAttributes{}

// ProfilePartialUpdateQueryResourceObjectAttributes struct for ProfilePartialUpdateQueryResourceObjectAttributes
type ProfilePartialUpdateQueryResourceObjectAttributes struct {
	// Individual's email address
	Email NullableString `json:"email,omitempty"`
	// Individual's phone number in E.164 format
	PhoneNumber NullableString `json:"phone_number,omitempty"`
	// A unique identifier used by customers to associate Klaviyo profiles with profiles in an external system, such as a point-of-sale system. Format varies based on the external system.
	ExternalId NullableString `json:"external_id,omitempty"`
	AnonymousId NullableString `json:"anonymous_id,omitempty"`
	// Individual's first name
	FirstName NullableString `json:"first_name,omitempty"`
	// Individual's last name
	LastName NullableString `json:"last_name,omitempty"`
	// Name of the company or organization within the company for whom the individual works
	Organization NullableString `json:"organization,omitempty"`
	// Individual's job title
	Title NullableString `json:"title,omitempty"`
	// URL pointing to the location of a profile image
	Image NullableString `json:"image,omitempty"`
	Location *ProfileLocation `json:"location,omitempty"`
	// An object containing key/value pairs for any custom properties assigned to this profile
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// NewProfilePartialUpdateQueryResourceObjectAttributes instantiates a new ProfilePartialUpdateQueryResourceObjectAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfilePartialUpdateQueryResourceObjectAttributes() *ProfilePartialUpdateQueryResourceObjectAttributes {
	this := ProfilePartialUpdateQueryResourceObjectAttributes{}
	return &this
}

// NewProfilePartialUpdateQueryResourceObjectAttributesWithDefaults instantiates a new ProfilePartialUpdateQueryResourceObjectAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfilePartialUpdateQueryResourceObjectAttributesWithDefaults() *ProfilePartialUpdateQueryResourceObjectAttributes {
	this := ProfilePartialUpdateQueryResourceObjectAttributes{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) UnsetEmail() {
	o.Email.Unset()
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber.Get()) {
		var ret string
		return ret
	}
	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given NullableString and assigns it to the PhoneNumber field.
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) SetPhoneNumber(v string) {
	o.PhoneNumber.Set(&v)
}
// SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) SetPhoneNumberNil() {
	o.PhoneNumber.Set(nil)
}

// UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) UnsetPhoneNumber() {
	o.PhoneNumber.Unset()
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// HasExternalId returns a boolean if a field has been set.
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) HasExternalId() bool {
	if o != nil && o.ExternalId.IsSet() {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given NullableString and assigns it to the ExternalId field.
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}
// SetExternalIdNil sets the value for ExternalId to be an explicit nil
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) SetExternalIdNil() {
	o.ExternalId.Set(nil)
}

// UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) UnsetExternalId() {
	o.ExternalId.Unset()
}

// GetAnonymousId returns the AnonymousId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) GetAnonymousId() string {
	if o == nil || IsNil(o.AnonymousId.Get()) {
		var ret string
		return ret
	}
	return *o.AnonymousId.Get()
}

// GetAnonymousIdOk returns a tuple with the AnonymousId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) GetAnonymousIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AnonymousId.Get(), o.AnonymousId.IsSet()
}

// HasAnonymousId returns a boolean if a field has been set.
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) HasAnonymousId() bool {
	if o != nil && o.AnonymousId.IsSet() {
		return true
	}

	return false
}

// SetAnonymousId gets a reference to the given NullableString and assigns it to the AnonymousId field.
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) SetAnonymousId(v string) {
	o.AnonymousId.Set(&v)
}
// SetAnonymousIdNil sets the value for AnonymousId to be an explicit nil
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) SetAnonymousIdNil() {
	o.AnonymousId.Set(nil)
}

// UnsetAnonymousId ensures that no value is present for AnonymousId, not even an explicit nil
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) UnsetAnonymousId() {
	o.AnonymousId.Unset()
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) GetFirstName() string {
	if o == nil || IsNil(o.FirstName.Get()) {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) GetLastName() string {
	if o == nil || IsNil(o.LastName.Get()) {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) UnsetLastName() {
	o.LastName.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) GetOrganization() string {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret string
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) GetOrganizationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableString and assigns it to the Organization field.
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) SetOrganization(v string) {
	o.Organization.Set(&v)
}
// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) UnsetOrganization() {
	o.Organization.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) UnsetTitle() {
	o.Title.Unset()
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) GetImage() string {
	if o == nil || IsNil(o.Image.Get()) {
		var ret string
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableString and assigns it to the Image field.
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) SetImage(v string) {
	o.Image.Set(&v)
}
// SetImageNil sets the value for Image to be an explicit nil
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) UnsetImage() {
	o.Image.Unset()
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) GetLocation() ProfileLocation {
	if o == nil || IsNil(o.Location) {
		var ret ProfileLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) GetLocationOk() (*ProfileLocation, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given ProfileLocation and assigns it to the Location field.
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) SetLocation(v ProfileLocation) {
	o.Location = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) GetProperties() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Properties) {
		return map[string]interface{}{}, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *ProfilePartialUpdateQueryResourceObjectAttributes) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

func (o ProfilePartialUpdateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfilePartialUpdateQueryResourceObjectAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.PhoneNumber.IsSet() {
		toSerialize["phone_number"] = o.PhoneNumber.Get()
	}
	if o.ExternalId.IsSet() {
		toSerialize["external_id"] = o.ExternalId.Get()
	}
	if o.AnonymousId.IsSet() {
		toSerialize["anonymous_id"] = o.AnonymousId.Get()
	}
	if o.FirstName.IsSet() {
		toSerialize["first_name"] = o.FirstName.Get()
	}
	if o.LastName.IsSet() {
		toSerialize["last_name"] = o.LastName.Get()
	}
	if o.Organization.IsSet() {
		toSerialize["organization"] = o.Organization.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Image.IsSet() {
		toSerialize["image"] = o.Image.Get()
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableProfilePartialUpdateQueryResourceObjectAttributes struct {
	value *ProfilePartialUpdateQueryResourceObjectAttributes
	isSet bool
}

func (v NullableProfilePartialUpdateQueryResourceObjectAttributes) Get() *ProfilePartialUpdateQueryResourceObjectAttributes {
	return v.value
}

func (v *NullableProfilePartialUpdateQueryResourceObjectAttributes) Set(val *ProfilePartialUpdateQueryResourceObjectAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableProfilePartialUpdateQueryResourceObjectAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableProfilePartialUpdateQueryResourceObjectAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfilePartialUpdateQueryResourceObjectAttributes(val *ProfilePartialUpdateQueryResourceObjectAttributes) *NullableProfilePartialUpdateQueryResourceObjectAttributes {
	return &NullableProfilePartialUpdateQueryResourceObjectAttributes{value: val, isSet: true}
}

func (v NullableProfilePartialUpdateQueryResourceObjectAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfilePartialUpdateQueryResourceObjectAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


