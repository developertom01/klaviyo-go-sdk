# CatalogVariantDeleteQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CatalogVariantEnum**](CatalogVariantEnum.md) |  | 
**Id** | **string** | The catalog variant ID is a compound ID (string), with format: &#x60;{integration}:::{catalog}:::{external_id}&#x60;. Currently, the only supported integration type is &#x60;$custom&#x60;, and the only supported catalog is &#x60;$default&#x60;. | 

## Methods

### NewCatalogVariantDeleteQueryResourceObject

`func NewCatalogVariantDeleteQueryResourceObject(type_ CatalogVariantEnum, id string, ) *CatalogVariantDeleteQueryResourceObject`

NewCatalogVariantDeleteQueryResourceObject instantiates a new CatalogVariantDeleteQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogVariantDeleteQueryResourceObjectWithDefaults

`func NewCatalogVariantDeleteQueryResourceObjectWithDefaults() *CatalogVariantDeleteQueryResourceObject`

NewCatalogVariantDeleteQueryResourceObjectWithDefaults instantiates a new CatalogVariantDeleteQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogVariantDeleteQueryResourceObject) GetType() CatalogVariantEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogVariantDeleteQueryResourceObject) GetTypeOk() (*CatalogVariantEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogVariantDeleteQueryResourceObject) SetType(v CatalogVariantEnum)`

SetType sets Type field to given value.


### GetId

`func (o *CatalogVariantDeleteQueryResourceObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogVariantDeleteQueryResourceObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogVariantDeleteQueryResourceObject) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


