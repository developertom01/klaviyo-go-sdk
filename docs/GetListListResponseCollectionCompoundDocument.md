# GetListListResponseCollectionCompoundDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GetListListResponseCollectionCompoundDocumentDataInner**](GetListListResponseCollectionCompoundDocumentDataInner.md) |  | 
**Links** | [**CollectionLinks**](CollectionLinks.md) |  | 
**Included** | Pointer to [**[]TagResponseObjectResource**](TagResponseObjectResource.md) |  | [optional] 

## Methods

### NewGetListListResponseCollectionCompoundDocument

`func NewGetListListResponseCollectionCompoundDocument(data []GetListListResponseCollectionCompoundDocumentDataInner, links CollectionLinks, ) *GetListListResponseCollectionCompoundDocument`

NewGetListListResponseCollectionCompoundDocument instantiates a new GetListListResponseCollectionCompoundDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetListListResponseCollectionCompoundDocumentWithDefaults

`func NewGetListListResponseCollectionCompoundDocumentWithDefaults() *GetListListResponseCollectionCompoundDocument`

NewGetListListResponseCollectionCompoundDocumentWithDefaults instantiates a new GetListListResponseCollectionCompoundDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetListListResponseCollectionCompoundDocument) GetData() []GetListListResponseCollectionCompoundDocumentDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetListListResponseCollectionCompoundDocument) GetDataOk() (*[]GetListListResponseCollectionCompoundDocumentDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetListListResponseCollectionCompoundDocument) SetData(v []GetListListResponseCollectionCompoundDocumentDataInner)`

SetData sets Data field to given value.


### GetLinks

`func (o *GetListListResponseCollectionCompoundDocument) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetListListResponseCollectionCompoundDocument) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetListListResponseCollectionCompoundDocument) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.


### GetIncluded

`func (o *GetListListResponseCollectionCompoundDocument) GetIncluded() []TagResponseObjectResource`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GetListListResponseCollectionCompoundDocument) GetIncludedOk() (*[]TagResponseObjectResource, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GetListListResponseCollectionCompoundDocument) SetIncluded(v []TagResponseObjectResource)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GetListListResponseCollectionCompoundDocument) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


