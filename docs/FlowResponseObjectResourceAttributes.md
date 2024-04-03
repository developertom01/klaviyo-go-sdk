# FlowResponseObjectResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Archived** | Pointer to **NullableBool** |  | [optional] 
**Created** | Pointer to **NullableTime** |  | [optional] 
**Updated** | Pointer to **NullableTime** |  | [optional] 
**TriggerType** | Pointer to **NullableString** | Corresponds to the object which triggered the flow. | [optional] 

## Methods

### NewFlowResponseObjectResourceAttributes

`func NewFlowResponseObjectResourceAttributes() *FlowResponseObjectResourceAttributes`

NewFlowResponseObjectResourceAttributes instantiates a new FlowResponseObjectResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowResponseObjectResourceAttributesWithDefaults

`func NewFlowResponseObjectResourceAttributesWithDefaults() *FlowResponseObjectResourceAttributes`

NewFlowResponseObjectResourceAttributesWithDefaults instantiates a new FlowResponseObjectResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FlowResponseObjectResourceAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlowResponseObjectResourceAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlowResponseObjectResourceAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlowResponseObjectResourceAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FlowResponseObjectResourceAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FlowResponseObjectResourceAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *FlowResponseObjectResourceAttributes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlowResponseObjectResourceAttributes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlowResponseObjectResourceAttributes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FlowResponseObjectResourceAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *FlowResponseObjectResourceAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *FlowResponseObjectResourceAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetArchived

`func (o *FlowResponseObjectResourceAttributes) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *FlowResponseObjectResourceAttributes) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *FlowResponseObjectResourceAttributes) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *FlowResponseObjectResourceAttributes) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### SetArchivedNil

`func (o *FlowResponseObjectResourceAttributes) SetArchivedNil(b bool)`

 SetArchivedNil sets the value for Archived to be an explicit nil

### UnsetArchived
`func (o *FlowResponseObjectResourceAttributes) UnsetArchived()`

UnsetArchived ensures that no value is present for Archived, not even an explicit nil
### GetCreated

`func (o *FlowResponseObjectResourceAttributes) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FlowResponseObjectResourceAttributes) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FlowResponseObjectResourceAttributes) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FlowResponseObjectResourceAttributes) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *FlowResponseObjectResourceAttributes) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *FlowResponseObjectResourceAttributes) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetUpdated

`func (o *FlowResponseObjectResourceAttributes) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *FlowResponseObjectResourceAttributes) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *FlowResponseObjectResourceAttributes) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *FlowResponseObjectResourceAttributes) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *FlowResponseObjectResourceAttributes) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *FlowResponseObjectResourceAttributes) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
### GetTriggerType

`func (o *FlowResponseObjectResourceAttributes) GetTriggerType() string`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *FlowResponseObjectResourceAttributes) GetTriggerTypeOk() (*string, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *FlowResponseObjectResourceAttributes) SetTriggerType(v string)`

SetTriggerType sets TriggerType field to given value.

### HasTriggerType

`func (o *FlowResponseObjectResourceAttributes) HasTriggerType() bool`

HasTriggerType returns a boolean if a field has been set.

### SetTriggerTypeNil

`func (o *FlowResponseObjectResourceAttributes) SetTriggerTypeNil(b bool)`

 SetTriggerTypeNil sets the value for TriggerType to be an explicit nil

### UnsetTriggerType
`func (o *FlowResponseObjectResourceAttributes) UnsetTriggerType()`

UnsetTriggerType ensures that no value is present for TriggerType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


