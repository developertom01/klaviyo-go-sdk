# ProfileSubscriptionCreateQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ProfileEnum**](ProfileEnum.md) |  | 
**Id** | Pointer to **NullableString** | The ID of the profile to subscribe. If provided, this will be used to perform the lookup. | [optional] 
**Attributes** | [**ProfileSubscriptionCreateQueryResourceObjectAttributes**](ProfileSubscriptionCreateQueryResourceObjectAttributes.md) |  | 

## Methods

### NewProfileSubscriptionCreateQueryResourceObject

`func NewProfileSubscriptionCreateQueryResourceObject(type_ ProfileEnum, attributes ProfileSubscriptionCreateQueryResourceObjectAttributes, ) *ProfileSubscriptionCreateQueryResourceObject`

NewProfileSubscriptionCreateQueryResourceObject instantiates a new ProfileSubscriptionCreateQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileSubscriptionCreateQueryResourceObjectWithDefaults

`func NewProfileSubscriptionCreateQueryResourceObjectWithDefaults() *ProfileSubscriptionCreateQueryResourceObject`

NewProfileSubscriptionCreateQueryResourceObjectWithDefaults instantiates a new ProfileSubscriptionCreateQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ProfileSubscriptionCreateQueryResourceObject) GetType() ProfileEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProfileSubscriptionCreateQueryResourceObject) GetTypeOk() (*ProfileEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProfileSubscriptionCreateQueryResourceObject) SetType(v ProfileEnum)`

SetType sets Type field to given value.


### GetId

`func (o *ProfileSubscriptionCreateQueryResourceObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProfileSubscriptionCreateQueryResourceObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProfileSubscriptionCreateQueryResourceObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProfileSubscriptionCreateQueryResourceObject) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ProfileSubscriptionCreateQueryResourceObject) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ProfileSubscriptionCreateQueryResourceObject) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *ProfileSubscriptionCreateQueryResourceObject) GetAttributes() ProfileSubscriptionCreateQueryResourceObjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ProfileSubscriptionCreateQueryResourceObject) GetAttributesOk() (*ProfileSubscriptionCreateQueryResourceObjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ProfileSubscriptionCreateQueryResourceObject) SetAttributes(v ProfileSubscriptionCreateQueryResourceObjectAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


