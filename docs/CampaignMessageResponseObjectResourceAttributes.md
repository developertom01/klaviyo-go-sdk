# CampaignMessageResponseObjectResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | The label or name on the message | 
**Channel** | **string** | The channel the message is to be sent on | 
**Content** | [**CampaignMessageResponseObjectResourceAttributesContent**](CampaignMessageResponseObjectResourceAttributesContent.md) |  | 
**SendTimes** | Pointer to [**[]SendTimeSubObject**](SendTimeSubObject.md) | The list of appropriate Send Time Sub-objects associated with the message | [optional] 
**RenderOptions** | Pointer to [**RenderOptionsSubObject**](RenderOptionsSubObject.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** | The datetime when the message was created | [optional] 
**UpdatedAt** | Pointer to **NullableTime** | The datetime when the message was last updated | [optional] 

## Methods

### NewCampaignMessageResponseObjectResourceAttributes

`func NewCampaignMessageResponseObjectResourceAttributes(label string, channel string, content CampaignMessageResponseObjectResourceAttributesContent, ) *CampaignMessageResponseObjectResourceAttributes`

NewCampaignMessageResponseObjectResourceAttributes instantiates a new CampaignMessageResponseObjectResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignMessageResponseObjectResourceAttributesWithDefaults

`func NewCampaignMessageResponseObjectResourceAttributesWithDefaults() *CampaignMessageResponseObjectResourceAttributes`

NewCampaignMessageResponseObjectResourceAttributesWithDefaults instantiates a new CampaignMessageResponseObjectResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CampaignMessageResponseObjectResourceAttributes) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CampaignMessageResponseObjectResourceAttributes) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CampaignMessageResponseObjectResourceAttributes) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetChannel

`func (o *CampaignMessageResponseObjectResourceAttributes) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CampaignMessageResponseObjectResourceAttributes) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CampaignMessageResponseObjectResourceAttributes) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetContent

`func (o *CampaignMessageResponseObjectResourceAttributes) GetContent() CampaignMessageResponseObjectResourceAttributesContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CampaignMessageResponseObjectResourceAttributes) GetContentOk() (*CampaignMessageResponseObjectResourceAttributesContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CampaignMessageResponseObjectResourceAttributes) SetContent(v CampaignMessageResponseObjectResourceAttributesContent)`

SetContent sets Content field to given value.


### GetSendTimes

`func (o *CampaignMessageResponseObjectResourceAttributes) GetSendTimes() []SendTimeSubObject`

GetSendTimes returns the SendTimes field if non-nil, zero value otherwise.

### GetSendTimesOk

`func (o *CampaignMessageResponseObjectResourceAttributes) GetSendTimesOk() (*[]SendTimeSubObject, bool)`

GetSendTimesOk returns a tuple with the SendTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendTimes

`func (o *CampaignMessageResponseObjectResourceAttributes) SetSendTimes(v []SendTimeSubObject)`

SetSendTimes sets SendTimes field to given value.

### HasSendTimes

`func (o *CampaignMessageResponseObjectResourceAttributes) HasSendTimes() bool`

HasSendTimes returns a boolean if a field has been set.

### SetSendTimesNil

`func (o *CampaignMessageResponseObjectResourceAttributes) SetSendTimesNil(b bool)`

 SetSendTimesNil sets the value for SendTimes to be an explicit nil

### UnsetSendTimes
`func (o *CampaignMessageResponseObjectResourceAttributes) UnsetSendTimes()`

UnsetSendTimes ensures that no value is present for SendTimes, not even an explicit nil
### GetRenderOptions

`func (o *CampaignMessageResponseObjectResourceAttributes) GetRenderOptions() RenderOptionsSubObject`

GetRenderOptions returns the RenderOptions field if non-nil, zero value otherwise.

### GetRenderOptionsOk

`func (o *CampaignMessageResponseObjectResourceAttributes) GetRenderOptionsOk() (*RenderOptionsSubObject, bool)`

GetRenderOptionsOk returns a tuple with the RenderOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderOptions

`func (o *CampaignMessageResponseObjectResourceAttributes) SetRenderOptions(v RenderOptionsSubObject)`

SetRenderOptions sets RenderOptions field to given value.

### HasRenderOptions

`func (o *CampaignMessageResponseObjectResourceAttributes) HasRenderOptions() bool`

HasRenderOptions returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CampaignMessageResponseObjectResourceAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CampaignMessageResponseObjectResourceAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CampaignMessageResponseObjectResourceAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CampaignMessageResponseObjectResourceAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *CampaignMessageResponseObjectResourceAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *CampaignMessageResponseObjectResourceAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *CampaignMessageResponseObjectResourceAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CampaignMessageResponseObjectResourceAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CampaignMessageResponseObjectResourceAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CampaignMessageResponseObjectResourceAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *CampaignMessageResponseObjectResourceAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *CampaignMessageResponseObjectResourceAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


