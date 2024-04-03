# GetSegmentResponseCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GetSegmentResponseCollectionDataInner**](GetSegmentResponseCollectionDataInner.md) |  | 
**Links** | [**CollectionLinks**](CollectionLinks.md) |  | 

## Methods

### NewGetSegmentResponseCollection

`func NewGetSegmentResponseCollection(data []GetSegmentResponseCollectionDataInner, links CollectionLinks, ) *GetSegmentResponseCollection`

NewGetSegmentResponseCollection instantiates a new GetSegmentResponseCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSegmentResponseCollectionWithDefaults

`func NewGetSegmentResponseCollectionWithDefaults() *GetSegmentResponseCollection`

NewGetSegmentResponseCollectionWithDefaults instantiates a new GetSegmentResponseCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetSegmentResponseCollection) GetData() []GetSegmentResponseCollectionDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetSegmentResponseCollection) GetDataOk() (*[]GetSegmentResponseCollectionDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetSegmentResponseCollection) SetData(v []GetSegmentResponseCollectionDataInner)`

SetData sets Data field to given value.


### GetLinks

`func (o *GetSegmentResponseCollection) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetSegmentResponseCollection) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetSegmentResponseCollection) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


