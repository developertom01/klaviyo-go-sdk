# SendStrategySubObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | Describes the shape of the options object. Allowed values: [&#39;static&#39;, &#39;throttled&#39;, &#39;immediate&#39;, &#39;smart_send_time&#39;] | 
**OptionsStatic** | Pointer to [**StaticScheduleOptions**](StaticScheduleOptions.md) |  | [optional] 
**OptionsThrottled** | Pointer to [**ThrottledScheduleOptions**](ThrottledScheduleOptions.md) |  | [optional] 
**OptionsSto** | Pointer to [**STOScheduleOptions**](STOScheduleOptions.md) |  | [optional] 

## Methods

### NewSendStrategySubObject

`func NewSendStrategySubObject(method string, ) *SendStrategySubObject`

NewSendStrategySubObject instantiates a new SendStrategySubObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendStrategySubObjectWithDefaults

`func NewSendStrategySubObjectWithDefaults() *SendStrategySubObject`

NewSendStrategySubObjectWithDefaults instantiates a new SendStrategySubObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *SendStrategySubObject) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SendStrategySubObject) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *SendStrategySubObject) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetOptionsStatic

`func (o *SendStrategySubObject) GetOptionsStatic() StaticScheduleOptions`

GetOptionsStatic returns the OptionsStatic field if non-nil, zero value otherwise.

### GetOptionsStaticOk

`func (o *SendStrategySubObject) GetOptionsStaticOk() (*StaticScheduleOptions, bool)`

GetOptionsStaticOk returns a tuple with the OptionsStatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsStatic

`func (o *SendStrategySubObject) SetOptionsStatic(v StaticScheduleOptions)`

SetOptionsStatic sets OptionsStatic field to given value.

### HasOptionsStatic

`func (o *SendStrategySubObject) HasOptionsStatic() bool`

HasOptionsStatic returns a boolean if a field has been set.

### GetOptionsThrottled

`func (o *SendStrategySubObject) GetOptionsThrottled() ThrottledScheduleOptions`

GetOptionsThrottled returns the OptionsThrottled field if non-nil, zero value otherwise.

### GetOptionsThrottledOk

`func (o *SendStrategySubObject) GetOptionsThrottledOk() (*ThrottledScheduleOptions, bool)`

GetOptionsThrottledOk returns a tuple with the OptionsThrottled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsThrottled

`func (o *SendStrategySubObject) SetOptionsThrottled(v ThrottledScheduleOptions)`

SetOptionsThrottled sets OptionsThrottled field to given value.

### HasOptionsThrottled

`func (o *SendStrategySubObject) HasOptionsThrottled() bool`

HasOptionsThrottled returns a boolean if a field has been set.

### GetOptionsSto

`func (o *SendStrategySubObject) GetOptionsSto() STOScheduleOptions`

GetOptionsSto returns the OptionsSto field if non-nil, zero value otherwise.

### GetOptionsStoOk

`func (o *SendStrategySubObject) GetOptionsStoOk() (*STOScheduleOptions, bool)`

GetOptionsStoOk returns a tuple with the OptionsSto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsSto

`func (o *SendStrategySubObject) SetOptionsSto(v STOScheduleOptions)`

SetOptionsSto sets OptionsSto field to given value.

### HasOptionsSto

`func (o *SendStrategySubObject) HasOptionsSto() bool`

HasOptionsSto returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


