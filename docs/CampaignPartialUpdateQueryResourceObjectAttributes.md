# CampaignPartialUpdateQueryResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The campaign name | [optional] 
**Audiences** | Pointer to [**AudiencesSubObject**](AudiencesSubObject.md) |  | [optional] 
**SendOptions** | Pointer to [**NullableCampaignCreateQueryResourceObjectAttributesSendOptions**](CampaignCreateQueryResourceObjectAttributesSendOptions.md) |  | [optional] 
**TrackingOptions** | Pointer to [**NullableCampaignCreateQueryResourceObjectAttributesTrackingOptions**](CampaignCreateQueryResourceObjectAttributesTrackingOptions.md) |  | [optional] 
**SendStrategy** | Pointer to [**SendStrategySubObject**](SendStrategySubObject.md) |  | [optional] 

## Methods

### NewCampaignPartialUpdateQueryResourceObjectAttributes

`func NewCampaignPartialUpdateQueryResourceObjectAttributes() *CampaignPartialUpdateQueryResourceObjectAttributes`

NewCampaignPartialUpdateQueryResourceObjectAttributes instantiates a new CampaignPartialUpdateQueryResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignPartialUpdateQueryResourceObjectAttributesWithDefaults

`func NewCampaignPartialUpdateQueryResourceObjectAttributesWithDefaults() *CampaignPartialUpdateQueryResourceObjectAttributes`

NewCampaignPartialUpdateQueryResourceObjectAttributesWithDefaults instantiates a new CampaignPartialUpdateQueryResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CampaignPartialUpdateQueryResourceObjectAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignPartialUpdateQueryResourceObjectAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignPartialUpdateQueryResourceObjectAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CampaignPartialUpdateQueryResourceObjectAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CampaignPartialUpdateQueryResourceObjectAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CampaignPartialUpdateQueryResourceObjectAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAudiences

`func (o *CampaignPartialUpdateQueryResourceObjectAttributes) GetAudiences() AudiencesSubObject`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *CampaignPartialUpdateQueryResourceObjectAttributes) GetAudiencesOk() (*AudiencesSubObject, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *CampaignPartialUpdateQueryResourceObjectAttributes) SetAudiences(v AudiencesSubObject)`

SetAudiences sets Audiences field to given value.

### HasAudiences

`func (o *CampaignPartialUpdateQueryResourceObjectAttributes) HasAudiences() bool`

HasAudiences returns a boolean if a field has been set.

### GetSendOptions

`func (o *CampaignPartialUpdateQueryResourceObjectAttributes) GetSendOptions() CampaignCreateQueryResourceObjectAttributesSendOptions`

GetSendOptions returns the SendOptions field if non-nil, zero value otherwise.

### GetSendOptionsOk

`func (o *CampaignPartialUpdateQueryResourceObjectAttributes) GetSendOptionsOk() (*CampaignCreateQueryResourceObjectAttributesSendOptions, bool)`

GetSendOptionsOk returns a tuple with the SendOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendOptions

`func (o *CampaignPartialUpdateQueryResourceObjectAttributes) SetSendOptions(v CampaignCreateQueryResourceObjectAttributesSendOptions)`

SetSendOptions sets SendOptions field to given value.

### HasSendOptions

`func (o *CampaignPartialUpdateQueryResourceObjectAttributes) HasSendOptions() bool`

HasSendOptions returns a boolean if a field has been set.

### SetSendOptionsNil

`func (o *CampaignPartialUpdateQueryResourceObjectAttributes) SetSendOptionsNil(b bool)`

 SetSendOptionsNil sets the value for SendOptions to be an explicit nil

### UnsetSendOptions
`func (o *CampaignPartialUpdateQueryResourceObjectAttributes) UnsetSendOptions()`

UnsetSendOptions ensures that no value is present for SendOptions, not even an explicit nil
### GetTrackingOptions

`func (o *CampaignPartialUpdateQueryResourceObjectAttributes) GetTrackingOptions() CampaignCreateQueryResourceObjectAttributesTrackingOptions`

GetTrackingOptions returns the TrackingOptions field if non-nil, zero value otherwise.

### GetTrackingOptionsOk

`func (o *CampaignPartialUpdateQueryResourceObjectAttributes) GetTrackingOptionsOk() (*CampaignCreateQueryResourceObjectAttributesTrackingOptions, bool)`

GetTrackingOptionsOk returns a tuple with the TrackingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingOptions

`func (o *CampaignPartialUpdateQueryResourceObjectAttributes) SetTrackingOptions(v CampaignCreateQueryResourceObjectAttributesTrackingOptions)`

SetTrackingOptions sets TrackingOptions field to given value.

### HasTrackingOptions

`func (o *CampaignPartialUpdateQueryResourceObjectAttributes) HasTrackingOptions() bool`

HasTrackingOptions returns a boolean if a field has been set.

### SetTrackingOptionsNil

`func (o *CampaignPartialUpdateQueryResourceObjectAttributes) SetTrackingOptionsNil(b bool)`

 SetTrackingOptionsNil sets the value for TrackingOptions to be an explicit nil

### UnsetTrackingOptions
`func (o *CampaignPartialUpdateQueryResourceObjectAttributes) UnsetTrackingOptions()`

UnsetTrackingOptions ensures that no value is present for TrackingOptions, not even an explicit nil
### GetSendStrategy

`func (o *CampaignPartialUpdateQueryResourceObjectAttributes) GetSendStrategy() SendStrategySubObject`

GetSendStrategy returns the SendStrategy field if non-nil, zero value otherwise.

### GetSendStrategyOk

`func (o *CampaignPartialUpdateQueryResourceObjectAttributes) GetSendStrategyOk() (*SendStrategySubObject, bool)`

GetSendStrategyOk returns a tuple with the SendStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendStrategy

`func (o *CampaignPartialUpdateQueryResourceObjectAttributes) SetSendStrategy(v SendStrategySubObject)`

SetSendStrategy sets SendStrategy field to given value.

### HasSendStrategy

`func (o *CampaignPartialUpdateQueryResourceObjectAttributes) HasSendStrategy() bool`

HasSendStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


