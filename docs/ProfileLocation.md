# ProfileLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address1** | Pointer to **NullableString** | First line of street address | [optional] 
**Address2** | Pointer to **NullableString** | Second line of street address | [optional] 
**City** | Pointer to **NullableString** | City name | [optional] 
**Country** | Pointer to **NullableString** | Country name | [optional] 
**Latitude** | Pointer to [**NullableProfileLocationLatitude**](ProfileLocationLatitude.md) |  | [optional] 
**Longitude** | Pointer to [**NullableProfileLocationLongitude**](ProfileLocationLongitude.md) |  | [optional] 
**Region** | Pointer to **NullableString** | Region within a country, such as state or province | [optional] 
**Zip** | Pointer to **NullableString** | Zip code | [optional] 
**Timezone** | Pointer to **NullableString** | Time zone name. We recommend using time zones from the IANA Time Zone Database. | [optional] 
**Ip** | Pointer to **NullableString** | IP Address | [optional] 

## Methods

### NewProfileLocation

`func NewProfileLocation() *ProfileLocation`

NewProfileLocation instantiates a new ProfileLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileLocationWithDefaults

`func NewProfileLocationWithDefaults() *ProfileLocation`

NewProfileLocationWithDefaults instantiates a new ProfileLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress1

`func (o *ProfileLocation) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *ProfileLocation) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *ProfileLocation) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *ProfileLocation) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### SetAddress1Nil

`func (o *ProfileLocation) SetAddress1Nil(b bool)`

 SetAddress1Nil sets the value for Address1 to be an explicit nil

### UnsetAddress1
`func (o *ProfileLocation) UnsetAddress1()`

UnsetAddress1 ensures that no value is present for Address1, not even an explicit nil
### GetAddress2

`func (o *ProfileLocation) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *ProfileLocation) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *ProfileLocation) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *ProfileLocation) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### SetAddress2Nil

`func (o *ProfileLocation) SetAddress2Nil(b bool)`

 SetAddress2Nil sets the value for Address2 to be an explicit nil

### UnsetAddress2
`func (o *ProfileLocation) UnsetAddress2()`

UnsetAddress2 ensures that no value is present for Address2, not even an explicit nil
### GetCity

`func (o *ProfileLocation) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ProfileLocation) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ProfileLocation) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ProfileLocation) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *ProfileLocation) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *ProfileLocation) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCountry

`func (o *ProfileLocation) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ProfileLocation) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ProfileLocation) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ProfileLocation) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *ProfileLocation) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *ProfileLocation) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetLatitude

`func (o *ProfileLocation) GetLatitude() ProfileLocationLatitude`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *ProfileLocation) GetLatitudeOk() (*ProfileLocationLatitude, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *ProfileLocation) SetLatitude(v ProfileLocationLatitude)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *ProfileLocation) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *ProfileLocation) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *ProfileLocation) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *ProfileLocation) GetLongitude() ProfileLocationLongitude`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *ProfileLocation) GetLongitudeOk() (*ProfileLocationLongitude, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *ProfileLocation) SetLongitude(v ProfileLocationLongitude)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *ProfileLocation) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *ProfileLocation) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *ProfileLocation) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetRegion

`func (o *ProfileLocation) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ProfileLocation) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ProfileLocation) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ProfileLocation) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *ProfileLocation) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *ProfileLocation) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetZip

`func (o *ProfileLocation) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *ProfileLocation) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *ProfileLocation) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *ProfileLocation) HasZip() bool`

HasZip returns a boolean if a field has been set.

### SetZipNil

`func (o *ProfileLocation) SetZipNil(b bool)`

 SetZipNil sets the value for Zip to be an explicit nil

### UnsetZip
`func (o *ProfileLocation) UnsetZip()`

UnsetZip ensures that no value is present for Zip, not even an explicit nil
### GetTimezone

`func (o *ProfileLocation) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ProfileLocation) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ProfileLocation) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ProfileLocation) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *ProfileLocation) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *ProfileLocation) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetIp

`func (o *ProfileLocation) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ProfileLocation) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ProfileLocation) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ProfileLocation) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *ProfileLocation) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *ProfileLocation) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


