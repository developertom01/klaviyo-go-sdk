# MetricResponseObjectResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MetricEnum**](MetricEnum.md) |  | 
**Id** | **string** | The Metric ID | 
**Attributes** | [**MetricResponseObjectResourceAttributes**](MetricResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewMetricResponseObjectResource

`func NewMetricResponseObjectResource(type_ MetricEnum, id string, attributes MetricResponseObjectResourceAttributes, links ObjectLinks, ) *MetricResponseObjectResource`

NewMetricResponseObjectResource instantiates a new MetricResponseObjectResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricResponseObjectResourceWithDefaults

`func NewMetricResponseObjectResourceWithDefaults() *MetricResponseObjectResource`

NewMetricResponseObjectResourceWithDefaults instantiates a new MetricResponseObjectResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MetricResponseObjectResource) GetType() MetricEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricResponseObjectResource) GetTypeOk() (*MetricEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricResponseObjectResource) SetType(v MetricEnum)`

SetType sets Type field to given value.


### GetId

`func (o *MetricResponseObjectResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricResponseObjectResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricResponseObjectResource) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *MetricResponseObjectResource) GetAttributes() MetricResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *MetricResponseObjectResource) GetAttributesOk() (*MetricResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *MetricResponseObjectResource) SetAttributes(v MetricResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *MetricResponseObjectResource) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MetricResponseObjectResource) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MetricResponseObjectResource) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


