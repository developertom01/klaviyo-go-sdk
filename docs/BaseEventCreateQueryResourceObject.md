# BaseEventCreateQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EventEnum**](EventEnum.md) |  | 
**Attributes** | [**BaseEventCreateQueryResourceObjectAttributes**](BaseEventCreateQueryResourceObjectAttributes.md) |  | 

## Methods

### NewBaseEventCreateQueryResourceObject

`func NewBaseEventCreateQueryResourceObject(type_ EventEnum, attributes BaseEventCreateQueryResourceObjectAttributes, ) *BaseEventCreateQueryResourceObject`

NewBaseEventCreateQueryResourceObject instantiates a new BaseEventCreateQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEventCreateQueryResourceObjectWithDefaults

`func NewBaseEventCreateQueryResourceObjectWithDefaults() *BaseEventCreateQueryResourceObject`

NewBaseEventCreateQueryResourceObjectWithDefaults instantiates a new BaseEventCreateQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BaseEventCreateQueryResourceObject) GetType() EventEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseEventCreateQueryResourceObject) GetTypeOk() (*EventEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BaseEventCreateQueryResourceObject) SetType(v EventEnum)`

SetType sets Type field to given value.


### GetAttributes

`func (o *BaseEventCreateQueryResourceObject) GetAttributes() BaseEventCreateQueryResourceObjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BaseEventCreateQueryResourceObject) GetAttributesOk() (*BaseEventCreateQueryResourceObjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BaseEventCreateQueryResourceObject) SetAttributes(v BaseEventCreateQueryResourceObjectAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


