# MetricCreateQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MetricEnum**](MetricEnum.md) |  | 
**Attributes** | [**MetricCreateQueryResourceObjectAttributes**](MetricCreateQueryResourceObjectAttributes.md) |  | 

## Methods

### NewMetricCreateQueryResourceObject

`func NewMetricCreateQueryResourceObject(type_ MetricEnum, attributes MetricCreateQueryResourceObjectAttributes, ) *MetricCreateQueryResourceObject`

NewMetricCreateQueryResourceObject instantiates a new MetricCreateQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricCreateQueryResourceObjectWithDefaults

`func NewMetricCreateQueryResourceObjectWithDefaults() *MetricCreateQueryResourceObject`

NewMetricCreateQueryResourceObjectWithDefaults instantiates a new MetricCreateQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MetricCreateQueryResourceObject) GetType() MetricEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricCreateQueryResourceObject) GetTypeOk() (*MetricEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricCreateQueryResourceObject) SetType(v MetricEnum)`

SetType sets Type field to given value.


### GetAttributes

`func (o *MetricCreateQueryResourceObject) GetAttributes() MetricCreateQueryResourceObjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MetricCreateQueryResourceObject) GetAttributesOk() (*MetricCreateQueryResourceObjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MetricCreateQueryResourceObject) SetAttributes(v MetricCreateQueryResourceObjectAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


