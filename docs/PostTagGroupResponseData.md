# PostTagGroupResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**TagGroupEnum**](TagGroupEnum.md) |  | 
**Id** | **string** | The Tag Group ID | 
**Attributes** | [**TagGroupResponseObjectResourceAttributes**](TagGroupResponseObjectResourceAttributes.md) |  | 
**Relationships** | Pointer to [**PostTagGroupResponseDataRelationships**](PostTagGroupResponseDataRelationships.md) |  | [optional] 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewPostTagGroupResponseData

`func NewPostTagGroupResponseData(type_ TagGroupEnum, id string, attributes TagGroupResponseObjectResourceAttributes, links ObjectLinks, ) *PostTagGroupResponseData`

NewPostTagGroupResponseData instantiates a new PostTagGroupResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTagGroupResponseDataWithDefaults

`func NewPostTagGroupResponseDataWithDefaults() *PostTagGroupResponseData`

NewPostTagGroupResponseDataWithDefaults instantiates a new PostTagGroupResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PostTagGroupResponseData) GetType() TagGroupEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostTagGroupResponseData) GetTypeOk() (*TagGroupEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostTagGroupResponseData) SetType(v TagGroupEnum)`

SetType sets Type field to given value.


### GetId

`func (o *PostTagGroupResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostTagGroupResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostTagGroupResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *PostTagGroupResponseData) GetAttributes() TagGroupResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PostTagGroupResponseData) GetAttributesOk() (*TagGroupResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PostTagGroupResponseData) SetAttributes(v TagGroupResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PostTagGroupResponseData) GetRelationships() PostTagGroupResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PostTagGroupResponseData) GetRelationshipsOk() (*PostTagGroupResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PostTagGroupResponseData) SetRelationships(v PostTagGroupResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PostTagGroupResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *PostTagGroupResponseData) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PostTagGroupResponseData) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PostTagGroupResponseData) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


