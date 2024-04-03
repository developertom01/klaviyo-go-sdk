# MetricCreateQueryResourceObjectAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the event. Must be less than 128 characters. | 
**Service** | Pointer to **NullableString** | This is for advanced usage. For api requests, this should use the default, which is set to api. | [optional] 

## Methods

### NewMetricCreateQueryResourceObjectAttributes

`func NewMetricCreateQueryResourceObjectAttributes(name string, ) *MetricCreateQueryResourceObjectAttributes`

NewMetricCreateQueryResourceObjectAttributes instantiates a new MetricCreateQueryResourceObjectAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricCreateQueryResourceObjectAttributesWithDefaults

`func NewMetricCreateQueryResourceObjectAttributesWithDefaults() *MetricCreateQueryResourceObjectAttributes`

NewMetricCreateQueryResourceObjectAttributesWithDefaults instantiates a new MetricCreateQueryResourceObjectAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MetricCreateQueryResourceObjectAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricCreateQueryResourceObjectAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricCreateQueryResourceObjectAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetService

`func (o *MetricCreateQueryResourceObjectAttributes) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *MetricCreateQueryResourceObjectAttributes) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *MetricCreateQueryResourceObjectAttributes) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *MetricCreateQueryResourceObjectAttributes) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *MetricCreateQueryResourceObjectAttributes) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *MetricCreateQueryResourceObjectAttributes) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


