# GetImageResponseCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]ImageResponseObjectResource**](ImageResponseObjectResource.md) |  | 
**Links** | [**CollectionLinks**](CollectionLinks.md) |  | 

## Methods

### NewGetImageResponseCollection

`func NewGetImageResponseCollection(data []ImageResponseObjectResource, links CollectionLinks, ) *GetImageResponseCollection`

NewGetImageResponseCollection instantiates a new GetImageResponseCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetImageResponseCollectionWithDefaults

`func NewGetImageResponseCollectionWithDefaults() *GetImageResponseCollection`

NewGetImageResponseCollectionWithDefaults instantiates a new GetImageResponseCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetImageResponseCollection) GetData() []ImageResponseObjectResource`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetImageResponseCollection) GetDataOk() (*[]ImageResponseObjectResource, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetImageResponseCollection) SetData(v []ImageResponseObjectResource)`

SetData sets Data field to given value.


### GetLinks

`func (o *GetImageResponseCollection) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetImageResponseCollection) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetImageResponseCollection) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


