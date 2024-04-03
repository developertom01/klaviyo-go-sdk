# PredictiveAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HistoricClv** | Pointer to **NullableFloat32** | Total value of all historically placed orders | [optional] 
**PredictedClv** | Pointer to **NullableFloat32** | Predicted value of all placed orders in the next 365 days | [optional] 
**TotalClv** | Pointer to **NullableFloat32** | Sum of historic and predicted CLV | [optional] 
**HistoricNumberOfOrders** | Pointer to **NullableInt32** | Number of already placed orders | [optional] 
**PredictedNumberOfOrders** | Pointer to **NullableFloat32** | Predicted number of placed orders in the next 365 days | [optional] 
**AverageDaysBetweenOrders** | Pointer to **NullableFloat32** | Average number of days between orders (None if only one order has been placed) | [optional] 
**AverageOrderValue** | Pointer to **NullableFloat32** | Average value of placed orders | [optional] 
**ChurnProbability** | Pointer to **NullableFloat32** | Probability the customer has churned | [optional] 
**ExpectedDateOfNextOrder** | Pointer to **NullableTime** | Expected date of next order, as calculated at the time of their most recent order | [optional] 

## Methods

### NewPredictiveAnalytics

`func NewPredictiveAnalytics() *PredictiveAnalytics`

NewPredictiveAnalytics instantiates a new PredictiveAnalytics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPredictiveAnalyticsWithDefaults

`func NewPredictiveAnalyticsWithDefaults() *PredictiveAnalytics`

NewPredictiveAnalyticsWithDefaults instantiates a new PredictiveAnalytics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistoricClv

`func (o *PredictiveAnalytics) GetHistoricClv() float32`

GetHistoricClv returns the HistoricClv field if non-nil, zero value otherwise.

### GetHistoricClvOk

`func (o *PredictiveAnalytics) GetHistoricClvOk() (*float32, bool)`

GetHistoricClvOk returns a tuple with the HistoricClv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricClv

`func (o *PredictiveAnalytics) SetHistoricClv(v float32)`

SetHistoricClv sets HistoricClv field to given value.

### HasHistoricClv

`func (o *PredictiveAnalytics) HasHistoricClv() bool`

HasHistoricClv returns a boolean if a field has been set.

### SetHistoricClvNil

`func (o *PredictiveAnalytics) SetHistoricClvNil(b bool)`

 SetHistoricClvNil sets the value for HistoricClv to be an explicit nil

### UnsetHistoricClv
`func (o *PredictiveAnalytics) UnsetHistoricClv()`

UnsetHistoricClv ensures that no value is present for HistoricClv, not even an explicit nil
### GetPredictedClv

`func (o *PredictiveAnalytics) GetPredictedClv() float32`

GetPredictedClv returns the PredictedClv field if non-nil, zero value otherwise.

### GetPredictedClvOk

`func (o *PredictiveAnalytics) GetPredictedClvOk() (*float32, bool)`

GetPredictedClvOk returns a tuple with the PredictedClv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedClv

`func (o *PredictiveAnalytics) SetPredictedClv(v float32)`

SetPredictedClv sets PredictedClv field to given value.

### HasPredictedClv

`func (o *PredictiveAnalytics) HasPredictedClv() bool`

HasPredictedClv returns a boolean if a field has been set.

### SetPredictedClvNil

`func (o *PredictiveAnalytics) SetPredictedClvNil(b bool)`

 SetPredictedClvNil sets the value for PredictedClv to be an explicit nil

### UnsetPredictedClv
`func (o *PredictiveAnalytics) UnsetPredictedClv()`

UnsetPredictedClv ensures that no value is present for PredictedClv, not even an explicit nil
### GetTotalClv

`func (o *PredictiveAnalytics) GetTotalClv() float32`

GetTotalClv returns the TotalClv field if non-nil, zero value otherwise.

### GetTotalClvOk

`func (o *PredictiveAnalytics) GetTotalClvOk() (*float32, bool)`

GetTotalClvOk returns a tuple with the TotalClv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalClv

`func (o *PredictiveAnalytics) SetTotalClv(v float32)`

SetTotalClv sets TotalClv field to given value.

### HasTotalClv

`func (o *PredictiveAnalytics) HasTotalClv() bool`

