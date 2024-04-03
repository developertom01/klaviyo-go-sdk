# GetCatalogCategoryResponseCollectionDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CatalogCategoryEnum**](CatalogCategoryEnum.md) |  | 
**Id** | **string** | The catalog category ID is a compound ID (string), with format: &#x60;{integration}:::{catalog}:::{external_id}&#x60;. Currently, the only supported integration type is &#x60;$custom&#x60;, and the only supported catalog is &#x60;$default&#x60;. | 
**Attributes** | [**CatalogCategoryResponseObjectResourceAttributes**](CatalogCategoryResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 
**Relationships** | Pointer to [**GetCatalogCategoryResponseCollectionDataInnerAllOfRelationships**](GetCatalogCategoryResponseCollectionDataInnerAllOfRelationships.md) |  | [optional] 

## Methods

### NewGetCatalogCategoryResponseCollectionDataInner

`func NewGetCatalogCategoryResponseCollectionDataInner(type_ CatalogCategoryEnum, id string, attributes CatalogCategoryResponseObjectResourceAttributes, links ObjectLinks, ) *GetCatalogCategoryResponseCollectionDataInner`

NewGetCatalogCategoryResponseCollectionDataInner instantiates a new GetCatalogCategoryResponseCollectionDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCatalogCategoryResponseCollectionDataInnerWithDefaults

`func NewGetCatalogCategoryResponseCollectionDataInnerWithDefaults() *GetCatalogCategoryResponseCollectionDataInner`

NewGetCatalogCategoryResponseCollectionDataInnerWithDefaults instantiates a new GetCatalogCategoryResponseCollectionDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetCatalogCategoryResponseCollectionDataInner) GetType() CatalogCategoryEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCatalogCategoryResponseCollectionDataInner) GetTypeOk() (*CatalogCategoryEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCatalogCategoryResponseCollectionDataInner) SetType(v CatalogCategoryEnum)`

SetType sets Type field to given value.


### GetId

`func (o *GetCatalogCategoryResponseCollectionDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCatalogCategoryResponseCollectionDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCatalogCategoryResponseCollectionDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GetCatalogCategoryResponseCollectionDataInner) GetAttributes() CatalogCategoryResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetCatalogCategoryResponseCollectionDataInner) GetAttributesOk() (*CatalogCategoryResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetCatalogCategoryResponseCollectionDataInner) SetAttributes(v CatalogCategoryResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *GetCatalogCategoryResponseCollectionDataInner) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetCatalogCategoryResponseCollectionDataInner) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetCatalogCategoryResponseCollectionDataInner) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.


### GetRelationships

`func (o *GetCatalogCategoryResponseCollectionDataInner) GetRelationships() GetCatalogCategoryResponseCollectionDataInnerAllOfRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GetCatalogCategoryResponseCollectionDataInner) GetRelationshipsOk() (*GetCatalogCategoryResponseCollectionDataInnerAllOfRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GetCatalogCategoryResponseCollectionDataInner) SetRelationships(v GetCatalogCategoryResponseCollectionDataInnerAllOfRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GetCatalogCategoryResponseCollectionDataInner) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


