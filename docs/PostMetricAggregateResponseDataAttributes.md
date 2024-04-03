# PostMetricAggregateResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dates** | [**[]time.Time**](time.Time.md) | The dates of the query range | 
**Data** | [**[]MetricAggregateRowDTO**](MetricAggregateRowDTO.md) | Aggregation result data | 

## Methods

### NewPostMetricAggregateResponseDataAttributes

`func NewPostMetricAggregateResponseDataAttributes(dates []time.Time, data []MetricAggregateRowDTO, ) *PostMetricAggregateResponseDataAttributes`

NewPostMetricAggregateResponseDataAttributes instantiates a new PostMetricAggregateResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostMetricAggregateResponseDataAttributesWithDefaults

`func NewPostMetricAggregateResponseDataAttributesWithDefaults() *PostMetricAggregateResponseDataAttributes`

NewPostMetricAggregateResponseDataAttributesWithDefaults instantiates a new PostMetricAggregateResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDates

`func (o *PostMetricAggregateResponseDataAttributes) GetDates() []time.Time`

GetDates returns the Dates field if non-nil, zero value otherwise.

### GetDatesOk

`func (o *PostMetricAggregateResponseDataAttributes) GetDatesOk() (*[]time.Time, bool)`

GetDatesOk returns a tuple with the Dates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDates

`func (o *PostMetricAggregateResponseDataAttributes) SetDates(v []time.Time)`

SetDates sets Dates field to given value.


### GetData

`func (o *PostMetricAggregateResponseDataAttributes) GetData() []MetricAggregateRowDTO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PostMetricAggregateResponseDataAttributes) GetDataOk() (*[]MetricAggregateRowDTO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PostMetricAggregateResponseDataAttributes) SetData(v []MetricAggregateRowDTO)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


