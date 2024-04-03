# MetricAggregateRowDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimensions** | **[]string** | List of dimensions associated with this set of measurements | 
**Measurements** | **map[string]interface{}** | Dictionary of measurement_key, values | 

## Methods

### NewMetricAggregateRowDTO

`func NewMetricAggregateRowDTO(dimensions []string, measurements map[string]interface{}, ) *MetricAggregateRowDTO`

NewMetricAggregateRowDTO instantiates a new MetricAggregateRowDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricAggregateRowDTOWithDefaults

`func NewMetricAggregateRowDTOWithDefaults() *MetricAggregateRowDTO`

NewMetricAggregateRowDTOWithDefaults instantiates a new MetricAggregateRowDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimensions

`func (o *MetricAggregateRowDTO) GetDimensions() []string`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *MetricAggregateRowDTO) GetDimensionsOk() (*[]string, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *MetricAggregateRowDTO) SetDimensions(v []string)`

SetDimensions sets Dimensions field to given value.


### GetMeasurements

`func (o *MetricAggregateRowDTO) GetMeasurements() map[string]interface{}`

GetMeasurements returns the Measurements field if non-nil, zero value otherwise.

### GetMeasurementsOk

`func (o *MetricAggregateRowDTO) GetMeasurementsOk() (*map[string]interface{}, bool)`

GetMeasurementsOk returns a tuple with the Measurements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurements

`func (o *MetricAggregateRowDTO) SetMeasurements(v map[string]interface{})`

SetMeasurements sets Measurements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


