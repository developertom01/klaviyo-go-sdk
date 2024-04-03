# CampaignResponseObjectResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The campaign name | 
**Status** | **string** | The current status of the campaign | 
**Archived** | **bool** | Whether the campaign has been archived or not | 
**Audiences** | [**AudiencesSubObject**](AudiencesSubObject.md) |  | 
**SendOptions** | [**CampaignResponseObjectResourceAttributesSendOptions**](CampaignResponseObjectResourceAttributesSendOptions.md) |  | 
**TrackingOptions** | [**CampaignResponseObjectResourceAttributesTrackingOptions**](CampaignResponseObjectResourceAttributesTrackingOptions.md) |  | 
**SendStrategy** | [**SendStrategySubObject**](SendStrategySubObject.md) |  | 
**CreatedAt** | **time.Time** | The datetime when the campaign was created | 
**ScheduledAt** | **time.Time** | The datetime when the campaign was scheduled for future sending | 
**UpdatedAt** | **time.Time** | The datetime when the campaign was last updated by a user or the system | 
**SendTime** | **time.Time** | The datetime when the campaign will be / was sent or None if not yet scheduled by a send_job. | 

## Methods

### NewCampaignResponseObjectResourceAttributes

`func NewCampaignResponseObjectResourceAttributes(name string, status string, archived bool, audiences AudiencesSubObject, sendOptions CampaignResponseObjectResourceAttributesSendOptions, trackingOptions CampaignResponseObjectResourceAttributesTrackingOptions, sendStrategy SendStrategySubObject, createdAt time.Time, scheduledAt time.Time, updatedAt time.Time, sendTime time.Time, ) *CampaignResponseObjectResourceAttributes`

NewCampaignResponseObjectResourceAttributes instantiates a new CampaignResponseObjectResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignResponseObjectResourceAttributesWithDefaults

`func NewCampaignResponseObjectResourceAttributesWithDefaults() *CampaignResponseObjectResourceAttributes`

NewCampaignResponseObjectResourceAttributesWithDefaults instantiates a new CampaignResponseObjectResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CampaignResponseObjectResourceAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignResponseObjectResourceAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignResponseObjectResourceAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *CampaignResponseObjectResourceAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CampaignResponseObjectResourceAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CampaignResponseObjectResourceAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetArchived

`func (o *CampaignResponseObjectResourceAttributes) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *CampaignResponseObjectResourceAttributes) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *CampaignResponseObjectResourceAttributes) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetAudiences

`func (o *CampaignResponseObjectResourceAttributes) GetAudiences() AudiencesSubObject`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *CampaignResponseObjectResourceAttributes) GetAudiencesOk() (*AudiencesSubObject, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *CampaignResponseObjectResourceAttributes) SetAudiences(v AudiencesSubObject)`

SetAudiences sets Audiences field to given value.


### GetSendOptions

`func (o *CampaignResponseObjectResourceAttributes) GetSendOptions() CampaignResponseObjectResourceAttributesSendOptions`

GetSendOptions returns the SendOptions field if non-nil, zero value otherwise.

### GetSendOptionsOk

`func (o *CampaignResponseObjectResourceAttributes) GetSendOptionsOk() (*CampaignResponseObjectResourceAttributesSendOptions, bool)`

GetSendOptionsOk returns a tuple with the SendOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendOptions

`func (o *CampaignResponseObjectResourceAttributes) SetSendOptions(v CampaignResponseObjectResourceAttributesSendOptions)`

SetSendOptions sets SendOptions field to given value.


### GetTrackingOptions

`func (o *CampaignResponseObjectResourceAttributes) GetTrackingOptions() CampaignResponseObjectResourceAttributesTrackingOptions`

GetTrackingOptions returns the TrackingOptions field if non-nil, zero value otherwise.

### GetTrackingOptionsOk

`func (o *CampaignResponseObjectResourceAttributes) GetTrackingOptionsOk() (*CampaignResponseObjectResourceAttributesTrackingOptions, bool)`

GetTrackingOptionsOk returns a tuple with the TrackingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingOptions

`func (o *CampaignResponseObjectResourceAttributes) SetTrackingOptions(v CampaignResponseObjectResourceAttributesTrackingOptions)`

SetTrackingOptions sets TrackingOptions field to given value.


### GetSendStrategy

`func (o *CampaignResponseObjectResourceAttributes) GetSendStrategy() SendStrategySubObject`

GetSendStrategy returns the SendStrategy field if non-nil, zero value otherwise.

### GetSendStrategyOk

`func (o *CampaignResponseObjectResourceAttributes) GetSendStrategyOk() (*SendStrategySubObject, bool)`

GetSendStrategyOk returns a tuple with the SendStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendStrategy

`func (o *CampaignResponseObjectResourceAttributes) SetSendStrategy(v SendStrategySubObject)`

SetSendStrategy sets SendStrategy field to given value.


### GetCreatedAt

`func (o *CampaignResponseObjectResourceAttributes) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CampaignResponseObjectResourceAttributes) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CampaignResponseObjectResourceAttributes) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetScheduledAt

`func (o *CampaignResponseObjectResourceAttributes) GetScheduledAt() time.Time`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *CampaignResponseObjectResourceAttributes) GetScheduledAtOk() (*time.Time, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *CampaignResponseObjectResourceAttributes) SetScheduledAt(v time.Time)`

SetScheduledAt sets ScheduledAt field to given value.


### GetUpdatedAt

`func (o *CampaignResponseObjectResourceAttributes) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CampaignResponseObjectResourceAttributes) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CampaignResponseObjectResourceAttributes) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetSendTime

`func (o *CampaignResponseObjectResourceAttributes) GetSendTime() time.Time`

GetSendTime returns the SendTime field if non-nil, zero value otherwise.

### GetSendTimeOk

`func (o *CampaignResponseObjectResourceAttributes) GetSendTimeOk() (*time.Time, bool)`

GetSendTimeOk returns a tuple with the SendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendTime

`func (o *CampaignResponseObjectResourceAttributes) SetSendTime(v time.Time)`

SetSendTime sets SendTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


