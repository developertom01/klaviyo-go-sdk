# ProfileImportJobCreateQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ProfileBulkImportJobEnum**](ProfileBulkImportJobEnum.md) |  | 
**Attributes** | [**ProfileImportJobCreateQueryResourceObjectAttributes**](ProfileImportJobCreateQueryResourceObjectAttributes.md) |  | 
**Relationships** | Pointer to [**ProfileImportJobCreateQueryResourceObjectRelationships**](ProfileImportJobCreateQueryResourceObjectRelationships.md) |  | [optional] 

## Methods

### NewProfileImportJobCreateQueryResourceObject

`func NewProfileImportJobCreateQueryResourceObject(type_ ProfileBulkImportJobEnum, attributes ProfileImportJobCreateQueryResourceObjectAttributes, ) *ProfileImportJobCreateQueryResourceObject`

NewProfileImportJobCreateQueryResourceObject instantiates a new ProfileImportJobCreateQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileImportJobCreateQueryResourceObjectWithDefaults

`func NewProfileImportJobCreateQueryResourceObjectWithDefaults() *ProfileImportJobCreateQueryResourceObject`

NewProfileImportJobCreateQueryResourceObjectWithDefaults instantiates a new ProfileImportJobCreateQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ProfileImportJobCreateQueryResourceObject) GetType() ProfileBulkImportJobEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProfileImportJobCreateQueryResourceObject) GetTypeOk() (*ProfileBulkImportJobEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProfileImportJobCreateQueryResourceObject) SetType(v ProfileBulkImportJobEnum)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ProfileImportJobCreateQueryResourceObject) GetAttributes() ProfileImportJobCreateQueryResourceObjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ProfileImportJobCreateQueryResourceObject) GetAttributesOk() (*ProfileImportJobCreateQueryResourceObjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ProfileImportJobCreateQueryResourceObject) SetAttributes(v ProfileImportJobCreateQueryResourceObjectAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ProfileImportJobCreateQueryResourceObject) GetRelationships() ProfileImportJobCreateQueryResourceObjectRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ProfileImportJobCreateQueryResourceObject) GetRelationshipsOk() (*ProfileImportJobCreateQueryResourceObjectRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ProfileImportJobCreateQueryResourceObject) SetRelationships(v ProfileImportJobCreateQueryResourceObjectRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ProfileImportJobCreateQueryResourceObject) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


