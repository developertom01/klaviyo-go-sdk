# CatalogCategoryDeleteQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CatalogCategoryEnum**](CatalogCategoryEnum.md) |  | 
**Id** | **string** | The catalog category ID is a compound ID (string), with format: &#x60;{integration}:::{catalog}:::{external_id}&#x60;. Currently, the only supported integration type is &#x60;$custom&#x60;, and the only supported catalog is &#x60;$default&#x60;. | 

## Methods

### NewCatalogCategoryDeleteQueryResourceObject

`func NewCatalogCategoryDeleteQueryResourceObject(type_ CatalogCategoryEnum, id string, ) *CatalogCategoryDeleteQueryResourceObject`

NewCatalogCategoryDeleteQueryResourceObject instantiates a new CatalogCategoryDeleteQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogCategoryDeleteQueryResourceObjectWithDefaults

`func NewCatalogCategoryDeleteQueryResourceObjectWithDefaults() *CatalogCategoryDeleteQueryResourceObject`

NewCatalogCategoryDeleteQueryResourceObjectWithDefaults instantiates a new CatalogCategoryDeleteQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogCategoryDeleteQueryResourceObject) GetType() CatalogCategoryEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogCategoryDeleteQueryResourceObject) GetTypeOk() (*CatalogCategoryEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogCategoryDeleteQueryResourceObject) SetType(v CatalogCategoryEnum)`

SetType sets Type field to given value.


### GetId

`func (o *CatalogCategoryDeleteQueryResourceObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogCategoryDeleteQueryResourceObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogCategoryDeleteQueryResourceObject) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


