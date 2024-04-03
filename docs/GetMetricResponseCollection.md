# GetMetricResponseCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]MetricResponseObjectResource**](MetricResponseObjectResource.md) |  | 
**Links** | [**CollectionLinks**](CollectionLinks.md) |  | 

## Methods

### NewGetMetricResponseCollection

`func NewGetMetricResponseCollection(data []MetricResponseObjectResource, links CollectionLinks, ) *GetMetricResponseCollection`

NewGetMetricResponseCollection instantiates a new GetMetricResponseCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMetricResponseCollectionWithDefaults

`func NewGetMetricResponseCollectionWithDefaults() *GetMetricResponseCollection`

NewGetMetricResponseCollectionWithDefaults instantiates a new GetMetricResponseCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetMetricResponseCollection) GetData() []MetricResponseObjectResource`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetMetricResponseCollection) GetDataOk() (*[]MetricResponseObjectResource, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetMetricResponseCollection) SetData(v []MetricResponseObjectResource)`

SetData sets Data field to given value.


### GetLinks

`func (o *GetMetricResponseCollection) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetMetricResponseCollection) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetMetricResponseCollection) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


