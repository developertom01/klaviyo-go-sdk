# GetCatalogItemResponseCollectionCompoundDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GetCatalogItemResponseCollectionCompoundDocumentDataInner**](GetCatalogItemResponseCollectionCompoundDocumentDataInner.md) |  | 
**Links** | [**CollectionLinks**](CollectionLinks.md) |  | 
**Included** | Pointer to [**[]CatalogVariantResponseObjectResource**](CatalogVariantResponseObjectResource.md) |  | [optional] 

## Methods

### NewGetCatalogItemResponseCollectionCompoundDocument

`func NewGetCatalogItemResponseCollectionCompoundDocument(data []GetCatalogItemResponseCollectionCompoundDocumentDataInner, links CollectionLinks, ) *GetCatalogItemResponseCollectionCompoundDocument`

NewGetCatalogItemResponseCollectionCompoundDocument instantiates a new GetCatalogItemResponseCollectionCompoundDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCatalogItemResponseCollectionCompoundDocumentWithDefaults

`func NewGetCatalogItemResponseCollectionCompoundDocumentWithDefaults() *GetCatalogItemResponseCollectionCompoundDocument`

NewGetCatalogItemResponseCollectionCompoundDocumentWithDefaults instantiates a new GetCatalogItemResponseCollectionCompoundDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetCatalogItemResponseCollectionCompoundDocument) GetData() []GetCatalogItemResponseCollectionCompoundDocumentDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCatalogItemResponseCollectionCompoundDocument) GetDataOk() (*[]GetCatalogItemResponseCollectionCompoundDocumentDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCatalogItemResponseCollectionCompoundDocument) SetData(v []GetCatalogItemResponseCollectionCompoundDocumentDataInner)`

SetData sets Data field to given value.


### GetLinks

`func (o *GetCatalogItemResponseCollectionCompoundDocument) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetCatalogItemResponseCollectionCompoundDocument) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetCatalogItemResponseCollectionCompoundDocument) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.


### GetIncluded

`func (o *GetCatalogItemResponseCollectionCompoundDocument) GetIncluded() []CatalogVariantResponseObjectResource`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GetCatalogItemResponseCollectionCompoundDocument) GetIncludedOk() (*[]CatalogVariantResponseObjectResource, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GetCatalogItemResponseCollectionCompoundDocument) SetIncluded(v []CatalogVariantResponseObjectResource)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GetCatalogItemResponseCollectionCompoundDocument) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


