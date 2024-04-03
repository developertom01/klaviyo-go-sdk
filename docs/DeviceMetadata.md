# DeviceMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **NullableString** | Relatively stable ID for the device. Will update on app uninstall and reinstall | [optional] 
**KlaviyoSdk** | Pointer to **NullableString** | The name of the SDK used to create the push token. | [optional] 
**SdkVersion** | Pointer to **NullableString** | The version of the SDK used to create the push token | [optional] 
**DeviceModel** | Pointer to **NullableString** | The model of the device | [optional] 
**OsName** | Pointer to **NullableString** | The name of the operating system on the device. | [optional] 
**OsVersion** | Pointer to **NullableString** | The version of the operating system on the device | [optional] 
**Manufacturer** | Pointer to **NullableString** | The manufacturer of the device | [optional] 
**AppName** | Pointer to **NullableString** | The name of the app that created the push token | [optional] 
**AppVersion** | Pointer to **NullableString** | The version of the app that created the push token | [optional] 
**AppBuild** | Pointer to **NullableString** | The build of the app that created the push token | [optional] 
**AppId** | Pointer to **NullableString** | The ID of the app that created the push token | [optional] 
**Environment** | Pointer to **NullableString** | The environment in which the push token was created | [optional] 

## Methods

### NewDeviceMetadata

`func NewDeviceMetadata() *DeviceMetadata`

NewDeviceMetadata instantiates a new DeviceMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceMetadataWithDefaults

`func NewDeviceMetadataWithDefaults() *DeviceMetadata`

NewDeviceMetadataWithDefaults instantiates a new DeviceMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *DeviceMetadata) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DeviceMetadata) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DeviceMetadata) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *DeviceMetadata) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *DeviceMetadata) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *DeviceMetadata) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetKlaviyoSdk

`func (o *DeviceMetadata) GetKlaviyoSdk() string`

GetKlaviyoSdk returns the KlaviyoSdk field if non-nil, zero value otherwise.

### GetKlaviyoSdkOk

`func (o *DeviceMetadata) GetKlaviyoSdkOk() (*string, bool)`

GetKlaviyoSdkOk returns a tuple with the KlaviyoSdk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKlaviyoSdk

`func (o *DeviceMetadata) SetKlaviyoSdk(v string)`

SetKlaviyoSdk sets KlaviyoSdk field to given value.

### HasKlaviyoSdk

`func (o *DeviceMetadata) HasKlaviyoSdk() bool`

HasKlaviyoSdk returns a boolean if a field has been set.

### SetKlaviyoSdkNil

`func (o *DeviceMetadata) SetKlaviyoSdkNil(b bool)`

 SetKlaviyoSdkNil sets the value for KlaviyoSdk to be an explicit nil

### UnsetKlaviyoSdk
`func (o *DeviceMetadata) UnsetKlaviyoSdk()`

UnsetKlaviyoSdk ensures that no value is present for KlaviyoSdk, not even an explicit nil
### GetSdkVersion

`func (o *DeviceMetadata) GetSdkVersion() string`

GetSdkVersion returns the SdkVersion field if non-nil, zero value otherwise.

### GetSdkVersionOk

`func (o *DeviceMetadata) GetSdkVersionOk() (*string, bool)`

GetSdkVersionOk returns a tuple with the SdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkVersion

`func (o *DeviceMetadata) SetSdkVersion(v string)`

SetSdkVersion sets SdkVersion field to given value.

### HasSdkVersion

`func (o *DeviceMetadata) HasSdkVersion() bool`

HasSdkVersion returns a boolean if a field has been set.

### SetSdkVersionNil

`func (o *DeviceMetadata) SetSdkVersionNil(b bool)`

 SetSdkVersionNil sets the value for SdkVersion to be an explicit nil

### UnsetSdkVersion
`func (o *DeviceMetadata) UnsetSdkVersion()`

UnsetSdkVersion ensures that no value is present for SdkVersion, not even an explicit nil
### GetDeviceModel

