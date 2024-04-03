# GetCatalogVariantUpdateJobResponseCompoundDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInner**](GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInner.md) |  | 
**Included** | Pointer to [**[]CatalogVariantResponseObjectResource**](CatalogVariantResponseObjectResource.md) |  | [optional] 

## Methods

### NewGetCatalogVariantUpdateJobResponseCompoundDocument

`func NewGetCatalogVariantUpdateJobResponseCompoundDocument(data GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInner, ) *GetCatalogVariantUpdateJobResponseCompoundDocument`

NewGetCatalogVariantUpdateJobResponseCompoundDocument instantiates a new GetCatalogVariantUpdateJobResponseCompoundDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCatalogVariantUpdateJobResponseCompoundDocumentWithDefaults

`func NewGetCatalogVariantUpdateJobResponseCompoundDocumentWithDefaults() *GetCatalogVariantUpdateJobResponseCompoundDocument`

NewGetCatalogVariantUpdateJobResponseCompoundDocumentWithDefaults instantiates a new GetCatalogVariantUpdateJobResponseCompoundDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetCatalogVariantUpdateJobResponseCompoundDocument) GetData() GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCatalogVariantUpdateJobResponseCompoundDocument) GetDataOk() (*GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCatalogVariantUpdateJobResponseCompoundDocument) SetData(v GetCatalogVariantUpdateJobResponseCollectionCompoundDocumentDataInner)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GetCatalogVariantUpdateJobResponseCompoundDocument) GetIncluded() []CatalogVariantResponseObjectResource`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GetCatalogVariantUpdateJobResponseCompoundDocument) GetIncludedOk() (*[]CatalogVariantResponseObjectResource, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GetCatalogVariantUpdateJobResponseCompoundDocument) SetIncluded(v []CatalogVariantResponseObjectResource)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GetCatalogVariantUpdateJobResponseCompoundDocument) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


