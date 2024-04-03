# PushTokenUnregisterQueryResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | A push token from APNS or FCM. | 
**Platform** | **string** | The platform on which the push token was created. | 
**Vendor** | Pointer to **NullableString** | The vendor of the push token. | [optional] 
**Profile** | [**PushTokenCreateQueryResourceObjectAttributesProfile**](PushTokenCreateQueryResourceObjectAttributesProfile.md) |  | 

## Methods

### NewPushTokenUnregisterQueryResourceObjectAttributes

`func NewPushTokenUnregisterQueryResourceObjectAttributes(token string, platform string, profile PushTokenCreateQueryResourceObjectAttributesProfile, ) *PushTokenUnregisterQueryResourceObjectAttributes`

NewPushTokenUnregisterQueryResourceObjectAttributes instantiates a new PushTokenUnregisterQueryResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushTokenUnregisterQueryResourceObjectAttributesWithDefaults

`func NewPushTokenUnregisterQueryResourceObjectAttributesWithDefaults() *PushTokenUnregisterQueryResourceObjectAttributes`

NewPushTokenUnregisterQueryResourceObjectAttributesWithDefaults instantiates a new PushTokenUnregisterQueryResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *PushTokenUnregisterQueryResourceObjectAttributes) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PushTokenUnregisterQueryResourceObjectAttributes) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PushTokenUnregisterQueryResourceObjectAttributes) SetToken(v string)`

SetToken sets Token field to given value.


### GetPlatform

`func (o *PushTokenUnregisterQueryResourceObjectAttributes) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PushTokenUnregisterQueryResourceObjectAttributes) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PushTokenUnregisterQueryResourceObjectAttributes) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetVendor

`func (o *PushTokenUnregisterQueryResourceObjectAttributes) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *PushTokenUnregisterQueryResourceObjectAttributes) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *PushTokenUnregisterQueryResourceObjectAttributes) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *PushTokenUnregisterQueryResourceObjectAttributes) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### SetVendorNil

`func (o *PushTokenUnregisterQueryResourceObjectAttributes) SetVendorNil(b bool)`

 SetVendorNil sets the value for Vendor to be an explicit nil

### UnsetVendor
`func (o *PushTokenUnregisterQueryResourceObjectAttributes) UnsetVendor()`

UnsetVendor ensures that no value is present for Vendor, not even an explicit nil
### GetProfile

`func (o *PushTokenUnregisterQueryResourceObjectAttributes) GetProfile() PushTokenCreateQueryResourceObjectAttributesProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PushTokenUnregisterQueryResourceObjectAttributes) GetProfileOk() (*PushTokenCreateQueryResourceObjectAttributesProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PushTokenUnregisterQueryResourceObjectAttributes) SetProfile(v PushTokenCreateQueryResourceObjectAttributesProfile)`

SetProfile sets Profile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


