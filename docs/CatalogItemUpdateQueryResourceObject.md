# CatalogItemUpdateQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CatalogItemEnum**](CatalogItemEnum.md) |  | 
**Id** | **string** | The catalog item ID is a compound ID (string), with format: &#x60;{integration}:::{catalog}:::{external_id}&#x60;. Currently, the only supported integration type is &#x60;$custom&#x60;, and the only supported catalog is &#x60;$default&#x60;. | 
**Attributes** | [**CatalogItemUpdateQueryResourceObjectAttributes**](CatalogItemUpdateQueryResourceObjectAttributes.md) |  | 
**Relationships** | Pointer to [**CatalogItemCreateQueryResourceObjectRelationships**](CatalogItemCreateQueryResourceObjectRelationships.md) |  | [optional] 

## Methods

### NewCatalogItemUpdateQueryResourceObject

`func NewCatalogItemUpdateQueryResourceObject(type_ CatalogItemEnum, id string, attributes CatalogItemUpdateQueryResourceObjectAttributes, ) *CatalogItemUpdateQueryResourceObject`

NewCatalogItemUpdateQueryResourceObject instantiates a new CatalogItemUpdateQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogItemUpdateQueryResourceObjectWithDefaults

`func NewCatalogItemUpdateQueryResourceObjectWithDefaults() *CatalogItemUpdateQueryResourceObject`

NewCatalogItemUpdateQueryResourceObjectWithDefaults instantiates a new CatalogItemUpdateQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogItemUpdateQueryResourceObject) GetType() CatalogItemEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogItemUpdateQueryResourceObject) GetTypeOk() (*CatalogItemEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogItemUpdateQueryResourceObject) SetType(v CatalogItemEnum)`

SetType sets Type field to given value.


### GetId

`func (o *CatalogItemUpdateQueryResourceObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogItemUpdateQueryResourceObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogItemUpdateQueryResourceObject) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CatalogItemUpdateQueryResourceObject) GetAttributes() CatalogItemUpdateQueryResourceObjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CatalogItemUpdateQueryResourceObject) GetAttributesOk() (*CatalogItemUpdateQueryResourceObjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CatalogItemUpdateQueryResourceObject) SetAttributes(v CatalogItemUpdateQueryResourceObjectAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CatalogItemUpdateQueryResourceObject) GetRelationships() CatalogItemCreateQueryResourceObjectRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CatalogItemUpdateQueryResourceObject) GetRelationshipsOk() (*CatalogItemCreateQueryResourceObjectRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CatalogItemUpdateQueryResourceObject) SetRelationships(v CatalogItemCreateQueryResourceObjectRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CatalogItemUpdateQueryResourceObject) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