HasTotalClv returns a boolean if a field has been set.

### SetTotalClvNil

`func (o *PredictiveAnalytics) SetTotalClvNil(b bool)`

 SetTotalClvNil sets the value for TotalClv to be an explicit nil

### UnsetTotalClv
`func (o *PredictiveAnalytics) UnsetTotalClv()`

UnsetTotalClv ensures that no value is present for TotalClv, not even an explicit nil
### GetHistoricNumberOfOrders

`func (o *PredictiveAnalytics) GetHistoricNumberOfOrders() int32`

GetHistoricNumberOfOrders returns the HistoricNumberOfOrders field if non-nil, zero value otherwise.

### GetHistoricNumberOfOrdersOk

`func (o *PredictiveAnalytics) GetHistoricNumberOfOrdersOk() (*int32, bool)`

GetHistoricNumberOfOrdersOk returns a tuple with the HistoricNumberOfOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricNumberOfOrders

`func (o *PredictiveAnalytics) SetHistoricNumberOfOrders(v int32)`

SetHistoricNumberOfOrders sets HistoricNumberOfOrders field to given value.

### HasHistoricNumberOfOrders

`func (o *PredictiveAnalytics) HasHistoricNumberOfOrders() bool`

HasHistoricNumberOfOrders returns a boolean if a field has been set.

### SetHistoricNumberOfOrdersNil

`func (o *PredictiveAnalytics) SetHistoricNumberOfOrdersNil(b bool)`

 SetHistoricNumberOfOrdersNil sets the value for HistoricNumberOfOrders to be an explicit nil

### UnsetHistoricNumberOfOrders
`func (o *PredictiveAnalytics) UnsetHistoricNumberOfOrders()`

UnsetHistoricNumberOfOrders ensures that no value is present for HistoricNumberOfOrders, not even an explicit nil
### GetPredictedNumberOfOrders

`func (o *PredictiveAnalytics) GetPredictedNumberOfOrders() float32`

GetPredictedNumberOfOrders returns the PredictedNumberOfOrders field if non-nil, zero value otherwise.

### GetPredictedNumberOfOrdersOk

`func (o *PredictiveAnalytics) GetPredictedNumberOfOrdersOk() (*float32, bool)`

GetPredictedNumberOfOrdersOk returns a tuple with the PredictedNumberOfOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedNumberOfOrders

`func (o *PredictiveAnalytics) SetPredictedNumberOfOrders(v float32)`

SetPredictedNumberOfOrders sets PredictedNumberOfOrders field to given value.

### HasPredictedNumberOfOrders

`func (o *PredictiveAnalytics) HasPredictedNumberOfOrders() bool`

HasPredictedNumberOfOrders returns a boolean if a field has been set.

### SetPredictedNumberOfOrdersNil

`func (o *PredictiveAnalytics) SetPredictedNumberOfOrdersNil(b bool)`

 SetPredictedNumberOfOrdersNil sets the value for PredictedNumberOfOrders to be an explicit nil

### UnsetPredictedNumberOfOrders
`func (o *PredictiveAnalytics) UnsetPredictedNumberOfOrders()`

UnsetPredictedNumberOfOrders ensures that no value is present for PredictedNumberOfOrders, not even an explicit nil
### GetAverageDaysBetweenOrders

`func (o *PredictiveAnalytics) GetAverageDaysBetweenOrders() float32`

GetAverageDaysBetweenOrders returns the AverageDaysBetweenOrders field if non-nil, zero value otherwise.

### GetAverageDaysBetweenOrdersOk

`func (o *PredictiveAnalytics) GetAverageDaysBetweenOrdersOk() (*float32, bool)`

GetAverageDaysBetweenOrdersOk returns a tuple with the AverageDaysBetweenOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageDaysBetweenOrders

`func (o *PredictiveAnalytics) SetAverageDaysBetweenOrders(v float32)`

SetAverageDaysBetweenOrders sets AverageDaysBetweenOrders field to given value.

### HasAverageDaysBetweenOrders

`func (o *PredictiveAnalytics) HasAverageDaysBetweenOrders() bool`

HasAverageDaysBetweenOrders returns a boolean if a field has been set.

### SetAverageDaysBetweenOrdersNil

