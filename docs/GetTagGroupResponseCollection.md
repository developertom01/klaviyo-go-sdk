# GetTagGroupResponseCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GetTagGroupResponseCollectionDataInner**](GetTagGroupResponseCollectionDataInner.md) |  | 
**Links** | [**CollectionLinks**](CollectionLinks.md) |  | 

## Methods

### NewGetTagGroupResponseCollection

`func NewGetTagGroupResponseCollection(data []GetTagGroupResponseCollectionDataInner, links CollectionLinks, ) *GetTagGroupResponseCollection`

NewGetTagGroupResponseCollection instantiates a new GetTagGroupResponseCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTagGroupResponseCollectionWithDefaults

`func NewGetTagGroupResponseCollectionWithDefaults() *GetTagGroupResponseCollection`

NewGetTagGroupResponseCollectionWithDefaults instantiates a new GetTagGroupResponseCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetTagGroupResponseCollection) GetData() []GetTagGroupResponseCollectionDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTagGroupResponseCollection) GetDataOk() (*[]GetTagGroupResponseCollectionDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTagGroupResponseCollection) SetData(v []GetTagGroupResponseCollectionDataInner)`

SetData sets Data field to given value.


### GetLinks

`func (o *GetTagGroupResponseCollection) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetTagGroupResponseCollection) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetTagGroupResponseCollection) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


