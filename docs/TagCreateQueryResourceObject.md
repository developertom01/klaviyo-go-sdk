# TagCreateQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**TagEnum**](TagEnum.md) |  | 
**Attributes** | [**TagResponseObjectResourceAttributes**](TagResponseObjectResourceAttributes.md) |  | 
**Relationships** | Pointer to [**TagCreateQueryResourceObjectRelationships**](TagCreateQueryResourceObjectRelationships.md) |  | [optional] 

## Methods

### NewTagCreateQueryResourceObject

`func NewTagCreateQueryResourceObject(type_ TagEnum, attributes TagResponseObjectResourceAttributes, ) *TagCreateQueryResourceObject`

NewTagCreateQueryResourceObject instantiates a new TagCreateQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagCreateQueryResourceObjectWithDefaults

`func NewTagCreateQueryResourceObjectWithDefaults() *TagCreateQueryResourceObject`

NewTagCreateQueryResourceObjectWithDefaults instantiates a new TagCreateQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TagCreateQueryResourceObject) GetType() TagEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TagCreateQueryResourceObject) GetTypeOk() (*TagEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TagCreateQueryResourceObject) SetType(v TagEnum)`

SetType sets Type field to given value.


### GetAttributes

`func (o *TagCreateQueryResourceObject) GetAttributes() TagResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TagCreateQueryResourceObject) GetAttributesOk() (*TagResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TagCreateQueryResourceObject) SetAttributes(v TagResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *TagCreateQueryResourceObject) GetRelationships() TagCreateQueryResourceObjectRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TagCreateQueryResourceObject) GetRelationshipsOk() (*TagCreateQueryResourceObjectRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TagCreateQueryResourceObject) SetRelationships(v TagCreateQueryResourceObjectRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TagCreateQueryResourceObject) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


