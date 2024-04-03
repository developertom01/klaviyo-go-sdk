# SubscriptionDeleteJobCreateQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ProfileSubscriptionBulkDeleteJobEnum**](ProfileSubscriptionBulkDeleteJobEnum.md) |  | 
**Attributes** | [**SubscriptionDeleteJobCreateQueryResourceObjectAttributes**](SubscriptionDeleteJobCreateQueryResourceObjectAttributes.md) |  | 
**Relationships** | Pointer to [**SubscriptionDeleteJobCreateQueryResourceObjectRelationships**](SubscriptionDeleteJobCreateQueryResourceObjectRelationships.md) |  | [optional] 

## Methods

### NewSubscriptionDeleteJobCreateQueryResourceObject

`func NewSubscriptionDeleteJobCreateQueryResourceObject(type_ ProfileSubscriptionBulkDeleteJobEnum, attributes SubscriptionDeleteJobCreateQueryResourceObjectAttributes, ) *SubscriptionDeleteJobCreateQueryResourceObject`

NewSubscriptionDeleteJobCreateQueryResourceObject instantiates a new SubscriptionDeleteJobCreateQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionDeleteJobCreateQueryResourceObjectWithDefaults

`func NewSubscriptionDeleteJobCreateQueryResourceObjectWithDefaults() *SubscriptionDeleteJobCreateQueryResourceObject`

NewSubscriptionDeleteJobCreateQueryResourceObjectWithDefaults instantiates a new SubscriptionDeleteJobCreateQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SubscriptionDeleteJobCreateQueryResourceObject) GetType() ProfileSubscriptionBulkDeleteJobEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubscriptionDeleteJobCreateQueryResourceObject) GetTypeOk() (*ProfileSubscriptionBulkDeleteJobEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubscriptionDeleteJobCreateQueryResourceObject) SetType(v ProfileSubscriptionBulkDeleteJobEnum)`

SetType sets Type field to given value.


### GetAttributes

`func (o *SubscriptionDeleteJobCreateQueryResourceObject) GetAttributes() SubscriptionDeleteJobCreateQueryResourceObjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SubscriptionDeleteJobCreateQueryResourceObject) GetAttributesOk() (*SubscriptionDeleteJobCreateQueryResourceObjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SubscriptionDeleteJobCreateQueryResourceObject) SetAttributes(v SubscriptionDeleteJobCreateQueryResourceObjectAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SubscriptionDeleteJobCreateQueryResourceObject) GetRelationships() SubscriptionDeleteJobCreateQueryResourceObjectRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SubscriptionDeleteJobCreateQueryResourceObject) GetRelationshipsOk() (*SubscriptionDeleteJobCreateQueryResourceObjectRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SubscriptionDeleteJobCreateQueryResourceObject) SetRelationships(v SubscriptionDeleteJobCreateQueryResourceObjectRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SubscriptionDeleteJobCreateQueryResourceObject) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


