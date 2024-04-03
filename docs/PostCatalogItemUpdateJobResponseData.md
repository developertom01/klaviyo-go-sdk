# PostCatalogItemUpdateJobResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CatalogItemBulkUpdateJobEnum**](CatalogItemBulkUpdateJobEnum.md) |  | 
**Id** | **string** | Unique identifier for retrieving the job. Generated by Klaviyo. | 
**Attributes** | [**CouponCodeCreateJobResponseObjectResourceAttributes**](CouponCodeCreateJobResponseObjectResourceAttributes.md) |  | 
**Relationships** | Pointer to [**GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships**](GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships.md) |  | [optional] 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewPostCatalogItemUpdateJobResponseData

`func NewPostCatalogItemUpdateJobResponseData(type_ CatalogItemBulkUpdateJobEnum, id string, attributes CouponCodeCreateJobResponseObjectResourceAttributes, links ObjectLinks, ) *PostCatalogItemUpdateJobResponseData`

NewPostCatalogItemUpdateJobResponseData instantiates a new PostCatalogItemUpdateJobResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCatalogItemUpdateJobResponseDataWithDefaults

`func NewPostCatalogItemUpdateJobResponseDataWithDefaults() *PostCatalogItemUpdateJobResponseData`

NewPostCatalogItemUpdateJobResponseDataWithDefaults instantiates a new PostCatalogItemUpdateJobResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PostCatalogItemUpdateJobResponseData) GetType() CatalogItemBulkUpdateJobEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostCatalogItemUpdateJobResponseData) GetTypeOk() (*CatalogItemBulkUpdateJobEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostCatalogItemUpdateJobResponseData) SetType(v CatalogItemBulkUpdateJobEnum)`

SetType sets Type field to given value.


### GetId

`func (o *PostCatalogItemUpdateJobResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostCatalogItemUpdateJobResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostCatalogItemUpdateJobResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *PostCatalogItemUpdateJobResponseData) GetAttributes() CouponCodeCreateJobResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PostCatalogItemUpdateJobResponseData) GetAttributesOk() (*CouponCodeCreateJobResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PostCatalogItemUpdateJobResponseData) SetAttributes(v CouponCodeCreateJobResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PostCatalogItemUpdateJobResponseData) GetRelationships() GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PostCatalogItemUpdateJobResponseData) GetRelationshipsOk() (*GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PostCatalogItemUpdateJobResponseData) SetRelationships(v GetCatalogItemUpdateJobResponseCollectionCompoundDocumentDataInnerAllOfRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PostCatalogItemUpdateJobResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *PostCatalogItemUpdateJobResponseData) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PostCatalogItemUpdateJobResponseData) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PostCatalogItemUpdateJobResponseData) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

