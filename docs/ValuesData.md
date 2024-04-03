# ValuesData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groupings** | **map[string]interface{}** | Applied groupings and the values for this object | 
**Statistics** | **map[string]interface{}** | Requested statistics and their values results | 

## Methods

### NewValuesData

`func NewValuesData(groupings map[string]interface{}, statistics map[string]interface{}, ) *ValuesData`

NewValuesData instantiates a new ValuesData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValuesDataWithDefaults

`func NewValuesDataWithDefaults() *ValuesData`

NewValuesDataWithDefaults instantiates a new ValuesData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupings

`func (o *ValuesData) GetGroupings() map[string]interface{}`

GetGroupings returns the Groupings field if non-nil, zero value otherwise.

### GetGroupingsOk

`func (o *ValuesData) GetGroupingsOk() (*map[string]interface{}, bool)`

GetGroupingsOk returns a tuple with the Groupings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupings

`func (o *ValuesData) SetGroupings(v map[string]interface{})`

SetGroupings sets Groupings field to given value.


### GetStatistics

`func (o *ValuesData) GetStatistics() map[string]interface{}`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *ValuesData) GetStatisticsOk() (*map[string]interface{}, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *ValuesData) SetStatistics(v map[string]interface{})`

SetStatistics sets Statistics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


