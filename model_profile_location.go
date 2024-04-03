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

// checks if the ProfileLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileLocation{}

// ProfileLocation struct for ProfileLocation
type ProfileLocation struct {
	// First line of street address
	Address1 NullableString `json:"address1,omitempty"`
	// Second line of street address
	Address2 NullableString `json:"address2,omitempty"`
	// City name
	City NullableString `json:"city,omitempty"`
	// Country name
	Country NullableString `json:"country,omitempty"`
	Latitude NullableProfileLocationLatitude `json:"latitude,omitempty"`
	Longitude NullableProfileLocationLongitude `json:"longitude,omitempty"`
	// Region within a country, such as state or province
	Region NullableString `json:"region,omitempty"`
	// Zip code
	Zip NullableString `json:"zip,omitempty"`
	// Time zone name. We recommend using time zones from the IANA Time Zone Database.
	Timezone NullableString `json:"timezone,omitempty"`
	// IP Address
	Ip NullableString `json:"ip,omitempty"`
}

// NewProfileLocation instantiates a new ProfileLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileLocation() *ProfileLocation {
	this := ProfileLocation{}
	return &this
}

// NewProfileLocationWithDefaults instantiates a new ProfileLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileLocationWithDefaults() *ProfileLocation {
	this := ProfileLocation{}
	return &this
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfileLocation) GetAddress1() string {
	if o == nil || IsNil(o.Address1.Get()) {
		var ret string
		return ret
	}
	return *o.Address1.Get()
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfileLocation) GetAddress1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address1.Get(), o.Address1.IsSet()
}

// HasAddress1 returns a boolean if a field has been set.
func (o *ProfileLocation) HasAddress1() bool {
	if o != nil && o.Address1.IsSet() {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given NullableString and assigns it to the Address1 field.
func (o *ProfileLocation) SetAddress1(v string) {
	o.Address1.Set(&v)
}
// SetAddress1Nil sets the value for Address1 to be an explicit nil
func (o *ProfileLocation) SetAddress1Nil() {
	o.Address1.Set(nil)
}

// UnsetAddress1 ensures that no value is present for Address1, not even an explicit nil
func (o *ProfileLocation) UnsetAddress1() {
	o.Address1.Unset()
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfileLocation) GetAddress2() string {
	if o == nil || IsNil(o.Address2.Get()) {
		var ret string
		return ret
	}
	return *o.Address2.Get()
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfileLocation) GetAddress2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address2.Get(), o.Address2.IsSet()
}

// HasAddress2 returns a boolean if a field has been set.
func (o *ProfileLocation) HasAddress2() bool {
	if o != nil && o.Address2.IsSet() {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given NullableString and assigns it to the Address2 field.
func (o *ProfileLocation) SetAddress2(v string) {
	o.Address2.Set(&v)
}
// SetAddress2Nil sets the value for Address2 to be an explicit nil
func (o *ProfileLocation) SetAddress2Nil() {
	o.Address2.Set(nil)
}

// UnsetAddress2 ensures that no value is present for Address2, not even an explicit nil
func (o *ProfileLocation) UnsetAddress2() {
	o.Address2.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfileLocation) GetCity() string {
	if o == nil || IsNil(o.City.Get()) {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfileLocation) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *ProfileLocation) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *ProfileLocation) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *ProfileLocation) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *ProfileLocation) UnsetCity() {
	o.City.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfileLocation) GetCountry() string {
	if o == nil || IsNil(o.Country.Get()) {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfileLocation) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *ProfileLocation) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *ProfileLocation) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *ProfileLocation) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *ProfileLocation) UnsetCountry() {
	o.Country.Unset()
}

// GetLatitude returns the Latitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfileLocation) GetLatitude() ProfileLocationLatitude {
	if o == nil || IsNil(o.Latitude.Get()) {
		var ret ProfileLocationLatitude
		return ret
	}
	return *o.Latitude.Get()
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfileLocation) GetLatitudeOk() (*ProfileLocationLatitude, bool) {
	if o == nil {
		return nil, false
	}
	return o.Latitude.Get(), o.Latitude.IsSet()
}

