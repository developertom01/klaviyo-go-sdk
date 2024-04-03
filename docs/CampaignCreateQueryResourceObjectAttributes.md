# CampaignCreateQueryResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The campaign name | 
**Audiences** | [**AudiencesSubObject**](AudiencesSubObject.md) |  | 
**SendStrategy** | Pointer to [**SendStrategySubObject**](SendStrategySubObject.md) |  | [optional] 
**SendOptions** | Pointer to [**NullableCampaignCreateQueryResourceObjectAttributesSendOptions**](CampaignCreateQueryResourceObjectAttributesSendOptions.md) |  | [optional] 
**TrackingOptions** | Pointer to [**NullableCampaignCreateQueryResourceObjectAttributesTrackingOptions**](CampaignCreateQueryResourceObjectAttributesTrackingOptions.md) |  | [optional] 
**CampaignMessages** | [**CampaignCreateQueryResourceObjectAttributesCampaignMessages**](CampaignCreateQueryResourceObjectAttributesCampaignMessages.md) |  | 

## Methods

### NewCampaignCreateQueryResourceObjectAttributes

`func NewCampaignCreateQueryResourceObjectAttributes(name string, audiences AudiencesSubObject, campaignMessages CampaignCreateQueryResourceObjectAttributesCampaignMessages, ) *CampaignCreateQueryResourceObjectAttributes`

NewCampaignCreateQueryResourceObjectAttributes instantiates a new CampaignCreateQueryResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignCreateQueryResourceObjectAttributesWithDefaults

`func NewCampaignCreateQueryResourceObjectAttributesWithDefaults() *CampaignCreateQueryResourceObjectAttributes`

NewCampaignCreateQueryResourceObjectAttributesWithDefaults instantiates a new CampaignCreateQueryResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CampaignCreateQueryResourceObjectAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignCreateQueryResourceObjectAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignCreateQueryResourceObjectAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetAudiences

`func (o *CampaignCreateQueryResourceObjectAttributes) GetAudiences() AudiencesSubObject`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *CampaignCreateQueryResourceObjectAttributes) GetAudiencesOk() (*AudiencesSubObject, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *CampaignCreateQueryResourceObjectAttributes) SetAudiences(v AudiencesSubObject)`

SetAudiences sets Audiences field to given value.


### GetSendStrategy

`func (o *CampaignCreateQueryResourceObjectAttributes) GetSendStrategy() SendStrategySubObject`

GetSendStrategy returns the SendStrategy field if non-nil, zero value otherwise.

### GetSendStrategyOk

`func (o *CampaignCreateQueryResourceObjectAttributes) GetSendStrategyOk() (*SendStrategySubObject, bool)`

GetSendStrategyOk returns a tuple with the SendStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendStrategy

`func (o *CampaignCreateQueryResourceObjectAttributes) SetSendStrategy(v SendStrategySubObject)`

SetSendStrategy sets SendStrategy field to given value.

### HasSendStrategy

`func (o *CampaignCreateQueryResourceObjectAttributes) HasSendStrategy() bool`

HasSendStrategy returns a boolean if a field has been set.

### GetSendOptions

`func (o *CampaignCreateQueryResourceObjectAttributes) GetSendOptions() CampaignCreateQueryResourceObjectAttributesSendOptions`

GetSendOptions returns the SendOptions field if non-nil, zero value otherwise.

### GetSendOptionsOk

`func (o *CampaignCreateQueryResourceObjectAttributes) GetSendOptionsOk() (*CampaignCreateQueryResourceObjectAttributesSendOptions, bool)`

GetSendOptionsOk returns a tuple with the SendOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendOptions

`func (o *CampaignCreateQueryResourceObjectAttributes) SetSendOptions(v CampaignCreateQueryResourceObjectAttributesSendOptions)`

SetSendOptions sets SendOptions field to given value.

### HasSendOptions

`func (o *CampaignCreateQueryResourceObjectAttributes) HasSendOptions() bool`

HasSendOptions returns a boolean if a field has been set.

### SetSendOptionsNil

`func (o *CampaignCreateQueryResourceObjectAttributes) SetSendOptionsNil(b bool)`

 SetSendOptionsNil sets the value for SendOptions to be an explicit nil

### UnsetSendOptions
`func (o *CampaignCreateQueryResourceObjectAttributes) UnsetSendOptions()`

UnsetSendOptions ensures that no value is present for SendOptions, not even an explicit nil
### GetTrackingOptions

`func (o *CampaignCreateQueryResourceObjectAttributes) GetTrackingOptions() CampaignCreateQueryResourceObjectAttributesTrackingOptions`

GetTrackingOptions returns the TrackingOptions field if non-nil, zero value otherwise.

### GetTrackingOptionsOk

`func (o *CampaignCreateQueryResourceObjectAttributes) GetTrackingOptionsOk() (*CampaignCreateQueryResourceObjectAttributesTrackingOptions, bool)`

GetTrackingOptionsOk returns a tuple with the TrackingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingOptions

`func (o *CampaignCreateQueryResourceObjectAttributes) SetTrackingOptions(v CampaignCreateQueryResourceObjectAttributesTrackingOptions)`

SetTrackingOptions sets TrackingOptions field to given value.

### HasTrackingOptions

`func (o *CampaignCreateQueryResourceObjectAttributes) HasTrackingOptions() bool`

HasTrackingOptions returns a boolean if a field has been set.

### SetTrackingOptionsNil

`func (o *CampaignCreateQueryResourceObjectAttributes) SetTrackingOptionsNil(b bool)`

 SetTrackingOptionsNil sets the value for TrackingOptions to be an explicit nil

### UnsetTrackingOptions
`func (o *CampaignCreateQueryResourceObjectAttributes) UnsetTrackingOptions()`

UnsetTrackingOptions ensures that no value is present for TrackingOptions, not even an explicit nil
### GetCampaignMessages

`func (o *CampaignCreateQueryResourceObjectAttributes) GetCampaignMessages() CampaignCreateQueryResourceObjectAttributesCampaignMessages`

GetCampaignMessages returns the CampaignMessages field if non-nil, zero value otherwise.

### GetCampaignMessagesOk

`func (o *CampaignCreateQueryResourceObjectAttributes) GetCampaignMessagesOk() (*CampaignCreateQueryResourceObjectAttributesCampaignMessages, bool)`

GetCampaignMessagesOk returns a tuple with the CampaignMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignMessages

`func (o *CampaignCreateQueryResourceObjectAttributes) SetCampaignMessages(v CampaignCreateQueryResourceObjectAttributesCampaignMessages)`

SetCampaignMessages sets CampaignMessages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


