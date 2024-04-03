# MetricResponseObjectResourceAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The name of the metric | [optional] 
**Created** | Pointer to **NullableString** | Creation time in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm) | [optional] 
**Updated** | Pointer to **NullableString** | Last updated time in ISO 8601 format (YYYY-MM-DDTHH:MM:SS.mmmmmm) | [optional] 
**Integration** | Pointer to **map[string]interface{}** | The integration associated with the event | [optional] 

## Methods

### NewMetricResponseObjectResourceAttributes

`func NewMetricResponseObjectResourceAttributes() *MetricResponseObjectResourceAttributes`

NewMetricResponseObjectResourceAttributes instantiates a new MetricResponseObjectResourceAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricResponseObjectResourceAttributesWithDefaults

`func NewMetricResponseObjectResourceAttributesWithDefaults() *MetricResponseObjectResourceAttributes`

NewMetricResponseObjectResourceAttributesWithDefaults instantiates a new MetricResponseObjectResourceAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MetricResponseObjectResourceAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricResponseObjectResourceAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricResponseObjectResourceAttributes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetricResponseObjectResourceAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MetricResponseObjectResourceAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MetricResponseObjectResourceAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCreated

`func (o *MetricResponseObjectResourceAttributes) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *MetricResponseObjectResourceAttributes) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *MetricResponseObjectResourceAttributes) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *MetricResponseObjectResourceAttributes) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *MetricResponseObjectResourceAttributes) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *MetricResponseObjectResourceAttributes) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetUpdated

`func (o *MetricResponseObjectResourceAttributes) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *MetricResponseObjectResourceAttributes) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *MetricResponseObjectResourceAttributes) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *MetricResponseObjectResourceAttributes) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### SetUpdatedNil

`func (o *MetricResponseObjectResourceAttributes) SetUpdatedNil(b bool)`

 SetUpdatedNil sets the value for Updated to be an explicit nil

### UnsetUpdated
`func (o *MetricResponseObjectResourceAttributes) UnsetUpdated()`

UnsetUpdated ensures that no value is present for Updated, not even an explicit nil
### GetIntegration

`func (o *MetricResponseObjectResourceAttributes) GetIntegration() map[string]interface{}`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *MetricResponseObjectResourceAttributes) GetIntegrationOk() (*map[string]interface{}, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *MetricResponseObjectResourceAttributes) SetIntegration(v map[string]interface{})`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *MetricResponseObjectResourceAttributes) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### SetIntegrationNil

`func (o *MetricResponseObjectResourceAttributes) SetIntegrationNil(b bool)`

 SetIntegrationNil sets the value for Integration to be an explicit nil

### UnsetIntegration
`func (o *MetricResponseObjectResourceAttributes) UnsetIntegration()`

UnsetIntegration ensures that no value is present for Integration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


