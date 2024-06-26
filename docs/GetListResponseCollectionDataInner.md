# GetListResponseCollectionDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ListEnum**](ListEnum.md) |  | 
**Id** | **string** | Primary key that uniquely identifies this list. Generated by Klaviyo. | 
**Attributes** | [**ListListResponseObjectResourceAttributes**](ListListResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 
**Relationships** | Pointer to [**GetListResponseCollectionDataInnerAllOfRelationships**](GetListResponseCollectionDataInnerAllOfRelationships.md) |  | [optional] 

## Methods

### NewGetListResponseCollectionDataInner

`func NewGetListResponseCollectionDataInner(type_ ListEnum, id string, attributes ListListResponseObjectResourceAttributes, links ObjectLinks, ) *GetListResponseCollectionDataInner`

NewGetListResponseCollectionDataInner instantiates a new GetListResponseCollectionDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetListResponseCollectionDataInnerWithDefaults

`func NewGetListResponseCollectionDataInnerWithDefaults() *GetListResponseCollectionDataInner`

NewGetListResponseCollectionDataInnerWithDefaults instantiates a new GetListResponseCollectionDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetListResponseCollectionDataInner) GetType() ListEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetListResponseCollectionDataInner) GetTypeOk() (*ListEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetListResponseCollectionDataInner) SetType(v ListEnum)`

SetType sets Type field to given value.


### GetId

`func (o *GetListResponseCollectionDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetListResponseCollectionDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetListResponseCollectionDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GetListResponseCollectionDataInner) GetAttributes() ListListResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetListResponseCollectionDataInner) GetAttributesOk() (*ListListResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetListResponseCollectionDataInner) SetAttributes(v ListListResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *GetListResponseCollectionDataInner) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetListResponseCollectionDataInner) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetListResponseCollectionDataInner) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.


### GetRelationships

`func (o *GetListResponseCollectionDataInner) GetRelationships() GetListResponseCollectionDataInnerAllOfRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GetListResponseCollectionDataInner) GetRelationshipsOk() (*GetListResponseCollectionDataInnerAllOfRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GetListResponseCollectionDataInner) SetRelationships(v GetListResponseCollectionDataInnerAllOfRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GetListResponseCollectionDataInner) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


