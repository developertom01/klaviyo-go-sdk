# GetTagGroupResponseCollectionDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**TagGroupEnum**](TagGroupEnum.md) |  | 
**Id** | **string** | The Tag Group ID | 
**Attributes** | [**TagGroupResponseObjectResourceAttributes**](TagGroupResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 
**Relationships** | Pointer to [**GetTagGroupResponseCollectionDataInnerAllOfRelationships**](GetTagGroupResponseCollectionDataInnerAllOfRelationships.md) |  | [optional] 

## Methods

### NewGetTagGroupResponseCollectionDataInner

`func NewGetTagGroupResponseCollectionDataInner(type_ TagGroupEnum, id string, attributes TagGroupResponseObjectResourceAttributes, links ObjectLinks, ) *GetTagGroupResponseCollectionDataInner`

NewGetTagGroupResponseCollectionDataInner instantiates a new GetTagGroupResponseCollectionDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTagGroupResponseCollectionDataInnerWithDefaults

`func NewGetTagGroupResponseCollectionDataInnerWithDefaults() *GetTagGroupResponseCollectionDataInner`

NewGetTagGroupResponseCollectionDataInnerWithDefaults instantiates a new GetTagGroupResponseCollectionDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetTagGroupResponseCollectionDataInner) GetType() TagGroupEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetTagGroupResponseCollectionDataInner) GetTypeOk() (*TagGroupEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetTagGroupResponseCollectionDataInner) SetType(v TagGroupEnum)`

SetType sets Type field to given value.


### GetId

`func (o *GetTagGroupResponseCollectionDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTagGroupResponseCollectionDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTagGroupResponseCollectionDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GetTagGroupResponseCollectionDataInner) GetAttributes() TagGroupResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetTagGroupResponseCollectionDataInner) GetAttributesOk() (*TagGroupResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetTagGroupResponseCollectionDataInner) SetAttributes(v TagGroupResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *GetTagGroupResponseCollectionDataInner) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetTagGroupResponseCollectionDataInner) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetTagGroupResponseCollectionDataInner) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.


### GetRelationships

`func (o *GetTagGroupResponseCollectionDataInner) GetRelationships() GetTagGroupResponseCollectionDataInnerAllOfRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GetTagGroupResponseCollectionDataInner) GetRelationshipsOk() (*GetTagGroupResponseCollectionDataInnerAllOfRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GetTagGroupResponseCollectionDataInner) SetRelationships(v GetTagGroupResponseCollectionDataInnerAllOfRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GetTagGroupResponseCollectionDataInner) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


