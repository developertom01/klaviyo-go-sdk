# SubscriptionCreateJobCreateQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ProfileSubscriptionBulkCreateJobEnum**](ProfileSubscriptionBulkCreateJobEnum.md) |  | 
**Attributes** | [**SubscriptionCreateJobCreateQueryResourceObjectAttributes**](SubscriptionCreateJobCreateQueryResourceObjectAttributes.md) |  | 
**Relationships** | Pointer to [**SubscriptionCreateJobCreateQueryResourceObjectRelationships**](SubscriptionCreateJobCreateQueryResourceObjectRelationships.md) |  | [optional] 

## Methods

### NewSubscriptionCreateJobCreateQueryResourceObject

`func NewSubscriptionCreateJobCreateQueryResourceObject(type_ ProfileSubscriptionBulkCreateJobEnum, attributes SubscriptionCreateJobCreateQueryResourceObjectAttributes, ) *SubscriptionCreateJobCreateQueryResourceObject`

NewSubscriptionCreateJobCreateQueryResourceObject instantiates a new SubscriptionCreateJobCreateQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionCreateJobCreateQueryResourceObjectWithDefaults

`func NewSubscriptionCreateJobCreateQueryResourceObjectWithDefaults() *SubscriptionCreateJobCreateQueryResourceObject`

NewSubscriptionCreateJobCreateQueryResourceObjectWithDefaults instantiates a new SubscriptionCreateJobCreateQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SubscriptionCreateJobCreateQueryResourceObject) GetType() ProfileSubscriptionBulkCreateJobEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubscriptionCreateJobCreateQueryResourceObject) GetTypeOk() (*ProfileSubscriptionBulkCreateJobEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubscriptionCreateJobCreateQueryResourceObject) SetType(v ProfileSubscriptionBulkCreateJobEnum)`

SetType sets Type field to given value.


### GetAttributes

`func (o *SubscriptionCreateJobCreateQueryResourceObject) GetAttributes() SubscriptionCreateJobCreateQueryResourceObjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SubscriptionCreateJobCreateQueryResourceObject) GetAttributesOk() (*SubscriptionCreateJobCreateQueryResourceObjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SubscriptionCreateJobCreateQueryResourceObject) SetAttributes(v SubscriptionCreateJobCreateQueryResourceObjectAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SubscriptionCreateJobCreateQueryResourceObject) GetRelationships() SubscriptionCreateJobCreateQueryResourceObjectRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SubscriptionCreateJobCreateQueryResourceObject) GetRelationshipsOk() (*SubscriptionCreateJobCreateQueryResourceObjectRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SubscriptionCreateJobCreateQueryResourceObject) SetRelationships(v SubscriptionCreateJobCreateQueryResourceObjectRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SubscriptionCreateJobCreateQueryResourceObject) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


