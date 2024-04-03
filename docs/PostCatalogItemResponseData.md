# PostCatalogItemResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CatalogItemEnum**](CatalogItemEnum.md) |  | 
**Id** | **string** | The catalog item ID is a compound ID (string), with format: &#x60;{integration}:::{catalog}:::{external_id}&#x60;. Currently, the only supported integration type is &#x60;$custom&#x60;, and the only supported catalog is &#x60;$default&#x60;. | 
**Attributes** | [**CatalogItemResponseObjectResourceAttributes**](CatalogItemResponseObjectResourceAttributes.md) |  | 
**Relationships** | Pointer to [**GetCatalogItemResponseCollectionCompoundDocumentDataInnerAllOfRelationships**](GetCatalogItemResponseCollectionCompoundDocumentDataInnerAllOfRelationships.md) |  | [optional] 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewPostCatalogItemResponseData

`func NewPostCatalogItemResponseData(type_ CatalogItemEnum, id string, attributes CatalogItemResponseObjectResourceAttributes, links ObjectLinks, ) *PostCatalogItemResponseData`

NewPostCatalogItemResponseData instantiates a new PostCatalogItemResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCatalogItemResponseDataWithDefaults

`func NewPostCatalogItemResponseDataWithDefaults() *PostCatalogItemResponseData`

NewPostCatalogItemResponseDataWithDefaults instantiates a new PostCatalogItemResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PostCatalogItemResponseData) GetType() CatalogItemEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostCatalogItemResponseData) GetTypeOk() (*CatalogItemEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostCatalogItemResponseData) SetType(v CatalogItemEnum)`

SetType sets Type field to given value.


### GetId

`func (o *PostCatalogItemResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostCatalogItemResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostCatalogItemResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *PostCatalogItemResponseData) GetAttributes() CatalogItemResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PostCatalogItemResponseData) GetAttributesOk() (*CatalogItemResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PostCatalogItemResponseData) SetAttributes(v CatalogItemResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PostCatalogItemResponseData) GetRelationships() GetCatalogItemResponseCollectionCompoundDocumentDataInnerAllOfRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PostCatalogItemResponseData) GetRelationshipsOk() (*GetCatalogItemResponseCollectionCompoundDocumentDataInnerAllOfRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PostCatalogItemResponseData) SetRelationships(v GetCatalogItemResponseCollectionCompoundDocumentDataInnerAllOfRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PostCatalogItemResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *PostCatalogItemResponseData) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PostCatalogItemResponseData) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PostCatalogItemResponseData) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


