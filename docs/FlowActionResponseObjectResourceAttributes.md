# FlowActionResponseObjectResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Created** | Pointer to **NullableTime** |  | [optional] 
**Updated** | Pointer to **NullableTime** |  | [optional] 
**Settings** | Pointer to **map[string]interface{}** |  | [optional] 
**TrackingOptions** | Pointer to [**NullableFlowActionResponseObjectResourceAttributesTrackingOptions**](FlowActionResponseObjectResourceAttributesTrackingOptions.md) |  | [optional] 
**SendOptions** | Pointer to [**SendOptions**](SendOptions.md) |  | [optional] 
**RenderOptions** | Pointer to [**SMSRenderOptions**](SMSRenderOptions.md) |  | [optional] 

## Methods

### NewFlowActionResponseObjectResourceAttributes

`func NewFlowActionResponseObjectResourceAttributes() *FlowActionResponseObjectResourceAttributes`

NewFlowActionResponseObjectResourceAttributes instantiates a new FlowActionResponseObjectResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowActionResponseObjectResourceAttributesWithDefaults

`func NewFlowActionResponseObjectResourceAttributesWithDefaults() *FlowActionResponseObjectResourceAttributes`

NewFlowActionResponseObjectResourceAttributesWithDefaults instantiates a new FlowActionResponseObjectResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *FlowActionResponseObjectResourceAttributes) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *FlowActionResponseObjectResourceAttributes) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *FlowActionResponseObjectResourceAttributes) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *FlowActionResponseObjectResourceAttributes) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### SetActionTypeNil

`func (o *FlowActionResponseObjectResourceAttributes) SetActionTypeNil(b bool)`

 SetActionTypeNil sets the value for ActionType to be an explicit nil

### UnsetActionType
`func (o *FlowActionResponseObjectResourceAttributes) UnsetActionType()`

UnsetActionType ensures that no value is present for ActionType, not even an explicit nil
### GetStatus

`func (o *FlowActionResponseObjectResourceAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlowActionResponseObjectResourceAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlowActionResponseObjectResourceAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FlowActionResponseObjectResourceAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *FlowActionResponseObjectResourceAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *FlowActionResponseObjectResourceAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetCreated

`func (o *FlowActionResponseObjectResourceAttributes) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FlowActionResponseObjectResourceAttributes) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FlowActionResponseObjectResourceAttributes) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FlowActionResponseObjectResourceAttributes) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *FlowActionResponseObjectResourceAttributes) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *FlowActionResponseObjectResourceAttributes) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetUpdated

`func (o *FlowActionResponseObjectResourceAttributes) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *FlowActionResponseObjectResourceAttributes) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *FlowActionResponseObjectResourceAttributes) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *FlowActionResponseObjectResourceAttributes) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *FlowActionResponseObjectResourceAttributes) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *FlowActionResponseObjectResourceAttributes) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
### GetSettings

`func (o *FlowActionResponseObjectResourceAttributes) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *FlowActionResponseObjectResourceAttributes) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *FlowActionResponseObjectResourceAttributes) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *FlowActionResponseObjectResourceAttributes) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *FlowActionResponseObjectResourceAttributes) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *FlowActionResponseObjectResourceAttributes) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil
### GetTrackingOptions

`func (o *FlowActionResponseObjectResourceAttributes) GetTrackingOptions() FlowActionResponseObjectResourceAttributesTrackingOptions`

GetTrackingOptions returns the TrackingOptions field if non-nil, zero value otherwise.

### GetTrackingOptionsOk

`func (o *FlowActionResponseObjectResourceAttributes) GetTrackingOptionsOk() (*FlowActionResponseObjectResourceAttributesTrackingOptions, bool)`

GetTrackingOptionsOk returns a tuple with the TrackingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingOptions

`func (o *FlowActionResponseObjectResourceAttributes) SetTrackingOptions(v FlowActionResponseObjectResourceAttributesTrackingOptions)`

SetTrackingOptions sets TrackingOptions field to given value.

### HasTrackingOptions

`func (o *FlowActionResponseObjectResourceAttributes) HasTrackingOptions() bool`

HasTrackingOptions returns a boolean if a field has been set.

### SetTrackingOptionsNil

`func (o *FlowActionResponseObjectResourceAttributes) SetTrackingOptionsNil(b bool)`

 SetTrackingOptionsNil sets the value for TrackingOptions to be an explicit nil

### UnsetTrackingOptions
`func (o *FlowActionResponseObjectResourceAttributes) UnsetTrackingOptions()`

UnsetTrackingOptions ensures that no value is present for TrackingOptions, not even an explicit nil
### GetSendOptions

`func (o *FlowActionResponseObjectResourceAttributes) GetSendOptions() SendOptions`

GetSendOptions returns the SendOptions field if non-nil, zero value otherwise.

### GetSendOptionsOk

`func (o *FlowActionResponseObjectResourceAttributes) GetSendOptionsOk() (*SendOptions, bool)`

GetSendOptionsOk returns a tuple with the SendOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendOptions

`func (o *FlowActionResponseObjectResourceAttributes) SetSendOptions(v SendOptions)`

SetSendOptions sets SendOptions field to given value.

### HasSendOptions

`func (o *FlowActionResponseObjectResourceAttributes) HasSendOptions() bool`

HasSendOptions returns a boolean if a field has been set.

### GetRenderOptions

`func (o *FlowActionResponseObjectResourceAttributes) GetRenderOptions() SMSRenderOptions`

GetRenderOptions returns the RenderOptions field if non-nil, zero value otherwise.

### GetRenderOptionsOk

`func (o *FlowActionResponseObjectResourceAttributes) GetRenderOptionsOk() (*SMSRenderOptions, bool)`

GetRenderOptionsOk returns a tuple with the RenderOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderOptions

`func (o *FlowActionResponseObjectResourceAttributes) SetRenderOptions(v SMSRenderOptions)`

SetRenderOptions sets RenderOptions field to given value.

### HasRenderOptions

`func (o *FlowActionResponseObjectResourceAttributes) HasRenderOptions() bool`

HasRenderOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


