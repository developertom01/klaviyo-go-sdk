# CatalogItemDeleteQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CatalogItemEnum**](CatalogItemEnum.md) |  | 
**Id** | **string** | The catalog item ID is a compound ID (string), with format: &#x60;{integration}:::{catalog}:::{external_id}&#x60;. Currently, the only supported integration type is &#x60;$custom&#x60;, and the only supported catalog is &#x60;$default&#x60;. | 

## Methods

### NewCatalogItemDeleteQueryResourceObject

`func NewCatalogItemDeleteQueryResourceObject(type_ CatalogItemEnum, id string, ) *CatalogItemDeleteQueryResourceObject`

NewCatalogItemDeleteQueryResourceObject instantiates a new CatalogItemDeleteQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogItemDeleteQueryResourceObjectWithDefaults

`func NewCatalogItemDeleteQueryResourceObjectWithDefaults() *CatalogItemDeleteQueryResourceObject`

NewCatalogItemDeleteQueryResourceObjectWithDefaults instantiates a new CatalogItemDeleteQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogItemDeleteQueryResourceObject) GetType() CatalogItemEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogItemDeleteQueryResourceObject) GetTypeOk() (*CatalogItemEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogItemDeleteQueryResourceObject) SetType(v CatalogItemEnum)`

SetType sets Type field to given value.


### GetId

`func (o *CatalogItemDeleteQueryResourceObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogItemDeleteQueryResourceObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogItemDeleteQueryResourceObject) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


