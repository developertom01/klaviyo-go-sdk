# GetCatalogCategoryItemListResponseCollectionDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CatalogItemEnum**](CatalogItemEnum.md) |  | 
**Id** | **string** | A list of catalog item IDs that are in the given category. | 

## Methods

### NewGetCatalogCategoryItemListResponseCollectionDataInner

`func NewGetCatalogCategoryItemListResponseCollectionDataInner(type_ CatalogItemEnum, id string, ) *GetCatalogCategoryItemListResponseCollectionDataInner`

NewGetCatalogCategoryItemListResponseCollectionDataInner instantiates a new GetCatalogCategoryItemListResponseCollectionDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCatalogCategoryItemListResponseCollectionDataInnerWithDefaults

`func NewGetCatalogCategoryItemListResponseCollectionDataInnerWithDefaults() *GetCatalogCategoryItemListResponseCollectionDataInner`

NewGetCatalogCategoryItemListResponseCollectionDataInnerWithDefaults instantiates a new GetCatalogCategoryItemListResponseCollectionDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetCatalogCategoryItemListResponseCollectionDataInner) GetType() CatalogItemEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCatalogCategoryItemListResponseCollectionDataInner) GetTypeOk() (*CatalogItemEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCatalogCategoryItemListResponseCollectionDataInner) SetType(v CatalogItemEnum)`

SetType sets Type field to given value.


### GetId

`func (o *GetCatalogCategoryItemListResponseCollectionDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCatalogCategoryItemListResponseCollectionDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCatalogCategoryItemListResponseCollectionDataInner) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


