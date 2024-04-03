# GetCatalogItemResponseCollectionCompoundDocumentDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CatalogItemEnum**](CatalogItemEnum.md) |  | 
**Id** | **string** | The catalog item ID is a compound ID (string), with format: &#x60;{integration}:::{catalog}:::{external_id}&#x60;. Currently, the only supported integration type is &#x60;$custom&#x60;, and the only supported catalog is &#x60;$default&#x60;. | 
**Attributes** | [**CatalogItemResponseObjectResourceAttributes**](CatalogItemResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 
**Relationships** | Pointer to [**GetCatalogItemResponseCollectionCompoundDocumentDataInnerAllOfRelationships**](GetCatalogItemResponseCollectionCompoundDocumentDataInnerAllOfRelationships.md) |  | [optional] 

## Methods

### NewGetCatalogItemResponseCollectionCompoundDocumentDataInner

`func NewGetCatalogItemResponseCollectionCompoundDocumentDataInner(type_ CatalogItemEnum, id string, attributes CatalogItemResponseObjectResourceAttributes, links ObjectLinks, ) *GetCatalogItemResponseCollectionCompoundDocumentDataInner`

NewGetCatalogItemResponseCollectionCompoundDocumentDataInner instantiates a new GetCatalogItemResponseCollectionCompoundDocumentDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCatalogItemResponseCollectionCompoundDocumentDataInnerWithDefaults

`func NewGetCatalogItemResponseCollectionCompoundDocumentDataInnerWithDefaults() *GetCatalogItemResponseCollectionCompoundDocumentDataInner`

NewGetCatalogItemResponseCollectionCompoundDocumentDataInnerWithDefaults instantiates a new GetCatalogItemResponseCollectionCompoundDocumentDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetCatalogItemResponseCollectionCompoundDocumentDataInner) GetType() CatalogItemEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCatalogItemResponseCollectionCompoundDocumentDataInner) GetTypeOk() (*CatalogItemEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCatalogItemResponseCollectionCompoundDocumentDataInner) SetType(v CatalogItemEnum)`

SetType sets Type field to given value.


### GetId

`func (o *GetCatalogItemResponseCollectionCompoundDocumentDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCatalogItemResponseCollectionCompoundDocumentDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCatalogItemResponseCollectionCompoundDocumentDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GetCatalogItemResponseCollectionCompoundDocumentDataInner) GetAttributes() CatalogItemResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetCatalogItemResponseCollectionCompoundDocumentDataInner) GetAttributesOk() (*CatalogItemResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetCatalogItemResponseCollectionCompoundDocumentDataInner) SetAttributes(v CatalogItemResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *GetCatalogItemResponseCollectionCompoundDocumentDataInner) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetCatalogItemResponseCollectionCompoundDocumentDataInner) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetCatalogItemResponseCollectionCompoundDocumentDataInner) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.


### GetRelationships

`func (o *GetCatalogItemResponseCollectionCompoundDocumentDataInner) GetRelationships() GetCatalogItemResponseCollectionCompoundDocumentDataInnerAllOfRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GetCatalogItemResponseCollectionCompoundDocumentDataInner) GetRelationshipsOk() (*GetCatalogItemResponseCollectionCompoundDocumentDataInnerAllOfRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GetCatalogItemResponseCollectionCompoundDocumentDataInner) SetRelationships(v GetCatalogItemResponseCollectionCompoundDocumentDataInnerAllOfRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GetCatalogItemResponseCollectionCompoundDocumentDataInner) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


