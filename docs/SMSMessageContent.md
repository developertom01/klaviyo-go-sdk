# SMSMessageContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **string** |  | 
**MediaUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSMSMessageContent

`func NewSMSMessageContent(body string, ) *SMSMessageContent`

NewSMSMessageContent instantiates a new SMSMessageContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMSMessageContentWithDefaults

`func NewSMSMessageContentWithDefaults() *SMSMessageContent`

NewSMSMessageContentWithDefaults instantiates a new SMSMessageContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *SMSMessageContent) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SMSMessageContent) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SMSMessageContent) SetBody(v string)`

SetBody sets Body field to given value.


### GetMediaUrl

`func (o *SMSMessageContent) GetMediaUrl() string`

GetMediaUrl returns the MediaUrl field if non-nil, zero value otherwise.

### GetMediaUrlOk

`func (o *SMSMessageContent) GetMediaUrlOk() (*string, bool)`

GetMediaUrlOk returns a tuple with the MediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrl

`func (o *SMSMessageContent) SetMediaUrl(v string)`

SetMediaUrl sets MediaUrl field to given value.

### HasMediaUrl

`func (o *SMSMessageContent) HasMediaUrl() bool`

HasMediaUrl returns a boolean if a field has been set.

### SetMediaUrlNil

`func (o *SMSMessageContent) SetMediaUrlNil(b bool)`

 SetMediaUrlNil sets the value for MediaUrl to be an explicit nil

### UnsetMediaUrl
`func (o *SMSMessageContent) UnsetMediaUrl()`

UnsetMediaUrl ensures that no value is present for MediaUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


