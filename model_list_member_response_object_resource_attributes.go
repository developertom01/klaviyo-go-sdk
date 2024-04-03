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

// checks if the ListMemberResponseObjectResourceAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListMemberResponseObjectResourceAttributes{}

// ListMemberResponseObjectResourceAttributes struct for ListMemberResponseObjectResourceAttributes
type ListMemberResponseObjectResourceAttributes struct {
	// Individual's email address
	Email NullableString `json:"email,omitempty"`
	// Individual's phone number in E.164 format
	PhoneNumber NullableString `json:"phone_number,omitempty"`
	// A unique identifier used by customers to associate Klaviyo profiles with profiles in an external system, such as a point-of-sale system. Format varies based on the external system.
	ExternalId NullableString `json:"external_id,omitempty"`
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
	// Date and time when the profile was created, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm)
	Created NullableTime `json:"created,omitempty"`
	// Date and time when the profile was last updated, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm)
	Updated NullableTime `json:"updated,omitempty"`
	// Date and time of the most recent event the triggered an update to the profile, in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm)
	LastEventDate NullableTime `json:"last_event_date,omitempty"`
	Location *ProfileLocation `json:"location,omitempty"`
	// An object containing key/value pairs for any custom properties assigned to this profile
	Properties map[string]interface{} `json:"properties,omitempty"`
	// The datetime when this profile most recently joined the list.
	JoinedGroupAt time.Time `json:"joined_group_at"`
}

type _ListMemberResponseObjectResourceAttributes ListMemberResponseObjectResourceAttributes

// NewListMemberResponseObjectResourceAttributes instantiates a new ListMemberResponseObjectResourceAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMemberResponseObjectResourceAttributes(joinedGroupAt time.Time) *ListMemberResponseObjectResourceAttributes {
	this := ListMemberResponseObjectResourceAttributes{}
	this.JoinedGroupAt = joinedGroupAt
	return &this
}

