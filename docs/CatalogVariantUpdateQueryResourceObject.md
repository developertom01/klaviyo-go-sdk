# CatalogVariantUpdateQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CatalogVariantEnum**](CatalogVariantEnum.md) |  | 
**Id** | **string** | The catalog variant ID is a compound ID (string), with format: &#x60;{integration}:::{catalog}:::{external_id}&#x60;. Currently, the only supported integration type is &#x60;$custom&#x60;, and the only supported catalog is &#x60;$default&#x60;. | 
**Attributes** | [**CatalogVariantUpdateQueryResourceObjectAttributes**](CatalogVariantUpdateQueryResourceObjectAttributes.md) |  | 

## Methods

### NewCatalogVariantUpdateQueryResourceObject

`func NewCatalogVariantUpdateQueryResourceObject(type_ CatalogVariantEnum, id string, attributes CatalogVariantUpdateQueryResourceObjectAttributes, ) *CatalogVariantUpdateQueryResourceObject`

NewCatalogVariantUpdateQueryResourceObject instantiates a new CatalogVariantUpdateQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogVariantUpdateQueryResourceObjectWithDefaults

`func NewCatalogVariantUpdateQueryResourceObjectWithDefaults() *CatalogVariantUpdateQueryResourceObject`

NewCatalogVariantUpdateQueryResourceObjectWithDefaults instantiates a new CatalogVariantUpdateQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogVariantUpdateQueryResourceObject) GetType() CatalogVariantEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogVariantUpdateQueryResourceObject) GetTypeOk() (*CatalogVariantEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogVariantUpdateQueryResourceObject) SetType(v CatalogVariantEnum)`

SetType sets Type field to given value.


### GetId

`func (o *CatalogVariantUpdateQueryResourceObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogVariantUpdateQueryResourceObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogVariantUpdateQueryResourceObject) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CatalogVariantUpdateQueryResourceObject) GetAttributes() CatalogVariantUpdateQueryResourceObjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CatalogVariantUpdateQueryResourceObject) GetAttributesOk() (*CatalogVariantUpdateQueryResourceObjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CatalogVariantUpdateQueryResourceObject) SetAttributes(v CatalogVariantUpdateQueryResourceObjectAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


