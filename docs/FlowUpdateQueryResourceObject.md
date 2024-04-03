# FlowUpdateQueryResourceObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FlowEnum**](FlowEnum.md) |  | 
**Id** | **string** | ID of the Flow to update. Ex: XVTP5Q | 
**Attributes** | [**FlowUpdateQueryResourceObjectAttributes**](FlowUpdateQueryResourceObjectAttributes.md) |  | 

## Methods

### NewFlowUpdateQueryResourceObject

`func NewFlowUpdateQueryResourceObject(type_ FlowEnum, id string, attributes FlowUpdateQueryResourceObjectAttributes, ) *FlowUpdateQueryResourceObject`

NewFlowUpdateQueryResourceObject instantiates a new FlowUpdateQueryResourceObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowUpdateQueryResourceObjectWithDefaults

`func NewFlowUpdateQueryResourceObjectWithDefaults() *FlowUpdateQueryResourceObject`

NewFlowUpdateQueryResourceObjectWithDefaults instantiates a new FlowUpdateQueryResourceObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FlowUpdateQueryResourceObject) GetType() FlowEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlowUpdateQueryResourceObject) GetTypeOk() (*FlowEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlowUpdateQueryResourceObject) SetType(v FlowEnum)`

SetType sets Type field to given value.


### GetId

`func (o *FlowUpdateQueryResourceObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlowUpdateQueryResourceObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlowUpdateQueryResourceObject) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *FlowUpdateQueryResourceObject) GetAttributes() FlowUpdateQueryResourceObjectAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FlowUpdateQueryResourceObject) GetAttributesOk() (*FlowUpdateQueryResourceObjectAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FlowUpdateQueryResourceObject) SetAttributes(v FlowUpdateQueryResourceObjectAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


