# GetCatalogItemUpdateJobResponseCompoundDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInner**](GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInner.md) |  | 
**Included** | Pointer to [**[]CatalogItemResponseObjectResource**](CatalogItemResponseObjectResource.md) |  | [optional] 

## Methods

### NewGetCatalogItemUpdateJobResponseCompoundDocument

`func NewGetCatalogItemUpdateJobResponseCompoundDocument(data GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInner, ) *GetCatalogItemUpdateJobResponseCompoundDocument`

NewGetCatalogItemUpdateJobResponseCompoundDocument instantiates a new GetCatalogItemUpdateJobResponseCompoundDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCatalogItemUpdateJobResponseCompoundDocumentWithDefaults

`func NewGetCatalogItemUpdateJobResponseCompoundDocumentWithDefaults() *GetCatalogItemUpdateJobResponseCompoundDocument`

NewGetCatalogItemUpdateJobResponseCompoundDocumentWithDefaults instantiates a new GetCatalogItemUpdateJobResponseCompoundDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetCatalogItemUpdateJobResponseCompoundDocument) GetData() GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCatalogItemUpdateJobResponseCompoundDocument) GetDataOk() (*GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCatalogItemUpdateJobResponseCompoundDocument) SetData(v GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInner)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GetCatalogItemUpdateJobResponseCompoundDocument) GetIncluded() []CatalogItemResponseObjectResource`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GetCatalogItemUpdateJobResponseCompoundDocument) GetIncludedOk() (*[]CatalogItemResponseObjectResource, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GetCatalogItemUpdateJobResponseCompoundDocument) SetIncluded(v []CatalogItemResponseObjectResource)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GetCatalogItemUpdateJobResponseCompoundDocument) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


