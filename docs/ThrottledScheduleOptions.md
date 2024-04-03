# ThrottledScheduleOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **time.Time** | The time to send at | 
**ThrottlePercentage** | **int32** | The percentage of recipients per hour to send to. Allowed values: [10, 11, 13, 14, 17, 20, 25, 33, 50] | 

## Methods

### NewThrottledScheduleOptions

`func NewThrottledScheduleOptions(datetime time.Time, throttlePercentage int32, ) *ThrottledScheduleOptions`

NewThrottledScheduleOptions instantiates a new ThrottledScheduleOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThrottledScheduleOptionsWithDefaults

`func NewThrottledScheduleOptionsWithDefaults() *ThrottledScheduleOptions`

NewThrottledScheduleOptionsWithDefaults instantiates a new ThrottledScheduleOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *ThrottledScheduleOptions) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *ThrottledScheduleOptions) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *ThrottledScheduleOptions) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.


### GetThrottlePercentage

`func (o *ThrottledScheduleOptions) GetThrottlePercentage() int32`

GetThrottlePercentage returns the ThrottlePercentage field if non-nil, zero value otherwise.

### GetThrottlePercentageOk

`func (o *ThrottledScheduleOptions) GetThrottlePercentageOk() (*int32, bool)`

GetThrottlePercentageOk returns a tuple with the ThrottlePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlePercentage

`func (o *ThrottledScheduleOptions) SetThrottlePercentage(v int32)`

SetThrottlePercentage sets ThrottlePercentage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


