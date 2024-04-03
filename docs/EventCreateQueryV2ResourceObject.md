# EventCreateQueryV2ResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EventEnum**](EventEnum.md) |  | 
**Attributes** | [**EventCreateQueryV2ResourceObjectAttributes**](EventCreateQueryV2ResourceObjectAttributes.md) |  | 

## Methods

### NewEventCreateQueryV2ResourceObject

`func NewEventCreateQueryV2ResourceObject(type_ EventEnum, attributes EventCreateQueryV2ResourceObjectAttributes, ) *EventCreateQueryV2ResourceObject`

NewEventCreateQueryV2ResourceObject instantiates a new EventCreateQueryV2ResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventCreateQueryV2ResourceObjectWithDefaults

`func NewEventCreateQueryV2ResourceObjectWithDefaults() *EventCreateQueryV2ResourceObject`

NewEventCreateQueryV2ResourceObjectWithDefaults instantiates a new EventCreateQueryV2ResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EventCreateQueryV2ResourceObject) GetType() EventEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventCreateQueryV2ResourceObject) GetTypeOk() (*EventEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventCreateQueryV2ResourceObject) SetType(v EventEnum)`

SetType sets Type field to given value.


### GetAttributes

`func (o *EventCreateQueryV2ResourceObject) GetAttributes() EventCreateQueryV2ResourceObjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EventCreateQueryV2ResourceObject) GetAttributesOk() (*EventCreateQueryV2ResourceObjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EventCreateQueryV2ResourceObject) SetAttributes(v EventCreateQueryV2ResourceObjectAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


