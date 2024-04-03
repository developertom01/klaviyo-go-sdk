# CatalogCategoryUpdateQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CatalogCategoryEnum**](CatalogCategoryEnum.md) |  | 
**Id** | **string** | The catalog category ID is a compound ID (string), with format: &#x60;{integration}:::{catalog}:::{external_id}&#x60;. Currently, the only supported integration type is &#x60;$custom&#x60;, and the only supported catalog is &#x60;$default&#x60;. | 
**Attributes** | [**CatalogCategoryUpdateQueryResourceObjectAttributes**](CatalogCategoryUpdateQueryResourceObjectAttributes.md) |  | 
**Relationships** | Pointer to [**CatalogCategoryCreateQueryResourceObjectRelationships**](CatalogCategoryCreateQueryResourceObjectRelationships.md) |  | [optional] 

## Methods

### NewCatalogCategoryUpdateQueryResourceObject

`func NewCatalogCategoryUpdateQueryResourceObject(type_ CatalogCategoryEnum, id string, attributes CatalogCategoryUpdateQueryResourceObjectAttributes, ) *CatalogCategoryUpdateQueryResourceObject`

NewCatalogCategoryUpdateQueryResourceObject instantiates a new CatalogCategoryUpdateQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogCategoryUpdateQueryResourceObjectWithDefaults

`func NewCatalogCategoryUpdateQueryResourceObjectWithDefaults() *CatalogCategoryUpdateQueryResourceObject`

NewCatalogCategoryUpdateQueryResourceObjectWithDefaults instantiates a new CatalogCategoryUpdateQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogCategoryUpdateQueryResourceObject) GetType() CatalogCategoryEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogCategoryUpdateQueryResourceObject) GetTypeOk() (*CatalogCategoryEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogCategoryUpdateQueryResourceObject) SetType(v CatalogCategoryEnum)`

SetType sets Type field to given value.


### GetId

`func (o *CatalogCategoryUpdateQueryResourceObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogCategoryUpdateQueryResourceObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogCategoryUpdateQueryResourceObject) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CatalogCategoryUpdateQueryResourceObject) GetAttributes() CatalogCategoryUpdateQueryResourceObjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CatalogCategoryUpdateQueryResourceObject) GetAttributesOk() (*CatalogCategoryUpdateQueryResourceObjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CatalogCategoryUpdateQueryResourceObject) SetAttributes(v CatalogCategoryUpdateQueryResourceObjectAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CatalogCategoryUpdateQueryResourceObject) GetRelationships() CatalogCategoryCreateQueryResourceObjectRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CatalogCategoryUpdateQueryResourceObject) GetRelationshipsOk() (*CatalogCategoryCreateQueryResourceObjectRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CatalogCategoryUpdateQueryResourceObject) SetRelationships(v CatalogCategoryCreateQueryResourceObjectRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CatalogCategoryUpdateQueryResourceObject) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


