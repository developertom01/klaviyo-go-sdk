# GetTagResponseCollectionCompoundDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GetTagResponseCollectionCompoundDocumentDataInner**](GetTagResponseCollectionCompoundDocumentDataInner.md) |  | 
**Links** | [**CollectionLinks**](CollectionLinks.md) |  | 
**Included** | Pointer to [**[]TagGroupResponseObjectResource**](TagGroupResponseObjectResource.md) |  | [optional] 

## Methods

### NewGetTagResponseCollectionCompoundDocument

`func NewGetTagResponseCollectionCompoundDocument(data []GetTagResponseCollectionCompoundDocumentDataInner, links CollectionLinks, ) *GetTagResponseCollectionCompoundDocument`

NewGetTagResponseCollectionCompoundDocument instantiates a new GetTagResponseCollectionCompoundDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTagResponseCollectionCompoundDocumentWithDefaults

`func NewGetTagResponseCollectionCompoundDocumentWithDefaults() *GetTagResponseCollectionCompoundDocument`

NewGetTagResponseCollectionCompoundDocumentWithDefaults instantiates a new GetTagResponseCollectionCompoundDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetTagResponseCollectionCompoundDocument) GetData() []GetTagResponseCollectionCompoundDocumentDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTagResponseCollectionCompoundDocument) GetDataOk() (*[]GetTagResponseCollectionCompoundDocumentDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTagResponseCollectionCompoundDocument) SetData(v []GetTagResponseCollectionCompoundDocumentDataInner)`

SetData sets Data field to given value.


### GetLinks

`func (o *GetTagResponseCollectionCompoundDocument) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetTagResponseCollectionCompoundDocument) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetTagResponseCollectionCompoundDocument) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.


### GetIncluded

`func (o *GetTagResponseCollectionCompoundDocument) GetIncluded() []TagGroupResponseObjectResource`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GetTagResponseCollectionCompoundDocument) GetIncludedOk() (*[]TagGroupResponseObjectResource, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GetTagResponseCollectionCompoundDocument) SetIncluded(v []TagGroupResponseObjectResource)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GetTagResponseCollectionCompoundDocument) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


