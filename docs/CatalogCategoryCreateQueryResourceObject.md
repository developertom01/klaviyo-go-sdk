# CatalogCategoryCreateQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CatalogCategoryEnum**](CatalogCategoryEnum.md) |  | 
**Attributes** | [**CatalogCategoryCreateQueryResourceObjectAttributes**](CatalogCategoryCreateQueryResourceObjectAttributes.md) |  | 
**Relationships** | Pointer to [**CatalogCategoryCreateQueryResourceObjectRelationships**](CatalogCategoryCreateQueryResourceObjectRelationships.md) |  | [optional] 

## Methods

### NewCatalogCategoryCreateQueryResourceObject

`func NewCatalogCategoryCreateQueryResourceObject(type_ CatalogCategoryEnum, attributes CatalogCategoryCreateQueryResourceObjectAttributes, ) *CatalogCategoryCreateQueryResourceObject`

NewCatalogCategoryCreateQueryResourceObject instantiates a new CatalogCategoryCreateQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogCategoryCreateQueryResourceObjectWithDefaults

`func NewCatalogCategoryCreateQueryResourceObjectWithDefaults() *CatalogCategoryCreateQueryResourceObject`

NewCatalogCategoryCreateQueryResourceObjectWithDefaults instantiates a new CatalogCategoryCreateQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogCategoryCreateQueryResourceObject) GetType() CatalogCategoryEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogCategoryCreateQueryResourceObject) GetTypeOk() (*CatalogCategoryEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogCategoryCreateQueryResourceObject) SetType(v CatalogCategoryEnum)`

SetType sets Type field to given value.


### GetAttributes

`func (o *CatalogCategoryCreateQueryResourceObject) GetAttributes() CatalogCategoryCreateQueryResourceObjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CatalogCategoryCreateQueryResourceObject) GetAttributesOk() (*CatalogCategoryCreateQueryResourceObjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CatalogCategoryCreateQueryResourceObject) SetAttributes(v CatalogCategoryCreateQueryResourceObjectAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CatalogCategoryCreateQueryResourceObject) GetRelationships() CatalogCategoryCreateQueryResourceObjectRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CatalogCategoryCreateQueryResourceObject) GetRelationshipsOk() (*CatalogCategoryCreateQueryResourceObjectRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CatalogCategoryCreateQueryResourceObject) SetRelationships(v CatalogCategoryCreateQueryResourceObjectRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CatalogCategoryCreateQueryResourceObject) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


