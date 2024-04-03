# CampaignMessageCreateQueryResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** | The channel the message is to be sent on (email or sms, for example) | 
**Label** | Pointer to **NullableString** | The label or name on the message | [optional] 
**Content** | Pointer to [**NullableCampaignMessageCreateQueryResourceObjectAttributesContent**](CampaignMessageCreateQueryResourceObjectAttributesContent.md) |  | [optional] 
**RenderOptions** | Pointer to [**RenderOptionsSubObject**](RenderOptionsSubObject.md) |  | [optional] 

## Methods

### NewCampaignMessageCreateQueryResourceObjectAttributes

`func NewCampaignMessageCreateQueryResourceObjectAttributes(channel string, ) *CampaignMessageCreateQueryResourceObjectAttributes`

NewCampaignMessageCreateQueryResourceObjectAttributes instantiates a new CampaignMessageCreateQueryResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignMessageCreateQueryResourceObjectAttributesWithDefaults

`func NewCampaignMessageCreateQueryResourceObjectAttributesWithDefaults() *CampaignMessageCreateQueryResourceObjectAttributes`

NewCampaignMessageCreateQueryResourceObjectAttributesWithDefaults instantiates a new CampaignMessageCreateQueryResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *CampaignMessageCreateQueryResourceObjectAttributes) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CampaignMessageCreateQueryResourceObjectAttributes) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CampaignMessageCreateQueryResourceObjectAttributes) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetLabel

`func (o *CampaignMessageCreateQueryResourceObjectAttributes) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CampaignMessageCreateQueryResourceObjectAttributes) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CampaignMessageCreateQueryResourceObjectAttributes) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CampaignMessageCreateQueryResourceObjectAttributes) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *CampaignMessageCreateQueryResourceObjectAttributes) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *CampaignMessageCreateQueryResourceObjectAttributes) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetContent

`func (o *CampaignMessageCreateQueryResourceObjectAttributes) GetContent() CampaignMessageCreateQueryResourceObjectAttributesContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CampaignMessageCreateQueryResourceObjectAttributes) GetContentOk() (*CampaignMessageCreateQueryResourceObjectAttributesContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CampaignMessageCreateQueryResourceObjectAttributes) SetContent(v CampaignMessageCreateQueryResourceObjectAttributesContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *CampaignMessageCreateQueryResourceObjectAttributes) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *CampaignMessageCreateQueryResourceObjectAttributes) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *CampaignMessageCreateQueryResourceObjectAttributes) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetRenderOptions

`func (o *CampaignMessageCreateQueryResourceObjectAttributes) GetRenderOptions() RenderOptionsSubObject`

GetRenderOptions returns the RenderOptions field if non-nil, zero value otherwise.

### GetRenderOptionsOk

`func (o *CampaignMessageCreateQueryResourceObjectAttributes) GetRenderOptionsOk() (*RenderOptionsSubObject, bool)`

GetRenderOptionsOk returns a tuple with the RenderOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderOptions

`func (o *CampaignMessageCreateQueryResourceObjectAttributes) SetRenderOptions(v RenderOptionsSubObject)`

SetRenderOptions sets RenderOptions field to given value.

### HasRenderOptions

`func (o *CampaignMessageCreateQueryResourceObjectAttributes) HasRenderOptions() bool`

HasRenderOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