// HasLatitude returns a boolean if a field has been set.
func (o *ProfileLocation) HasLatitude() bool {
	if o != nil && o.Latitude.IsSet() {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given NullableProfileLocationLatitude and assigns it to the Latitude field.
func (o *ProfileLocation) SetLatitude(v ProfileLocationLatitude) {
	o.Latitude.Set(&v)
}
// SetLatitudeNil sets the value for Latitude to be an explicit nil
func (o *ProfileLocation) SetLatitudeNil() {
	o.Latitude.Set(nil)
}

// UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
func (o *ProfileLocation) UnsetLatitude() {
	o.Latitude.Unset()
}

// GetLongitude returns the Longitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfileLocation) GetLongitude() ProfileLocationLongitude {
	if o == nil || IsNil(o.Longitude.Get()) {
		var ret ProfileLocationLongitude
		return ret
	}
	return *o.Longitude.Get()
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfileLocation) GetLongitudeOk() (*ProfileLocationLongitude, bool) {
	if o == nil {
		return nil, false
	}
	return o.Longitude.Get(), o.Longitude.IsSet()
}

// HasLongitude returns a boolean if a field has been set.
func (o *ProfileLocation) HasLongitude() bool {
	if o != nil && o.Longitude.IsSet() {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given NullableProfileLocationLongitude and assigns it to the Longitude field.
func (o *ProfileLocation) SetLongitude(v ProfileLocationLongitude) {
	o.Longitude.Set(&v)
}
// SetLongitudeNil sets the value for Longitude to be an explicit nil
func (o *ProfileLocation) SetLongitudeNil() {
	o.Longitude.Set(nil)
}

// UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
func (o *ProfileLocation) UnsetLongitude() {
	o.Longitude.Unset()
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfileLocation) GetRegion() string {
	if o == nil || IsNil(o.Region.Get()) {
		var ret string
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfileLocation) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *ProfileLocation) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableString and assigns it to the Region field.
func (o *ProfileLocation) SetRegion(v string) {
	o.Region.Set(&v)
}
// SetRegionNil sets the value for Region to be an explicit nil
func (o *ProfileLocation) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *ProfileLocation) UnsetRegion() {
	o.Region.Unset()
}

// GetZip returns the Zip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfileLocation) GetZip() string {
	if o == nil || IsNil(o.Zip.Get()) {
		var ret string
		return ret
	}
	return *o.Zip.Get()
}

// GetZipOk returns a tuple with the Zip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfileLocation) GetZipOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Zip.Get(), o.Zip.IsSet()
}

// HasZip returns a boolean if a field has been set.
func (o *ProfileLocation) HasZip() bool {
	if o != nil && o.Zip.IsSet() {
		return true
	}

	return false
}

// SetZip gets a reference to the given NullableString and assigns it to the Zip field.
func (o *ProfileLocation) SetZip(v string) {
	o.Zip.Set(&v)
}
// SetZipNil sets the value for Zip to be an explicit nil
func (o *ProfileLocation) SetZipNil() {
	o.Zip.Set(nil)
}

// UnsetZip ensures that no value is present for Zip, not even an explicit nil
func (o *ProfileLocation) UnsetZip() {
	o.Zip.Unset()
}

// GetTimezone returns the Timezone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfileLocation) GetTimezone() string {
	if o == nil || IsNil(o.Timezone.Get()) {
		var ret string
		return ret
	}
	return *o.Timezone.Get()
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfileLocation) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timezone.Get(), o.Timezone.IsSet()
}

// HasTimezone returns a boolean if a field has been set.
func (o *ProfileLocation) HasTimezone() bool {
	if o != nil && o.Timezone.IsSet() {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given NullableString and assigns it to the Timezone field.
func (o *ProfileLocation) SetTimezone(v string) {
	o.Timezone.Set(&v)
}
// SetTimezoneNil sets the value for Timezone to be an explicit nil
func (o *ProfileLocation) SetTimezoneNil() {
	o.Timezone.Set(nil)
}

// UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
func (o *ProfileLocation) UnsetTimezone() {
	o.Timezone.Unset()
}

// GetIp returns the Ip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfileLocation) GetIp() string {
	if o == nil || IsNil(o.Ip.Get()) {
		var ret string
		return ret
	}
	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfileLocation) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// HasIp returns a boolean if a field has been set.
func (o *ProfileLocation) HasIp() bool {
	if o != nil && o.Ip.IsSet() {
		return true
	}

	return false
}

// SetIp gets a reference to the given NullableString and assigns it to the Ip field.
func (o *ProfileLocation) SetIp(v string) {
	o.Ip.Set(&v)
}
// SetIpNil sets the value for Ip to be an explicit nil
func (o *ProfileLocation) SetIpNil() {
	o.Ip.Set(nil)
}

// UnsetIp ensures that no value is present for Ip, not even an explicit nil
func (o *ProfileLocation) UnsetIp() {
	o.Ip.Unset()
}

func (o ProfileLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Address1.IsSet() {
		toSerialize["address1"] = o.Address1.Get()
	}
	if o.Address2.IsSet() {
		toSerialize["address2"] = o.Address2.Get()
	}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
	}
	if o.Latitude.IsSet() {
		toSerialize["latitude"] = o.Latitude.Get()
	}
	if o.Longitude.IsSet() {
		toSerialize["longitude"] = o.Longitude.Get()
	}
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	if o.Zip.IsSet() {
		toSerialize["zip"] = o.Zip.Get()
	}
	if o.Timezone.IsSet() {
		toSerialize["timezone"] = o.Timezone.Get()
	}
	if o.Ip.IsSet() {
		toSerialize["ip"] = o.Ip.Get()
	}
	return toSerialize, nil
}

type NullableProfileLocation struct {
	value *ProfileLocation
	isSet bool
}

func (v NullableProfileLocation) Get() *ProfileLocation {
	return v.value
}

func (v *NullableProfileLocation) Set(val *ProfileLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileLocation(val *ProfileLocation) *NullableProfileLocation {
	return &NullableProfileLocation{value: val, isSet: true}
}

func (v NullableProfileLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


