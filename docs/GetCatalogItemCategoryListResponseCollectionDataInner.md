# GetCatalogItemCategoryListResponseCollectionDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CatalogCategoryEnum**](CatalogCategoryEnum.md) |  | 
**Id** | **string** | A list of catalog category IDs representing the categories the item is in | 

## Methods

### NewGetCatalogItemCategoryListResponseCollectionDataInner

`func NewGetCatalogItemCategoryListResponseCollectionDataInner(type_ CatalogCategoryEnum, id string, ) *GetCatalogItemCategoryListResponseCollectionDataInner`

NewGetCatalogItemCategoryListResponseCollectionDataInner instantiates a new GetCatalogItemCategoryListResponseCollectionDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCatalogItemCategoryListResponseCollectionDataInnerWithDefaults

`func NewGetCatalogItemCategoryListResponseCollectionDataInnerWithDefaults() *GetCatalogItemCategoryListResponseCollectionDataInner`

NewGetCatalogItemCategoryListResponseCollectionDataInnerWithDefaults instantiates a new GetCatalogItemCategoryListResponseCollectionDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetCatalogItemCategoryListResponseCollectionDataInner) GetType() CatalogCategoryEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCatalogItemCategoryListResponseCollectionDataInner) GetTypeOk() (*CatalogCategoryEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCatalogItemCategoryListResponseCollectionDataInner) SetType(v CatalogCategoryEnum)`

SetType sets Type field to given value.


### GetId

`func (o *GetCatalogItemCategoryListResponseCollectionDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCatalogItemCategoryListResponseCollectionDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCatalogItemCategoryListResponseCollectionDataInner) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