`func (o *DeviceMetadata) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *DeviceMetadata) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *DeviceMetadata) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *DeviceMetadata) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### SetDeviceModelNil

`func (o *DeviceMetadata) SetDeviceModelNil(b bool)`

 SetDeviceModelNil sets the value for DeviceModel to be an explicit nil

### UnsetDeviceModel
`func (o *DeviceMetadata) UnsetDeviceModel()`

UnsetDeviceModel ensures that no value is present for DeviceModel, not even an explicit nil
### GetOsName

`func (o *DeviceMetadata) GetOsName() string`

GetOsName returns the OsName field if non-nil, zero value otherwise.

### GetOsNameOk

`func (o *DeviceMetadata) GetOsNameOk() (*string, bool)`

GetOsNameOk returns a tuple with the OsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsName

`func (o *DeviceMetadata) SetOsName(v string)`

SetOsName sets OsName field to given value.

### HasOsName

`func (o *DeviceMetadata) HasOsName() bool`

HasOsName returns a boolean if a field has been set.

### SetOsNameNil

`func (o *DeviceMetadata) SetOsNameNil(b bool)`

 SetOsNameNil sets the value for OsName to be an explicit nil

### UnsetOsName
`func (o *DeviceMetadata) UnsetOsName()`

UnsetOsName ensures that no value is present for OsName, not even an explicit nil
### GetOsVersion

`func (o *DeviceMetadata) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *DeviceMetadata) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *DeviceMetadata) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *DeviceMetadata) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### SetOsVersionNil

`func (o *DeviceMetadata) SetOsVersionNil(b bool)`

 SetOsVersionNil sets the value for OsVersion to be an explicit nil

### UnsetOsVersion
`func (o *DeviceMetadata) UnsetOsVersion()`

UnsetOsVersion ensures that no value is present for OsVersion, not even an explicit nil
### GetManufacturer

`func (o *DeviceMetadata) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *DeviceMetadata) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *DeviceMetadata) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *DeviceMetadata) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *DeviceMetadata) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *DeviceMetadata) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetAppName

`func (o *DeviceMetadata) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *DeviceMetadata) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *DeviceMetadata) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *DeviceMetadata) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### SetAppNameNil

`func (o *DeviceMetadata) SetAppNameNil(b bool)`

 SetAppNameNil sets the value for AppName to be an explicit nil

### UnsetAppName
`func (o *DeviceMetadata) UnsetAppName()`

UnsetAppName ensures that no value is present for AppName, not even an explicit nil
### GetAppVersion

`func (o *DeviceMetadata) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *DeviceMetadata) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *DeviceMetadata) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *DeviceMetadata) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### SetAppVersionNil

`func (o *DeviceMetadata) SetAppVersionNil(b bool)`

 SetAppVersionNil sets the value for AppVersion to be an explicit nil

### UnsetAppVersion
`func (o *DeviceMetadata) UnsetAppVersion()`

UnsetAppVersion ensures that no value is present for AppVersion, not even an explicit nil
### GetAppBuild

`func (o *DeviceMetadata) GetAppBuild() string`

GetAppBuild returns the AppBuild field if non-nil, zero value otherwise.

### GetAppBuildOk

`func (o *DeviceMetadata) GetAppBuildOk() (*string, bool)`

GetAppBuildOk returns a tuple with the AppBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppBuild

`func (o *DeviceMetadata) SetAppBuild(v string)`

SetAppBuild sets AppBuild field to given value.

### HasAppBuild

`func (o *DeviceMetadata) HasAppBuild() bool`

HasAppBuild returns a boolean if a field has been set.

### SetAppBuildNil

`func (o *DeviceMetadata) SetAppBuildNil(b bool)`

 SetAppBuildNil sets the value for AppBuild to be an explicit nil

### UnsetAppBuild
`func (o *DeviceMetadata) UnsetAppBuild()`

UnsetAppBuild ensures that no value is present for AppBuild, not even an explicit nil
### GetAppId

`func (o *DeviceMetadata) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *DeviceMetadata) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *DeviceMetadata) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *DeviceMetadata) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### SetAppIdNil

`func (o *DeviceMetadata) SetAppIdNil(b bool)`

 SetAppIdNil sets the value for AppId to be an explicit nil

### UnsetAppId
`func (o *DeviceMetadata) UnsetAppId()`

UnsetAppId ensures that no value is present for AppId, not even an explicit nil
### GetEnvironment

`func (o *DeviceMetadata) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DeviceMetadata) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DeviceMetadata) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *DeviceMetadata) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *DeviceMetadata) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *DeviceMetadata) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