`func (o *PredictiveAnalytics) SetAverageDaysBetweenOrdersNil(b bool)`

 SetAverageDaysBetweenOrdersNil sets the value for AverageDaysBetweenOrders to be an explicit nil

### UnsetAverageDaysBetweenOrders
`func (o *PredictiveAnalytics) UnsetAverageDaysBetweenOrders()`

UnsetAverageDaysBetweenOrders ensures that no value is present for AverageDaysBetweenOrders, not even an explicit nil
### GetAverageOrderValue

`func (o *PredictiveAnalytics) GetAverageOrderValue() float32`

GetAverageOrderValue returns the AverageOrderValue field if non-nil, zero value otherwise.

### GetAverageOrderValueOk

`func (o *PredictiveAnalytics) GetAverageOrderValueOk() (*float32, bool)`

GetAverageOrderValueOk returns a tuple with the AverageOrderValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageOrderValue

`func (o *PredictiveAnalytics) SetAverageOrderValue(v float32)`

SetAverageOrderValue sets AverageOrderValue field to given value.

### HasAverageOrderValue

`func (o *PredictiveAnalytics) HasAverageOrderValue() bool`

HasAverageOrderValue returns a boolean if a field has been set.

### SetAverageOrderValueNil

`func (o *PredictiveAnalytics) SetAverageOrderValueNil(b bool)`

 SetAverageOrderValueNil sets the value for AverageOrderValue to be an explicit nil

### UnsetAverageOrderValue
`func (o *PredictiveAnalytics) UnsetAverageOrderValue()`

UnsetAverageOrderValue ensures that no value is present for AverageOrderValue, not even an explicit nil
### GetChurnProbability

`func (o *PredictiveAnalytics) GetChurnProbability() float32`

GetChurnProbability returns the ChurnProbability field if non-nil, zero value otherwise.

### GetChurnProbabilityOk

`func (o *PredictiveAnalytics) GetChurnProbabilityOk() (*float32, bool)`

GetChurnProbabilityOk returns a tuple with the ChurnProbability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChurnProbability

`func (o *PredictiveAnalytics) SetChurnProbability(v float32)`

SetChurnProbability sets ChurnProbability field to given value.

### HasChurnProbability

`func (o *PredictiveAnalytics) HasChurnProbability() bool`

HasChurnProbability returns a boolean if a field has been set.

### SetChurnProbabilityNil

`func (o *PredictiveAnalytics) SetChurnProbabilityNil(b bool)`

 SetChurnProbabilityNil sets the value for ChurnProbability to be an explicit nil

### UnsetChurnProbability
`func (o *PredictiveAnalytics) UnsetChurnProbability()`

UnsetChurnProbability ensures that no value is present for ChurnProbability, not even an explicit nil
### GetExpectedDateOfNextOrder

`func (o *PredictiveAnalytics) GetExpectedDateOfNextOrder() time.Time`

GetExpectedDateOfNextOrder returns the ExpectedDateOfNextOrder field if non-nil, zero value otherwise.

### GetExpectedDateOfNextOrderOk

`func (o *PredictiveAnalytics) GetExpectedDateOfNextOrderOk() (*time.Time, bool)`

GetExpectedDateOfNextOrderOk returns a tuple with the ExpectedDateOfNextOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedDateOfNextOrder

`func (o *PredictiveAnalytics) SetExpectedDateOfNextOrder(v time.Time)`

SetExpectedDateOfNextOrder sets ExpectedDateOfNextOrder field to given value.

### HasExpectedDateOfNextOrder

`func (o *PredictiveAnalytics) HasExpectedDateOfNextOrder() bool`

HasExpectedDateOfNextOrder returns a boolean if a field has been set.

### SetExpectedDateOfNextOrderNil

`func (o *PredictiveAnalytics) SetExpectedDateOfNextOrderNil(b bool)`

 SetExpectedDateOfNextOrderNil sets the value for ExpectedDateOfNextOrder to be an explicit nil

### UnsetExpectedDateOfNextOrder
`func (o *PredictiveAnalytics) UnsetExpectedDateOfNextOrder()`

UnsetExpectedDateOfNextOrder ensures that no value is present for ExpectedDateOfNextOrder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


