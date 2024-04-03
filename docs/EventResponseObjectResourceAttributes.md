# EventResponseObjectResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **NullableInt32** | Event timestamp in seconds | [optional] 
**EventProperties** | Pointer to **map[string]interface{}** | Event properties, can include identifiers and extra properties | [optional] 
**Datetime** | Pointer to **NullableTime** | Event timestamp in ISO8601 format (YYYY-MM-DDTHH:MM:SS+hh:mm) | [optional] 
**Uuid** | Pointer to **NullableString** | A unique identifier for the event, this can be used as a cursor in pagination | [optional] 

## Methods

### NewEventResponseObjectResourceAttributes

`func NewEventResponseObjectResourceAttributes() *EventResponseObjectResourceAttributes`

NewEventResponseObjectResourceAttributes instantiates a new EventResponseObjectResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventResponseObjectResourceAttributesWithDefaults

`func NewEventResponseObjectResourceAttributesWithDefaults() *EventResponseObjectResourceAttributes`

NewEventResponseObjectResourceAttributesWithDefaults instantiates a new EventResponseObjectResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *EventResponseObjectResourceAttributes) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EventResponseObjectResourceAttributes) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EventResponseObjectResourceAttributes) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *EventResponseObjectResourceAttributes) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *EventResponseObjectResourceAttributes) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *EventResponseObjectResourceAttributes) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetEventProperties

`func (o *EventResponseObjectResourceAttributes) GetEventProperties() map[string]interface{}`

GetEventProperties returns the EventProperties field if non-nil, zero value otherwise.

### GetEventPropertiesOk

`func (o *EventResponseObjectResourceAttributes) GetEventPropertiesOk() (*map[string]interface{}, bool)`

GetEventPropertiesOk returns a tuple with the EventProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventProperties

`func (o *EventResponseObjectResourceAttributes) SetEventProperties(v map[string]interface{})`

SetEventProperties sets EventProperties field to given value.

### HasEventProperties

`func (o *EventResponseObjectResourceAttributes) HasEventProperties() bool`

HasEventProperties returns a boolean if a field has been set.

### SetEventPropertiesNil

`func (o *EventResponseObjectResourceAttributes) SetEventPropertiesNil(b bool)`

 SetEventPropertiesNil sets the value for EventProperties to be an explicit nil

### UnsetEventProperties
`func (o *EventResponseObjectResourceAttributes) UnsetEventProperties()`

UnsetEventProperties ensures that no value is present for EventProperties, not even an explicit nil
### GetDatetime

`func (o *EventResponseObjectResourceAttributes) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *EventResponseObjectResourceAttributes) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *EventResponseObjectResourceAttributes) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *EventResponseObjectResourceAttributes) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### SetDatetimeNil

`func (o *EventResponseObjectResourceAttributes) SetDatetimeNil(b bool)`

 SetDatetimeNil sets the value for Datetime to be an explicit nil

### UnsetDatetime
`func (o *EventResponseObjectResourceAttributes) UnsetDatetime()`

UnsetDatetime ensures that no value is present for Datetime, not even an explicit nil
### GetUuid

`func (o *EventResponseObjectResourceAttributes) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *EventResponseObjectResourceAttributes) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *EventResponseObjectResourceAttributes) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *EventResponseObjectResourceAttributes) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *EventResponseObjectResourceAttributes) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *EventResponseObjectResourceAttributes) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