// NewListMemberResponseObjectResourceAttributesWithDefaults instantiates a new ListMemberResponseObjectResourceAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMemberResponseObjectResourceAttributesWithDefaults() *ListMemberResponseObjectResourceAttributes {
	this := ListMemberResponseObjectResourceAttributes{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListMemberResponseObjectResourceAttributes) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListMemberResponseObjectResourceAttributes) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *ListMemberResponseObjectResourceAttributes) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *ListMemberResponseObjectResourceAttributes) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *ListMemberResponseObjectResourceAttributes) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *ListMemberResponseObjectResourceAttributes) UnsetEmail() {
	o.Email.Unset()
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListMemberResponseObjectResourceAttributes) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber.Get()) {
		var ret string
		return ret
	}
	return *o.PhoneNumber.Get()
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListMemberResponseObjectResourceAttributes) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhoneNumber.Get(), o.PhoneNumber.IsSet()
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *ListMemberResponseObjectResourceAttributes) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber.IsSet() {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given NullableString and assigns it to the PhoneNumber field.
func (o *ListMemberResponseObjectResourceAttributes) SetPhoneNumber(v string) {
	o.PhoneNumber.Set(&v)
}
// SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil
func (o *ListMemberResponseObjectResourceAttributes) SetPhoneNumberNil() {
	o.PhoneNumber.Set(nil)
}

// UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
func (o *ListMemberResponseObjectResourceAttributes) UnsetPhoneNumber() {
	o.PhoneNumber.Unset()
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListMemberResponseObjectResourceAttributes) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListMemberResponseObjectResourceAttributes) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// HasExternalId returns a boolean if a field has been set.
func (o *ListMemberResponseObjectResourceAttributes) HasExternalId() bool {
	if o != nil && o.ExternalId.IsSet() {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given NullableString and assigns it to the ExternalId field.
func (o *ListMemberResponseObjectResourceAttributes) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}
// SetExternalIdNil sets the value for ExternalId to be an explicit nil
func (o *ListMemberResponseObjectResourceAttributes) SetExternalIdNil() {
	o.ExternalId.Set(nil)
}

// UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
func (o *ListMemberResponseObjectResourceAttributes) UnsetExternalId() {
	o.ExternalId.Unset()
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListMemberResponseObjectResourceAttributes) GetFirstName() string {
	if o == nil || IsNil(o.FirstName.Get()) {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListMemberResponseObjectResourceAttributes) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *ListMemberResponseObjectResourceAttributes) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *ListMemberResponseObjectResourceAttributes) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *ListMemberResponseObjectResourceAttributes) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *ListMemberResponseObjectResourceAttributes) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListMemberResponseObjectResourceAttributes) GetLastName() string {
	if o == nil || IsNil(o.LastName.Get()) {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListMemberResponseObjectResourceAttributes) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *ListMemberResponseObjectResourceAttributes) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *ListMemberResponseObjectResourceAttributes) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *ListMemberResponseObjectResourceAttributes) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *ListMemberResponseObjectResourceAttributes) UnsetLastName() {
	o.LastName.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListMemberResponseObjectResourceAttributes) GetOrganization() string {
	if o == nil || IsNil(o.Organization.Get()) {
		var ret string
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListMemberResponseObjectResourceAttributes) GetOrganizationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *ListMemberResponseObjectResourceAttributes) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableString and assigns it to the Organization field.
func (o *ListMemberResponseObjectResourceAttributes) SetOrganization(v string) {
	o.Organization.Set(&v)
}
// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *ListMemberResponseObjectResourceAttributes) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *ListMemberResponseObjectResourceAttributes) UnsetOrganization() {
	o.Organization.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListMemberResponseObjectResourceAttributes) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListMemberResponseObjectResourceAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ListMemberResponseObjectResourceAttributes) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *ListMemberResponseObjectResourceAttributes) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *ListMemberResponseObjectResourceAttributes) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *ListMemberResponseObjectResourceAttributes) UnsetTitle() {
	o.Title.Unset()
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListMemberResponseObjectResourceAttributes) GetImage() string {
	if o == nil || IsNil(o.Image.Get()) {
		var ret string
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListMemberResponseObjectResourceAttributes) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *ListMemberResponseObjectResourceAttributes) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableString and assigns it to the Image field.
func (o *ListMemberResponseObjectResourceAttributes) SetImage(v string) {
	o.Image.Set(&v)
}
// SetImageNil sets the value for Image to be an explicit nil
func (o *ListMemberResponseObjectResourceAttributes) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *ListMemberResponseObjectResourceAttributes) UnsetImage() {
	o.Image.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListMemberResponseObjectResourceAttributes) GetCreated() time.Time {
	if o == nil || IsNil(o.Created.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListMemberResponseObjectResourceAttributes) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *ListMemberResponseObjectResourceAttributes) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *ListMemberResponseObjectResourceAttributes) SetCreated(v time.Time) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *ListMemberResponseObjectResourceAttributes) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *ListMemberResponseObjectResourceAttributes) UnsetCreated() {
	o.Created.Unset()
}

// GetUpdated returns the Updated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListMemberResponseObjectResourceAttributes) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Updated.Get()
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListMemberResponseObjectResourceAttributes) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Updated.Get(), o.Updated.IsSet()
}

