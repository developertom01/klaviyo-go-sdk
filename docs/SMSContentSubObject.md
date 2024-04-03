# SMSContentSubObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **NullableString** | The message body | [optional] 
**MediaUrl** | Pointer to **NullableString** | URL for included media | [optional] 

## Methods

### NewSMSContentSubObject

`func NewSMSContentSubObject() *SMSContentSubObject`

NewSMSContentSubObject instantiates a new SMSContentSubObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMSContentSubObjectWithDefaults

`func NewSMSContentSubObjectWithDefaults() *SMSContentSubObject`

NewSMSContentSubObjectWithDefaults instantiates a new SMSContentSubObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *SMSContentSubObject) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SMSContentSubObject) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SMSContentSubObject) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *SMSContentSubObject) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *SMSContentSubObject) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *SMSContentSubObject) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetMediaUrl

`func (o *SMSContentSubObject) GetMediaUrl() string`

GetMediaUrl returns the MediaUrl field if non-nil, zero value otherwise.

### GetMediaUrlOk

`func (o *SMSContentSubObject) GetMediaUrlOk() (*string, bool)`

GetMediaUrlOk returns a tuple with the MediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrl

`func (o *SMSContentSubObject) SetMediaUrl(v string)`

SetMediaUrl sets MediaUrl field to given value.

### HasMediaUrl

`func (o *SMSContentSubObject) HasMediaUrl() bool`

HasMediaUrl returns a boolean if a field has been set.

### SetMediaUrlNil

`func (o *SMSContentSubObject) SetMediaUrlNil(b bool)`

 SetMediaUrlNil sets the value for MediaUrl to be an explicit nil

### UnsetMediaUrl
`func (o *SMSContentSubObject) UnsetMediaUrl()`

UnsetMediaUrl ensures that no value is present for MediaUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


