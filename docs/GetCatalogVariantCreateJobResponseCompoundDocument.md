# GetCatalogVariantCreateJobResponseCompoundDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GetCatalogVariantCreateJobResponseCollectionCompoundDocumentDataInner**](GetCatalogVariantCreateJobResponseCollectionCompoundDocumentDataInner.md) |  | 
**Included** | Pointer to [**[]CatalogVariantResponseObjectResource**](CatalogVariantResponseObjectResource.md) |  | [optional] 

## Methods

### NewGetCatalogVariantCreateJobResponseCompoundDocument

`func NewGetCatalogVariantCreateJobResponseCompoundDocument(data GetCatalogVariantCreateJobResponseCollectionCompoundDocumentDataInner, ) *GetCatalogVariantCreateJobResponseCompoundDocument`

NewGetCatalogVariantCreateJobResponseCompoundDocument instantiates a new GetCatalogVariantCreateJobResponseCompoundDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCatalogVariantCreateJobResponseCompoundDocumentWithDefaults

`func NewGetCatalogVariantCreateJobResponseCompoundDocumentWithDefaults() *GetCatalogVariantCreateJobResponseCompoundDocument`

NewGetCatalogVariantCreateJobResponseCompoundDocumentWithDefaults instantiates a new GetCatalogVariantCreateJobResponseCompoundDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetCatalogVariantCreateJobResponseCompoundDocument) GetData() GetCatalogVariantCreateJobResponseCollectionCompoundDocumentDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCatalogVariantCreateJobResponseCompoundDocument) GetDataOk() (*GetCatalogVariantCreateJobResponseCollectionCompoundDocumentDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCatalogVariantCreateJobResponseCompoundDocument) SetData(v GetCatalogVariantCreateJobResponseCollectionCompoundDocumentDataInner)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GetCatalogVariantCreateJobResponseCompoundDocument) GetIncluded() []CatalogVariantResponseObjectResource`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GetCatalogVariantCreateJobResponseCompoundDocument) GetIncludedOk() (*[]CatalogVariantResponseObjectResource, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GetCatalogVariantCreateJobResponseCompoundDocument) SetIncluded(v []CatalogVariantResponseObjectResource)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GetCatalogVariantCreateJobResponseCompoundDocument) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


