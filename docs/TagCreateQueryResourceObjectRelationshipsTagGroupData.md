# TagCreateQueryResourceObjectRelationshipsTagGroupData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**TagGroupEnum**](TagGroupEnum.md) |  | 
**Id** | **string** | The ID of the Tag Group to associate the Tag with. If this field is not specified, the Tag will be associated with the company&#39;s Default Tag Group. | 

## Methods

### NewTagCreateQueryResourceObjectRelationshipsTagGroupData

`func NewTagCreateQueryResourceObjectRelationshipsTagGroupData(type_ TagGroupEnum, id string, ) *TagCreateQueryResourceObjectRelationshipsTagGroupData`

NewTagCreateQueryResourceObjectRelationshipsTagGroupData instantiates a new TagCreateQueryResourceObjectRelationshipsTagGroupData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagCreateQueryResourceObjectRelationshipsTagGroupDataWithDefaults

`func NewTagCreateQueryResourceObjectRelationshipsTagGroupDataWithDefaults() *TagCreateQueryResourceObjectRelationshipsTagGroupData`

NewTagCreateQueryResourceObjectRelationshipsTagGroupDataWithDefaults instantiates a new TagCreateQueryResourceObjectRelationshipsTagGroupData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TagCreateQueryResourceObjectRelationshipsTagGroupData) GetType() TagGroupEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TagCreateQueryResourceObjectRelationshipsTagGroupData) GetTypeOk() (*TagGroupEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TagCreateQueryResourceObjectRelationshipsTagGroupData) SetType(v TagGroupEnum)`

SetType sets Type field to given value.


### GetId

`func (o *TagCreateQueryResourceObjectRelationshipsTagGroupData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagCreateQueryResourceObjectRelationshipsTagGroupData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagCreateQueryResourceObjectRelationshipsTagGroupData) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


