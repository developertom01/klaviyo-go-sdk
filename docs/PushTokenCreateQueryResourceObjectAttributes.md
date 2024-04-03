# PushTokenCreateQueryResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | A push token from APNS or FCM. | 
**Platform** | **string** | The platform on which the push token was created. | 
**EnablementStatus** | Pointer to **NullableString** | This is the enablement status for the individual push token. | [optional] [default to "AUTHORIZED"]
**Vendor** | **string** | The vendor of the push token. | 
**Background** | Pointer to **NullableString** | The background state of the push token. | [optional] [default to "AVAILABLE"]
**DeviceMetadata** | Pointer to [**DeviceMetadata**](DeviceMetadata.md) |  | [optional] 
**Profile** | [**PushTokenCreateQueryResourceObjectAttributesProfile**](PushTokenCreateQueryResourceObjectAttributesProfile.md) |  | 

## Methods

### NewPushTokenCreateQueryResourceObjectAttributes

`func NewPushTokenCreateQueryResourceObjectAttributes(token string, platform string, vendor string, profile PushTokenCreateQueryResourceObjectAttributesProfile, ) *PushTokenCreateQueryResourceObjectAttributes`

NewPushTokenCreateQueryResourceObjectAttributes instantiates a new PushTokenCreateQueryResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushTokenCreateQueryResourceObjectAttributesWithDefaults

`func NewPushTokenCreateQueryResourceObjectAttributesWithDefaults() *PushTokenCreateQueryResourceObjectAttributes`

NewPushTokenCreateQueryResourceObjectAttributesWithDefaults instantiates a new PushTokenCreateQueryResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *PushTokenCreateQueryResourceObjectAttributes) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PushTokenCreateQueryResourceObjectAttributes) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PushTokenCreateQueryResourceObjectAttributes) SetToken(v string)`

SetToken sets Token field to given value.


### GetPlatform

`func (o *PushTokenCreateQueryResourceObjectAttributes) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PushTokenCreateQueryResourceObjectAttributes) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PushTokenCreateQueryResourceObjectAttributes) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetEnablementStatus

`func (o *PushTokenCreateQueryResourceObjectAttributes) GetEnablementStatus() string`

GetEnablementStatus returns the EnablementStatus field if non-nil, zero value otherwise.

### GetEnablementStatusOk

`func (o *PushTokenCreateQueryResourceObjectAttributes) GetEnablementStatusOk() (*string, bool)`

GetEnablementStatusOk returns a tuple with the EnablementStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablementStatus

`func (o *PushTokenCreateQueryResourceObjectAttributes) SetEnablementStatus(v string)`

SetEnablementStatus sets EnablementStatus field to given value.

### HasEnablementStatus

`func (o *PushTokenCreateQueryResourceObjectAttributes) HasEnablementStatus() bool`

HasEnablementStatus returns a boolean if a field has been set.

### SetEnablementStatusNil

`func (o *PushTokenCreateQueryResourceObjectAttributes) SetEnablementStatusNil(b bool)`

 SetEnablementStatusNil sets the value for EnablementStatus to be an explicit nil

### UnsetEnablementStatus
`func (o *PushTokenCreateQueryResourceObjectAttributes) UnsetEnablementStatus()`

UnsetEnablementStatus ensures that no value is present for EnablementStatus, not even an explicit nil
### GetVendor

`func (o *PushTokenCreateQueryResourceObjectAttributes) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *PushTokenCreateQueryResourceObjectAttributes) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *PushTokenCreateQueryResourceObjectAttributes) SetVendor(v string)`

SetVendor sets Vendor field to given value.


### GetBackground

`func (o *PushTokenCreateQueryResourceObjectAttributes) GetBackground() string`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *PushTokenCreateQueryResourceObjectAttributes) GetBackgroundOk() (*string, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *PushTokenCreateQueryResourceObjectAttributes) SetBackground(v string)`

SetBackground sets Background field to given value.

### HasBackground

`func (o *PushTokenCreateQueryResourceObjectAttributes) HasBackground() bool`

HasBackground returns a boolean if a field has been set.

### SetBackgroundNil

`func (o *PushTokenCreateQueryResourceObjectAttributes) SetBackgroundNil(b bool)`

 SetBackgroundNil sets the value for Background to be an explicit nil

### UnsetBackground
`func (o *PushTokenCreateQueryResourceObjectAttributes) UnsetBackground()`

UnsetBackground ensures that no value is present for Background, not even an explicit nil
### GetDeviceMetadata

`func (o *PushTokenCreateQueryResourceObjectAttributes) GetDeviceMetadata() DeviceMetadata`

GetDeviceMetadata returns the DeviceMetadata field if non-nil, zero value otherwise.

### GetDeviceMetadataOk

`func (o *PushTokenCreateQueryResourceObjectAttributes) GetDeviceMetadataOk() (*DeviceMetadata, bool)`

GetDeviceMetadataOk returns a tuple with the DeviceMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMetadata

`func (o *PushTokenCreateQueryResourceObjectAttributes) SetDeviceMetadata(v DeviceMetadata)`

SetDeviceMetadata sets DeviceMetadata field to given value.

### HasDeviceMetadata

`func (o *PushTokenCreateQueryResourceObjectAttributes) HasDeviceMetadata() bool`

HasDeviceMetadata returns a boolean if a field has been set.

### GetProfile

`func (o *PushTokenCreateQueryResourceObjectAttributes) GetProfile() PushTokenCreateQueryResourceObjectAttributesProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PushTokenCreateQueryResourceObjectAttributes) GetProfileOk() (*PushTokenCreateQueryResourceObjectAttributesProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PushTokenCreateQueryResourceObjectAttributes) SetProfile(v PushTokenCreateQueryResourceObjectAttributesProfile)`

SetProfile sets Profile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


