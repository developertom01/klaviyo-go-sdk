# SendTimeSubObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **time.Time** | The datetime that the message is to be sent | 
**IsLocal** | **bool** | Whether that datetime is to be a local datetime for the recipient | 

## Methods

### NewSendTimeSubObject

`func NewSendTimeSubObject(datetime time.Time, isLocal bool, ) *SendTimeSubObject`

NewSendTimeSubObject instantiates a new SendTimeSubObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendTimeSubObjectWithDefaults

`func NewSendTimeSubObjectWithDefaults() *SendTimeSubObject`

NewSendTimeSubObjectWithDefaults instantiates a new SendTimeSubObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *SendTimeSubObject) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *SendTimeSubObject) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *SendTimeSubObject) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.


### GetIsLocal

`func (o *SendTimeSubObject) GetIsLocal() bool`

GetIsLocal returns the IsLocal field if non-nil, zero value otherwise.

### GetIsLocalOk

`func (o *SendTimeSubObject) GetIsLocalOk() (*bool, bool)`

GetIsLocalOk returns a tuple with the IsLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocal

`func (o *SendTimeSubObject) SetIsLocal(v bool)`

SetIsLocal sets IsLocal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


