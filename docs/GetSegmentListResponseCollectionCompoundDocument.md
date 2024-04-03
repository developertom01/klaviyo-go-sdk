# GetSegmentListResponseCollectionCompoundDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GetSegmentListResponseCollectionCompoundDocumentDataInner**](GetSegmentListResponseCollectionCompoundDocumentDataInner.md) |  | 
**Links** | [**CollectionLinks**](CollectionLinks.md) |  | 
**Included** | Pointer to [**[]TagResponseObjectResource**](TagResponseObjectResource.md) |  | [optional] 

## Methods

### NewGetSegmentListResponseCollectionCompoundDocument

`func NewGetSegmentListResponseCollectionCompoundDocument(data []GetSegmentListResponseCollectionCompoundDocumentDataInner, links CollectionLinks, ) *GetSegmentListResponseCollectionCompoundDocument`

NewGetSegmentListResponseCollectionCompoundDocument instantiates a new GetSegmentListResponseCollectionCompoundDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSegmentListResponseCollectionCompoundDocumentWithDefaults

`func NewGetSegmentListResponseCollectionCompoundDocumentWithDefaults() *GetSegmentListResponseCollectionCompoundDocument`

NewGetSegmentListResponseCollectionCompoundDocumentWithDefaults instantiates a new GetSegmentListResponseCollectionCompoundDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetSegmentListResponseCollectionCompoundDocument) GetData() []GetSegmentListResponseCollectionCompoundDocumentDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetSegmentListResponseCollectionCompoundDocument) GetDataOk() (*[]GetSegmentListResponseCollectionCompoundDocumentDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetSegmentListResponseCollectionCompoundDocument) SetData(v []GetSegmentListResponseCollectionCompoundDocumentDataInner)`

SetData sets Data field to given value.


### GetLinks

`func (o *GetSegmentListResponseCollectionCompoundDocument) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetSegmentListResponseCollectionCompoundDocument) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetSegmentListResponseCollectionCompoundDocument) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.


### GetIncluded

`func (o *GetSegmentListResponseCollectionCompoundDocument) GetIncluded() []TagResponseObjectResource`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GetSegmentListResponseCollectionCompoundDocument) GetIncludedOk() (*[]TagResponseObjectResource, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GetSegmentListResponseCollectionCompoundDocument) SetIncluded(v []TagResponseObjectResource)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GetSegmentListResponseCollectionCompoundDocument) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


