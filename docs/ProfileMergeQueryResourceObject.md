# ProfileMergeQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ProfileMergeEnum**](ProfileMergeEnum.md) |  | 
**Id** | **string** | The ID of the destination profile to merge into | 
**Relationships** | [**ProfileMergeQueryResourceObjectRelationships**](ProfileMergeQueryResourceObjectRelationships.md) |  | 

## Methods

### NewProfileMergeQueryResourceObject

`func NewProfileMergeQueryResourceObject(type_ ProfileMergeEnum, id string, relationships ProfileMergeQueryResourceObjectRelationships, ) *ProfileMergeQueryResourceObject`

NewProfileMergeQueryResourceObject instantiates a new ProfileMergeQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileMergeQueryResourceObjectWithDefaults

`func NewProfileMergeQueryResourceObjectWithDefaults() *ProfileMergeQueryResourceObject`

NewProfileMergeQueryResourceObjectWithDefaults instantiates a new ProfileMergeQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ProfileMergeQueryResourceObject) GetType() ProfileMergeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProfileMergeQueryResourceObject) GetTypeOk() (*ProfileMergeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProfileMergeQueryResourceObject) SetType(v ProfileMergeEnum)`

SetType sets Type field to given value.


### GetId

`func (o *ProfileMergeQueryResourceObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProfileMergeQueryResourceObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProfileMergeQueryResourceObject) SetId(v string)`

SetId sets Id field to given value.


### GetRelationships

`func (o *ProfileMergeQueryResourceObject) GetRelationships() ProfileMergeQueryResourceObjectRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ProfileMergeQueryResourceObject) GetRelationshipsOk() (*ProfileMergeQueryResourceObjectRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ProfileMergeQueryResourceObject) SetRelationships(v ProfileMergeQueryResourceObjectRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


