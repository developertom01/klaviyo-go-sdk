# StaticScheduleOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **time.Time** | The time to send at | 
**IsLocal** | Pointer to **NullableBool** | If the campaign should be sent with local recipient timezone send (requires UTC time) or statically sent at the given time. Defaults to False. | [optional] 
**SendPastRecipientsImmediately** | Pointer to **NullableBool** | Determines if we should send to local recipient timezone if the given time has passed. Only applicable to local sends. Defaults to False. | [optional] 

## Methods

### NewStaticScheduleOptions

`func NewStaticScheduleOptions(datetime time.Time, ) *StaticScheduleOptions`

NewStaticScheduleOptions instantiates a new StaticScheduleOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticScheduleOptionsWithDefaults

`func NewStaticScheduleOptionsWithDefaults() *StaticScheduleOptions`

NewStaticScheduleOptionsWithDefaults instantiates a new StaticScheduleOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *StaticScheduleOptions) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *StaticScheduleOptions) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *StaticScheduleOptions) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.


### GetIsLocal

`func (o *StaticScheduleOptions) GetIsLocal() bool`

GetIsLocal returns the IsLocal field if non-nil, zero value otherwise.

### GetIsLocalOk

`func (o *StaticScheduleOptions) GetIsLocalOk() (*bool, bool)`

GetIsLocalOk returns a tuple with the IsLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocal

`func (o *StaticScheduleOptions) SetIsLocal(v bool)`

SetIsLocal sets IsLocal field to given value.

### HasIsLocal

`func (o *StaticScheduleOptions) HasIsLocal() bool`

HasIsLocal returns a boolean if a field has been set.

### SetIsLocalNil

`func (o *StaticScheduleOptions) SetIsLocalNil(b bool)`

 SetIsLocalNil sets the value for IsLocal to be an explicit nil

### UnsetIsLocal
`func (o *StaticScheduleOptions) UnsetIsLocal()`

UnsetIsLocal ensures that no value is present for IsLocal, not even an explicit nil
### GetSendPastRecipientsImmediately

`func (o *StaticScheduleOptions) GetSendPastRecipientsImmediately() bool`

GetSendPastRecipientsImmediately returns the SendPastRecipientsImmediately field if non-nil, zero value otherwise.

### GetSendPastRecipientsImmediatelyOk

`func (o *StaticScheduleOptions) GetSendPastRecipientsImmediatelyOk() (*bool, bool)`

GetSendPastRecipientsImmediatelyOk returns a tuple with the SendPastRecipientsImmediately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendPastRecipientsImmediately

`func (o *StaticScheduleOptions) SetSendPastRecipientsImmediately(v bool)`

SetSendPastRecipientsImmediately sets SendPastRecipientsImmediately field to given value.

### HasSendPastRecipientsImmediately

`func (o *StaticScheduleOptions) HasSendPastRecipientsImmediately() bool`

HasSendPastRecipientsImmediately returns a boolean if a field has been set.

### SetSendPastRecipientsImmediatelyNil

`func (o *StaticScheduleOptions) SetSendPastRecipientsImmediatelyNil(b bool)`

 SetSendPastRecipientsImmediatelyNil sets the value for SendPastRecipientsImmediately to be an explicit nil

### UnsetSendPastRecipientsImmediately
`func (o *StaticScheduleOptions) UnsetSendPastRecipientsImmediately()`

UnsetSendPastRecipientsImmediately ensures that no value is present for SendPastRecipientsImmediately, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


