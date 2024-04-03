# CatalogItemCreateQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CatalogItemEnum**](CatalogItemEnum.md) |  | 
**Attributes** | [**CatalogItemCreateQueryResourceObjectAttributes**](CatalogItemCreateQueryResourceObjectAttributes.md) |  | 
**Relationships** | Pointer to [**CatalogItemCreateQueryResourceObjectRelationships**](CatalogItemCreateQueryResourceObjectRelationships.md) |  | [optional] 

## Methods

### NewCatalogItemCreateQueryResourceObject

`func NewCatalogItemCreateQueryResourceObject(type_ CatalogItemEnum, attributes CatalogItemCreateQueryResourceObjectAttributes, ) *CatalogItemCreateQueryResourceObject`

NewCatalogItemCreateQueryResourceObject instantiates a new CatalogItemCreateQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogItemCreateQueryResourceObjectWithDefaults

`func NewCatalogItemCreateQueryResourceObjectWithDefaults() *CatalogItemCreateQueryResourceObject`

NewCatalogItemCreateQueryResourceObjectWithDefaults instantiates a new CatalogItemCreateQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogItemCreateQueryResourceObject) GetType() CatalogItemEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogItemCreateQueryResourceObject) GetTypeOk() (*CatalogItemEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogItemCreateQueryResourceObject) SetType(v CatalogItemEnum)`

SetType sets Type field to given value.


### GetAttributes

`func (o *CatalogItemCreateQueryResourceObject) GetAttributes() CatalogItemCreateQueryResourceObjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CatalogItemCreateQueryResourceObject) GetAttributesOk() (*CatalogItemCreateQueryResourceObjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CatalogItemCreateQueryResourceObject) SetAttributes(v CatalogItemCreateQueryResourceObjectAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CatalogItemCreateQueryResourceObject) GetRelationships() CatalogItemCreateQueryResourceObjectRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CatalogItemCreateQueryResourceObject) GetRelationshipsOk() (*CatalogItemCreateQueryResourceObjectRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CatalogItemCreateQueryResourceObject) SetRelationships(v CatalogItemCreateQueryResourceObjectRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CatalogItemCreateQueryResourceObject) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


