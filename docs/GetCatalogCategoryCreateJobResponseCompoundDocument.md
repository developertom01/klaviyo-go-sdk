# GetCatalogCategoryCreateJobResponseCompoundDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GetCatalogCategoryCreateJobResponseCollectionCompoundDocumentDataInner**](GetCatalogCategoryCreateJobResponseCollectionCompoundDocumentDataInner.md) |  | 
**Included** | Pointer to [**[]CatalogCategoryResponseObjectResource**](CatalogCategoryResponseObjectResource.md) |  | [optional] 

## Methods

### NewGetCatalogCategoryCreateJobResponseCompoundDocument

`func NewGetCatalogCategoryCreateJobResponseCompoundDocument(data GetCatalogCategoryCreateJobResponseCollectionCompoundDocumentDataInner, ) *GetCatalogCategoryCreateJobResponseCompoundDocument`

NewGetCatalogCategoryCreateJobResponseCompoundDocument instantiates a new GetCatalogCategoryCreateJobResponseCompoundDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCatalogCategoryCreateJobResponseCompoundDocumentWithDefaults

`func NewGetCatalogCategoryCreateJobResponseCompoundDocumentWithDefaults() *GetCatalogCategoryCreateJobResponseCompoundDocument`

NewGetCatalogCategoryCreateJobResponseCompoundDocumentWithDefaults instantiates a new GetCatalogCategoryCreateJobResponseCompoundDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetCatalogCategoryCreateJobResponseCompoundDocument) GetData() GetCatalogCategoryCreateJobResponseCollectionCompoundDocumentDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCatalogCategoryCreateJobResponseCompoundDocument) GetDataOk() (*GetCatalogCategoryCreateJobResponseCollectionCompoundDocumentDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCatalogCategoryCreateJobResponseCompoundDocument) SetData(v GetCatalogCategoryCreateJobResponseCollectionCompoundDocumentDataInner)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GetCatalogCategoryCreateJobResponseCompoundDocument) GetIncluded() []CatalogCategoryResponseObjectResource`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GetCatalogCategoryCreateJobResponseCompoundDocument) GetIncludedOk() (*[]CatalogCategoryResponseObjectResource, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GetCatalogCategoryCreateJobResponseCompoundDocument) SetIncluded(v []CatalogCategoryResponseObjectResource)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GetCatalogCategoryCreateJobResponseCompoundDocument) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


