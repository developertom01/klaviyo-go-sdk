# TagGroupResponseObjectResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**TagGroupEnum**](TagGroupEnum.md) |  | 
**Id** | **string** | The Tag Group ID | 
**Attributes** | [**TagGroupResponseObjectResourceAttributes**](TagGroupResponseObjectResourceAttributes.md) |  | 
**Links** | [**ObjectLinks**](ObjectLinks.md) |  | 

## Methods

### NewTagGroupResponseObjectResource

`func NewTagGroupResponseObjectResource(type_ TagGroupEnum, id string, attributes TagGroupResponseObjectResourceAttributes, links ObjectLinks, ) *TagGroupResponseObjectResource`

NewTagGroupResponseObjectResource instantiates a new TagGroupResponseObjectResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagGroupResponseObjectResourceWithDefaults

`func NewTagGroupResponseObjectResourceWithDefaults() *TagGroupResponseObjectResource`

NewTagGroupResponseObjectResourceWithDefaults instantiates a new TagGroupResponseObjectResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TagGroupResponseObjectResource) GetType() TagGroupEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TagGroupResponseObjectResource) GetTypeOk() (*TagGroupEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TagGroupResponseObjectResource) SetType(v TagGroupEnum)`

SetType sets Type field to given value.


### GetId

`func (o *TagGroupResponseObjectResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagGroupResponseObjectResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagGroupResponseObjectResource) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *TagGroupResponseObjectResource) GetAttributes() TagGroupResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TagGroupResponseObjectResource) GetAttributesOk() (*TagGroupResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TagGroupResponseObjectResource) SetAttributes(v TagGroupResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetLinks

`func (o *TagGroupResponseObjectResource) GetLinks() ObjectLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TagGroupResponseObjectResource) GetLinksOk() (*ObjectLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TagGroupResponseObjectResource) SetLinks(v ObjectLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


