# GetCatalogItemResponseCompoundDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GetCatalogItemResponseCollectionCompoundDocumentDataInner**](GetCatalogItemResponseCollectionCompoundDocumentDataInner.md) |  | 
**Included** | Pointer to [**[]CatalogVariantResponseObjectResource**](CatalogVariantResponseObjectResource.md) |  | [optional] 

## Methods

### NewGetCatalogItemResponseCompoundDocument

`func NewGetCatalogItemResponseCompoundDocument(data GetCatalogItemResponseCollectionCompoundDocumentDataInner, ) *GetCatalogItemResponseCompoundDocument`

NewGetCatalogItemResponseCompoundDocument instantiates a new GetCatalogItemResponseCompoundDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCatalogItemResponseCompoundDocumentWithDefaults

`func NewGetCatalogItemResponseCompoundDocumentWithDefaults() *GetCatalogItemResponseCompoundDocument`

NewGetCatalogItemResponseCompoundDocumentWithDefaults instantiates a new GetCatalogItemResponseCompoundDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetCatalogItemResponseCompoundDocument) GetData() GetCatalogItemResponseCollectionCompoundDocumentDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCatalogItemResponseCompoundDocument) GetDataOk() (*GetCatalogItemResponseCollectionCompoundDocumentDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCatalogItemResponseCompoundDocument) SetData(v GetCatalogItemResponseCollectionCompoundDocumentDataInner)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GetCatalogItemResponseCompoundDocument) GetIncluded() []CatalogVariantResponseObjectResource`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GetCatalogItemResponseCompoundDocument) GetIncludedOk() (*[]CatalogVariantResponseObjectResource, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GetCatalogItemResponseCompoundDocument) SetIncluded(v []CatalogVariantResponseObjectResource)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GetCatalogItemResponseCompoundDocument) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


