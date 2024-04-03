# PostMetricAggregateResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MetricAggregateEnum**](MetricAggregateEnum.md) |  | 
**Id** | **string** | Ephemeral ID associated with the aggregation query | 
**Attributes** | [**PostMetricAggregateResponseDataAttributes**](PostMetricAggregateResponseDataAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewPostMetricAggregateResponseData

`func NewPostMetricAggregateResponseData(type_ MetricAggregateEnum, id string, attributes PostMetricAggregateResponseDataAttributes, links ObjectLinks, ) *PostMetricAggregateResponseData`

NewPostMetricAggregateResponseData instantiates a new PostMetricAggregateResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostMetricAggregateResponseDataWithDefaults

`func NewPostMetricAggregateResponseDataWithDefaults() *PostMetricAggregateResponseData`

NewPostMetricAggregateResponseDataWithDefaults instantiates a new PostMetricAggregateResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PostMetricAggregateResponseData) GetType() MetricAggregateEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostMetricAggregateResponseData) GetTypeOk() (*MetricAggregateEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostMetricAggregateResponseData) SetType(v MetricAggregateEnum)`

SetType sets Type field to given value.


### GetId

`func (o *PostMetricAggregateResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostMetricAggregateResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostMetricAggregateResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *PostMetricAggregateResponseData) GetAttributes() PostMetricAggregateResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PostMetricAggregateResponseData) GetAttributesOk() (*PostMetricAggregateResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PostMetricAggregateResponseData) SetAttributes(v PostMetricAggregateResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *PostMetricAggregateResponseData) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PostMetricAggregateResponseData) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PostMetricAggregateResponseData) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


