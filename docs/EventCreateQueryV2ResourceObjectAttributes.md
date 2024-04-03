# EventCreateQueryV2ResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | **map[string]interface{}** | Properties of this event. Any top level property (that are not objects) can be used to create segments. The $extra property is a special property. This records any non-segmentable values that can be referenced later. For example, HTML templates are useful on a segment but are not used to create a segment. There are limits placed onto the size of the data present. This must not exceed 5 MB. This must not exceed 300 event properties. A single string cannot be larger than 100 KB. Each array must not exceed 4000 elements. The properties cannot contain more than 10 nested levels. | 
**Time** | Pointer to **NullableTime** | When this event occurred. By default, the time the request was received will be used. The time is truncated to the second. The time must be after the year 2000 and can only be up to 1 year in the future. | [optional] 
**Value** | Pointer to **NullableFloat32** | A numeric, monetary value to associate with this event. For example, the dollar amount of a purchase. | [optional] 
**ValueCurrency** | Pointer to **NullableString** | The ISO 4217 currency code of the value associated with the event. | [optional] 
**UniqueId** | Pointer to **NullableString** | A unique identifier for an event. If the unique_id is repeated for the same profile and metric, only the first processed event will be recorded. If this is not present, this will use the time to the second. Using the default, this limits only one event per profile per second. | [optional] 
**Metric** | [**EventCreateQueryV2ResourceObjectAttributesMetric**](EventCreateQueryV2ResourceObjectAttributesMetric.md) |  | 
**Profile** | [**EventCreateQueryV2ResourceObjectAttributesProfile**](EventCreateQueryV2ResourceObjectAttributesProfile.md) |  | 

## Methods

### NewEventCreateQueryV2ResourceObjectAttributes

`func NewEventCreateQueryV2ResourceObjectAttributes(properties map[string]interface{}, metric EventCreateQueryV2ResourceObjectAttributesMetric, profile EventCreateQueryV2ResourceObjectAttributesProfile, ) *EventCreateQueryV2ResourceObjectAttributes`

NewEventCreateQueryV2ResourceObjectAttributes instantiates a new EventCreateQueryV2ResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventCreateQueryV2ResourceObjectAttributesWithDefaults

`func NewEventCreateQueryV2ResourceObjectAttributesWithDefaults() *EventCreateQueryV2ResourceObjectAttributes`

NewEventCreateQueryV2ResourceObjectAttributesWithDefaults instantiates a new EventCreateQueryV2ResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *EventCreateQueryV2ResourceObjectAttributes) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *EventCreateQueryV2ResourceObjectAttributes) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *EventCreateQueryV2ResourceObjectAttributes) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.


### GetTime

`func (o *EventCreateQueryV2ResourceObjectAttributes) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *EventCreateQueryV2ResourceObjectAttributes) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *EventCreateQueryV2ResourceObjectAttributes) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *EventCreateQueryV2ResourceObjectAttributes) HasTime() bool`

HasTime returns a boolean if a field has been set.

### SetTimeNil

`func (o *EventCreateQueryV2ResourceObjectAttributes) SetTimeNil(b bool)`

 SetTimeNil sets the value for Time to be an explicit nil

### UnsetTime
`func (o *EventCreateQueryV2ResourceObjectAttributes) UnsetTime()`

UnsetTime ensures that no value is present for Time, not even an explicit nil
### GetValue

`func (o *EventCreateQueryV2ResourceObjectAttributes) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EventCreateQueryV2ResourceObjectAttributes) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EventCreateQueryV2ResourceObjectAttributes) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *EventCreateQueryV2ResourceObjectAttributes) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *EventCreateQueryV2ResourceObjectAttributes) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *EventCreateQueryV2ResourceObjectAttributes) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetValueCurrency

`func (o *EventCreateQueryV2ResourceObjectAttributes) GetValueCurrency() string`

GetValueCurrency returns the ValueCurrency field if non-nil, zero value otherwise.

### GetValueCurrencyOk

`func (o *EventCreateQueryV2ResourceObjectAttributes) GetValueCurrencyOk() (*string, bool)`

GetValueCurrencyOk returns a tuple with the ValueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueCurrency

`func (o *EventCreateQueryV2ResourceObjectAttributes) SetValueCurrency(v string)`

SetValueCurrency sets ValueCurrency field to given value.

### HasValueCurrency

`func (o *EventCreateQueryV2ResourceObjectAttributes) HasValueCurrency() bool`

HasValueCurrency returns a boolean if a field has been set.

### SetValueCurrencyNil

`func (o *EventCreateQueryV2ResourceObjectAttributes) SetValueCurrencyNil(b bool)`

 SetValueCurrencyNil sets the value for ValueCurrency to be an explicit nil

### UnsetValueCurrency
`func (o *EventCreateQueryV2ResourceObjectAttributes) UnsetValueCurrency()`

UnsetValueCurrency ensures that no value is present for ValueCurrency, not even an explicit nil
### GetUniqueId

`func (o *EventCreateQueryV2ResourceObjectAttributes) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *EventCreateQueryV2ResourceObjectAttributes) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *EventCreateQueryV2ResourceObjectAttributes) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *EventCreateQueryV2ResourceObjectAttributes) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### SetUniqueIdNil

`func (o *EventCreateQueryV2ResourceObjectAttributes) SetUniqueIdNil(b bool)`

 SetUniqueIdNil sets the value for UniqueId to be an explicit nil

### UnsetUniqueId
`func (o *EventCreateQueryV2ResourceObjectAttributes) UnsetUniqueId()`

UnsetUniqueId ensures that no value is present for UniqueId, not even an explicit nil
### GetMetric

`func (o *EventCreateQueryV2ResourceObjectAttributes) GetMetric() EventCreateQueryV2ResourceObjectAttributesMetric`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *EventCreateQueryV2ResourceObjectAttributes) GetMetricOk() (*EventCreateQueryV2ResourceObjectAttributesMetric, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *EventCreateQueryV2ResourceObjectAttributes) SetMetric(v EventCreateQueryV2ResourceObjectAttributesMetric)`

SetMetric sets Metric field to given value.


### GetProfile

`func (o *EventCreateQueryV2ResourceObjectAttributes) GetProfile() EventCreateQueryV2ResourceObjectAttributesProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *EventCreateQueryV2ResourceObjectAttributes) GetProfileOk() (*EventCreateQueryV2ResourceObjectAttributesProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *EventCreateQueryV2ResourceObjectAttributes) SetProfile(v EventCreateQueryV2ResourceObjectAttributesProfile)`

SetProfile sets Profile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


