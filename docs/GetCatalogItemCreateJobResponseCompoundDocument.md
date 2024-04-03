# GetCatalogItemCreateJobResponseCompoundDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner**](GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner.md) |  | 
**Included** | Pointer to [**[]CatalogItemResponseObjectResource**](CatalogItemResponseObjectResource.md) |  | [optional] 

## Methods

### NewGetCatalogItemCreateJobResponseCompoundDocument

`func NewGetCatalogItemCreateJobResponseCompoundDocument(data GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner, ) *GetCatalogItemCreateJobResponseCompoundDocument`

NewGetCatalogItemCreateJobResponseCompoundDocument instantiates a new GetCatalogItemCreateJobResponseCompoundDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCatalogItemCreateJobResponseCompoundDocumentWithDefaults

`func NewGetCatalogItemCreateJobResponseCompoundDocumentWithDefaults() *GetCatalogItemCreateJobResponseCompoundDocument`

NewGetCatalogItemCreateJobResponseCompoundDocumentWithDefaults instantiates a new GetCatalogItemCreateJobResponseCompoundDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetCatalogItemCreateJobResponseCompoundDocument) GetData() GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCatalogItemCreateJobResponseCompoundDocument) GetDataOk() (*GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCatalogItemCreateJobResponseCompoundDocument) SetData(v GetCatalogItemCreateJobResponseCollectionCompoundDocumentDataInner)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GetCatalogItemCreateJobResponseCompoundDocument) GetIncluded() []CatalogItemResponseObjectResource`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GetCatalogItemCreateJobResponseCompoundDocument) GetIncludedOk() (*[]CatalogItemResponseObjectResource, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GetCatalogItemCreateJobResponseCompoundDocument) SetIncluded(v []CatalogItemResponseObjectResource)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GetCatalogItemCreateJobResponseCompoundDocument) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


