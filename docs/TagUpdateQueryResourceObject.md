# TagUpdateQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**TagEnum**](TagEnum.md) |  | 
**Id** | **string** | The Tag ID | 
**Attributes** | [**TagResponseObjectResourceAttributes**](TagResponseObjectResourceAttributes.md) |  | 

## Methods

### NewTagUpdateQueryResourceObject

`func NewTagUpdateQueryResourceObject(type_ TagEnum, id string, attributes TagResponseObjectResourceAttributes, ) *TagUpdateQueryResourceObject`

NewTagUpdateQueryResourceObject instantiates a new TagUpdateQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagUpdateQueryResourceObjectWithDefaults

`func NewTagUpdateQueryResourceObjectWithDefaults() *TagUpdateQueryResourceObject`

NewTagUpdateQueryResourceObjectWithDefaults instantiates a new TagUpdateQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TagUpdateQueryResourceObject) GetType() TagEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TagUpdateQueryResourceObject) GetTypeOk() (*TagEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TagUpdateQueryResourceObject) SetType(v TagEnum)`

SetType sets Type field to given value.


### GetId

`func (o *TagUpdateQueryResourceObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagUpdateQueryResourceObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagUpdateQueryResourceObject) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *TagUpdateQueryResourceObject) GetAttributes() TagResponseObjectResourceAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TagUpdateQueryResourceObject) GetAttributesOk() (*TagResponseObjectResourceAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TagUpdateQueryResourceObject) SetAttributes(v TagResponseObjectResourceAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


