# PostFlowSeriesResponseDTODataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]SeriesData**](SeriesData.md) | An array of all the returned values data. Each object in the array represents one unique grouping and the results for that grouping. Each value in the results array corresponds to the date time at the same index. | 
**DateTimes** | [**[]time.Time**](time.Time.md) | An array of date times which correspond to the equivalent index in the results data. | 

## Methods

### NewPostFlowSeriesResponseDTODataAttributes

`func NewPostFlowSeriesResponseDTODataAttributes(results []SeriesData, dateTimes []time.Time, ) *PostFlowSeriesResponseDTODataAttributes`

NewPostFlowSeriesResponseDTODataAttributes instantiates a new PostFlowSeriesResponseDTODataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostFlowSeriesResponseDTODataAttributesWithDefaults

`func NewPostFlowSeriesResponseDTODataAttributesWithDefaults() *PostFlowSeriesResponseDTODataAttributes`

NewPostFlowSeriesResponseDTODataAttributesWithDefaults instantiates a new PostFlowSeriesResponseDTODataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *PostFlowSeriesResponseDTODataAttributes) GetResults() []SeriesData`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PostFlowSeriesResponseDTODataAttributes) GetResultsOk() (*[]SeriesData, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PostFlowSeriesResponseDTODataAttributes) SetResults(v []SeriesData)`

SetResults sets Results field to given value.


### GetDateTimes

`func (o *PostFlowSeriesResponseDTODataAttributes) GetDateTimes() []time.Time`

GetDateTimes returns the DateTimes field if non-nil, zero value otherwise.

### GetDateTimesOk

`func (o *PostFlowSeriesResponseDTODataAttributes) GetDateTimesOk() (*[]time.Time, bool)`

GetDateTimesOk returns a tuple with the DateTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimes

`func (o *PostFlowSeriesResponseDTODataAttributes) SetDateTimes(v []time.Time)`

SetDateTimes sets DateTimes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