// HasUpdated returns a boolean if a field has been set.
func (o *ListMemberResponseObjectResourceAttributes) HasUpdated() bool {
	if o != nil && o.Updated.IsSet() {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given NullableTime and assigns it to the Updated field.
func (o *ListMemberResponseObjectResourceAttributes) SetUpdated(v time.Time) {
	o.Updated.Set(&v)
}
// SetUpdatedNil sets the value for Updated to be an explicit nil
func (o *ListMemberResponseObjectResourceAttributes) SetUpdatedNil() {
	o.Updated.Set(nil)
}

// UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
func (o *ListMemberResponseObjectResourceAttributes) UnsetUpdated() {
	o.Updated.Unset()
}

// GetLastEventDate returns the LastEventDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListMemberResponseObjectResourceAttributes) GetLastEventDate() time.Time {
	if o == nil || IsNil(o.LastEventDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastEventDate.Get()
}

// GetLastEventDateOk returns a tuple with the LastEventDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListMemberResponseObjectResourceAttributes) GetLastEventDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastEventDate.Get(), o.LastEventDate.IsSet()
}

// HasLastEventDate returns a boolean if a field has been set.
func (o *ListMemberResponseObjectResourceAttributes) HasLastEventDate() bool {
	if o != nil && o.LastEventDate.IsSet() {
		return true
	}

	return false
}

// SetLastEventDate gets a reference to the given NullableTime and assigns it to the LastEventDate field.
func (o *ListMemberResponseObjectResourceAttributes) SetLastEventDate(v time.Time) {
	o.LastEventDate.Set(&v)
}
// SetLastEventDateNil sets the value for LastEventDate to be an explicit nil
func (o *ListMemberResponseObjectResourceAttributes) SetLastEventDateNil() {
	o.LastEventDate.Set(nil)
}

// UnsetLastEventDate ensures that no value is present for LastEventDate, not even an explicit nil
func (o *ListMemberResponseObjectResourceAttributes) UnsetLastEventDate() {
	o.LastEventDate.Unset()
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ListMemberResponseObjectResourceAttributes) GetLocation() ProfileLocation {
	if o == nil || IsNil(o.Location) {
		var ret ProfileLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMemberResponseObjectResourceAttributes) GetLocationOk() (*ProfileLocation, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ListMemberResponseObjectResourceAttributes) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given ProfileLocation and assigns it to the Location field.
func (o *ListMemberResponseObjectResourceAttributes) SetLocation(v ProfileLocation) {
	o.Location = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListMemberResponseObjectResourceAttributes) GetProperties() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListMemberResponseObjectResourceAttributes) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Properties) {
		return map[string]interface{}{}, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ListMemberResponseObjectResourceAttributes) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *ListMemberResponseObjectResourceAttributes) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

// GetJoinedGroupAt returns the JoinedGroupAt field value
func (o *ListMemberResponseObjectResourceAttributes) GetJoinedGroupAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.JoinedGroupAt
}

// GetJoinedGroupAtOk returns a tuple with the JoinedGroupAt field value
// and a boolean to check if the value has been set.
func (o *ListMemberResponseObjectResourceAttributes) GetJoinedGroupAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinedGroupAt, true
}

// SetJoinedGroupAt sets field value
func (o *ListMemberResponseObjectResourceAttributes) SetJoinedGroupAt(v time.Time) {
	o.JoinedGroupAt = v
}

func (o ListMemberResponseObjectResourceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListMemberResponseObjectResourceAttributes) ToMap() (map[string]interface{}, error) {
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
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if o.Updated.IsSet() {
		toSerialize["updated"] = o.Updated.Get()
	}
	if o.LastEventDate.IsSet() {
		toSerialize["last_event_date"] = o.LastEventDate.Get()
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	toSerialize["joined_group_at"] = o.JoinedGroupAt
	return toSerialize, nil
}

func (o *ListMemberResponseObjectResourceAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"joined_group_at",
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

	varListMemberResponseObjectResourceAttributes := _ListMemberResponseObjectResourceAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListMemberResponseObjectResourceAttributes)

	if err != nil {
		return err
	}

	*o = ListMemberResponseObjectResourceAttributes(varListMemberResponseObjectResourceAttributes)

	return err
}

type NullableListMemberResponseObjectResourceAttributes struct {
	value *ListMemberResponseObjectResourceAttributes
	isSet bool
}

func (v NullableListMemberResponseObjectResourceAttributes) Get() *ListMemberResponseObjectResourceAttributes {
	return v.value
}

func (v *NullableListMemberResponseObjectResourceAttributes) Set(val *ListMemberResponseObjectResourceAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableListMemberResponseObjectResourceAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableListMemberResponseObjectResourceAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMemberResponseObjectResourceAttributes(val *ListMemberResponseObjectResourceAttributes) *NullableListMemberResponseObjectResourceAttributes {
	return &NullableListMemberResponseObjectResourceAttributes{value: val, isSet: true}
}

func (v NullableListMemberResponseObjectResourceAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMemberResponseObjectResourceAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

