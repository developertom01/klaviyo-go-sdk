# CustomTimeframe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | **time.Time** |  | 
**End** | **time.Time** |  | 

## Methods

### NewCustomTimeframe

`func NewCustomTimeframe(start time.Time, end time.Time, ) *CustomTimeframe`

NewCustomTimeframe instantiates a new CustomTimeframe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomTimeframeWithDefaults

`func NewCustomTimeframeWithDefaults() *CustomTimeframe`

NewCustomTimeframeWithDefaults instantiates a new CustomTimeframe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *CustomTimeframe) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CustomTimeframe) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CustomTimeframe) SetStart(v time.Time)`

SetStart sets Start field to given value.


### GetEnd

`func (o *CustomTimeframe) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *CustomTimeframe) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *CustomTimeframe) SetEnd(v time.Time)`

SetEnd sets End field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


